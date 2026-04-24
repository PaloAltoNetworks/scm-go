/*
 * Mobile Agent Testing
 *
 * UserLocationsAPIService
 */

package mobile_agent

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/mobile_agent"
)

// newTestUserLocation returns a minimal valid ForwardingProfileUserLocations.
// The API allows all fields to be optional, so we create a minimal object with just a name.
func newTestUserLocation(name string) mobile_agent.ForwardingProfileUserLocations {
	userLoc := mobile_agent.NewForwardingProfileUserLocations()
	userLoc.SetName(name)
	return *userLoc
}

// Test_mobile_agent_UserLocationsAPIService_Create tests the creation of a user location
// with all possible attributes including description and ip_addresses.
func Test_mobile_agent_UserLocationsAPIService_Create(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	locationName := "test-userloc-create-" + randomSuffix

	// Create a user location with description and IP addresses
	userLoc := mobile_agent.NewForwardingProfileUserLocations()
	userLoc.SetName(locationName)
	userLoc.SetDescription("Test user location for create with all attributes")
	userLoc.SetIpAddresses([]string{"192.168.1.0/24", "10.0.0.0/8", "172.16.0.1"})

	fmt.Printf("Attempting to create User Location with name: %s\n", userLoc.GetName())

	req := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(*userLoc)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create User Location")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created User Location should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up User Location with ID: %s", createdID)
		_, errDel := client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete User Location during cleanup")
	}()

	assert.Equal(t, locationName, res.GetName(), "Created User Location name should match")
	assert.Equal(t, "Test user location for create with all attributes", res.GetDescription(), "Description should match")

	require.NotNil(t, res.IpAddresses, "IP addresses list should not be nil")
	assert.ElementsMatch(t, []string{"192.168.1.0/24", "10.0.0.0/8", "172.16.0.1"}, res.IpAddresses, "IP addresses should match")

	t.Logf("Successfully created and validated User Location: %s with ID: %s", userLoc.GetName(), createdID)
}

// Test_mobile_agent_UserLocationsAPIService_CreateWithInternalHostDetection tests the creation of a user location
// with internal host detection (without ip_addresses).
func Test_mobile_agent_UserLocationsAPIService_CreateWithInternalHostDetection(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	locationName := "test-userloc-inthost-" + randomSuffix

	// Create a user location with internal host detection
	userLoc := mobile_agent.NewForwardingProfileUserLocations()
	userLoc.SetName(locationName)
	userLoc.SetDescription("Test user location with internal host detection")

	// Set internal host detection
	internalHostDetection := mobile_agent.NewForwardingProfileUserLocationsInternalHostDetection(
		"internal.example.com",
		"192.168.100.1",
	)
	userLoc.SetInternalHostDetection(*internalHostDetection)

	fmt.Printf("Attempting to create User Location with internal host detection: %s\n", userLoc.GetName())

	req := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(*userLoc)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create User Location")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created User Location should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up User Location with ID: %s", createdID)
		_, errDel := client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete User Location during cleanup")
	}()

	assert.Equal(t, locationName, res.GetName(), "Created User Location name should match")
	assert.Equal(t, "Test user location with internal host detection", res.GetDescription(), "Description should match")

	require.NotNil(t, res.InternalHostDetection, "Internal host detection should not be nil")
	assert.Equal(t, "internal.example.com", res.InternalHostDetection.GetFqdn(), "FQDN should match")
	assert.Equal(t, "192.168.100.1", res.InternalHostDetection.GetIpAddress(), "IP address should match")

	t.Logf("Successfully created and validated User Location with internal host detection: %s with ID: %s", userLoc.GetName(), createdID)
}

