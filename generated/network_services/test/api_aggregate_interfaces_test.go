package network_services

/*
 * Network Services Testing
 *
 * AggregateInterfacesAPIService
 */

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

// --- Helper Functions ---

// generateAggregateName creates a unique name for the resource.
func generateAggregateName(base string) string {
	// Name must match: ^\$[a-zA-Z\d\-_.]+
	return "$" + base + common.GenerateRandomString(4)
}

// createBaseAggregateInterface creates a base AggregateInterfaces object WITHOUT providing an Id.
func createBaseAggregateInterface(t *testing.T, baseName string) network_services.AggregateInterfaces {
	name := generateAggregateName(baseName)
	intf := *network_services.NewAggregateInterfaces(name) // 'name' is required

	// Add common optional fields
	intf.SetComment("Managed by Go Test - Aggregate")
	intf.SetFolder("All")

	return intf
}

// ---------------------------------------------------------------------------------------------------------------------
// --- Test Cases for different 'Create' modes (L2 vs. L3 exclusivity) ---
// ---------------------------------------------------------------------------------------------------------------------

// Test_CreateAggregateInterfaces_L2 tests creation of a Layer 2 Aggregate Interface.
func Test_CreateAggregateInterfaces_L2(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseAggregateInterface(t, "ae-l2-")

	// Set Layer 2 mode with LACP config (from your L2 LACP test data)
	layer2Config := network_services.NewAggregateInterfacesLayer2()
	layer2Config.SetLacp(*network_services.NewLacpWithDefaults()) // Assuming Lacp constructor exists
	layer2Config.SetVlanTag("6666")

	intf.SetLayer2(*layer2Config)

	res, httpRes, err := client.AggregateInterfacesAPI.
		CreateAggregateInterfaces(context.Background()).
		AggregateInterfaces(intf).
		Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create L2 Aggregate Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	assert.NotEmpty(t, res.GetId(), "The server must return a generated Id")
	createdID := res.GetId()

	// Cleanup the created resource
	defer func() {
		client.AggregateInterfacesAPI.DeleteAggregateInterfacesByID(context.Background(), createdID).Execute()
	}()

	// *** Exclusivity Assertion ***
	layer2Val, ok := res.GetLayer2Ok()
	require.True(t, ok, "**Layer2 field must be set**")
	assert.NotNil(t, layer2Val, "Layer2 configuration value should not be nil")
	assert.False(t, res.HasLayer3(), "**Layer3 field must NOT be set**")

	// L2 Specific Assertion
	assert.Equal(t, "6666", layer2Val.GetVlanTag(), "VLAN tag must be set correctly")
}

// Test_CreateAggregateInterfaces_L3Static tests creation of a Layer 3 Aggregate Interface with static IP.
func Test_CreateAggregateInterfaces_L3Static(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseAggregateInterface(t, "ae-l3-static-")

	// Set Layer 3 mode with static IP configuration (from your L3 static test data)
	layer3Config := network_services.NewAggregateInterfacesLayer3()
	layer3Config.SetIp([]network_services.AggregateInterfacesLayer3IpInner{
		{
			Name: "198.18.1.1/24",
		},
	})

	// Add LACP to the L3 configuration
	lacpConfig := network_services.NewLacpWithDefaults()
	lacpConfig.SetEnable(true)
	layer3Config.SetLacp(*lacpConfig)

	intf.SetLayer3(*layer3Config)

	res, httpRes, err := client.AggregateInterfacesAPI.
		CreateAggregateInterfaces(context.Background()).
		AggregateInterfaces(intf).
		Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create L3 Static Aggregate Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	createdID := res.GetId()

	// Cleanup the created resource
	defer func() {
		client.AggregateInterfacesAPI.DeleteAggregateInterfacesByID(context.Background(), createdID).Execute()
	}()

	// *** Exclusivity Assertion ***
	layer3Val, ok := res.GetLayer3Ok()
	require.True(t, ok, "**Layer3 field must be set**")
	assert.False(t, res.HasLayer2(), "**Layer2 field must NOT be set**")

	// L3 Static Specific Assertion
	require.NotEmpty(t, layer3Val.GetIp(), "Static IP list must be set")
	assert.Equal(t, "198.18.1.1/24", layer3Val.GetIp()[0].GetName(), "IP address must match")
	assert.True(t, layer3Val.HasLacp(), "LACP config must be present")

	// *** Inner Exclusivity Check ***
	assert.False(t, layer3Val.HasDhcpClient(), "DHCP client must NOT be set when static IP is configured")
}

