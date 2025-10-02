/*
Objects Testing ApplicationGroupsAPIService
*/
package objects

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)

// Helper function to create a test application object.
func createTestApplication(t *testing.T, client *objects.APIClient, name string) string {
	app := objects.Applications{
		Folder:      common.StringPtr("Shared"),
		Name:        name,
		Category:    "business-systems",
		Subcategory: common.StringPtr("database"),
		Technology:  common.StringPtr("client-server"),
		Risk:        1,
	}
	req := client.ApplicationsAPI.CreateApplications(context.Background()).Applications(app)
	res, _, err := req.Execute()
	require.NoError(t, err, "Failed to create test application")
	require.NotNil(t, res, "Test application create response should not be nil")
	t.Logf("Created test application %s with ID %s", name, *res.Id)
	return *res.Id
}

// Helper function to delete a test application object.
func deleteTestApplication(t *testing.T, client *objects.APIClient, id, name string) {
	req := client.ApplicationsAPI.DeleteApplicationsByID(context.Background(), id)
	_, err := req.Execute()
	require.NoError(t, err, "Failed to delete test application %s", name)
	t.Logf("Deleted test application %s", name)
}

// Test_objects_ApplicationGroupsAPIService_Create tests the creation of an application group object.
func Test_objects_ApplicationGroupsAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create test applications first.
	randomSuffix := common.GenerateRandomString(6)
	app1Name := "test-app-1-" + randomSuffix
	app2Name := "test-app-2-" + randomSuffix
	app1ID := createTestApplication(t, client, app1Name)
	app2ID := createTestApplication(t, client, app2Name)

	// Create a valid application group object with a unique name.
	createdAppGroupName := "test-app-group-create-" + randomSuffix
	applicationGroup := objects.ApplicationGroups{
		Folder:  common.StringPtr("Shared"),
		Name:    createdAppGroupName,
		Members: []string{app1Name, app2Name},
	}

	fmt.Printf("Creating application group with name: %s\n", applicationGroup.Name)

	// Make the create request to the API.
	req := client.ApplicationGroupsAPI.CreateApplicationGroups(context.Background()).ApplicationGroups(applicationGroup)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create application group")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties.
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdAppGroupName, res.Name, "Created application group name should match")
	assert.ElementsMatch(t, []string{app1Name, app2Name}, res.Members, "Members list should match")
	assert.NotEmpty(t, res.Id, "Created application group should have an ID")

	createdAppGroupID := res.Id
	t.Logf("Successfully created application group: %s with ID: %s", applicationGroup.Name, createdAppGroupID)

	// Cleanup: Delete the created application group and test applications.
	delReq := client.ApplicationGroupsAPI.DeleteApplicationGroupsByID(context.Background(), createdAppGroupID)
	_, errDel := delReq.Execute()
	if errDel == nil {
		t.Logf("Successfully deleted application group: %s with ID: %s", applicationGroup.Name, createdAppGroupID)
	}
	require.NoError(t, errDel, "Failed to delete application group during cleanup")
	deleteTestApplication(t, client, app1ID, app1Name)
	deleteTestApplication(t, client, app2ID, app2Name)
}

