/*
 * Network Services Testing
 *
 * Layer2SubInterfacesAPIService
 */

package network_services

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// NOTE: Placeholder functions and types are assumed to be available in the test environment.
// - SetupNetworkSvcTestClient(t) -> returns a *network_services.APIClient (with Layer2SubinterfacesAPIService)
// - handleAPIError(err)
// - common.GenerateRandomString(6)

// --- Helper Functions ---

// generateLayer2SubinterfacesName creates a unique name for the resource.
func generateLayer2SubinterfacesName(base string) string {
	return "$" + base
}

// createMinimalLayer2Subinterface creates a minimal Layer2Subinterfaces object for testing.
func createMinimalLayer2Subinterface(t *testing.T, name string, vlanTag string) network_services.Layer2Subinterfaces {
	// 'name' and 'vlan_tag' are the required fields.
	return *network_services.NewLayer2Subinterfaces(name, vlanTag)
}

// createFullLayer2Subinterface creates a comprehensive Layer2Subinterfaces object for update/get testing.
func createFullLayer2Subinterface(t *testing.T, name string, vlanTag string) network_services.Layer2Subinterfaces {
	// 1. Build the base object
	subIf := createMinimalLayer2Subinterface(t, name, vlanTag)

	// 2. Setup optional fields
	parentInterface := "$scm_parent_interface_go"
	comment := "L2 test subinterface for " + name
	folder := "All"

	subIf.SetParentInterface(parentInterface)
	subIf.SetComment(comment)
	subIf.SetFolder(folder)

	return subIf
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_Layer2SubinterfacesAPIService_Create tests the creation of a Layer 2 Subinterface.
func Test_network_services_Layer2SubinterfacesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	subIfName := generateLayer2SubinterfacesName("scm_parent_interface_go.150")
	vlanTag := "100" // Use a common VLAN tag
	subIf := createFullLayer2Subinterface(t, subIfName, vlanTag)

	t.Logf("Creating Layer 2 Subinterface with name: %s and VLAN: %s", subIfName, vlanTag)
	req := client.Layer2SubinterfacesAPI.CreateLayer2Subinterfaces(context.Background()).Layer2Subinterfaces(subIf)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Layer 2 Subinterface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created subinterface should have an ID")
	createdID := *res.Id

	// Cleanup the created subinterface
	defer func() {
		t.Logf("Cleaning up Layer 2 Subinterface with ID: %s", createdID)
		_, errDel := client.Layer2SubinterfacesAPI.DeleteLayer2SubinterfacesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete Layer 2 Subinterface during cleanup")
	}()

	t.Logf("Successfully created Layer 2 Subinterface: %s with ID: %s", subIfName, createdID)

	// Verify key fields in the response
	assert.Equal(t, subIfName, res.GetName(), "Created name should match")
	assert.Equal(t, subIf.GetParentInterface(), res.GetParentInterface(), "Parent interface should match")
}

// Test_network_services_Layer2SubinterfacesAPIService_GetByID tests retrieving a Layer 2 Subinterface by ID.
func Test_network_services_Layer2SubinterfacesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	subIfName := generateLayer2SubinterfacesName("scm_parent_interface_go.400")
	vlanTag := "200"
	subIf := createFullLayer2Subinterface(t, subIfName, vlanTag)

	// Setup: Create a subinterface first
	createRes, _, err := client.Layer2SubinterfacesAPI.CreateLayer2Subinterfaces(context.Background()).Layer2Subinterfaces(subIf).Execute()
	require.NoError(t, err, "Failed to create subinterface for get test setup")
	createdID := *createRes.Id

	defer func() {
		client.Layer2SubinterfacesAPI.DeleteLayer2SubinterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the subinterface
	getRes, httpResGet, errGet := client.Layer2SubinterfacesAPI.GetLayer2SubinterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get Layer 2 Subinterface by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, subIfName, getRes.GetName(), "Subinterface name should match")
	assert.Equal(t, subIf.GetComment(), getRes.GetComment(), "Comment should be preserved")
}

// Test_network_services_Layer2SubinterfacesAPIService_Update tests updating a Layer 2 Subinterface.
func Test_network_services_Layer2SubinterfacesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	subIfName := generateLayer2SubinterfacesName("scm_parent_interface_go.1000")
	vlanTag := "210"
	subIf := createFullLayer2Subinterface(t, subIfName, vlanTag)

	// Setup: Create a subinterface first
	createRes, _, err := client.Layer2SubinterfacesAPI.CreateLayer2Subinterfaces(context.Background()).Layer2Subinterfaces(subIf).Execute()

	require.NoError(t, err, "Failed to create subinterface for put test setup")
	createdID := *createRes.Id

	defer func() {
		client.Layer2SubinterfacesAPI.DeleteLayer2SubinterfacesByID(context.Background(), createdID).Execute()
	}()

	updatedSubIf := *createRes // Dereference the pointer and copy the created object (includes ParentInterface, Folder, etc.)

	updatedComment := "This comment was updated during the test."
	updatedSubIf.SetComment(updatedComment)

	updateRes, httpResUpdate, errUpdate := client.Layer2SubinterfacesAPI.
		UpdateLayer2SubinterfacesByID(context.Background(), createdID).
		Layer2Subinterfaces(updatedSubIf).
		Execute()

	require.NoError(t, errUpdate, "Failed to update Layer 2 Subinterface")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
}

// Test_network_services_Layer2SubinterfacesAPIService_DeleteByID tests deleting a Layer 2 Subinterface.
func Test_network_services_Layer2SubinterfacesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	subIfName := generateLayer2SubinterfacesName("scm_parent_interface_go.500")
	vlanTag := "210"
	subIf := createFullLayer2Subinterface(t, subIfName, vlanTag)

	// Setup: Create a subinterface first
	createRes, _, err := client.Layer2SubinterfacesAPI.CreateLayer2Subinterfaces(context.Background()).Layer2Subinterfaces(subIf).Execute()
	require.NoError(t, err, "Failed to create subinterface for delete test setup")
	createdID := *createRes.Id

	// Test: Delete the subinterface
	httpResDel, errDel := client.Layer2SubinterfacesAPI.DeleteLayer2SubinterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete Layer 2 Subinterface")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}

// Test_network_services_Layer2SubinterfacesAPIService_List tests listing Layer 2 Subinterfaces.
func Test_network_services_Layer2SubinterfacesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a unique resource to ensure the list test finds it and is not empty.
	subIfName := generateLayer2SubinterfacesName("scm_parent_interface_go.600")
	vlanTag := "230" // Use a unique VLAN tag for this resource
	subIf := createFullLayer2Subinterface(t, subIfName, vlanTag)

	// Setup: Create a subinterface first
	createRes, _, err := client.Layer2SubinterfacesAPI.CreateLayer2Subinterfaces(context.Background()).Layer2Subinterfaces(subIf).Execute()
	require.NoError(t, err, "Failed to create subinterface for list test setup")
	createdID := *createRes.Id

	// Cleanup the created subinterface
	defer func() {
		t.Logf("Cleaning up Layer 2 Subinterface with ID: %s", createdID)
		_, errDel := client.Layer2SubinterfacesAPI.DeleteLayer2SubinterfacesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete Layer 2 Subinterface during cleanup")
	}()

	// Test: List the subinterfaces, filtering by known parameters if possible (like folder)
	listRes, httpResList, errList := client.Layer2SubinterfacesAPI.ListLayer2Subinterfaces(context.Background()).
		Folder("All"). // Filter by the folder we set (e.g., "Shared")
		Limit(10).
		Execute()

	require.NoError(t, errList, "Failed to list Layer 2 Subinterfaces")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}
