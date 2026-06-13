/*
 * Mobile Agent Testing
 *
 * DestinationsAPIService
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

// newTestForwardingProfileDestinations returns a minimal valid ForwardingProfileDestinations. The API requires at
// least one FQDN or IP address entry in addition to 'name'.
func newTestForwardingProfileDestinations(name string) mobile_agent.ForwardingProfileDestinations {
	fqdnEntry := mobile_agent.NewForwardingProfileDestinationFqdnEntry("www.google.com")
	fqdnEntry.SetPort(80)
	d := mobile_agent.NewForwardingProfileDestinations(name)
	d.Fqdn = []mobile_agent.ForwardingProfileDestinationFqdnEntry{*fqdnEntry}
	return *d
}

// Test_mobile_agent_DestinationsAPIService_Create tests the creation of a destination
// with both FQDN and IP address entries fully populated.
func Test_mobile_agent_DestinationsAPIService_Create(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	destName := "test-dest-create-" + randomSuffix

	fqdnEntry := mobile_agent.NewForwardingProfileDestinationFqdnEntry("www.google.com")
	fqdnEntry.SetPort(80)

	ipEntry := mobile_agent.NewForwardingProfileDestinationIpEntry("10.2.3.4")
	ipEntry.SetPort(345)

	dest := mobile_agent.ForwardingProfileDestinations{
		Name:        destName,
		Description: common.StringPtr("test"),
		Fqdn:        []mobile_agent.ForwardingProfileDestinationFqdnEntry{*fqdnEntry},
		IpAddresses: []mobile_agent.ForwardingProfileDestinationIpEntry{*ipEntry},
	}

	fmt.Printf("Attempting to create ForwardingProfileDestinations with name: %s\n", dest.Name)

	res, httpRes, err := client.DestinationsAPI.CreateGlobalProtectDestination(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileDestinations(dest).
		Execute()

	common.LogRequestIDOnFailure(t, httpRes)
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create ForwardingProfileDestinations")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created ForwardingProfileDestinations should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up ForwardingProfileDestinations with ID: %s", createdID)
		_, errDel := client.DestinationsAPI.DeleteGlobalProtectDestination(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete ForwardingProfileDestinations during cleanup")
	}()

	assert.Equal(t, destName, res.Name, "Created ForwardingProfileDestinations name should match")
	assert.Equal(t, "test", *res.Description, "Description should match")

	require.Len(t, res.Fqdn, 1, "Should have exactly one FQDN entry")
	assert.Equal(t, "www.google.com", res.Fqdn[0].Name, "FQDN name should match")
	assert.Equal(t, int32(80), res.Fqdn[0].GetPort(), "FQDN port should match")

	require.Len(t, res.IpAddresses, 1, "Should have exactly one IP address entry")
	assert.Equal(t, "10.2.3.4", res.IpAddresses[0].Name, "IP address name should match")
	assert.Equal(t, int32(345), res.IpAddresses[0].GetPort(), "IP address port should match")

	t.Logf("Successfully created and validated ForwardingProfileDestinations: %s with ID: %s", dest.Name, createdID)
}

// Test_mobile_agent_DestinationsAPIService_GetByID tests retrieving a destination by its ID.
func Test_mobile_agent_DestinationsAPIService_GetByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	destName := "test-dest-get-" + randomSuffix

	createRes, httpResCreate, err := client.DestinationsAPI.CreateGlobalProtectDestination(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileDestinations(newTestForwardingProfileDestinations(destName)).
		Execute()
	common.LogRequestIDOnFailure(t, httpResCreate)
	require.NoError(t, err, "Failed to create ForwardingProfileDestinations for get test")
	createdID := *createRes.Id

	defer func() {
		client.DestinationsAPI.DeleteGlobalProtectDestination(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.DestinationsAPI.GetGlobalProtectDestinationByID(context.Background(), createdID).Execute()
	common.LogRequestIDOnFailure(t, httpResGet)
	require.NoError(t, errGet, "Failed to get ForwardingProfileDestinations by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, destName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_mobile_agent_DestinationsAPIService_Update tests updating an existing destination.
func Test_mobile_agent_DestinationsAPIService_Update(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	destName := "test-dest-update-" + randomSuffix

	createRes, httpResCreate, err := client.DestinationsAPI.CreateGlobalProtectDestination(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileDestinations(newTestForwardingProfileDestinations(destName)).
		Execute()
	common.LogRequestIDOnFailure(t, httpResCreate)
	require.NoError(t, err, "Failed to create ForwardingProfileDestinations for update test")
	createdID := *createRes.Id

	defer func() {
		client.DestinationsAPI.DeleteGlobalProtectDestination(context.Background(), createdID).Execute()
	}()

	updatedDest := newTestForwardingProfileDestinations(destName)
	updatedDest.Description = common.StringPtr("Updated description")

	updateRes, httpResUpdate, errUpdate := client.DestinationsAPI.UpdateGlobalProtectDestinationByID(context.Background(), createdID).
		ForwardingProfileDestinations(updatedDest).
		Execute()
	common.LogRequestIDOnFailure(t, httpResUpdate)
	require.NoError(t, errUpdate, "Failed to update ForwardingProfileDestinations")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, destName, updateRes.Name, "Name should remain the same after update")
	assert.Equal(t, "Updated description", *updateRes.Description, "Description should be updated")
}

// Test_mobile_agent_DestinationsAPIService_List tests listing destinations.
func Test_mobile_agent_DestinationsAPIService_List(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	destName := "test-dest-list-" + randomSuffix

	createRes, httpResCreate, err := client.DestinationsAPI.CreateGlobalProtectDestination(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileDestinations(newTestForwardingProfileDestinations(destName)).
		Execute()
	common.LogRequestIDOnFailure(t, httpResCreate)
	require.NoError(t, err, "Failed to create ForwardingProfileDestinations for list test")
	createdID := *createRes.Id

	defer func() {
		client.DestinationsAPI.DeleteGlobalProtectDestination(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.DestinationsAPI.ListGlobalProtectDestinations(context.Background()).
		Folder("Mobile Users").
		Limit(10000).
		Execute()
	common.LogRequestIDOnFailure(t, httpResList)
	require.NoError(t, errList, "Failed to list Destinations")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundDest := false
	for _, d := range listRes.Data {
		if d.Name == destName {
			foundDest = true
			break
		}
	}
	assert.True(t, foundDest, "Created ForwardingProfileDestinations should be found in the list")
}

// Test_mobile_agent_DestinationsAPIService_DeleteByID tests deleting a destination by ID.
func Test_mobile_agent_DestinationsAPIService_DeleteByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	destName := "test-dest-delete-" + randomSuffix

	createRes, httpResCreate, err := client.DestinationsAPI.CreateGlobalProtectDestination(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileDestinations(newTestForwardingProfileDestinations(destName)).
		Execute()
	common.LogRequestIDOnFailure(t, httpResCreate)
	require.NoError(t, err, "Failed to create ForwardingProfileDestinations for delete test")
	createdID := *createRes.Id

	httpResDel, errDel := client.DestinationsAPI.DeleteGlobalProtectDestination(context.Background(), createdID).Execute()
	common.LogRequestIDOnFailure(t, httpResDel)
	require.NoError(t, errDel, "Failed to delete ForwardingProfileDestinations")
}

// Test_mobile_agent_DestinationsAPIService_GetByID_NotFound tests that fetching a non-existent destination returns 404.
func Test_mobile_agent_DestinationsAPIService_GetByID_NotFound(t *testing.T) {
	client := SetupMobileAgentTestClient(t)

	nonExistentID := "00000000-0000-0000-0000-000000000000"
	_, httpRes, err := client.DestinationsAPI.GetGlobalProtectDestinationByID(context.Background(), nonExistentID).Execute()
	common.LogRequestIDOnFailure(t, httpRes)
	assert.Equal(t, 404, httpRes.StatusCode, "Expected 404 Not Found status")
}

// Test_mobile_agent_DestinationsAPIService_UpdateByID_NotFound tests that updating a non-existent destination returns 404.
func Test_mobile_agent_DestinationsAPIService_UpdateByID_NotFound(t *testing.T) {
	client := SetupMobileAgentTestClient(t)

	nonExistentID := "00000000-0000-0000-0000-000000000000"
	_, httpRes, _ := client.DestinationsAPI.UpdateGlobalProtectDestinationByID(context.Background(), nonExistentID).
		ForwardingProfileDestinations(newTestForwardingProfileDestinations("non-existent-dest")).
		Execute()
	common.LogRequestIDOnFailure(t, httpRes)
	assert.Equal(t, 404, httpRes.StatusCode, "Expected 404 Not Found status")
}

// Test_mobile_agent_DestinationsAPIService_DeleteByID_NotFound tests that deleting a non-existent destination returns 404.
func Test_mobile_agent_DestinationsAPIService_DeleteByID_NotFound(t *testing.T) {
	client := SetupMobileAgentTestClient(t)

	nonExistentID := "00000000-0000-0000-0000-000000000000"
	httpRes, err := client.DestinationsAPI.DeleteGlobalProtectDestination(context.Background(), nonExistentID).Execute()
	common.LogRequestIDOnFailure(t, httpRes)
	assert.Equal(t, 404, httpRes.StatusCode, "Expected 404 Not Found status")
}

// Test_mobile_agent_DestinationsAPIService_DeleteByID_VerifyGone tests that a deleted destination is no longer retrievable.
func Test_mobile_agent_DestinationsAPIService_DeleteByID_VerifyGone(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	destName := "test-dest-gone-" + randomSuffix

	createRes, httpResCreate, err := client.DestinationsAPI.CreateGlobalProtectDestination(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileDestinations(newTestForwardingProfileDestinations(destName)).
		Execute()
	common.LogRequestIDOnFailure(t, httpResCreate)
	require.NoError(t, err, "Failed to create ForwardingProfileDestinations for delete-verify test")
	createdID := *createRes.Id

	httpResDel, errDel := client.DestinationsAPI.DeleteGlobalProtectDestination(context.Background(), createdID).Execute()
	common.LogRequestIDOnFailure(t, httpResDel)
	require.NoError(t, errDel, "Failed to delete ForwardingProfileDestinations")

	_, httpRes, errGet := client.DestinationsAPI.GetGlobalProtectDestinationByID(context.Background(), createdID).Execute()
	common.LogRequestIDOnFailure(t, httpRes)
	require.Error(t, errGet, "Expected an error after deletion")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 404, httpRes.StatusCode, "Expected 404 after deletion")
}

// Test_mobile_agent_DestinationsAPIService_FetchDestinations tests the FetchDestinations convenience method.
func Test_mobile_agent_DestinationsAPIService_FetchDestinations(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-dest-fetch-" + randomSuffix

	testObj := newTestForwardingProfileDestinations(testName)
	testObj.Description = common.StringPtr("Test destination for fetch")

	createRes, httpResCreate, err := client.DestinationsAPI.CreateGlobalProtectDestination(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileDestinations(testObj).
		Execute()
	common.LogRequestIDOnFailure(t, httpResCreate)
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		client.DestinationsAPI.DeleteGlobalProtectDestination(context.Background(), createdID).Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.DestinationsAPI.FetchDestinations(
		context.Background(),
		testName,
		common.StringPtr("Mobile Users"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch destinations by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchDestinations found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.DestinationsAPI.FetchDestinations(
		context.Background(),
		"non-existent-destination-xyz-12345",
		common.StringPtr("Mobile Users"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchDestinations correctly returned nil for non-existent object")
}
