/*
 * Network Services Testing
 *
 * Layer2SubInterfacesAPIService
 */

package network_services

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// NOTE: Placeholder functions and types are assumed to be available in the test environment.
// - SetupNetworkSvcTestClient(t) -> returns a *network_services.APIClient
// - handleAPIError(err)

// --- Helper Functions ---

// generateUniqueName creates a unique name using a timestamp.
func generateUniqueNameforSubInterfaces(base string) string {
	return base + fmt.Sprintf("%d", time.Now().UnixNano()/1e6) // Use milliseconds for uniqueness
}

// generateEthernetInterfaceName creates a unique name for the Ethernet Interface.
func generateEthernetInterfaceNameforSubInterfaces(base string) string {
	// The '$' prefix is often used for dynamic/local interface names in SCM environments.
	return "$" + generateUniqueNameforSubInterfaces(base)
}

// createBaseEthernetInterface creates a base EthernetInterface object WITHOUT providing an Id.
func createBaseEthernetInterfaceforSubInterfaces(t *testing.T, baseName string) network_services.EthernetInterfaces {
	name := generateEthernetInterfaceNameforSubInterfaces(baseName)

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

// setupL2EthernetInterface creates a new Ethernet Interface configured for Layer 2
// mode via API, asserts success, and returns its unique name and a cleanup function.
func setupL2EthernetInterface(t *testing.T, client *network_services.APIClient) (string, func()) {
	// 1. Create the base interface object
	intf := createBaseEthernetInterfaceforSubInterfaces(t, "l2-parent-intf-")

	// 2. Set Layer 2 mode (Minimal required config for L2)
	intf.SetLayer2(*network_services.NewEthernetInterfacesLayer2WithDefaults())

	// Get the generated name (e.g., $l2-parent-intf-1700000000)
	intfName := intf.GetName()

	t.Logf("Creating Layer 2 Parent Interface with name: %s", intfName)

	// 3. Execute the API call
	createRes, _, err := client.EthernetInterfacesAPI.
		CreateEthernetInterfaces(context.Background()).
		EthernetInterfaces(intf).
		Execute()

	// Assert creation success
	require.NoError(t, err, "Failed to create L2 Parent Ethernet Interface for test setup")

	// 4. Get the generated ID for cleanup
	createdID := createRes.GetId()
	require.NotEmpty(t, createdID, "Created interface must have a generated ID")

	// 5. Define and return cleanup function
	cleanup := func() {
		t.Logf("Cleaning up L2 Parent Interface with ID: %s", createdID)
		_, errDel := client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), createdID).Execute()
		// Note: We ignore cleanup error for prerequisite interfaces if the main test passed.
		if errDel != nil {
			t.Logf("Warning: Failed to clean up L2 Parent Interface: %v", errDel)
		}
	}

	return intfName, cleanup
}

// createMinimalLayer2Subinterface creates a minimal Layer2Subinterfaces object for testing.
// NOTE: This version is minimal for completeness but the full creation is done in the next function.
func createMinimalLayer2Subinterface(t *testing.T, name string, vlanTag string) network_services.Layer2Subinterfaces {
	// Assumes NewLayer2Subinterfaces(name, vlan_tag) is the constructor
	return *network_services.NewLayer2Subinterfaces(name, vlanTag)
}

