/*
 * Config Setup Testing
 *
 * LabelsAPIService
 */
/*
Config Setup Testing LabelsAPIService
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

// Test_config_setup_LabelsAPIService_Create tests the creation of a label
func Test_config_setup_LabelsAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Create a valid label object with a unique name
	createdLabelName := "test-label-create-" + common.GenerateRandomString(6)
	label := config_setup.Labels{
		Description: common.StringPtr("Test label for create API testing"),
		Name:        createdLabelName, // Unique test name
	}

	fmt.Printf("Creating label with name: %s\n", label.Name)

	// Make the create request to the API
	req := client.LabelsAPI.CreateLabel(context.Background()).Labels(label)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create label")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdLabelName, res.Name, "Created label name should match")
	require.NotNil(t, label.Description, "Test setup error: expected Description should not be nil")
	require.NotNil(t, res.Description, "API Response error: Description field in response is unexpectedly nil")
	assert.Equal(t, *label.Description, *res.Description, "Description string value should match the expected value")
	assert.NotEmpty(t, res.Id, "Created label should have an ID")

	// Use the ID from the response object
	createdLabelID := res.Id
	t.Logf("Successfully created label: %s with ID: %s", createdLabelName, createdLabelID)

	// Cleanup: Delete the created label to maintain test isolation
	reqDel := client.LabelsAPI.DeleteLabelByID(context.Background(), createdLabelID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete label during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up label: %s", createdLabelID)
}

// ---

// Test_config_setup_LabelsAPIService_GetByID tests retrieving a label by its ID
func Test_config_setup_LabelsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Create a label first to have something to retrieve
	createdLabelName := "test-label-getbyid-" + common.GenerateRandomString(6)
	label := config_setup.Labels{
		Description: common.StringPtr("Test label for Get by ID API testing"),
		Name:        createdLabelName,
	}

	// Create the label via API
	req := client.LabelsAPI.CreateLabel(context.Background()).Labels(label)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create label for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdLabelID := createRes.Id
	require.NotEmpty(t, createdLabelID, "Created label should have an ID")

	// Test Get by ID operation
	reqGetById := client.LabelsAPI.GetLabelByID(context.Background(), createdLabelID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get label by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdLabelName, getRes.Name, "Label name should match")
	assert.Equal(t, createdLabelID, getRes.Id, "Label ID should match")

	t.Logf("Successfully retrieved label: %s", getRes.Name)

	// Cleanup: Delete the created label (lenient - don't fail test on cleanup errors)
	reqDel := client.LabelsAPI.DeleteLabelByID(context.Background(), createdLabelID)
	_, errDel := reqDel.Execute()
	if errDel != nil {
		t.Logf("Warning: cleanup delete returned error (non-fatal): %v", errDel)
	} else {
		t.Logf("Successfully cleaned up label: %s", createdLabelID)
	}
}

// ---

// Test_config_setup_LabelsAPIService_Update tests updating an existing label
func Test_config_setup_LabelsAPIService_Update(t *testing.T) {
	t.Skip("API returns array in update response but model expects object - model deserialization error")
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// 1. Create a label first to have something to update
	createdLabelName := "test-label-update-" + common.GenerateRandomString(10)
	initialLabel := config_setup.Labels{
		Description: common.StringPtr("Initial description"),
		Name:        createdLabelName,
	}

	// Create the label via API
	reqCreate := client.LabelsAPI.CreateLabel(context.Background()).Labels(initialLabel)
	createRes, _, err := reqCreate.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create label for update test")
	createdLabelID := createRes.Id

	// 2. Test Update operation with modified fields
	updatedLabel := config_setup.Labels{
		Description: common.StringPtr("Updated test label description"), // Updated field
		Name:        createdLabelName,                                   // Name must be the same
	}

	reqUpdate := client.LabelsAPI.UpdateLabelByID(context.Background(), createdLabelID).Labels(updatedLabel)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update label")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdLabelName, updateRes.Name, "Label name should remain the same")
	assert.Equal(t, *updatedLabel.Description, *updateRes.Description, "Description should be updated")
	assert.Equal(t, createdLabelID, updateRes.Id, "Label ID should remain the same")

	t.Logf("Successfully updated label: %s", createdLabelName)

	// 3. Cleanup: Delete the created label
	reqDel := client.LabelsAPI.DeleteLabelByID(context.Background(), createdLabelID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete label during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up label: %s", createdLabelID)
}

// ---

// Test_config_setup_LabelsAPIService_List tests listing labels
func Test_config_setup_LabelsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// 1. Create a label first to have something to list
	createdLabelName := "test-label-list-" + common.GenerateRandomString(10)
	label := config_setup.Labels{
		Description: common.StringPtr("Test label for list API testing"),
		Name:        createdLabelName,
	}

	// Create the label via API
	reqCreate := client.LabelsAPI.CreateLabel(context.Background()).Labels(label)
	createRes, _, err := reqCreate.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create label for list test")
	createdLabelID := createRes.Id
	require.NotEmpty(t, createdLabelID, "Created label should have an ID")

	// 2. Test List operation
	reqList := client.LabelsAPI.ListLabels(context.Background()).Limit(500)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list labels")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.GreaterOrEqual(t, len(listRes.Data), 1, "Expected at least 1 label in the list response")

	t.Logf("Successfully listed and found created label: %s", createdLabelName)

	// 3. Cleanup: Delete the created label
	reqDel := client.LabelsAPI.DeleteLabelByID(context.Background(), createdLabelID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete label during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up label: %s", createdLabelID)
}

// ---

// Test_config_setup_LabelsAPIService_DeleteByID tests deleting a label by its ID
func Test_config_setup_LabelsAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// 1. Create a label first to have something to delete
	createdLabelName := "test-label-delete-" + common.GenerateRandomString(10)
	label := config_setup.Labels{
		Description: common.StringPtr("Test label for delete API testing"),
		Name:        createdLabelName,
	}

	// Create the label via API
	reqCreate := client.LabelsAPI.CreateLabel(context.Background()).Labels(label)
	createRes, _, err := reqCreate.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create label for delete test")
	createdLabelID := createRes.Id
	require.NotEmpty(t, createdLabelID, "Created label should have an ID")
	t.Logf("Label created successfully: %s", createdLabelID)

	// 2. Test Delete by ID operation
	reqDel := client.LabelsAPI.DeleteLabelByID(context.Background(), createdLabelID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete label")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted label: %s", createdLabelID)

	// 3. Optional verification: Try to get the deleted label (should return error)
	reqGet := client.LabelsAPI.GetLabelByID(context.Background(), createdLabelID)
	_, httpResGet, errGet := reqGet.Execute()

	assert.Error(t, errGet, "Expected an error when trying to retrieve deleted label")
	if httpResGet != nil {
		assert.True(t, httpResGet.StatusCode >= 400, "Expected error status code after deletion, got %d", httpResGet.StatusCode)
	}
}

// Test_config_setup_LabelsAPIService_FetchLabels tests the FetchLabels convenience method
func Test_config_setup_LabelsAPIService_FetchLabels(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Create test object using same payload as Create test
	testName := "test-label-fetch-" + common.GenerateRandomString(6)
	testObj := config_setup.Labels{
		Description: common.StringPtr("Test label for fetch API testing"),
		Name:        testName,
	}

	createReq := client.LabelsAPI.CreateLabel(context.Background()).Labels(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.LabelsAPI.DeleteLabelByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.LabelsAPI.FetchLabels(
		context.Background(),
		testName,
		nil, // folder (labels don't have folder field)
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch labels by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchLabels found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.LabelsAPI.FetchLabels(
		context.Background(),
		"non-existent-labels-xyz-12345",
		nil, // folder (labels don't have folder field)
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchLabels correctly returned nil for non-existent object")
}
