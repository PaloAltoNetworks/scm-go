/*
 /*
 * Deployment Service Testing
 *
 * ServiceConnectionGroupsAPIService
*/

package deployment_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	// NOTE: If CreateTestServiceConnection and related utils are in a different package,
	// you may need to import that package here. Assuming they are accessible or mocked.
)

// --- SERVICE CONNECTION GROUP TEST SUITE ---

// Helper to create a single Service Connection dependency.
// NOTE: This assumes CreateTestServiceConnection is a utility method defined in a local or imported utility file.
func createTestServiceConnection(t *testing.T, depSvcClient *deployment_services.APIClient, suffix string) (string, func()) {
	// A SC requires an IPsec Tunnel, which itself requires an IKE Gateway and Crypto Profile.
	// We call the utility function that creates the full dependency chain.
	tunnel_suffix := common.GenerateRandomString(6)
	tunnelName, cleanupDeps := createSCTestIPsecTunnelAndDeps(t, "scg-"+tunnel_suffix)

	// Create the Service Connection that will be a target for the group.
	scName := "test-sc-for-group-" + suffix
	sc := deployment_services.ServiceConnections{
		Name:        scName,
		IpsecTunnel: tunnelName,      // Dependency
		Region:      "us-central1-a", // Required field
	}

	req := depSvcClient.ServiceConnectionsAPI.CreateServiceConnections(context.Background()).ServiceConnections(sc)
	res, _, err := req.Execute()

	// Use the friendly error handling logic defined in your utils
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Service Connection dependency for group test")

	createdSCID := res.Id

	// Cleanup function deletes the SC and all its underlying dependencies.
	cleanup := func() {
		// 1. Delete the Service Connection
		t.Logf("Cleaning up Service Connection dependency: %s", createdSCID)
		_, errDel := depSvcClient.ServiceConnectionsAPI.DeleteServiceConnectionsByID(context.Background(), createdSCID).Execute()
		require.NoError(t, errDel, "Failed to delete Service Connection during cleanup")

		// 2. Clean up SC's underlying dependencies (Tunnel, Gateway, Crypto Profile)
		cleanupDeps()
	}

	return scName, cleanup
}

// Test_deployment_services_ServiceConnectionGroupsAPIService_Create tests the creation of a Service Connection Group.
func Test_deployment_services_ServiceConnectionGroupsAPIService_Create(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// 1. Create dependency (at least one Service Connection)
	targetSCName, cleanupSC := createTestServiceConnection(t, depSvcClient, "create-"+randomSuffix)
	defer cleanupSC()

	// 2. Define a valid Service Connection Group object
	groupName := "test-sc-group-" + randomSuffix

	scGroup := deployment_services.ServiceConnectionGroups{
		Name: groupName,
		// The 'Target' array must contain the name of the Service Connection we just created.
		Target: []string{targetSCName},
	}

	// Set optional fields
	scGroup.SetDisableSnat(true)

	fmt.Printf("Attempting to create Service Connection Group with name: %s\n", scGroup.Name)

	// Make the create request to the API
	req := depSvcClient.ServiceConnectionGroupsAPI.CreateServiceConnectionGroups(context.Background()).ServiceConnectionGroups(scGroup)
	res, httpRes, err := req.Execute()

	// Use the friendly error handler
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create Service Connection Group")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotEmpty(t, res.Id, "Created SC Group should have an ID")

	createdGroupID := res.Id

	// Defer the cleanup of the created Service Connection Group.
	defer func() {
		t.Logf("Cleaning up Service Connection Group with ID: %s", createdGroupID)
		_, errDel := depSvcClient.ServiceConnectionGroupsAPI.DeleteServiceConnectionGroupsByID(context.Background(), createdGroupID).Execute()
		require.NoError(t, errDel, "Failed to delete Service Connection Group during cleanup")
	}()

	// Assert response object properties
	assert.Equal(t, groupName, res.Name, "Created Group name should match")
	assert.Contains(t, res.Target, targetSCName, "Target list should contain the created Service Connection")
	assert.Equal(t, true, res.GetDisableSnat(), "Disable SNAT should match input")

	t.Logf("Successfully created and validated SC Group: %s with ID: %s", scGroup.Name, createdGroupID)
}

// -------------------------------------------------------------------------------------------

