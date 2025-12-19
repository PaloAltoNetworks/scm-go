/*
 * Security Services Testing
 *
 * HTTPHeaderProfilesAPIService
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

// Test_security_services_HTTPHeaderProfilesAPIService_Create tests the creation of an HTTP Header Profile
func Test_security_services_HTTPHeaderProfilesAPIService_Create(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdProfileName := "test-http-header-create-" + common.GenerateRandomString(6)

	// define headers
	headerOne := security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{
		Header: "Authorization",
		Name:   "Authorization",
		Value:  "general auth",
	}

	headerTwo := security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{
		Header: "X-Authenticated-User",
		Name:   "X-Authenticated-User",
		Value:  "auth for users",
		Log:    common.BoolPtr(true),
	}

	headerThree := security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{
		Header: "YouTube-Restrict",
		Name:   "YouTube-Restrict",
		Value:  "192.168.1.1",
		Log:    common.BoolPtr(true),
	}

	// define types
	typeOne := security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner{
		Name:    "Dynamic Fields",
		Domains: []string{"custom_domain"},
		Headers: []security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{headerOne, headerTwo},
	}

	typeTwo := security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner{
		Name:    "Youtube Safe Search",
		Domains: []string{"m.youtube.com", "www.youtube.com", "youtubei.googleapis.com"},
		Headers: []security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{headerThree},
	}

	// define http header insertion
	httpHeaderInsertions := []security_services.HttpHeaderProfilesHttpHeaderInsertionInner{
		{
			Name: "RuleOne",
			Type: []security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner{typeOne},
		},
		{
			Name: "RuleTwo",
			Type: []security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner{typeTwo},
		},
	}

	// create http header profile
	profile := security_services.HttpHeaderProfiles{
		Folder:              common.StringPtr("Shared"),
		Name:                createdProfileName,
		Description:         common.StringPtr("Test HTTP Header Profile for create API"),
		HttpHeaderInsertion: httpHeaderInsertions,
	}

	fmt.Printf("Creating HTTP Header Profile with name: %s\n", profile.Name)
	req := client.HTTPHeaderProfilesAPI.CreateHTTPHeaderProfiles(context.Background()).HttpHeaderProfiles(profile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create HTTP Header Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")
	createdProfileID := *res.Id

	defer func() {
		t.Logf("Cleaning up HTTP Header Profile with ID: %s", createdProfileID)
		_, errDel := client.HTTPHeaderProfilesAPI.DeleteHTTPHeaderProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete HTTP Header Profile during cleanup")
	}()

	t.Logf("Successfully created HTTP Header Profile: %s with ID: %s", profile.Name, createdProfileID)

	assert.Equal(t, "Shared", *res.Folder, "Folder should match")
	assert.Equal(t, "Test HTTP Header Profile for create API", *res.Description, "Description should match")
	assert.Len(t, res.HttpHeaderInsertion, 2, "Should have 2 header insertion rules")
}

func Test_security_services_HTTPHeaderProfilesAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-http-header-get-" + common.GenerateRandomString(6)

	profile := security_services.HttpHeaderProfiles{
		Folder:      common.StringPtr("Shared"),
		Name:        profileName,
		Description: common.StringPtr("Test HTTP Header Profile for get API"),
	}

	createRes, _, err := client.HTTPHeaderProfilesAPI.CreateHTTPHeaderProfiles(context.Background()).HttpHeaderProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create HTTP Header Profile for get test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up HTTP Header Profile with ID: %s", createdProfileID)
		_, errDel := client.HTTPHeaderProfilesAPI.DeleteHTTPHeaderProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete HTTP Header Profile during cleanup")
	}()

	getRes, httpResGet, errGet := client.HTTPHeaderProfilesAPI.GetHTTPHeaderProfilesByID(context.Background(), createdProfileID).Execute()

	require.NoError(t, errGet, "Failed to get HTTP Header Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	assert.Equal(t, createdProfileID, getRes.GetId(), "Returned profile ID should match the created ID")
	assert.Equal(t, profileName, getRes.GetName(), "Returned profile name should match the created name")
	assert.Equal(t, 0, len(getRes.GetHttpHeaderInsertion()), "HttpHeaderInsertion should be empty since we didn't define it")
}

func Test_security_services_HTTPHeaderProfilesAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-http-header-update-" + common.GenerateRandomString(6)
	initialDescription := "Initial description"

	profile := security_services.HttpHeaderProfiles{
		Folder:      common.StringPtr("Shared"),
		Name:        profileName,
		Description: common.StringPtr(initialDescription),
	}

	createRes, _, err := client.HTTPHeaderProfilesAPI.CreateHTTPHeaderProfiles(context.Background()).HttpHeaderProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create HTTP Header Profile for update test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up HTTP Header Profile with ID: %s", createdProfileID)
		_, errDel := client.HTTPHeaderProfilesAPI.DeleteHTTPHeaderProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete HTTP Header Profile during cleanup")
	}()

	updatedDescription := "Updated HTTP header profile description"

	updatedProfile := security_services.HttpHeaderProfiles{
		Name:        profileName,
		Description: common.StringPtr(updatedDescription),
		HttpHeaderInsertion: []security_services.HttpHeaderProfilesHttpHeaderInsertionInner{
			{
				Name: "New_Rule",
				Type: []security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInner{
					{
						Name:    "Dropbox Network Control",
						Domains: []string{"*.dropboxapi.com", "*.db.tt"},
						Headers: []security_services.HttpHeaderProfilesHttpHeaderInsertionInnerTypeInnerHeadersInner{

							{
								Header: "X-Dropbox-allowed-Team-Ids",
								Name:   "X-Dropbox-allowed-Team-Ids",
								Value:  "dropbox",
								Log:    common.BoolPtr(true),
							},
						},
					},
				},
			},
		},
	}

	updateRes, httpResUpdate, errUpdate := client.HTTPHeaderProfilesAPI.UpdateHTTPHeaderProfilesByID(context.Background(), createdProfileID).HttpHeaderProfiles(updatedProfile).Execute()

	require.NoError(t, errUpdate, "Failed to update HTTP Header Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
	assert.Len(t, updateRes.HttpHeaderInsertion, 1, "Should have 1 header insertion rule after update")
	assert.Equal(t, "New_Rule", updateRes.HttpHeaderInsertion[0].Name, "Rule name should match updated value")
}

// Test_security_services_HTTPHeaderProfilesAPIService_List tests listing HTTP Header Profiles.
func Test_security_services_HTTPHeaderProfilesAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-http-header-list-" + common.GenerateRandomString(6)
	profile := security_services.HttpHeaderProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
	}

	// create a profile to ensure list returns data
	createRes, _, err := client.HTTPHeaderProfilesAPI.CreateHTTPHeaderProfiles(context.Background()).HttpHeaderProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create HTTP Header Profile for list test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up HTTP Header Profile with ID: %s", createdProfileID)
		_, errDel := client.HTTPHeaderProfilesAPI.DeleteHTTPHeaderProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete HTTP Header Profile during cleanup")
	}()

	// list profiles, filtered by the folder 'Shared'
	listRes, httpResList, errList := client.HTTPHeaderProfilesAPI.ListHTTPHeaderProfiles(context.Background()).Folder("Shared").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list HTTP Header Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.True(t, len(listRes.Data) > 0, "List should return at least one profile")

	foundObject := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created HTTP Header Profile should be found in the list")
}

// Test_security_services_HTTPHeaderProfilesAPIService_DeleteByID tests deleting an HTTP Header Profile by its ID.
func Test_security_services_HTTPHeaderProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-http-header-delete-" + common.GenerateRandomString(6)

	profile := security_services.HttpHeaderProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
	}
	createRes, _, err := client.HTTPHeaderProfilesAPI.CreateHTTPHeaderProfiles(context.Background()).HttpHeaderProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create HTTP Header Profile for delete test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	httpResDel, errDel := client.HTTPHeaderProfilesAPI.DeleteHTTPHeaderProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete HTTP Header Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status on successful delete")
}
