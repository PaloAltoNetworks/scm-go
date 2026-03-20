/*
 * Network Services Testing
 *
 * SystemMatchListAPIService
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

// Test_networkservices_SystemMatchListAPIService_Create tests the creation of a System Match List.
func Test_networkservices_SystemMatchListAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupNetworkSvcTestClient(t)

	// Create a valid System Match List object with a unique name.
	matchListName := "test-system-list-" + common.GenerateRandomString(10)

	matchList := network_services.SystemMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("System match list for capturing system-level events and forwarding to monitoring platforms"),
		Folder:         common.StringPtr("ngfw-shared"),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{"test-syslog"},
		SendHttp:       []string{"some-http-profile"},
		SendSnmptrap:   []string{"snmp_test"},
		SendEmail:      []string{"test-email"},
		SendToPanorama: common.BoolPtr(false),
	}

	fmt.Printf("Attempting to create System Match List with name: %s\n", matchList.Name)

	// Make the create request to the API.
	req := client.SystemMatchListAPI.CreateSystemMatchList(context.Background()).SystemMatchList(matchList)
	res, httpRes, err := req.Execute()

	// Defer cleanup for the System Match List.
	if res != nil && res.Id != nil {
		defer func() {
			t.Logf("Cleaning up System Match List with ID: %s", *res.Id)
			delReq := client.SystemMatchListAPI.DeleteSystemMatchListByID(context.Background(), *res.Id)
			_, errDel := delReq.Execute()
			if errDel != nil {
				t.Logf("Failed to delete System Match List during cleanup: %v", errDel)
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

	t.Logf("Successfully created System Match List with ID: %s", *res.Id)
}

// Test_networkservices_SystemMatchListAPIService_GetByID tests the retrieval of a System Match List by its ID.
func Test_networkservices_SystemMatchListAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a match list to retrieve.
	matchListName := "test-system-list-" + common.GenerateRandomString(10)

	matchList := network_services.SystemMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("System match list for get by ID test"),
		Folder:         common.StringPtr("ngfw-shared"),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{"test-syslog"},
		SendHttp:       []string{"some-http-profile"},
		SendSnmptrap:   []string{"snmp_test"},
		SendEmail:      []string{"test-email"},
		SendToPanorama: common.BoolPtr(false),
	}

	createRes, _, err := client.SystemMatchListAPI.CreateSystemMatchList(context.Background()).SystemMatchList(matchList).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create match list for get test")
	createdMatchListID := *createRes.Id

	// Defer cleanup for the System Match List.
	defer func() {
		t.Logf("Cleaning up System Match List with ID: %s", createdMatchListID)
		_, errDel := client.SystemMatchListAPI.DeleteSystemMatchListByID(context.Background(), createdMatchListID).Execute()
		if errDel != nil {
			t.Logf("Failed to delete match list during cleanup: %v", errDel)
		}
	}()

	t.Logf("Created System Match List for Get test with ID: %s", createdMatchListID)

	// Test the Get by ID operation.
	fmt.Printf("Attempting to get System Match List with ID: %s\n", createdMatchListID)
	req := client.SystemMatchListAPI.GetSystemMatchListByID(context.Background(), createdMatchListID)
	getRes, httpRes, err := req.Execute()

	// Verify the retrieval was successful.
	handleAPIError(err)
	require.NoError(t, err, "Get by ID request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "The response from get should not be nil")
	assert.Equal(t, createdMatchListID, *getRes.Id, "The ID of the retrieved match list should match")
	assert.Equal(t, matchListName, getRes.Name, "The name of the retrieved match list should match")
}

// Test_networkservices_SystemMatchListAPIService_Update tests updating a System Match List.
func Test_networkservices_SystemMatchListAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a match list to update.
	matchListName := "test-system-list-" + common.GenerateRandomString(10)

	matchList := network_services.SystemMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("System match list for update test"),
		Folder:         common.StringPtr("ngfw-shared"),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{"test-syslog"},
		SendHttp:       []string{"some-http-profile"},
		SendSnmptrap:   []string{"snmp_test"},
		SendEmail:      []string{"test-email"},
		SendToPanorama: common.BoolPtr(false),
	}

	createRes, _, err := client.SystemMatchListAPI.CreateSystemMatchList(context.Background()).SystemMatchList(matchList).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create match list for update test")
	createdMatchListID := *createRes.Id

	// Defer cleanup for the System Match List.
	defer func() {
		t.Logf("Cleaning up System Match List with ID: %s", createdMatchListID)
		_, errDel := client.SystemMatchListAPI.DeleteSystemMatchListByID(context.Background(), createdMatchListID).Execute()
		if errDel != nil {
			t.Logf("Failed to delete match list during cleanup: %v", errDel)
		}
	}()

	t.Logf("Created System Match List for Update test with ID: %s", createdMatchListID)

	// Update the match list object.
	updatedMatchList := network_services.SystemMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("Updated description for System match list"),
		Folder:         common.StringPtr("ngfw-shared"),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{"test-syslog"},
		SendHttp:       []string{"some-http-profile"},
		SendSnmptrap:   []string{"snmp_test"},
		SendEmail:      []string{"test-email"},
		SendToPanorama: common.BoolPtr(false),
	}

	// Test the Update by ID operation.
	fmt.Printf("Attempting to update System Match List with ID: %s\n", createdMatchListID)
	reqUpdate := client.SystemMatchListAPI.UpdateSystemMatchListByID(context.Background(), createdMatchListID).SystemMatchList(updatedMatchList)
	updateRes, httpRes, err := reqUpdate.Execute()

	// Verify the update was successful.
	handleAPIError(err)
	require.NoError(t, err, "Update request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "The response from update should not be nil")
	assert.Equal(t, common.StringPtr("Updated description for System match list"), updateRes.Description, "Description should be updated")
}

// Test_networkservices_SystemMatchListAPIService_List tests listing System Match Lists.
func Test_networkservices_SystemMatchListAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Read-only test: list existing objects (no Create needed)
	listRes, httpResList, errList := client.SystemMatchListAPI.ListSystemMatchList(context.Background()).Folder("ngfw-shared").Limit(200).Offset(0).Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list System match lists")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed System match lists")
}

// Test_networkservices_SystemMatchListAPIService_Fetch tests the fetch convenience method.
func Test_networkservices_SystemMatchListAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Read-only test: Fetch non-existent object (should return nil, nil)
	notFound, err := client.SystemMatchListAPI.FetchSystemMatchList(
		context.Background(),
		"non-existent-system-match-list-xyz-12345",
		common.StringPtr("ngfw-shared"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchSystemMatchList correctly returned nil for non-existent object")
}

// Test_networkservices_SystemMatchListAPIService_DeleteByID tests deleting a System Match List.
func Test_networkservices_SystemMatchListAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a match list to delete.
	matchListName := "test-system-list-" + common.GenerateRandomString(10)

	matchList := network_services.SystemMatchList{
		Name:           matchListName,
		Description:    common.StringPtr("System match list for delete test"),
		Folder:         common.StringPtr("ngfw-shared"),
		Filter:         common.StringPtr("All Logs"),
		SendSyslog:     []string{"test-syslog"},
		SendHttp:       []string{"some-http-profile"},
		SendSnmptrap:   []string{"snmp_test"},
		SendEmail:      []string{"test-email"},
		SendToPanorama: common.BoolPtr(false),
	}

	createRes, _, err := client.SystemMatchListAPI.CreateSystemMatchList(context.Background()).SystemMatchList(matchList).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create match list for delete test")
	createdMatchListID := *createRes.Id

	t.Logf("Created System Match List for Delete test with ID: %s", createdMatchListID)

	// Test the Delete by ID operation.
	fmt.Printf("Attempting to delete System Match List with ID: %s\n", createdMatchListID)
	reqDel := client.SystemMatchListAPI.DeleteSystemMatchListByID(context.Background(), createdMatchListID)
	httpResDel, errDel := reqDel.Execute()

	// Verify the delete operation was successful.
	handleAPIError(errDel)
	require.NoError(t, errDel, "Failed to delete match list")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
	t.Logf("Successfully deleted System Match List: %s", createdMatchListID)
}
