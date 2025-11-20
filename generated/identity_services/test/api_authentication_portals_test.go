package identity_services

/*
 * Authentication Portals Testing
 *
 * Test_identityservices_AuthenticationPortalsAPIService_
 */

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Assuming these imports are necessary for your environment

	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// --- Test Constants ---

// ⚠️ IMPORTANT: Use this static IP for all portal creation tests.
const TEST_REDIRECT_HOST = "192.168.255.254"

// createTestAuthPortal creates an AuthenticationPortals object using the fixed host IP.
func createTestAuthPortal(t *testing.T) identity_services.AuthenticationPortals {
	// RedirectHost is the only required field in the constructor
	p := identity_services.NewAuthenticationPortals(TEST_REDIRECT_HOST)

	// Set optional fields
	var gpPort int32 = 10
	var idleT int32 = 10
	var timer int32 = 12

	p.SetFolder("All")
	p.SetAuthenticationProfile("test_auth_profile")
	p.SetCertificateProfile("EDL-Hosting-Service-Profile")
	p.SetTlsServiceProfile("test_svc_profile")
	p.SetGpUdpPort(gpPort)
	p.SetIdleTimer(idleT)
	p.SetTimer(timer)

	return *p
}

// cleanupPortal ensures the portal is deleted after a test attempt.
func cleanupPortal(t *testing.T, client *identity_services.APIClient, id string) {
	if id == "" {
		t.Log("Cleanup skipped: Portal ID is empty.")
		return
	}

	t.Logf("Cleaning up Authentication Portal with ID: %s", id)
	_, errDel := client.AuthenticationPortalsAPI.DeleteAuthenticationPortalsByID(context.Background(), id).Execute()

	// Log cleanup errors but don't fail the test suite on cleanup issues
	if errDel != nil {
		t.Logf("Warning: Failed to delete portal ID %s during cleanup: %v", id, errDel)
	} else {
		t.Logf("Cleanup successful for portal ID: %s", id)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationPortalsAPIService__Create tests the creation of an Auth Portal.
func Test_identityservices_AuthenticationPortalsAPIService__Create(t *testing.T) {

	client := SetupIdentitySvcTestClient(t)
	authPortal := createTestAuthPortal(t)
	var createdID string

	// Cleanup will run regardless of test outcome
	defer func() {
		cleanupPortal(t, client, createdID)
	}()

	t.Logf("Creating Authentication Portal with fixed host: %s", TEST_REDIRECT_HOST)
	req := client.AuthenticationPortalsAPI.CreateAuthenticationPortals(context.Background()).AuthenticationPortals(authPortal)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Authentication Portal")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	// Capture the ID for cleanup
	createdID = res.GetId()
	require.NotEmpty(t, createdID, "Created portal should have a generated ID")

	// Verify the response
	assert.Equal(t, TEST_REDIRECT_HOST, res.RedirectHost, "Created portal host should match fixed IP")
	assert.Equal(t, int32(10), res.GetGpUdpPort(), "GP UDP Port should be 10")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationPortalsAPIService__GetByID tests retrieving an Auth Portal by ID.
func Test_identityservices_AuthenticationPortalsAPIService__GetByID(t *testing.T) {

	client := SetupIdentitySvcTestClient(t)
	authPortal := createTestAuthPortal(t)
	var createdID string

	// Setup: Create a portal first and capture the generated ID
	createRes, _, err := client.AuthenticationPortalsAPI.CreateAuthenticationPortals(context.Background()).AuthenticationPortals(authPortal).Execute()
	require.NoError(t, err, "Failed to create portal for get test setup")
	createdID = createRes.GetId()

	defer func() {
		cleanupPortal(t, client, createdID)
	}()

	// Test: Retrieve the portal
	getRes, httpResGet, errGet := client.AuthenticationPortalsAPI.GetAuthenticationPortalsByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get Authentication Portal by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, createdID, getRes.GetId(), "Retrieved ID should match the created ID")
	assert.Equal(t, TEST_REDIRECT_HOST, getRes.RedirectHost, "Retrieved host should match fixed IP")
	assert.Equal(t, int32(12), getRes.GetTimer(), "Retrieved timer should be preserved")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationPortalsAPIService__Update tests updating an Auth Portal.
func Test_identityservices_AuthenticationPortalsAPIService__Update(t *testing.T) {

	client := SetupIdentitySvcTestClient(t)
	authPortal := createTestAuthPortal(t)
	var createdID string

	// 1. Setup: Create a portal first
	createRes, _, err := client.AuthenticationPortalsAPI.CreateAuthenticationPortals(context.Background()).AuthenticationPortals(authPortal).Execute()
	require.NoError(t, err, "Failed to create portal for update test setup")
	createdID = createRes.GetId()

	defer func() {
		cleanupPortal(t, client, createdID)
	}()

	// 2. Prepare updated portal object (Change port and timer)
	updatedGpPort := int32(20) // CHANGE
	updatedTimer := int32(30)  // CHANGE

	// Start with the base object
	updatedPortal := createTestAuthPortal(t)

	// Apply updates
	updatedPortal.SetGpUdpPort(updatedGpPort)
	updatedPortal.SetTimer(updatedTimer)

	// 3. Test: Update the portal
	updateRes, httpResUpdate, errUpdate := client.AuthenticationPortalsAPI.UpdateAuthenticationPortalsByID(context.Background(), createdID).
		AuthenticationPortals(updatedPortal).
		Execute()

	require.NoError(t, errUpdate, "Failed to update Authentication Portal")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// 4. Verify the changes
	assert.Equal(t, createdID, updateRes.GetId(), "ID should be present in the response")
	assert.Equal(t, updatedGpPort, updateRes.GetGpUdpPort(), "GP UDP Port should be updated")
	assert.Equal(t, updatedTimer, updateRes.GetTimer(), "Timer should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationPortalsAPIService__DeleteByID tests deleting an Auth Portal.
func Test_identityservices_AuthenticationPortalsAPIService__DeleteByID(t *testing.T) {

	client := SetupIdentitySvcTestClient(t)
	authPortal := createTestAuthPortal(t)
	var createdID string

	// Setup: Create a portal first and capture the generated ID
	createRes, _, err := client.AuthenticationPortalsAPI.CreateAuthenticationPortals(context.Background()).AuthenticationPortals(authPortal).Execute()
	require.NoError(t, err, "Failed to create portal for delete test setup")
	createdID = createRes.GetId()

	// Test: Delete the portal
	t.Logf("Deleting Authentication Portal with ID: %s and host: %s", createdID, TEST_REDIRECT_HOST)
	httpResDel, errDel := client.AuthenticationPortalsAPI.DeleteAuthenticationPortalsByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete Authentication Portal")

	// Status 200 OK or 204 No Content are typical for successful delete.
	deleteSuccess := httpResDel.StatusCode == http.StatusOK || httpResDel.StatusCode == http.StatusNoContent
	assert.True(t, deleteSuccess, "Expected 200 OK or 204 No Content status for deletion, got %d", httpResDel.StatusCode)
}

// Test_identityservices_AuthenticationPortalsAPIService__List tests listing Auth Portals.
func Test_identityservices_AuthenticationPortalsAPIService__List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	testFolderName := "All"
	authPortal := createTestAuthPortal(t)
	var createdID string

	// 1. Setup: Create a portal
	t.Logf("Setup: Creating unique portal for list verification using host: %s", TEST_REDIRECT_HOST)
	createRes, _, err := client.AuthenticationPortalsAPI.CreateAuthenticationPortals(context.Background()).AuthenticationPortals(authPortal).Execute()
	require.NoError(t, err, "Failed to create portal for list test setup")
	createdID = createRes.GetId()

	// Teardown: Cleanup the created resource using defer
	defer func() {
		cleanupPortal(t, client, createdID)
	}()

	// 2. Test: List the portals, filtering by folder and host.
	t.Logf("Test: Listing portals filtered by folder: %s and host: %s", testFolderName, TEST_REDIRECT_HOST)
	listRes, httpResList, errList := client.AuthenticationPortalsAPI.ListAuthenticationPortals(context.Background()).
		Folder(testFolderName).
		Execute()

	// 3. Assertions
	require.NoError(t, errList, "Failed to list Authentication Portals")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}
