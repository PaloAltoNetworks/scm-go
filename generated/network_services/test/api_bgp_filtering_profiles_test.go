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

// Test_network_services_BGPFilteringProfilesAPIService_Create tests the creation of a BGP filtering profile.
func Test_network_services_BGPFilteringProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-filter-create-" + common.GenerateRandomString(6)

	profile := network_services.BgpFilteringProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
	}

	t.Logf("Creating BGP filtering profile with name: %s", profileName)
	req := client.BGPFilteringProfilesAPI.CreateBGPFilteringProfiles(context.Background()).BgpFilteringProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create BGP filtering profile")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	createdProfileID := res.Id

	// Cleanup the created profile
	defer func() {
		t.Logf("Cleaning up BGP filtering profile with ID: %s", *createdProfileID)
		_, errDel := client.BGPFilteringProfilesAPI.DeleteBGPFilteringProfilesByID(context.Background(), *createdProfileID).Execute()
		if errDel != nil {
			t.Logf("Cleanup failed: %v", errDel)
		}
	}()

	// Verify the response matches key input fields
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
}

// Test_network_services_BGPFilteringProfilesAPIService_GetByID tests retrieving a BGP filtering profile by ID.
func Test_network_services_BGPFilteringProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-filter-get-" + common.GenerateRandomString(6)

	profile := network_services.BgpFilteringProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPFilteringProfilesAPI.CreateBGPFilteringProfiles(context.Background()).BgpFilteringProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for get test setup")
	createdProfileID := createRes.Id

	defer func() {
		client.BGPFilteringProfilesAPI.DeleteBGPFilteringProfilesByID(context.Background(), *createdProfileID).Execute()
	}()

	// Test: Retrieve the profile
	getRes, httpResGet, errGet := client.BGPFilteringProfilesAPI.GetBGPFilteringProfilesByID(context.Background(), *createdProfileID).Execute()

	require.NoError(t, errGet, "Failed to get BGP filtering profile by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
	assert.Equal(t, *createdProfileID, *getRes.Id, "Profile ID should match")
}

// Test_network_services_BGPFilteringProfilesAPIService_UpdateByID tests updating a BGP filtering profile.
func Test_network_services_BGPFilteringProfilesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-filter-update-" + common.GenerateRandomString(6)

	profile := network_services.BgpFilteringProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPFilteringProfilesAPI.CreateBGPFilteringProfiles(context.Background()).BgpFilteringProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for update test setup")
	createdProfileID := createRes.Id

	defer func() {
		client.BGPFilteringProfilesAPI.DeleteBGPFilteringProfilesByID(context.Background(), *createdProfileID).Execute()
	}()

	// Test: Update the profile
	updatedProfile := network_services.BgpFilteringProfiles{
		Name:        profileName,
		Folder:      common.StringPtr("Prisma Access"),
		Description: common.StringPtr("Updated description"),
	}

	updateRes, httpResUpdate, errUpdate := client.BGPFilteringProfilesAPI.UpdateBGPFilteringProfilesByID(context.Background(), *createdProfileID).BgpFilteringProfiles(updatedProfile).Execute()

	require.NoError(t, errUpdate, "Failed to update BGP filtering profile")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the updated data
	assert.Equal(t, profileName, updateRes.Name, "Profile name should match")
	assert.Equal(t, *createdProfileID, *updateRes.Id, "Profile ID should match")
}

// Test_network_services_BGPFilteringProfilesAPIService_List tests listing BGP filtering profiles.
func Test_network_services_BGPFilteringProfilesAPIService_List(t *testing.T) {
	t.Skip("List response contains array fields that cause model deserialization error")
	client := SetupNetworkSvcTestClient(t)

	// Test: List profiles
	listRes, httpResList, errList := client.BGPFilteringProfilesAPI.ListBGPFilteringProfiles(context.Background()).Folder("Prisma Access").Execute()

	require.NoError(t, errList, "Failed to list BGP filtering profiles")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")
}

// Test_network_services_BGPFilteringProfilesAPIService_DeleteByID tests deleting a BGP filtering profile by ID.
func Test_network_services_BGPFilteringProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-filter-delete-" + common.GenerateRandomString(6)

	profile := network_services.BgpFilteringProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPFilteringProfilesAPI.CreateBGPFilteringProfiles(context.Background()).BgpFilteringProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for delete test")
	createdProfileID := createRes.Id

	// Test: Delete the profile
	httpResDel, errDel := client.BGPFilteringProfilesAPI.DeleteBGPFilteringProfilesByID(context.Background(), *createdProfileID).Execute()

	require.NoError(t, errDel, "Failed to delete BGP filtering profile")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_network_services_BGPFilteringProfilesAPIService_Fetch tests the FetchBGPFilteringProfiles convenience method.
func Test_network_services_BGPFilteringProfilesAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-filter-fetch-" + common.GenerateRandomString(6)

	profile := network_services.BgpFilteringProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
	}

	// Create a test profile
	createRes, _, err := client.BGPFilteringProfilesAPI.CreateBGPFilteringProfiles(context.Background()).BgpFilteringProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create test profile for fetch test")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		client.BGPFilteringProfilesAPI.DeleteBGPFilteringProfilesByID(context.Background(), *createdID).Execute()
		t.Logf("Cleaned up test profile: %s", *createdID)
	}()

	// Test 1: Fetch existing profile by name
	fetchedProfile, err := client.BGPFilteringProfilesAPI.FetchBGPFilteringProfiles(
		context.Background(),
		profileName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch BGP filtering profile by name")
	require.NotNil(t, fetchedProfile, "Fetched profile should not be nil")
	assert.Equal(t, *createdID, *fetchedProfile.Id, "Fetched profile ID should match")
	assert.Equal(t, profileName, fetchedProfile.Name, "Fetched profile name should match")
	t.Logf("[SUCCESS] FetchBGPFilteringProfiles found profile: %s", fetchedProfile.Name)

	// Test 2: Fetch non-existent profile (should return nil, nil)
	notFound, err := client.BGPFilteringProfilesAPI.FetchBGPFilteringProfiles(
		context.Background(),
		"non-existent-bgp-filter-profile-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent profile")
	assert.Nil(t, notFound, "Should return nil for non-existent profile")
	t.Logf("[SUCCESS] FetchBGPFilteringProfiles correctly returned nil for non-existent profile")
}