// Test_mobile_agent_UserLocationsAPIService_CreateMinimal tests the creation of a user location
// with only a name (minimal required fields).
func Test_mobile_agent_UserLocationsAPIService_CreateMinimal(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	locationName := "test-userloc-minimal-" + randomSuffix

	userLoc := newTestUserLocation(locationName)

	fmt.Printf("Attempting to create minimal User Location with name: %s\n", userLoc.GetName())

	req := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(userLoc)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create minimal User Location")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created User Location should have an ID")

	createdID := *res.Id

	defer func() {
		_, errDel := client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete User Location during cleanup")
	}()

	assert.Equal(t, locationName, res.GetName(), "Created User Location name should match")

	t.Logf("Successfully created minimal User Location: %s with ID: %s", userLoc.GetName(), *res.Id)
}

// Test_mobile_agent_UserLocationsAPIService_CreateWithIPAddressesOnly tests creating a user location
// with just name and IP addresses (no internal host detection).
func Test_mobile_agent_UserLocationsAPIService_CreateWithIPAddressesOnly(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	locationName := "test-userloc-iponly-" + randomSuffix

	userLoc := newTestUserLocation(locationName)
	userLoc.SetDescription("User location with IP addresses only")
	userLoc.SetIpAddresses([]string{"10.10.0.0/16", "192.168.5.0/24"})

	req := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(userLoc)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create User Location with IP addresses only")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")

	defer func() {
		if res.Id != nil {
			client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), *res.Id).Execute()
		}
	}()

	assert.Equal(t, locationName, res.GetName(), "Name should match")
	assert.ElementsMatch(t, []string{"10.10.0.0/16", "192.168.5.0/24"}, res.IpAddresses, "IP addresses should match")

	t.Logf("Successfully created User Location with IP addresses only: %s", locationName)
}

// Test_mobile_agent_UserLocationsAPIService_List tests listing user locations.
func Test_mobile_agent_UserLocationsAPIService_List(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	locationName := "test-userloc-list-" + randomSuffix

	createRes, _, err := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(newTestUserLocation(locationName)).
		Execute()
	require.NoError(t, err, "Failed to create User Location for list test")
	require.NotNil(t, createRes.Id, "Created User Location should have an ID")

	defer func() {
		client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), *createRes.Id).Execute()
	}()

	listRes, httpResList, errList := client.UserLocationsAPI.ListGlobalProtectUserLocations(context.Background()).
		Folder("Mobile Users").
		Limit(10000).
		Execute()
	require.NoError(t, errList, "Failed to list User Locations")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundLocation := false
	for _, loc := range listRes.Data {
		if loc.GetName() == locationName {
			foundLocation = true
			assert.Equal(t, *createRes.Id, *loc.Id, "Listed location ID should match created ID")
			break
		}
	}
	assert.True(t, foundLocation, "Created User Location should be found in the list")
}

// Test_mobile_agent_UserLocationsAPIService_ListWithNameFilter tests listing user locations
// with a name filter.
func Test_mobile_agent_UserLocationsAPIService_ListWithNameFilter(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	locationName := "test-userloc-filter-" + randomSuffix

	createRes, _, err := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(newTestUserLocation(locationName)).
		Execute()
	require.NoError(t, err, "Failed to create User Location for filter test")
	require.NotNil(t, createRes.Id, "Created User Location should have an ID")

	defer func() {
		client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), *createRes.Id).Execute()
	}()

	listRes, httpResList, errList := client.UserLocationsAPI.ListGlobalProtectUserLocations(context.Background()).
		Folder("Mobile Users").
		Name(locationName).
		Limit(10).
		Execute()
	require.NoError(t, errList, "Failed to list User Locations with name filter")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundLocation := false
	for _, loc := range listRes.Data {
		if loc.GetName() == locationName {
			foundLocation = true
			assert.Equal(t, *createRes.Id, *loc.Id, "Filtered location ID should match created ID")
			break
		}
	}
	assert.True(t, foundLocation, "Created User Location should be found with name filter")
}

