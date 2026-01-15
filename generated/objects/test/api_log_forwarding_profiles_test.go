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

// Helper function to create a COMPLEX LogForwardingProfiles object for testing.
func createComplexTestLogForwardingProfile(nameSuffix string, folder string) objects.LogForwardingProfiles {

	// 1. define the list of match list objects
	matchList := []objects.LogForwardingProfilesMatchListInner{
		{
			Name:       "profile-match-1",
			ActionDesc: common.StringPtr("profile match for tunnel"),
			LogType:    "tunnel",
			Filter:     "(tunnelid neq 123) or (zone.dst eq 192.5.125.155)",
			SendSyslog: []string{"syslog-server-prof-mixed"},
			SendHttp:   []string{"test_http"},
		},
		{
			Name:         "profile-match-2",
			ActionDesc:   common.StringPtr("profile match w/ snmp and email"),
			LogType:      "decryption",
			Filter:       "(addr.src in 10.0.0.0/8)",
			SendSnmptrap: []string{"snmp_test"},
			SendEmail:    []string{"email_test", "email_test_2"},
		},
		{
			Name:         "profile-match-3",
			ActionDesc:   common.StringPtr("profile match w/ all server profiles"),
			LogType:      "traffic",
			Filter:       "(device_name eq test_device)",
			SendSyslog:   []string{"syslog-server-prof-mixed", "syslog-server-prof-complete"},
			SendHttp:     []string{"test_http", "t10", "t5"},
			SendSnmptrap: []string{"snmp_test"},
			SendEmail:    []string{"email_test", "email_test_2"},
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

	// create a profile object
	profile := createComplexTestLogForwardingProfile("create", "All")
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

	// cleanup: delete the created profile to maintain test isolation
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

	// 2. prepare the updated object (adding a second match list entry)
	updatedProfile := *createRes
	updatedProfile.Description = common.StringPtr("Updated Description")
	updatedProfile.MatchList = append(updatedProfile.MatchList, objects.LogForwardingProfilesMatchListInner{
		Name:     "added-match-during-update",
		LogType:  "wildfire",
		Filter:   "(imei contains test_server)",
		SendHttp: []string{"t20"},
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
	reqList := client.LogForwardingProfilesAPI.ListLogForwardingProfiles(context.Background()).Folder("All")
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
