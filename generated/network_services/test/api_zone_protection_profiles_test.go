/*
 * Network Services Testing
 *
 * ZoneProtectionProfilesAPIService
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

// Test_network_services_ZoneProtectionProfilesAPIService_Create tests the creation of a Zone Protection Profile.
func Test_network_services_ZoneProtectionProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	profileName := "test-zpp-create-" + randomSuffix

	profile := network_services.ZoneProtectionProfiles{
		Name:                     profileName,
		Folder:                   common.StringPtr("All"),
		Description:              common.StringPtr("Test zone protection profile for create"),
		DiscardIcmpEmbeddedError: common.BoolPtr(true),
		IcmpFragDiscard:          common.BoolPtr(true),
	}

	fmt.Printf("Attempting to create Zone Protection Profile with name: %s\n", profile.Name)

	req := client.ZoneProtectionProfilesAPI.CreateZoneProtectionProfiles(context.Background()).ZoneProtectionProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Zone Protection Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created zone protection profile should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up Zone Protection Profile with ID: %s", createdID)
		_, errDel := client.ZoneProtectionProfilesAPI.DeleteZoneProtectionProfilesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete zone protection profile during cleanup")
	}()

	assert.Equal(t, profileName, res.Name, "Created zone protection profile name should match")
	t.Logf("Successfully created and validated Zone Protection Profile: %s with ID: %s", profile.Name, createdID)
}

// Test_network_services_ZoneProtectionProfilesAPIService_GetByID tests retrieving a zone protection profile by its ID.
func Test_network_services_ZoneProtectionProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	profileName := "test-zpp-get-" + randomSuffix

	profile := network_services.ZoneProtectionProfiles{
		Name:            profileName,
		Folder:          common.StringPtr("All"),
		IcmpFragDiscard: common.BoolPtr(false),
	}

	createRes, _, err := client.ZoneProtectionProfilesAPI.CreateZoneProtectionProfiles(context.Background()).ZoneProtectionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create zone protection profile for get test")
	createdID := *createRes.Id

	defer func() {
		client.ZoneProtectionProfilesAPI.DeleteZoneProtectionProfilesByID(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.ZoneProtectionProfilesAPI.GetZoneProtectionProfilesByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get zone protection profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_network_services_ZoneProtectionProfilesAPIService_Update tests updating an existing zone protection profile.
func Test_network_services_ZoneProtectionProfilesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	profileName := "test-zpp-update-" + randomSuffix

	profile := network_services.ZoneProtectionProfiles{
		Name:            profileName,
		Folder:          common.StringPtr("All"),
		IcmpFragDiscard: common.BoolPtr(false),
	}

	createRes, _, err := client.ZoneProtectionProfilesAPI.CreateZoneProtectionProfiles(context.Background()).ZoneProtectionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create zone protection profile for update test")
	createdID := *createRes.Id

	defer func() {
		client.ZoneProtectionProfilesAPI.DeleteZoneProtectionProfilesByID(context.Background(), createdID).Execute()
	}()

	updatedProfile := network_services.ZoneProtectionProfiles{
		Name:                     profileName,
		Folder:                   common.StringPtr("All"),
		Description:              common.StringPtr("Updated zone protection profile"),
		IcmpFragDiscard:          common.BoolPtr(true),
		DiscardIcmpEmbeddedError: common.BoolPtr(true),
	}

	updateRes, httpResUpdate, errUpdate := client.ZoneProtectionProfilesAPI.UpdateZoneProtectionProfilesByID(context.Background(), createdID).ZoneProtectionProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update zone protection profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, true, *updateRes.IcmpFragDiscard, "IcmpFragDiscard should be updated")
}

// Test_network_services_ZoneProtectionProfilesAPIService_List tests listing Zone Protection Profiles.
func Test_network_services_ZoneProtectionProfilesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	profileName := "test-zpp-list-" + randomSuffix

	profile := network_services.ZoneProtectionProfiles{
		Name:   profileName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.ZoneProtectionProfilesAPI.CreateZoneProtectionProfiles(context.Background()).ZoneProtectionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create zone protection profile for list test")
	createdID := *createRes.Id

	defer func() {
		client.ZoneProtectionProfilesAPI.DeleteZoneProtectionProfilesByID(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.ZoneProtectionProfilesAPI.ListZoneProtectionProfiles(context.Background()).Folder("All").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list zone protection profiles")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundProfile := false
	for _, prof := range listRes.Data {
		if prof.Name == profileName {
			foundProfile = true
			break
		}
	}
	assert.True(t, foundProfile, "Created zone protection profile should be found in the list")
}

// Test_network_services_ZoneProtectionProfilesAPIService_DeleteByID tests deleting a zone protection profile by ID.
func Test_network_services_ZoneProtectionProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	profileName := "test-zpp-delete-" + randomSuffix

	profile := network_services.ZoneProtectionProfiles{
		Name:   profileName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.ZoneProtectionProfilesAPI.CreateZoneProtectionProfiles(context.Background()).ZoneProtectionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create zone protection profile for delete test")
	createdID := *createRes.Id

	_, errDel := client.ZoneProtectionProfilesAPI.DeleteZoneProtectionProfilesByID(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete zone protection profile")
}

// Test_network_services_ZoneProtectionProfilesAPIService_FetchZoneProtectionProfiles tests the FetchZoneProtectionProfiles convenience method
func Test_network_services_ZoneProtectionProfilesAPIService_FetchZoneProtectionProfiles(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-zpp-fetch-" + randomSuffix

	testObj := network_services.ZoneProtectionProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
	}

	createReq := client.ZoneProtectionProfilesAPI.CreateZoneProtectionProfiles(context.Background()).ZoneProtectionProfiles(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		deleteReq := client.ZoneProtectionProfilesAPI.DeleteZoneProtectionProfilesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.ZoneProtectionProfilesAPI.FetchZoneProtectionProfiles(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch zone_protection_profiles by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchZoneProtectionProfiles found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.ZoneProtectionProfilesAPI.FetchZoneProtectionProfiles(
		context.Background(),
		"non-existent-zone_protection_profiles-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchZoneProtectionProfiles correctly returned nil for non-existent object")
}
