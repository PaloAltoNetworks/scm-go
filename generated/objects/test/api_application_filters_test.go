/*
Objects Testing ApplicationFiltersAPIService
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

// Test_objects_ApplicationFiltersAPIService_Create tests the creation of an application filter object
// This test creates a new application filter and then deletes it to ensure proper cleanup
func Test_objects_ApplicationFiltersAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a valid application filter object with unique name to avoid conflicts
	createdFilterName := "test-app-filter-create-" + common.GenerateRandomString(6)
	applicationFilter := objects.ApplicationFilters{
		Folder:   common.StringPtr("Prisma Access"),
		Name:     createdFilterName,
		Category: []string{"business-systems"},
		Risk:     []int32{1},
		Evasive:  common.BoolPtr(true),

		// Only set boolean fields to true (never false)
		AdditionalProperties: nil,
	}

	fmt.Printf("Creating application filter with name: %s\n", applicationFilter.Name)

	// Make the create request to the API
	req := client.ApplicationFiltersAPI.CreateApplicationFilters(context.Background()).ApplicationFilters(applicationFilter)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create application filter")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdFilterName, res.Name, "Created application filter name should match")
	assert.True(t, *res.Folder == "Shared" || *res.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, []string{"business-systems"}, res.Category, "Category should match")
	assert.Equal(t, common.BoolPtr(true), res.Evasive, "Evasive should match")
	assert.NotEmpty(t, res.Id, "Created application filter should have an ID")

	// Use the ID from the response object
	createdFilterID := res.Id
	t.Logf("Successfully created application filter: %s with ID: %s", applicationFilter.Name, *createdFilterID)

	// Cleanup: Delete the created application filter to maintain test isolation
	reqDel := client.ApplicationFiltersAPI.DeleteApplicationFiltersByID(context.Background(), *createdFilterID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete application filter during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up application filter: %s", *createdFilterID)
}

// Test_objects_ApplicationFiltersAPIService_GetByID tests retrieving an application filter by its ID
// This test creates an application filter, retrieves it by ID, then deletes it
func Test_objects_ApplicationFiltersAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create an application filter first to have something to retrieve
	createdFilterName := "test-app-filter-getbyid-" + common.GenerateRandomString(6)
	applicationFilter := objects.ApplicationFilters{
		Folder:     common.StringPtr("Prisma Access"), // Using Prisma Access folder scope
		Name:       createdFilterName,                 // Unique test name
		Category:   []string{"business-systems"},      // Sample category
		Risk:       []int32{2},                        // Sample risk level
		Technology: []string{"client-server"},         // Sample technology
	}

	// Create the application filter via API
	req := client.ApplicationFiltersAPI.CreateApplicationFilters(context.Background()).ApplicationFilters(applicationFilter)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create application filter for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdFilterID := createRes.Id
	require.NotEmpty(t, createdFilterID, "Created application filter should have an ID")

	// Test Get by ID operation
	reqGetById := client.ApplicationFiltersAPI.GetApplicationFiltersByID(context.Background(), *createdFilterID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get application filter by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdFilterName, getRes.Name, "Application filter name should match")
	assert.True(t, *getRes.Folder == "Shared" || *getRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, []string{"business-systems"}, getRes.Category, "Category should match")
	assert.Equal(t, createdFilterID, getRes.Id, "Application filter ID should match")

	t.Logf("Successfully retrieved application filter: %s", getRes.Name)

	// Cleanup: Delete the created application filter
	reqDel := client.ApplicationFiltersAPI.DeleteApplicationFiltersByID(context.Background(), *createdFilterID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete application filter during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up application filter: %s", *createdFilterID)
}

// Test_objects_ApplicationFiltersAPIService_Update tests updating an existing application filter
// This test creates an application filter, updates it, then deletes it
func Test_objects_ApplicationFiltersAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create an application filter first to have something to update
	createdFilterName := "test-app-filter-update-" + common.GenerateRandomString(6)
	applicationFilter := objects.ApplicationFilters{
		Folder:     common.StringPtr("Prisma Access"), // Using Prisma Access folder scope
		Name:       createdFilterName,                 // Unique test name
		Category:   []string{"business-systems"},      // Initial category
		Risk:       []int32{2},                        // Initial risk level
		Technology: []string{"client-server"},         // Initial technology
	}

	// Create the application filter via API
	req := client.ApplicationFiltersAPI.CreateApplicationFilters(context.Background()).ApplicationFilters(applicationFilter)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create application filter for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdFilterID := createRes.Id
	require.NotEmpty(t, createdFilterID, "Created application filter should have an ID")

	// Test Update operation with modified fields
	updatedFilter := objects.ApplicationFilters{
		Folder:     common.StringPtr("Prisma Access"),          // Keep same folder scope
		Name:       createdFilterName,                          // Keep same name (required for update)
		Category:   []string{"business-systems", "networking"}, // Updated categories
		Risk:       []int32{3, 4},                              // Updated risk levels
		Technology: []string{"client-server", "peer-to-peer"},  // Updated technologies
		Exclude:    []string{"ftp"},                            // New exclusion list
	}

	reqUpdate := client.ApplicationFiltersAPI.UpdateApplicationFiltersByID(context.Background(), *createdFilterID).ApplicationFilters(updatedFilter)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update application filter")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")

	assert.Equal(t, createdFilterName, updateRes.Name, "Application filter name should remain the same")
	assert.True(t, *updateRes.Folder == "Shared" || *updateRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, []string{"business-systems", "networking"}, updateRes.Category, "Category should be updated")
	assert.Equal(t, []int32{3, 4}, updateRes.Risk, "Risk should be updated")
	assert.Equal(t, []string{"client-server", "peer-to-peer"}, updateRes.Technology, "Technology should be updated")
	assert.Equal(t, createdFilterID, updateRes.Id, "Application filter ID should remain the same")

	t.Logf("Successfully updated application filter: %s", createdFilterName)

	// Cleanup: Delete the created application filter
	reqDel := client.ApplicationFiltersAPI.DeleteApplicationFiltersByID(context.Background(), *createdFilterID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete application filter during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up application filter: %s", *createdFilterID)
}

// Test_objects_ApplicationFiltersAPIService_List tests listing application filters with folder filter
// This test creates an application filter, lists filters to verify it's included, then deletes it
func Test_objects_ApplicationFiltersAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create an application filter first to have something to list
	createdFilterName := "test-app-filter-list-" + common.GenerateRandomString(6)
	applicationFilter := objects.ApplicationFilters{
		Folder:     common.StringPtr("Prisma Access"), // Using Prisma Access folder scope
		Name:       createdFilterName,                 // Unique test name
		Category:   []string{"business-systems"},      // Sample category
		Risk:       []int32{2},                        // Sample risk level
		Technology: []string{"client-server"},         // Sample technology
	}

	// Create the application filter via API
	req := client.ApplicationFiltersAPI.CreateApplicationFilters(context.Background()).ApplicationFilters(applicationFilter)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create application filter for list test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdFilterID := createRes.Id
	require.NotEmpty(t, createdFilterID, "Created application filter should have an ID")

	// Test List operation with folder filter
	reqList := client.ApplicationFiltersAPI.ListApplicationFilters(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list application filters")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one application filter in the list")

	// Verify our created application filter is in the list
	foundFilter := false
	for _, filter := range listRes.Data {
		if filter.Name == createdFilterName {
			foundFilter = true
			assert.True(t, *filter.Folder == "Shared" || *filter.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
			assert.Equal(t, []string{"business-systems"}, filter.Category, "Category should match")
			break
		}
	}
	assert.True(t, foundFilter, "Created application filter should be found in the list")

	t.Logf("Successfully listed application filters, found created filter: %s", createdFilterName)

	// Cleanup: Delete the created application filter
	reqDel := client.ApplicationFiltersAPI.DeleteApplicationFiltersByID(context.Background(), *createdFilterID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete application filter during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up application filter: %s", *createdFilterID)
}

// Test_objects_ApplicationFiltersAPIService_DeleteByID tests deleting an application filter by its ID
// This test creates an application filter, deletes it, then verifies the deletion was successful
func Test_objects_ApplicationFiltersAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create an application filter first to have something to delete
	createdFilterName := "test-app-filter-delete-" + common.GenerateRandomString(6)
	applicationFilter := objects.ApplicationFilters{
		Folder:     common.StringPtr("Prisma Access"), // Using Prisma Access folder scope
		Name:       createdFilterName,                 // Unique test name
		Category:   []string{"business-systems"},      // Sample category
		Risk:       []int32{2},                        // Sample risk level
		Technology: []string{"client-server"},         // Sample technology
	}

	// Create the application filter via API
	req := client.ApplicationFiltersAPI.CreateApplicationFilters(context.Background()).ApplicationFilters(applicationFilter)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create application filter for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdFilterID := createRes.Id
	require.NotEmpty(t, createdFilterID, "Created application filter should have an ID")

	// Test Delete by ID operation
	reqDel := client.ApplicationFiltersAPI.DeleteApplicationFiltersByID(context.Background(), *createdFilterID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete application filter")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted application filter: %s", *createdFilterID)
}

// Test_objects_ApplicationFiltersAPIService_FetchApplicationFilters tests the FetchApplicationFilters convenience method
func Test_objects_ApplicationFiltersAPIService_FetchApplicationFilters(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first using same payload as Create test
	testName := "test-appflt-fetch-" + common.GenerateRandomString(6)
	testObj := objects.ApplicationFilters{
		Folder:               common.StringPtr("Prisma Access"),
		Name:                 testName,
		Category:             []string{"business-systems"},
		Risk:                 []int32{1},
		Evasive:              common.BoolPtr(true),
		AdditionalProperties: nil,
	}

	createReq := client.ApplicationFiltersAPI.CreateApplicationFilters(context.Background()).ApplicationFilters(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.ApplicationFiltersAPI.DeleteApplicationFiltersByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.ApplicationFiltersAPI.FetchApplicationFilters(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch application_filters by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchApplicationFilters found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.ApplicationFiltersAPI.FetchApplicationFilters(
		context.Background(),
		"non-existent-application_filters-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchApplicationFilters correctly returned nil for non-existent object")
}