// Test_objects_ApplicationGroupsAPIService_GetByID tests retrieving an application group by its ID.
func Test_objects_ApplicationGroupsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create test applications first.
	randomSuffix := common.GenerateRandomString(6)
	app1Name := "test-app-get-1-" + randomSuffix
	app1ID := createTestApplication(t, client, app1Name)

	// Create an application group to retrieve.
	createdAppGroupName := "test-app-group-get-" + randomSuffix
	applicationGroup := objects.ApplicationGroups{
		Folder:  common.StringPtr("Shared"),
		Name:    createdAppGroupName,
		Members: []string{app1Name},
	}
	createRes, _, err := client.ApplicationGroupsAPI.CreateApplicationGroups(context.Background()).ApplicationGroups(applicationGroup).Execute()
	if err == nil {
		t.Logf("Created test application group %s with ID %s", createdAppGroupName, createRes.Id)
	}
	require.NoError(t, err, "Failed to create application group for get test")
	createdAppGroupID := createRes.Id

	// Test Get by ID operation.
	reqGetById := client.ApplicationGroupsAPI.GetApplicationGroupsByID(context.Background(), createdAppGroupID)
	getRes, httpResGet, errGet := reqGetById.Execute()
	if errGet == nil {
		t.Logf("Retrieved test application group %s with ID %s", createdAppGroupName, createdAppGroupID)
	}
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful.
	require.NoError(t, errGet, "Failed to get application group by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdAppGroupName, getRes.Name, "Application group name should match")
	assert.Equal(t, createdAppGroupID, getRes.Id, "Application group ID should match")

	// Cleanup.
	delReq := client.ApplicationGroupsAPI.DeleteApplicationGroupsByID(context.Background(), createdAppGroupID)
	_, errDel := delReq.Execute()
	if errDel == nil {
		t.Logf("Deleted test application group %s with ID %s", createdAppGroupName, createdAppGroupID)
	}
	require.NoError(t, errDel)
	deleteTestApplication(t, client, app1ID, app1Name)
}

// Test_objects_ApplicationGroupsAPIService_Update tests updating an existing application group.
func Test_objects_ApplicationGroupsAPIService_Update(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create initial and new member applications.
	randomSuffix := common.GenerateRandomString(6)
	app1Name := "test-app-update-1-" + randomSuffix
	app2Name := "test-app-update-2-" + randomSuffix
	app1ID := createTestApplication(t, client, app1Name)
	app2ID := createTestApplication(t, client, app2Name)

	// Create an application group to update.
	createdAppGroupName := "test-app-group-update-" + randomSuffix
	applicationGroup := objects.ApplicationGroups{
		Folder:  common.StringPtr("Shared"),
		Name:    createdAppGroupName,
		Members: []string{app1Name},
	}
	createRes, _, err := client.ApplicationGroupsAPI.CreateApplicationGroups(context.Background()).ApplicationGroups(applicationGroup).Execute()
	if err == nil {
		t.Logf("Successfully created application group: %s with ID: %s", applicationGroup.Name, createRes.Id)
	}
	require.NoError(t, err, "Failed to create application group for update test")
	createdAppGroupID := createRes.Id

	// Test Update operation with a new member list.
	updatedGroup := objects.ApplicationGroups{
		Name:    createdAppGroupName,
		Members: []string{app1Name, app2Name},
	}
	reqUpdate := client.ApplicationGroupsAPI.UpdateApplicationGroupsByID(context.Background(), createdAppGroupID).ApplicationGroups(updatedGroup)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate == nil {
		t.Logf("Successfully updated application group: %s with ID: %s", createdAppGroupName, createdAppGroupID)
	}
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation.
	require.NoError(t, errUpdate, "Failed to update application group")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.ElementsMatch(t, []string{app1Name, app2Name}, updateRes.Members, "Members should be updated")

	// Cleanup.
	delReq := client.ApplicationGroupsAPI.DeleteApplicationGroupsByID(context.Background(), createdAppGroupID)
	_, errDel := delReq.Execute()
	if errDel == nil {
		t.Logf("Successfully deleted application group: %s with ID: %s", createdAppGroupName, createdAppGroupID)
	}
	require.NoError(t, errDel)
	deleteTestApplication(t, client, app1ID, app1Name)
	deleteTestApplication(t, client, app2ID, app2Name)
}

