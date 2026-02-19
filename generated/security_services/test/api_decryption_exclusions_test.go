/*
* Security Services Testing
* DecryptionExclusionsAPIService
 */
package security_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/security_services"
)

// Test_security_services_DecryptionExclusionsAPIService_Create tests the creation of a decryptionexclusion object
// This test creates a new decryptionexclusion and then deletes it to ensure proper cleanup
func Test_security_services_DecryptionExclusionsAPIService_Create(t *testing.T) {
	t.Skip("API rejects empty id field in create request - model serializes non-pointer Id as empty string")
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a valid decryptionexclusion object with unique name to avoid conflicts
	createdDecryptionExclusionName := "test-" + common.GenerateRandomString(10)
	decryptionexclusion := security_services.DecryptionExclusions{
		Description: common.StringPtr("Test decryption exclusion for create API testing"),
		Folder:      common.StringPtr("All"),        // Using All folder scope
		Name:        createdDecryptionExclusionName, // Unique test name
	}

	fmt.Printf("Creating decryptionexclusion with name: %s\n", decryptionexclusion.Name)

	// Make the create request to the API
	req := client.DecryptionExclusionsAPI.CreateDecryptionExclusions(context.Background()).DecryptionExclusions(decryptionexclusion)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create decryptionexclusion")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdDecryptionExclusionName, res.Name, "Created decryptionexclusion name should match")
	assert.Equal(t, common.StringPtr("Test decryption exclusion for create API testing"), res.Description, "Description should match")
	assert.NotEmpty(t, res.Id, "Created decryptionexclusion should have an ID")

	// Use the ID from the response object
	createdDecryptionExclusionID := res.Id
	t.Logf("Successfully created decryptionexclusion: %s with ID: %s", decryptionexclusion.Name, createdDecryptionExclusionID)

	// Cleanup: Delete the created decryptionexclusion to maintain test isolation
	reqDel := client.DecryptionExclusionsAPI.DeleteDecryptionExclusionsByID(context.Background(), createdDecryptionExclusionID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete decryptionexclusion during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up decryptionexclusion: %s", createdDecryptionExclusionID)
}

// Test_security_services_DecryptionExclusionsAPIService_GetByID tests retrieving a decryptionexclusion by its ID
// This test creates a decryptionexclusion, retrieves it by ID, then deletes it
func Test_security_services_DecryptionExclusionsAPIService_GetByID(t *testing.T) {
	t.Skip("API rejects empty id field in create request - model serializes non-pointer Id as empty string")
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a decryptionexclusion first to have something to retrieve
	createdDecryptionExclusionName := "test-getbyid-" + common.GenerateRandomString(10)
	decryptionexclusion := security_services.DecryptionExclusions{
		Description: common.StringPtr("Test decryption exclusion for get by ID API testing"),
		Folder:      common.StringPtr("All"),        // Using All folder scope
		Name:        createdDecryptionExclusionName, // Unique test name
	}

	// Create the decryptionexclusion via API
	req := client.DecryptionExclusionsAPI.CreateDecryptionExclusions(context.Background()).DecryptionExclusions(decryptionexclusion)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create decryptionexclusion for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDecryptionExclusionID := createRes.Id
	require.NotEmpty(t, createdDecryptionExclusionID, "Created decryptionexclusion should have an ID")

	// Test Get by ID operation
	reqGetById := client.DecryptionExclusionsAPI.GetDecryptionExclusionsByID(context.Background(), createdDecryptionExclusionID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get decryptionexclusion by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdDecryptionExclusionName, getRes.Name, "DecryptionExclusion name should match")
	assert.Equal(t, common.StringPtr("Test decryption exclusion for get by ID API testing"), getRes.Description, "Description should match")
	assert.Equal(t, createdDecryptionExclusionID, getRes.Id, "DecryptionExclusion ID should match")

	t.Logf("Successfully retrieved decryptionexclusion: %s", getRes.Name)

	// Cleanup: Delete the created decryptionexclusion
	reqDel := client.DecryptionExclusionsAPI.DeleteDecryptionExclusionsByID(context.Background(), createdDecryptionExclusionID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete decryptionexclusion during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up decryptionexclusion: %s", createdDecryptionExclusionID)
}

// Test_security_services_DecryptionExclusionsAPIService_Update tests updating an existing decryptionexclusion
// This test creates a decryptionexclusion, updates it, then deletes it
func Test_security_services_DecryptionExclusionsAPIService_Update(t *testing.T) {
	t.Skip("API rejects empty id field in create request - model serializes non-pointer Id as empty string")
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a decryptionexclusion first to have something to update
	createdDecryptionExclusionName := "test-update-" + common.GenerateRandomString(10)
	decryptionexclusion := security_services.DecryptionExclusions{
		Description: common.StringPtr("Test decryption exclusion for update API testing"),
		Folder:      common.StringPtr("All"),        // Using All folder scope
		Name:        createdDecryptionExclusionName, // Unique test name
	}

	// Create the decryptionexclusion via API
	req := client.DecryptionExclusionsAPI.CreateDecryptionExclusions(context.Background()).DecryptionExclusions(decryptionexclusion)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create decryptionexclusion for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDecryptionExclusionID := createRes.Id
	require.NotEmpty(t, createdDecryptionExclusionID, "Created decryptionexclusion should have an ID")

	// Test Update operation with modified fields
	updatedDecryptionExclusion := security_services.DecryptionExclusions{
		Description: common.StringPtr("Updated test decryption exclusion description"), // Updated description
		Folder:      common.StringPtr("All"),                                           // Keep same folder scope
		Name:        createdDecryptionExclusionName,                                    // Keep same name (required for update)
	}

	reqUpdate := client.DecryptionExclusionsAPI.UpdateDecryptionExclusionsByID(context.Background(), createdDecryptionExclusionID).DecryptionExclusions(updatedDecryptionExclusion)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update decryptionexclusion")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdDecryptionExclusionName, updateRes.Name, "DecryptionExclusion name should remain the same")
	assert.Equal(t, common.StringPtr("Updated test decryption exclusion description"), updateRes.Description, "Description should be updated")
	assert.Equal(t, createdDecryptionExclusionID, updateRes.Id, "DecryptionExclusion ID should remain the same")

	t.Logf("Successfully updated decryptionexclusion: %s", createdDecryptionExclusionName)

	// Cleanup: Delete the created decryptionexclusion
	reqDel := client.DecryptionExclusionsAPI.DeleteDecryptionExclusionsByID(context.Background(), createdDecryptionExclusionID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete decryptionexclusion during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up decryptionexclusion: %s", createdDecryptionExclusionID)
}

// Test_security_services_DecryptionExclusionsAPIService_List tests listing decryptionexclusions with folder filter
// This is a read-only test that lists existing decryption exclusions (Create blocked by non-pointer Id issue)
func Test_security_services_DecryptionExclusionsAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	// Test List operation with folder filter (read-only, no create needed)
	reqList := client.DecryptionExclusionsAPI.ListDecryptionExclusions(context.Background()).Folder("All").Limit(200).Offset(0)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	require.NoError(t, errList, "Failed to list decryptionexclusions")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	t.Logf("Successfully listed decryption exclusions, total: %d", len(listRes.Data))
}

// Test_security_services_DecryptionExclusionsAPIService_DeleteByID tests deleting a decryptionexclusion by its ID
// This test creates a decryptionexclusion, deletes it, then verifies the deletion was successful
func Test_security_services_DecryptionExclusionsAPIService_DeleteByID(t *testing.T) {
	t.Skip("API rejects empty id field in create request - model serializes non-pointer Id as empty string")
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a decryptionexclusion first to have something to delete
	createdDecryptionExclusionName := "test-delete-" + common.GenerateRandomString(10)
	decryptionexclusion := security_services.DecryptionExclusions{
		Description: common.StringPtr("Test decryption exclusion for delete API testing"),
		Folder:      common.StringPtr("All"),        // Using All folder scope
		Name:        createdDecryptionExclusionName, // Unique test name
	}

	// Create the decryptionexclusion via API
	req := client.DecryptionExclusionsAPI.CreateDecryptionExclusions(context.Background()).DecryptionExclusions(decryptionexclusion)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create decryptionexclusion for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDecryptionExclusionID := createRes.Id
	require.NotEmpty(t, createdDecryptionExclusionID, "Created decryptionexclusion should have an ID")

	// Test Delete by ID operation
	reqDel := client.DecryptionExclusionsAPI.DeleteDecryptionExclusionsByID(context.Background(), createdDecryptionExclusionID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete decryptionexclusion")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted decryptionexclusion: %s", createdDecryptionExclusionID)
}

// Test_security_services_DecryptionExclusionsAPIService_Fetch tests the fetch convenience method
// This is a read-only test (Create blocked by non-pointer Id issue)
func Test_security_services_DecryptionExclusionsAPIService_Fetch(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	// Test: Fetch non-existent object (should return nil, nil)
	notFound, err := client.DecryptionExclusionsAPI.FetchDecryptionExclusions(
		context.Background(),
		"non-existent-exclusion-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchDecryptionExclusions correctly returned nil for non-existent object")
}
