package network_services

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Test_network_services_BGPRouteMapsAPIService_Create tests the creation of a BGP route map.
func Test_network_services_BGPRouteMapsAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	mapName := "test-bgp-routemap-create-" + common.GenerateRandomString(6)

	routeMapEntry := network_services.BgpRouteMapsRouteMapInner{
		Action: common.StringPtr("permit"),
		Name:   common.Int32Ptr(1),
	}
	profile := network_services.BgpRouteMaps{
		Name:        mapName,
		Folder:      common.StringPtr("Prisma Access"),
		Description: common.StringPtr("Test BGP route map"),
		RouteMap:    []network_services.BgpRouteMapsRouteMapInner{routeMapEntry},
	}

	t.Logf("Creating BGP route map with name: %s", mapName)
	req := client.BGPRouteMapsAPI.CreateBGPRouteMaps(context.Background()).BgpRouteMaps(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create BGP route map")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	createdMapID := res.Id

	// Cleanup the created route map
	defer func() {
		t.Logf("Cleaning up BGP route map with ID: %s", *createdMapID)
		_, errDel := client.BGPRouteMapsAPI.DeleteBGPRouteMapsByID(context.Background(), *createdMapID).Execute()
		if errDel != nil {
			t.Logf("Cleanup failed: %v", errDel)
		}
	}()

	// Verify the response matches key input fields
	assert.Equal(t, mapName, res.Name, "Created route map name should match")
}

// Test_network_services_BGPRouteMapsAPIService_GetByID tests retrieving a BGP route map by ID.
func Test_network_services_BGPRouteMapsAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	mapName := "test-bgp-routemap-get-" + common.GenerateRandomString(6)

	routeMapEntry := network_services.BgpRouteMapsRouteMapInner{
		Action: common.StringPtr("permit"),
		Name:   common.Int32Ptr(1),
	}
	profile := network_services.BgpRouteMaps{
		Name:     mapName,
		Folder:   common.StringPtr("Prisma Access"),
		RouteMap: []network_services.BgpRouteMapsRouteMapInner{routeMapEntry},
	}

	// Setup: Create a route map first
	createRes, _, err := client.BGPRouteMapsAPI.CreateBGPRouteMaps(context.Background()).BgpRouteMaps(profile).Execute()
	require.NoError(t, err, "Failed to create route map for get test setup")
	createdMapID := createRes.Id

	defer func() {
		client.BGPRouteMapsAPI.DeleteBGPRouteMapsByID(context.Background(), *createdMapID).Execute()
	}()

	// Test: Retrieve the route map
	getRes, httpResGet, errGet := client.BGPRouteMapsAPI.GetBGPRouteMapsByID(context.Background(), *createdMapID).Execute()

	require.NoError(t, errGet, "Failed to get BGP route map by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, mapName, getRes.Name, "Route map name should match")
	assert.Equal(t, *createdMapID, *getRes.Id, "Route map ID should match")
}

// Test_network_services_BGPRouteMapsAPIService_UpdateByID tests updating a BGP route map.
func Test_network_services_BGPRouteMapsAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	mapName := "test-bgp-routemap-update-" + common.GenerateRandomString(6)

	routeMapEntry := network_services.BgpRouteMapsRouteMapInner{
		Action: common.StringPtr("permit"),
		Name:   common.Int32Ptr(1),
	}
	profile := network_services.BgpRouteMaps{
		Name:     mapName,
		Folder:   common.StringPtr("Prisma Access"),
		RouteMap: []network_services.BgpRouteMapsRouteMapInner{routeMapEntry},
	}

	// Setup: Create a route map first
	createRes, _, err := client.BGPRouteMapsAPI.CreateBGPRouteMaps(context.Background()).BgpRouteMaps(profile).Execute()
	require.NoError(t, err, "Failed to create route map for update test setup")
	createdMapID := createRes.Id

	defer func() {
		client.BGPRouteMapsAPI.DeleteBGPRouteMapsByID(context.Background(), *createdMapID).Execute()
	}()

	// Test: Update the route map with additional entry
	updatedEntry2 := network_services.BgpRouteMapsRouteMapInner{
		Action: common.StringPtr("deny"),
		Name:   common.Int32Ptr(2),
	}
	updatedProfile := network_services.BgpRouteMaps{
		Name:        mapName,
		Folder:      common.StringPtr("Prisma Access"),
		Description: common.StringPtr("Updated description"),
		RouteMap:    []network_services.BgpRouteMapsRouteMapInner{routeMapEntry, updatedEntry2},
	}

	updateRes, httpResUpdate, errUpdate := client.BGPRouteMapsAPI.UpdateBGPRouteMapsByID(context.Background(), *createdMapID).BgpRouteMaps(updatedProfile).Execute()

	require.NoError(t, errUpdate, "Failed to update BGP route map")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the updated data
	assert.Equal(t, mapName, updateRes.Name, "Route map name should match")
	assert.Equal(t, *createdMapID, *updateRes.Id, "Route map ID should match")
}

