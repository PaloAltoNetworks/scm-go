/*
* Security Services Testing
* AntiSpywareProfilesAPIService
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

// Test_security_services_AntiSpywareProfilesAPIService_Create tests the creation of an antispywareprofile object
// This test creates a new antispywareprofile and then deletes it to ensure proper cleanup
func Test_security_services_AntiSpywareProfilesAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a valid antispywareprofile object with unique name to avoid conflicts
	createdAntiSpywareProfileName := "test-" + common.GenerateRandomString(10)
	antispywareprofile := security_services.AntiSpywareProfiles{
		Description:         common.StringPtr("Test antispywareprofile for create API testing"),
		Folder:              common.StringPtr("Prisma Access"), // Using Prisma Access folder scope
		CloudInlineAnalysis: common.BoolPtr(true),              // Enable Inline Cloud Analysis
		Name:                createdAntiSpywareProfileName,     // Unique test name
	}

	fmt.Printf("Creating antispywareprofile with name: %s\n", antispywareprofile.Name)

	// Make the create request to the API
	req := client.AntiSpywareProfilesAPI.CreateAntiSpywareProfiles(context.Background()).AntiSpywareProfiles(antispywareprofile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create antispywareprofile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdAntiSpywareProfileName, res.Name, "Created antispywareprofile name should match")
	assert.Equal(t, common.StringPtr("Test antispywareprofile for create API testing"), res.Description, "Description should match")
	assert.True(t, *res.Folder == "Shared" || *res.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.True(t, *res.CloudInlineAnalysis, "CloudInlineAnalysis status should match")
	assert.NotEmpty(t, res.Id, "Created antispywareprofile should have an ID")

	// Use the ID from the response object
	createdAntiSpywareProfileID := res.Id
	t.Logf("Successfully created antispywareprofile: %s with ID: %s", antispywareprofile.Name, createdAntiSpywareProfileID)

	// Cleanup: Delete the created antispywareprofile to maintain test isolation
	reqDel := client.AntiSpywareProfilesAPI.DeleteAntiSpywareProfilesByID(context.Background(), createdAntiSpywareProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete antispywareprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up antispywareprofile: %s", createdAntiSpywareProfileID)
}

// Test_security_services_AntiSpywareProfilesAPIService_GetByID tests retrieving an antispywareprofile by its ID
// This test creates an antispywareprofile, retrieves it by ID, then deletes it
func Test_security_services_AntiSpywareProfilesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create an antispywareprofile first to have something to retrieve
	createdAntiSpywareProfileName := "test-getbyid-" + common.GenerateRandomString(10)
	antispywareprofile := security_services.AntiSpywareProfiles{
		Description:         common.StringPtr("Test antispywareprofile for get by ID API testing"),
		Folder:              common.StringPtr("Prisma Access"), // Using Prisma Access folder scope
		CloudInlineAnalysis: common.BoolPtr(true),              // Enable Inline Cloud Analysis
		Name:                createdAntiSpywareProfileName,     // Unique test name
	}

	// Create the antispywareprofile via API
	req := client.AntiSpywareProfilesAPI.CreateAntiSpywareProfiles(context.Background()).AntiSpywareProfiles(antispywareprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create antispywareprofile for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAntiSpywareProfileID := createRes.Id
	require.NotEmpty(t, createdAntiSpywareProfileID, "Created antispywareprofile should have an ID")

	// Test Get by ID operation
	reqGetById := client.AntiSpywareProfilesAPI.GetAntiSpywareProfilesByID(context.Background(), createdAntiSpywareProfileID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get antispywareprofile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdAntiSpywareProfileName, getRes.Name, "AntiSpywareProfile name should match")
	assert.Equal(t, common.StringPtr("Test antispywareprofile for get by ID API testing"), getRes.Description, "Description should match")
	assert.True(t, *getRes.Folder == "Shared" || *getRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.True(t, *getRes.CloudInlineAnalysis, "Inline Cloud Analysis should be enabled")
	assert.Equal(t, createdAntiSpywareProfileID, getRes.Id, "AntiSpywareProfile ID should match")

	t.Logf("Successfully retrieved antispywareprofile: %s", getRes.Name)

	// Cleanup: Delete the created antispywareprofile
	reqDel := client.AntiSpywareProfilesAPI.DeleteAntiSpywareProfilesByID(context.Background(), createdAntiSpywareProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete antispywareprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up antispywareprofile: %s", createdAntiSpywareProfileID)
}

// Test_security_services_AntiSpywareProfilesAPIService_Update tests updating an existing antispywareprofile
// This test creates an antispywareprofile, updates it, then deletes it
func Test_security_services_AntiSpywareProfilesAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create an antispywareprofile first to have something to update
	createdAntiSpywareProfileName := "test-update-" + common.GenerateRandomString(10)
	antispywareprofile := security_services.AntiSpywareProfiles{
		Description:         common.StringPtr("Test antispywareprofile for update API testing"),
		Folder:              common.StringPtr("Prisma Access"), // Using Prisma Access folder scope
		CloudInlineAnalysis: common.BoolPtr(true),              // Enable Inline Cloud Analysis
		Name:                createdAntiSpywareProfileName,     // Unique test name
	}

	// Create the antispywareprofile via API
	req := client.AntiSpywareProfilesAPI.CreateAntiSpywareProfiles(context.Background()).AntiSpywareProfiles(antispywareprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create antispywareprofile for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAntiSpywareProfileID := createRes.Id
	require.NotEmpty(t, createdAntiSpywareProfileID, "Created antispywareprofile should have an ID")

	// Test Update operation with modified fields
	updatedAntiSpywareProfile := security_services.AntiSpywareProfiles{
		Description:         common.StringPtr("Updated test antispywareprofile description"), // Updated description
		Folder:              common.StringPtr("Prisma Access"),                               // Keep same folder scope
		CloudInlineAnalysis: common.BoolPtr(false),                                           // Updated Inline Cloud Analysis
		Name:                createdAntiSpywareProfileName,                                   // Keep same name (required for update)
	}

	reqUpdate := client.AntiSpywareProfilesAPI.UpdateAntiSpywareProfilesByID(context.Background(), createdAntiSpywareProfileID).AntiSpywareProfiles(updatedAntiSpywareProfile)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update antispywareprofile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdAntiSpywareProfileName, updateRes.Name, "AntiSpywareProfile name should remain the same")
	assert.Equal(t, common.StringPtr("Updated test antispywareprofile description"), updateRes.Description, "Description should be updated")
	assert.True(t, *updateRes.Folder == "Shared" || *updateRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.False(t, *updateRes.CloudInlineAnalysis, "CloudInlineAnalysis should be updated")
	assert.Equal(t, createdAntiSpywareProfileID, updateRes.Id, "AntiSpywareProfile ID should remain the same")

	t.Logf("Successfully updated antispywareprofile: %s", createdAntiSpywareProfileName)

	// Cleanup: Delete the created antispywareprofile
	reqDel := client.AntiSpywareProfilesAPI.DeleteAntiSpywareProfilesByID(context.Background(), createdAntiSpywareProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete antispywareprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up antispywareprofile: %s", createdAntiSpywareProfileID)
}

// Test_security_services_AntiSpywareProfilesAPIService_List tests listing antispywareprofiles with folder filter
// This test creates an antispywareprofile, lists antispywareprofiles to verify it's included, then deletes it
func Test_security_services_AntiSpywareProfilesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create an antispywareprofile first to have something to list
	createdAntiSpywareProfileName := "test-list-" + common.GenerateRandomString(10)
	antispywareprofile := security_services.AntiSpywareProfiles{
		Description:         common.StringPtr("Test antispywareprofile for list API testing"),
		Folder:              common.StringPtr("Prisma Access"), // Using Prisma Access folder scope
		CloudInlineAnalysis: common.BoolPtr(true),              // Enable Inline Cloud Analysis
		Name:                createdAntiSpywareProfileName,     // Unique test name
	}

	// Create the antispywareprofile via API
	req := client.AntiSpywareProfilesAPI.CreateAntiSpywareProfiles(context.Background()).AntiSpywareProfiles(antispywareprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create antispywareprofile for list test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAntiSpywareProfileID := createRes.Id
	require.NotEmpty(t, createdAntiSpywareProfileID, "Created antispywareprofile should have an ID")

	// Test List operation with folder filter
	reqList := client.AntiSpywareProfilesAPI.ListAntiSpywareProfiles(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list antispywareprofiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one antispywareprofile in the list")

	// Verify our created antispywareprofile is in the list
	foundAntiSpywareProfile := false
	for _, profile := range listRes.Data {
		if profile.Name == createdAntiSpywareProfileName {
			foundAntiSpywareProfile = true
			assert.Equal(t, common.StringPtr("Test antispywareprofile for list API testing"), profile.Description, "Description should match")
			assert.True(t, *profile.Folder == "Shared" || *profile.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
			assert.True(t, *profile.CloudInlineAnalysis, "CloudInlineAnalysis should match")
			break
		}
	}
	assert.True(t, foundAntiSpywareProfile, "Created antispywareprofile should be found in the list")

	t.Logf("Successfully listed antispywareprofiles, found created antispywareprofile: %s", createdAntiSpywareProfileName)

	// Cleanup: Delete the created antispywareprofile
	reqDel := client.AntiSpywareProfilesAPI.DeleteAntiSpywareProfilesByID(context.Background(), createdAntiSpywareProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete antispywareprofile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up antispywareprofile: %s", createdAntiSpywareProfileID)
}

// Test_security_services_AntiSpywareProfilesAPIService_DeleteByID tests deleting an antispywareprofile by its ID
// This test creates an antispywareprofile, deletes it, then verifies the deletion was successful
func Test_security_services_AntiSpywareProfilesAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create an antispywareprofile first to have something to delete
	createdAntiSpywareProfileName := "test-delete-" + common.GenerateRandomString(10)
	antispywareprofile := security_services.AntiSpywareProfiles{
		Description:         common.StringPtr("Test antispywareprofile for delete API testing"),
		Folder:              common.StringPtr("Prisma Access"), // Using Prisma Access folder scope
		CloudInlineAnalysis: common.BoolPtr(true),              // Enable Inline Cloud Analysis
		Name:                createdAntiSpywareProfileName,     // Unique test name
	}

	// Create the antispywareprofile via API
	req := client.AntiSpywareProfilesAPI.CreateAntiSpywareProfiles(context.Background()).AntiSpywareProfiles(antispywareprofile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create antispywareprofile for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAntiSpywareProfileID := createRes.Id
	require.NotEmpty(t, createdAntiSpywareProfileID, "Created antispywareprofile should have an ID")

	// Test Delete by ID operation
	reqDel := client.AntiSpywareProfilesAPI.DeleteAntiSpywareProfilesByID(context.Background(), createdAntiSpywareProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete antispywareprofile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted antispywareprofile: %s", createdAntiSpywareProfileID)
}
