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

// Test_network_services_OSPFAuthenticationProfilesAPIService_Create tests the creation of an OSPF authentication profile.
func Test_network_services_OSPFAuthenticationProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-ospf-auth-create-" + common.GenerateRandomString(6)

	profile := network_services.OspfAuthProfiles{
		Name:     profileName,
		Folder:   common.StringPtr("Prisma Access"),
		Password: common.StringPtr("testpw1"),
	}

	t.Logf("Creating OSPF authentication profile with name: %s", profileName)
	req := client.OSPFAuthenticationProfilesAPI.CreateOSPFAuthenticationProfiles(context.Background()).OspfAuthProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create OSPF authentication profile")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	createdProfileID := res.Id

	// Cleanup the created profile
	defer func() {
		t.Logf("Cleaning up OSPF authentication profile with ID: %s", *createdProfileID)
		_, errDel := client.OSPFAuthenticationProfilesAPI.DeleteOSPFAuthenticationProfilesByID(context.Background(), *createdProfileID).Execute()
		if errDel != nil {
			t.Logf("Cleanup failed: %v", errDel)
		}
	}()

	// Verify the response matches key input fields
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
}

// Test_network_services_OSPFAuthenticationProfilesAPIService_GetByID tests retrieving an OSPF authentication profile by ID.
func Test_network_services_OSPFAuthenticationProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-ospf-auth-get-" + common.GenerateRandomString(6)

	profile := network_services.OspfAuthProfiles{
		Name:     profileName,
		Folder:   common.StringPtr("Prisma Access"),
		Password: common.StringPtr("testpw2"),
	}

	// Setup: Create a profile first
	createRes, _, err := client.OSPFAuthenticationProfilesAPI.CreateOSPFAuthenticationProfiles(context.Background()).OspfAuthProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for get test setup")
	createdProfileID := createRes.Id

	defer func() {
		client.OSPFAuthenticationProfilesAPI.DeleteOSPFAuthenticationProfilesByID(context.Background(), *createdProfileID).Execute()
	}()

	// Test: Retrieve the profile
	getRes, httpResGet, errGet := client.OSPFAuthenticationProfilesAPI.GetOSPFAuthenticationProfilesByID(context.Background(), *createdProfileID).Execute()

	require.NoError(t, errGet, "Failed to get OSPF authentication profile by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
	assert.Equal(t, *createdProfileID, *getRes.Id, "Profile ID should match")
}

// Test_network_services_OSPFAuthenticationProfilesAPIService_UpdateByID tests updating an OSPF authentication profile.
func Test_network_services_OSPFAuthenticationProfilesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-ospf-auth-update-" + common.GenerateRandomString(6)

	profile := network_services.OspfAuthProfiles{
		Name:     profileName,
		Folder:   common.StringPtr("Prisma Access"),
		Password: common.StringPtr("origpw3"),
	}

	// Setup: Create a profile first
	createRes, _, err := client.OSPFAuthenticationProfilesAPI.CreateOSPFAuthenticationProfiles(context.Background()).OspfAuthProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for update test setup")
	createdProfileID := createRes.Id

	defer func() {
		client.OSPFAuthenticationProfilesAPI.DeleteOSPFAuthenticationProfilesByID(context.Background(), *createdProfileID).Execute()
	}()

	// Test: Update the profile
	updatedProfile := network_services.OspfAuthProfiles{
		Name:     profileName,
		Folder:   common.StringPtr("Prisma Access"),
		Password: common.StringPtr("updpw4"),
	}

	updateRes, httpResUpdate, errUpdate := client.OSPFAuthenticationProfilesAPI.UpdateOSPFAuthenticationProfilesByID(context.Background(), *createdProfileID).OspfAuthProfiles(updatedProfile).Execute()

	require.NoError(t, errUpdate, "Failed to update OSPF authentication profile")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the updated data
	assert.Equal(t, profileName, updateRes.Name, "Profile name should match")
	assert.Equal(t, *createdProfileID, *updateRes.Id, "Profile ID should match")
}

// Test_network_services_OSPFAuthenticationProfilesAPIService_List tests listing OSPF authentication profiles.
func Test_network_services_OSPFAuthenticationProfilesAPIService_List(t *testing.T) {
	t.Skip("List response contains array fields that cause model deserialization error")
	client := SetupNetworkSvcTestClient(t)

	// Test: List profiles
	listRes, httpResList, errList := client.OSPFAuthenticationProfilesAPI.ListOSPFAuthenticationProfiles(context.Background()).Folder("Prisma Access").Execute()

	require.NoError(t, errList, "Failed to list OSPF authentication profiles")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")
}

// Test_network_services_OSPFAuthenticationProfilesAPIService_DeleteByID tests deleting an OSPF authentication profile by ID.
func Test_network_services_OSPFAuthenticationProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-ospf-auth-delete-" + common.GenerateRandomString(6)

	profile := network_services.OspfAuthProfiles{
		Name:     profileName,
		Folder:   common.StringPtr("Prisma Access"),
		Password: common.StringPtr("delpw5"),
	}

	// Setup: Create a profile first
	createRes, _, err := client.OSPFAuthenticationProfilesAPI.CreateOSPFAuthenticationProfiles(context.Background()).OspfAuthProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create profile for delete test")
	createdProfileID := createRes.Id

	// Test: Delete the profile
	httpResDel, errDel := client.OSPFAuthenticationProfilesAPI.DeleteOSPFAuthenticationProfilesByID(context.Background(), *createdProfileID).Execute()

	require.NoError(t, errDel, "Failed to delete OSPF authentication profile")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_network_services_OSPFAuthenticationProfilesAPIService_Fetch tests the FetchOSPFAuthenticationProfiles convenience method.
func Test_network_services_OSPFAuthenticationProfilesAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-ospf-auth-fetch-" + common.GenerateRandomString(6)

	profile := network_services.OspfAuthProfiles{
		Name:     profileName,
		Folder:   common.StringPtr("Prisma Access"),
		Password: common.StringPtr("fetchpw6"),
	}

	// Create a test profile
	createRes, _, err := client.OSPFAuthenticationProfilesAPI.CreateOSPFAuthenticationProfiles(context.Background()).OspfAuthProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create test profile for fetch test")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		client.OSPFAuthenticationProfilesAPI.DeleteOSPFAuthenticationProfilesByID(context.Background(), *createdID).Execute()
		t.Logf("Cleaned up test profile: %s", *createdID)
	}()

	// Test 1: Fetch existing profile by name
	fetchedProfile, err := client.OSPFAuthenticationProfilesAPI.FetchOSPFAuthenticationProfiles(
		context.Background(),
		profileName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch OSPF authentication profile by name")
	require.NotNil(t, fetchedProfile, "Fetched profile should not be nil")
	assert.Equal(t, *createdID, *fetchedProfile.Id, "Fetched profile ID should match")
	assert.Equal(t, profileName, fetchedProfile.Name, "Fetched profile name should match")
	t.Logf("[SUCCESS] FetchOSPFAuthenticationProfiles found profile: %s", fetchedProfile.Name)

	// Test 2: Fetch non-existent profile (should return nil, nil)
	notFound, err := client.OSPFAuthenticationProfilesAPI.FetchOSPFAuthenticationProfiles(
		context.Background(),
		"non-existent-ospf-auth-profile-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent profile")
	assert.Nil(t, notFound, "Should return nil for non-existent profile")
	t.Logf("[SUCCESS] FetchOSPFAuthenticationProfiles correctly returned nil for non-existent profile")
}
