/*
Objects Testing HTTPServerProfilesAPIService
*/
package objects

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)

// Test_objects_HTTPServerProfilesAPIService_Create tests creating an HTTP server profile
func Test_objects_HTTPServerProfilesAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object
	testName := "test-http-srv-" + common.GenerateRandomString(6)

	server := objects.HttpServerProfilesServerInner{
		Name:       common.StringPtr("test-server-1"),
		Address:    common.StringPtr("192.0.2.1"),
		Port:       common.Int32Ptr(443),
		Protocol:   common.StringPtr("HTTPS"),
		HttpMethod: common.StringPtr("GET"),
	}

	testObj := objects.HttpServerProfiles{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
		Server: []objects.HttpServerProfilesServerInner{server},
	}

	createReq := client.HTTPServerProfilesAPI.CreateHTTPServerProfiles(context.Background()).HttpServerProfiles(testObj)
	resp, httpResp, err := createReq.Execute()

	// Cleanup after test
	if resp != nil && resp.Id != "" {
		defer func() {
			deleteReq := client.HTTPServerProfilesAPI.DeleteHTTPServerProfilesByID(context.Background(), resp.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", resp.Id)
		}()
	}

	// Verify the response
	require.NoError(t, err, "Failed to create HTTP server profile")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 201, httpResp.StatusCode, "Expected 201 Created status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testName, resp.Name, "Created profile name should match")
	assert.NotEmpty(t, resp.Id, "Created profile should have an ID")
	t.Logf("[SUCCESS] Created HTTP server profile: %s (ID: %s)", resp.Name, resp.Id)
}

// Test_objects_HTTPServerProfilesAPIService_GetByID tests getting an HTTP server profile by ID
func Test_objects_HTTPServerProfilesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first
	testName := "test-http-srv-" + common.GenerateRandomString(6)

	server := objects.HttpServerProfilesServerInner{
		Name:       common.StringPtr("test-server-1"),
		Address:    common.StringPtr("192.0.2.1"),
		Port:       common.Int32Ptr(443),
		Protocol:   common.StringPtr("HTTPS"),
		HttpMethod: common.StringPtr("GET"),
	}

	testObj := objects.HttpServerProfiles{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
		Server: []objects.HttpServerProfilesServerInner{server},
	}

	createReq := client.HTTPServerProfilesAPI.CreateHTTPServerProfiles(context.Background()).HttpServerProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.HTTPServerProfilesAPI.DeleteHTTPServerProfilesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Get the object by ID
	getReq := client.HTTPServerProfilesAPI.GetHTTPServerProfilesByID(context.Background(), createdID)
	resp, httpResp, err := getReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to get HTTP server profile by ID")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, createdID, resp.Id, "Retrieved profile ID should match")
	assert.Equal(t, testName, resp.Name, "Retrieved profile name should match")
	t.Logf("[SUCCESS] Retrieved HTTP server profile by ID: %s", resp.Id)
}

// Test_objects_HTTPServerProfilesAPIService_Update tests updating an HTTP server profile
func Test_objects_HTTPServerProfilesAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first
	testName := "test-http-srv-" + common.GenerateRandomString(6)

	server := objects.HttpServerProfilesServerInner{
		Name:       common.StringPtr("test-server-1"),
		Address:    common.StringPtr("192.0.2.1"),
		Port:       common.Int32Ptr(443),
		Protocol:   common.StringPtr("HTTPS"),
		HttpMethod: common.StringPtr("GET"),
	}

	testObj := objects.HttpServerProfiles{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
		Server: []objects.HttpServerProfilesServerInner{server},
	}

	createReq := client.HTTPServerProfilesAPI.CreateHTTPServerProfiles(context.Background()).HttpServerProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.HTTPServerProfilesAPI.DeleteHTTPServerProfilesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Update the object with different server configuration
	updatedServer := objects.HttpServerProfilesServerInner{
		Name:       common.StringPtr("test-server-1"),
		Address:    common.StringPtr("192.0.2.2"),
		Port:       common.Int32Ptr(8443),
		Protocol:   common.StringPtr("HTTPS"),
		HttpMethod: common.StringPtr("POST"),
	}

	updateObj := objects.HttpServerProfiles{
		Id:     createdID,
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
		Server: []objects.HttpServerProfilesServerInner{updatedServer},
	}

	updateReq := client.HTTPServerProfilesAPI.UpdateHTTPServerProfilesByID(context.Background(), createdID).HttpServerProfiles(updateObj)
	resp, httpResp, err := updateReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to update HTTP server profile")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, createdID, resp.Id, "Updated profile ID should match")

	// Verify the server configuration was updated
	if len(resp.Server) > 0 {
		assert.Equal(t, "192.0.2.2", *resp.Server[0].Address, "Server address should be updated")
		assert.Equal(t, int32(8443), *resp.Server[0].Port, "Server port should be updated")
	}
	t.Logf("[SUCCESS] Updated HTTP server profile: %s", resp.Id)
}

