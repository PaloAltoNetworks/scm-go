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

// Test_network_services_BGPAddressFamilyProfilesAPIService_Create tests the creation of a BGP address family profile.
func Test_network_services_BGPAddressFamilyProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-addr-fam-create-" + common.GenerateRandomString(6)

	unicast := network_services.BgpAddressFamily{
		Enable: common.BoolPtr(true),
	}
	ipv4 := network_services.BgpAddressFamilyProfilesIpv4{
		Unicast: &unicast,
	}
	profile := network_services.BgpAddressFamilyProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   &ipv4,
	}

	t.Logf("Creating BGP address family profile with name: %s", profileName)
	req := client.BGPAddressFamilyProfilesAPI.CreateBGPAddressFamilyProfiles(context.Background()).BgpAddressFamilyProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create BGP address family profile")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	createdProfileID := res.Id

	// Cleanup the created profile
	defer func() {
		t.Logf("Cleaning up BGP address family profile with ID: %s", *createdProfileID)
		_, errDel := client.BGPAddressFamilyProfilesAPI.DeleteBGPAddressFamilyProfilesByID(context.Background(), *createdProfileID).Execute()
		if errDel != nil {
			t.Logf("Cleanup failed: %v", errDel)
		}
	}()

	// Verify the response matches key input fields
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
}

// Test_network_services_BGPAddressFamilyProfilesAPIService_GetByID tests retrieving a BGP address family profile by ID.
func Test_network_services_BGPAddressFamilyProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-addr-fam-get-" + common.GenerateRandomString(6)

	unicast := network_services.BgpAddressFamily{
		Enable: common.BoolPtr(true),
	}
	ipv4 := network_services.BgpAddressFamilyProfilesIpv4{
		Unicast: &unicast,
	}
	profile := network_services.BgpAddressFamilyProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   &ipv4,
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPAddressFamilyProfilesAPI.CreateBGPAddressFamilyProfiles(context.Background()).BgpAddressFamilyProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for get test setup")
	createdProfileID := createRes.Id

	defer func() {
		client.BGPAddressFamilyProfilesAPI.DeleteBGPAddressFamilyProfilesByID(context.Background(), *createdProfileID).Execute()
	}()

	// Test: Retrieve the profile
	getRes, httpResGet, errGet := client.BGPAddressFamilyProfilesAPI.GetBGPAddressFamilyProfilesByID(context.Background(), *createdProfileID).Execute()

	require.NoError(t, errGet, "Failed to get BGP address family profile by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
	assert.Equal(t, *createdProfileID, *getRes.Id, "Profile ID should match")
}

// Test_network_services_BGPAddressFamilyProfilesAPIService_UpdateByID tests updating a BGP address family profile.
func Test_network_services_BGPAddressFamilyProfilesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-addr-fam-update-" + common.GenerateRandomString(6)

	unicast := network_services.BgpAddressFamily{
		Enable: common.BoolPtr(true),
	}
	ipv4 := network_services.BgpAddressFamilyProfilesIpv4{
		Unicast: &unicast,
	}
	profile := network_services.BgpAddressFamilyProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   &ipv4,
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPAddressFamilyProfilesAPI.CreateBGPAddressFamilyProfiles(context.Background()).BgpAddressFamilyProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for update test setup")
	createdProfileID := createRes.Id

	defer func() {
		client.BGPAddressFamilyProfilesAPI.DeleteBGPAddressFamilyProfilesByID(context.Background(), *createdProfileID).Execute()
	}()

	// Test: Update the profile with multicast enabled
	multicast := network_services.BgpAddressFamily{
		Enable: common.BoolPtr(true),
	}
	updatedIpv4 := network_services.BgpAddressFamilyProfilesIpv4{
		Unicast:   &unicast,
		Multicast: &multicast,
	}
	updatedProfile := network_services.BgpAddressFamilyProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   &updatedIpv4,
	}

	updateRes, httpResUpdate, errUpdate := client.BGPAddressFamilyProfilesAPI.UpdateBGPAddressFamilyProfilesByID(context.Background(), *createdProfileID).BgpAddressFamilyProfiles(updatedProfile).Execute()

	require.NoError(t, errUpdate, "Failed to update BGP address family profile")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the updated data
	assert.Equal(t, profileName, updateRes.Name, "Profile name should match")
	assert.Equal(t, *createdProfileID, *updateRes.Id, "Profile ID should match")
}

