/*
 * Network Services Testing
 *
 * IKECryptoProfilesAPIService
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

// Test_networkservices_IKECryptoProfilesAPIService_Create tests the creation of an IKE Crypto Profile.
func Test_networkservices_IKECryptoProfilesAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupNetworkSvcTestClient(t)

	// Create a valid IKE Crypto Profile object with a unique name.
	profileName := "test-ike-create-" + common.GenerateRandomString(6)
	profile := network_services.IkeCryptoProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		Hash:   []string{"sha256", "sha384"},
		DhGroup: []string{
			"group14",
		},
		Encryption: []string{"aes-256-cbc"},
		Lifetime: &network_services.IkeCryptoProfilesLifetime{
			Hours: common.Int32Ptr(8),
		},
	}

	fmt.Printf("Attempting to create IKE Crypto Profile with name: %s\n", profile.Name)

	// Make the create request to the API.
	req := client.IKECryptoProfilesAPI.CreateIKECryptoProfiles(context.Background()).IkeCryptoProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create IKE Crypto Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created profile should have an ID")

	createdProfileID := *res.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer the cleanup of the created profile.
	defer func() {
		t.Logf("Cleaning up IKE Crypto Profile with ID: %s", createdProfileID)
		reqDel := client.IKECryptoProfilesAPI.DeleteIKECryptoProfilesByID(context.Background(), createdProfileID)
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
	assert.Equal(t, []string{"sha256", "sha384"}, res.Hash, "Hash algorithms should match")
	t.Logf("Successfully created and validated IKE Crypto Profile: %s with ID: %s", profile.Name, createdProfileID)
}

// Test_networkservices_IKECryptoProfilesAPIService_GetByID tests retrieving a profile by its ID.
func Test_networkservices_IKECryptoProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a profile to retrieve.
	profileName := "test-ike-get-" + common.GenerateRandomString(6)
	profile := network_services.IkeCryptoProfiles{
		Folder:     common.StringPtr("Shared"),
		Name:       profileName,
		Hash:       []string{"sha512"},
		DhGroup:    []string{"group20"},
		Encryption: []string{"aes-128-gcm"},
	}

	reqCreate := client.IKECryptoProfilesAPI.CreateIKECryptoProfiles(context.Background()).IkeCryptoProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer cleanup.
	defer func() {
		t.Logf("Cleaning up IKE Crypto Profile with ID: %s", createdProfileID)
		_, errDel := client.IKECryptoProfilesAPI.DeleteIKECryptoProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
	}()

	// Test Get by ID operation.
	fmt.Printf("Attempting to get IKE Crypto Profile with ID: %s\n", createdProfileID)
	reqGetById := client.IKECryptoProfilesAPI.GetIKECryptoProfilesByID(context.Background(), createdProfileID)
	getRes, httpResGet, errGet := reqGetById.Execute()

	if errGet != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful.
	require.NoError(t, errGet, "Failed to get profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
	assert.Equal(t, []string{"sha512"}, getRes.Hash, "Hash should match")
	assert.Equal(t, *createRes.Id, *getRes.Id, "Profile ID should match")
	t.Logf("Successfully retrieved IKE Crypto Profile: %s", getRes.Name)
}

// Test_networkservices_IKECryptoProfilesAPIService_Update tests updating an existing profile.
func Test_networkservices_IKECryptoProfilesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a profile to update.
	profileName := "test-ike-update-" + common.GenerateRandomString(6)
	profile := network_services.IkeCryptoProfiles{
		Folder:     common.StringPtr("Shared"),
		Name:       profileName,
		Hash:       []string{"sha256"},
		DhGroup:    []string{"group14"},
		Encryption: []string{"aes-256-cbc"},
	}

	reqCreate := client.IKECryptoProfilesAPI.CreateIKECryptoProfiles(context.Background()).IkeCryptoProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for update test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer cleanup.
	defer func() {
		t.Logf("Cleaning up IKE Crypto Profile with ID: %s", createdProfileID)
		_, errDel := client.IKECryptoProfilesAPI.DeleteIKECryptoProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
	}()

	// Define the update payload.
	updatedProfile := network_services.IkeCryptoProfiles{
		Name:       profileName,
		Hash:       []string{"sha512"},
		DhGroup:    []string{"group20"},
		Encryption: []string{"aes-256-gcm"},
		Lifetime: &network_services.IkeCryptoProfilesLifetime{
			Hours: common.Int32Ptr(24),
		},
	}

	fmt.Printf("Attempting to update IKE Crypto Profile with ID: %s\n", createdProfileID)
	reqUpdate := client.IKECryptoProfilesAPI.UpdateIKECryptoProfilesByID(context.Background(), createdProfileID).IkeCryptoProfiles(updatedProfile)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()

	if errUpdate != nil {
		handleAPIError(err)
	}

	// Verify the update was successful.
	require.NoError(t, errUpdate, "Failed to update profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, []string{"sha512"}, updateRes.Hash, "Hash should be updated")
	assert.Equal(t, []string{"aes-256-gcm"}, updateRes.Encryption, "Encryption should be updated")
	require.NotNil(t, updateRes.Lifetime.Hours, "Lifetime hours should be updated")
	assert.Equal(t, int32(24), *updateRes.Lifetime.Hours, "Lifetime should be updated")

	t.Logf("Successfully updated IKE Crypto Profile: %s", profileName)
}

// Test_networkservices_IKECryptoProfilesAPIService_List tests listing IKE Crypto Profiles.
func Test_networkservices_IKECryptoProfilesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a profile to ensure it appears in the list.
	profileName := "test-ike-list-" + common.GenerateRandomString(6)
	profile := network_services.IkeCryptoProfiles{
		Folder:     common.StringPtr("Shared"),
		Name:       profileName,
		Hash:       []string{"sha256"},
		DhGroup:    []string{"group14"},
		Encryption: []string{"aes-256-cbc"},
	}

	reqCreate := client.IKECryptoProfilesAPI.CreateIKECryptoProfiles(context.Background()).IkeCryptoProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for list test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer cleanup.
	defer func() {
		t.Logf("Cleaning up IKE Crypto Profile with ID: %s", createdProfileID)
		_, errDel := client.IKECryptoProfilesAPI.DeleteIKECryptoProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
	}()

	// Test List operation.
	fmt.Println("Attempting to list IKE Crypto Profiles")
	reqList := client.IKECryptoProfilesAPI.ListIKECryptoProfiles(context.Background()).Folder("Shared").Limit(10000)
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
			assert.Equal(t, []string{"sha256"}, p.Hash, "Hash should match")
			break
		}
	}
	assert.True(t, foundProfile, "Created profile should be found in the list")
	t.Logf("Successfully listed IKE Crypto Profiles, found created profile: %s", profileName)
}

// Test_networkservices_IKECryptoProfilesAPIService_DeleteByID tests deleting a profile by ID.
func Test_networkservices_IKECryptoProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a profile to delete.
	profileName := "test-ike-delete-" + common.GenerateRandomString(6)
	profile := network_services.IkeCryptoProfiles{
		Folder:     common.StringPtr("Shared"),
		Name:       profileName,
		Hash:       []string{"sha1"},
		DhGroup:    []string{"group5"},
		Encryption: []string{"3des"},
	}
	reqCreate := client.IKECryptoProfilesAPI.CreateIKECryptoProfiles(context.Background()).IkeCryptoProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for delete test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	t.Logf("Created IKE Crypto Profile for Delete test with ID: %s", createdProfileID)

	// Test Delete by ID operation.
	fmt.Printf("Attempting to delete IKE Crypto Profile with ID: %s\n", createdProfileID)
	reqDel := client.IKECryptoProfilesAPI.DeleteIKECryptoProfilesByID(context.Background(), createdProfileID)
	httpResDel, errDel := reqDel.Execute()

	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful.
	require.NoError(t, errDel, "Failed to delete profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
	t.Logf("Successfully deleted IKE Crypto Profile: %s", createdProfileID)
}
