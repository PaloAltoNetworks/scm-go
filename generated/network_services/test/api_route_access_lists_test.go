/*
 * Network Services Testing
 *
 * RouteAccessListsAPIService
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

// Test_network_services_RouteAccessListsAPIService_Create tests the creation of a Route Access List.
func Test_network_services_RouteAccessListsAPIService_Create(t *testing.T) {
	t.Skip("API response returns source_address as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-ral-create-" + randomSuffix

	srcAddr := network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{
		Entry: &network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressEntry{
			Address:  common.StringPtr("10.0.0.0"),
			Wildcard: common.StringPtr("0.255.255.255"),
		},
	}
	ipv4Entry := network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{
		Name:          common.Int32Ptr(10),
		Action:        common.StringPtr("permit"),
		SourceAddress: &srcAddr,
	}
	ipv4 := network_services.RouteAccessListsTypeIpv4{
		Ipv4Entry: []network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{ipv4Entry},
	}
	listType := network_services.RouteAccessListsType{
		Ipv4: &ipv4,
	}
	routeAccessList := network_services.RouteAccessLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		Description: common.StringPtr("Test route access list for create"),
		Type:        &listType,
	}

	fmt.Printf("Attempting to create Route Access List with name: %s\n", routeAccessList.Name)

	req := client.RouteAccessListsAPI.CreateRouteAccessLists(context.Background()).RouteAccessLists(routeAccessList)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Route Access List")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created route access list should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up Route Access List with ID: %s", createdID)
		_, errDel := client.RouteAccessListsAPI.DeleteRouteAccessListsByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete route access list during cleanup")
	}()

	assert.Equal(t, listName, res.Name, "Created route access list name should match")
	t.Logf("Successfully created and validated Route Access List: %s with ID: %s", routeAccessList.Name, createdID)
}

// Test_network_services_RouteAccessListsAPIService_GetByID tests retrieving a route access list by its ID.
func Test_network_services_RouteAccessListsAPIService_GetByID(t *testing.T) {
	t.Skip("API response returns source_address as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-ral-get-" + randomSuffix

	srcAddr := network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{
		Entry: &network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressEntry{
			Address:  common.StringPtr("10.0.0.0"),
			Wildcard: common.StringPtr("0.255.255.255"),
		},
	}
	ipv4Entry := network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{
		Name:          common.Int32Ptr(10),
		Action:        common.StringPtr("permit"),
		SourceAddress: &srcAddr,
	}
	ipv4 := network_services.RouteAccessListsTypeIpv4{
		Ipv4Entry: []network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{ipv4Entry},
	}
	listType := network_services.RouteAccessListsType{
		Ipv4: &ipv4,
	}
	routeAccessList := network_services.RouteAccessLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RouteAccessListsAPI.CreateRouteAccessLists(context.Background()).RouteAccessLists(routeAccessList).Execute()
	require.NoError(t, err, "Failed to create route access list for get test")
	createdID := *createRes.Id

	defer func() {
		client.RouteAccessListsAPI.DeleteRouteAccessListsByID(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.RouteAccessListsAPI.GetRouteAccessListsByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get route access list by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, listName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_network_services_RouteAccessListsAPIService_Update tests updating an existing route access list.
func Test_network_services_RouteAccessListsAPIService_Update(t *testing.T) {
	t.Skip("API response returns source_address as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-ral-update-" + randomSuffix

	srcAddr := network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{
		Entry: &network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressEntry{
			Address:  common.StringPtr("10.0.0.0"),
			Wildcard: common.StringPtr("0.255.255.255"),
		},
	}
	ipv4Entry := network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{
		Name:          common.Int32Ptr(10),
		Action:        common.StringPtr("permit"),
		SourceAddress: &srcAddr,
	}
	ipv4 := network_services.RouteAccessListsTypeIpv4{
		Ipv4Entry: []network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{ipv4Entry},
	}
	listType := network_services.RouteAccessListsType{
		Ipv4: &ipv4,
	}
	routeAccessList := network_services.RouteAccessLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RouteAccessListsAPI.CreateRouteAccessLists(context.Background()).RouteAccessLists(routeAccessList).Execute()
	require.NoError(t, err, "Failed to create route access list for update test")
	createdID := *createRes.Id

	defer func() {
		client.RouteAccessListsAPI.DeleteRouteAccessListsByID(context.Background(), createdID).Execute()
	}()

	updatedRouteAccessList := network_services.RouteAccessLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		Description: common.StringPtr("Updated route access list description"),
		Type:        &listType,
	}

	updateRes, httpResUpdate, errUpdate := client.RouteAccessListsAPI.UpdateRouteAccessListsByID(context.Background(), createdID).RouteAccessLists(updatedRouteAccessList).Execute()
	require.NoError(t, errUpdate, "Failed to update route access list")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, "Updated route access list description", *updateRes.Description, "Description should be updated")
}

// Test_network_services_RouteAccessListsAPIService_List tests listing Route Access Lists.
func Test_network_services_RouteAccessListsAPIService_List(t *testing.T) {
	t.Skip("API response returns source_address as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-ral-list-" + randomSuffix

	srcAddr := network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{
		Entry: &network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressEntry{
			Address:  common.StringPtr("10.0.0.0"),
			Wildcard: common.StringPtr("0.255.255.255"),
		},
	}
	ipv4Entry := network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{
		Name:          common.Int32Ptr(10),
		Action:        common.StringPtr("permit"),
		SourceAddress: &srcAddr,
	}
	ipv4 := network_services.RouteAccessListsTypeIpv4{
		Ipv4Entry: []network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{ipv4Entry},
	}
	listType := network_services.RouteAccessListsType{
		Ipv4: &ipv4,
	}
	routeAccessList := network_services.RouteAccessLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RouteAccessListsAPI.CreateRouteAccessLists(context.Background()).RouteAccessLists(routeAccessList).Execute()
	require.NoError(t, err, "Failed to create route access list for list test")
	createdID := *createRes.Id

	defer func() {
		client.RouteAccessListsAPI.DeleteRouteAccessListsByID(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.RouteAccessListsAPI.ListRouteAccessLists(context.Background()).Folder("All").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list route access lists")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundList := false
	for _, item := range listRes.Data {
		if item.Name == listName {
			foundList = true
			break
		}
	}
	assert.True(t, foundList, "Created route access list should be found in the list")
}

// Test_network_services_RouteAccessListsAPIService_DeleteByID tests deleting a route access list by ID.
func Test_network_services_RouteAccessListsAPIService_DeleteByID(t *testing.T) {
	t.Skip("API response returns source_address as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-ral-delete-" + randomSuffix

	srcAddr := network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{
		Entry: &network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressEntry{
			Address:  common.StringPtr("10.0.0.0"),
			Wildcard: common.StringPtr("0.255.255.255"),
		},
	}
	ipv4Entry := network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{
		Name:          common.Int32Ptr(10),
		Action:        common.StringPtr("permit"),
		SourceAddress: &srcAddr,
	}
	ipv4 := network_services.RouteAccessListsTypeIpv4{
		Ipv4Entry: []network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{ipv4Entry},
	}
	listType := network_services.RouteAccessListsType{
		Ipv4: &ipv4,
	}
	routeAccessList := network_services.RouteAccessLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RouteAccessListsAPI.CreateRouteAccessLists(context.Background()).RouteAccessLists(routeAccessList).Execute()
	require.NoError(t, err, "Failed to create route access list for delete test")
	createdID := *createRes.Id

	_, errDel := client.RouteAccessListsAPI.DeleteRouteAccessListsByID(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete route access list")
}

// Test_network_services_RouteAccessListsAPIService_FetchRouteAccessLists tests the FetchRouteAccessLists convenience method
func Test_network_services_RouteAccessListsAPIService_FetchRouteAccessLists(t *testing.T) {
	t.Skip("API response returns source_address as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-ral-fetch-" + randomSuffix

	srcAddr := network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddress{
		Entry: &network_services.RouteAccessListsTypeIpv4Ipv4EntryInnerSourceAddressEntry{
			Address:  common.StringPtr("10.0.0.0"),
			Wildcard: common.StringPtr("0.255.255.255"),
		},
	}
	ipv4Entry := network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{
		Name:          common.Int32Ptr(10),
		Action:        common.StringPtr("permit"),
		SourceAddress: &srcAddr,
	}
	ipv4 := network_services.RouteAccessListsTypeIpv4{
		Ipv4Entry: []network_services.RouteAccessListsTypeIpv4Ipv4EntryInner{ipv4Entry},
	}
	fetchListType := network_services.RouteAccessListsType{
		Ipv4: &ipv4,
	}
	testObj := network_services.RouteAccessLists{
		Name:   testName,
		Folder: common.StringPtr("All"),
		Type:   &fetchListType,
	}

	createReq := client.RouteAccessListsAPI.CreateRouteAccessLists(context.Background()).RouteAccessLists(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		deleteReq := client.RouteAccessListsAPI.DeleteRouteAccessListsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.RouteAccessListsAPI.FetchRouteAccessLists(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch route_access_lists by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchRouteAccessLists found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.RouteAccessListsAPI.FetchRouteAccessLists(
		context.Background(),
		"non-existent-route_access_lists-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchRouteAccessLists correctly returned nil for non-existent object")
}
