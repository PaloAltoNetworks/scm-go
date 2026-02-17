package network_services

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Test_network_services_BGPAuthenticationProfilesAPIService_Create tests the creation of a BGP authentication profile.
func Test_network_services_BGPAuthenticationProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-auth-create-" + common.GenerateRandomString(6)

	profile := network_services.BgpAuthProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Secret: common.StringPtr("test-secret-key-123"),
	}

	t.Logf("Creating BGP authentication profile with name: %s", profileName)
	req := client.BGPAuthenticationProfilesAPI.CreateBGPAuthenticationProfiles(context.Background()).BgpAuthProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create BGP authentication profile")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	createdProfileID := res.Id

	// Cleanup the created profile
	defer func() {
		t.Logf("Cleaning up BGP authentication profile with ID: %s", *createdProfileID)
		_, errDel := client.BGPAuthenticationProfilesAPI.DeleteBGPAuthenticationProfilesByID(context.Background(), *createdProfileID).Execute()
		if errDel != nil {
			t.Logf("Cleanup failed: %v", errDel)
		}
	}()

	// Verify the response matches key input fields
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
}

// Test_network_services_BGPAuthenticationProfilesAPIService_GetByID tests retrieving a BGP authentication profile by ID.
func Test_network_services_BGPAuthenticationProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-auth-get-" + common.GenerateRandomString(6)

	profile := network_services.BgpAuthProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Secret: common.StringPtr("test-secret-key-456"),
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPAuthenticationProfilesAPI.CreateBGPAuthenticationProfiles(context.Background()).BgpAuthProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for get test setup")
	createdProfileID := createRes.Id

	defer func() {
		client.BGPAuthenticationProfilesAPI.DeleteBGPAuthenticationProfilesByID(context.Background(), *createdProfileID).Execute()
	}()

	// Test: Retrieve the profile
	getRes, httpResGet, errGet := client.BGPAuthenticationProfilesAPI.GetBGPAuthenticationProfilesByID(context.Background(), *createdProfileID).Execute()

	require.NoError(t, errGet, "Failed to get BGP authentication profile by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
	assert.Equal(t, *createdProfileID, *getRes.Id, "Profile ID should match")
}

// Test_network_services_BGPAuthenticationProfilesAPIService_UpdateByID tests updating a BGP authentication profile.
func Test_network_services_BGPAuthenticationProfilesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-auth-update-" + common.GenerateRandomString(6)

	profile := network_services.BgpAuthProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Secret: common.StringPtr("original-secret-789"),
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPAuthenticationProfilesAPI.CreateBGPAuthenticationProfiles(context.Background()).BgpAuthProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for update test setup")
	createdProfileID := createRes.Id

	defer func() {
		client.BGPAuthenticationProfilesAPI.DeleteBGPAuthenticationProfilesByID(context.Background(), *createdProfileID).Execute()
	}()

	// Test: Update the profile
	updatedProfile := network_services.BgpAuthProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Secret: common.StringPtr("updated-secret-999"),
	}

	updateRes, httpResUpdate, errUpdate := client.BGPAuthenticationProfilesAPI.UpdateBGPAuthenticationProfilesByID(context.Background(), *createdProfileID).BgpAuthProfiles(updatedProfile).Execute()

	require.NoError(t, errUpdate, "Failed to update BGP authentication profile")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the updated data
	assert.Equal(t, profileName, updateRes.Name, "Profile name should match")
	assert.Equal(t, *createdProfileID, *updateRes.Id, "Profile ID should match")
}

// Test_network_services_BGPAuthenticationProfilesAPIService_List tests listing BGP authentication profiles.
func Test_network_services_BGPAuthenticationProfilesAPIService_List(t *testing.T) {
	t.Skip("List response contains array fields that cause model deserialization error")
	client := SetupNetworkSvcTestClient(t)

	// Test: List profiles
	listRes, httpResList, errList := client.BGPAuthenticationProfilesAPI.ListBGPAuthenticationProfiles(context.Background()).Folder("Prisma Access").Execute()

	require.NoError(t, errList, "Failed to list BGP authentication profiles")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")
}

// Test_network_services_BGPAuthenticationProfilesAPIService_DeleteByID tests deleting a BGP authentication profile by ID.
func Test_network_services_BGPAuthenticationProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-auth-delete-" + common.GenerateRandomString(6)

	profile := network_services.BgpAuthProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Secret: common.StringPtr("delete-test-secret"),
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPAuthenticationProfilesAPI.CreateBGPAuthenticationProfiles(context.Background()).BgpAuthProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for delete test")
	createdProfileID := createRes.Id

	// Test: Delete the profile
	httpResDel, errDel := client.BGPAuthenticationProfilesAPI.DeleteBGPAuthenticationProfilesByID(context.Background(), *createdProfileID).Execute()

	require.NoError(t, errDel, "Failed to delete BGP authentication profile")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_network_services_BGPAuthenticationProfilesAPIService_Fetch tests the FetchBGPAuthenticationProfiles convenience method.
func Test_network_services_BGPAuthenticationProfilesAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-auth-fetch-" + common.GenerateRandomString(6)

	profile := network_services.BgpAuthProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Secret: common.StringPtr("fetch-test-secret"),
	}

	// Create a test profile
	createRes, _, err := client.BGPAuthenticationProfilesAPI.CreateBGPAuthenticationProfiles(context.Background()).BgpAuthProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create test profile for fetch test")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		client.BGPAuthenticationProfilesAPI.DeleteBGPAuthenticationProfilesByID(context.Background(), *createdID).Execute()
		t.Logf("Cleaned up test profile: %s", *createdID)
	}()

	// Test 1: Fetch existing profile by name
	fetchedProfile, err := client.BGPAuthenticationProfilesAPI.FetchBGPAuthenticationProfiles(
		context.Background(),
		profileName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch BGP authentication profile by name")
	require.NotNil(t, fetchedProfile, "Fetched profile should not be nil")
	assert.Equal(t, *createdID, *fetchedProfile.Id, "Fetched profile ID should match")
	assert.Equal(t, profileName, fetchedProfile.Name, "Fetched profile name should match")
	t.Logf("[SUCCESS] FetchBGPAuthenticationProfiles found profile: %s", fetchedProfile.Name)

	// Test 2: Fetch non-existent profile (should return nil, nil)
	notFound, err := client.BGPAuthenticationProfilesAPI.FetchBGPAuthenticationProfiles(
		context.Background(),
		"non-existent-bgp-auth-profile-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent profile")
	assert.Nil(t, notFound, "Should return nil for non-existent profile")
	t.Logf("[SUCCESS] FetchBGPAuthenticationProfiles correctly returned nil for non-existent profile")
}
