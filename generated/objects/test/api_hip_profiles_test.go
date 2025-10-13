/*
Objects Testing HIPProfilesAPIService
*/
package objects

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)

// Test_objects_HIPProfilesAPIService_Create tests the creation of a HIP profile object.
// This test creates a new HIP profile and then deletes it to ensure proper cleanup.
func Test_objects_HIPProfilesAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a valid HIP profile object with a unique name to avoid conflicts.
	createdHipProfileName := "test-hip-create-" + common.GenerateRandomString(6)
	hipProfile := objects.HipProfiles{
		Folder:      common.StringPtr("Prisma Access"),
		Name:        createdHipProfileName,
		Description: common.StringPtr("Test HIP profile for create API"),
		Match:       "\"is-win\" and \"is-anti-malware-and-rtp-enabled\"",
	}

	fmt.Printf("Creating HIP profile with name: %s\n", hipProfile.Name)

	// Make the create request to the API.
	req := client.HIPProfilesAPI.CreateHIPProfiles(context.Background()).HipProfiles(hipProfile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create HIP profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties.
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdHipProfileName, res.Name, "Created HIP profile name should match")
	assert.True(t, *res.Folder == "Shared" || *res.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, hipProfile.Match, res.Match, "Match criteria should match")
	require.NotNil(t, res.Id, "Created HIP profile should have an ID")
	assert.NotEmpty(t, res.Id, "Created HIP profile ID should not be empty")

	// Use the ID from the response object.
	createdHipProfileID := res.Id
	t.Logf("Successfully created HIP profile: %s with ID: %s", hipProfile.Name, createdHipProfileID)

	// Cleanup: Delete the created HIP profile to maintain test isolation.
	reqDel := client.HIPProfilesAPI.DeleteHIPProfilesByID(context.Background(), createdHipProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete HIP profile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up HIP profile: %s", createdHipProfileID)
}

// Test_objects_HIPProfilesAPIService_GetByID tests retrieving a HIP profile by its ID.
// This test creates a HIP profile, retrieves it by ID, then deletes it.
func Test_objects_HIPProfilesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a HIP profile first to have something to retrieve.
	createdHipProfileName := "test-hip-get-" + common.GenerateRandomString(6)
	hipProfile := objects.HipProfiles{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdHipProfileName,
		Match:  "\"is-win\" and \"is-anti-malware-and-rtp-enabled\"",
	}

	// Create the HIP profile via API.
	req := client.HIPProfilesAPI.CreateHIPProfiles(context.Background()).HipProfiles(hipProfile)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create HIP profile for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdHipProfileID := createRes.Id
	require.NotEmpty(t, createdHipProfileID, "Created HIP profile ID should not be empty")

	// Test Get by ID operation.
	reqGetById := client.HIPProfilesAPI.GetHIPProfilesByID(context.Background(), createdHipProfileID)
	getRes, httpResGet, errGet := reqGetById.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful.
	require.NoError(t, errGet, "Failed to get HIP profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties.
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdHipProfileName, getRes.Name, "HIP profile name should match")
	assert.Equal(t, hipProfile.Match, getRes.Match, "Match criteria should match")
	assert.Equal(t, createdHipProfileID, getRes.Id, "HIP profile ID should match")

	t.Logf("Successfully retrieved HIP profile: %s", getRes.Name)

	// Cleanup: Delete the created HIP profile.
	reqDel := client.HIPProfilesAPI.DeleteHIPProfilesByID(context.Background(), createdHipProfileID)
	_, errDel := reqDel.Execute()
	require.NoError(t, errDel, "Failed to delete HIP profile during cleanup")

	t.Logf("Successfully cleaned up HIP profile: %s", createdHipProfileID)
}

