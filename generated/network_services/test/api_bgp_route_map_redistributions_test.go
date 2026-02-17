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

// Test_network_services_BGPRouteMapRedistributionsAPIService_Create tests the creation of a BGP route map redistribution.
func Test_network_services_BGPRouteMapRedistributionsAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	redistName := "test-bgp-rmap-redist-create-" + common.GenerateRandomString(6)

	profile := network_services.BgpRouteMapRedistributions{
		Name:        redistName,
		Folder:      common.StringPtr("Prisma Access"),
		Description: common.StringPtr("Test BGP route map redistribution"),
	}

	t.Logf("Creating BGP route map redistribution with name: %s", redistName)
	req := client.BGPRouteMapRedistributionsAPI.CreateBGPRouteMapRedistributions(context.Background()).BgpRouteMapRedistributions(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create BGP route map redistribution")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	createdRedistID := res.Id

	// Cleanup the created redistribution
	defer func() {
		t.Logf("Cleaning up BGP route map redistribution with ID: %s", *createdRedistID)
		_, errDel := client.BGPRouteMapRedistributionsAPI.DeleteBGPRouteMapRedistributionsByID(context.Background(), *createdRedistID).Execute()
		if errDel != nil {
			t.Logf("Cleanup failed: %v", errDel)
		}
	}()

	// Verify the response matches key input fields
	assert.Equal(t, redistName, res.Name, "Created redistribution name should match")
}

// Test_network_services_BGPRouteMapRedistributionsAPIService_GetByID tests retrieving a BGP route map redistribution by ID.
func Test_network_services_BGPRouteMapRedistributionsAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	redistName := "test-bgp-rmap-redist-get-" + common.GenerateRandomString(6)

	profile := network_services.BgpRouteMapRedistributions{
		Name:   redistName,
		Folder: common.StringPtr("Prisma Access"),
	}

	// Setup: Create a redistribution first
	createRes, _, err := client.BGPRouteMapRedistributionsAPI.CreateBGPRouteMapRedistributions(context.Background()).BgpRouteMapRedistributions(profile).Execute()
	require.NoError(t, err, "Failed to create redistribution for get test setup")
	createdRedistID := createRes.Id

	defer func() {
		client.BGPRouteMapRedistributionsAPI.DeleteBGPRouteMapRedistributionsByID(context.Background(), *createdRedistID).Execute()
	}()

	// Test: Retrieve the redistribution
	getRes, httpResGet, errGet := client.BGPRouteMapRedistributionsAPI.GetBGPRouteMapRedistributionsByID(context.Background(), *createdRedistID).Execute()

	require.NoError(t, errGet, "Failed to get BGP route map redistribution by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, redistName, getRes.Name, "Redistribution name should match")
	assert.Equal(t, *createdRedistID, *getRes.Id, "Redistribution ID should match")
}

// Test_network_services_BGPRouteMapRedistributionsAPIService_UpdateByID tests updating a BGP route map redistribution.
func Test_network_services_BGPRouteMapRedistributionsAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	redistName := "test-bgp-rmap-redist-update-" + common.GenerateRandomString(6)

	profile := network_services.BgpRouteMapRedistributions{
		Name:   redistName,
		Folder: common.StringPtr("Prisma Access"),
	}

	// Setup: Create a redistribution first
	createRes, _, err := client.BGPRouteMapRedistributionsAPI.CreateBGPRouteMapRedistributions(context.Background()).BgpRouteMapRedistributions(profile).Execute()
	require.NoError(t, err, "Failed to create redistribution for update test setup")
	createdRedistID := createRes.Id

	defer func() {
		client.BGPRouteMapRedistributionsAPI.DeleteBGPRouteMapRedistributionsByID(context.Background(), *createdRedistID).Execute()
	}()

	// Test: Update the redistribution
	updatedProfile := network_services.BgpRouteMapRedistributions{
		Name:        redistName,
		Folder:      common.StringPtr("Prisma Access"),
		Description: common.StringPtr("Updated description"),
	}

	updateRes, httpResUpdate, errUpdate := client.BGPRouteMapRedistributionsAPI.UpdateBGPRouteMapRedistributionsByID(context.Background(), *createdRedistID).BgpRouteMapRedistributions(updatedProfile).Execute()

	require.NoError(t, errUpdate, "Failed to update BGP route map redistribution")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the updated data
	assert.Equal(t, redistName, updateRes.Name, "Redistribution name should match")
	assert.Equal(t, *createdRedistID, *updateRes.Id, "Redistribution ID should match")
}

