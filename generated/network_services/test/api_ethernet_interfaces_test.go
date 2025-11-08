/*
 * Network Services Testing
 *
 * EthernetInterfacesAPIService
 */

package network_services

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Replace with your actual imports
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// NOTE: Placeholder functions and types are assumed to be available in the test environment.
// - SetupNetworkSvcTestClient(t) -> returns a *network_services.APIClient (with EthernetInterfacesAPIService)
// - handleAPIError(err)
// - common.GenerateRandomString(6)
// - common.StringPtr(s) -> returns *string

// --- Helper Functions ---

// generateEthernetInterfaceName creates a unique name for the resource.
func generateEthernetInterfaceName(base string) string {
	return "$" + base + common.GenerateRandomString(4)
}

// createBaseEthernetInterface creates a base EthernetInterface object WITHOUT providing an Id.
// The API is expected to generate the Id on creation.
func createBaseEthernetInterface(t *testing.T, baseName string) network_services.EthernetInterfaces {
	name := generateEthernetInterfaceName(baseName)

	// Use the constructor with defaults, then manually set the required fields ('Name').
	// The 'Id' will be left as its zero value ("") and should be ignored/overwritten by the server.
	intf := *network_services.NewEthernetInterfacesWithDefaults()

	// Set the required 'Name' field
	intf.SetName(name)

	// Add common optional fields
	intf.SetComment("Managed by Go Test")
	intf.SetFolder("All")

	// Set link settings
	intf.SetLinkDuplex("full")
	intf.SetLinkSpeed("auto")
	intf.SetLinkState("up")

	return intf
}

// --- Test Cases for different 'Create' modes ---

// Test_CreateEthernetInterfaces_L2 tests creation of a Layer 2 Ethernet Interface.
func Test_CreateEthernetInterfaces_L2(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseEthernetInterface(t, "l2-intf-")

	// Set Layer 2 mode
	intf.SetLayer2(*network_services.NewEthernetInterfacesLayer2WithDefaults())

	res, httpRes, err := client.EthernetInterfacesAPI.
		CreateEthernetInterfaces(context.Background()).
		EthernetInterfaces(intf).
		Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create L2 Ethernet Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	assert.NotEmpty(t, res.GetId(), "The server must return a generated Id")

	// Cleanup the created resource using the generated Id
	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), res.GetId()).Execute()
	}()

	layer2Val, ok := res.GetLayer2Ok()
	require.True(t, ok, "Layer2 field must be set (ok boolean must be true)")
	assert.NotNil(t, layer2Val, "Layer2 configuration value should not be nil")
	assert.False(t, res.HasLayer3(), "Layer3 field must NOT be set")
}

// Test_CreateEthernetInterfaces_L3Static tests creation of a Layer 3 Ethernet Interface with static IP.
func Test_CreateEthernetInterfaces_L3Static(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseEthernetInterface(t, "l3-static-")

	// Set Layer 3 mode with static IP configuration
	layer3Config := network_services.NewEthernetInterfacesLayer3WithDefaults()
	layer3Config.SetIp([]network_services.EthernetInterfacesLayer3IpInner{
		{
			// Assuming the Layer3 IP inner structure has a 'Name' field for IP/Mask
			Name: "198.18.1.1/24",
		},
	})

	intf.SetLayer3(*layer3Config)

	res, httpRes, err := client.EthernetInterfacesAPI.
		CreateEthernetInterfaces(context.Background()).
		EthernetInterfaces(intf).
		Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create L3 Static Ethernet Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	assert.NotEmpty(t, res.GetId(), "The server must return a generated Id")

	// Cleanup the created resource
	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), res.GetId()).Execute()
	}()

	layer3Val, ok := res.GetLayer3Ok()
	require.True(t, ok, "Layer3 field must be explicitly set (ok boolean must be true)")
	assert.NotNil(t, layer3Val, "Layer3 configuration pointer should not be nil")
}