// Test_CreateAggregateInterfaces_L3DHCP tests creation of a Layer 3 Aggregate Interface with DHCP client.
func Test_CreateAggregateInterfaces_L3DHCP(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseAggregateInterface(t, "ae-l3-dhcp-")

	// Set Layer 3 mode with DHCP client configuration (from your L3 DHCP test data)
	layer3Config := network_services.NewAggregateInterfacesLayer3()

	dhcpClientConfig := network_services.NewAggEthernetDhcpClientDhcpClient() // Assuming a basic constructor
	dhcpClientConfig.SetEnable(true)
	dhcpClientConfig.SetCreateDefaultRoute(true)
	dhcpClientConfig.SetDefaultRouteMetric(int32(10))

	layer3Config.SetDhcpClient(*dhcpClientConfig)
	intf.SetLayer3(*layer3Config)

	res, httpRes, err := client.AggregateInterfacesAPI.
		CreateAggregateInterfaces(context.Background()).
		AggregateInterfaces(intf).
		Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create L3 DHCP Aggregate Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	createdID := res.GetId()

	// Cleanup the created resource
	defer func() {
		client.AggregateInterfacesAPI.DeleteAggregateInterfacesByID(context.Background(), createdID).Execute()
	}()

	// *** Exclusivity Assertion ***
	layer3Val, ok := res.GetLayer3Ok()
	require.True(t, ok, "**Layer3 field must be set**")
	assert.False(t, res.HasLayer2(), "**Layer2 field must NOT be set**")

	// L3 DHCP Specific Assertion
	dhcpClientVal, dhcpOk := layer3Val.GetDhcpClientOk()
	require.True(t, dhcpOk, "DhcpClient must be explicitly set")
	assert.True(t, dhcpClientVal.GetEnable(), "DHCP client must be enabled")
	assert.Equal(t, int32(10), dhcpClientVal.GetDefaultRouteMetric(), "Route metric must match")

	// *** Inner Exclusivity Check ***
	require.Empty(t, layer3Val.GetIp(), "Static IP list must be empty when DHCP is configured")
}

// ---------------------------------------------------------------------------------------------------------------------
// --- CRUD Tests (Using L3 Static as the standard setup) ---
// ---------------------------------------------------------------------------------------------------------------------

// Test_AggregateInterfacesAPIService_GetByID tests retrieving an Aggregate Interface by ID.
func Test_AggregateInterfacesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseAggregateInterface(t, "ae-get-")
	intf.SetComment("Original Comment for Get Test")

	// Set L3 Static mode for setup
	layer3Config := network_services.NewAggregateInterfacesLayer3()
	layer3Config.SetIp([]network_services.AggregateInterfacesLayer3IpInner{
		{
			Name: "198.18.1.1/24",
		},
	})
	intf.SetLayer3(*layer3Config)

	// Setup: Create an interface first
	createRes, _, err := client.AggregateInterfacesAPI.CreateAggregateInterfaces(context.Background()).AggregateInterfaces(intf).Execute()
	require.NoError(t, err, "Failed to create interface for get test setup")
	createdID := createRes.GetId()

	defer func() {
		client.AggregateInterfacesAPI.DeleteAggregateInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the interface
	getRes, httpResGet, errGet := client.AggregateInterfacesAPI.GetAggregateInterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get Aggregate Interface by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	assert.Equal(t, intf.GetName(), getRes.GetName(), "Interface name should match")
	assert.Equal(t, "Original Comment for Get Test", getRes.GetComment(), "Comment should be preserved")
	require.True(t, getRes.HasLayer3(), "Layer3 mode should be preserved")
}

