package identity_services

/*
 * Authentication Sequences Testing
 *
 * Test_identityservices_AuthenticationSequencesAPIService_
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
	// NOTE: AuthenticationProfiles and AuthenticationSequences structs are assumed available
	// from the surrounding package context.
)

// --- Helper Functions ---

// generateSequenceName creates a unique name for the sequence.
func generateSequenceName(base string) string {
	return base + common.GenerateRandomString(4)
}

// generateProfileName creates a unique name for the profile.
func generateProfileName(base string) string {
	return base + common.GenerateRandomString(4)
}

// setupTestAuthProfile creates an authentication profile and returns its name and a cleanup function
func setupTestAuthProfile(t *testing.T, client *identity_services.APIClient) (string, func()) {
	profileName := generateProfileName("scm-authprofile-")
	authProfile := createTestLocalDBFullConfigProfile(t, profileName)

	t.Logf("Creating Authentication Profile with name: %s", profileName)
	createRes, _, err := client.AuthenticationProfilesAPI.CreateAuthenticationProfiles(context.Background()).
		AuthenticationProfiles(authProfile).
		Execute()
	require.NoError(t, err, "Failed to create Authentication Profile for test setup")

	profileID := createRes.GetId()
	require.NotEmpty(t, profileID, "Created profile should have a generated ID")

	// Return cleanup function
	cleanup := func() {
		t.Logf("Cleaning up Authentication Profile with ID: %s", profileID)
		_, errDel := client.AuthenticationProfilesAPI.DeleteAuthenticationProfilesByID(context.Background(), profileID).Execute()
		require.NoError(t, errDel, "Failed to delete Authentication Profile during cleanup")
	}

	return profileName, cleanup
}

// createTestAuthSequence creates an AuthenticationSequences object.
// IMPORTANT: It no longer sets the AuthenticationProfiles field, allowing the caller to set it dynamically.
func createTestAuthSequence(t *testing.T, sequenceName string) identity_services.AuthenticationSequences {
	// Use NewAuthenticationSequencesWithDefaults for a clean starting point.
	p := identity_services.NewAuthenticationSequencesWithDefaults()

	// Set required and explicit fields:
	p.SetName(sequenceName)
	p.SetFolder("All")
	p.SetUseDomainFindProfile(false)
	// p.SetAuthenticationProfiles([]string{"Test_UI"}) <-- REMOVED THIS HARDCODED VALUE

	// NOTE: For a POST request, the 'Id' field is typically omitted.

	return *p
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationSequencesAPIService__Create tests the creation of an Auth Sequence.
func Test_identityservices_AuthenticationSequencesAPIService__Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	sequenceName := generateSequenceName("scm-authseq-create-")

	// --- 1. SETUP PREREQUISITE: Create the Authentication Profile ---
	profileName, profileCleanup := setupTestAuthProfile(t, client)

	// Ensure the profile is deleted after this test runs.
	defer profileCleanup()

	// --- 2. CREATE SEQUENCE OBJECT: Use the temporary profile name ---
	authSequence := createTestAuthSequence(t, sequenceName)

	// ASSIGN: Override the default list with the dynamically created profile name.
	authSequence.SetAuthenticationProfiles([]string{profileName})

	// --- 3. TEST EXECUTION ---
	t.Logf("Creating Authentication Sequence with name: %s", sequenceName)
	req := client.AuthenticationSequencesAPI.CreateAuthenticationSequences(context.Background()).AuthenticationSequences(authSequence)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Authentication Sequence")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	// REQUIRE a generated ID from the API
	createdID := res.GetId()
	require.NotEmpty(t, createdID, "Created sequence should have a generated ID")

	// --- 4. CLEANUP THE SEQUENCE ---
	defer func() {
		t.Logf("Cleaning up Authentication Sequence with ID: %s", createdID)
		_, errDel := client.AuthenticationSequencesAPI.DeleteAuthenticationSequencesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete Authentication Sequence during cleanup")
	}()

	t.Logf("Successfully created Authentication Sequence ID: %s", createdID)

	// --- 5. VERIFY RESPONSE ---
	assert.Equal(t, sequenceName, res.Name, "Created sequence name should match")
	assert.Equal(t, false, res.GetUseDomainFindProfile(), "use_domain_find_profile should be false")

	// VERIFY: The returned profile list contains the dynamic profile name
	assert.Equal(t, []string{profileName}, res.GetAuthenticationProfiles(), "Authentication profiles list should match the created profile name")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationSequencesAPIService__GetByID tests retrieving an Auth Sequence by ID.
func Test_identityservices_AuthenticationSequencesAPIService__GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	sequenceName := generateSequenceName("scm-authseq-get-")

	// Setup prerequisite profile
	profileName, profileCleanup := setupTestAuthProfile(t, client)
	defer profileCleanup()

	// Create sequence with the profile
	authSequence := createTestAuthSequence(t, sequenceName)
	authSequence.SetAuthenticationProfiles([]string{profileName})

	// Setup: Create a sequence first and capture the generated ID
	createRes, _, err := client.AuthenticationSequencesAPI.CreateAuthenticationSequences(context.Background()).AuthenticationSequences(authSequence).Execute()
	require.NoError(t, err, "Failed to create sequence for get test setup")
	createdID := createRes.GetId()

	defer func() {
		client.AuthenticationSequencesAPI.DeleteAuthenticationSequencesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the sequence
	getRes, httpResGet, errGet := client.AuthenticationSequencesAPI.GetAuthenticationSequencesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get Authentication Sequence by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, createdID, getRes.GetId(), "Retrieved ID should match the created ID")
	assert.Equal(t, sequenceName, getRes.Name, "Retrieved name should match")
	assert.Equal(t, []string{profileName}, getRes.GetAuthenticationProfiles(), "Authentication profiles list should be preserved")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationSequencesAPIService__Update tests updating an Auth Sequence.
func Test_identityservices_AuthenticationSequencesAPIService__Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	sequenceName := generateSequenceName("scm-authseq-update-")
	targetFolder := "All"

	// Setup prerequisite profiles
	profileName1, profileCleanup1 := setupTestAuthProfile(t, client)
	defer profileCleanup1()
	profileName2, profileCleanup2 := setupTestAuthProfile(t, client)
	defer profileCleanup2()

	// 1. Setup: Create a sequence first
	authSequence := createTestAuthSequence(t, sequenceName)
	authSequence.SetAuthenticationProfiles([]string{profileName1}) // Start with first profile
	authSequence.SetFolder(targetFolder)

	createRes, _, err := client.AuthenticationSequencesAPI.CreateAuthenticationSequences(context.Background()).AuthenticationSequences(authSequence).Execute()
	require.NoError(t, err, "Failed to create sequence for update test setup")
	createdID := createRes.GetId()

	defer func() {
		client.AuthenticationSequencesAPI.DeleteAuthenticationSequencesByID(context.Background(), createdID).Execute()
	}()

	// 2. Prepare updated sequence object
	updatedAuthProfiles := []string{profileName2} // Use the second profile for update
	updatedUseDomain := true

	// Use the original helper structure
	updatedSequence := createTestAuthSequence(t, sequenceName)

	// Set the fields we want to explicitly update. ID is handled by the API path.
	updatedSequence.SetAuthenticationProfiles(updatedAuthProfiles)
	updatedSequence.SetUseDomainFindProfile(updatedUseDomain)
	updatedSequence.SetFolder(targetFolder)

	// 3. Test: Update the sequence
	updateRes, httpResUpdate, errUpdate := client.AuthenticationSequencesAPI.UpdateAuthenticationSequencesByID(context.Background(), createdID).
		AuthenticationSequences(updatedSequence).
		Execute()

	require.NoError(t, errUpdate, "Failed to update Authentication Sequence")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// 4. Verify the changes
	assert.Equal(t, createdID, updateRes.GetId(), "ID should be present in the response")
	assert.Equal(t, updatedUseDomain, updateRes.GetUseDomainFindProfile(), "use_domain_find_profile should be updated")
	assert.Equal(t, updatedAuthProfiles, updateRes.GetAuthenticationProfiles(), "Authentication profiles list should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationSequencesAPIService__List tests listing Auth Sequences.
func Test_identityservices_AuthenticationSequencesAPIService__List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	sequenceName := generateSequenceName("scm-authseq-list-")

	// Setup prerequisite profile
	profileName, profileCleanup := setupTestAuthProfile(t, client)
	defer profileCleanup()

	authSequence := createTestAuthSequence(t, sequenceName)
	authSequence.SetAuthenticationProfiles([]string{profileName})

	// Setup: Create a unique sequence to ensure the list filter works
	createRes, _, err := client.AuthenticationSequencesAPI.CreateAuthenticationSequences(context.Background()).AuthenticationSequences(authSequence).Execute()
	require.NoError(t, err, "Failed to create sequence for list test setup")
	createdID := createRes.GetId()

	defer func() {
		client.AuthenticationSequencesAPI.DeleteAuthenticationSequencesByID(context.Background(), createdID).Execute()
	}()

	// Test: List the sequences, filtering by folder
	listRes, httpResList, errList := client.AuthenticationSequencesAPI.ListAuthenticationSequences(context.Background()).
		Folder("All").
		Execute()

	require.NoError(t, errList, "Failed to list Authentication Sequences")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationSequencesAPIService__DeleteByID tests deleting an Auth Sequence.
func Test_identityservices_AuthenticationSequencesAPIService__DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	sequenceName := generateSequenceName("scm-authseq-delete-")

	// Setup prerequisite profile
	profileName, profileCleanup := setupTestAuthProfile(t, client)
	defer profileCleanup()

	authSequence := createTestAuthSequence(t, sequenceName)
	authSequence.SetAuthenticationProfiles([]string{profileName})

	// Setup: Create a sequence first and capture the generated ID
	createRes, _, err := client.AuthenticationSequencesAPI.CreateAuthenticationSequences(context.Background()).AuthenticationSequences(authSequence).Execute()
	require.NoError(t, err, "Failed to create sequence for delete test setup")
	createdID := createRes.GetId()

	// Test: Delete the sequence
	httpResDel, errDel := client.AuthenticationSequencesAPI.DeleteAuthenticationSequencesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete Authentication Sequence")
	// Status 200 OK or 204 No Content are typical for successful delete.
	assert.True(t, httpResDel.StatusCode == http.StatusOK || httpResDel.StatusCode == http.StatusNoContent, "Expected 200 OK or 204 No Content status for deletion")
}
