/*
Network Services Testing SDWANSaaSQualityProfilesAPIService
*/
package network_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Test_networkservices_SDWANSaaSQualityProfilesAPIService_Create tests creating an SD-WAN SaaS quality profile
func Test_networkservices_SDWANSaaSQualityProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object
	testName := "test-sdwan-sqp-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanSaasQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		MonitorMode: network_services.SdwanSaasQualityProfilesMonitorMode{
			Adaptive: map[string]interface{}{},
		},
	}

	createReq := client.SDWANSaaSQualityProfilesAPI.CreateSDWANSaaSQualityProfiles(context.Background()).SdwanSaasQualityProfiles(testObj)
	resp, httpResp, err := createReq.Execute()

	// Cleanup after test
	if resp != nil && resp.Id != nil {
		defer func() {
			deleteReq := client.SDWANSaaSQualityProfilesAPI.DeleteSDWANSaaSQualityProfilesByID(context.Background(), *resp.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *resp.Id)
		}()
	}

	// Verify the response
	require.NoError(t, err, "Failed to create SD-WAN SaaS quality profile")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 201, httpResp.StatusCode, "Expected 201 Created status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testName, resp.Name, "Created profile name should match")
	t.Logf("[SUCCESS] Created SD-WAN SaaS quality profile: %s", resp.Name)
}

// Test_networkservices_SDWANSaaSQualityProfilesAPIService_List tests listing SD-WAN SaaS quality profiles
func Test_networkservices_SDWANSaaSQualityProfilesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-sqp-list-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanSaasQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		MonitorMode: network_services.SdwanSaasQualityProfilesMonitorMode{
			Adaptive: map[string]interface{}{},
		},
	}

	createReq := client.SDWANSaaSQualityProfilesAPI.CreateSDWANSaaSQualityProfiles(context.Background()).SdwanSaasQualityProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test
	defer func() {
		if createRes.Id != nil {
			deleteReq := client.SDWANSaaSQualityProfilesAPI.DeleteSDWANSaaSQualityProfilesByID(context.Background(), *createRes.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *createRes.Id)
		}
	}()

	// List SD-WAN SaaS quality profiles
	listReq := client.SDWANSaaSQualityProfilesAPI.ListSDWANSaaSQualityProfiles(context.Background()).Folder("All").Limit(200).Offset(0)
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list SD-WAN SaaS quality profiles")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")

	// Verify our created object is in the list
	found := false
	if resp.Data != nil {
		for _, item := range resp.Data {
			if item.Name == testName {
				found = true
				break
			}
		}
	}
	assert.True(t, found, "Created profile should be in the list")
	t.Logf("[SUCCESS] Listed SD-WAN SaaS quality profiles, found test object: %s", testName)
}

// Test_networkservices_SDWANSaaSQualityProfilesAPIService_GetByID tests retrieving an SD-WAN SaaS quality profile by ID
func Test_networkservices_SDWANSaaSQualityProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-sqp-get-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanSaasQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		MonitorMode: network_services.SdwanSaasQualityProfilesMonitorMode{
			Adaptive: map[string]interface{}{},
		},
	}

	createReq := client.SDWANSaaSQualityProfilesAPI.CreateSDWANSaaSQualityProfiles(context.Background()).SdwanSaasQualityProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Cleanup after test
	defer func() {
		deleteReq := client.SDWANSaaSQualityProfilesAPI.DeleteSDWANSaaSQualityProfilesByID(context.Background(), *createRes.Id)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", *createRes.Id)
	}()

	// Get by ID
	getReq := client.SDWANSaaSQualityProfilesAPI.GetSDWANSaaSQualityProfilesByID(context.Background(), *createRes.Id)
	getRes, httpRes, err := getReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to get SD-WAN SaaS quality profile by ID")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Response should not be nil")
	assert.Equal(t, *createRes.Id, *getRes.Id, "Retrieved profile ID should match")
	assert.Equal(t, testName, getRes.Name, "Retrieved profile name should match")
	t.Logf("[SUCCESS] Retrieved SD-WAN SaaS quality profile by ID: %s", *getRes.Id)
}

