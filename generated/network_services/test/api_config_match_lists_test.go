/*
 * Network Services Testing
 *
 * ConfigMatchListAPIService
 */

package network_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	setup "github.com/paloaltonetworks/scm-go"
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)

// Helper function to get objects API client for creating dependencies
func getObjectsClient(t *testing.T) *objects.APIClient {
	configPath := common.GetConfigPath()
	setupClient := &setup.Client{
		AuthFile:         configPath,
		CheckEnvironment: false,
	}
	err := setupClient.Setup()
	require.NoError(t, err, "Failed to setup objects client")
	ctx := context.Background()
	if setupClient.Jwt == "" {
		err = setupClient.RefreshJwt(ctx)
		require.NoError(t, err, "Failed to refresh JWT for objects client")
	}
	return setup.GetObjectsAPIClient(setupClient)
}

// Helper function to create an HTTP server profile for testing
func createTestHTTPServerProfile(objClient *objects.APIClient, t *testing.T, nameSuffix string, folder string) (string, string) {
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

	req := objClient.HTTPServerProfilesAPI.CreateHTTPServerProfiles(context.Background()).HttpServerProfiles(httpProfile)
	res, _, err := req.Execute()
	if err != nil {
		if apiErr, ok := err.(*objects.GenericOpenAPIError); ok {
			t.Logf("API Error Body: %s", string(apiErr.Body()))
		}
		t.Logf("Failed to create HTTP server profile: %v", err)
	}
	require.NoError(t, err, "Failed to create HTTP server profile")
	t.Logf("Created HTTP server profile: %s with ID: %s", httpProfile.Name, res.Id)
	return httpProfile.Name, res.Id
}

// Helper function to create a Syslog server profile for testing
func createTestSyslogServerProfile(objClient *objects.APIClient, t *testing.T, nameSuffix string, folder string) (string, string) {
	syslogProfile := objects.SyslogServerProfiles{
		Name:   "sys-" + nameSuffix + "-" + common.GenerateRandomString(5),
		Folder: common.StringPtr(folder),
		Server: []objects.SyslogServerProfilesServerInner{
			{
				Name:      common.StringPtr("syslog-server-1"),
				Server:    common.StringPtr("192.168.1.101"),
				Port:      common.Int32Ptr(514),
				Format:    common.StringPtr("BSD"),
				Facility:  common.StringPtr("LOG_USER"),
				Transport: common.StringPtr("UDP"),
			},
		},
	}

	req := objClient.SyslogServerProfilesAPI.CreateSyslogServerProfiles(context.Background()).SyslogServerProfiles(syslogProfile)
	res, _, err := req.Execute()
	if err != nil {
		if apiErr, ok := err.(*objects.GenericOpenAPIError); ok {
			t.Logf("API Error Body: %s", string(apiErr.Body()))
		}
		t.Logf("Failed to create Syslog server profile: %v", err)
	}
	require.NoError(t, err, "Failed to create Syslog server profile")
	t.Logf("Created Syslog server profile: %s with ID: %s", syslogProfile.Name, res.Id)
	return syslogProfile.Name, res.Id
}

// Helper function to delete an HTTP server profile
func deleteTestHTTPServerProfile(objClient *objects.APIClient, t *testing.T, id string) {
	req := objClient.HTTPServerProfilesAPI.DeleteHTTPServerProfilesByID(context.Background(), id)
	_, err := req.Execute()
	if err != nil {
		t.Logf("Warning: Failed to delete HTTP server profile %s: %v", id, err)
	} else {
		t.Logf("Deleted HTTP server profile: %s", id)
	}
}

// Helper function to delete a Syslog server profile
func deleteTestSyslogServerProfile(objClient *objects.APIClient, t *testing.T, id string) {
	req := objClient.SyslogServerProfilesAPI.DeleteSyslogServerProfilesByID(context.Background(), id)
	_, err := req.Execute()
	if err != nil {
		t.Logf("Warning: Failed to delete Syslog server profile %s: %v", id, err)
	} else {
		t.Logf("Deleted Syslog server profile: %s", id)
	}
}

