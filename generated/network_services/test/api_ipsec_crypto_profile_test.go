/*
 * Network Services Testing
 *
 * IPsecCryptoProfilesAPIService
 */

package network_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Test_networkservices_IPsecCryptoProfilesAPIService_Create tests the creation of an IPsec Crypto Profile.
func Test_networkservices_IPsecCryptoProfilesAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupNetworkSvcTestClient(t)

	// Create a valid IPsec Crypto Profile object with a unique name.
	profileName := "test-ipsec-create-" + common.GenerateRandomString(6)
	profile := network_services.IpsecCryptoProfiles{
		Folder:  common.StringPtr("Prisma Access"),
		Name:    profileName,
		DhGroup: common.StringPtr("group14"),
		Esp: &network_services.IpsecCryptoProfilesEsp{
			Authentication: []string{"sha256"},
			Encryption:     []string{"aes-256-gcm"},
		},
		Lifetime: network_services.IpsecCryptoProfilesLifetime{
			Hours: common.Int32Ptr(8),
		},
	}

	fmt.Printf("Attempting to create IPsec Crypto Profile with name: %s\n", profile.Name)

	// Make the create request to the API.
	req := client.IPsecCryptoProfilesAPI.CreateIPsecCryptoProfiles(context.Background()).IpsecCryptoProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create IPsec Crypto Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created profile should have an ID")

	createdProfileID := *res.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer the cleanup of the created profile.
	defer func() {
		t.Logf("Cleaning up IPsec Crypto Profile with ID: %s", createdProfileID)
		reqDel := client.IPsecCryptoProfilesAPI.DeleteIPsecCryptoProfilesByID(context.Background(), createdProfileID)
		httpResDel, errDel := reqDel.Execute()
		if errDel != nil {
			handleAPIError(err)
		}
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
		assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK for cleanup delete")
	}()

	// Assert response object properties.
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
	require.NotNil(t, res.Folder, "Folder should not be nil in response")
	assert.True(t, *res.Folder == "Shared" || *res.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, "group14", *res.DhGroup, "DH Group should match")
	require.NotNil(t, res.Esp, "ESP should not be nil")
	assert.Equal(t, []string{"aes-256-gcm"}, res.Esp.Encryption, "ESP encryption should match")
	t.Logf("Successfully created and validated IPsec Crypto Profile: %s with ID: %s", profile.Name, createdProfileID)
}

// Test_networkservices_IPsecCryptoProfilesAPIService_GetByID tests retrieving a profile by its ID.
func Test_networkservices_IPsecCryptoProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a profile to retrieve.
	profileName := "test-ipsec-get-" + common.GenerateRandomString(6)
	profile := network_services.IpsecCryptoProfiles{
		Folder:  common.StringPtr("Prisma Access"),
		Name:    profileName,
		DhGroup: common.StringPtr("group20"),
		Esp: &network_services.IpsecCryptoProfilesEsp{
			Authentication: []string{"sha512"},
			Encryption:     []string{"aes-128-cbc"},
		},
		Lifetime: network_services.IpsecCryptoProfilesLifetime{Hours: common.Int32Ptr(1)},
	}

	reqCreate := client.IPsecCryptoProfilesAPI.CreateIPsecCryptoProfiles(context.Background()).IpsecCryptoProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer cleanup.
	defer func() {
		t.Logf("Cleaning up IPsec Crypto Profile with ID: %s", createdProfileID)
		_, errDel := client.IPsecCryptoProfilesAPI.DeleteIPsecCryptoProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
	}()

	// Test Get by ID operation.
	fmt.Printf("Attempting to get IPsec Crypto Profile with ID: %s\n", createdProfileID)
	reqGetById := client.IPsecCryptoProfilesAPI.GetIPsecCrytoProfilesByID(context.Background(), createdProfileID)
	getRes, httpResGet, errGet := reqGetById.Execute()

	if errGet != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful.
	require.NoError(t, errGet, "Failed to get profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
	assert.Equal(t, "group20", *getRes.DhGroup, "DH Group should match")
	require.NotNil(t, getRes.Esp, "ESP should not be nil")
	assert.Equal(t, []string{"sha512"}, getRes.Esp.Authentication, "ESP authentication should match")
	assert.Equal(t, *createRes.Id, *getRes.Id, "Profile ID should match")
	t.Logf("Successfully retrieved IPsec Crypto Profile: %s", getRes.Name)
}

