/*
Objects Testing DynamicUserGroupsAPIService
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

// Helper function to create a test tag object.
func createTestTag(t *testing.T, client *objects.APIClient, name string, color string) string {
	tag := objects.Tags{
		Folder: common.StringPtr("Shared"),
		Name:   name,
		Color:  common.StringPtr(color),
	}
	req := client.TagsAPI.CreateTags(context.Background()).Tags(tag)
	res, _, err := req.Execute()
	if err == nil {
		t.Logf("Successfully created test tag %s with ID %s", name, *res.Id)
	}
	require.NoError(t, err, "Failed to create test tag")
	require.NotNil(t, res, "Test tag create response should not be nil")
	return *res.Id
}

// Test_objects_DynamicUserGroupsAPIService_Create tests the creation of a dynamic user group object.
func Test_objects_DynamicUserGroupsAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a test tag first.
	randomSuffix := common.GenerateRandomString(6)
	tagName := "tag-for-dug-" + randomSuffix
	tagID := createTestTag(t, client, tagName, "Red")

	// Create a valid dynamic user group object with a unique name.
	createdDUGName := "test-dug-create-" + randomSuffix
	dynamicUserGroup := objects.DynamicUserGroups{
		Folder:      common.StringPtr("Shared"),
		Name:        createdDUGName,
		Description: common.StringPtr("Test DUG for create API"),
		Filter:      fmt.Sprintf("'Microsoft 365 Access' and '%s'", tagName),
	}

	fmt.Printf("Creating dynamic user group with name: %s\n", dynamicUserGroup.Name)

	// Make the create request to the API.
	req := client.DynamicUserGroupsAPI.CreateDynamicUserGroups(context.Background()).DynamicUserGroups(dynamicUserGroup)
	res, httpRes, err := req.Execute()
	if err == nil {
		t.Logf("Successfully created dynamic user group: %s with ID: %s", dynamicUserGroup.Name, res.Id)
	}
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create dynamic user group")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties.
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdDUGName, res.Name, "Created DUG name should match")
	assert.Equal(t, fmt.Sprintf("'Microsoft 365 Access' and '%s'", tagName), res.Filter, "Filter should match")
	assert.NotEmpty(t, res.Id, "Created DUG should have an ID")

	createdDUGID := res.Id

	// Cleanup: Delete the created DUG and test tag.
	delReq := client.DynamicUserGroupsAPI.DeleteDynamicUserGroupsByID(context.Background(), createdDUGID)
	_, errDel := delReq.Execute()
	if errDel == nil {
		t.Logf("Successfully deleted dynamic user group: %s", createdDUGID)
	}
	require.NoError(t, errDel, "Failed to delete dynamic user group during cleanup")
	deleteTestTag(t, client, tagID, tagName)
}

// Test_objects_DynamicUserGroupsAPIService_GetByID tests retrieving a dynamic user group by its ID.
func Test_objects_DynamicUserGroupsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a test tag and DUG to retrieve.
	randomSuffix := common.GenerateRandomString(6)
	tagName := "tag-for-dug-get-" + randomSuffix
	tagID := createTestTag(t, client, tagName, "Blue")
	createdDUGName := "test-dug-get-" + randomSuffix
	dynamicUserGroup := objects.DynamicUserGroups{
		Folder: common.StringPtr("Shared"),
		Name:   createdDUGName,
		Filter: fmt.Sprintf("'%s'", tagName),
	}
	createRes, _, err := client.DynamicUserGroupsAPI.CreateDynamicUserGroups(context.Background()).DynamicUserGroups(dynamicUserGroup).Execute()
	if err == nil {
		t.Logf("Successfully created dynamic user group for get test: %s", createdDUGName)
	}
	require.NoError(t, err, "Failed to create dynamic user group for get test")
	createdDUGID := createRes.Id

	// Test Get by ID operation.
	reqGetById := client.DynamicUserGroupsAPI.GetDynamicUserGroupsByID(context.Background(), createdDUGID)
	getRes, httpResGet, errGet := reqGetById.Execute()
	if errGet == nil {
		t.Logf("Successfully retrieved dynamic user group: %s", createdDUGName)
	}
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful.
	require.NoError(t, errGet, "Failed to get dynamic user group by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdDUGName, getRes.Name, "DUG name should match")
	assert.Equal(t, createdDUGID, getRes.Id, "DUG ID should match")

	// Cleanup.
	delReq := client.DynamicUserGroupsAPI.DeleteDynamicUserGroupsByID(context.Background(), createdDUGID)
	_, errDel := delReq.Execute()
	require.NoError(t, errDel)
	deleteTestTag(t, client, tagID, tagName)
}

// Test_objects_DynamicUserGroupsAPIService_Update tests updating an existing dynamic user group.
func Test_objects_DynamicUserGroupsAPIService_Update(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create tags for initial and updated state.
	randomSuffix := common.GenerateRandomString(6)
	tag1Name := "tag-dug-update-1-" + randomSuffix
	tag2Name := "tag-dug-update-2-" + randomSuffix
	tag1ID := createTestTag(t, client, tag1Name, "Green")
	tag2ID := createTestTag(t, client, tag2Name, "Yellow")

	// Create a DUG to update.
	createdDUGName := "test-dug-update-" + randomSuffix
	dynamicUserGroup := objects.DynamicUserGroups{
		Folder: common.StringPtr("Shared"),
		Name:   createdDUGName,
		Filter: fmt.Sprintf("'%s'", tag1Name),
	}
	createRes, _, err := client.DynamicUserGroupsAPI.CreateDynamicUserGroups(context.Background()).DynamicUserGroups(dynamicUserGroup).Execute()
	if err == nil {
		t.Logf("Successfully created dynamic user group for update test: %s", createdDUGName)
	}
	require.NoError(t, err, "Failed to create dynamic user group for update test")
	createdDUGID := createRes.Id

	// Test Update operation with a new filter.
	updatedGroup := objects.DynamicUserGroups{
		Name:   createdDUGName,
		Filter: fmt.Sprintf("'%s' or '%s'", tag1Name, tag2Name),
	}
	reqUpdate := client.DynamicUserGroupsAPI.UpdateDynamicUserGroupsByID(context.Background(), createdDUGID).DynamicUserGroups(updatedGroup)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate == nil {
		t.Logf("Successfully updated dynamic user group: %s", createdDUGName)
	}
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation.
	require.NoError(t, errUpdate, "Failed to update dynamic user group")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, fmt.Sprintf("'%s' or '%s'", tag1Name, tag2Name), updateRes.Filter, "Filter should be updated")

	// Cleanup.
	delReq := client.DynamicUserGroupsAPI.DeleteDynamicUserGroupsByID(context.Background(), createdDUGID)
	_, errDel := delReq.Execute()
	require.NoError(t, errDel)
	deleteTestTag(t, client, tag1ID, tag1Name)
	deleteTestTag(t, client, tag2ID, tag2Name)
}

// Test_objects_DynamicUserGroupsAPIService_List tests listing dynamic user groups.
func Test_objects_DynamicUserGroupsAPIService_List(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a test tag and DUG to find in the list.
	randomSuffix := common.GenerateRandomString(6)
	tagName := "tag-dug-list-" + randomSuffix
	tagID := createTestTag(t, client, tagName, "Purple")
	createdDUGName := "test-dug-list-" + randomSuffix
	dynamicUserGroup := objects.DynamicUserGroups{
		Folder: common.StringPtr("Shared"),
		Name:   createdDUGName,
		Filter: fmt.Sprintf("'%s'", tagName),
	}
	createRes, _, err := client.DynamicUserGroupsAPI.CreateDynamicUserGroups(context.Background()).DynamicUserGroups(dynamicUserGroup).Execute()
	if err == nil {
		t.Logf("Successfully created dynamic user group for list test: %s", createdDUGName)
	}
	require.NoError(t, err, "Failed to create dynamic user group for list test")
	createdDUGID := createRes.Id

	// Test List operation.
	reqList := client.DynamicUserGroupsAPI.ListDynamicUserGroups(context.Background()).Folder("Shared")
	listRes, httpResList, errList := reqList.Execute()
	if errList == nil {
		t.Logf("Successfully listed dynamic user groups")
	}
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation.
	require.NoError(t, errList, "Failed to list dynamic user groups")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	assert.NotNil(t, listRes.Data)

	// Verify our created group is in the list.
	foundGroup := false
	for _, group := range listRes.Data {
		if group.Name == createdDUGName {
			foundGroup = true
			break
		}
	}
	assert.True(t, foundGroup, "Created dynamic user group should be found in the list")

	// Cleanup.
	delReq := client.DynamicUserGroupsAPI.DeleteDynamicUserGroupsByID(context.Background(), createdDUGID)
	_, errDel := delReq.Execute()
	require.NoError(t, errDel)
	deleteTestTag(t, client, tagID, tagName)
}

// Test_objects_DynamicUserGroupsAPIService_DeleteByID tests deleting a dynamic user group.
func Test_objects_DynamicUserGroupsAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a test tag and DUG to delete.
	randomSuffix := common.GenerateRandomString(6)
	tagName := "tag-dug-delete-" + randomSuffix
	tagID := createTestTag(t, client, tagName, "Orange")
	createdDUGName := "test-dug-delete-" + randomSuffix
	dynamicUserGroup := objects.DynamicUserGroups{
		Folder: common.StringPtr("Shared"),
		Name:   createdDUGName,
		Filter: fmt.Sprintf("'%s'", tagName),
	}
	createRes, _, err := client.DynamicUserGroupsAPI.CreateDynamicUserGroups(context.Background()).DynamicUserGroups(dynamicUserGroup).Execute()
	if err == nil {
		t.Logf("Successfully created dynamic user group for delete test: %s", createdDUGName)
	}
	require.NoError(t, err, "Failed to create dynamic user group for delete test")
	createdDUGID := createRes.Id

	// Test Delete operation.
	reqDel := client.DynamicUserGroupsAPI.DeleteDynamicUserGroupsByID(context.Background(), createdDUGID)
	httpResDel, errDel := reqDel.Execute()
	if errDel == nil {
		t.Logf("Successfully deleted dynamic user group: %s", createdDUGID)
	}
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation.
	require.NoError(t, errDel, "Failed to delete dynamic user group")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	// Cleanup the test tag.
	deleteTestTag(t, client, tagID, tagName)
}
