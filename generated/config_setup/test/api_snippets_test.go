/*
 * Config Setup Testing
 *
 * SnippetsAPIService
 */
/*
Objects Testing SnippetsAPIService
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

// Test_config_setup_SnippetsAPIService_Create tests the creation of a snippet
func Test_config_setup_SnippetsAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Create a valid snippet object with a unique name
	createdSnippetName := "test-snippet-create-" + common.GenerateRandomString(6)
	snippet := config_setup.Snippets{
		Description: common.StringPtr("Test snippet for create API testing"),
		Name:        createdSnippetName, // Unique test name
	}

	fmt.Printf("Creating snippet with name: %s\n", snippet.Name)

	// Make the create request to the API
	req := client.SnippetsAPI.CreateSnippet(context.Background()).Snippets(snippet)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create snippet")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdSnippetName, res.Name, "Created snippet name should match")
	require.NotNil(t, snippet.Description, "Test setup error: expected Description should not be nil")
	require.NotNil(t, res.Description, "API Response error: Description field in response is unexpectedly nil")
	assert.Equal(t, *snippet.Description, *res.Description, "Description string value should match the expected value")
	assert.NotEmpty(t, res.Id, "Created snippet should have an ID")

	// Use the ID from the response object
	createdSnippetID := res.Id
	t.Logf("Successfully created snippet: %s with ID: %s", createdSnippetName, createdSnippetID)

	// Cleanup: Delete the created snippet to maintain test isolation
	reqDel := client.SnippetsAPI.DeleteSnippetByID(context.Background(), createdSnippetID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete snippet during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up snippet: %s", createdSnippetID)
}

// ---

// Test_config_setup_SnippetsAPIService_GetByID tests retrieving a snippet by its ID
func Test_config_setup_SnippetsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Create a snippet first to have something to retrieve
	createdSnippetName := "test-snippet-getbyid-" + common.GenerateRandomString(6)
	snippet := config_setup.Snippets{
		Description: common.StringPtr("Test snippet for Get by ID API testing"),
		Name:        createdSnippetName,
	}

	// Create the snippet via API
	req := client.SnippetsAPI.CreateSnippet(context.Background()).Snippets(snippet)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create snippet for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdSnippetID := createRes.Id
	require.NotEmpty(t, createdSnippetID, "Created snippet should have an ID")

	// Test Get by ID operation
	reqGetById := client.SnippetsAPI.GetSnippetByID(context.Background(), createdSnippetID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get snippet by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdSnippetName, getRes.Name, "Snippet name should match")
	assert.Equal(t, createdSnippetID, getRes.Id, "Snippet ID should match")

	t.Logf("Successfully retrieved snippet: %s", getRes.Name)

	// Cleanup: Delete the created snippet
	reqDel := client.SnippetsAPI.DeleteSnippetByID(context.Background(), createdSnippetID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete snippet during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up snippet: %s", createdSnippetID)
}

// ---

// Test_config_setup_SnippetsAPIService_Update tests updating an existing snippet
func Test_config_setup_SnippetsAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// 1. Create an address first to have something to update
	createdSnippetName := "test-snippet-update-" + common.GenerateRandomString(10)
	initialSnippet := config_setup.Snippets{
		Description: common.StringPtr("Initial description"),
		Name:        createdSnippetName,
		Type:        common.StringPtr("custom"),
	}

	// Create the snippet via API
	reqCreate := client.SnippetsAPI.CreateSnippet(context.Background()).Snippets(initialSnippet)
	createRes, _, err := reqCreate.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create snippet for update test")
	createdSnippetID := createRes.Id

	// 2. Test Update operation with modified fields
	updatedSnippet := config_setup.Snippets{
		Description: common.StringPtr("Updated test snippet description"), // Updated field
		Name:        createdSnippetName,                                   // Name must be the same
		Type:        common.StringPtr("predefined"),                       // update type
	}

	reqUpdate := client.SnippetsAPI.UpdateSnippetByID(context.Background(), createdSnippetID).Snippets(updatedSnippet)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update snippet")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdSnippetName, updateRes.Name, "Snippet name should remain the same")
	assert.Equal(t, *updatedSnippet.Description, *updateRes.Description, "Description should be updated")
	assert.Equal(t, createdSnippetID, updateRes.Id, "Snippet ID should remain the same")

	t.Logf("Successfully updated snippet: %s", createdSnippetName)

	// 3. Cleanup: Delete the created snippet
	reqDel := client.SnippetsAPI.DeleteSnippetByID(context.Background(), createdSnippetID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete snippet during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up snippet: %s", createdSnippetID)
}

// ---

// Test_config_setup_SnippetsAPIService_List tests listing snippets with name filter
func Test_config_setup_SnippetsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// 1. Create a snippet first to have something to list
	createdSnippetName := "test-snippet-list-" + common.GenerateRandomString(10)
	snippet := config_setup.Snippets{
		Description: common.StringPtr("Test snippet for list API testing"),
		Name:        createdSnippetName,
	}

	// Create the snippet via API
	reqCreate := client.SnippetsAPI.CreateSnippet(context.Background()).Snippets(snippet)
	createRes, _, err := reqCreate.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create snippet for list test")
	createdSnippetID := createRes.Id
	require.NotEmpty(t, createdSnippetID, "Created snippet should have an ID")

	// 2. Test List operation with name filter
	// The List API supports filtering by 'name'
	reqList := client.SnippetsAPI.ListSnippets(context.Background()).Limit(500)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list snippets")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Equal(t, 500, len(listRes.Data), "Expected exactly 1 snippet matching the unique name filter")

	t.Logf("Successfully listed and found created snippet: %s", createdSnippetName)

	// 3. Cleanup: Delete the created snippet
	reqDel := client.SnippetsAPI.DeleteSnippetByID(context.Background(), createdSnippetID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete snippet during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up snippet: %s", createdSnippetID)
}

// ---

// Test_config_setup_SnippetsAPIService_DeleteByID tests deleting a snippet by its ID
func Test_config_setup_SnippetsAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// 1. Create a snippet first to have something to delete
	createdSnippetName := "test-snippet-delete-" + common.GenerateRandomString(10)
	snippet := config_setup.Snippets{
		Description: common.StringPtr("Test snippet for delete API testing"),
		Name:        createdSnippetName,
	}

	// Create the snippet via API
	reqCreate := client.SnippetsAPI.CreateSnippet(context.Background()).Snippets(snippet)
	createRes, _, err := reqCreate.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create snippet for delete test")
	createdSnippetID := createRes.Id
	require.NotEmpty(t, createdSnippetID, "Created snippet should have an ID")
	t.Logf("Address created successfully: %s", createdSnippetID)

	// 2. Test Delete by ID operation
	reqDel := client.SnippetsAPI.DeleteSnippetByID(context.Background(), createdSnippetID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete snippet")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted snippet: %s", createdSnippetID)

	// 3. Optional verification: Try to get the deleted snippet (should return 404)
	reqGet := client.SnippetsAPI.GetSnippetByID(context.Background(), createdSnippetID)
	_, httpResGet, errGet := reqGet.Execute()

	assert.Error(t, errGet, "Expected an error when trying to retrieve deleted snippet")
	assert.Equal(t, 404, httpResGet.StatusCode, "Expected 404 Not Found status after deletion")
}
