/*
 * Security Services Testing
 *
 * FileBlockingProfilesAPIService
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

// Test_security_services_FileBlockingProfilesAPIService_Create tests the creation of a File Blocking Profile.
func Test_security_services_FileBlockingProfilesAPIService_Create(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdProfileName := "test-file-block-create-" + common.GenerateRandomString(6)

	// define rules (one of each action: alert, block, continue)
	testRules := []security_services.FileBlockingProfilesRulesInner{
		{
			// alert action
			Action: "alert",
			Application: []string{
				"any",
			},
			Direction: "both",
			FileType: []string{
				"any",
			},
			Name: "alert_rule",
		},
		{
			// block action
			Action: "block",
			Application: []string{
				"web-browsing",
			},
			Direction: "upload",
			FileType: []string{
				"PE",
			},
			Name: "block_rule",
		},
		{
			// continue action
			Action: "continue",
			Application: []string{
				"adobe-update", "aim-file-transfer", "anyword-base",
			},
			Direction: "download",
			FileType: []string{
				"7z", "bat", "chm", "class", "cpl", "dll", "hlp", "hta", "jar", "ocx", "pif", "scr", "torrent", "vbe", "wsf", // best practice
			},
			Name: "continue_rule",
		},
	}

	// define a file blocking profile
	profile := security_services.FileBlockingProfiles{
		Folder:      common.StringPtr("All"),
		Name:        createdProfileName,
		Description: common.StringPtr("Test File Blocking Profile for create API"),
		Rules:       testRules,
	}

	fmt.Printf("Creating File Blocking Profile with name: %s\n", profile.Name)
	req := client.FileBlockingProfilesAPI.CreateFileBlockingProfiles(context.Background()).FileBlockingProfiles(profile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create File Blocking Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")
	assert.Len(t, res.Rules, 3, "Created profile should have 3 rules")
	assert.Equal(t, "alert_rule", res.Rules[0].Name, "Rule name should match")
	assert.Equal(t, "block_rule", res.Rules[1].Name, "Rule name should match")
	assert.Equal(t, "continue_rule", res.Rules[2].Name, "Rule name should match")

	createdProfileID := res.Id

	defer func() {
		t.Logf("Cleaning up File Blocking Profile with ID: %s", *createdProfileID)
		_, errDel := client.FileBlockingProfilesAPI.DeleteFileBlockingProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete File Blocking Profile during cleanup")
	}()

	t.Logf("Successfully created File Blocking Profile: %s with ID: %s", profile.Name, *createdProfileID)
}

// Test_security_services_FileBlockingProfilesAPIService_GetByID tests retrieving a File Blocking Profile by its ID.
func Test_security_services_FileBlockingProfilesAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-file-block-get-" + common.GenerateRandomString(6)
	profile := security_services.FileBlockingProfiles{
		Folder: common.StringPtr("All"),
		Name:   profileName,
	}

	createRes, _, err := client.FileBlockingProfilesAPI.CreateFileBlockingProfiles(context.Background()).FileBlockingProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create File Blocking Profile for get test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up File Blocking Profile with ID: %s", *createdProfileID)
		_, errDel := client.FileBlockingProfilesAPI.DeleteFileBlockingProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete File Blocking Profile during cleanup")
	}()

	getRes, httpResGet, errGet := client.FileBlockingProfilesAPI.GetFileBlockingProfilesByID(context.Background(), *createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get File Blocking Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
}

// Test_security_services_FileBlockingProfilesAPIService_Update tests updating an existing File Blocking Profile.
func Test_security_services_FileBlockingProfilesAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-file-block-update-" + common.GenerateRandomString(6)
	profile := security_services.FileBlockingProfiles{
		Folder: common.StringPtr("All"),
		Name:   profileName,
	}

	createRes, _, err := client.FileBlockingProfilesAPI.CreateFileBlockingProfiles(context.Background()).FileBlockingProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create File Blocking Profile for update test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up File Blocking Profile with ID: %s", *createdProfileID)
		_, errDel := client.FileBlockingProfilesAPI.DeleteFileBlockingProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete File Blocking Profile during cleanup")
	}()

	updatedDescription := "Updated file blocking profile description"
	// new rule to test the update operation
	newRule := security_services.FileBlockingProfilesRulesInner{
		Action: "alert",
		Application: []string{
			"blackboard-collaborate", "bugzilla", "chatbot-download",
		},
		Direction: "both",
		FileType: []string{
			"pdf", "zip", "shk",
		},
		Name: "alert_new_rule",
	}

	updatedProfile := security_services.FileBlockingProfiles{
		Name:        profileName,
		Description: common.StringPtr(updatedDescription),
		Rules:       []security_services.FileBlockingProfilesRulesInner{newRule},
	}

	updateRes, httpResUpdate, errUpdate := client.FileBlockingProfilesAPI.UpdateFileBlockingProfilesByID(context.Background(), *createdProfileID).FileBlockingProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update File Blocking Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
	assert.Len(t, updateRes.Rules, 1, "Updated profile should have 1 rule")
	assert.Equal(t, "alert_new_rule", updateRes.Rules[0].Name, "Rule name should match the update")
}

// Test_security_services_FileBlockingProfilesAPIService_List tests listing File Blocking Profiles.
func Test_security_services_FileBlockingProfilesAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-file-block-list-" + common.GenerateRandomString(6)
	profile := security_services.FileBlockingProfiles{
		Folder: common.StringPtr("All"),
		Name:   profileName,
	}

	createRes, _, err := client.FileBlockingProfilesAPI.CreateFileBlockingProfiles(context.Background()).FileBlockingProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create File Blocking Profile for list test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up File Blocking Profile with ID: %s", *createdProfileID)
		_, errDel := client.FileBlockingProfilesAPI.DeleteFileBlockingProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete File Blocking Profile during cleanup")
	}()

	listRes, httpResList, errList := client.FileBlockingProfilesAPI.ListFileBlockingProfiles(context.Background()).Folder("Shared").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list File Blocking Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	foundObject := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created File Blocking Profile should be found in the list")
}

// Test_security_services_FileBlockingProfilesAPIService_DeleteByID tests deleting a File Blocking Profile by its ID.
func Test_security_services_FileBlockingProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-file-block-delete-" + common.GenerateRandomString(6)
	profile := security_services.FileBlockingProfiles{
		Folder: common.StringPtr("All"),
		Name:   profileName,
	}

	createRes, _, err := client.FileBlockingProfilesAPI.CreateFileBlockingProfiles(context.Background()).FileBlockingProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create File Blocking Profile for delete test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	httpResDel, errDel := client.FileBlockingProfilesAPI.DeleteFileBlockingProfilesByID(context.Background(), *createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete File Blocking Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}
