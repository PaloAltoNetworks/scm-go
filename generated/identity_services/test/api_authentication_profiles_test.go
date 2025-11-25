package identity_services

/*
 * Authentication Profiles - Local Database Testing
 *
 * Test_identityservices_AuthenticationProfilesAPIService__LocalDatabase
 */

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Assuming 'common' contains utilities like GenerateRandomString and SetupClient
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
	// Ensure this import path is correct for your generated client
)

// --- Helper Functions ---

// generateAuthProfileName creates a unique name for the profile.
func generateAuthProfileName(base string) string {
	return base + common.GenerateRandomString(4)
}

// createTestLocalDBFullConfigProfile creates an AuthenticationProfiles object
// with all the specified fields, using Local Database for the method.
func createTestLocalDBFullConfigProfile(t *testing.T, profileName string) identity_services.AuthenticationProfiles {
	// 1. Define the Method struct for Local Database
	method := identity_services.NewAuthenticationProfilesMethod()
	// LocalDatabase is map[string]interface{}
	localDbConfig := make(map[string]interface{})
	method.SetLocalDatabase(localDbConfig)

	// 2. Define Lockout
	lockout := identity_services.NewAuthenticationProfilesLockoutWithDefaults() // Assuming a NewWithDefaults helper exists for sub-models
	failedAttempts := int32(9)
	lockoutTime := int32(5)
	lockout.SetFailedAttempts(failedAttempts)
	lockout.SetLockoutTime(lockoutTime)

	// 3. Define Single Sign-On
	sso := identity_services.NewAuthenticationProfilesSingleSignOnWithDefaults() // Assuming NewWithDefaults exists
	realm := "EXAMPLE.COM"
	sso.SetRealm(realm)

	// 4. Create the main profile and set all fields
	profile := identity_services.NewAuthenticationProfiles(profileName)

	// Set the required Name
	profile.SetName(profileName)

	// Set Allow List
	profile.SetAllowList([]string{"all"})

	// Set Lockout, Method, Single Sign-On, Folder
	profile.SetLockout(*lockout)
	profile.SetMethod(*method)
	profile.SetSingleSignOn(*sso)

	// Set other optional fields
	userDomain := "default"
	usernameModifier := "%USERINPUT%"

	profile.SetUserDomain(userDomain)
	profile.SetUsernameModifier(usernameModifier)
	profile.SetFolder("All")

	return *profile
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationProfilesAPIService__CreateLocalDB_Full tests creation with the comprehensive payload.
func Test_identityservices_AuthenticationProfilesAPIService__CreateLocalDB_Full(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := generateAuthProfileName("scm-auth-localdb-full-")

	authProfile := createTestLocalDBFullConfigProfile(t, profileName)

	t.Logf("Creating Authentication Profile (Local DB, Full Config) with name: %s", profileName)
	req := client.AuthenticationProfilesAPI.CreateAuthenticationProfiles(context.Background()).AuthenticationProfiles(authProfile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Authentication Profile (Local DB, Full Config)")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	// REQUIRE a generated ID from the API
	require.NotNil(t, res.Id, "Created profile should have a generated ID")
	createdID := *res.Id

	// Cleanup the created profile
	defer func() {
		t.Logf("Cleaning up Authentication Profile with ID: %s", createdID)
		_, errDel := client.AuthenticationProfilesAPI.DeleteAuthenticationProfilesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete Authentication Profile during cleanup")
	}()

	t.Logf("Successfully created Authentication Profile ID: %s", createdID)

	// --- Comprehensive Assertion Block ---

	// 1. Assert top-level fields
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
	assert.Equal(t, "default", res.GetUserDomain(), "User domain should match 'default'")
	assert.Equal(t, "%USERINPUT%", res.GetUsernameModifier(), "Username modifier should match '%USERINPUT%'")

	// 2. Assert Allow List
	require.NotEmpty(t, res.AllowList, "Allow list should not be empty")
	assert.Contains(t, res.AllowList, "all", "Allow list should contain 'all'")

	// 3. Assert Lockout configuration
	require.NotNil(t, res.Lockout, "Lockout object should not be nil")
	assert.Equal(t, int32(9), res.Lockout.GetFailedAttempts(), "Lockout failed_attempts should be 9")
	assert.Equal(t, int32(5), res.Lockout.GetLockoutTime(), "Lockout lockout_time should be 5")

	// 4. Assert Method (Local Database)
	require.NotNil(t, res.Method, "Method object should not be nil")
	assert.True(t, res.Method.HasLocalDatabase(), "Authentication method should be Local Database")
	assert.False(t, res.Method.HasRadius(), "Authentication method should NOT be Radius")

	// 5. Assert Single Sign-On (SSO) configuration
	require.NotNil(t, res.SingleSignOn, "Single Sign-On object should not be nil")
	assert.Equal(t, "EXAMPLE.COM", res.SingleSignOn.GetRealm(), "SSO realm should match 'EXAMPLE.COM'")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationProfilesAPIService__GetByIDLocalDB tests retrieving a Local DB profile.
func Test_identityservices_AuthenticationProfilesAPIService__GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := generateAuthProfileName("scm-auth-localdb-get-")
	authProfile := createTestLocalDBFullConfigProfile(t, profileName)

	// Setup: Create a profile first and capture the generated ID
	createRes, _, err := client.AuthenticationProfilesAPI.CreateAuthenticationProfiles(context.Background()).AuthenticationProfiles(authProfile).Execute()
	require.NoError(t, err, "Failed to create profile for get test setup")
	createdID := *createRes.Id

	defer func() {
		client.AuthenticationProfilesAPI.DeleteAuthenticationProfilesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the profile
	getRes, httpResGet, errGet := client.AuthenticationProfilesAPI.GetAuthenticationProfilesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get Authentication Profile by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, createdID, getRes.GetId(), "Retrieved ID should match the created ID")
	assert.Equal(t, profileName, getRes.Name, "Retrieved name should match")
	assert.True(t, getRes.Method.HasLocalDatabase(), "Retrieved method should be Local Database")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationProfilesAPIService__UpdateLocalDB tests updating a Local DB profile.
func Test_identityservices_AuthenticationProfilesAPIService__UpdateLocalDB(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := generateAuthProfileName("scm-auth-localdb-update-")
	targetFolder := "All"

	// 1. Setup: Create a profile first
	authProfile := createTestLocalDBFullConfigProfile(t, profileName)
	authProfile.SetFolder(targetFolder)

	createRes, _, err := client.AuthenticationProfilesAPI.CreateAuthenticationProfiles(context.Background()).AuthenticationProfiles(authProfile).Execute()
	require.NoError(t, err, "Failed to create profile for update test setup")
	createdID := *createRes.Id

	defer func() {
		client.AuthenticationProfilesAPI.DeleteAuthenticationProfilesByID(context.Background(), createdID).Execute()
	}()

	// 2. Prepare updated profile object
	updatedUserDomain := "paloaltonetworks.com"

	// Create a new payload using the ORIGINAL profile name
	updatedProfile := createTestLocalDBFullConfigProfile(t, profileName)

	// Set the fields we want to explicitly update
	updatedProfile.SetUserDomain(updatedUserDomain)
	updatedProfile.SetFolder(targetFolder)

	// 3. Test: Update the profile
	updateRes, httpResUpdate, errUpdate := client.AuthenticationProfilesAPI.UpdateAuthenticationProfilesByID(context.Background(), createdID).
		AuthenticationProfiles(updatedProfile).
		Execute()

	require.NoError(t, errUpdate, "Failed to update Authentication Profile")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// 4. Verify the changes
	assert.Equal(t, profileName, updateRes.Name, "Name should remain the original profile name")
	assert.Equal(t, updatedUserDomain, updateRes.GetUserDomain(), "User domain should be updated")
	assert.True(t, updateRes.Method.HasLocalDatabase(), "Authentication method should remain Local Database")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationProfilesAPIService__ListLocalDB tests listing profiles.
func Test_identityservices_AuthenticationProfilesAPIService__ListLocalDB(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := generateAuthProfileName("scm-auth-localdb-list-")
	authProfile := createTestLocalDBFullConfigProfile(t, profileName)

	// Setup: Create a unique profile to ensure the list filter works
	createRes, _, err := client.AuthenticationProfilesAPI.CreateAuthenticationProfiles(context.Background()).AuthenticationProfiles(authProfile).Execute()
	require.NoError(t, err, "Failed to create profile for list test setup")
	createdID := *createRes.Id

	defer func() {
		client.AuthenticationProfilesAPI.DeleteAuthenticationProfilesByID(context.Background(), createdID).Execute()
	}()

	// Test: List the profiles, filtering by folder
	listRes, httpResList, errList := client.AuthenticationProfilesAPI.ListAuthenticationProfiles(context.Background()).
		Folder("All").
		Execute()

	require.NoError(t, errList, "Failed to list Authentication Profiles")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationProfilesAPIService__DeleteByIDLocalDB tests deleting a Local DB profile.
func Test_identityservices_AuthenticationProfilesAPIService__DeleteByIDLocalDB(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := generateAuthProfileName("scm-auth-localdb-delete-")
	authProfile := createTestLocalDBFullConfigProfile(t, profileName)

	// Setup: Create a profile first and capture the generated ID
	createRes, _, err := client.AuthenticationProfilesAPI.CreateAuthenticationProfiles(context.Background()).AuthenticationProfiles(authProfile).Execute()
	require.NoError(t, err, "Failed to create profile for delete test setup")
	createdID := *createRes.Id

	// Test: Delete the profile
	httpResDel, errDel := client.AuthenticationProfilesAPI.DeleteAuthenticationProfilesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete Authentication Profile")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}
