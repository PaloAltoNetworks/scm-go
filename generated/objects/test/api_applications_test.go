/*
Objects Testing ApplicationsAPIService
*/
package objects

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)

// Test_objects_ApplicationsAPIService_Create tests the creation of an application object.
// This test creates a new application and then deletes it to ensure proper cleanup.
func Test_objects_ApplicationsAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a valid application object with a unique name to avoid conflicts.
	createdAppName := "test-app-create-" + common.GenerateRandomString(6)
	application := objects.Applications{
		Folder:      common.StringPtr("Prisma Access"),
		Name:        createdAppName,
		Description: common.StringPtr("Test application for create API"),
		Category:    "business-systems",
		Subcategory: common.StringPtr("ics-protocols"),
		Technology:  common.StringPtr("client-server"),
		Risk:        3,
		Default: &objects.ApplicationsDefault{
			Port: []string{"tcp/80", "tcp/443"},
		},
	}

	fmt.Printf("Creating application with name: %s\n", application.Name)

	// Make the create request to the API.
	req := client.ApplicationsAPI.CreateApplications(context.Background()).Applications(application)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create application")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties.
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdAppName, res.Name, "Created application name should match")
	assert.True(t, *res.Folder == "Shared" || *res.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, "business-systems", res.Category, "Category should match")
	assert.Equal(t, float64(3), res.Risk, "Risk should match")
	require.NotNil(t, res.Id, "Created application should have an ID")
	assert.NotEmpty(t, *res.Id, "Created application ID should not be empty")

	// Use the ID from the response object.
	createdAppID := *res.Id
	t.Logf("Successfully created application: %s with ID: %s", application.Name, createdAppID)

	// Cleanup: Delete the created application to maintain test isolation.
	reqDel := client.ApplicationsAPI.DeleteApplicationsByID(context.Background(), createdAppID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete application during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up application: %s", createdAppID)
}

// Test_objects_ApplicationsAPIService_GetByID tests retrieving an application by its ID.
// This test creates an application, retrieves it by ID, then deletes it.
func Test_objects_ApplicationsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create an application first to have something to retrieve.
	createdAppName := "test-app-get-" + common.GenerateRandomString(6)
	application := objects.Applications{
		Folder:      common.StringPtr("Prisma Access"),
		Name:        createdAppName,
		Description: common.StringPtr("Test application for get API"),
		Category:    "business-systems",
		Subcategory: common.StringPtr("ics-protocols"),
		Technology:  common.StringPtr("client-server"),
		Risk:        4,
	}

	// Create the application via API.
	req := client.ApplicationsAPI.CreateApplications(context.Background()).Applications(application)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create application for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created application should have an ID")
	createdAppID := *createRes.Id
	require.NotEmpty(t, createdAppID, "Created application ID should not be empty")

	// Test Get by ID operation.
	reqGetById := client.ApplicationsAPI.GetApplicationsByID(context.Background(), createdAppID)
	getRes, httpResGet, errGet := reqGetById.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful.
	require.NoError(t, errGet, "Failed to get application by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties.
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdAppName, getRes.Name, "Application name should match")
	assert.True(t, *getRes.Folder == "Shared" || *getRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, common.StringPtr("ics-protocols"), getRes.Subcategory, "Subcategory should match")
	require.NotNil(t, getRes.Id, "Application ID should not be nil")
	assert.Equal(t, createdAppID, *getRes.Id, "Application ID should match")

	t.Logf("Successfully retrieved application: %s", getRes.Name)

	// Cleanup: Delete the created application.
	reqDel := client.ApplicationsAPI.DeleteApplicationsByID(context.Background(), createdAppID)
	_, errDel := reqDel.Execute()
	require.NoError(t, errDel, "Failed to delete application during cleanup")

	t.Logf("Successfully cleaned up application: %s", createdAppID)
}

// Test_objects_ApplicationsAPIService_Update tests updating an existing application.
// This test creates an application, updates it, then deletes it.
func Test_objects_ApplicationsAPIService_Update(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create an application first to have something to update.
	createdAppName := "test-app-update-" + common.GenerateRandomString(6)
	application := objects.Applications{
		Folder:      common.StringPtr("Prisma Access"),
		Name:        createdAppName,
		Category:    "business-systems",
		Subcategory: common.StringPtr("ics-protocols"),
		Technology:  common.StringPtr("client-server"),
		Risk:        3,
	}

	// Create the application via API.
	req := client.ApplicationsAPI.CreateApplications(context.Background()).Applications(application)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create application for update test")
	require.NotNil(t, createRes.Id, "Created application should have an ID")
	createdAppID := *createRes.Id
	require.NotEmpty(t, createdAppID, "Created application ID should not be empty")

	// Test Update operation with modified fields.
	updatedApplication := objects.Applications{
		Description:           common.StringPtr("Updated description"),
		Category:              "networking",
		Subcategory:           common.StringPtr("encrypted-tunnel"),
		Technology:            common.StringPtr("peer-to-peer"),
		Risk:                  5,
		AbleToTransferFile:    common.BoolPtr(true),
		HasKnownVulnerability: common.BoolPtr(true),
		Name:                  createdAppName,
	}

	reqUpdate := client.ApplicationsAPI.UpdateApplicationsByID(context.Background(), createdAppID).Applications(updatedApplication)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful.
	require.NoError(t, errUpdate, "Failed to update application")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties.
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, common.StringPtr("Updated description"), updateRes.Description, "Description should be updated")
	assert.Equal(t, "networking", updateRes.Category, "Category should be updated")
	assert.Equal(t, float64(5), updateRes.Risk, "Risk should be updated")
	assert.Equal(t, common.BoolPtr(true), updateRes.AbleToTransferFile, "AbleToTransferFile should be updated")
	require.NotNil(t, updateRes.Id, "Application ID should not be nil")
	assert.Equal(t, createdAppID, *updateRes.Id, "Application ID should remain the same")

	t.Logf("Successfully updated application: %s", createdAppName)

	// Cleanup: Delete the created application.
	reqDel := client.ApplicationsAPI.DeleteApplicationsByID(context.Background(), createdAppID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete application during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up application: %s", createdAppID)
}