// Test_networkservices_ConfigMatchListAPIService_Create tests the creation of a Config Match List.
func Test_networkservices_ConfigMatchListAPIService_Create(t *testing.T) {
	// Setup clients
	client := SetupNetworkSvcTestClient(t)
	objClient := getObjectsClient(t)
	folder := "ngfw-shared"

	// Step 1: Create dependency profiles first
	httpProfileName, httpProfileID := createTestHTTPServerProfile(objClient, t, "cfgmatch-create", folder)
	syslogProfileName, syslogProfileID := createTestSyslogServerProfile(objClient, t, "cfgmatch-create", folder)

	// Defer cleanup of dependencies (will run last)
	defer deleteTestHTTPServerProfile(objClient, t, httpProfileID)
	defer deleteTestSyslogServerProfile(objClient, t, syslogProfileID)

	// Step 2: Create Config Match List with dependencies
	matchListName := "test-config-list-" + common.GenerateRandomString(10)

	matchList := network_services.ConfigMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("Config match list for tracking configuration changes and audit log forwarding"),
		Folder:         common.StringPtr(folder),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{syslogProfileName},
		SendHttp:       []string{httpProfileName},
		SendToPanorama: common.BoolPtr(false),
	}

	fmt.Printf("Attempting to create Config Match List with name: %s\n", matchList.Name)

	// Make the create request to the API.
	req := client.ConfigMatchListAPI.CreateConfigMatchList(context.Background()).ConfigMatchList(matchList)
	res, httpRes, err := req.Execute()

	// Defer cleanup for the Config Match List (runs before dependency cleanup)
	if res != nil && res.Id != nil {
		defer func() {
			t.Logf("Cleaning up Config Match List with ID: %s", *res.Id)
			delReq := client.ConfigMatchListAPI.DeleteConfigMatchListByID(context.Background(), *res.Id)
			_, errDel := delReq.Execute()
			if errDel != nil {
				t.Logf("Failed to delete Config Match List during cleanup: %v", errDel)
			}
		}()
	}

	// Verify the request was successful.
	handleAPIError(err)
	require.NoError(t, err, "Create request should not return an error")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "The response from create should not be nil")
	assert.Equal(t, matchListName, res.Name, "The name of the created match list should match")
	assert.NotEmpty(t, *res.Id, "The ID of the created match list should not be empty")

	t.Logf("Successfully created Config Match List with ID: %s", *res.Id)
}

// Test_networkservices_ConfigMatchListAPIService_GetByID tests the retrieval of a Config Match List by its ID.
func Test_networkservices_ConfigMatchListAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	folder := "ngfw-shared"

	// Create a simple match list to retrieve (no dependencies needed for basic test)
	matchListName := "test-config-list-" + common.GenerateRandomString(10)

	matchList := network_services.ConfigMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("Config match list for get by ID test"),
		Folder:         common.StringPtr(folder),
		Filter:         common.StringPtr("All Logs"),
		SendToPanorama: common.BoolPtr(false),
	}

	createRes, _, err := client.ConfigMatchListAPI.CreateConfigMatchList(context.Background()).ConfigMatchList(matchList).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create match list for get test")
	createdMatchListID := *createRes.Id

	// Defer cleanup for the Config Match List.
	defer func() {
		t.Logf("Cleaning up Config Match List with ID: %s", createdMatchListID)
		_, errDel := client.ConfigMatchListAPI.DeleteConfigMatchListByID(context.Background(), createdMatchListID).Execute()
		if errDel != nil {
			t.Logf("Failed to delete match list during cleanup: %v", errDel)
		}
	}()

	t.Logf("Created Config Match List for Get test with ID: %s", createdMatchListID)

	// Test the Get by ID operation.
	fmt.Printf("Attempting to get Config Match List with ID: %s\n", createdMatchListID)
	req := client.ConfigMatchListAPI.GetConfigMatchListByID(context.Background(), createdMatchListID)
	getRes, httpRes, err := req.Execute()

	// Verify the retrieval was successful.
	handleAPIError(err)
	require.NoError(t, err, "Get by ID request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "The response from get should not be nil")
	assert.Equal(t, createdMatchListID, *getRes.Id, "The ID of the retrieved match list should match")
	assert.Equal(t, matchListName, getRes.Name, "The name of the retrieved match list should match")
}

