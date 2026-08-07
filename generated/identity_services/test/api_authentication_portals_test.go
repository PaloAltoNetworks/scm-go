package identity_services

/*
 * Authentication Portals Testing
 *
 * Test_identityservices_AuthenticationPortalsAPIService_
 *
 * Note: AuthenticationPortals is a singleton object (one per folder).
 * Tests handle this by checking if a portal exists before creating.
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

// createTestAuthPortal creates an AuthenticationPortals object.
func createTestAuthPortal(t *testing.T) identity_services.AuthenticationPortals {
	p := identity_services.NewAuthenticationPortals()

	// Set optional fields
	var gpPort int32 = 10
	var idleT int32 = 10
	var timer int32 = 12

	p.SetFolder("All")
	p.SetCertificateProfile("EDL-Hosting-Service-Profile")
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
// Note: AuthenticationPortals is a singleton object (one per folder).
// This test only runs if no portal exists, and does NOT cleanup so the portal remains for other tests.
func Test_identityservices_AuthenticationPortalsAPIService__Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	testFolderName := "All"

	// First check if a portal already exists
	listRes, _, errList := client.AuthenticationPortalsAPI.ListAuthenticationPortals(context.Background()).
		Folder(testFolderName).
		Execute()
	require.NoError(t, errList, "Failed to list Authentication Portals")

	if len(listRes.Data) > 0 {
		t.Log("AuthenticationPortals already exists - singleton object, test passes")
		return
	}

	// Create bare minimum portal
	authPortal := identity_services.NewAuthenticationPortals()
	authPortal.SetFolder(testFolderName)

	t.Log("Creating Authentication Portal (bare minimum)")
	req := client.AuthenticationPortalsAPI.CreateAuthenticationPortals(context.Background()).AuthenticationPortals(*authPortal)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Authentication Portal")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	// Verify the response
	require.NotEmpty(t, res.GetId(), "Created portal should have a generated ID")
	t.Logf("Successfully created Authentication Portal with ID: %s", res.GetId())
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationPortalsAPIService__GetByID tests retrieving an Auth Portal by ID.
// Note: AuthenticationPortals is a singleton object, so we list first to get an existing ID.
func Test_identityservices_AuthenticationPortalsAPIService__GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	testFolderName := "All"

	// First, list to get an existing portal ID
	listRes, _, errList := client.AuthenticationPortalsAPI.ListAuthenticationPortals(context.Background()).
		Folder(testFolderName).
		Execute()
	require.NoError(t, errList, "Failed to list Authentication Portals")
	require.NotNil(t, listRes, "List response should not be nil")

	// Skip if no portal exists (singleton may have been deleted)
	if len(listRes.Data) == 0 {
		t.Skip("Skipping: No AuthenticationPortals exist - singleton may have been deleted")
	}

	// Get the ID from the first portal in the list
	existingID := listRes.Data[0].GetId()
	t.Logf("Using existing portal ID: %s", existingID)

	// Test: Retrieve the portal by ID
	getRes, httpResGet, errGet := client.AuthenticationPortalsAPI.GetAuthenticationPortalsByID(context.Background(), existingID).Execute()

	require.NoError(t, errGet, "Failed to get Authentication Portal by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, existingID, getRes.GetId(), "Retrieved ID should match the existing ID")
	t.Logf("Successfully retrieved Authentication Portal with ID: %s", getRes.GetId())
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationPortalsAPIService__Update tests updating an Auth Portal.
// Note: AuthenticationPortals is a singleton object, so we list first to get an existing ID.
func Test_identityservices_AuthenticationPortalsAPIService__Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	testFolderName := "All"

	// 1. List to get an existing portal ID
	listRes, _, errList := client.AuthenticationPortalsAPI.ListAuthenticationPortals(context.Background()).
		Folder(testFolderName).
		Execute()
	require.NoError(t, errList, "Failed to list Authentication Portals")
	require.NotNil(t, listRes, "List response should not be nil")

	// Skip if no portal exists (singleton may have been deleted)
	if len(listRes.Data) == 0 {
		t.Skip("Skipping: No AuthenticationPortals exist - singleton may have been deleted")
	}

	// Get the existing portal
	existingPortal := listRes.Data[0]
	existingID := existingPortal.GetId()
	t.Logf("Using existing portal ID: %s", existingID)

	// 2. Prepare updated portal object - change idle_timer to 100
	updatedIdleTimer := int32(100)

	// Create update payload from existing portal
	updatedPortal := identity_services.NewAuthenticationPortals()
	updatedPortal.SetIdleTimer(updatedIdleTimer)

	// 3. Test: Update the portal
	t.Logf("Updating portal idle_timer to: %d", updatedIdleTimer)
	updateRes, httpResUpdate, errUpdate := client.AuthenticationPortalsAPI.UpdateAuthenticationPortalsByID(context.Background(), existingID).
		AuthenticationPortals(*updatedPortal).
		Execute()

	require.NoError(t, errUpdate, "Failed to update Authentication Portal")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// 4. Verify the changes
	assert.Equal(t, existingID, updateRes.GetId(), "ID should be present in the response")
	assert.Equal(t, updatedIdleTimer, updateRes.GetIdleTimer(), "Idle timer should be updated to 100")
	t.Logf("Successfully updated portal idle_timer to: %d", updateRes.GetIdleTimer())
}

// ---------------------------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------------------------

// Test_identityservices_AuthenticationPortalsAPIService__DeleteByID tests deleting an Auth Portal.
// Note: AuthenticationPortals is a singleton object, so we list first to get an existing ID.
func Test_identityservices_AuthenticationPortalsAPIService__DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	testFolderName := "All"

	// 1. List to get an existing portal ID
	listRes, _, errList := client.AuthenticationPortalsAPI.ListAuthenticationPortals(context.Background()).
		Folder(testFolderName).
		Execute()
	require.NoError(t, errList, "Failed to list Authentication Portals")
	require.NotNil(t, listRes, "List response should not be nil")
	require.Greater(t, len(listRes.Data), 0, "Expected at least one portal in the list")

	// Get the existing portal ID
	existingID := listRes.Data[0].GetId()
	t.Logf("Attempting to delete existing portal ID: %s", existingID)

	// Test: Delete the portal
	httpResDel, errDel := client.AuthenticationPortalsAPI.DeleteAuthenticationPortalsByID(context.Background(), existingID).Execute()

	require.NoError(t, errDel, "Failed to delete Authentication Portal")

	// Status 200 OK or 204 No Content are typical for successful delete.
	deleteSuccess := httpResDel.StatusCode == http.StatusOK || httpResDel.StatusCode == http.StatusNoContent
	assert.True(t, deleteSuccess, "Expected 200 OK or 204 No Content status for deletion, got %d", httpResDel.StatusCode)
	t.Logf("Successfully deleted portal ID: %s", existingID)
}

// Test_identityservices_AuthenticationPortalsAPIService__List tests listing Auth Portals.
// Note: AuthenticationPortals is a singleton object (one per folder), so we just list without creating.
func Test_identityservices_AuthenticationPortalsAPIService__List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	testFolderName := "All"

	// Test: List the portals, filtering by folder
	t.Logf("Test: Listing portals filtered by folder: %s", testFolderName)
	listRes, httpResList, errList := client.AuthenticationPortalsAPI.ListAuthenticationPortals(context.Background()).
		Folder(testFolderName).
		Execute()

	// Assertions
	require.NoError(t, errList, "Failed to list Authentication Portals")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed Authentication Portals, total: %d", listRes.GetTotal())
}
