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

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Test_networkservices_ConfigMatchListAPIService_Create tests the creation of a Config Match List.
func Test_networkservices_ConfigMatchListAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupNetworkSvcTestClient(t)

	// Create a valid Config Match List object with a unique name.
	matchListName := "test-config-list-" + common.GenerateRandomString(10)

	matchList := network_services.ConfigMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("Config match list for tracking configuration changes and audit log forwarding"),
		Folder:         common.StringPtr("ngfw-shared"),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{"test-syslog"},
		SendHttp:       []string{"some-http-profile"},
		SendSnmptrap:   []string{"snmp_test"},
		SendEmail:      []string{"test-email"},
		SendToPanorama: common.BoolPtr(false),
	}

	fmt.Printf("Attempting to create Config Match List with name: %s\n", matchList.Name)

	// Make the create request to the API.
	req := client.ConfigMatchListAPI.CreateConfigMatchList(context.Background()).ConfigMatchList(matchList)
	res, httpRes, err := req.Execute()

	// Defer cleanup for the Config Match List.
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

	// Create a match list to retrieve.
	matchListName := "test-config-list-" + common.GenerateRandomString(10)

	matchList := network_services.ConfigMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("Config match list for get by ID test"),
		Folder:         common.StringPtr("ngfw-shared"),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{"test-syslog"},
		SendHttp:       []string{"some-http-profile"},
		SendSnmptrap:   []string{"snmp_test"},
		SendEmail:      []string{"test-email"},
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

	// Create a match list to update.
	matchListName := "test-config-list-" + common.GenerateRandomString(10)

	matchList := network_services.ConfigMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("Config match list for update test"),
		Folder:         common.StringPtr("ngfw-shared"),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{"test-syslog"},
		SendHttp:       []string{"some-http-profile"},
		SendSnmptrap:   []string{"snmp_test"},
		SendEmail:      []string{"test-email"},
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
		Folder:         common.StringPtr("ngfw-shared"),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{"test-syslog"},
		SendHttp:       []string{"some-http-profile"},
		SendSnmptrap:   []string{"snmp_test"},
		SendEmail:      []string{"test-email"},
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

	// Create a match list to delete.
	matchListName := "test-config-list-" + common.GenerateRandomString(10)

	matchList := network_services.ConfigMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("Config match list for delete test"),
		Folder:         common.StringPtr("ngfw-shared"),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{"test-syslog"},
		SendHttp:       []string{"some-http-profile"},
		SendSnmptrap:   []string{"snmp_test"},
		SendEmail:      []string{"test-email"},
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
