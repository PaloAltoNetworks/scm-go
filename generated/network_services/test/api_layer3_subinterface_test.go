/*
 * Network Services Testing
 *
 * Layer3SubInterfacesAPIService
 */

package network_services

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time" // Needed for unique name generation

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// NOTE: Placeholder functions and types are assumed to be available in the test environment.
// - SetupNetworkSvcTestClient(t) -> returns a *network_services.APIClient
// - handleAPIError(err)

// --- Shared Base Helper Functions (Assumed to be defined/reused from L2 file) ---
// We define them here to ensure the logic flow is complete and correct for this file.

// generateUniqueName creates a unique name using a timestamp.
func generateUniqueName(base string) string {
	return base + fmt.Sprintf("%d", time.Now().UnixNano()/1e6) // Use milliseconds for uniqueness
}

// generateEthernetInterfaceName creates a unique name for the Ethernet Interface.
func generateEthernetInterfaceNameforL3(base string) string {
	// The '$' prefix is often used for dynamic/local interface names in SCM environments.
	return "$" + generateUniqueName(base)
}

// createBaseEthernetInterface creates a base EthernetInterface object WITHOUT providing an Id.
func createBaseEthernetInterfaceforL3(t *testing.T, baseName string) network_services.EthernetInterfaces {
	name := generateEthernetInterfaceNameforL3(baseName)

	// Use the constructor with defaults, then manually set the required fields ('Name').
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

// setupL3EthernetInterface creates a new Ethernet Interface configured for Layer 3
// mode via API, asserts success, and returns its unique name and a cleanup function.
func setupL3EthernetInterface(t *testing.T, client *network_services.APIClient) (string, func()) {
	// 1. Create the base interface object using the shared helper
	intf := createBaseEthernetInterfaceforL3(t, "l3-parent-intf-")

	// 2. Set Layer 3 mode (Minimal required config for L3)
	intf.SetLayer3(*network_services.NewEthernetInterfacesLayer3WithDefaults())

	// Get the generated name
	intfName := intf.GetName()

	t.Logf("Creating Layer 3 Parent Interface with name: %s", intfName)

	// 3. Execute the API call
	createRes, _, err := client.EthernetInterfacesAPI.
		CreateEthernetInterfaces(context.Background()).
		EthernetInterfaces(intf).
		Execute()

	// Assert creation success
	require.NoError(t, err, "Failed to create L3 Parent Ethernet Interface for test setup")

	// 4. Get the generated ID for cleanup
	createdID := createRes.GetId()
	require.NotEmpty(t, createdID, "Created interface must have a generated ID")

	// 5. Define and return cleanup function
	cleanup := func() {
		t.Logf("Cleaning up L3 Parent Interface with ID: %s", createdID)
		_, errDel := client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), createdID).Execute()
		if errDel != nil {
			t.Logf("Warning: Failed to clean up L3 Parent Interface: %v", errDel)
		}
	}

	return intfName, cleanup
}

// createMinimalLayer3Subinterface creates a minimal Layer3Subinterfaces object for testing.
func createMinimalLayer3Subinterface(t *testing.T, name string) network_services.Layer3Subinterfaces {
	// 'name' is the only required field in the constructor.
	return *network_services.NewLayer3Subinterfaces(name)
}

// createFullLayer3Subinterface creates a comprehensive Layer3Subinterfaces object.
// It calculates the final subinterface name internally as [ParentIfName].[VLAN Tag].
// MODIFIED: BaseName parameter is removed, now takes parentIfName and vlanTag directly.
func createFullLayer3Subinterface(t *testing.T, parentIfName string, vlanTag string) network_services.Layer3Subinterfaces {
	// 1. Calculate the final required name: ParentName.VLAN_Tag (e.g., $l3-parent-intf-123.400)
	subIfName := fmt.Sprintf("%s.%s", parentIfName, vlanTag)

	// 2. Build the base object
	subIf := createMinimalLayer3Subinterface(t, subIfName)

	// 3. Setup required/optional fields
	comment := "L3 test subinterface for " + subIfName
	folder := "All"

	// The VLAN tag is usually passed as an int32 field
	tag, err := assertInt32(vlanTag) // Assuming assertInt32 helper exists or use direct conversion/const
	if err != nil {
		t.Fatalf("Failed to convert VLAN tag string to int32: %v", err)
	}

	mtu := int32(580)

	// Complex IP Configuration structure
	ipConfig := []network_services.Layer3SubinterfacesIpInner{
		{
			// NOTE: Name field used for CIDR address.
			Name: "192.168.10.1/24",
		},
	}

	subIf.SetParentInterface(parentIfName)
	subIf.SetComment(comment)
	subIf.SetFolder(folder)
	subIf.SetTag(tag)
	subIf.SetMtu(mtu)
	subIf.SetIp(ipConfig)

	return subIf
}

