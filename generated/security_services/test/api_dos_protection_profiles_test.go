/*
* Security Services Testing
* DoSProtectionProfilesAPIService
 */
package security_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/security_services"
)

// Test_security_services_DoSProtectionProfilesAPIService_Create tests the creation of a dosprotectionprofile object
// This test creates a new dosprotectionprofile and then deletes it to ensure proper cleanup
func Test_security_services_DoSProtectionProfilesAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a valid dosprotectionprofile object with unique name to avoid conflicts
	createdDoSProtectionProfileName := "test-" + common.GenerateRandomString(10)
	dosprotectionprofile := security_services.DosProtectionProfiles{
		Description: common.StringPtr("Test DoS protection profile for create API testing"),
		Folder:      common.StringPtr("All"),         // Using All folder scope
		Name:        createdDoSProtectionProfileName, // Unique test name
		Type:        "aggregate",                     // Required field - aggregate or classified
	}

	fmt.Printf("Creating dosprotectionprofile with name: %s\n", dosprotectionprofile.Name)

	// Make the create request to the API
	req := client.DoSProtectionProfilesAPI.CreateDoSProtectionProfiles(context.Background()).DosProtectionProfiles(dosprotectionprofile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create dosprotectionprofile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdDoSProtectionProfileName, res.Name, "Created dosprotectionprofile name should match")
	assert.Equal(t, common.StringPtr("Test DoS protection profile for create API testing"), res.Description, "Description should match")
	assert.Equal(t, "aggregate", res.Type, "Type should match")
	assert.NotEmpty(t, res.Id, "Created dosprotectionprofile should have an ID")

	// Use the ID from the response object
	createdDoSProtectionProfileID := res.Id
	t.Logf("Successfully created dosprotectionprofile: %s with ID: %s", dosprotectionprofile.Name, *createdDoSProtectionProfileID)

	// Cleanup: Delete the created dosprotectionprofile to maintain test isolation
	reqDel := client.DoSProtectionProfilesAPI.DeleteDoSProtectionProfilesByID(context.Background(), *createdDoSProtectionProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete dosprotectionprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up dosprotectionprofile: %s", *createdDoSProtectionProfileID)
}