// Test_AggregateInterfacesAPIService_Update tests updating an Aggregate Interface.
func Test_AggregateInterfacesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseAggregateInterface(t, "ae-update-")

	// Setup: Create an L2 interface
	intf.SetLayer2(*network_services.NewAggregateInterfacesLayer2())
	createRes, _, err := client.AggregateInterfacesAPI.CreateAggregateInterfaces(context.Background()).AggregateInterfaces(intf).Execute()
	require.NoError(t, err, "Failed to create interface for update test setup")
	createdID := createRes.GetId()

	defer func() {
		client.AggregateInterfacesAPI.DeleteAggregateInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Prepare updated object: Change configuration from L2 to L3 (Static IP)
	updatedIntf := *createRes
	updatedComment := "Updated to Layer3 Static."
	updatedIntf.SetComment(updatedComment)

	// Remove L2 and add L3 configuration
	updatedIntf.Layer2 = nil
	layer3Config := network_services.NewAggregateInterfacesLayer3()
	layer3Config.SetIp([]network_services.AggregateInterfacesLayer3IpInner{
		{
			Name: "198.18.1.1/24",
		},
	})
	updatedIntf.SetLayer3(*layer3Config)

	// Test: Update the interface
	updateRes, httpResUpdate, errUpdate := client.AggregateInterfacesAPI.
		UpdateAggregateInterfacesByID(context.Background(), createdID).
		AggregateInterfaces(updatedIntf).
		Execute()

	require.NoError(t, errUpdate, "Failed to update Aggregate Interface")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Verify the update
	assert.Equal(t, updatedComment, updateRes.GetComment(), "Comment should be updated")
	assert.False(t, updateRes.HasLayer2(), "Layer2 config must be removed")
	require.True(t, updateRes.HasLayer3(), "Layer3 config must be added")
	layer3Val, layer3Ok := updateRes.GetLayer3Ok()
	require.True(t, layer3Ok, "Layer3 config must be present after update (ok boolean)")
	require.NotNil(t, layer3Val, "Layer3 config pointer should not be nil")

	require.NotEmpty(t, layer3Val.GetIp(), "IP list must not be empty")
	assert.Equal(t, "198.18.1.1/24", layer3Val.GetIp()[0].GetName(), "New IP must be set")
}

// Test_AggregateInterfacesAPIService_DeleteByID tests deleting an Aggregate Interface.
func Test_AggregateInterfacesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseAggregateInterface(t, "ae-delete-")
	intf.SetLayer2(*network_services.NewAggregateInterfacesLayer2())

	// Setup: Create an interface first
	createRes, _, err := client.AggregateInterfacesAPI.CreateAggregateInterfaces(context.Background()).AggregateInterfaces(intf).Execute()
	require.NoError(t, err, "Failed to create interface for delete test setup")
	createdID := createRes.GetId()

	// Test: Delete the interface
	httpResDel, errDel := client.AggregateInterfacesAPI.DeleteAggregateInterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete Aggregate Interface")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}

// Test_AggregateInterfacesAPIService_List tests listing Aggregate Interfaces.
func Test_AggregateInterfacesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	intf := createBaseAggregateInterface(t, "ae-list-")
	intf.SetLayer2(*network_services.NewAggregateInterfacesLayer2())

	// Setup: Create an interface first
	createRes, _, err := client.AggregateInterfacesAPI.CreateAggregateInterfaces(context.Background()).AggregateInterfaces(intf).Execute()
	require.NoError(t, err, "Failed to create interface for list test setup")
	createdID := createRes.GetId()

	defer func() {
		client.AggregateInterfacesAPI.DeleteAggregateInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: List the interfaces, filtering by folder
	listRes, httpResList, errList := client.AggregateInterfacesAPI.ListAggregateInterfaces(context.Background()).
		Folder(intf.GetFolder()). // Filter by a unique folder used in setup
		Limit(10).
		Execute()

	require.NoError(t, errList, "Failed to list Aggregate Interfaces")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	assert.GreaterOrEqual(t, len(listRes.Data), 1, "Expected at least one Aggregate Interface in the list")
}