// Assumed helper to convert string to int32 (since we removed common.Int32Ptr)
func assertInt32(s string) (int32, error) {
	// Simple placeholder conversion logic
	var i int32 = 0
	fmt.Sscanf(s, "%d", &i)
	return i, nil // Ignoring actual error handling for test code simplicity
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_Layer3SubinterfacesAPIService_Create tests the creation of a Layer 3 Subinterface.
func Test_network_services_Layer3SubinterfacesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// --- 1. SETUP PREREQUISITE: Create the L3 Parent Interface ---
	parentIfName, parentCleanup := setupL3EthernetInterface(t, client)
	defer parentCleanup()

	vlanTag := "400"

	// 2. CREATE SUBINTERFACE OBJECT: The helper calculates the name: [ParentName].400
	subIf := createFullLayer3Subinterface(t, parentIfName, vlanTag)
	subIfName := subIf.GetName()

	t.Logf("Creating Layer 3 Subinterface with name: %s", subIfName)
	req := client.Layer3SubinterfacesAPI.CreateLayer3Subinterfaces(context.Background()).Layer3Subinterfaces(subIf)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Layer 3 Subinterface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created subinterface should have an ID")
	createdID := *res.Id

	// Cleanup the created subinterface
	defer func() {
		t.Logf("Cleaning up Layer 3 Subinterface with ID: %s", createdID)
		_, errDel := client.Layer3SubinterfacesAPI.DeleteLayer3SubinterfacesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete Layer 3 Subinterface during cleanup")
	}()

	t.Logf("Successfully created Layer 3 Subinterface: %s with ID: %s", subIfName, createdID)

	// Verify key fields in the response
	assert.Equal(t, subIfName, res.GetName(), "Created name should match the calculated Parent.Tag name")
	assert.Equal(t, subIf.GetTag(), res.GetTag(), "VLAN Tag should match")
	assert.Equal(t, parentIfName, res.GetParentInterface(), "Parent interface should match the dynamically created parent")
	assert.Len(t, res.GetIp(), 1, "IP config list should have one entry")
}

// Test_network_services_Layer3SubinterfacesAPIService_GetByID tests retrieving a Layer 3 Subinterface by ID.
func Test_network_services_Layer3SubinterfacesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// --- 1. SETUP PREREQUISITE: Create the L3 Parent Interface ---
	parentIfName, parentCleanup := setupL3EthernetInterface(t, client)
	defer parentCleanup()

	vlanTag := "200"
	subIf := createFullLayer3Subinterface(t, parentIfName, vlanTag)
	subIfName := subIf.GetName()

	// Setup: Create a subinterface first
	createRes, _, err := client.Layer3SubinterfacesAPI.CreateLayer3Subinterfaces(context.Background()).Layer3Subinterfaces(subIf).Execute()
	require.NoError(t, err, "Failed to create subinterface for get test setup")
	createdID := *createRes.Id

	defer func() {
		client.Layer3SubinterfacesAPI.DeleteLayer3SubinterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the subinterface
	getRes, httpResGet, errGet := client.Layer3SubinterfacesAPI.GetLayer3SubinterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get Layer 3 Subinterface by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, subIfName, getRes.GetName(), "Subinterface name should match")
	assert.Equal(t, subIf.GetComment(), getRes.GetComment(), "Comment should be preserved")
	assert.True(t, getRes.HasMtu(), "MTU field should be present")
}