// Test_security_services_DoSProtectionProfilesAPIService_GetByID tests retrieving a dosprotectionprofile by its ID
// This test creates a dosprotectionprofile, retrieves it by ID, then deletes it
func Test_security_services_DoSProtectionProfilesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a dosprotectionprofile first to have something to retrieve
	createdDoSProtectionProfileName := "test-getbyid-" + common.GenerateRandomString(10)
	dosprotectionprofile := security_services.DosProtectionProfiles{
		Description: common.StringPtr("Test DoS protection profile for get by ID API testing"),
		Folder:      common.StringPtr("All"),         // Using All folder scope
		Name:        createdDoSProtectionProfileName, // Unique test name
		Type:        "aggregate",                     // Required field
	}

	// Create the dosprotectionprofile via API
	req := client.DoSProtectionProfilesAPI.CreateDoSProtectionProfiles(context.Background()).DosProtectionProfiles(dosprotectionprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create dosprotectionprofile for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDoSProtectionProfileID := createRes.Id
	require.NotEmpty(t, createdDoSProtectionProfileID, "Created dosprotectionprofile should have an ID")

	// Test Get by ID operation
	reqGetById := client.DoSProtectionProfilesAPI.GetDoSProtectionProfilesByID(context.Background(), *createdDoSProtectionProfileID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get dosprotectionprofile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdDoSProtectionProfileName, getRes.Name, "DoSProtectionProfile name should match")
	assert.Equal(t, common.StringPtr("Test DoS protection profile for get by ID API testing"), getRes.Description, "Description should match")
	assert.Equal(t, "aggregate", getRes.Type, "Type should match")
	assert.Equal(t, *createdDoSProtectionProfileID, *getRes.Id, "DoSProtectionProfile ID should match")

	t.Logf("Successfully retrieved dosprotectionprofile: %s", getRes.Name)

	// Cleanup: Delete the created dosprotectionprofile
	reqDel := client.DoSProtectionProfilesAPI.DeleteDoSProtectionProfilesByID(context.Background(), *createdDoSProtectionProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete dosprotectionprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up dosprotectionprofile: %s", *createdDoSProtectionProfileID)
}

// Test_security_services_DoSProtectionProfilesAPIService_Update tests updating an existing dosprotectionprofile
// This test creates a dosprotectionprofile, updates it, then deletes it
func Test_security_services_DoSProtectionProfilesAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a dosprotectionprofile first to have something to update
	createdDoSProtectionProfileName := "test-update-" + common.GenerateRandomString(10)
	dosprotectionprofile := security_services.DosProtectionProfiles{
		Description: common.StringPtr("Test DoS protection profile for update API testing"),
		Folder:      common.StringPtr("All"),         // Using All folder scope
		Name:        createdDoSProtectionProfileName, // Unique test name
		Type:        "aggregate",                     // Required field
	}

	// Create the dosprotectionprofile via API
	req := client.DoSProtectionProfilesAPI.CreateDoSProtectionProfiles(context.Background()).DosProtectionProfiles(dosprotectionprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create dosprotectionprofile for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDoSProtectionProfileID := createRes.Id
	require.NotEmpty(t, createdDoSProtectionProfileID, "Created dosprotectionprofile should have an ID")

	// Test Update operation with modified fields
	updatedDoSProtectionProfile := security_services.DosProtectionProfiles{
		Description: common.StringPtr("Updated test DoS protection profile description"), // Updated description
		Folder:      common.StringPtr("All"),                                             // Keep same folder scope
		Name:        createdDoSProtectionProfileName,                                     // Keep same name (required for update)
		Type:        "aggregate",                                                         // Keep same type (required)
	}

	reqUpdate := client.DoSProtectionProfilesAPI.UpdateDoSProtectionProfilesByID(context.Background(), *createdDoSProtectionProfileID).DosProtectionProfiles(updatedDoSProtectionProfile)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update dosprotectionprofile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdDoSProtectionProfileName, updateRes.Name, "DoSProtectionProfile name should remain the same")
	assert.Equal(t, common.StringPtr("Updated test DoS protection profile description"), updateRes.Description, "Description should be updated")
	assert.Equal(t, "aggregate", updateRes.Type, "Type should remain the same")
	assert.Equal(t, *createdDoSProtectionProfileID, *updateRes.Id, "DoSProtectionProfile ID should remain the same")

	t.Logf("Successfully updated dosprotectionprofile: %s", createdDoSProtectionProfileName)

	// Cleanup: Delete the created dosprotectionprofile
	reqDel := client.DoSProtectionProfilesAPI.DeleteDoSProtectionProfilesByID(context.Background(), *createdDoSProtectionProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete dosprotectionprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up dosprotectionprofile: %s", *createdDoSProtectionProfileID)
}

// Test_security_services_DoSProtectionProfilesAPIService_List tests listing dosprotectionprofiles with folder filter
// This test creates a dosprotectionprofile, lists dosprotectionprofiles to verify it's included, then deletes it
func Test_security_services_DoSProtectionProfilesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a dosprotectionprofile first to have something to list
	createdDoSProtectionProfileName := "test-list-" + common.GenerateRandomString(10)
	dosprotectionprofile := security_services.DosProtectionProfiles{
		Description: common.StringPtr("Test DoS protection profile for list API testing"),
		Folder:      common.StringPtr("All"),         // Using All folder scope
		Name:        createdDoSProtectionProfileName, // Unique test name
		Type:        "aggregate",                     // Required field
	}

	// Create the dosprotectionprofile via API
	req := client.DoSProtectionProfilesAPI.CreateDoSProtectionProfiles(context.Background()).DosProtectionProfiles(dosprotectionprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create dosprotectionprofile for list test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDoSProtectionProfileID := createRes.Id
	require.NotEmpty(t, createdDoSProtectionProfileID, "Created dosprotectionprofile should have an ID")

	// Test List operation with folder filter
	reqList := client.DoSProtectionProfilesAPI.ListDoSProtectionProfiles(context.Background()).Folder("All").Limit(200).Offset(0)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list dosprotectionprofiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one dosprotectionprofile in the list")

	// Verify our created dosprotectionprofile is in the list
	foundDoSProtectionProfile := false
	for _, profile := range listRes.Data {
		if profile.Name == createdDoSProtectionProfileName {
			foundDoSProtectionProfile = true
			assert.Equal(t, common.StringPtr("Test DoS protection profile for list API testing"), profile.Description, "Description should match")
			assert.Equal(t, "aggregate", profile.Type, "Type should match")
			break
		}
	}
	assert.True(t, foundDoSProtectionProfile, "Created dosprotectionprofile should be found in the list")

	t.Logf("Successfully listed dosprotectionprofiles, found created dosprotectionprofile: %s", createdDoSProtectionProfileName)

	// Cleanup: Delete the created dosprotectionprofile
	reqDel := client.DoSProtectionProfilesAPI.DeleteDoSProtectionProfilesByID(context.Background(), *createdDoSProtectionProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete dosprotectionprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up dosprotectionprofile: %s", *createdDoSProtectionProfileID)
}

// Test_security_services_DoSProtectionProfilesAPIService_DeleteByID tests deleting a dosprotectionprofile by its ID
// This test creates a dosprotectionprofile, deletes it, then verifies the deletion was successful
func Test_security_services_DoSProtectionProfilesAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a dosprotectionprofile first to have something to delete
	createdDoSProtectionProfileName := "test-delete-" + common.GenerateRandomString(10)
	dosprotectionprofile := security_services.DosProtectionProfiles{
		Description: common.StringPtr("Test DoS protection profile for delete API testing"),
		Folder:      common.StringPtr("All"),         // Using All folder scope
		Name:        createdDoSProtectionProfileName, // Unique test name
		Type:        "aggregate",                     // Required field
	}

	// Create the dosprotectionprofile via API
	req := client.DoSProtectionProfilesAPI.CreateDoSProtectionProfiles(context.Background()).DosProtectionProfiles(dosprotectionprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create dosprotectionprofile for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDoSProtectionProfileID := createRes.Id
	require.NotEmpty(t, createdDoSProtectionProfileID, "Created dosprotectionprofile should have an ID")

	// Test Delete by ID operation
	reqDel := client.DoSProtectionProfilesAPI.DeleteDoSProtectionProfilesByID(context.Background(), *createdDoSProtectionProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete dosprotectionprofile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted dosprotectionprofile: %s", *createdDoSProtectionProfileID)
}

// Test_security_services_DoSProtectionProfilesAPIService_Fetch tests the fetch convenience method
func Test_security_services_DoSProtectionProfilesAPIService_Fetch(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	// Create a DoS protection profile to fetch
	profileName := "test-fetch-" + common.GenerateRandomString(10)
	profile := security_services.DosProtectionProfiles{
		Description: common.StringPtr("Test DoS protection profile for fetch"),
		Folder:      common.StringPtr("All"),
		Name:        profileName,
		Type:        "aggregate",
	}

	reqCreate := client.DoSProtectionProfilesAPI.CreateDoSProtectionProfiles(context.Background()).DosProtectionProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create DoS protection profile for fetch test")
	createdID := createRes.Id
	createdFolder := createRes.Folder
	require.NotEmpty(t, createdID, "Created DoS protection profile ID should not be empty")

	// Defer cleanup
	defer func() {
		t.Logf("Cleaning up DoS protection profile with ID: %s", *createdID)
		_, errDel := client.DoSProtectionProfilesAPI.DeleteDoSProtectionProfilesByID(context.Background(), *createdID).Execute()
		require.NoError(t, errDel, "Failed to delete DoS protection profile during cleanup")
	}()

	// Test Fetch by name operation
	fmt.Printf("Attempting to fetch DoS protection profile with name: %s\n", profileName)
	fetchedProfile, errFetch := client.DoSProtectionProfilesAPI.FetchDoSProtectionProfiles(context.Background(), profileName, createdFolder, nil, nil)

	// Verify the fetch operation was successful
	require.NoError(t, errFetch, "Failed to fetch DoS protection profile by name")
	require.NotNil(t, fetchedProfile, "Fetched DoS protection profile should not be nil")
	assert.Equal(t, profileName, fetchedProfile.Name, "DoS protection profile name should match")
	assert.Equal(t, *createdID, *fetchedProfile.Id, "DoS protection profile ID should match")
	assert.Equal(t, *createdFolder, *fetchedProfile.Folder, "Folder should match")
	t.Logf("Successfully fetched DoS protection profile: %s", profileName)

	// Test fetching non-existent DoS protection profile (should return nil)
	nonExistentName := "non-existent-dosprofile-xyz-12345"
	notFoundProfile, errNotFound := client.DoSProtectionProfilesAPI.FetchDoSProtectionProfiles(context.Background(), nonExistentName, createdFolder, nil, nil)
	require.NoError(t, errNotFound, "Fetch for non-existent DoS protection profile should not error")
	assert.Nil(t, notFoundProfile, "Non-existent DoS protection profile should return nil")
	t.Logf("Successfully verified fetch returns nil for non-existent DoS protection profile")
}
