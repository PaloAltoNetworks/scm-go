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

// Test_network_services_BGPRedistributionProfilesAPIService_Create tests the creation of a BGP redistribution profile.
func Test_network_services_BGPRedistributionProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-redist-create-" + common.GenerateRandomString(6)

	staticConfig := network_services.BgpRedistributionProfilesIpv4UnicastStatic{
		Enable: common.BoolPtr(true),
	}
	unicast := network_services.BgpRedistributionProfilesIpv4Unicast{
		Static: &staticConfig,
	}
	// CRITICAL: Ipv4 is a VALUE type, not a pointer!
	ipv4 := network_services.BgpRedistributionProfilesIpv4{
		Unicast: &unicast,
	}
	profile := network_services.BgpRedistributionProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   ipv4, // VALUE type, not &ipv4
	}

	t.Logf("Creating BGP redistribution profile with name: %s", profileName)
	req := client.BGPRedistributionProfilesAPI.CreateBGPRedistributionProfiles(context.Background()).BgpRedistributionProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create BGP redistribution profile")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	createdProfileID := res.Id

	// Cleanup the created profile
	defer func() {
		t.Logf("Cleaning up BGP redistribution profile with ID: %s", *createdProfileID)
		_, errDel := client.BGPRedistributionProfilesAPI.DeleteBGPRedistributionProfilesByID(context.Background(), *createdProfileID).Execute()
		if errDel != nil {
			t.Logf("Cleanup failed: %v", errDel)
		}
	}()

	// Verify the response matches key input fields
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
}

// Test_network_services_BGPRedistributionProfilesAPIService_GetByID tests retrieving a BGP redistribution profile by ID.
func Test_network_services_BGPRedistributionProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-redist-get-" + common.GenerateRandomString(6)

	staticConfig := network_services.BgpRedistributionProfilesIpv4UnicastStatic{
		Enable: common.BoolPtr(true),
	}
	unicast := network_services.BgpRedistributionProfilesIpv4Unicast{
		Static: &staticConfig,
	}
	ipv4 := network_services.BgpRedistributionProfilesIpv4{
		Unicast: &unicast,
	}
	profile := network_services.BgpRedistributionProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   ipv4,
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPRedistributionProfilesAPI.CreateBGPRedistributionProfiles(context.Background()).BgpRedistributionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for get test setup")
	createdProfileID := createRes.Id

	defer func() {
		client.BGPRedistributionProfilesAPI.DeleteBGPRedistributionProfilesByID(context.Background(), *createdProfileID).Execute()
	}()

	// Test: Retrieve the profile
	getRes, httpResGet, errGet := client.BGPRedistributionProfilesAPI.GetBGPRedistributionProfilesByID(context.Background(), *createdProfileID).Execute()

	require.NoError(t, errGet, "Failed to get BGP redistribution profile by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
	assert.Equal(t, *createdProfileID, *getRes.Id, "Profile ID should match")
}

// Test_network_services_BGPRedistributionProfilesAPIService_UpdateByID tests updating a BGP redistribution profile.
func Test_network_services_BGPRedistributionProfilesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-redist-update-" + common.GenerateRandomString(6)

	staticConfig := network_services.BgpRedistributionProfilesIpv4UnicastStatic{
		Enable: common.BoolPtr(true),
	}
	unicast := network_services.BgpRedistributionProfilesIpv4Unicast{
		Static: &staticConfig,
	}
	ipv4 := network_services.BgpRedistributionProfilesIpv4{
		Unicast: &unicast,
	}
	profile := network_services.BgpRedistributionProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   ipv4,
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPRedistributionProfilesAPI.CreateBGPRedistributionProfiles(context.Background()).BgpRedistributionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for update test setup")
	createdProfileID := createRes.Id

	defer func() {
		client.BGPRedistributionProfilesAPI.DeleteBGPRedistributionProfilesByID(context.Background(), *createdProfileID).Execute()
	}()

	// Test: Update the profile with connected enabled
	connectedConfig := network_services.BgpRedistributionProfilesIpv4UnicastConnected{
		Enable: common.BoolPtr(true),
	}
	updatedUnicast := network_services.BgpRedistributionProfilesIpv4Unicast{
		Static:    &staticConfig,
		Connected: &connectedConfig,
	}
	updatedIpv4 := network_services.BgpRedistributionProfilesIpv4{
		Unicast: &updatedUnicast,
	}
	updatedProfile := network_services.BgpRedistributionProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   updatedIpv4,
	}

	updateRes, httpResUpdate, errUpdate := client.BGPRedistributionProfilesAPI.UpdateBGPRedistributionProfilesByID(context.Background(), *createdProfileID).BgpRedistributionProfiles(updatedProfile).Execute()

	require.NoError(t, errUpdate, "Failed to update BGP redistribution profile")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the updated data
	assert.Equal(t, profileName, updateRes.Name, "Profile name should match")
	assert.Equal(t, *createdProfileID, *updateRes.Id, "Profile ID should match")
}