// Test_objects_HIPProfilesAPIService_Update tests updating an existing HIP profile.
// This test creates a HIP profile, updates it, then deletes it.
func Test_objects_HIPProfilesAPIService_Update(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a HIP profile first to have something to update.
	createdHipProfileName := "test-hip-update-" + common.GenerateRandomString(6)
	hipProfile := objects.HipProfiles{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdHipProfileName,
		Match:  "\"is-win\" and \"is-anti-malware-and-rtp-enabled\"",
	}

	// Create the HIP profile via API.
	req := client.HIPProfilesAPI.CreateHIPProfiles(context.Background()).HipProfiles(hipProfile)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create HIP profile for update test")
	createdHipProfileID := createRes.Id
	require.NotEmpty(t, createdHipProfileID, "Created HIP profile ID should not be empty")

	// Test Update operation with modified fields.
	updatedHipProfile := objects.HipProfiles{
		Description: common.StringPtr("Updated description"),
		Match:       "\"is-win\" and \"is-rtp-enabled\"",
		Name:        createdHipProfileName, // Name must be included in update payload
	}

	reqUpdate := client.HIPProfilesAPI.UpdateHIPProfilesByID(context.Background(), createdHipProfileID).HipProfiles(updatedHipProfile)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful.
	require.NoError(t, errUpdate, "Failed to update HIP profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties.
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, common.StringPtr("Updated description"), updateRes.Description, "Description should be updated")
	assert.Equal(t, updatedHipProfile.Match, updateRes.Match, "Match criteria should be updated")
	assert.Equal(t, createdHipProfileID, updateRes.Id, "HIP profile ID should remain the same")

	t.Logf("Successfully updated HIP profile: %s", createdHipProfileName)

	// Cleanup: Delete the created HIP profile.
	reqDel := client.HIPProfilesAPI.DeleteHIPProfilesByID(context.Background(), createdHipProfileID)
	_, errDel := reqDel.Execute()
	require.NoError(t, errDel, "Failed to delete HIP profile during cleanup")

	t.Logf("Successfully cleaned up HIP profile: %s", createdHipProfileID)
}

// Test_objects_HIPProfilesAPIService_List tests listing HIP profiles.
// This test creates a HIP profile, lists profiles to verify it's included, then deletes it.
func Test_objects_HIPProfilesAPIService_List(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a HIP profile first to have something to list.
	createdHipProfileName := "test-hip-list-" + common.GenerateRandomString(6)
	hipProfile := objects.HipProfiles{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdHipProfileName,
		Match:  "\"is-win\" and \"is-anti-malware-and-rtp-enabled\"",
	}

	// Create the HIP profile via API.
	req := client.HIPProfilesAPI.CreateHIPProfiles(context.Background()).HipProfiles(hipProfile)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create HIP profile for list test")
	createdHipProfileID := createRes.Id
	require.NotEmpty(t, createdHipProfileID, "Created HIP profile ID should not be empty")

	// Test List operation.
	reqList := client.HIPProfilesAPI.ListHIPProfiles(context.Background()).Folder("Prisma Access").Limit(10000)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful.
	require.NoError(t, errList, "Failed to list HIP profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one HIP profile in the list")

	// Verify our created HIP profile is in the list.
	foundProfile := false
	for _, profile := range listRes.Data {
		if profile.Name == createdHipProfileName {
			foundProfile = true
			assert.Equal(t, hipProfile.Match, profile.Match, "Match criteria should match in list")
			assert.True(t, *profile.Folder == "Shared" || *profile.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
			break
		}
	}
	assert.True(t, foundProfile, "Created HIP profile should be found in the list")

	t.Logf("Successfully listed HIP profiles, found created profile: %s", createdHipProfileName)

	// Cleanup: Delete the created HIP profile.
	reqDel := client.HIPProfilesAPI.DeleteHIPProfilesByID(context.Background(), createdHipProfileID)
	_, errDel := reqDel.Execute()
	require.NoError(t, errDel, "Failed to delete HIP profile during cleanup")

	t.Logf("Successfully cleaned up HIP profile: %s", createdHipProfileID)
}

// Test_objects_HIPProfilesAPIService_DeleteByID tests deleting a HIP profile by its ID.
// This test creates a HIP profile, deletes it, then verifies the deletion was successful.
func Test_objects_HIPProfilesAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a HIP profile first to have something to delete.
	createdHipProfileName := "test-hip-delete-" + common.GenerateRandomString(6)
	hipProfile := objects.HipProfiles{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdHipProfileName,
		Match:  "\"is-win\" and \"is-anti-malware-and-rtp-enabled\"",
	}

	// Create the HIP profile via API.
	req := client.HIPProfilesAPI.CreateHIPProfiles(context.Background()).HipProfiles(hipProfile)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create HIP profile for delete test")
	createdHipProfileID := createRes.Id
	require.NotEmpty(t, createdHipProfileID, "Created HIP profile ID should not be empty")

	// Test Delete by ID operation.
	reqDel := client.HIPProfilesAPI.DeleteHIPProfilesByID(context.Background(), createdHipProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful.
	require.NoError(t, errDel, "Failed to delete HIP profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted HIP profile: %s", createdHipProfileID)
}