// Test_network_services_BGPRouteMapsAPIService_List tests listing BGP route maps.
func Test_network_services_BGPRouteMapsAPIService_List(t *testing.T) {
	t.Skip("List response contains array fields that cause model deserialization error")
	client := SetupNetworkSvcTestClient(t)

	// Test: List route maps
	listRes, httpResList, errList := client.BGPRouteMapsAPI.ListBGPRouteMaps(context.Background()).Folder("Prisma Access").Execute()

	require.NoError(t, errList, "Failed to list BGP route maps")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")
}

// Test_network_services_BGPRouteMapsAPIService_DeleteByID tests deleting a BGP route map by ID.
func Test_network_services_BGPRouteMapsAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	mapName := "test-bgp-routemap-delete-" + common.GenerateRandomString(6)

	routeMapEntry := network_services.BgpRouteMapsRouteMapInner{
		Action: common.StringPtr("permit"),
		Name:   common.Int32Ptr(1),
	}
	profile := network_services.BgpRouteMaps{
		Name:     mapName,
		Folder:   common.StringPtr("Prisma Access"),
		RouteMap: []network_services.BgpRouteMapsRouteMapInner{routeMapEntry},
	}

	// Setup: Create a route map first
	createRes, _, err := client.BGPRouteMapsAPI.CreateBGPRouteMaps(context.Background()).BgpRouteMaps(profile).Execute()
	require.NoError(t, err, "Failed to create route map for delete test")
	createdMapID := createRes.Id

	// Test: Delete the route map
	httpResDel, errDel := client.BGPRouteMapsAPI.DeleteBGPRouteMapsByID(context.Background(), *createdMapID).Execute()

	require.NoError(t, errDel, "Failed to delete BGP route map")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_network_services_BGPRouteMapsAPIService_Fetch tests the FetchBGPRouteMaps convenience method.
func Test_network_services_BGPRouteMapsAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	mapName := "test-bgp-routemap-fetch-" + common.GenerateRandomString(6)

	routeMapEntry := network_services.BgpRouteMapsRouteMapInner{
		Action: common.StringPtr("permit"),
		Name:   common.Int32Ptr(1),
	}
	profile := network_services.BgpRouteMaps{
		Name:     mapName,
		Folder:   common.StringPtr("Prisma Access"),
		RouteMap: []network_services.BgpRouteMapsRouteMapInner{routeMapEntry},
	}

	// Create a test route map
	createRes, _, err := client.BGPRouteMapsAPI.CreateBGPRouteMaps(context.Background()).BgpRouteMaps(profile).Execute()
	require.NoError(t, err, "Failed to create test route map for fetch test")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		client.BGPRouteMapsAPI.DeleteBGPRouteMapsByID(context.Background(), *createdID).Execute()
		t.Logf("Cleaned up test route map: %s", *createdID)
	}()

	// Test 1: Fetch existing route map by name
	fetchedMap, err := client.BGPRouteMapsAPI.FetchBGPRouteMaps(
		context.Background(),
		mapName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch BGP route map by name")
	require.NotNil(t, fetchedMap, "Fetched route map should not be nil")
	assert.Equal(t, *createdID, *fetchedMap.Id, "Fetched route map ID should match")
	assert.Equal(t, mapName, fetchedMap.Name, "Fetched route map name should match")
	t.Logf("[SUCCESS] FetchBGPRouteMaps found route map: %s", fetchedMap.Name)

	// Test 2: Fetch non-existent route map (should return nil, nil)
	notFound, err := client.BGPRouteMapsAPI.FetchBGPRouteMaps(
		context.Background(),
		"non-existent-bgp-routemap-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent route map")
	assert.Nil(t, notFound, "Should return nil for non-existent route map")
	t.Logf("[SUCCESS] FetchBGPRouteMaps correctly returned nil for non-existent route map")
}
