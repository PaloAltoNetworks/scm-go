/*
* Security Services Testing
* ProfileGroupsAPIService
 */
package security_services

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/security_services"
)

// Test_security_services_ProfileGroupsAPIService_Create tests the creation of a profilegroup object
// This test creates a new profilegroup and then deletes it to ensure proper cleanup
func Test_security_services_ProfileGroupsAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a valid profilegroup object with unique name to avoid conflicts
	createdProfileGroupName := "test-" + common.GenerateRandomString(10)
	profilegroup := security_services.ProfileGroups{
		Folder: common.StringPtr("All"), // Using All folder scope
		Name:   createdProfileGroupName, // Unique test name
	}

	fmt.Printf("Creating profilegroup with name: %s\n", profilegroup.Name)

	// Make the create request to the API
	req := client.ProfileGroupsAPI.CreateProfileGroups(context.Background()).ProfileGroups(profilegroup)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create profilegroup")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileGroupName, res.Name, "Created profilegroup name should match")
	assert.NotEmpty(t, res.Id, "Created profilegroup should have an ID")

	// Use the ID from the response object
	createdProfileGroupID := res.Id
	t.Logf("Successfully created profilegroup: %s with ID: %s", profilegroup.Name, *createdProfileGroupID)

	// Cleanup: Delete the created profilegroup to maintain test isolation
	reqDel := client.ProfileGroupsAPI.DeleteProfileGroupsByID(context.Background(), *createdProfileGroupID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete profilegroup during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up profilegroup: %s", *createdProfileGroupID)
}

// Test_security_services_ProfileGroupsAPIService_GetByID tests retrieving a profilegroup by its ID
// This test creates a profilegroup, retrieves it by ID, then deletes it
func Test_security_services_ProfileGroupsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a profilegroup first to have something to retrieve
	createdProfileGroupName := "test-getbyid-" + common.GenerateRandomString(10)
	profilegroup := security_services.ProfileGroups{
		Folder: common.StringPtr("All"), // Using All folder scope
		Name:   createdProfileGroupName, // Unique test name
	}

	// Create the profilegroup via API
	req := client.ProfileGroupsAPI.CreateProfileGroups(context.Background()).ProfileGroups(profilegroup)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create profilegroup for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdProfileGroupID := createRes.Id
	require.NotEmpty(t, createdProfileGroupID, "Created profilegroup should have an ID")

	// Test Get by ID operation
	reqGetById := client.ProfileGroupsAPI.GetProfileGroupsByID(context.Background(), *createdProfileGroupID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get profilegroup by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdProfileGroupName, getRes.Name, "ProfileGroup name should match")
	assert.Equal(t, *createdProfileGroupID, *getRes.Id, "ProfileGroup ID should match")

	t.Logf("Successfully retrieved profilegroup: %s", getRes.Name)

	// Cleanup: Delete the created profilegroup
	reqDel := client.ProfileGroupsAPI.DeleteProfileGroupsByID(context.Background(), *createdProfileGroupID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete profilegroup during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up profilegroup: %s", *createdProfileGroupID)
}

// Test_security_services_ProfileGroupsAPIService_Update tests updating an existing profilegroup
// This test creates a profilegroup, updates it, then deletes it
func Test_security_services_ProfileGroupsAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a profilegroup first to have something to update
	createdProfileGroupName := "test-update-" + common.GenerateRandomString(10)
	profilegroup := security_services.ProfileGroups{
		Folder: common.StringPtr("All"), // Using All folder scope
		Name:   createdProfileGroupName, // Unique test name
	}

	// Create the profilegroup via API
	req := client.ProfileGroupsAPI.CreateProfileGroups(context.Background()).ProfileGroups(profilegroup)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create profilegroup for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdProfileGroupID := createRes.Id
	require.NotEmpty(t, createdProfileGroupID, "Created profilegroup should have an ID")

	// Test Update operation with modified fields
	updatedProfileGroup := security_services.ProfileGroups{
		Folder:  common.StringPtr("All"),   // Keep same folder scope
		Name:    createdProfileGroupName,   // Keep same name (required for update)
		Spyware: []string{"best-practice"}, // Add a spyware profile reference
	}

	reqUpdate := client.ProfileGroupsAPI.UpdateProfileGroupsByID(context.Background(), *createdProfileGroupID).ProfileGroups(updatedProfileGroup)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update profilegroup")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdProfileGroupName, updateRes.Name, "ProfileGroup name should remain the same")
	assert.Equal(t, *createdProfileGroupID, *updateRes.Id, "ProfileGroup ID should remain the same")
	assert.NotNil(t, updateRes.Spyware, "Spyware should be set")
	if updateRes.Spyware != nil && len(updateRes.Spyware) > 0 {
		assert.Equal(t, "best-practice", updateRes.Spyware[0], "Spyware profile should be updated")
	}

	t.Logf("Successfully updated profilegroup: %s", createdProfileGroupName)

	// Cleanup: Delete the created profilegroup
	reqDel := client.ProfileGroupsAPI.DeleteProfileGroupsByID(context.Background(), *createdProfileGroupID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete profilegroup during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up profilegroup: %s", *createdProfileGroupID)
}

