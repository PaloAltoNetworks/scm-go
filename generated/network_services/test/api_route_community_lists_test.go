/*
 * Network Services Testing
 *
 * RouteCommunityListsAPIService
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

// Test_network_services_RouteCommunityListsAPIService_Create tests the creation of a Route Community List.
func Test_network_services_RouteCommunityListsAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rcl-create-" + randomSuffix

	entry := network_services.RouteCommunityListsTypeRegularRegularEntryInner{
		Name:      common.Int32Ptr(10),
		Action:    common.StringPtr("permit"),
		Community: []string{"65001:100"},
	}
	regular := network_services.RouteCommunityListsTypeRegular{
		RegularEntry: []network_services.RouteCommunityListsTypeRegularRegularEntryInner{entry},
	}
	listType := network_services.RouteCommunityListsType{
		Regular: &regular,
	}
	routeCommunityList := network_services.RouteCommunityLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		Description: common.StringPtr("Test route community list for create"),
		Type:        &listType,
	}

	fmt.Printf("Attempting to create Route Community List with name: %s\n", routeCommunityList.Name)

	req := client.RouteCommunityListsAPI.CreateRouteCommunityLists(context.Background()).RouteCommunityLists(routeCommunityList)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Route Community List")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created route community list should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up Route Community List with ID: %s", createdID)
		_, errDel := client.RouteCommunityListsAPI.DeleteRouteCommunityListsByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete route community list during cleanup")
	}()

	assert.Equal(t, listName, res.Name, "Created route community list name should match")
	t.Logf("Successfully created and validated Route Community List: %s with ID: %s", routeCommunityList.Name, createdID)
}

// Test_network_services_RouteCommunityListsAPIService_GetByID tests retrieving a route community list by its ID.
func Test_network_services_RouteCommunityListsAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rcl-get-" + randomSuffix

	entry := network_services.RouteCommunityListsTypeRegularRegularEntryInner{
		Name:      common.Int32Ptr(10),
		Action:    common.StringPtr("permit"),
		Community: []string{"65002:100"},
	}
	regular := network_services.RouteCommunityListsTypeRegular{
		RegularEntry: []network_services.RouteCommunityListsTypeRegularRegularEntryInner{entry},
	}
	listType := network_services.RouteCommunityListsType{
		Regular: &regular,
	}
	routeCommunityList := network_services.RouteCommunityLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RouteCommunityListsAPI.CreateRouteCommunityLists(context.Background()).RouteCommunityLists(routeCommunityList).Execute()
	require.NoError(t, err, "Failed to create route community list for get test")
	createdID := *createRes.Id

	defer func() {
		client.RouteCommunityListsAPI.DeleteRouteCommunityListsByID(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.RouteCommunityListsAPI.GetRouteCommunityListsByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get route community list by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, listName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_network_services_RouteCommunityListsAPIService_Update tests updating an existing route community list.
func Test_network_services_RouteCommunityListsAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rcl-update-" + randomSuffix

	entry := network_services.RouteCommunityListsTypeRegularRegularEntryInner{
		Name:      common.Int32Ptr(10),
		Action:    common.StringPtr("permit"),
		Community: []string{"65003:100"},
	}
	regular := network_services.RouteCommunityListsTypeRegular{
		RegularEntry: []network_services.RouteCommunityListsTypeRegularRegularEntryInner{entry},
	}
	listType := network_services.RouteCommunityListsType{
		Regular: &regular,
	}
	routeCommunityList := network_services.RouteCommunityLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RouteCommunityListsAPI.CreateRouteCommunityLists(context.Background()).RouteCommunityLists(routeCommunityList).Execute()
	require.NoError(t, err, "Failed to create route community list for update test")
	createdID := *createRes.Id

	defer func() {
		client.RouteCommunityListsAPI.DeleteRouteCommunityListsByID(context.Background(), createdID).Execute()
	}()

	updatedRouteCommunityList := network_services.RouteCommunityLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		Description: common.StringPtr("Updated route community list description"),
		Type:        &listType,
	}

	updateRes, httpResUpdate, errUpdate := client.RouteCommunityListsAPI.UpdateRouteCommunityListsByID(context.Background(), createdID).RouteCommunityLists(updatedRouteCommunityList).Execute()
	require.NoError(t, errUpdate, "Failed to update route community list")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, "Updated route community list description", *updateRes.Description, "Description should be updated")
}

// Test_network_services_RouteCommunityListsAPIService_List tests listing Route Community Lists.
func Test_network_services_RouteCommunityListsAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rcl-list-" + randomSuffix

	entry := network_services.RouteCommunityListsTypeRegularRegularEntryInner{
		Name:      common.Int32Ptr(10),
		Action:    common.StringPtr("permit"),
		Community: []string{"65004:100"},
	}
	regular := network_services.RouteCommunityListsTypeRegular{
		RegularEntry: []network_services.RouteCommunityListsTypeRegularRegularEntryInner{entry},
	}
	listType := network_services.RouteCommunityListsType{
		Regular: &regular,
	}
	routeCommunityList := network_services.RouteCommunityLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RouteCommunityListsAPI.CreateRouteCommunityLists(context.Background()).RouteCommunityLists(routeCommunityList).Execute()
	require.NoError(t, err, "Failed to create route community list for list test")
	createdID := *createRes.Id

	defer func() {
		client.RouteCommunityListsAPI.DeleteRouteCommunityListsByID(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.RouteCommunityListsAPI.ListRouteCommunityLists(context.Background()).Folder("All").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list route community lists")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundList := false
	for _, item := range listRes.Data {
		if item.Name == listName {
			foundList = true
			break
		}
	}
	assert.True(t, foundList, "Created route community list should be found in the list")
}

// Test_network_services_RouteCommunityListsAPIService_DeleteByID tests deleting a route community list by ID.
func Test_network_services_RouteCommunityListsAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rcl-delete-" + randomSuffix

	entry := network_services.RouteCommunityListsTypeRegularRegularEntryInner{
		Name:      common.Int32Ptr(10),
		Action:    common.StringPtr("permit"),
		Community: []string{"65005:100"},
	}
	regular := network_services.RouteCommunityListsTypeRegular{
		RegularEntry: []network_services.RouteCommunityListsTypeRegularRegularEntryInner{entry},
	}
	listType := network_services.RouteCommunityListsType{
		Regular: &regular,
	}
	routeCommunityList := network_services.RouteCommunityLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RouteCommunityListsAPI.CreateRouteCommunityLists(context.Background()).RouteCommunityLists(routeCommunityList).Execute()
	require.NoError(t, err, "Failed to create route community list for delete test")
	createdID := *createRes.Id

	_, errDel := client.RouteCommunityListsAPI.DeleteRouteCommunityListsByID(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete route community list")
}

// Test_network_services_RouteCommunityListsAPIService_FetchRouteCommunityLists tests the FetchRouteCommunityLists convenience method
func Test_network_services_RouteCommunityListsAPIService_FetchRouteCommunityLists(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-rcl-fetch-" + randomSuffix

	entry := network_services.RouteCommunityListsTypeRegularRegularEntryInner{
		Name:      common.Int32Ptr(10),
		Action:    common.StringPtr("permit"),
		Community: []string{"65006:100"},
	}
	regular := network_services.RouteCommunityListsTypeRegular{
		RegularEntry: []network_services.RouteCommunityListsTypeRegularRegularEntryInner{entry},
	}
	listType := network_services.RouteCommunityListsType{
		Regular: &regular,
	}
	testObj := network_services.RouteCommunityLists{
		Name:   testName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createReq := client.RouteCommunityListsAPI.CreateRouteCommunityLists(context.Background()).RouteCommunityLists(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		deleteReq := client.RouteCommunityListsAPI.DeleteRouteCommunityListsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.RouteCommunityListsAPI.FetchRouteCommunityLists(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch route_community_lists by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchRouteCommunityLists found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.RouteCommunityListsAPI.FetchRouteCommunityLists(
		context.Background(),
		"non-existent-route_community_lists-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchRouteCommunityLists correctly returned nil for non-existent object")
}
