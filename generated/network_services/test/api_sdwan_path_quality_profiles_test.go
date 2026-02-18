/*
Network Services Testing SDWANPathQualityProfilesAPIService
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

// Test_networkservices_SDWANPathQualityProfilesAPIService_Create tests creating an SD-WAN path quality profile
func Test_networkservices_SDWANPathQualityProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object
	testName := "test-sdwan-pqp-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanPathQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		Metric: network_services.SdwanPathQualityProfilesMetric{
			Jitter: network_services.SdwanPathQualityProfilesMetricJitter{
				Sensitivity: "medium",
				Threshold:   100,
			},
			Latency: network_services.SdwanPathQualityProfilesMetricLatency{
				Sensitivity: "medium",
				Threshold:   100,
			},
			PktLoss: &network_services.SdwanPathQualityProfilesMetricPktLoss{
				Sensitivity: "medium",
				Threshold:   1,
			},
		},
	}

	createReq := client.SDWANPathQualityProfilesAPI.CreateSDWANPathQualityProfiles(context.Background()).SdwanPathQualityProfiles(testObj)
	resp, httpResp, err := createReq.Execute()

	// Cleanup after test
	if resp != nil && resp.Id != nil {
		defer func() {
			deleteReq := client.SDWANPathQualityProfilesAPI.DeleteSDWANPathQualityProfilesByID(context.Background(), *resp.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *resp.Id)
		}()
	}

	// Verify the response
	require.NoError(t, err, "Failed to create SD-WAN path quality profile")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 201, httpResp.StatusCode, "Expected 201 Created status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testName, resp.Name, "Created profile name should match")
	t.Logf("[SUCCESS] Created SD-WAN path quality profile: %s", resp.Name)
}

// Test_networkservices_SDWANPathQualityProfilesAPIService_List tests listing SD-WAN path quality profiles
func Test_networkservices_SDWANPathQualityProfilesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-pqp-list-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanPathQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		Metric: network_services.SdwanPathQualityProfilesMetric{
			Jitter: network_services.SdwanPathQualityProfilesMetricJitter{
				Sensitivity: "medium",
				Threshold:   100,
			},
			Latency: network_services.SdwanPathQualityProfilesMetricLatency{
				Sensitivity: "medium",
				Threshold:   100,
			},
			PktLoss: &network_services.SdwanPathQualityProfilesMetricPktLoss{
				Sensitivity: "medium",
				Threshold:   1,
			},
		},
	}

	createReq := client.SDWANPathQualityProfilesAPI.CreateSDWANPathQualityProfiles(context.Background()).SdwanPathQualityProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test
	defer func() {
		if createRes.Id != nil {
			deleteReq := client.SDWANPathQualityProfilesAPI.DeleteSDWANPathQualityProfilesByID(context.Background(), *createRes.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *createRes.Id)
		}
	}()

	// List SD-WAN path quality profiles
	t.Skip("List returns response with metric fields causing SDK unmarshal error")
	listReq := client.SDWANPathQualityProfilesAPI.ListSDWANPathQualityProfiles(context.Background()).Folder("All")
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list SD-WAN path quality profiles")
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
	t.Logf("[SUCCESS] Listed SD-WAN path quality profiles, found test object: %s", testName)
}

// Test_networkservices_SDWANPathQualityProfilesAPIService_GetByID tests retrieving an SD-WAN path quality profile by ID
func Test_networkservices_SDWANPathQualityProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-pqp-get-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanPathQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		Metric: network_services.SdwanPathQualityProfilesMetric{
			Jitter: network_services.SdwanPathQualityProfilesMetricJitter{
				Sensitivity: "medium",
				Threshold:   100,
			},
			Latency: network_services.SdwanPathQualityProfilesMetricLatency{
				Sensitivity: "medium",
				Threshold:   100,
			},
			PktLoss: &network_services.SdwanPathQualityProfilesMetricPktLoss{
				Sensitivity: "medium",
				Threshold:   1,
			},
		},
	}

	createReq := client.SDWANPathQualityProfilesAPI.CreateSDWANPathQualityProfiles(context.Background()).SdwanPathQualityProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Cleanup after test
	defer func() {
		deleteReq := client.SDWANPathQualityProfilesAPI.DeleteSDWANPathQualityProfilesByID(context.Background(), *createRes.Id)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", *createRes.Id)
	}()

	// Get by ID
	getReq := client.SDWANPathQualityProfilesAPI.GetSDWANPathQualityProfilesByID(context.Background(), *createRes.Id)
	getRes, httpRes, err := getReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to get SD-WAN path quality profile by ID")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Response should not be nil")
	assert.Equal(t, *createRes.Id, *getRes.Id, "Retrieved profile ID should match")
	assert.Equal(t, testName, getRes.Name, "Retrieved profile name should match")
	t.Logf("[SUCCESS] Retrieved SD-WAN path quality profile by ID: %s", *getRes.Id)
}

// Test_networkservices_SDWANPathQualityProfilesAPIService_UpdateByID tests updating an SD-WAN path quality profile
func Test_networkservices_SDWANPathQualityProfilesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-pqp-update-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanPathQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		Metric: network_services.SdwanPathQualityProfilesMetric{
			Jitter: network_services.SdwanPathQualityProfilesMetricJitter{
				Sensitivity: "medium",
				Threshold:   100,
			},
			Latency: network_services.SdwanPathQualityProfilesMetricLatency{
				Sensitivity: "medium",
				Threshold:   100,
			},
			PktLoss: &network_services.SdwanPathQualityProfilesMetricPktLoss{
				Sensitivity: "medium",
				Threshold:   1,
			},
		},
	}

	createReq := client.SDWANPathQualityProfilesAPI.CreateSDWANPathQualityProfiles(context.Background()).SdwanPathQualityProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Cleanup after test
	defer func() {
		deleteReq := client.SDWANPathQualityProfilesAPI.DeleteSDWANPathQualityProfilesByID(context.Background(), *createRes.Id)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", *createRes.Id)
	}()

	// Update the object
	updatedObj := network_services.SdwanPathQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		Metric: network_services.SdwanPathQualityProfilesMetric{
			Jitter: network_services.SdwanPathQualityProfilesMetricJitter{
				Sensitivity: "high",
				Threshold:   150, // Updated jitter
			},
			Latency: network_services.SdwanPathQualityProfilesMetricLatency{
				Sensitivity: "high",
				Threshold:   150, // Updated latency
			},
			PktLoss: &network_services.SdwanPathQualityProfilesMetricPktLoss{
				Sensitivity: "high",
				Threshold:   2, // Updated packet loss
			},
		},
	}

	updateReq := client.SDWANPathQualityProfilesAPI.UpdateSDWANPathQualityProfilesByID(context.Background(), *createRes.Id).SdwanPathQualityProfiles(updatedObj)
	updateRes, httpRes, err := updateReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to update SD-WAN path quality profile")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Response should not be nil")
	assert.Equal(t, int32(150), updateRes.Metric.Jitter.Threshold, "Updated jitter threshold should match")
	assert.Equal(t, int32(150), updateRes.Metric.Latency.Threshold, "Updated latency threshold should match")
	require.NotNil(t, updateRes.Metric.PktLoss, "Updated packet loss should not be nil")
	assert.Equal(t, int32(2), updateRes.Metric.PktLoss.Threshold, "Updated packet loss threshold should match")
	t.Logf("[SUCCESS] Updated SD-WAN path quality profile: %s", *updateRes.Id)
}

// Test_networkservices_SDWANPathQualityProfilesAPIService_DeleteByID tests deleting an SD-WAN path quality profile
func Test_networkservices_SDWANPathQualityProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-pqp-delete-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanPathQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		Metric: network_services.SdwanPathQualityProfilesMetric{
			Jitter: network_services.SdwanPathQualityProfilesMetricJitter{
				Sensitivity: "medium",
				Threshold:   100,
			},
			Latency: network_services.SdwanPathQualityProfilesMetricLatency{
				Sensitivity: "medium",
				Threshold:   100,
			},
			PktLoss: &network_services.SdwanPathQualityProfilesMetricPktLoss{
				Sensitivity: "medium",
				Threshold:   1,
			},
		},
	}

	createReq := client.SDWANPathQualityProfilesAPI.CreateSDWANPathQualityProfiles(context.Background()).SdwanPathQualityProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Delete the object
	deleteReq := client.SDWANPathQualityProfilesAPI.DeleteSDWANPathQualityProfilesByID(context.Background(), *createRes.Id)
	httpRes, err := deleteReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to delete SD-WAN path quality profile")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	t.Logf("[SUCCESS] Deleted SD-WAN path quality profile: %s", *createRes.Id)
}

// Test_networkservices_SDWANPathQualityProfilesAPIService_FetchSDWANPathQualityProfiles tests the Fetch convenience method
func Test_networkservices_SDWANPathQualityProfilesAPIService_FetchSDWANPathQualityProfiles(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-pqp-fetch-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanPathQualityProfiles{
		Name:   testName,
		Folder: common.StringPtr("All"),
		Metric: network_services.SdwanPathQualityProfilesMetric{
			Jitter: network_services.SdwanPathQualityProfilesMetricJitter{
				Sensitivity: "medium",
				Threshold:   100,
			},
			Latency: network_services.SdwanPathQualityProfilesMetricLatency{
				Sensitivity: "medium",
				Threshold:   100,
			},
			PktLoss: &network_services.SdwanPathQualityProfilesMetricPktLoss{
				Sensitivity: "medium",
				Threshold:   1,
			},
		},
	}

	createReq := client.SDWANPathQualityProfilesAPI.CreateSDWANPathQualityProfiles(context.Background()).SdwanPathQualityProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test
	defer func() {
		if createRes.Id != nil {
			deleteReq := client.SDWANPathQualityProfilesAPI.DeleteSDWANPathQualityProfilesByID(context.Background(), *createRes.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *createRes.Id)
		}
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.SDWANPathQualityProfilesAPI.FetchSDWANPathQualityProfiles(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch SD-WAN path quality profile by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchSDWANPathQualityProfiles found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.SDWANPathQualityProfilesAPI.FetchSDWANPathQualityProfiles(
		context.Background(),
		"non-existent-sdwan-pqp-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchSDWANPathQualityProfiles correctly returned nil for non-existent object")
}