// Test_CreateEthernetInterfaces_L3DHCP tests creation of a Layer 3 Ethernet Interface with DHCP client.
// Test_CreateEthernetInterfaces_L3DHCP tests creation of a Layer 3 Ethernet Interface with DHCP client.
func Test_CreateEthernetInterfaces_L3DHCP(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseEthernetInterface(t, "l3-dhcp-") // Base interface setup (no ID, just Name, Comment, etc.)

	// 1. Create the Layer 3 container object
	layer3Config := network_services.NewEthernetInterfacesLayer3WithDefaults()

	// 2. Create the DHCP client configuration.
	// The model provided shows NewEthernetInterfacesLayer3DhcpClient() takes no arguments
	// but sets 'Enable', 'CreateDefaultRoute', and 'DefaultRouteMetric' to true/defaults.
	dhcpClientConfig := network_services.NewEthernetInterfacesLayer3DhcpClient()

	// Optional: Override defaults if needed.
	// dhcpClientConfig.SetCreateDefaultRoute(false)

	// 3. Set the DHCP client config onto the Layer 3 config
	layer3Config.SetDhcpClient(*dhcpClientConfig)

	// 4. Set the Layer 3 config onto the main interface object
	intf.SetLayer3(*layer3Config)

	// Execute the creation request
	res, httpRes, err := client.EthernetInterfacesAPI.
		CreateEthernetInterfaces(context.Background()).
		EthernetInterfaces(intf).
		Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create L3 DHCP Ethernet Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	assert.NotEmpty(t, res.GetId(), "The server must return a generated Id")

	// Cleanup the created resource
	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), res.GetId()).Execute()
	}()

	// Assertion: Verify that the Layer 3 object was created and contains the DHCP config
	layer3Val, ok := res.GetLayer3Ok()
	require.True(t, ok, "Layer3 field must be explicitly set (ok boolean)")
	require.NotNil(t, layer3Val, "Layer3 configuration pointer should not be nil")

	// Assertion: Verify that the DhcpClient field is set within Layer 3
	dhcpClientVal, dhcpOk := layer3Val.GetDhcpClientOk()
	require.True(t, dhcpOk, "DhcpClient must be explicitly set")
	require.NotNil(t, dhcpClientVal, "DhcpClient configuration pointer should not be nil")

	// Assertion: Verify a specific default was applied
	require.True(t, dhcpClientVal.GetEnable(), "DHCP client must be enabled")
}

// Test_CreateEthernetInterfaces_L3PPPoE tests creation of a Layer 3 Ethernet Interface with PPPoE.
func Test_CreateEthernetInterfaces_L3PPPoE(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseEthernetInterface(t, "l3-pppoe-")

	// Set Layer 3 mode with PPPoE configuration
	layer3Config := network_services.NewEthernetInterfacesLayer3WithDefaults()

	// Assuming NewEthernetInterfacesPppoeWithDefaults exists and sets required fields
	pppoeConfig := network_services.NewEthernetInterfacesLayer3PppoeWithDefaults()
	pppoeConfig.SetEnable(true)
	pppoeConfig.SetUsername("testuser")
	pppoeConfig.SetPassword("testpass")

	layer3Config.SetPppoe(*pppoeConfig)

	intf.SetLayer3(*layer3Config)

	res, httpRes, err := client.EthernetInterfacesAPI.
		CreateEthernetInterfaces(context.Background()).
		EthernetInterfaces(intf).
		Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create L3 PPPoE Ethernet Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	assert.NotEmpty(t, res.GetId(), "The server must return a generated Id")

	// Cleanup the created resource
	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), res.GetId()).Execute()
	}()

	// --- CORRECTION 1: Check Layer3 field ---
	layer3Val, layer3Ok := res.GetLayer3Ok()
	require.True(t, layer3Ok, "Layer3 field must be explicitly set (ok boolean)")
	require.NotNil(t, layer3Val, "Layer3 configuration pointer should not be nil")

	// --- CORRECTION 2: Check PPPoE field ---
	pppoeVal, pppoeOk := layer3Val.GetPppoeOk()
	require.True(t, pppoeOk, "PPPoE field must be explicitly set (ok boolean)")
	require.NotNil(t, pppoeVal, "PPPoE configuration pointer should not be nil")

	// Verify PPPoE specific config
	assert.True(t, pppoeVal.GetEnable(), "PPPoE enable flag should be true")
	assert.Equal(t, "testuser", pppoeVal.GetUsername(), "PPPoE username must match setup")
}

