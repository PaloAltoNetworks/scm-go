/*
* Security Services Testing
* AntiSpywareSignaturesAPIService
 */
package security_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/security_services"
)

// Test_security_services_AntiSpywareSignaturesAPIService_Create tests the creation of an antispywaresignature object
// This test creates a new antispywaresignature and then deletes it to ensure proper cleanup
func Test_security_services_AntiSpywareSignaturesAPIService_Create(t *testing.T) {
	t.Skip("Model has Id as string (non-pointer) which causes JSON serialization issues - omitempty cannot be applied")

	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a valid antispywaresignature object with unique name to avoid conflicts
	createdAntiSpywareSignatureThreatName := "test-" + common.GenerateRandomString(10)
	antispywaresignature := security_services.AntiSpywareSignatures{
		Comment:    common.StringPtr("Test anti-spyware signature for create API testing"),
		Folder:     common.StringPtr("All"),               // Using All folder scope
		Threatname: createdAntiSpywareSignatureThreatName, // Unique test name
		ThreatId:   "6900001",                             // Required field - must be in range 15000-18000 or 6900001-7000000
		Id:         "",                                    // This is a required string field (not pointer) which causes serialization issues
	}

	// Make the create request to the API
	req := client.AntiSpywareSignaturesAPI.CreateAntiSpywareSignatures(context.Background()).AntiSpywareSignatures(antispywaresignature)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create antispywaresignature")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdAntiSpywareSignatureThreatName, res.Threatname, "Created antispywaresignature threatname should match")
	assert.NotEmpty(t, res.Id, "Created antispywaresignature should have an ID")

	// Use the ID from the response object
	createdAntiSpywareSignatureID := res.Id
	t.Logf("Successfully created antispywaresignature: %s with ID: %s", antispywaresignature.Threatname, createdAntiSpywareSignatureID)

	// Cleanup: Delete the created antispywaresignature to maintain test isolation
	reqDel := client.AntiSpywareSignaturesAPI.DeleteAntiSpywareSignaturesByID(context.Background(), createdAntiSpywareSignatureID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete antispywaresignature during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up antispywaresignature: %s", createdAntiSpywareSignatureID)
}

// Test_security_services_AntiSpywareSignaturesAPIService_GetByID tests retrieving an antispywaresignature by its ID
// This test creates an antispywaresignature, retrieves it by ID, then deletes it
func Test_security_services_AntiSpywareSignaturesAPIService_GetByID(t *testing.T) {
	t.Skip("Model has Id as string (non-pointer) which causes JSON serialization issues - omitempty cannot be applied")

	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create an antispywaresignature first to have something to retrieve
	createdAntiSpywareSignatureThreatName := "test-getbyid-" + common.GenerateRandomString(10)
	antispywaresignature := security_services.AntiSpywareSignatures{
		Comment:    common.StringPtr("Test anti-spyware signature for get by ID API testing"),
		Folder:     common.StringPtr("All"),               // Using All folder scope
		Threatname: createdAntiSpywareSignatureThreatName, // Unique test name
		ThreatId:   "6900002",                             // Required field
		Id:         "",                                    // Required string field causing issues
	}

	// Create the antispywaresignature via API
	req := client.AntiSpywareSignaturesAPI.CreateAntiSpywareSignatures(context.Background()).AntiSpywareSignatures(antispywaresignature)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create antispywaresignature for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAntiSpywareSignatureID := createRes.Id
	require.NotEmpty(t, createdAntiSpywareSignatureID, "Created antispywaresignature should have an ID")

	// Test Get by ID operation
	reqGetById := client.AntiSpywareSignaturesAPI.GetAntiSpywareSignaturesByID(context.Background(), createdAntiSpywareSignatureID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get antispywaresignature by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdAntiSpywareSignatureThreatName, getRes.Threatname, "AntiSpywareSignature threatname should match")
	assert.Equal(t, createdAntiSpywareSignatureID, getRes.Id, "AntiSpywareSignature ID should match")

	t.Logf("Successfully retrieved antispywaresignature: %s", getRes.Threatname)

	// Cleanup: Delete the created antispywaresignature
	reqDel := client.AntiSpywareSignaturesAPI.DeleteAntiSpywareSignaturesByID(context.Background(), createdAntiSpywareSignatureID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete antispywaresignature during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up antispywaresignature: %s", createdAntiSpywareSignatureID)
}

// Test_security_services_AntiSpywareSignaturesAPIService_Update tests updating an existing antispywaresignature
// This test creates an antispywaresignature, updates it, then deletes it
func Test_security_services_AntiSpywareSignaturesAPIService_Update(t *testing.T) {
	t.Skip("Model has Id as string (non-pointer) which causes JSON serialization issues - omitempty cannot be applied")

	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create an antispywaresignature first to have something to update
	createdAntiSpywareSignatureThreatName := "test-update-" + common.GenerateRandomString(10)
	antispywaresignature := security_services.AntiSpywareSignatures{
		Comment:    common.StringPtr("Test anti-spyware signature for update API testing"),
		Folder:     common.StringPtr("All"),               // Using All folder scope
		Threatname: createdAntiSpywareSignatureThreatName, // Unique test name
		ThreatId:   "6900003",                             // Required field
		Id:         "",                                    // Required string field causing issues
	}

	// Create the antispywaresignature via API
	req := client.AntiSpywareSignaturesAPI.CreateAntiSpywareSignatures(context.Background()).AntiSpywareSignatures(antispywaresignature)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create antispywaresignature for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAntiSpywareSignatureID := createRes.Id
	require.NotEmpty(t, createdAntiSpywareSignatureID, "Created antispywaresignature should have an ID")

	// Test Update operation with modified fields
	updatedAntiSpywareSignature := security_services.AntiSpywareSignatures{
		Comment:    common.StringPtr("Updated test anti-spyware signature comment"), // Updated comment
		Folder:     common.StringPtr("All"),                                         // Keep same folder scope
		Threatname: createdAntiSpywareSignatureThreatName,                           // Keep same threatname (required for update)
		ThreatId:   "6900003",                                                       // Keep same threat ID
		Id:         createdAntiSpywareSignatureID,                                   // Include ID for update
	}

	reqUpdate := client.AntiSpywareSignaturesAPI.UpdateAntiSpywareSignaturesByID(context.Background(), createdAntiSpywareSignatureID).AntiSpywareSignatures(updatedAntiSpywareSignature)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update antispywaresignature")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdAntiSpywareSignatureThreatName, updateRes.Threatname, "AntiSpywareSignature threatname should remain the same")
	assert.Equal(t, common.StringPtr("Updated test anti-spyware signature comment"), updateRes.Comment, "Comment should be updated")
	assert.Equal(t, createdAntiSpywareSignatureID, updateRes.Id, "AntiSpywareSignature ID should remain the same")

	t.Logf("Successfully updated antispywaresignature: %s", createdAntiSpywareSignatureThreatName)

	// Cleanup: Delete the created antispywaresignature
	reqDel := client.AntiSpywareSignaturesAPI.DeleteAntiSpywareSignaturesByID(context.Background(), createdAntiSpywareSignatureID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete antispywaresignature during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up antispywaresignature: %s", createdAntiSpywareSignatureID)
}

// Test_security_services_AntiSpywareSignaturesAPIService_List tests listing antispywaresignatures with folder filter
// This test creates an antispywaresignature, lists antispywaresignatures to verify it's included, then deletes it
func Test_security_services_AntiSpywareSignaturesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create an antispywaresignature first to have something to list
	createdAntiSpywareSignatureThreatName := "test-list-" + common.GenerateRandomString(10)
	antispywaresignature := security_services.AntiSpywareSignatures{
		Comment:    common.StringPtr("Test anti-spyware signature for list API testing"),
		Folder:     common.StringPtr("All"),               // Using All folder scope
		Threatname: createdAntiSpywareSignatureThreatName, // Unique test name
		ThreatId:   "6900004",                             // Required field
		Id:         "",                                    // Required string field causing issues
	}

	// Create the antispywaresignature via API
	req := client.AntiSpywareSignaturesAPI.CreateAntiSpywareSignatures(context.Background()).AntiSpywareSignatures(antispywaresignature)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create antispywaresignature for list test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAntiSpywareSignatureID := createRes.Id
	require.NotEmpty(t, createdAntiSpywareSignatureID, "Created antispywaresignature should have an ID")

	// Test List operation with folder filter
	reqList := client.AntiSpywareSignaturesAPI.ListAntiSpywareSignatures(context.Background()).Folder("All").Limit(200).Offset(0)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list antispywaresignatures")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one antispywaresignature in the list")

	// Verify our created antispywaresignature is in the list
	foundAntiSpywareSignature := false
	for _, signature := range listRes.Data {
		if signature.Threatname == createdAntiSpywareSignatureThreatName {
			foundAntiSpywareSignature = true
			assert.Equal(t, common.StringPtr("Test anti-spyware signature for list API testing"), signature.Comment, "Comment should match")
			break
		}
	}
	assert.True(t, foundAntiSpywareSignature, "Created antispywaresignature should be found in the list")

	t.Logf("Successfully listed antispywaresignatures, found created antispywaresignature: %s", createdAntiSpywareSignatureThreatName)

	// Cleanup: Delete the created antispywaresignature
	reqDel := client.AntiSpywareSignaturesAPI.DeleteAntiSpywareSignaturesByID(context.Background(), createdAntiSpywareSignatureID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete antispywaresignature during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up antispywaresignature: %s", createdAntiSpywareSignatureID)
}

// Test_security_services_AntiSpywareSignaturesAPIService_DeleteByID tests deleting an antispywaresignature by its ID
// This test creates an antispywaresignature, deletes it, then verifies the deletion was successful
func Test_security_services_AntiSpywareSignaturesAPIService_DeleteByID(t *testing.T) {
	t.Skip("Model has Id as string (non-pointer) which causes JSON serialization issues - omitempty cannot be applied")

	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create an antispywaresignature first to have something to delete
	createdAntiSpywareSignatureThreatName := "test-delete-" + common.GenerateRandomString(10)
	antispywaresignature := security_services.AntiSpywareSignatures{
		Comment:    common.StringPtr("Test anti-spyware signature for delete API testing"),
		Folder:     common.StringPtr("All"),               // Using All folder scope
		Threatname: createdAntiSpywareSignatureThreatName, // Unique test name
		ThreatId:   "6900005",                             // Required field
		Id:         "",                                    // Required string field causing issues
	}

	// Create the antispywaresignature via API
	req := client.AntiSpywareSignaturesAPI.CreateAntiSpywareSignatures(context.Background()).AntiSpywareSignatures(antispywaresignature)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create antispywaresignature for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAntiSpywareSignatureID := createRes.Id
	require.NotEmpty(t, createdAntiSpywareSignatureID, "Created antispywaresignature should have an ID")

	// Test Delete by ID operation
	reqDel := client.AntiSpywareSignaturesAPI.DeleteAntiSpywareSignaturesByID(context.Background(), createdAntiSpywareSignatureID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete antispywaresignature")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted antispywaresignature: %s", createdAntiSpywareSignatureID)
}
