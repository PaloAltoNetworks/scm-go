/*
 * Config Setup Testing
 *
 * FoldersAPIService
 */
/*
Config Setup Testing FoldersAPIService
*/
package config_setup

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Replace with your actual common and generated package paths
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/config_setup"
)

// --- Test Cases ---

// Test_config_setup_FoldersAPIService_Create tests the creation of a folder
func Test_config_setup_FoldersAPIService_Create(t *testing.T) {
	t.Skip("API returns 500 Internal Server Error on folder create - folder creation not supported in test environment")
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Create a valid folder object with a unique name
	createdFolderName := "test-folder-create-" + common.GenerateRandomString(6)
	folder := config_setup.Folders{
		Description: common.StringPtr("Test folder for create API testing"),
		Name:        createdFolderName, // Unique test name
		Parent:      "Shared",          // Folders require a parent - using "Shared" as default
	}

	fmt.Printf("Creating folder with name: %s\n", folder.Name)

	// Make the create request to the API
	req := client.FoldersAPI.CreateFolder(context.Background()).Folders(folder)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create folder")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdFolderName, res.Name, "Created folder name should match")
	require.NotNil(t, folder.Description, "Test setup error: expected Description should not be nil")
	require.NotNil(t, res.Description, "API Response error: Description field in response is unexpectedly nil")
	assert.Equal(t, *folder.Description, *res.Description, "Description string value should match the expected value")
	assert.Equal(t, folder.Parent, res.Parent, "Parent should match")
	assert.NotEmpty(t, res.Id, "Created folder should have an ID")

	// Use the ID from the response object
	createdFolderID := res.Id
	t.Logf("Successfully created folder: %s with ID: %s", createdFolderName, createdFolderID)

	// Cleanup: Delete the created folder to maintain test isolation
	reqDel := client.FoldersAPI.DeleteFolderByID(context.Background(), createdFolderID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete folder during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up folder: %s", createdFolderID)
}

// ---

// Test_config_setup_FoldersAPIService_GetByID tests retrieving a folder by its ID
func Test_config_setup_FoldersAPIService_GetByID(t *testing.T) {
	t.Skip("API returns 500 Internal Server Error on folder create - folder creation not supported in test environment")
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Create a folder first to have something to retrieve
	createdFolderName := "test-folder-getbyid-" + common.GenerateRandomString(6)
	folder := config_setup.Folders{
		Description: common.StringPtr("Test folder for Get by ID API testing"),
		Name:        createdFolderName,
		Parent:      "Shared", // Folders require a parent
	}

	// Create the folder via API
	req := client.FoldersAPI.CreateFolder(context.Background()).Folders(folder)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create folder for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdFolderID := createRes.Id
	require.NotEmpty(t, createdFolderID, "Created folder should have an ID")

	// Test Get by ID operation
	reqGetById := client.FoldersAPI.GetFolderByID(context.Background(), createdFolderID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get folder by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdFolderName, getRes.Name, "Folder name should match")
	assert.Equal(t, createdFolderID, getRes.Id, "Folder ID should match")
	assert.Equal(t, "Shared", getRes.Parent, "Parent should match")

	t.Logf("Successfully retrieved folder: %s", getRes.Name)

	// Cleanup: Delete the created folder
	reqDel := client.FoldersAPI.DeleteFolderByID(context.Background(), createdFolderID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete folder during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up folder: %s", createdFolderID)
}

// ---

// Test_config_setup_FoldersAPIService_Update tests updating an existing folder
func Test_config_setup_FoldersAPIService_Update(t *testing.T) {
	t.Skip("API returns 500 Internal Server Error on folder create - folder creation not supported in test environment")
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// 1. Create a folder first to have something to update
	createdFolderName := "test-folder-update-" + common.GenerateRandomString(10)
	initialFolder := config_setup.Folders{
		Description: common.StringPtr("Initial description"),
		Name:        createdFolderName,
		Parent:      "Shared",
	}

	// Create the folder via API
	reqCreate := client.FoldersAPI.CreateFolder(context.Background()).Folders(initialFolder)
	createRes, _, err := reqCreate.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create folder for update test")
	createdFolderID := createRes.Id

	// 2. Test Update operation with modified fields
	updatedFolder := config_setup.Folders{
		Description: common.StringPtr("Updated test folder description"), // Updated field
		Name:        createdFolderName,                                   // Name must be the same
		Parent:      "Shared",                                            // Parent remains the same
	}

	reqUpdate := client.FoldersAPI.UpdateFolderByID(context.Background(), createdFolderID).Folders(updatedFolder)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update folder")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdFolderName, updateRes.Name, "Folder name should remain the same")
	assert.Equal(t, *updatedFolder.Description, *updateRes.Description, "Description should be updated")
	assert.Equal(t, createdFolderID, updateRes.Id, "Folder ID should remain the same")

	t.Logf("Successfully updated folder: %s", createdFolderName)

	// 3. Cleanup: Delete the created folder
	reqDel := client.FoldersAPI.DeleteFolderByID(context.Background(), createdFolderID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete folder during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up folder: %s", createdFolderID)
}

// ---

// Test_config_setup_FoldersAPIService_List tests listing folders (read-only)
func Test_config_setup_FoldersAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Read-only test: list existing folders (Create gives 500, but List works independently)
	listRes, httpResList, errList := client.FoldersAPI.ListFolders(context.Background()).Limit(200).Offset(0).Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list folders")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed folders, total: %d", listRes.GetTotal())
}

// ---

// Test_config_setup_FoldersAPIService_DeleteByID tests deleting a folder by its ID
func Test_config_setup_FoldersAPIService_DeleteByID(t *testing.T) {
	t.Skip("API returns 500 Internal Server Error on folder create - folder creation not supported in test environment")
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// 1. Create a folder first to have something to delete
	createdFolderName := "test-folder-delete-" + common.GenerateRandomString(10)
	folder := config_setup.Folders{
		Description: common.StringPtr("Test folder for delete API testing"),
		Name:        createdFolderName,
		Parent:      "Shared",
	}

	// Create the folder via API
	reqCreate := client.FoldersAPI.CreateFolder(context.Background()).Folders(folder)
	createRes, _, err := reqCreate.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create folder for delete test")
	createdFolderID := createRes.Id
	require.NotEmpty(t, createdFolderID, "Created folder should have an ID")
	t.Logf("Folder created successfully: %s", createdFolderID)

	// 2. Test Delete by ID operation
	reqDel := client.FoldersAPI.DeleteFolderByID(context.Background(), createdFolderID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete folder")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted folder: %s", createdFolderID)

	// 3. Optional verification: Try to get the deleted folder (should return 404)
	reqGet := client.FoldersAPI.GetFolderByID(context.Background(), createdFolderID)
	_, httpResGet, errGet := reqGet.Execute()

	assert.Error(t, errGet, "Expected an error when trying to retrieve deleted folder")
	assert.Equal(t, 404, httpResGet.StatusCode, "Expected 404 Not Found status after deletion")
}

// Test_config_setup_FoldersAPIService_FetchFolders tests the FetchFolders convenience method (read-only)
func Test_config_setup_FoldersAPIService_FetchFolders(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Read-only test: Fetch non-existent object (should return nil, nil)
	notFound, err := client.FoldersAPI.FetchFolders(
		context.Background(),
		"non-existent-folder-xyz-12345",
		nil, // folder
		nil, // snippet
		nil, // device
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchFolders correctly returned nil for non-existent object")
}