// Test_networkservices_IPsecCryptoProfilesAPIService_Update tests updating an existing profile.
func Test_networkservices_IPsecCryptoProfilesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a profile to update.
	profileName := "test-ipsec-update-" + common.GenerateRandomString(6)
	profile := network_services.IpsecCryptoProfiles{
		Folder:  common.StringPtr("Prisma Access"),
		Name:    profileName,
		DhGroup: common.StringPtr("group14"),
		Esp: &network_services.IpsecCryptoProfilesEsp{
			Authentication: []string{"sha256"},
			Encryption:     []string{"aes-256-cbc"},
		},
		Lifetime: network_services.IpsecCryptoProfilesLifetime{Hours: common.Int32Ptr(8)},
	}

	reqCreate := client.IPsecCryptoProfilesAPI.CreateIPsecCryptoProfiles(context.Background()).IpsecCryptoProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for update test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer cleanup.
	defer func() {
		t.Logf("Cleaning up IPsec Crypto Profile with ID: %s", createdProfileID)
		_, errDel := client.IPsecCryptoProfilesAPI.DeleteIPsecCryptoProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
	}()

	// Define the update payload.
	updatedProfile := network_services.IpsecCryptoProfiles{
		Name:    profileName,
		DhGroup: common.StringPtr("group5"),
		Esp: &network_services.IpsecCryptoProfilesEsp{
			Authentication: []string{"sha384"},
			Encryption:     []string{"aes-192-cbc"},
		},
		Lifetime: network_services.IpsecCryptoProfilesLifetime{
			Minutes: common.Int32Ptr(60),
		},
	}

	fmt.Printf("Attempting to update IPsec Crypto Profile with ID: %s\n", createdProfileID)
	reqUpdate := client.IPsecCryptoProfilesAPI.UpdateIPsecCryptoProfilesByID(context.Background(), createdProfileID).IpsecCryptoProfiles(updatedProfile)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()

	if errUpdate != nil {
		handleAPIError(err)
	}

	// Verify the update was successful.
	require.NoError(t, errUpdate, "Failed to update profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, "group5", *updateRes.DhGroup, "DH Group should be updated")
	require.NotNil(t, updateRes.Esp, "ESP should not be nil")
	assert.Equal(t, []string{"sha384"}, updateRes.Esp.Authentication, "ESP authentication should be updated")
	require.NotNil(t, updateRes.Lifetime.Minutes, "Lifetime minutes should be updated")
	assert.Equal(t, int32(60), *updateRes.Lifetime.Minutes, "Lifetime should be updated")

	t.Logf("Successfully updated IPsec Crypto Profile: %s", profileName)
}

// Test_networkservices_IPsecCryptoProfilesAPIService_List tests listing IPsec Crypto Profiles.
func Test_networkservices_IPsecCryptoProfilesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a profile to ensure it appears in the list.
	profileName := "test-ipsec-list-" + common.GenerateRandomString(6)
	profile := network_services.IpsecCryptoProfiles{
		Folder:  common.StringPtr("Prisma Access"),
		Name:    profileName,
		DhGroup: common.StringPtr("group2"),
		Esp: &network_services.IpsecCryptoProfilesEsp{
			Authentication: []string{"sha256"},
			Encryption:     []string{"aes-256-cbc"},
		},
		Lifetime: network_services.IpsecCryptoProfilesLifetime{Hours: common.Int32Ptr(8)},
	}

	reqCreate := client.IPsecCryptoProfilesAPI.CreateIPsecCryptoProfiles(context.Background()).IpsecCryptoProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for list test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer cleanup.
	defer func() {
		t.Logf("Cleaning up IPsec Crypto Profile with ID: %s", createdProfileID)
		_, errDel := client.IPsecCryptoProfilesAPI.DeleteIPsecCryptoProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
	}()

	// Test List operation.
	fmt.Println("Attempting to list IPsec Crypto Profiles")
	reqList := client.IPsecCryptoProfilesAPI.ListIPsecCryptoProfiles(context.Background()).Folder("Prisma Access").Limit(10000)
	listRes, httpResList, errList := reqList.Execute()

	if errList != nil {
		handleAPIError(err)
	}

	// Verify the list operation was successful.
	require.NoError(t, errList, "Failed to list profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	assert.True(t, len(listRes.Data) > 0, "Should have at least one profile in the list")

	// Verify our created profile is in the list.
	foundProfile := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			foundProfile = true
			assert.Equal(t, "group2", *p.DhGroup, "DH Group should match")
			break
		}
	}
	assert.True(t, foundProfile, "Created profile should be found in the list")
	t.Logf("Successfully listed IPsec Crypto Profiles, found created profile: %s", profileName)
}

// Test_networkservices_IPsecCryptoProfilesAPIService_DeleteByID tests deleting a profile by ID.
func Test_networkservices_IPsecCryptoProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a profile to delete.
	profileName := "test-ipsec-delete-" + common.GenerateRandomString(6)

	// CORRECTED: Use secure and valid algorithms for the test profile.
	profile := network_services.IpsecCryptoProfiles{
		Folder:  common.StringPtr("Prisma Access"),
		Name:    profileName,
		DhGroup: common.StringPtr("group14"),
		Esp: &network_services.IpsecCryptoProfilesEsp{
			Authentication: []string{"sha256"},      // Changed from "md5"
			Encryption:     []string{"aes-256-cbc"}, // Changed from "des"
		},
		Lifetime: network_services.IpsecCryptoProfilesLifetime{Hours: common.Int32Ptr(4)},
	}
	reqCreate := client.IPsecCryptoProfilesAPI.CreateIPsecCryptoProfiles(context.Background()).IpsecCryptoProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for delete test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	t.Logf("Created IPsec Crypto Profile for Delete test with ID: %s", createdProfileID)

	// Test Delete by ID operation.
	fmt.Printf("Attempting to delete IPsec Crypto Profile with ID: %s\n", createdProfileID)
	reqDel := client.IPsecCryptoProfilesAPI.DeleteIPsecCryptoProfilesByID(context.Background(), createdProfileID)
	httpResDel, errDel := reqDel.Execute()

	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful.
	require.NoError(t, errDel, "Failed to delete profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
	t.Logf("Successfully deleted IPsec Crypto Profile: %s", createdProfileID)

	// Verify deletion by trying to get the profile again (should fail).
	reqGetById := client.IPsecCryptoProfilesAPI.GetIPsecCrytoProfilesByID(context.Background(), createdProfileID)
	getRes, httpResGet, errGet := reqGetById.Execute()

	assert.Error(t, errGet, "Getting deleted profile should fail")
	if httpResGet != nil {
		assert.NotEqual(t, 200, httpResGet.StatusCode, "Should not return 200 for deleted profile")
	}
	assert.Nil(t, getRes, "Response should be nil for a deleted profile")
	t.Logf("Verified IPsec Crypto Profile deletion: %s", createdProfileID)
}