// Test_network_services_BGPAddressFamilyProfilesAPIService_List tests listing BGP address family profiles.
func Test_network_services_BGPAddressFamilyProfilesAPIService_List(t *testing.T) {
	t.Skip("List response contains array fields that cause model deserialization error")
	client := SetupNetworkSvcTestClient(t)

	// Test: List profiles
	listRes, httpResList, errList := client.BGPAddressFamilyProfilesAPI.ListBGPAddressFamilyProfiles(context.Background()).Folder("Prisma Access").Execute()

	require.NoError(t, errList, "Failed to list BGP address family profiles")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")
}

// Test_network_services_BGPAddressFamilyProfilesAPIService_DeleteByID tests deleting a BGP address family profile by ID.
func Test_network_services_BGPAddressFamilyProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-addr-fam-delete-" + common.GenerateRandomString(6)

	unicast := network_services.BgpAddressFamily{
		Enable: common.BoolPtr(true),
	}
	ipv4 := network_services.BgpAddressFamilyProfilesIpv4{
		Unicast: &unicast,
	}
	profile := network_services.BgpAddressFamilyProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   &ipv4,
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPAddressFamilyProfilesAPI.CreateBGPAddressFamilyProfiles(context.Background()).BgpAddressFamilyProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for delete test")
	createdProfileID := createRes.Id

	// Test: Delete the profile
	httpResDel, errDel := client.BGPAddressFamilyProfilesAPI.DeleteBGPAddressFamilyProfilesByID(context.Background(), *createdProfileID).Execute()

	require.NoError(t, errDel, "Failed to delete BGP address family profile")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_network_services_BGPAddressFamilyProfilesAPIService_Fetch tests the FetchBGPAddressFamilyProfiles convenience method.
func Test_network_services_BGPAddressFamilyProfilesAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-addr-fam-fetch-" + common.GenerateRandomString(6)

	unicast := network_services.BgpAddressFamily{
		Enable: common.BoolPtr(true),
	}
	ipv4 := network_services.BgpAddressFamilyProfilesIpv4{
		Unicast: &unicast,
	}
	profile := network_services.BgpAddressFamilyProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   &ipv4,
	}

	// Create a test profile
	createRes, _, err := client.BGPAddressFamilyProfilesAPI.CreateBGPAddressFamilyProfiles(context.Background()).BgpAddressFamilyProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create test profile for fetch test")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		client.BGPAddressFamilyProfilesAPI.DeleteBGPAddressFamilyProfilesByID(context.Background(), *createdID).Execute()
		t.Logf("Cleaned up test profile: %s", *createdID)
	}()

	// Test 1: Fetch existing profile by name
	fetchedProfile, err := client.BGPAddressFamilyProfilesAPI.FetchBGPAddressFamilyProfiles(
		context.Background(),
		profileName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch BGP address family profile by name")
	require.NotNil(t, fetchedProfile, "Fetched profile should not be nil")
	assert.Equal(t, *createdID, *fetchedProfile.Id, "Fetched profile ID should match")
	assert.Equal(t, profileName, fetchedProfile.Name, "Fetched profile name should match")
	t.Logf("[SUCCESS] FetchBGPAddressFamilyProfiles found profile: %s", fetchedProfile.Name)

	// Test 2: Fetch non-existent profile (should return nil, nil)
	notFound, err := client.BGPAddressFamilyProfilesAPI.FetchBGPAddressFamilyProfiles(
		context.Background(),
		"non-existent-bgp-addr-fam-profile-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent profile")
	assert.Nil(t, notFound, "Should return nil for non-existent profile")
	t.Logf("[SUCCESS] FetchBGPAddressFamilyProfiles correctly returned nil for non-existent profile")
}
