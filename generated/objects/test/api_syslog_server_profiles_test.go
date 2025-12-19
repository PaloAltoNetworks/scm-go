/*
Objects Testing SyslogServerProfilesAPIService
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

// Helper function to create a minimal SyslogServerProfiles object for testing
func createTestSyslogProfile(nameSuffix string, folder string) objects.SyslogServerProfiles {
	serverList := []objects.SyslogServerProfilesServerInner{
		{
			Name:   common.StringPtr("TestServer"),
			Server: common.StringPtr("192.0.2.1"),
		},
	}

	return objects.SyslogServerProfiles{
		Name:   "test-syslog-" + nameSuffix + "-" + common.GenerateRandomString(5),
		Folder: common.StringPtr(folder),
		Server: serverList,
	}
}

// Helper function to create a COMPLEX SyslogServerProfiles object for testing.
func createComplexTestSyslogProfile(nameSuffix string, folder string) objects.SyslogServerProfiles {

	// 1. define the list of server objects (two servers)
	serverList := []objects.SyslogServerProfilesServerInner{
		{
			Name:      common.StringPtr("Server-A"),
			Server:    common.StringPtr("172.16.10.1"),
			Transport: common.StringPtr("UDP"),
			Port:      common.Int32Ptr(514),
			Format:    common.StringPtr("BSD"),
			Facility:  common.StringPtr("LOG_LOCAL7"),
		},
		{
			Name:      common.StringPtr("Server-B"),
			Server:    common.StringPtr("172.16.10.2"),
			Transport: common.StringPtr("TCP"),
			Port:      common.Int32Ptr(6514),
			Format:    common.StringPtr("IETF"),
			Facility:  common.StringPtr("LOG_LOCAL3"),
		},
	}

	// 2. define the format object, including escaping
	formatConfig := &objects.SyslogServerProfilesFormat{
		Escaping: &objects.SyslogServerProfilesFormatEscaping{
			EscapeCharacter:   common.StringPtr("*"),
			EscapedCharacters: common.StringPtr("&\\#"),
		},
		Traffic:       common.StringPtr("$error + $errorcode"),
		Threat:        common.StringPtr("$client_os"),
		Wildfire:      common.StringPtr("default"),
		Url:           common.StringPtr("$device_name and $contenttype"),
		Data:          common.StringPtr("$status"),
		Gtp:           common.StringPtr("dg_hier_level_4"),
		Sctp:          common.StringPtr("$srcregion"),
		Tunnel:        common.StringPtr("$tunnel_type"),
		Auth:          common.StringPtr("$location"),
		Userid:        common.StringPtr("$host_id"),
		Iptag:         common.StringPtr("$vsys_name"),
		Decryption:    common.StringPtr("default"),
		Config:        common.StringPtr("custom"),
		System:        common.StringPtr("default"),
		Globalprotect: common.StringPtr("$type"),
		HipMatch:      common.StringPtr("$actionflags"),
		Correlation:   common.StringPtr("$error"),
	}

	// 3. return the complete syslog server profiles object
	return objects.SyslogServerProfiles{
		Name:   "test-syslog-" + nameSuffix + "-" + common.GenerateRandomString(5),
		Folder: common.StringPtr(folder),
		Server: serverList,
		Format: formatConfig,
	}
}

// Test_objects_SyslogServerProfilesAPIService_Create tests the creation of a syslog server profile
func Test_objects_SyslogServerProfilesAPIService_Create(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	// create a profile object
	profile := createComplexTestSyslogProfile("create", "All")
	profileName := profile.Name

	// make the create request to the API
	req := client.SyslogServerProfilesAPI.CreateSyslogServerProfiles(context.Background()).SyslogServerProfiles(profile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// verify the creation was successful
	require.NoError(t, err, "Failed to create syslog server profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
	assert.NotEmpty(t, res.Id, "Created profile should have an ID")
	assert.NotEmpty(t, res.Server, "Server list should not be empty")
	assert.Equal(t, 2, len(res.Server), "Server list should contain 2 item")

	// use the id from the response object
	createdProfileID := res.Id
	t.Logf("Successfully created syslog profile: %s with ID: %s", profileName, createdProfileID)

	// cleanup: delete the created profile to maintain test isolation
	reqDel := client.SyslogServerProfilesAPI.DeleteSyslogServerProfilesByID(context.Background(), createdProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete syslog profile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up syslog profile: %s", createdProfileID)
}

// Test_objects_SyslogServerProfilesAPIService_GetByID tests retrieving a syslog server profile by its ID
func Test_objects_SyslogServerProfilesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// 1. Create a profile first to have something to retrieve
	profile := createTestSyslogProfile("getbyid", "All")
	req := client.SyslogServerProfilesAPI.CreateSyslogServerProfiles(context.Background()).SyslogServerProfiles(profile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create profile for get test")
	createdProfileID := createRes.Id

	// 2. Test Get by ID operation
	reqGetById := client.SyslogServerProfilesAPI.GetSyslogServerProfilesByID(context.Background(), createdProfileID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// 3. Verify the get operation was successful
	require.NoError(t, err, "Failed to get syslog server profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// 4. Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profile.Name, getRes.Name, "Profile name should match")
	assert.Equal(t, createdProfileID, getRes.Id, "Profile ID should match")
	assert.Equal(t, 1, len(getRes.Server), "Server list should contain 1 item")

	t.Logf("Successfully retrieved syslog profile: %s", getRes.Name)

	// 5. Cleanup: Delete the created profile
	reqDel := client.SyslogServerProfilesAPI.DeleteSyslogServerProfilesByID(context.Background(), createdProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete profile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
}

// Test_objects_SyslogServerProfilesAPIService_Update tests updating an existing syslog server profile
func Test_objects_SyslogServerProfilesAPIService_Update(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	// 1. create a profile first
	profile := createTestSyslogProfile("update", "All")
	req := client.SyslogServerProfilesAPI.CreateSyslogServerProfiles(context.Background()).SyslogServerProfiles(profile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create profile for update test")
	createdProfileID := createRes.Id

	// 2. prepare the updated object (adding a second server and a format config)
	updatedServerList := append(profile.Server, objects.SyslogServerProfilesServerInner{
		Name:      common.StringPtr("TestServer-B"),
		Server:    common.StringPtr("192.0.2.2"),
		Transport: common.StringPtr("TCP"),
		Port:      common.Int32Ptr(601),
		Format:    common.StringPtr("IETF"),
		Facility:  common.StringPtr("LOG_LOCAL7"),
	})

	updatedProfile := objects.SyslogServerProfiles{
		Folder: common.StringPtr("All"),
		Name:   profile.Name,
		Server: updatedServerList,
		Format: &objects.SyslogServerProfilesFormat{
			Traffic: common.StringPtr("default"),
			Threat:  common.StringPtr("default"),
			Escaping: &objects.SyslogServerProfilesFormatEscaping{
				EscapeCharacter:   common.StringPtr("\\"),
				EscapedCharacters: common.StringPtr("&"),
			},
		},
	}

	// 3. test update operation
	reqUpdate := client.SyslogServerProfilesAPI.UpdateSyslogServerProfilesByID(context.Background(), createdProfileID).SyslogServerProfiles(updatedProfile)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// 4. verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update syslog server profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// 5. assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, 2, len(updateRes.Server), "Server list should now contain 2 items")
	require.NotNil(t, updateRes.Format, "Format should be present after update")
	require.NotNil(t, updateRes.Format.Escaping, "Escaping should be present after update")
	assert.Equal(t, common.StringPtr("\\"), updateRes.Format.Escaping.EscapeCharacter, "Escape character should be updated")

	t.Logf("Successfully updated syslog profile: %s", profile.Name)

	// 6. cleanup: delete the created profile
	reqDel := client.SyslogServerProfilesAPI.DeleteSyslogServerProfilesByID(context.Background(), createdProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete profile during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
}

// Test_objects_SyslogServerProfilesAPIService_List tests listing syslog server profiles with folder filter
func Test_objects_SyslogServerProfilesAPIService_List(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	// test List operation with folder filter
	reqList := client.SyslogServerProfilesAPI.ListSyslogServerProfiles(context.Background()).Folder("All")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// verify the list operation was successful
	require.NoError(t, errList, "Failed to list syslog server profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.GreaterOrEqual(t, len(listRes.Data), 0, "Should handle zero or more profiles in the list")
	t.Logf("Found %d syslog server profiles in the list.", len(listRes.Data))
}

// Test_objects_SyslogServerProfilesAPIService_DeleteByID tests deleting a syslog server profile by its ID
func Test_objects_SyslogServerProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	// 1. create a profile first to have something to delete
	profile := createTestSyslogProfile("delete", "All")
	req := client.SyslogServerProfilesAPI.CreateSyslogServerProfiles(context.Background()).SyslogServerProfiles(profile)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create profile for delete test")
	createdProfileID := createRes.Id

	// 2. test Delete by ID operation
	reqDel := client.SyslogServerProfilesAPI.DeleteSyslogServerProfilesByID(context.Background(), createdProfileID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// 3. verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete syslog server profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted syslog profile: %s", createdProfileID)
}
