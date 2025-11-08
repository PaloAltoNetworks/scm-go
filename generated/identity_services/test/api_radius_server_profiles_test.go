package identity_services

/*
 * Identity Services Testing
 *
 * Test_identityservices_RADIUSServerProfilesAPIService_
 */

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Replace with your actual imports for common utils and generated client
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// --- Helper Functions ---

// generateRadiusServerProfileName creates a unique name for the profile.
func generateRadiusServerProfileName(base string) string {
	return base + common.GenerateRandomString(4)
}

// createTestRadiusServerProfile creates a minimal RadiusServerProfiles object for testing.
// NOTE: The 'id' field is now optional and should be left unset for creation.
func createTestRadiusServerProfile(t *testing.T, profileName string) identity_services.RadiusServerProfiles {
	// 1. Create a minimal server inner struct
	serverName := "radius-server-1"
	ipAddress := "10.1.1.1"
	secret := "mySecureSecret123"
	port := int32(1812)

	serverInner := identity_services.NewRadiusServerProfilesServerInner()
	serverInner.SetName(serverName)
	serverInner.SetIpAddress(ipAddress)
	serverInner.SetSecret(secret)
	serverInner.SetPort(port)

	// 2. Define the Protocol (using PAP for simplicity in required fields)
	protocol := identity_services.NewRadiusServerProfilesProtocol()
	protocol.SetPAP(make(map[string]interface{}))

	// 3. Create the main profile
	// Required fields: name, protocol, server
	profile := identity_services.NewRadiusServerProfiles(
		profileName,
		*protocol,
		[]identity_services.RadiusServerProfilesServerInner{*serverInner},
	)

	// Set optional fields
	retries := int32(5)
	timeout := int32(10)
	profile.SetRetries(retries)
	profile.SetTimeout(timeout)
	profile.SetFolder("All")

	return *profile
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_RADIUSServerProfilesAPIService__Create tests the creation of a RADIUS Server Profile.
func Test_identityservices_RADIUSServerProfilesAPIService__Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := generateRadiusServerProfileName("scm-radius-create-")

	// ID is NOT provided in the payload now
	radiusProfile := createTestRadiusServerProfile(t, profileName)

	t.Logf("Creating RADIUS Server Profile with name: %s", profileName)
	req := client.RADIUSServerProfilesAPI.CreateRADIUSServerProfiles(context.Background()).RadiusServerProfiles(radiusProfile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create RADIUS Server Profile")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	// REQUIRE a generated ID from the API
	require.NotNil(t, res.Id, "Created profile should have a generated ID")
	createdID := *res.Id

	// Cleanup the created profile
	defer func() {
		t.Logf("Cleaning up RADIUS Server Profile with ID: %s", createdID)
		_, errDel := client.RADIUSServerProfilesAPI.DeleteRADIUSServerProfilesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete RADIUS Server Profile during cleanup")
	}()

	t.Logf("Successfully created RADIUS Server Profile ID: %s", createdID)

	// Verify the response matches key input fields
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
	assert.Equal(t, int32(5), res.GetRetries(), "Retries should match the set value")
	require.NotEmpty(t, res.Server, "Server list should not be empty")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_RADIUSServerProfilesAPIService__GetByID tests retrieving a RADIUS Server Profile by ID.
func Test_identityservices_RADIUSServerProfilesAPIService__GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := generateRadiusServerProfileName("scm-radius-get-")
	radiusProfile := createTestRadiusServerProfile(t, profileName)

	// Setup: Create a profile first and capture the generated ID
	createRes, _, err := client.RADIUSServerProfilesAPI.CreateRADIUSServerProfiles(context.Background()).RadiusServerProfiles(radiusProfile).Execute()
	require.NoError(t, err, "Failed to create profile for get test setup")
	createdID := *createRes.Id

	defer func() {
		client.RADIUSServerProfilesAPI.DeleteRADIUSServerProfilesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the profile
	getRes, httpResGet, errGet := client.RADIUSServerProfilesAPI.GetRADIUSServerProfilesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get RADIUS Server Profile by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, createdID, getRes.GetId(), "Retrieved ID should match the created ID")
	assert.Equal(t, profileName, getRes.Name, "Retrieved name should match")
	assert.Equal(t, int32(10), getRes.GetTimeout(), "Timeout should match setup value")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_RADIUSServerProfilesAPIService__Update tests updating a RADIUS Server Profile.
// Test_identityservices_RADIUSServerProfilesAPIService__Update tests updating a RADIUS Server Profile.
func Test_identityservices_RADIUSServerProfilesAPIService__Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := generateRadiusServerProfileName("scm-radius-update-")
	targetFolder := "All"

	// 1. Setup: Create a profile first
	radiusProfile := createTestRadiusServerProfile(t, profileName)
	// Ensure the folder is set for CREATION, as this usually determines the resource scope.
	radiusProfile.SetFolder(targetFolder)

	createRes, _, err := client.RADIUSServerProfilesAPI.CreateRADIUSServerProfiles(context.Background()).RadiusServerProfiles(radiusProfile).Execute()
	require.NoError(t, err, "Failed to create profile for update test setup")
	createdID := *createRes.Id

	defer func() {
		// Cleanup logic uses the ID
		client.RADIUSServerProfilesAPI.DeleteRADIUSServerProfilesByID(context.Background(), createdID).Execute()
	}()

	// 2. Prepare updated profile object
	updatedRetries := int32(2)  // Value to update
	updatedTimeout := int32(90) // Value to update

	// Create a new payload using the ORIGINAL profile name
	updatedProfile := createTestRadiusServerProfile(t, profileName)

	// *** CRITICAL CHANGE: DO NOT call SetId(createdID) or SetFolder(targetFolder) ***
	// We are only setting the fields that we want to explicitly update.

	updatedProfile.SetRetries(updatedRetries)
	updatedProfile.SetTimeout(updatedTimeout)
	updatedProfile.SetFolder(targetFolder)

	// 3. Test: Update the profile
	// The createdID still identifies the target resource in the URL path.
	updateRes, httpResUpdate, errUpdate := client.RADIUSServerProfilesAPI.UpdateRADIUSServerProfilesByID(context.Background(), createdID).
		RadiusServerProfiles(updatedProfile).
		Execute()

	require.NoError(t, errUpdate, "Failed to update RADIUS Server Profile")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// 4. Verify the changes
	// Check original fields are preserved and updated fields are changed.
	assert.Equal(t, profileName, updateRes.Name, "Name should remain the original profile name")
	assert.Equal(t, createdID, updateRes.GetId(), "ID should be present in the response")
	assert.Equal(t, targetFolder, updateRes.GetFolder(), "Folder should be present in the response")

	assert.Equal(t, updatedRetries, updateRes.GetRetries(), "Retries should be updated")
	assert.Equal(t, updatedTimeout, updateRes.GetTimeout(), "Timeout should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_RADIUSServerProfilesAPIService__List tests listing RADIUS Server Profiles.
func Test_identityservices_RADIUSServerProfilesAPIService__List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := generateRadiusServerProfileName("scm-radius-list-")
	radiusProfile := createTestRadiusServerProfile(t, profileName)

	// Setup: Create a unique profile to ensure the list filter works and capture the ID
	createRes, _, err := client.RADIUSServerProfilesAPI.CreateRADIUSServerProfiles(context.Background()).RadiusServerProfiles(radiusProfile).Execute()
	require.NoError(t, err, "Failed to create profile for list test setup")
	createdID := *createRes.Id

	defer func() {
		client.RADIUSServerProfilesAPI.DeleteRADIUSServerProfilesByID(context.Background(), createdID).Execute()
	}()

	// Test: List the profiles, filtering by name and folder (if applicable)
	listRes, httpResList, errList := client.RADIUSServerProfilesAPI.ListRADIUSServerProfiles(context.Background()).
		Folder("All").
		Execute()

	require.NoError(t, errList, "Failed to list RADIUS Server Profiles")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_RADIUSServerProfilesAPIService__DeleteByID tests deleting a RADIUS Server Profile.
func Test_identityservices_RADIUSServerProfilesAPIService__DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := generateRadiusServerProfileName("scm-radius-delete-")
	radiusProfile := createTestRadiusServerProfile(t, profileName)

	// Setup: Create a profile first and capture the generated ID
	createRes, _, err := client.RADIUSServerProfilesAPI.CreateRADIUSServerProfiles(context.Background()).RadiusServerProfiles(radiusProfile).Execute()
	require.NoError(t, err, "Failed to create profile for delete test setup")
	createdID := *createRes.Id

	// Test: Delete the profile
	httpResDel, errDel := client.RADIUSServerProfilesAPI.DeleteRADIUSServerProfilesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete RADIUS Server Profile")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}
