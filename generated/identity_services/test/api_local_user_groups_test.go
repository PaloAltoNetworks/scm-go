/*
Identity Services Testing LocalUserGroupsAPIService
*/
package identity_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// Test_identity_services_LocalUserGroupsAPIService_Create tests creating a local user group
func Test_identity_services_LocalUserGroupsAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Create a test object
	testName := "test-user-grp-" + common.GenerateRandomString(6)
	testObj := identity_services.LocalUserGroups{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
	}

	createReq := client.LocalUserGroupsAPI.CreateLocalUserGroups(context.Background()).LocalUserGroups(testObj)
	resp, httpResp, err := createReq.Execute()

	// Cleanup after test
	if resp != nil && resp.Id != "" {
		defer func() {
			deleteReq := client.LocalUserGroupsAPI.DeleteLocalUserGroupsByID(context.Background(), resp.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", resp.Id)
		}()
	}

	// Verify the response
	require.NoError(t, err, "Failed to create local user group")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 201, httpResp.StatusCode, "Expected 201 Created status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testName, resp.Name, "Created group name should match")
	assert.NotEmpty(t, resp.Id, "Created group should have an ID")
	t.Logf("[SUCCESS] Created local user group: %s (ID: %s)", resp.Name, resp.Id)
}

// Test_identity_services_LocalUserGroupsAPIService_GetByID tests getting a local user group by ID
func Test_identity_services_LocalUserGroupsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Create a test object first
	testName := "test-user-grp-" + common.GenerateRandomString(6)
	testObj := identity_services.LocalUserGroups{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
	}

	createReq := client.LocalUserGroupsAPI.CreateLocalUserGroups(context.Background()).LocalUserGroups(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.LocalUserGroupsAPI.DeleteLocalUserGroupsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Get the object by ID
	getReq := client.LocalUserGroupsAPI.GetLocalUserGroupsByID(context.Background(), createdID)
	resp, httpResp, err := getReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to get local user group by ID")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, createdID, resp.Id, "Retrieved group ID should match")
	assert.Equal(t, testName, resp.Name, "Retrieved group name should match")
	t.Logf("[SUCCESS] Retrieved local user group by ID: %s", resp.Id)
}

// Test_identity_services_LocalUserGroupsAPIService_Update tests updating a local user group
func Test_identity_services_LocalUserGroupsAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Create a test object first
	testName := "test-user-grp-" + common.GenerateRandomString(6)
	testObj := identity_services.LocalUserGroups{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
	}

	createReq := client.LocalUserGroupsAPI.CreateLocalUserGroups(context.Background()).LocalUserGroups(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.LocalUserGroupsAPI.DeleteLocalUserGroupsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Update the group (no-op update to verify API endpoint works)
	updatedObj := identity_services.LocalUserGroups{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
	}

	updateReq := client.LocalUserGroupsAPI.UpdateLocalUserGroupsByID(context.Background(), createdID).LocalUserGroups(updatedObj)
	updateRes, httpResUpdate, errUpdate := updateReq.Execute()
	require.NoError(t, errUpdate, "Failed to update local user group")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, testName, updateRes.Name, "Name should match after update")
	t.Logf("[SUCCESS] Updated local user group: %s", createdID)
}

// Test_identity_services_LocalUserGroupsAPIService_List tests listing local user groups
func Test_identity_services_LocalUserGroupsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Create a test object first
	testName := "test-user-grp-" + common.GenerateRandomString(6)
	testObj := identity_services.LocalUserGroups{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
	}

	createReq := client.LocalUserGroupsAPI.CreateLocalUserGroups(context.Background()).LocalUserGroups(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.LocalUserGroupsAPI.DeleteLocalUserGroupsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// List groups with folder filter
	listReq := client.LocalUserGroupsAPI.ListLocalUserGroups(context.Background()).Folder("Prisma Access")
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list local user groups")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")

	// Verify our created object is in the list
	found := false
	if resp.Data != nil {
		for _, item := range resp.Data {
			if item.Id == createdID {
				found = true
				assert.Equal(t, testName, item.Name, "Listed group name should match")
				break
			}
		}
	}
	assert.True(t, found, "Created group should be in the list")
	t.Logf("[SUCCESS] Listed local user groups, found test object: %s", createdID)
}

// Test_identity_services_LocalUserGroupsAPIService_DeleteByID tests deleting a local user group
func Test_identity_services_LocalUserGroupsAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Create a test object first
	testName := "test-user-grp-" + common.GenerateRandomString(6)
	testObj := identity_services.LocalUserGroups{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
	}

	createReq := client.LocalUserGroupsAPI.CreateLocalUserGroups(context.Background()).LocalUserGroups(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Delete the object
	deleteReq := client.LocalUserGroupsAPI.DeleteLocalUserGroupsByID(context.Background(), createdID)
	httpResp, err := deleteReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to delete local user group")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	t.Logf("[SUCCESS] Deleted local user group: %s", createdID)
}

// Test_identity_services_LocalUserGroupsAPIService_FetchLocalUserGroups tests the FetchLocalUserGroups convenience method
func Test_identity_services_LocalUserGroupsAPIService_FetchLocalUserGroups(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Create a test object first
	testName := "fetch-user-group-" + common.GenerateRandomString(6)
	testObj := identity_services.LocalUserGroups{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
	}

	createReq := client.LocalUserGroupsAPI.CreateLocalUserGroups(context.Background()).LocalUserGroups(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.LocalUserGroupsAPI.DeleteLocalUserGroupsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.LocalUserGroupsAPI.FetchLocalUserGroups(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch local user groups by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchLocalUserGroups found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.LocalUserGroupsAPI.FetchLocalUserGroups(
		context.Background(),
		"non-existent-user-group-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchLocalUserGroups correctly returned nil for non-existent object")
}