// createFullLayer2Subinterface creates a comprehensive Layer2Subinterfaces object.
// It calculates the final subinterface name internally as [ParentIfName].[VLAN Tag].
func createFullLayer2Subinterface(t *testing.T, parentIfName string, vlanTag string) network_services.Layer2Subinterfaces {
	// 1. Calculate the final required name: ParentName.VLAN_Tag (e.g., $l2-parent-intf-123.400)
	subIfName := fmt.Sprintf("%s.%s", parentIfName, vlanTag)

	// 2. Build the base object
	subIf := createMinimalLayer2Subinterface(t, subIfName, vlanTag)

	// 3. Setup optional fields
	comment := "L2 test subinterface for " + subIfName
	folder := "All"

	// Set the dynamic parent interface name
	subIf.SetParentInterface(parentIfName)

	subIf.SetComment(comment)
	subIf.SetFolder(folder)

	return subIf
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_Layer2SubinterfacesAPIService_Create tests the creation of a Layer 2 Subinterface.
func Test_network_services_Layer2SubinterfacesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// --- 1. SETUP PREREQUISITE: Create the L2 Parent Interface ---
	parentIfName, parentCleanup := setupL2EthernetInterface(t, client)
	defer parentCleanup()

	vlanTag := "400" // Use a consistent VLAN tag for this test

	// 2. CREATE SUBINTERFACE OBJECT: The helper calculates the name: [ParentName].400
	subIf := createFullLayer2Subinterface(t, parentIfName, vlanTag)
	subIfName := subIf.GetName() // Get the calculated name for logging/verification

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
	assert.Equal(t, subIfName, res.GetName(), "Created name should match the calculated Parent.Tag name")
	assert.Equal(t, parentIfName, res.GetParentInterface(), "Parent interface should match the dynamically created parent")
}

// Test_network_services_Layer2SubinterfacesAPIService_GetByID tests retrieving a Layer 2 Subinterface by ID.
func Test_network_services_Layer2SubinterfacesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// --- 1. SETUP PREREQUISITE: Create the L2 Parent Interface ---
	parentIfName, parentCleanup := setupL2EthernetInterface(t, client)
	defer parentCleanup()

	vlanTag := "200"
	subIf := createFullLayer2Subinterface(t, parentIfName, vlanTag)
	subIfName := subIf.GetName()

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
	assert.Equal(t, parentIfName, getRes.GetParentInterface(), "Parent interface should be preserved")
}

// Test_network_services_Layer2SubinterfacesAPIService_Update tests updating a Layer 2 Subinterface.
func Test_network_services_Layer2SubinterfacesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// --- 1. SETUP PREREQUISITE: Create the L2 Parent Interface ---
	parentIfName, parentCleanup := setupL2EthernetInterface(t, client)
	defer parentCleanup()

	vlanTag := "210"
	subIf := createFullLayer2Subinterface(t, parentIfName, vlanTag)

	// Setup: Create a subinterface first
	createRes, _, err := client.Layer2SubinterfacesAPI.CreateLayer2Subinterfaces(context.Background()).Layer2Subinterfaces(subIf).Execute()

	require.NoError(t, err, "Failed to create subinterface for put test setup")
	createdID := *createRes.Id

	defer func() {
		client.Layer2SubinterfacesAPI.DeleteLayer2SubinterfacesByID(context.Background(), createdID).Execute()
	}()

	updatedSubIf := *createRes // Copy the created object

	updatedComment := "This comment was updated during the test."
	updatedSubIf.SetComment(updatedComment)

	updateRes, httpResUpdate, errUpdate := client.Layer2SubinterfacesAPI.
		UpdateLayer2SubinterfacesByID(context.Background(), createdID).
		Layer2Subinterfaces(updatedSubIf).
		Execute()

	require.NoError(t, errUpdate, "Failed to update Layer 2 Subinterface")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the change
	assert.Equal(t, updatedComment, updateRes.GetComment(), "Comment should be updated")
	assert.Equal(t, parentIfName, updateRes.GetParentInterface(), "Parent interface should remain the same")
}

// Test_network_services_Layer2SubinterfacesAPIService_DeleteByID tests deleting a Layer 2 Subinterface.
func Test_network_services_Layer2SubinterfacesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// --- 1. SETUP PREREQUISITE: Create the L2 Parent Interface ---
	parentIfName, parentCleanup := setupL2EthernetInterface(t, client)
	defer parentCleanup()

	vlanTag := "500"
	subIf := createFullLayer2Subinterface(t, parentIfName, vlanTag)

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

	// --- 1. SETUP PREREQUISITE: Create the L2 Parent Interface ---
	parentIfName, parentCleanup := setupL2EthernetInterface(t, client)
	defer parentCleanup()

	vlanTag := "600"
	subIf := createFullLayer2Subinterface(t, parentIfName, vlanTag)

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