// Test_objects_ApplicationsAPIService_List tests listing applications.
// This test creates an application, lists applications to verify it's included, then deletes it.
func Test_objects_ApplicationsAPIService_List(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create an application first to have something to list.
	createdAppName := "test-app-list-" + common.GenerateRandomString(6)
	application := objects.Applications{
		Folder:      common.StringPtr("Prisma Access"),
		Name:        createdAppName,
		Description: common.StringPtr("Test application for create API"),
		Category:    "business-systems",
		Subcategory: common.StringPtr("ics-protocols"),
		Technology:  common.StringPtr("client-server"),
		Risk:        3,
		Default: &objects.ApplicationsDefault{
			Port: []string{"tcp/80", "tcp/443"},
		},
	}

	// Create the application via API.
	req := client.ApplicationsAPI.CreateApplications(context.Background()).Applications(application)
	createRes, _, err := req.Execute()
	if err != nil {
		require.NoError(t, err, "Failed to create application for list test")
		return
	}
	require.NoError(t, err, "Failed to create application for list test")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created application should have an ID")
	createdAppID := *createRes.Id
	require.NotEmpty(t, createdAppID, "Created application ID should not be empty")

	// Test List operation.
	reqList := client.ApplicationsAPI.ListApplications(context.Background()).Folder("Prisma Access").Limit(10000)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful.
	require.NoError(t, errList, "Failed to list applications")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one application in the list")

	// Verify our created application is in the list.
	foundApp := false
	for _, app := range listRes.Data {
		if app.Name == createdAppName {
			foundApp = true
			assert.Equal(t, "business-systems", app.Category, "Category should match")
			assert.True(t, *app.Folder == "Shared" || *app.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
			break
		}
	}
	assert.True(t, foundApp, "Created application should be found in the list")

	t.Logf("Successfully listed applications, found created application: %s %t", createdAppName, foundApp)

	// Cleanup: Delete the created application.
	reqDel := client.ApplicationsAPI.DeleteApplicationsByID(context.Background(), createdAppID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete application during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up application: %s", createdAppID)
}

// Test_objects_ApplicationsAPIService_DeleteByID tests deleting an application by its ID.
// This test creates an application, deletes it, then verifies the deletion was successful.
func Test_objects_ApplicationsAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create an application first to have something to delete.
	createdAppName := "test-app-delete-" + common.GenerateRandomString(6)
	application := objects.Applications{
		Folder:      common.StringPtr("Prisma Access"),
		Name:        createdAppName,
		Description: common.StringPtr("Test application for create API"),
		Category:    "business-systems",
		Subcategory: common.StringPtr("ics-protocols"),
		Technology:  common.StringPtr("client-server"),
		Risk:        3,
		Default: &objects.ApplicationsDefault{
			Port: []string{"tcp/80", "tcp/443"},
		},
	}

	// Create the application via API.
	req := client.ApplicationsAPI.CreateApplications(context.Background()).Applications(application)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create application for delete test")
	require.NotNil(t, createRes.Id, "Created application should have an ID")
	createdAppID := *createRes.Id
	require.NotEmpty(t, createdAppID, "Created application ID should not be empty")

	// Test Delete by ID operation.
	reqDel := client.ApplicationsAPI.DeleteApplicationsByID(context.Background(), createdAppID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful.
	require.NoError(t, errDel, "Failed to delete application")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted application: %s", createdAppID)
}

// Test_objects_ApplicationsAPIService_FetchApplications tests the FetchApplications convenience method
func Test_objects_ApplicationsAPIService_FetchApplications(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create test object using same payload as Create test
	testName := "test-app-fetch-" + common.GenerateRandomString(6)
	testObj := objects.Applications{
		Folder:      common.StringPtr("Prisma Access"),
		Name:        testName,
		Description: common.StringPtr("Test application for fetch API"),
		Category:    "business-systems",
		Subcategory: common.StringPtr("ics-protocols"),
		Technology:  common.StringPtr("client-server"),
		Risk:        3,
		Default: &objects.ApplicationsDefault{
			Port: []string{"tcp/80", "tcp/443"},
		},
	}

	createReq := client.ApplicationsAPI.CreateApplications(context.Background()).Applications(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.ApplicationsAPI.DeleteApplicationsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.ApplicationsAPI.FetchApplications(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch applications by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchApplications found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.ApplicationsAPI.FetchApplications(
		context.Background(),
		"non-existent-applications-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchApplications correctly returned nil for non-existent object")
}