// Test_mobile_agent_UserLocationsAPIService_GetByID tests retrieving a user location by its ID.
func Test_mobile_agent_UserLocationsAPIService_GetByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	locationName := "test-userloc-getbyid-" + randomSuffix

	createRes, _, err := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(newTestUserLocation(locationName)).
		Execute()
	require.NoError(t, err, "Failed to create User Location for get by ID test")
	require.NotNil(t, createRes.Id, "Created User Location should have an ID")
	createdID := *createRes.Id

	defer func() {
		client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.UserLocationsAPI.GetGlobalProtectUserLocationByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get User Location by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, locationName, getRes.GetName(), "Name should match")
	assert.Equal(t, createdID, *getRes.Id, "ID should match")

	t.Logf("Successfully retrieved User Location by ID: %s", createdID)
}

// Test_mobile_agent_UserLocationsAPIService_Update tests updating an existing user location.
func Test_mobile_agent_UserLocationsAPIService_Update(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	locationName := "test-userloc-update-" + randomSuffix

	createRes, _, err := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(newTestUserLocation(locationName)).
		Execute()
	require.NoError(t, err, "Failed to create User Location for update test")
	require.NotNil(t, createRes.Id, "Created User Location should have an ID")
	createdID := *createRes.Id

	defer func() {
		client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), createdID).Execute()
	}()

	updatedLocation := newTestUserLocation(locationName)
	updatedLocation.SetDescription("Updated description for user location")
	updatedLocation.SetIpAddresses([]string{"172.20.0.0/16", "10.5.5.0/24"})

	updateRes, httpResUpdate, errUpdate := client.UserLocationsAPI.UpdateGlobalProtectUserLocationByID(context.Background(), createdID).
		ForwardingProfileUserLocations(updatedLocation).
		Execute()
	require.NoError(t, errUpdate, "Failed to update User Location")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, locationName, updateRes.GetName(), "Name should remain the same after update")
	assert.Equal(t, "Updated description for user location", updateRes.GetDescription(), "Description should be updated")
	assert.ElementsMatch(t, []string{"172.20.0.0/16", "10.5.5.0/24"}, updateRes.IpAddresses, "IP addresses should be updated")

	t.Logf("Successfully updated User Location: %s with ID: %s", locationName, createdID)
}

// Test_mobile_agent_UserLocationsAPIService_UpdateWithInternalHostDetection tests updating a user location
// to use internal host detection instead of IP addresses.
func Test_mobile_agent_UserLocationsAPIService_UpdateWithInternalHostDetection(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	locationName := "test-userloc-upd-inthost-" + randomSuffix

	// Create initial location with IP addresses
	initialLoc := newTestUserLocation(locationName)
	initialLoc.SetIpAddresses([]string{"10.1.0.0/16"})

	createRes, _, err := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(initialLoc).
		Execute()
	require.NoError(t, err, "Failed to create User Location for update test")
	require.NotNil(t, createRes.Id, "Created User Location should have an ID")
	createdID := *createRes.Id

	defer func() {
		client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), createdID).Execute()
	}()

	// Update to use internal host detection
	updatedLocation := newTestUserLocation(locationName)
	updatedLocation.SetDescription("Updated to use internal host detection")

	internalHostDetection := mobile_agent.NewForwardingProfileUserLocationsInternalHostDetection(
		"updated.example.com",
		"172.20.1.1",
	)
	updatedLocation.SetInternalHostDetection(*internalHostDetection)

	updateRes, httpResUpdate, errUpdate := client.UserLocationsAPI.UpdateGlobalProtectUserLocationByID(context.Background(), createdID).
		ForwardingProfileUserLocations(updatedLocation).
		Execute()
	require.NoError(t, errUpdate, "Failed to update User Location")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, locationName, updateRes.GetName(), "Name should remain the same after update")
	assert.Equal(t, "Updated to use internal host detection", updateRes.GetDescription(), "Description should be updated")

	require.NotNil(t, updateRes.InternalHostDetection, "Internal host detection should not be nil")
	assert.Equal(t, "updated.example.com", updateRes.InternalHostDetection.GetFqdn(), "FQDN should be updated")
	assert.Equal(t, "172.20.1.1", updateRes.InternalHostDetection.GetIpAddress(), "IP address should be updated")

	t.Logf("Successfully updated User Location to use internal host detection: %s with ID: %s", locationName, createdID)
}