// Test_security_services_ProfileGroupsAPIService_List tests listing profilegroups with folder filter
// This test creates a profilegroup, lists profilegroups to verify it's included, then deletes it
func Test_security_services_ProfileGroupsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a profilegroup first to have something to list
	createdProfileGroupName := "test-list-" + common.GenerateRandomString(10)
	profilegroup := security_services.ProfileGroups{
		Folder: common.StringPtr("All"), // Using All folder scope
		Name:   createdProfileGroupName, // Unique test name
	}

	// Create the profilegroup via API
	req := client.ProfileGroupsAPI.CreateProfileGroups(context.Background()).ProfileGroups(profilegroup)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create profilegroup for list test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdProfileGroupID := createRes.Id
	require.NotEmpty(t, createdProfileGroupID, "Created profilegroup should have an ID")

	// Test List operation with folder filter
	reqList := client.ProfileGroupsAPI.ListProfileGroups(context.Background()).Folder("All").Limit(200).Offset(0)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list profilegroups")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one profilegroup in the list")

	// Verify our created profilegroup is in the list
	foundProfileGroup := false
	for _, profile := range listRes.Data {
		if profile.Name == createdProfileGroupName {
			foundProfileGroup = true
			break
		}
	}
	assert.True(t, foundProfileGroup, "Created profilegroup should be found in the list")

	t.Logf("Successfully listed profilegroups, found created profilegroup: %s", createdProfileGroupName)

	// Cleanup: Delete the created profilegroup
	reqDel := client.ProfileGroupsAPI.DeleteProfileGroupsByID(context.Background(), *createdProfileGroupID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete profilegroup during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up profilegroup: %s", *createdProfileGroupID)
}

// Test_security_services_ProfileGroupsAPIService_DeleteByID tests deleting a profilegroup by its ID
// This test creates a profilegroup, deletes it, then verifies the deletion was successful
func Test_security_services_ProfileGroupsAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a profilegroup first to have something to delete
	createdProfileGroupName := "test-delete-" + common.GenerateRandomString(10)
	profilegroup := security_services.ProfileGroups{
		Folder: common.StringPtr("All"), // Using All folder scope
		Name:   createdProfileGroupName, // Unique test name
	}

	// Create the profilegroup via API
	req := client.ProfileGroupsAPI.CreateProfileGroups(context.Background()).ProfileGroups(profilegroup)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create profilegroup for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdProfileGroupID := createRes.Id
	require.NotEmpty(t, createdProfileGroupID, "Created profilegroup should have an ID")

	// Small delay for eventual consistency
	time.Sleep(2 * time.Second)

	// Test Delete by ID operation
	reqDel := client.ProfileGroupsAPI.DeleteProfileGroupsByID(context.Background(), *createdProfileGroupID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete profilegroup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted profilegroup: %s", *createdProfileGroupID)
}

// Test_security_services_ProfileGroupsAPIService_Fetch tests the fetch convenience method
func Test_security_services_ProfileGroupsAPIService_Fetch(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	// Create a profile group to fetch
	profileGroupName := "test-fetch-" + common.GenerateRandomString(10)
	profileGroup := security_services.ProfileGroups{
		Folder: common.StringPtr("All"),
		Name:   profileGroupName,
	}

	reqCreate := client.ProfileGroupsAPI.CreateProfileGroups(context.Background()).ProfileGroups(profileGroup)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile group for fetch test")
	createdID := createRes.Id
	createdFolder := createRes.Folder
	require.NotEmpty(t, createdID, "Created profile group ID should not be empty")

	// Defer cleanup
	defer func() {
		t.Logf("Cleaning up profile group with ID: %s", *createdID)
		_, errDel := client.ProfileGroupsAPI.DeleteProfileGroupsByID(context.Background(), *createdID).Execute()
		require.NoError(t, errDel, "Failed to delete profile group during cleanup")
	}()

	// Test Fetch by name operation
	fmt.Printf("Attempting to fetch profile group with name: %s\n", profileGroupName)
	fetchedProfileGroup, errFetch := client.ProfileGroupsAPI.FetchProfileGroups(context.Background(), profileGroupName, createdFolder, nil, nil)

	// Verify the fetch operation was successful
	require.NoError(t, errFetch, "Failed to fetch profile group by name")
	require.NotNil(t, fetchedProfileGroup, "Fetched profile group should not be nil")
	assert.Equal(t, profileGroupName, fetchedProfileGroup.Name, "ProfileGroup name should match")
	assert.Equal(t, *createdID, *fetchedProfileGroup.Id, "ProfileGroup ID should match")
	assert.Equal(t, *createdFolder, *fetchedProfileGroup.Folder, "Folder should match")
	t.Logf("Successfully fetched profile group: %s", profileGroupName)

	// Test fetching non-existent profile group (should return nil)
	nonExistentName := "non-existent-profilegroup-xyz-12345"
	notFoundProfileGroup, errNotFound := client.ProfileGroupsAPI.FetchProfileGroups(context.Background(), nonExistentName, createdFolder, nil, nil)
	require.NoError(t, errNotFound, "Fetch for non-existent profile group should not error")
	assert.Nil(t, notFoundProfileGroup, "Non-existent profile group should return nil")
	t.Logf("Successfully verified fetch returns nil for non-existent profile group")
}