// Test_objects_ApplicationGroupsAPIService_List tests listing application groups.
func Test_objects_ApplicationGroupsAPIService_List(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a test application and group to find in the list.
	randomSuffix := common.GenerateRandomString(6)
	app1Name := "test-app-list-1-" + randomSuffix
	app1ID := createTestApplication(t, client, app1Name)
	// Log the app1ID
	t.Logf("Successfully created application with id and name: %s, %s", app1ID, app1Name)

	createdAppGroupName := "test-app-group-list-" + randomSuffix
	applicationGroup := objects.ApplicationGroups{
		Folder:  common.StringPtr("Shared"),
		Name:    createdAppGroupName,
		Members: []string{app1Name},
	}
	createRes, _, err := client.ApplicationGroupsAPI.CreateApplicationGroups(context.Background()).ApplicationGroups(applicationGroup).Execute()
	if err == nil {
		t.Logf("Successfully created application group: %s with ID: %s", createdAppGroupName, createRes.Id)
	}
	require.NoError(t, err, "Failed to create application group for list test")
	if err != nil {
		require.NoError(t, err, "Failed to create application group for list test")
		return
	}
	createdAppGroupID := createRes.Id

	// Test List operation.
	reqList := client.ApplicationGroupsAPI.ListApplicationGroups(context.Background()).Folder("Shared")
	listRes, httpResList, errList := reqList.Execute()
	if errList == nil {
		t.Logf("Successfully listed application groups")
	}
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation.
	require.NoError(t, errList, "Failed to list application groups")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one application in the list")

	// Verify our created group is in the list.
	foundGroup := false
	for _, group := range listRes.Data {
		if group.Name == createdAppGroupName {
			foundGroup = true
			break
		}
	}
	assert.True(t, foundGroup, "Created application group should be found in the list")

	t.Logf("Successfully listed application groups, found created application: %s %t", createdAppGroupName, foundGroup)

	// Cleanup.
	reqDel := client.ApplicationGroupsAPI.DeleteApplicationGroupsByID(context.Background(), createdAppGroupID)
	httpResDel, errDel := reqDel.Execute()
	if errDel == nil {
		t.Logf("Successfully deleted application group: %s with ID: %s", createdAppGroupName, createdAppGroupID)
	}
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete application group during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
}

// Test_objects_ApplicationGroupsAPIService_DeleteByID tests deleting an application group.
func Test_objects_ApplicationGroupsAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a test application and group to delete.
	randomSuffix := common.GenerateRandomString(6)
	app1Name := "test-app-delete-1-" + randomSuffix
	app1ID := createTestApplication(t, client, app1Name)
	createdAppGroupName := "test-app-group-delete-" + randomSuffix
	applicationGroup := objects.ApplicationGroups{
		Folder:  common.StringPtr("Shared"),
		Name:    createdAppGroupName,
		Members: []string{app1Name},
	}
	createRes, _, err := client.ApplicationGroupsAPI.CreateApplicationGroups(context.Background()).ApplicationGroups(applicationGroup).Execute()
	if err == nil {
		t.Logf("Successfully created application group: %s with ID: %s", createdAppGroupName, createRes.Id)
	}
	require.NoError(t, err, "Failed to create application group for delete test")
	createdAppGroupID := createRes.Id

	// Test Delete operation.
	reqDel := client.ApplicationGroupsAPI.DeleteApplicationGroupsByID(context.Background(), createdAppGroupID)
	httpResDel, errDel := reqDel.Execute()
	if errDel == nil {
		t.Logf("Successfully deleted application group: %s with ID: %s", createdAppGroupName, createdAppGroupID)
	}

	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation.
	require.NoError(t, errDel, "Failed to delete application group")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted application: %s", createdAppGroupID)

	// Verify deletion by trying to get the group (should fail).
	reqGetById := client.ApplicationGroupsAPI.GetApplicationGroupsByID(context.Background(), createdAppGroupID)
	getRes, httpResGet, errGet := reqGetById.Execute()
	if errGet == nil {
		t.Logf("Successfully retrieved application group: %s with ID: %s", createdAppGroupName, createdAppGroupID)
	}

	// We expect this to fail since the application group was deleted.
	assert.Error(t, errGet, "Getting deleted application group should fail")
	if httpResGet != nil {
		assert.NotEqual(t, 200, httpResGet.StatusCode, "Should not return 200 for deleted group")
	}
	assert.Nil(t, getRes, "Response should be nil for a deleted application group")

	t.Logf("Verified application group deletion: %s", createdAppGroupID)

	// Cleanup the test application.
	deleteTestApplication(t, client, app1ID, app1Name)
}
