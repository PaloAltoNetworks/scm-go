/*
 * Security Services Testing
 *
 * URLAccessProfilesAPIService
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

// Test_security_services_URLAccessProfilesAPIService_Create tests the creation of a URL Access Profile.
func Test_security_services_URLAccessProfilesAPIService_Create(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdProfileName := "test-url-access-create-" + common.GenerateRandomString(6)
	profile := security_services.UrlAccessProfiles{
		Folder:      common.StringPtr("Shared"),
		Name:        createdProfileName,
		Description: common.StringPtr("Test URL Access Profile for create API"),
		Block:       []string{"adult", "gambling"},
		Alert:       []string{"high-risk", "phishing"},
	}

	fmt.Printf("Creating URL Access Profile with name: %s\n", profile.Name)
	req := client.URLAccessProfilesAPI.CreateURLAccessProfiles(context.Background()).UrlAccessProfiles(profile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create URL Access Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")
	createdProfileID := *res.Id

	defer func() {
		t.Logf("Cleaning up URL Access Profile with ID: %s", createdProfileID)
		_, errDel := client.URLAccessProfilesAPI.DeleteURLAccessProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete URL Access Profile during cleanup")
	}()

	t.Logf("Successfully created URL Access Profile: %s with ID: %s", profile.Name, createdProfileID)
	assert.ElementsMatch(t, []string{"adult", "gambling"}, res.Block, "Block list should match")
	assert.ElementsMatch(t, []string{"high-risk", "phishing"}, res.Alert, "Alert list should match")
}

// Test_security_services_URLAccessProfilesAPIService_GetByID tests retrieving a URL Access Profile by its ID.
func Test_security_services_URLAccessProfilesAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-url-access-get-" + common.GenerateRandomString(6)
	profile := security_services.UrlAccessProfiles{
		Folder:      common.StringPtr("Shared"),
		Name:        profileName,
		Description: common.StringPtr("Test URL Access Profile for create API"),
		Block:       []string{"adult", "gambling"},
		Alert:       []string{"high-risk", "phishing"},
	}

	createRes, _, err := client.URLAccessProfilesAPI.CreateURLAccessProfiles(context.Background()).UrlAccessProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create URL Access Profile for get test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up URL Access Profile with ID: %s", createdProfileID)
		_, errDel := client.URLAccessProfilesAPI.DeleteURLAccessProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete URL Access Profile during cleanup")
	}()

	getRes, httpResGet, errGet := client.URLAccessProfilesAPI.GetURLAccessProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get URL Access Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
}

// Test_security_services_URLAccessProfilesAPIService_Update tests updating an existing URL Access Profile.
func Test_security_services_URLAccessProfilesAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-url-access-update-" + common.GenerateRandomString(6)
	profile := security_services.UrlAccessProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		Block:  []string{"hacking"},
	}

	createRes, _, err := client.URLAccessProfilesAPI.CreateURLAccessProfiles(context.Background()).UrlAccessProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create URL Access Profile for update test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up URL Access Profile with ID: %s", createdProfileID)
		_, errDel := client.URLAccessProfilesAPI.DeleteURLAccessProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete URL Access Profile during cleanup")
	}()

	updatedProfile := security_services.UrlAccessProfiles{
		Name:        profileName,
		Description: common.StringPtr("Updated description"),
		Block:       []string{"hacking", "malware"},
		Continue:    []string{"questionable"},
	}

	updateRes, httpResUpdate, errUpdate := client.URLAccessProfilesAPI.UpdateURLAccessProfilesByID(context.Background(), createdProfileID).UrlAccessProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update URL Access Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, "Updated description", *updateRes.Description, "Description should be updated")
	assert.ElementsMatch(t, []string{"hacking", "malware"}, updateRes.Block, "Block list should be updated")
	assert.ElementsMatch(t, []string{"questionable"}, updateRes.Continue, "Continue list should be added")
}

// Test_security_services_URLAccessProfilesAPIService_List tests listing URL Access Profiles.
func Test_security_services_URLAccessProfilesAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-url-access-list-" + common.GenerateRandomString(6)
	profile := security_services.UrlAccessProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
	}

	createRes, _, err := client.URLAccessProfilesAPI.CreateURLAccessProfiles(context.Background()).UrlAccessProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create URL Access Profile for list test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up URL Access Profile with ID: %s", createdProfileID)
		_, errDel := client.URLAccessProfilesAPI.DeleteURLAccessProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete URL Access Profile during cleanup")
	}()

	listRes, httpResList, errList := client.URLAccessProfilesAPI.ListURLAccessProfiles(context.Background()).Folder("Shared").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list URL Access Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	foundObject := false
	for _, p := range listRes.Data {
		if p.Name != "" && p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created URL Access Profile should be found in the list")
}

// Test_security_services_URLAccessProfilesAPIService_DeleteByID tests deleting a URL Access Profile by its ID.
func Test_security_services_URLAccessProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-url-access-delete-" + common.GenerateRandomString(6)
	profile := security_services.UrlAccessProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
	}

	createRes, _, err := client.URLAccessProfilesAPI.CreateURLAccessProfiles(context.Background()).UrlAccessProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create URL Access Profile for delete test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	httpResDel, errDel := client.URLAccessProfilesAPI.DeleteURLAccessProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete URL Access Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_security_services_URLAccessProfilesAPIService_FetchURLAccessProfiles tests the FetchURLAccessProfiles convenience method
func Test_security_services_URLAccessProfilesAPIService_FetchURLAccessProfiles(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create test object using same payload as Create test
	testName := "test-url-access-fetch-" + common.GenerateRandomString(6)
	testObj := security_services.UrlAccessProfiles{
		Folder:      common.StringPtr("Shared"),
		Name:        testName,
		Description: common.StringPtr("Test URL Access Profile for fetch API"),
		Block:       []string{"adult", "gambling"},
		Alert:       []string{"high-risk", "phishing"},
	}

	createReq := client.URLAccessProfilesAPI.CreateURLAccessProfiles(context.Background()).UrlAccessProfiles(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created profile should have an ID")
	createdID := *createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.URLAccessProfilesAPI.DeleteURLAccessProfilesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.URLAccessProfilesAPI.FetchURLAccessProfiles(
		context.Background(),
		testName,
		common.StringPtr("Shared"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch url_access_profiles by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	require.NotNil(t, fetchedObj.Id, "Fetched object should have an ID")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchURLAccessProfiles found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.URLAccessProfilesAPI.FetchURLAccessProfiles(
		context.Background(),
		"non-existent-url_access_profiles-xyz-12345",
		common.StringPtr("Shared"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchURLAccessProfiles correctly returned nil for non-existent object")
}