// Test_objects_HTTPServerProfilesAPIService_List tests listing HTTP server profiles
func Test_objects_HTTPServerProfilesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first
	testName := "test-http-srv-" + common.GenerateRandomString(6)

	server := objects.HttpServerProfilesServerInner{
		Name:       common.StringPtr("test-server-1"),
		Address:    common.StringPtr("192.0.2.1"),
		Port:       common.Int32Ptr(443),
		Protocol:   common.StringPtr("HTTPS"),
		HttpMethod: common.StringPtr("GET"),
	}

	testObj := objects.HttpServerProfiles{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
		Server: []objects.HttpServerProfilesServerInner{server},
	}

	createReq := client.HTTPServerProfilesAPI.CreateHTTPServerProfiles(context.Background()).HttpServerProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.HTTPServerProfilesAPI.DeleteHTTPServerProfilesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// List profiles with folder filter
	listReq := client.HTTPServerProfilesAPI.ListHTTPServerProfiles(context.Background()).Folder("Prisma Access")
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list HTTP server profiles")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")

	// Verify our created object is in the list
	found := false
	if resp.Data != nil {
		for _, item := range resp.Data {
			if item.Id == createdID {
				found = true
				assert.Equal(t, testName, item.Name, "Listed profile name should match")
				break
			}
		}
	}
	assert.True(t, found, "Created profile should be in the list")
	t.Logf("[SUCCESS] Listed HTTP server profiles, found test object: %s", createdID)
}

// Test_objects_HTTPServerProfilesAPIService_DeleteByID tests deleting an HTTP server profile
func Test_objects_HTTPServerProfilesAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first
	testName := "test-http-srv-" + common.GenerateRandomString(6)

	server := objects.HttpServerProfilesServerInner{
		Name:       common.StringPtr("test-server-1"),
		Address:    common.StringPtr("192.0.2.1"),
		Port:       common.Int32Ptr(443),
		Protocol:   common.StringPtr("HTTPS"),
		HttpMethod: common.StringPtr("GET"),
	}

	testObj := objects.HttpServerProfiles{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
		Server: []objects.HttpServerProfilesServerInner{server},
	}

	createReq := client.HTTPServerProfilesAPI.CreateHTTPServerProfiles(context.Background()).HttpServerProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Delete the object
	deleteReq := client.HTTPServerProfilesAPI.DeleteHTTPServerProfilesByID(context.Background(), createdID)
	httpResp, err := deleteReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to delete HTTP server profile")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	t.Logf("[SUCCESS] Deleted HTTP server profile: %s", createdID)
}

// Test_objects_HTTPServerProfilesAPIService_FetchHTTPServerProfiles tests the FetchHTTPServerProfiles convenience method
func Test_objects_HTTPServerProfilesAPIService_FetchHTTPServerProfiles(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first
	testName := "fetch-http-srv-" + common.GenerateRandomString(6)

	// Create minimal server configuration (required field)
	server := objects.HttpServerProfilesServerInner{
		Name:       common.StringPtr("test-server-1"),
		Address:    common.StringPtr("192.0.2.1"),
		Port:       common.Int32Ptr(443),
		Protocol:   common.StringPtr("HTTPS"),
		HttpMethod: common.StringPtr("GET"),
	}

	testObj := objects.HttpServerProfiles{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
		Server: []objects.HttpServerProfilesServerInner{server},
	}

	createReq := client.HTTPServerProfilesAPI.CreateHTTPServerProfiles(context.Background()).HttpServerProfiles(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.HTTPServerProfilesAPI.DeleteHTTPServerProfilesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.HTTPServerProfilesAPI.FetchHTTPServerProfiles(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch http_server_profiles by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchHTTPServerProfiles found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.HTTPServerProfilesAPI.FetchHTTPServerProfiles(
		context.Background(),
		"non-existent-http_server_profiles-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchHTTPServerProfiles correctly returned nil for non-existent object")
}