// Test_mobile_agent_UserLocationsAPIService_DeleteByID tests deleting a user location by ID.
func Test_mobile_agent_UserLocationsAPIService_DeleteByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	locationName := "test-userloc-delete-" + randomSuffix

	createRes, _, err := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(newTestUserLocation(locationName)).
		Execute()
	require.NoError(t, err, "Failed to create User Location for delete test")
	require.NotNil(t, createRes.Id, "Created User Location should have an ID")
	createdID := *createRes.Id

	_, errDel := client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete User Location")

	t.Logf("Successfully deleted User Location with ID: %s", createdID)
}

// Test_mobile_agent_UserLocationsAPIService_FetchUserLocations tests the FetchUserLocations convenience method.
func Test_mobile_agent_UserLocationsAPIService_FetchUserLocations(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-userloc-fetch-" + randomSuffix

	testObj := newTestUserLocation(testName)
	testObj.SetDescription("Test user location for fetch")
	testObj.SetIpAddresses([]string{"10.20.30.0/24"})

	createReq := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileUserLocations(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object should have an ID")
	createdID := *createRes.Id

	defer func() {
		client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), createdID).Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.UserLocationsAPI.FetchUserLocations(
		context.Background(),
		testName,
		common.StringPtr("Mobile Users"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch user_locations by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.GetName(), "Fetched object name should match")
	assert.Equal(t, "Test user location for fetch", fetchedObj.GetDescription(), "Description should match")
	assert.ElementsMatch(t, []string{"10.20.30.0/24"}, fetchedObj.IpAddresses, "IP addresses should match")
	t.Logf("[SUCCESS] FetchUserLocations found object: %s", fetchedObj.GetName())

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.UserLocationsAPI.FetchUserLocations(
		context.Background(),
		"non-existent-user-location-xyz-12345",
		common.StringPtr("Mobile Users"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchUserLocations correctly returned nil for non-existent object")
}

// Test_mobile_agent_UserLocationsAPIService_ListPagination tests pagination parameters.
func Test_mobile_agent_UserLocationsAPIService_ListPagination(t *testing.T) {
	client := SetupMobileAgentTestClient(t)

	// Create multiple test objects to test pagination
	randomSuffix := common.GenerateRandomString(6)
	createdIDs := []string{}
	for i := 0; i < 3; i++ {
		locationName := fmt.Sprintf("test-userloc-page-%s-%d", randomSuffix, i)
		createRes, _, err := client.UserLocationsAPI.CreateGlobalProtectUserLocation(context.Background()).
			Folder("Mobile Users").
			ForwardingProfileUserLocations(newTestUserLocation(locationName)).
			Execute()
		require.NoError(t, err, "Failed to create test User Location for pagination test")
		if createRes.Id != nil {
			createdIDs = append(createdIDs, *createRes.Id)
		}
	}

	defer func() {
		for _, id := range createdIDs {
			client.UserLocationsAPI.DeleteGlobalProtectUserLocation(context.Background(), id).Execute()
		}
	}()

	// Test with limit
	listRes, httpRes, err := client.UserLocationsAPI.ListGlobalProtectUserLocations(context.Background()).
		Folder("Mobile Users").
		Limit(2).
		Offset(0).
		Execute()
	require.NoError(t, err, "Failed to list User Locations with pagination")
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, listRes)

	t.Logf("Retrieved %d items with limit=2", len(listRes.Data))

	// Test with offset
	listRes2, httpRes2, err2 := client.UserLocationsAPI.ListGlobalProtectUserLocations(context.Background()).
		Folder("Mobile Users").
		Limit(10).
		Offset(1).
		Execute()
	require.NoError(t, err2, "Failed to list User Locations with offset")
	assert.Equal(t, 200, httpRes2.StatusCode)
	require.NotNil(t, listRes2)

	t.Logf("Retrieved %d items with offset=1", len(listRes2.Data))
	t.Logf("Pagination test completed successfully")
}
