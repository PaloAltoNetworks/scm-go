/*
Objects Testing LogForwardingProfilesAPIService
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

// Helper function to create a minimal LogForwardingProfiles object for testing
func createTestLogForwardingProfile(nameSuffix string, folder string) objects.LogForwardingProfiles {
	matchList := []objects.LogForwardingProfilesMatchListInner{
		{
			Name:    "profile-match",
			LogType: "auth",
			Filter:  "All Logs",
		},
	}

	return objects.LogForwardingProfiles{
		Name:      "test-log-fwd-" + nameSuffix + "-" + common.GenerateRandomString(5),
		Folder:    common.StringPtr(folder),
		MatchList: matchList,
	}
}

// Helper function to create an HTTP server profile for testing
func createTestHttpServerProfile(client *objects.APIClient, t *testing.T, nameSuffix string, folder string) (string, string) {
	httpProfile := objects.HttpServerProfiles{
		Name:   "test-http-" + nameSuffix + "-" + common.GenerateRandomString(5),
		Folder: common.StringPtr(folder),
		Server: []objects.HttpServerProfilesServerInner{
			{
				Name:       common.StringPtr("http-server-1"),
				Address:    common.StringPtr("192.168.1.100"),
				Port:       common.Int32Ptr(8080),
				Protocol:   common.StringPtr("HTTP"),
				HttpMethod: common.StringPtr("POST"),
			},
		},
	}

	req := client.HTTPServerProfilesAPI.CreateHTTPServerProfiles(context.Background()).HttpServerProfiles(httpProfile)
	res, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create HTTP server profile")
	t.Logf("Created HTTP server profile: %s with ID: %s", httpProfile.Name, res.Id)
	return httpProfile.Name, res.Id
}

// Helper function to create a Syslog server profile for testing
func createTestSyslogServerProfile(client *objects.APIClient, t *testing.T, nameSuffix string, folder string) (string, string) {
	syslogProfile := objects.SyslogServerProfiles{
		Name:   "test-syslog-" + nameSuffix + "-" + common.GenerateRandomString(5),
		Folder: common.StringPtr(folder),
		Server: []objects.SyslogServerProfilesServerInner{
			{
				Name:      common.StringPtr("syslog-server-1"),
				Server:    common.StringPtr("192.168.1.101"),
				Port:      common.Int32Ptr(514),
				Facility:  common.StringPtr("LOG_USER"),
				Transport: common.StringPtr("UDP"),
			},
		},
	}

	req := client.SyslogServerProfilesAPI.CreateSyslogServerProfiles(context.Background()).SyslogServerProfiles(syslogProfile)
	res, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Syslog server profile")
	t.Logf("Created Syslog server profile: %s with ID: %s", syslogProfile.Name, res.Id)
	return syslogProfile.Name, res.Id
}

// Helper function to delete an HTTP server profile
func deleteTestHttpServerProfile(client *objects.APIClient, t *testing.T, id string) {
	req := client.HTTPServerProfilesAPI.DeleteHTTPServerProfilesByID(context.Background(), id)
	_, err := req.Execute()
	if err != nil {
		t.Logf("Warning: Failed to delete HTTP server profile %s: %v", id, err)
	} else {
		t.Logf("Deleted HTTP server profile: %s", id)
	}
}

// Helper function to delete a Syslog server profile
func deleteTestSyslogServerProfile(client *objects.APIClient, t *testing.T, id string) {
	req := client.SyslogServerProfilesAPI.DeleteSyslogServerProfilesByID(context.Background(), id)
	_, err := req.Execute()
	if err != nil {
		t.Logf("Warning: Failed to delete Syslog server profile %s: %v", id, err)
	} else {
		t.Logf("Deleted Syslog server profile: %s", id)
	}
}

// Helper function to create a COMPLEX LogForwardingProfiles object for testing.
// Requires httpProfileName and syslogProfileName to be created first.
func createComplexTestLogForwardingProfile(nameSuffix string, folder string, httpProfileName string, syslogProfileName string) objects.LogForwardingProfiles {

	// 1. define the list of match list objects using the provided profile names
	matchList := []objects.LogForwardingProfilesMatchListInner{
		{
			Name:       "profile-match-1",
			ActionDesc: common.StringPtr("profile match for tunnel with syslog"),
			LogType:    "tunnel",
			Filter:     "(tunnelid neq 123) or (zone.dst eq 192.5.125.155)",
			SendSyslog: []string{syslogProfileName},
		},
		{
			Name:       "profile-match-2",
			ActionDesc: common.StringPtr("profile match with http"),
			LogType:    "decryption",
			Filter:     "(addr.src in 10.0.0.0/8)",
			SendHttp:   []string{httpProfileName},
		},
		{
			Name:       "profile-match-3",
			ActionDesc: common.StringPtr("profile match with both http and syslog"),
			LogType:    "traffic",
			Filter:     "(device_name eq test_device)",
			SendSyslog: []string{syslogProfileName},
			SendHttp:   []string{httpProfileName},
		},
	}

	// 2. return the complete log forwarding profiles object
	return objects.LogForwardingProfiles{
		Name:        "test-log-fwd-" + nameSuffix + "-" + common.GenerateRandomString(5),
		Folder:      common.StringPtr(folder),
		Description: common.StringPtr("Log Forwarding w/ Multiple Match Lists"),
		MatchList:   matchList,
	}
}

// Test_objects_LogForwardingProfilesAPIService_Create tests the creation of a log forwarding profile
func Test_objects_LogForwardingProfilesAPIService_Create(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	folder := "All"

	// Step 1: Create dependency profiles first
	httpProfileName, httpProfileID := createTestHttpServerProfile(client, t, "create", folder)
	syslogProfileName, syslogProfileID := createTestSyslogServerProfile(client, t, "create", folder)

	// Defer cleanup of dependencies (will run last, after log forwarding profile cleanup)
	defer deleteTestHttpServerProfile(client, t, httpProfileID)
	defer deleteTestSyslogServerProfile(client, t, syslogProfileID)

	// Step 2: Create the complex log forwarding profile using the dependency names
	profile := createComplexTestLogForwardingProfile("create", folder, httpProfileName, syslogProfileName)
	profileName := profile.Name

	// make the create request to the API
	req := client.LogForwardingProfilesAPI.CreateLogForwardingProfiles(context.Background()).LogForwardingProfiles(profile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// verify the creation was successful
	require.NoError(t, err, "Failed to create log forwarding profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
	assert.NotEmpty(t, res.Id, "Created profile should have an ID")
	assert.Equal(t, 3, len(res.MatchList), "Match list should contain 3 items")

	// use the id from the response object
	createdID := *res.Id
	t.Logf("Successfully created log forwarding profile: %s with ID: %s", profileName, createdID)

	// cleanup: delete the created log forwarding profile first (before dependencies)
	reqDel := client.LogForwardingProfilesAPI.DeleteLogForwardingProfilesByID(context.Background(), createdID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete log forwarding profile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up log forwarding profile: %s", createdID)
}

// Test_objects_LogForwardingProfilesAPIService_GetByID tests retrieving a log forwarding profile by its ID
func Test_objects_LogForwardingProfilesAPIService_GetByID(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	// 1. Create a profile first to have something to retrieve
	profile := createTestLogForwardingProfile("getbyid", "All")
	req := client.LogForwardingProfilesAPI.CreateLogForwardingProfiles(context.Background()).LogForwardingProfiles(profile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create profile for get test")
	createdID := *createRes.Id

	// 2. Test Get by ID operation
	reqGetById := client.LogForwardingProfilesAPI.GetLogForwardingProfilesByID(context.Background(), createdID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// 3. Verify the get operation was successful
	require.NoError(t, err, "Failed to get log forwarding profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// 4. Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profile.Name, getRes.Name, "Profile name should match")
	assert.Equal(t, createdID, *getRes.Id, "Profile ID should match")
	assert.Equal(t, 1, len(getRes.MatchList), "Match list should contain 1 item")

	t.Logf("Successfully retrieved log forwarding profile: %s", getRes.Name)

	// 5. Cleanup: Delete the created profile
	reqDel := client.LogForwardingProfilesAPI.DeleteLogForwardingProfilesByID(context.Background(), createdID)
	_, _ = reqDel.Execute()
}

// Test_objects_LogForwardingProfilesAPIService_Update tests updating an existing log forwarding profile
func Test_objects_LogForwardingProfilesAPIService_Update(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	// 1. create a profile first
	profile := createTestLogForwardingProfile("update", "All")
	req := client.LogForwardingProfilesAPI.CreateLogForwardingProfiles(context.Background()).LogForwardingProfiles(profile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create profile for update test")
	createdID := *createRes.Id

	// 2. prepare the updated object (adding a second match list entry without external references)
	updatedProfile := *createRes
	updatedProfile.Description = common.StringPtr("Updated Description")
	updatedProfile.MatchList = append(updatedProfile.MatchList, objects.LogForwardingProfilesMatchListInner{
		Name:    "added-match-during-update",
		LogType: "wildfire",
		Filter:  "(imei contains test_server)",
	})

	// 3. test update operation
	reqUpdate := client.LogForwardingProfilesAPI.UpdateLogForwardingProfilesByID(context.Background(), createdID).LogForwardingProfiles(updatedProfile)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// 4. verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update log forwarding profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// 5. assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, 2, len(updateRes.MatchList), "Match list should now contain 2 items")
	assert.Equal(t, "Updated Description", *updateRes.Description, "Description should be updated")

	t.Logf("Successfully updated log forwarding profile: %s", profile.Name)

	// 6. cleanup: delete the created profile
	reqDel := client.LogForwardingProfilesAPI.DeleteLogForwardingProfilesByID(context.Background(), createdID)
	_, _ = reqDel.Execute()
}

// Test_objects_LogForwardingProfilesAPIService_List tests listing log forwarding profiles with folder filter
func Test_objects_LogForwardingProfilesAPIService_List(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	// 1. Create a resource to ensure the list is non-empty
	profile := createTestLogForwardingProfile("list-test", "All")
	profileName := profile.Name

	createRes, _, err := client.LogForwardingProfilesAPI.CreateLogForwardingProfiles(context.Background()).LogForwardingProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create log forwarding profile for list test")
	createdID := *createRes.Id
	require.NotEmpty(t, createdID, "Created profile ID should not be empty")

	// 2. Setup deferred cleanup
	defer func() {
		t.Logf("Cleaning up log forwarding profile with ID: %s", createdID)
		_, errDel := client.LogForwardingProfilesAPI.DeleteLogForwardingProfilesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete log forwarding profile during cleanup")
	}()

	// 3. Test List operation with folder filter
	reqList := client.LogForwardingProfilesAPI.ListLogForwardingProfiles(context.Background()).Folder("All").Limit(200).Offset(0)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// 4. Verify the list operation was successful
	require.NoError(t, errList, "Failed to list log forwarding profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")

	// 5. Verify the created object is in the list
	foundObject := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created log forwarding profile should be found in the list")

	t.Logf("Found %d log forwarding profiles in the list.", len(listRes.Data))
}

// Test_objects_LogForwardingProfilesAPIService_DeleteByID tests deleting a log forwarding profile by its ID
func Test_objects_LogForwardingProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	// 1. create a profile first to have something to delete
	profile := createTestLogForwardingProfile("delete", "All")
	req := client.LogForwardingProfilesAPI.CreateLogForwardingProfiles(context.Background()).LogForwardingProfiles(profile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create profile for delete test")
	createdID := *createRes.Id

	// 2. test Delete by ID operation
	reqDel := client.LogForwardingProfilesAPI.DeleteLogForwardingProfilesByID(context.Background(), createdID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// 3. verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete log forwarding profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted log forwarding profile: %s", createdID)
}

// Test_objects_LogForwardingProfilesAPIService_FetchLogForwardingProfiles tests the FetchLogForwardingProfiles convenience method
func Test_objects_LogForwardingProfilesAPIService_FetchLogForwardingProfiles(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)
	folder := "All"

	// Step 1: Create dependency profiles first
	httpProfileName, httpProfileID := createTestHttpServerProfile(client, t, "fetch", folder)
	syslogProfileName, syslogProfileID := createTestSyslogServerProfile(client, t, "fetch", folder)

	// Defer cleanup of dependencies (will run last)
	defer deleteTestHttpServerProfile(client, t, httpProfileID)
	defer deleteTestSyslogServerProfile(client, t, syslogProfileID)

	// Step 2: Create test object using complex payload with dependencies
	testObj := createComplexTestLogForwardingProfile("fetch", folder, httpProfileName, syslogProfileName)
	testName := testObj.Name

	createReq := client.LogForwardingProfilesAPI.CreateLogForwardingProfiles(context.Background()).LogForwardingProfiles(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	// Cleanup log forwarding profile after test (before dependencies)
	defer func() {
		deleteReq := client.LogForwardingProfilesAPI.DeleteLogForwardingProfilesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.LogForwardingProfilesAPI.FetchLogForwardingProfiles(
		context.Background(),
		testName,
		common.StringPtr(folder),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch log_forwarding_profiles by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchLogForwardingProfiles found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.LogForwardingProfilesAPI.FetchLogForwardingProfiles(
		context.Background(),
		"non-existent-log_forwarding_profiles-xyz-12345",
		common.StringPtr(folder),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchLogForwardingProfiles correctly returned nil for non-existent object")
}