// Test_CreateEthernetInterfaces_Tap tests creation of a TAP Ethernet Interface.
func Test_CreateEthernetInterfaces_Tap(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseEthernetInterface(t, "tap-intf-")

	// Set TAP mode
	intf.SetTap(make(map[string]interface{}))

	res, httpRes, err := client.EthernetInterfacesAPI.
		CreateEthernetInterfaces(context.Background()).
		EthernetInterfaces(intf).
		Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create TAP Ethernet Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	assert.NotEmpty(t, res.GetId(), "The server must return a generated Id")

	// Cleanup the created resource
	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), res.GetId()).Execute()
	}()

	assert.True(t, res.HasTap(), "Tap field must be set")
	assert.False(t, res.HasLayer2(), "Layer2 field must NOT be set")
}

// Test_EthernetInterfacesAPIService_GetByID tests retrieving an Ethernet Interface by ID.
func Test_EthernetInterfacesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseEthernetInterface(t, "get-intf-")
	intf.SetComment("Original Comment for Get Test")

	// Set L2 mode for a simple setup
	intf.SetLayer2(*network_services.NewEthernetInterfacesLayer2WithDefaults())

	// Setup: Create an interface first to get the server-generated ID
	createRes, _, err := client.EthernetInterfacesAPI.CreateEthernetInterfaces(context.Background()).EthernetInterfaces(intf).Execute()
	require.NoError(t, err, "Failed to create interface for get test setup")
	createdID := createRes.GetId()

	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the interface
	getRes, httpResGet, errGet := client.EthernetInterfacesAPI.GetEthernetInterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get Ethernet Interface by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	assert.Equal(t, intf.GetName(), getRes.GetName(), "Interface name should match")
	assert.Equal(t, "Original Comment for Get Test", getRes.GetComment(), "Comment should be preserved")
}

// Test_EthernetInterfacesAPIService_Update tests updating an Ethernet Interface.
func Test_EthernetInterfacesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseEthernetInterface(t, "update-intf-")
	intf.SetLinkSpeed("auto")

	// Set L2 mode for a stable update base
	intf.SetLayer2(*network_services.NewEthernetInterfacesLayer2WithDefaults())

	// Setup: Create an interface first to get the server-generated ID
	createRes, _, err := client.EthernetInterfacesAPI.CreateEthernetInterfaces(context.Background()).EthernetInterfaces(intf).Execute()
	require.NoError(t, err, "Failed to create interface for update test setup")
	createdID := createRes.GetId()

	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), createdID).Execute()
	}()

	updatedIntf := *createRes

	updatedComment := "This interface was updated."
	updatedIntf.SetComment(updatedComment)

	// Test: Update the interface
	updateRes, httpResUpdate, errUpdate := client.EthernetInterfacesAPI.
		UpdateEthernetInterfacesByID(context.Background(), createdID).
		EthernetInterfaces(updatedIntf).
		Execute()

	require.NoError(t, errUpdate, "Failed to update Ethernet Interface")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Verify the update
	assert.Equal(t, updatedComment, updateRes.GetComment(), "Comment should be updated")
}

// Test_EthernetInterfacesAPIService_DeleteByID tests deleting an Ethernet Interface.
func Test_EthernetInterfacesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseEthernetInterface(t, "delete-intf-")
	intf.SetLayer2(*network_services.NewEthernetInterfacesLayer2WithDefaults())

	// Setup: Create an interface first
	createRes, _, err := client.EthernetInterfacesAPI.CreateEthernetInterfaces(context.Background()).EthernetInterfaces(intf).Execute()
	require.NoError(t, err, "Failed to create interface for delete test setup")
	createdID := createRes.GetId()

	// Note: No defer needed here, as the test is meant to perform the deletion.

	// Test: Delete the interface
	httpResDel, errDel := client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete Ethernet Interface")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}

// Test_EthernetInterfacesAPIService_List tests listing Ethernet Interfaces.
func Test_EthernetInterfacesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseEthernetInterface(t, "list-intf-")
	intf.SetLayer2(*network_services.NewEthernetInterfacesLayer2WithDefaults())

	// Setup: Create an interface first
	createRes, _, err := client.EthernetInterfacesAPI.CreateEthernetInterfaces(context.Background()).EthernetInterfaces(intf).Execute()
	require.NoError(t, err, "Failed to create interface for list test setup")
	createdID := createRes.GetId()

	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: List the interfaces, filtering by folder to reduce noise
	listRes, httpResList, errList := client.EthernetInterfacesAPI.ListEthernetInterfaces(context.Background()).
		Folder(intf.GetFolder()).
		Limit(10).
		Execute()

	require.NoError(t, errList, "Failed to list Ethernet Interfaces")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}