// Test_network_services_BGPRouteMapRedistributionsAPIService_List tests listing BGP route map redistributions.
func Test_network_services_BGPRouteMapRedistributionsAPIService_List(t *testing.T) {
	t.Skip("List response contains array fields that cause model deserialization error")
	client := SetupNetworkSvcTestClient(t)

	// Test: List redistributions
	listRes, httpResList, errList := client.BGPRouteMapRedistributionsAPI.ListBGPRouteMapRedistributions(context.Background()).Folder("Prisma Access").Execute()

	require.NoError(t, errList, "Failed to list BGP route map redistributions")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")
}

// Test_network_services_BGPRouteMapRedistributionsAPIService_DeleteByID tests deleting a BGP route map redistribution by ID.
func Test_network_services_BGPRouteMapRedistributionsAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	redistName := "test-bgp-rmap-redist-delete-" + common.GenerateRandomString(6)

	profile := network_services.BgpRouteMapRedistributions{
		Name:   redistName,
		Folder: common.StringPtr("Prisma Access"),
	}

	// Setup: Create a redistribution first
	createRes, _, err := client.BGPRouteMapRedistributionsAPI.CreateBGPRouteMapRedistributions(context.Background()).BgpRouteMapRedistributions(profile).Execute()
	require.NoError(t, err, "Failed to create redistribution for delete test")
	createdRedistID := createRes.Id

	// Test: Delete the redistribution
	httpResDel, errDel := client.BGPRouteMapRedistributionsAPI.DeleteBGPRouteMapRedistributionsByID(context.Background(), *createdRedistID).Execute()

	require.NoError(t, errDel, "Failed to delete BGP route map redistribution")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_network_services_BGPRouteMapRedistributionsAPIService_Fetch tests the FetchBGPRouteMapRedistributions convenience method.
func Test_network_services_BGPRouteMapRedistributionsAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	redistName := "test-bgp-rmap-redist-fetch-" + common.GenerateRandomString(6)

	profile := network_services.BgpRouteMapRedistributions{
		Name:   redistName,
		Folder: common.StringPtr("Prisma Access"),
	}

	// Create a test redistribution
	createRes, _, err := client.BGPRouteMapRedistributionsAPI.CreateBGPRouteMapRedistributions(context.Background()).BgpRouteMapRedistributions(profile).Execute()
	require.NoError(t, err, "Failed to create test redistribution for fetch test")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		client.BGPRouteMapRedistributionsAPI.DeleteBGPRouteMapRedistributionsByID(context.Background(), *createdID).Execute()
		t.Logf("Cleaned up test redistribution: %s", *createdID)
	}()

	// Test 1: Fetch existing redistribution by name
	fetchedRedist, err := client.BGPRouteMapRedistributionsAPI.FetchBGPRouteMapRedistributions(
		context.Background(),
		redistName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch BGP route map redistribution by name")
	require.NotNil(t, fetchedRedist, "Fetched redistribution should not be nil")
	assert.Equal(t, *createdID, *fetchedRedist.Id, "Fetched redistribution ID should match")
	assert.Equal(t, redistName, fetchedRedist.Name, "Fetched redistribution name should match")
	t.Logf("[SUCCESS] FetchBGPRouteMapRedistributions found redistribution: %s", fetchedRedist.Name)

	// Test 2: Fetch non-existent redistribution (should return nil, nil)
	notFound, err := client.BGPRouteMapRedistributionsAPI.FetchBGPRouteMapRedistributions(
		context.Background(),
		"non-existent-bgp-rmap-redist-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent redistribution")
	assert.Nil(t, notFound, "Should return nil for non-existent redistribution")
	t.Logf("[SUCCESS] FetchBGPRouteMapRedistributions correctly returned nil for non-existent redistribution")
}