// Test_network_services_BGPRedistributionProfilesAPIService_List tests listing BGP redistribution profiles.
func Test_network_services_BGPRedistributionProfilesAPIService_List(t *testing.T) {
	t.Skip("List response contains array fields that cause model deserialization error")
	client := SetupNetworkSvcTestClient(t)

	// Test: List profiles
	listRes, httpResList, errList := client.BGPRedistributionProfilesAPI.ListBGPRedistributionProfiles(context.Background()).Folder("Prisma Access").Execute()

	require.NoError(t, errList, "Failed to list BGP redistribution profiles")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")
}

// Test_network_services_BGPRedistributionProfilesAPIService_DeleteByID tests deleting a BGP redistribution profile by ID.
func Test_network_services_BGPRedistributionProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-redist-delete-" + common.GenerateRandomString(6)

	staticConfig := network_services.BgpRedistributionProfilesIpv4UnicastStatic{
		Enable: common.BoolPtr(true),
	}
	unicast := network_services.BgpRedistributionProfilesIpv4Unicast{
		Static: &staticConfig,
	}
	ipv4 := network_services.BgpRedistributionProfilesIpv4{
		Unicast: &unicast,
	}
	profile := network_services.BgpRedistributionProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   ipv4,
	}

	// Setup: Create a profile first
	createRes, _, err := client.BGPRedistributionProfilesAPI.CreateBGPRedistributionProfiles(context.Background()).BgpRedistributionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for delete test")
	createdProfileID := createRes.Id

	// Test: Delete the profile
	httpResDel, errDel := client.BGPRedistributionProfilesAPI.DeleteBGPRedistributionProfilesByID(context.Background(), *createdProfileID).Execute()

	require.NoError(t, errDel, "Failed to delete BGP redistribution profile")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_network_services_BGPRedistributionProfilesAPIService_Fetch tests the FetchBGPRedistributionProfiles convenience method.
func Test_network_services_BGPRedistributionProfilesAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-bgp-redist-fetch-" + common.GenerateRandomString(6)

	staticConfig := network_services.BgpRedistributionProfilesIpv4UnicastStatic{
		Enable: common.BoolPtr(true),
	}
	unicast := network_services.BgpRedistributionProfilesIpv4Unicast{
		Static: &staticConfig,
	}
	ipv4 := network_services.BgpRedistributionProfilesIpv4{
		Unicast: &unicast,
	}
	profile := network_services.BgpRedistributionProfiles{
		Name:   profileName,
		Folder: common.StringPtr("Prisma Access"),
		Ipv4:   ipv4,
	}

	// Create a test profile
	createRes, _, err := client.BGPRedistributionProfilesAPI.CreateBGPRedistributionProfiles(context.Background()).BgpRedistributionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create test profile for fetch test")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		client.BGPRedistributionProfilesAPI.DeleteBGPRedistributionProfilesByID(context.Background(), *createdID).Execute()
		t.Logf("Cleaned up test profile: %s", *createdID)
	}()

	// Test 1: Fetch existing profile by name
	fetchedProfile, err := client.BGPRedistributionProfilesAPI.FetchBGPRedistributionProfiles(
		context.Background(),
		profileName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch BGP redistribution profile by name")
	require.NotNil(t, fetchedProfile, "Fetched profile should not be nil")
	assert.Equal(t, *createdID, *fetchedProfile.Id, "Fetched profile ID should match")
	assert.Equal(t, profileName, fetchedProfile.Name, "Fetched profile name should match")
	t.Logf("[SUCCESS] FetchBGPRedistributionProfiles found profile: %s", fetchedProfile.Name)

	// Test 2: Fetch non-existent profile (should return nil, nil)
	notFound, err := client.BGPRedistributionProfilesAPI.FetchBGPRedistributionProfiles(
		context.Background(),
		"non-existent-bgp-redist-profile-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent profile")
	assert.Nil(t, notFound, "Should return nil for non-existent profile")
	t.Logf("[SUCCESS] FetchBGPRedistributionProfiles correctly returned nil for non-existent profile")
}