// Test_networkservices_SDWANSaaSQualityProfilesAPIService_UpdateByID tests updating an SD-WAN SaaS quality profile
func Test_networkservices_SDWANSaaSQualityProfilesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-sqp-update-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanSaasQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		MonitorMode: network_services.SdwanSaasQualityProfilesMonitorMode{
			Adaptive: map[string]interface{}{},
		},
	}

	createReq := client.SDWANSaaSQualityProfilesAPI.CreateSDWANSaaSQualityProfiles(context.Background()).SdwanSaasQualityProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Cleanup after test
	defer func() {
		deleteReq := client.SDWANSaaSQualityProfilesAPI.DeleteSDWANSaaSQualityProfilesByID(context.Background(), *createRes.Id)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", *createRes.Id)
	}()

	// Update the object (use same monitor mode for update)
	updatedObj := network_services.SdwanSaasQualityProfiles{
		Name: testName,
		MonitorMode: network_services.SdwanSaasQualityProfilesMonitorMode{
			Adaptive: map[string]interface{}{},
		},
	}

	updateReq := client.SDWANSaaSQualityProfilesAPI.UpdateSDWANSaaSQualityProfilesByID(context.Background(), *createRes.Id).SdwanSaasQualityProfiles(updatedObj)
	updateRes, httpRes, err := updateReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to update SD-WAN SaaS quality profile")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Response should not be nil")
	assert.Equal(t, testName, updateRes.Name, "Updated name should match")
	t.Logf("[SUCCESS] Updated SD-WAN SaaS quality profile: %s", *updateRes.Id)
}

// Test_networkservices_SDWANSaaSQualityProfilesAPIService_DeleteByID tests deleting an SD-WAN SaaS quality profile
func Test_networkservices_SDWANSaaSQualityProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-sqp-delete-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanSaasQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		MonitorMode: network_services.SdwanSaasQualityProfilesMonitorMode{
			Adaptive: map[string]interface{}{},
		},
	}

	createReq := client.SDWANSaaSQualityProfilesAPI.CreateSDWANSaaSQualityProfiles(context.Background()).SdwanSaasQualityProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Delete the object
	deleteReq := client.SDWANSaaSQualityProfilesAPI.DeleteSDWANSaaSQualityProfilesByID(context.Background(), *createRes.Id)
	httpRes, err := deleteReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to delete SD-WAN SaaS quality profile")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	t.Logf("[SUCCESS] Deleted SD-WAN SaaS quality profile: %s", *createRes.Id)
}

// Test_networkservices_SDWANSaaSQualityProfilesAPIService_FetchSDWANSaaSQualityProfiles tests the Fetch convenience method
func Test_networkservices_SDWANSaaSQualityProfilesAPIService_FetchSDWANSaaSQualityProfiles(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-sqp-fetch-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanSaasQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		MonitorMode: network_services.SdwanSaasQualityProfilesMonitorMode{
			Adaptive: map[string]interface{}{},
		},
	}

	createReq := client.SDWANSaaSQualityProfilesAPI.CreateSDWANSaaSQualityProfiles(context.Background()).SdwanSaasQualityProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test
	defer func() {
		if createRes.Id != nil {
			deleteReq := client.SDWANSaaSQualityProfilesAPI.DeleteSDWANSaaSQualityProfilesByID(context.Background(), *createRes.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *createRes.Id)
		}
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.SDWANSaaSQualityProfilesAPI.FetchSDWANSaaSQualityProfiles(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch SD-WAN SaaS quality profile by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchSDWANSaaSQualityProfiles found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.SDWANSaaSQualityProfilesAPI.FetchSDWANSaaSQualityProfiles(
		context.Background(),
		"non-existent-sdwan-sqp-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchSDWANSaaSQualityProfiles correctly returned nil for non-existent object")
}
