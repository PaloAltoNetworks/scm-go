/*
 * Network Services Testing
 *
 * Layer3SubInterfacesAPIService
 */

package network_services

import (
	"context"
	"net/http"
	"testing"

	// To handle integer Tag conversion
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Replace with your actual imports

	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// NOTE: Placeholder functions and types are assumed to be available in the test environment.
// - SetupNetworkSvcTestClient(t) -> returns a *network_services.APIClient (with Layer3SubinterfacesAPIService)
// - handleAPIError(err)
// - common.GenerateRandomString(6)
// - common.StringPtr(s) -> returns *string
// - common.Int32Ptr(i) -> returns *int32 (Needed for optional fields like Tag and Mtu)

// --- Helper Functions ---

// generateLayer3SubinterfacesName creates a unique name for the resource.
func generateLayer3SubinterfacesName(base string) string {
	// Replicating your dollar-sign prefix and dynamic suffix
	return "$" + base
}

// createMinimalLayer3Subinterface creates a minimal Layer3Subinterfaces object for testing.
func createMinimalLayer3Subinterface(t *testing.T, name string) network_services.Layer3Subinterfaces {
	// 'name' is the only required field in the constructor.
	return *network_services.NewLayer3Subinterfaces(name)
}

// createFullLayer3Subinterface creates a comprehensive Layer3Subinterfaces object for update/get testing.
func createFullLayer3Subinterface(t *testing.T, baseName string) network_services.Layer3Subinterfaces {
	name := generateLayer3SubinterfacesName(baseName)

	// 1. Build the base object
	subIf := createMinimalLayer3Subinterface(t, name)

	// 2. Setup optional fields
	parentInterface := "$scm_parent_interface_go_l3"
	comment := "L3 test subinterface for " + name
	folder := "All"
	tag := int32(100)
	mtu := int32(580)

	// --- Complex Child Structures ---
	// Assuming Layer3SubinterfacesIpInner has a V4 field for IPv4
	ipConfig := []network_services.Layer3SubinterfacesIpInner{
		{
			// Placeholder - replace 'V4' with the actual field name for IPv4 address/mask
			// This is critical for L3 interfaces.
			Name: "192.168.10.1/24",
		},
	}

	subIf.SetParentInterface(parentInterface)
	subIf.SetComment(comment)
	subIf.SetFolder(folder)
	subIf.SetTag(tag)
	subIf.SetMtu(mtu)
	subIf.SetIp(ipConfig)

	return subIf
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_Layer3SubinterfacesAPIService_Create tests the creation of a Layer 3 Subinterface.
func Test_network_services_Layer3SubinterfacesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	subIf := createFullLayer3Subinterface(t, "scm_parent_interface_go_l3.100")

	t.Logf("Creating Layer 3 Subinterface with name: %s", subIf.GetName())
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

	t.Logf("Successfully created Layer 3 Subinterface: %s with ID: %s", subIf.GetName(), createdID)

	// Verify key fields in the response
	assert.Equal(t, subIf.GetName(), res.GetName(), "Created name should match")
	assert.Equal(t, subIf.GetTag(), res.GetTag(), "VLAN Tag should match")
	assert.Len(t, res.GetIp(), 1, "IP config list should have one entry")
}

// Test_network_services_Layer3SubinterfacesAPIService_GetByID tests retrieving a Layer 3 Subinterface by ID.
func Test_network_services_Layer3SubinterfacesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	subIf := createFullLayer3Subinterface(t, "scm_parent_interface_go_l3.200")

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
	assert.Equal(t, subIf.GetName(), getRes.GetName(), "Subinterface name should match")
	assert.Equal(t, subIf.GetComment(), getRes.GetComment(), "Comment should be preserved")
	assert.True(t, getRes.HasMtu(), "MTU field should be present")
}

// Test_network_services_Layer3SubinterfacesAPIService_Update tests updating a Layer 3 Subinterface.
func Test_network_services_Layer3SubinterfacesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	subIf := createFullLayer3Subinterface(t, "scm_parent_interface_go_l3.900")

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
	subIf := createFullLayer3Subinterface(t, "scm_parent_interface_go_l3.400")

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

	// Create a unique resource to ensure the list test finds it.
	subIf := createFullLayer3Subinterface(t, "scm_parent_interface_go_l3.500")

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