// Test_networkservices_ConfigMatchListAPIService_Update tests updating a Config Match List.
func Test_networkservices_ConfigMatchListAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	folder := "ngfw-shared"

	// Create a simple match list to update
	matchListName := "test-config-list-" + common.GenerateRandomString(10)

	matchList := network_services.ConfigMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("Config match list for update test"),
		Folder:         common.StringPtr(folder),
		Filter:         common.StringPtr("All Logs"),
		SendToPanorama: common.BoolPtr(false),
	}

	createRes, _, err := client.ConfigMatchListAPI.CreateConfigMatchList(context.Background()).ConfigMatchList(matchList).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create match list for update test")
	createdMatchListID := *createRes.Id

	// Defer cleanup for the Config Match List.
	defer func() {
		t.Logf("Cleaning up Config Match List with ID: %s", createdMatchListID)
		_, errDel := client.ConfigMatchListAPI.DeleteConfigMatchListByID(context.Background(), createdMatchListID).Execute()
		if errDel != nil {
			t.Logf("Failed to delete match list during cleanup: %v", errDel)
		}
	}()

	t.Logf("Created Config Match List for Update test with ID: %s", createdMatchListID)

	// Update the match list object.
	updatedMatchList := network_services.ConfigMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("Updated description for Config match list"),
		Folder:         common.StringPtr(folder),
		Filter:         common.StringPtr("All Logs"),
		SendToPanorama: common.BoolPtr(false),
	}

	// Test the Update by ID operation.
	fmt.Printf("Attempting to update Config Match List with ID: %s\n", createdMatchListID)
	reqUpdate := client.ConfigMatchListAPI.UpdateConfigMatchListByID(context.Background(), createdMatchListID).ConfigMatchList(updatedMatchList)
	updateRes, httpRes, err := reqUpdate.Execute()

	// Verify the update was successful.
	handleAPIError(err)
	require.NoError(t, err, "Update request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "The response from update should not be nil")
	assert.Equal(t, common.StringPtr("Updated description for Config match list"), updateRes.Description, "Description should be updated")
}

// Test_networkservices_ConfigMatchListAPIService_List tests listing Config Match Lists.
func Test_networkservices_ConfigMatchListAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Read-only test: list existing objects (no Create needed)
	listRes, httpResList, errList := client.ConfigMatchListAPI.ListConfigMatchList(context.Background()).Folder("ngfw-shared").Limit(200).Offset(0).Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list Config match lists")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed Config match lists")
}

// Test_networkservices_ConfigMatchListAPIService_Fetch tests the fetch convenience method.
func Test_networkservices_ConfigMatchListAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Read-only test: Fetch non-existent object (should return nil, nil)
	notFound, err := client.ConfigMatchListAPI.FetchConfigMatchList(
		context.Background(),
		"non-existent-config-match-list-xyz-12345",
		common.StringPtr("ngfw-shared"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchConfigMatchList correctly returned nil for non-existent object")
}

// Test_networkservices_ConfigMatchListAPIService_DeleteByID tests deleting a Config Match List.
func Test_networkservices_ConfigMatchListAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	folder := "ngfw-shared"

	// Create a simple match list to delete
	matchListName := "test-config-list-" + common.GenerateRandomString(10)

	matchList := network_services.ConfigMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("Config match list for delete test"),
		Folder:         common.StringPtr(folder),
		Filter:         common.StringPtr("All Logs"),
		SendToPanorama: common.BoolPtr(false),
	}

	createRes, _, err := client.ConfigMatchListAPI.CreateConfigMatchList(context.Background()).ConfigMatchList(matchList).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create match list for delete test")
	createdMatchListID := *createRes.Id

	t.Logf("Created Config Match List for Delete test with ID: %s", createdMatchListID)

	// Test the Delete by ID operation.
	fmt.Printf("Attempting to delete Config Match List with ID: %s\n", createdMatchListID)
	reqDel := client.ConfigMatchListAPI.DeleteConfigMatchListByID(context.Background(), createdMatchListID)
	httpResDel, errDel := reqDel.Execute()

	// Verify the delete operation was successful.
	handleAPIError(errDel)
	require.NoError(t, errDel, "Failed to delete match list")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
	t.Logf("Successfully deleted Config Match List: %s", createdMatchListID)
}
