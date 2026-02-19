/*
 * Network Services Testing
 *
 * RoutePathAccessListsAPIService
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

// Test_network_services_RoutePathAccessListsAPIService_Create tests the creation of a Route Path Access List.
func Test_network_services_RoutePathAccessListsAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rpal-create-" + randomSuffix

	aspathEntry := network_services.RoutePathAccessListsAspathEntryInner{
		Name:        common.Int32Ptr(10),
		Action:      common.StringPtr("permit"),
		AspathRegex: common.StringPtr("^65001_"),
	}
	routePathAccessList := network_services.RoutePathAccessLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		Description: common.StringPtr("Test route path access list for create"),
		AspathEntry: []network_services.RoutePathAccessListsAspathEntryInner{aspathEntry},
	}

	fmt.Printf("Attempting to create Route Path Access List with name: %s\n", routePathAccessList.Name)

	req := client.RoutePathAccessListsAPI.CreateRoutePathAccessLists(context.Background()).RoutePathAccessLists(routePathAccessList)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Route Path Access List")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created route path access list should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up Route Path Access List with ID: %s", createdID)
		_, errDel := client.RoutePathAccessListsAPI.DeleteRoutePathAccessListsByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete route path access list during cleanup")
	}()

	assert.Equal(t, listName, res.Name, "Created route path access list name should match")
	t.Logf("Successfully created and validated Route Path Access List: %s with ID: %s", routePathAccessList.Name, createdID)
}

// Test_network_services_RoutePathAccessListsAPIService_GetByID tests retrieving a route path access list by its ID.
func Test_network_services_RoutePathAccessListsAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rpal-get-" + randomSuffix

	aspathEntry := network_services.RoutePathAccessListsAspathEntryInner{
		Name:        common.Int32Ptr(10),
		Action:      common.StringPtr("permit"),
		AspathRegex: common.StringPtr("^65002_"),
	}
	routePathAccessList := network_services.RoutePathAccessLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		AspathEntry: []network_services.RoutePathAccessListsAspathEntryInner{aspathEntry},
	}

	createRes, _, err := client.RoutePathAccessListsAPI.CreateRoutePathAccessLists(context.Background()).RoutePathAccessLists(routePathAccessList).Execute()
	require.NoError(t, err, "Failed to create route path access list for get test")
	createdID := *createRes.Id

	defer func() {
		client.RoutePathAccessListsAPI.DeleteRoutePathAccessListsByID(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.RoutePathAccessListsAPI.GetRoutePathAccessListsByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get route path access list by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, listName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_network_services_RoutePathAccessListsAPIService_Update tests updating an existing route path access list.
func Test_network_services_RoutePathAccessListsAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rpal-update-" + randomSuffix

	aspathEntry := network_services.RoutePathAccessListsAspathEntryInner{
		Name:        common.Int32Ptr(10),
		Action:      common.StringPtr("permit"),
		AspathRegex: common.StringPtr("^65003_"),
	}
	routePathAccessList := network_services.RoutePathAccessLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		AspathEntry: []network_services.RoutePathAccessListsAspathEntryInner{aspathEntry},
	}

	createRes, _, err := client.RoutePathAccessListsAPI.CreateRoutePathAccessLists(context.Background()).RoutePathAccessLists(routePathAccessList).Execute()
	require.NoError(t, err, "Failed to create route path access list for update test")
	createdID := *createRes.Id

	defer func() {
		client.RoutePathAccessListsAPI.DeleteRoutePathAccessListsByID(context.Background(), createdID).Execute()
	}()

	updatedRoutePathAccessList := network_services.RoutePathAccessLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		Description: common.StringPtr("Updated route path access list description"),
		AspathEntry: []network_services.RoutePathAccessListsAspathEntryInner{aspathEntry},
	}

	updateRes, httpResUpdate, errUpdate := client.RoutePathAccessListsAPI.UpdateRoutePathAccessListsByID(context.Background(), createdID).RoutePathAccessLists(updatedRoutePathAccessList).Execute()
	require.NoError(t, errUpdate, "Failed to update route path access list")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, "Updated route path access list description", *updateRes.Description, "Description should be updated")
}

// Test_network_services_RoutePathAccessListsAPIService_List tests listing Route Path Access Lists.
func Test_network_services_RoutePathAccessListsAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rpal-list-" + randomSuffix

	aspathEntry := network_services.RoutePathAccessListsAspathEntryInner{
		Name:        common.Int32Ptr(10),
		Action:      common.StringPtr("permit"),
		AspathRegex: common.StringPtr("^65004_"),
	}
	routePathAccessList := network_services.RoutePathAccessLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		AspathEntry: []network_services.RoutePathAccessListsAspathEntryInner{aspathEntry},
	}

	createRes, _, err := client.RoutePathAccessListsAPI.CreateRoutePathAccessLists(context.Background()).RoutePathAccessLists(routePathAccessList).Execute()
	require.NoError(t, err, "Failed to create route path access list for list test")
	createdID := *createRes.Id

	defer func() {
		client.RoutePathAccessListsAPI.DeleteRoutePathAccessListsByID(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.RoutePathAccessListsAPI.ListRoutePathAccessLists(context.Background()).Folder("All").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list route path access lists")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundList := false
	for _, item := range listRes.Data {
		if item.Name == listName {
			foundList = true
			break
		}
	}
	assert.True(t, foundList, "Created route path access list should be found in the list")
}

// Test_network_services_RoutePathAccessListsAPIService_DeleteByID tests deleting a route path access list by ID.
func Test_network_services_RoutePathAccessListsAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rpal-delete-" + randomSuffix

	aspathEntry := network_services.RoutePathAccessListsAspathEntryInner{
		Name:        common.Int32Ptr(10),
		Action:      common.StringPtr("permit"),
		AspathRegex: common.StringPtr("^65005_"),
	}
	routePathAccessList := network_services.RoutePathAccessLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		AspathEntry: []network_services.RoutePathAccessListsAspathEntryInner{aspathEntry},
	}

	createRes, _, err := client.RoutePathAccessListsAPI.CreateRoutePathAccessLists(context.Background()).RoutePathAccessLists(routePathAccessList).Execute()
	require.NoError(t, err, "Failed to create route path access list for delete test")
	createdID := *createRes.Id

	_, errDel := client.RoutePathAccessListsAPI.DeleteRoutePathAccessListsByID(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete route path access list")
}

// Test_network_services_RoutePathAccessListsAPIService_FetchRoutePathAccessLists tests the FetchRoutePathAccessLists convenience method
func Test_network_services_RoutePathAccessListsAPIService_FetchRoutePathAccessLists(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-rpal-fetch-" + randomSuffix

	aspathEntry := network_services.RoutePathAccessListsAspathEntryInner{
		Name:        common.Int32Ptr(10),
		Action:      common.StringPtr("permit"),
		AspathRegex: common.StringPtr("^65006_"),
	}
	testObj := network_services.RoutePathAccessLists{
		Name:        testName,
		Folder:      common.StringPtr("All"),
		AspathEntry: []network_services.RoutePathAccessListsAspathEntryInner{aspathEntry},
	}

	createReq := client.RoutePathAccessListsAPI.CreateRoutePathAccessLists(context.Background()).RoutePathAccessLists(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		deleteReq := client.RoutePathAccessListsAPI.DeleteRoutePathAccessListsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.RoutePathAccessListsAPI.FetchRoutePathAccessLists(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch route_path_access_lists by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchRoutePathAccessLists found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.RoutePathAccessListsAPI.FetchRoutePathAccessLists(
		context.Background(),
		"non-existent-route_path_access_lists-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchRoutePathAccessLists correctly returned nil for non-existent object")
}