// Test_deployment_services_ServiceConnectionGroupsAPIService_GetByID tests retrieving a SC Group by its ID.
func Test_deployment_services_ServiceConnectionGroupsAPIService_GetByID(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// 1. Create dependencies
	targetSCName, cleanupSC := createTestServiceConnection(t, depSvcClient, "get-"+randomSuffix)
	defer cleanupSC()

	// 2. Create a Service Connection Group to retrieve.
	groupName := "test-sc-group-get-" + randomSuffix
	scGroup := deployment_services.ServiceConnectionGroups{
		Name:   groupName,
		Target: []string{targetSCName},
	}

	createRes, _, err := depSvcClient.ServiceConnectionGroupsAPI.CreateServiceConnectionGroups(context.Background()).ServiceConnectionGroups(scGroup).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create SC Group for get test")
	createdGroupID := createRes.Id
	defer func() {
		depSvcClient.ServiceConnectionGroupsAPI.DeleteServiceConnectionGroupsByID(context.Background(), createdGroupID).Execute()
	}()

	// 3. Test Get by ID operation
	getRes, httpResGet, errGet := depSvcClient.ServiceConnectionGroupsAPI.GetServiceConnectionGroupsByID(context.Background(), createdGroupID).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful
	require.NoError(t, errGet, "Failed to get Service Connection Group by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Assert key properties
	assert.Equal(t, groupName, getRes.Name, "Group name should match")
	assert.Equal(t, createdGroupID, getRes.Id, "Group ID should match")
	assert.Contains(t, getRes.Target, targetSCName, "Target SC should be present")

	t.Logf("Successfully retrieved SC Group: %s", getRes.Name)
}

// -------------------------------------------------------------------------------------------

// Test_deployment_services_ServiceConnectionGroupsAPIService_List tests listing SC Groups with folder filter.
func Test_deployment_services_ServiceConnectionGroupsAPIService_List(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// 1. Create dependency
	targetSCName, cleanupSC := createTestServiceConnection(t, depSvcClient, "list-"+randomSuffix)
	defer cleanupSC()

	// 2. Create a unique SC Group to ensure it appears in the list.
	groupName := "test-sc-group-list-" + randomSuffix
	folder := "Service Connections"
	scGroup := deployment_services.ServiceConnectionGroups{
		Name:   groupName,
		Target: []string{targetSCName},
	}

	createRes, _, err := depSvcClient.ServiceConnectionGroupsAPI.CreateServiceConnectionGroups(context.Background()).ServiceConnectionGroups(scGroup).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create SC Group for list test")
	createdGroupID := createRes.Id
	defer func() {
		depSvcClient.ServiceConnectionGroupsAPI.DeleteServiceConnectionGroupsByID(context.Background(), createdGroupID).Execute()
	}()

	// 3. Test List operation, filtering by folder (required by List API signature)
	reqList := depSvcClient.ServiceConnectionGroupsAPI.ListServiceConnectionGroups(context.Background()).Folder(folder).Limit(100)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list Service Connection Groups")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes)

	// 4. Verify our created SC Group is in the list
	foundGroup := false
	for _, groupRes := range listRes.Data {
		if groupRes.Name == groupName {
			foundGroup = true
			assert.Contains(t, groupRes.Target, targetSCName, "Target SC should be present in listed group")
			break
		}
	}
	assert.True(t, foundGroup, "Created Service Connection Group should be found in the list")

	t.Logf("Successfully listed SC Groups, found created group: %s", groupName)
}

// -------------------------------------------------------------------------------------------

// Test_deployment_services_ServiceConnectionGroupsAPIService_DeleteByID tests deleting a SC Group by its ID.
func Test_deployment_services_ServiceConnectionGroupsAPIService_DeleteByID(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// 1. Create dependency
	targetSCName, cleanupSC := createTestServiceConnection(t, depSvcClient, "del-"+randomSuffix)
	defer cleanupSC()

	// 2. Create a Service Connection Group to delete.
	groupName := "test-sc-group-del-" + randomSuffix
	scGroup := deployment_services.ServiceConnectionGroups{
		Name:   groupName,
		Target: []string{targetSCName},
	}

	createRes, _, err := depSvcClient.ServiceConnectionGroupsAPI.CreateServiceConnectionGroups(context.Background()).ServiceConnectionGroups(scGroup).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create SC Group for delete test")
	createdGroupID := createRes.Id

	// 3. Test Delete by ID operation
	_, errDel := depSvcClient.ServiceConnectionGroupsAPI.DeleteServiceConnectionGroupsByID(context.Background(), createdGroupID).Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// 4. Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete Service Connection Group")

	t.Logf("Successfully deleted SC Group: %s", createdGroupID)

	t.Logf("Verified SC Group deletion: %s", createdGroupID)
}

// Test_deployment_services_ServiceConnectionGroupsAPIService_Update tests updating an existing SC Group.
func Test_deployment_services_ServiceConnectionGroupsAPIService_Update(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// 1. Create dependencies (two Service Connections)
	targetSCName1, cleanupSC1 := createTestServiceConnection(t, depSvcClient, "update-"+randomSuffix)
	defer cleanupSC1()
	targetSCName2, cleanupSC2 := createTestServiceConnection(t, depSvcClient, "update2-"+randomSuffix)
	defer cleanupSC2()

	// 2. Create the initial SC Group.
	groupName := "test-sc-group-upd-" + randomSuffix
	initialGroup := deployment_services.ServiceConnectionGroups{
		Name:        groupName,
		Target:      []string{targetSCName1}, // Initial target
		DisableSnat: common.BoolPtr(true),
	}

	createRes, _, err := depSvcClient.ServiceConnectionGroupsAPI.CreateServiceConnectionGroups(context.Background()).ServiceConnectionGroups(initialGroup).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create SC Group for update test")
	createdGroupID := createRes.Id
	defer func() {
		depSvcClient.ServiceConnectionGroupsAPI.DeleteServiceConnectionGroupsByID(context.Background(), createdGroupID).Execute()
	}()

	// 3. Define the update payload.
	updatedGroup := deployment_services.ServiceConnectionGroups{
		Name:        groupName,
		Id:          createdGroupID,
		Target:      []string{targetSCName1, targetSCName2}, // ADDED targetSCName2
		DisableSnat: common.BoolPtr(false),                  // Updated boolean
		PbfOnly:     common.BoolPtr(true),                   // Set optional field
	}

	// 4. Execute Update
	reqUpdate := depSvcClient.ServiceConnectionGroupsAPI.UpdateServiceConnectionGroupsByID(context.Background(), createdGroupID).ServiceConnectionGroups(updatedGroup)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update Service Connection Group")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// 5. Assert updated properties
	assert.Len(t, updateRes.Target, 2, "Should now have two targets")
	assert.Contains(t, updateRes.Target, targetSCName2, "New target should be added")
	assert.Equal(t, false, updateRes.GetDisableSnat(), "Disable SNAT should be updated to false")
	assert.Equal(t, true, updateRes.GetPbfOnly(), "PBF Only should be updated to true")

	t.Logf("Successfully updated SC Group: %s", groupName)
}