// Test_network_services_Layer3SubinterfacesAPIService_Update tests updating a Layer 3 Subinterface.
func Test_network_services_Layer3SubinterfacesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// --- 1. SETUP PREREQUISITE: Create the L3 Parent Interface ---
	parentIfName, parentCleanup := setupL3EthernetInterface(t, client)
	defer parentCleanup()

	vlanTag := "900"
	subIf := createFullLayer3Subinterface(t, parentIfName, vlanTag)

	// Setup: Create a subinterface first
	createRes, _, err := client.Layer3SubinterfacesAPI.CreateLayer3Subinterfaces(context.Background()).Layer3Subinterfaces(subIf).Execute()

	require.NoError(t, err, "Failed to create subinterface for update test setup")
	createdID := *createRes.Id

	defer func() {
		client.Layer3SubinterfacesAPI.DeleteLayer3SubinterfacesByID(context.Background(), createdID).Execute()
	}()

	// 1. START WITH THE OBJECT RETURNED FROM THE CREATE CALL
	updatedSubIf := *createRes

	// 2. MAKE YOUR MODIFICATION
	updatedComment := "This L3 comment was updated during the test."
	updatedMTU := int32(1400) // Change the MTU
	updatedSubIf.SetComment(updatedComment)
	updatedSubIf.SetMtu(updatedMTU)

	updateRes, httpResUpdate, errUpdate := client.Layer3SubinterfacesAPI.
		UpdateLayer3SubinterfacesByID(context.Background(), createdID).
		Layer3Subinterfaces(updatedSubIf).
		Execute()

	require.NoError(t, errUpdate, "Failed to update Layer 3 Subinterface")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the update in the response
	assert.Equal(t, updatedComment, updateRes.GetComment(), "Comment should be updated")
	assert.Equal(t, updatedMTU, updateRes.GetMtu(), "MTU should be updated")
}

// Test_network_services_Layer3SubinterfacesAPIService_DeleteByID tests deleting a Layer 3 Subinterface.
func Test_network_services_Layer3SubinterfacesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// --- 1. SETUP PREREQUISITE: Create the L3 Parent Interface ---
	parentIfName, parentCleanup := setupL3EthernetInterface(t, client)
	defer parentCleanup()

	vlanTag := "400"
	subIf := createFullLayer3Subinterface(t, parentIfName, vlanTag)

	// Setup: Create a subinterface first
	createRes, _, err := client.Layer3SubinterfacesAPI.CreateLayer3Subinterfaces(context.Background()).Layer3Subinterfaces(subIf).Execute()
	require.NoError(t, err, "Failed to create subinterface for delete test setup")
	createdID := *createRes.Id

	// Test: Delete the subinterface
	httpResDel, errDel := client.Layer3SubinterfacesAPI.DeleteLayer3SubinterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete Layer 3 Subinterface")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}

// Test_network_services_Layer3SubinterfacesAPIService_List tests listing Layer 3 Subinterfaces.
func Test_network_services_Layer3SubinterfacesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// --- 1. SETUP PREREQUISITE: Create the L3 Parent Interface ---
	parentIfName, parentCleanup := setupL3EthernetInterface(t, client)
	defer parentCleanup()

	vlanTag := "500"
	subIf := createFullLayer3Subinterface(t, parentIfName, vlanTag)

	// Setup: Create a subinterface first
	createRes, _, err := client.Layer3SubinterfacesAPI.CreateLayer3Subinterfaces(context.Background()).Layer3Subinterfaces(subIf).Execute()
	require.NoError(t, err, "Failed to create subinterface for list test setup")
	createdID := *createRes.Id

	// Cleanup the created subinterface
	defer func() {
		client.Layer3SubinterfacesAPI.DeleteLayer3SubinterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: List the subinterfaces
	listRes, httpResList, errList := client.Layer3SubinterfacesAPI.ListLayer3Subinterfaces(context.Background()).
		Folder(subIf.GetFolder()).
		Limit(10).
		Execute()

	require.NoError(t, errList, "Failed to list Layer 3 Subinterfaces")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}
