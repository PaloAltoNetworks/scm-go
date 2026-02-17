/*
 * Network Services Testing
 *
 * RoutePrefixListsAPIService
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

// Test_network_services_RoutePrefixListsAPIService_Create tests the creation of a Route Prefix List.
func Test_network_services_RoutePrefixListsAPIService_Create(t *testing.T) {
	t.Skip("API response returns prefix as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rpl-create-" + randomSuffix

	prefixEntry := network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{
		Name:   common.Int32Ptr(10),
		Action: common.StringPtr("permit"),
		Prefix: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{
			Entry: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{
				Network: common.StringPtr("10.0.0.0/8"),
			},
		},
	}
	listType := network_services.RoutePrefixListsType{
		Ipv4: network_services.RoutePrefixListsTypeIpv4{
			Ipv4Entry: []network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{prefixEntry},
		},
	}
	routePrefixList := network_services.RoutePrefixLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		Description: common.StringPtr("Test route prefix list for create"),
		Type:        &listType,
	}

	fmt.Printf("Attempting to create Route Prefix List with name: %s\n", routePrefixList.Name)

	req := client.RoutePrefixListsAPI.CreateRoutePrefixLists(context.Background()).RoutePrefixLists(routePrefixList)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Route Prefix List")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created route prefix list should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up Route Prefix List with ID: %s", createdID)
		_, errDel := client.RoutePrefixListsAPI.DeleteRoutePrefixListsByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete route prefix list during cleanup")
	}()

	assert.Equal(t, listName, res.Name, "Created route prefix list name should match")
	t.Logf("Successfully created and validated Route Prefix List: %s with ID: %s", routePrefixList.Name, createdID)
}

// Test_network_services_RoutePrefixListsAPIService_GetByID tests retrieving a route prefix list by its ID.
func Test_network_services_RoutePrefixListsAPIService_GetByID(t *testing.T) {
	t.Skip("API response returns prefix as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rpl-get-" + randomSuffix

	prefixEntry := network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{
		Name:   common.Int32Ptr(10),
		Action: common.StringPtr("permit"),
		Prefix: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{
			Entry: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{
				Network: common.StringPtr("10.0.0.0/8"),
			},
		},
	}
	listType := network_services.RoutePrefixListsType{
		Ipv4: network_services.RoutePrefixListsTypeIpv4{
			Ipv4Entry: []network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{prefixEntry},
		},
	}
	routePrefixList := network_services.RoutePrefixLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RoutePrefixListsAPI.CreateRoutePrefixLists(context.Background()).RoutePrefixLists(routePrefixList).Execute()
	require.NoError(t, err, "Failed to create route prefix list for get test")
	createdID := *createRes.Id

	defer func() {
		client.RoutePrefixListsAPI.DeleteRoutePrefixListsByID(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.RoutePrefixListsAPI.GetRoutePrefixListsByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get route prefix list by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, listName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_network_services_RoutePrefixListsAPIService_Update tests updating an existing route prefix list.
func Test_network_services_RoutePrefixListsAPIService_Update(t *testing.T) {
	t.Skip("API response returns prefix as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rpl-update-" + randomSuffix

	prefixEntry := network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{
		Name:   common.Int32Ptr(10),
		Action: common.StringPtr("permit"),
		Prefix: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{
			Entry: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{
				Network: common.StringPtr("10.0.0.0/8"),
			},
		},
	}
	listType := network_services.RoutePrefixListsType{
		Ipv4: network_services.RoutePrefixListsTypeIpv4{
			Ipv4Entry: []network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{prefixEntry},
		},
	}
	routePrefixList := network_services.RoutePrefixLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RoutePrefixListsAPI.CreateRoutePrefixLists(context.Background()).RoutePrefixLists(routePrefixList).Execute()
	require.NoError(t, err, "Failed to create route prefix list for update test")
	createdID := *createRes.Id

	defer func() {
		client.RoutePrefixListsAPI.DeleteRoutePrefixListsByID(context.Background(), createdID).Execute()
	}()

	updatedRoutePrefixList := network_services.RoutePrefixLists{
		Name:        listName,
		Folder:      common.StringPtr("All"),
		Description: common.StringPtr("Updated route prefix list description"),
		Type:        &listType,
	}

	updateRes, httpResUpdate, errUpdate := client.RoutePrefixListsAPI.UpdateRoutePrefixListsByID(context.Background(), createdID).RoutePrefixLists(updatedRoutePrefixList).Execute()
	require.NoError(t, errUpdate, "Failed to update route prefix list")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, "Updated route prefix list description", *updateRes.Description, "Description should be updated")
}

// Test_network_services_RoutePrefixListsAPIService_List tests listing Route Prefix Lists.
func Test_network_services_RoutePrefixListsAPIService_List(t *testing.T) {
	t.Skip("API response returns prefix as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rpl-list-" + randomSuffix

	prefixEntry := network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{
		Name:   common.Int32Ptr(10),
		Action: common.StringPtr("permit"),
		Prefix: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{
			Entry: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{
				Network: common.StringPtr("10.0.0.0/8"),
			},
		},
	}
	listType := network_services.RoutePrefixListsType{
		Ipv4: network_services.RoutePrefixListsTypeIpv4{
			Ipv4Entry: []network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{prefixEntry},
		},
	}
	routePrefixList := network_services.RoutePrefixLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RoutePrefixListsAPI.CreateRoutePrefixLists(context.Background()).RoutePrefixLists(routePrefixList).Execute()
	require.NoError(t, err, "Failed to create route prefix list for list test")
	createdID := *createRes.Id

	defer func() {
		client.RoutePrefixListsAPI.DeleteRoutePrefixListsByID(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.RoutePrefixListsAPI.ListRoutePrefixLists(context.Background()).Folder("All").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list route prefix lists")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundList := false
	for _, item := range listRes.Data {
		if item.Name == listName {
			foundList = true
			break
		}
	}
	assert.True(t, foundList, "Created route prefix list should be found in the list")
}

// Test_network_services_RoutePrefixListsAPIService_DeleteByID tests deleting a route prefix list by ID.
func Test_network_services_RoutePrefixListsAPIService_DeleteByID(t *testing.T) {
	t.Skip("API response returns prefix as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	listName := "test-rpl-delete-" + randomSuffix

	prefixEntry := network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{
		Name:   common.Int32Ptr(10),
		Action: common.StringPtr("permit"),
		Prefix: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{
			Entry: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{
				Network: common.StringPtr("10.0.0.0/8"),
			},
		},
	}
	listType := network_services.RoutePrefixListsType{
		Ipv4: network_services.RoutePrefixListsTypeIpv4{
			Ipv4Entry: []network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{prefixEntry},
		},
	}
	routePrefixList := network_services.RoutePrefixLists{
		Name:   listName,
		Folder: common.StringPtr("All"),
		Type:   &listType,
	}

	createRes, _, err := client.RoutePrefixListsAPI.CreateRoutePrefixLists(context.Background()).RoutePrefixLists(routePrefixList).Execute()
	require.NoError(t, err, "Failed to create route prefix list for delete test")
	createdID := *createRes.Id

	_, errDel := client.RoutePrefixListsAPI.DeleteRoutePrefixListsByID(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete route prefix list")
}

// Test_network_services_RoutePrefixListsAPIService_FetchRoutePrefixLists tests the FetchRoutePrefixLists convenience method
func Test_network_services_RoutePrefixListsAPIService_FetchRoutePrefixLists(t *testing.T) {
	t.Skip("API response returns prefix as array but model expects object - model deserialization error")
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-rpl-fetch-" + randomSuffix

	prefixEntry := network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{
		Name:   common.Int32Ptr(10),
		Action: common.StringPtr("permit"),
		Prefix: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefix{
			Entry: &network_services.RoutePrefixListsTypeIpv4Ipv4EntryInnerPrefixEntry{
				Network: common.StringPtr("10.0.0.0/8"),
			},
		},
	}
	fetchListType := network_services.RoutePrefixListsType{
		Ipv4: network_services.RoutePrefixListsTypeIpv4{
			Ipv4Entry: []network_services.RoutePrefixListsTypeIpv4Ipv4EntryInner{prefixEntry},
		},
	}
	testObj := network_services.RoutePrefixLists{
		Name:   testName,
		Folder: common.StringPtr("All"),
		Type:   &fetchListType,
	}

	createReq := client.RoutePrefixListsAPI.CreateRoutePrefixLists(context.Background()).RoutePrefixLists(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		deleteReq := client.RoutePrefixListsAPI.DeleteRoutePrefixListsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.RoutePrefixListsAPI.FetchRoutePrefixLists(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch route_prefix_lists by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchRoutePrefixLists found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.RoutePrefixListsAPI.FetchRoutePrefixLists(
		context.Background(),
		"non-existent-route_prefix_lists-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchRoutePrefixLists correctly returned nil for non-existent object")
}
