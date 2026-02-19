/*
* Security Services Testing
* WildFireAntiVirusProfilesAPIService
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

// Test_security_services_WildFireAntiVirusProfilesAPIService_Create tests the creation of a wildfireantivirusprofile object
// This test creates a new wildfireantivirusprofile and then deletes it to ensure proper cleanup
func Test_security_services_WildFireAntiVirusProfilesAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a valid wildfireantivirusprofile object with unique name to avoid conflicts
	createdWildFireAntiVirusProfileName := "test-" + common.GenerateRandomString(10)
	wildfireantivirusprofile := security_services.WildfireAntiVirusProfiles{
		Description: common.StringPtr("Test WildFire anti-virus profile for create API testing"),
		Folder:      common.StringPtr("All"),             // Using All folder scope
		Name:        createdWildFireAntiVirusProfileName, // Unique test name
	}

	fmt.Printf("Creating wildfireantivirusprofile with name: %s\n", wildfireantivirusprofile.Name)

	// Make the create request to the API
	req := client.WildFireAntiVirusProfilesAPI.CreateWildFireAntiVirusProfiles(context.Background()).WildfireAntiVirusProfiles(wildfireantivirusprofile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create wildfireantivirusprofile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdWildFireAntiVirusProfileName, res.Name, "Created wildfireantivirusprofile name should match")
	assert.Equal(t, common.StringPtr("Test WildFire anti-virus profile for create API testing"), res.Description, "Description should match")
	assert.NotEmpty(t, res.Id, "Created wildfireantivirusprofile should have an ID")

	// Use the ID from the response object
	createdWildFireAntiVirusProfileID := res.Id
	t.Logf("Successfully created wildfireantivirusprofile: %s with ID: %s", wildfireantivirusprofile.Name, *createdWildFireAntiVirusProfileID)

	// Cleanup: Delete the created wildfireantivirusprofile to maintain test isolation
	reqDel := client.WildFireAntiVirusProfilesAPI.DeleteWildFireAntiVirusProfilesByID(context.Background(), *createdWildFireAntiVirusProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete wildfireantivirusprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up wildfireantivirusprofile: %s", *createdWildFireAntiVirusProfileID)
}

// Test_security_services_WildFireAntiVirusProfilesAPIService_GetByID tests retrieving a wildfireantivirusprofile by its ID
// This test creates a wildfireantivirusprofile, retrieves it by ID, then deletes it
func Test_security_services_WildFireAntiVirusProfilesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a wildfireantivirusprofile first to have something to retrieve
	createdWildFireAntiVirusProfileName := "test-getbyid-" + common.GenerateRandomString(10)
	wildfireantivirusprofile := security_services.WildfireAntiVirusProfiles{
		Description: common.StringPtr("Test WildFire anti-virus profile for get by ID API testing"),
		Folder:      common.StringPtr("All"),             // Using All folder scope
		Name:        createdWildFireAntiVirusProfileName, // Unique test name
	}

	// Create the wildfireantivirusprofile via API
	req := client.WildFireAntiVirusProfilesAPI.CreateWildFireAntiVirusProfiles(context.Background()).WildfireAntiVirusProfiles(wildfireantivirusprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create wildfireantivirusprofile for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdWildFireAntiVirusProfileID := createRes.Id
	require.NotEmpty(t, createdWildFireAntiVirusProfileID, "Created wildfireantivirusprofile should have an ID")

	// Test Get by ID operation
	reqGetById := client.WildFireAntiVirusProfilesAPI.GetWildFireAntiVirusProfilesByID(context.Background(), *createdWildFireAntiVirusProfileID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get wildfireantivirusprofile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdWildFireAntiVirusProfileName, getRes.Name, "WildFireAntiVirusProfile name should match")
	assert.Equal(t, common.StringPtr("Test WildFire anti-virus profile for get by ID API testing"), getRes.Description, "Description should match")
	assert.Equal(t, *createdWildFireAntiVirusProfileID, *getRes.Id, "WildFireAntiVirusProfile ID should match")

	t.Logf("Successfully retrieved wildfireantivirusprofile: %s", getRes.Name)

	// Cleanup: Delete the created wildfireantivirusprofile
	reqDel := client.WildFireAntiVirusProfilesAPI.DeleteWildFireAntiVirusProfilesByID(context.Background(), *createdWildFireAntiVirusProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete wildfireantivirusprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up wildfireantivirusprofile: %s", *createdWildFireAntiVirusProfileID)
}

// Test_security_services_WildFireAntiVirusProfilesAPIService_Update tests updating an existing wildfireantivirusprofile
// This test creates a wildfireantivirusprofile, updates it, then deletes it
func Test_security_services_WildFireAntiVirusProfilesAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a wildfireantivirusprofile first to have something to update
	createdWildFireAntiVirusProfileName := "test-update-" + common.GenerateRandomString(10)
	wildfireantivirusprofile := security_services.WildfireAntiVirusProfiles{
		Description: common.StringPtr("Test WildFire anti-virus profile for update API testing"),
		Folder:      common.StringPtr("All"),             // Using All folder scope
		Name:        createdWildFireAntiVirusProfileName, // Unique test name
	}

	// Create the wildfireantivirusprofile via API
	req := client.WildFireAntiVirusProfilesAPI.CreateWildFireAntiVirusProfiles(context.Background()).WildfireAntiVirusProfiles(wildfireantivirusprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create wildfireantivirusprofile for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdWildFireAntiVirusProfileID := createRes.Id
	require.NotEmpty(t, createdWildFireAntiVirusProfileID, "Created wildfireantivirusprofile should have an ID")

	// Test Update operation with modified fields
	updatedWildFireAntiVirusProfile := security_services.WildfireAntiVirusProfiles{
		Description: common.StringPtr("Updated test WildFire anti-virus profile description"), // Updated description
		Folder:      common.StringPtr("All"),                                                  // Keep same folder scope
		Name:        createdWildFireAntiVirusProfileName,                                      // Keep same name (required for update)
	}

	reqUpdate := client.WildFireAntiVirusProfilesAPI.UpdateWildFireAntiVirusProfilesByID(context.Background(), *createdWildFireAntiVirusProfileID).WildfireAntiVirusProfiles(updatedWildFireAntiVirusProfile)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update wildfireantivirusprofile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdWildFireAntiVirusProfileName, updateRes.Name, "WildFireAntiVirusProfile name should remain the same")
	assert.Equal(t, common.StringPtr("Updated test WildFire anti-virus profile description"), updateRes.Description, "Description should be updated")
	assert.Equal(t, *createdWildFireAntiVirusProfileID, *updateRes.Id, "WildFireAntiVirusProfile ID should remain the same")

	t.Logf("Successfully updated wildfireantivirusprofile: %s", createdWildFireAntiVirusProfileName)

	// Cleanup: Delete the created wildfireantivirusprofile
	reqDel := client.WildFireAntiVirusProfilesAPI.DeleteWildFireAntiVirusProfilesByID(context.Background(), *createdWildFireAntiVirusProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete wildfireantivirusprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up wildfireantivirusprofile: %s", *createdWildFireAntiVirusProfileID)
}

// Test_security_services_WildFireAntiVirusProfilesAPIService_List tests listing wildfireantivirusprofiles with folder filter
// This test creates a wildfireantivirusprofile, lists wildfireantivirusprofiles to verify it's included, then deletes it
func Test_security_services_WildFireAntiVirusProfilesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a wildfireantivirusprofile first to have something to list
	createdWildFireAntiVirusProfileName := "test-list-" + common.GenerateRandomString(10)
	wildfireantivirusprofile := security_services.WildfireAntiVirusProfiles{
		Description: common.StringPtr("Test WildFire anti-virus profile for list API testing"),
		Folder:      common.StringPtr("All"),             // Using All folder scope
		Name:        createdWildFireAntiVirusProfileName, // Unique test name
	}

	// Create the wildfireantivirusprofile via API
	req := client.WildFireAntiVirusProfilesAPI.CreateWildFireAntiVirusProfiles(context.Background()).WildfireAntiVirusProfiles(wildfireantivirusprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create wildfireantivirusprofile for list test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdWildFireAntiVirusProfileID := createRes.Id
	require.NotEmpty(t, createdWildFireAntiVirusProfileID, "Created wildfireantivirusprofile should have an ID")

	// Test List operation with folder filter
	reqList := client.WildFireAntiVirusProfilesAPI.ListWildFireAntiVirusProfiles(context.Background()).Folder("All").Limit(200).Offset(0)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list wildfireantivirusprofiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one wildfireantivirusprofile in the list")

	// Verify our created wildfireantivirusprofile is in the list
	foundWildFireAntiVirusProfile := false
	for _, profile := range listRes.Data {
		if profile.Name == createdWildFireAntiVirusProfileName {
			foundWildFireAntiVirusProfile = true
			assert.Equal(t, common.StringPtr("Test WildFire anti-virus profile for list API testing"), profile.Description, "Description should match")
			break
		}
	}
	assert.True(t, foundWildFireAntiVirusProfile, "Created wildfireantivirusprofile should be found in the list")

	t.Logf("Successfully listed wildfireantivirusprofiles, found created wildfireantivirusprofile: %s", createdWildFireAntiVirusProfileName)

	// Cleanup: Delete the created wildfireantivirusprofile
	reqDel := client.WildFireAntiVirusProfilesAPI.DeleteWildFireAntiVirusProfilesByID(context.Background(), *createdWildFireAntiVirusProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete wildfireantivirusprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up wildfireantivirusprofile: %s", *createdWildFireAntiVirusProfileID)
}

// Test_security_services_WildFireAntiVirusProfilesAPIService_DeleteByID tests deleting a wildfireantivirusprofile by its ID
// This test creates a wildfireantivirusprofile, deletes it, then verifies the deletion was successful
func Test_security_services_WildFireAntiVirusProfilesAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a wildfireantivirusprofile first to have something to delete
	createdWildFireAntiVirusProfileName := "test-delete-" + common.GenerateRandomString(10)
	wildfireantivirusprofile := security_services.WildfireAntiVirusProfiles{
		Description: common.StringPtr("Test WildFire anti-virus profile for delete API testing"),
		Folder:      common.StringPtr("All"),             // Using All folder scope
		Name:        createdWildFireAntiVirusProfileName, // Unique test name
	}

	// Create the wildfireantivirusprofile via API
	req := client.WildFireAntiVirusProfilesAPI.CreateWildFireAntiVirusProfiles(context.Background()).WildfireAntiVirusProfiles(wildfireantivirusprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create wildfireantivirusprofile for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdWildFireAntiVirusProfileID := createRes.Id
	require.NotEmpty(t, createdWildFireAntiVirusProfileID, "Created wildfireantivirusprofile should have an ID")

	// Test Delete by ID operation
	reqDel := client.WildFireAntiVirusProfilesAPI.DeleteWildFireAntiVirusProfilesByID(context.Background(), *createdWildFireAntiVirusProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete wildfireantivirusprofile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted wildfireantivirusprofile: %s", *createdWildFireAntiVirusProfileID)
}

// Test_security_services_WildFireAntiVirusProfilesAPIService_Fetch tests the fetch convenience method
func Test_security_services_WildFireAntiVirusProfilesAPIService_Fetch(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	// Create a WildFire anti-virus profile to fetch
	profileName := "test-fetch-" + common.GenerateRandomString(10)
	profile := security_services.WildfireAntiVirusProfiles{
		Description: common.StringPtr("Test WildFire anti-virus profile for fetch"),
		Folder:      common.StringPtr("All"),
		Name:        profileName,
	}

	reqCreate := client.WildFireAntiVirusProfilesAPI.CreateWildFireAntiVirusProfiles(context.Background()).WildfireAntiVirusProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create WildFire anti-virus profile for fetch test")
	createdID := createRes.Id
	createdFolder := createRes.Folder
	require.NotEmpty(t, createdID, "Created WildFire anti-virus profile ID should not be empty")

	// Defer cleanup
	defer func() {
		t.Logf("Cleaning up WildFire anti-virus profile with ID: %s", *createdID)
		_, errDel := client.WildFireAntiVirusProfilesAPI.DeleteWildFireAntiVirusProfilesByID(context.Background(), *createdID).Execute()
		require.NoError(t, errDel, "Failed to delete WildFire anti-virus profile during cleanup")
	}()

	// Test Fetch by name operation
	fmt.Printf("Attempting to fetch WildFire anti-virus profile with name: %s\n", profileName)
	fetchedProfile, errFetch := client.WildFireAntiVirusProfilesAPI.FetchWildFireAntiVirusProfiles(context.Background(), profileName, createdFolder, nil, nil)

	// Verify the fetch operation was successful
	require.NoError(t, errFetch, "Failed to fetch WildFire anti-virus profile by name")
	require.NotNil(t, fetchedProfile, "Fetched WildFire anti-virus profile should not be nil")
	assert.Equal(t, profileName, fetchedProfile.Name, "WildFire anti-virus profile name should match")
	assert.Equal(t, *createdID, *fetchedProfile.Id, "WildFire anti-virus profile ID should match")
	assert.Equal(t, *createdFolder, *fetchedProfile.Folder, "Folder should match")
	t.Logf("Successfully fetched WildFire anti-virus profile: %s", profileName)

	// Test fetching non-existent WildFire anti-virus profile (should return nil)
	nonExistentName := "non-existent-wfavprofile-xyz-12345"
	notFoundProfile, errNotFound := client.WildFireAntiVirusProfilesAPI.FetchWildFireAntiVirusProfiles(context.Background(), nonExistentName, createdFolder, nil, nil)
	require.NoError(t, errNotFound, "Fetch for non-existent WildFire anti-virus profile should not error")
	assert.Nil(t, notFoundProfile, "Non-existent WildFire anti-virus profile should return nil")
	t.Logf("Successfully verified fetch returns nil for non-existent WildFire anti-virus profile")
}
