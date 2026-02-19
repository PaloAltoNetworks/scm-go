/*
Network Services Testing SDWANErrorCorrectionProfilesAPIService
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

// Test_networkservices_SDWANErrorCorrectionProfilesAPIService_Create tests creating an SD-WAN error correction profile
func Test_networkservices_SDWANErrorCorrectionProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object
	testName := "test-sdwan-ecp-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanErrorCorrectionProfiles{
		Name:                testName,
		Folder:              common.StringPtr("All"),
		ActivationThreshold: 2,
		Mode: network_services.SdwanErrorCorrectionProfilesMode{
			ForwardErrorCorrection: &network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{
				Ratio:            "10% (20:2)",
				RecoveryDuration: 1000,
			},
		},
	}

	createReq := client.SDWANErrorCorrectionProfilesAPI.CreateSDWANErrorCorrectionProfiles(context.Background()).SdwanErrorCorrectionProfiles(testObj)
	resp, httpResp, err := createReq.Execute()

	// Cleanup after test
	if resp != nil && resp.Id != nil {
		defer func() {
			deleteReq := client.SDWANErrorCorrectionProfilesAPI.DeleteSDWANErrorCorrectionProfilesByID(context.Background(), *resp.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *resp.Id)
		}()
	}

	// Verify the response
	require.NoError(t, err, "Failed to create SD-WAN error correction profile")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 201, httpResp.StatusCode, "Expected 201 Created status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testName, resp.Name, "Created profile name should match")
	t.Logf("[SUCCESS] Created SD-WAN error correction profile: %s", resp.Name)
}

// Test_networkservices_SDWANErrorCorrectionProfilesAPIService_List tests listing SD-WAN error correction profiles
func Test_networkservices_SDWANErrorCorrectionProfilesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-ecp-list-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanErrorCorrectionProfiles{
		Name:                testName,
		Folder:              common.StringPtr("All"),
		ActivationThreshold: 2,
		Mode: network_services.SdwanErrorCorrectionProfilesMode{
			ForwardErrorCorrection: &network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{
				Ratio:            "10% (20:2)",
				RecoveryDuration: 1000,
			},
		},
	}

	createReq := client.SDWANErrorCorrectionProfilesAPI.CreateSDWANErrorCorrectionProfiles(context.Background()).SdwanErrorCorrectionProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test
	defer func() {
		if createRes.Id != nil {
			deleteReq := client.SDWANErrorCorrectionProfilesAPI.DeleteSDWANErrorCorrectionProfilesByID(context.Background(), *createRes.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *createRes.Id)
		}
	}()

	// List SD-WAN error correction profiles
	listReq := client.SDWANErrorCorrectionProfilesAPI.ListSDWANErrorCorrectionProfiles(context.Background()).Folder("All").Limit(200).Offset(0)
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list SD-WAN error correction profiles")
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
	t.Logf("[SUCCESS] Listed SD-WAN error correction profiles, found test object: %s", testName)
}

// Test_networkservices_SDWANErrorCorrectionProfilesAPIService_GetByID tests retrieving an SD-WAN error correction profile by ID
func Test_networkservices_SDWANErrorCorrectionProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-ecp-get-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanErrorCorrectionProfiles{
		Name:                testName,
		Folder:              common.StringPtr("All"),
		ActivationThreshold: 2,
		Mode: network_services.SdwanErrorCorrectionProfilesMode{
			ForwardErrorCorrection: &network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{
				Ratio:            "10% (20:2)",
				RecoveryDuration: 1000,
			},
		},
	}

	createReq := client.SDWANErrorCorrectionProfilesAPI.CreateSDWANErrorCorrectionProfiles(context.Background()).SdwanErrorCorrectionProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Cleanup after test
	defer func() {
		deleteReq := client.SDWANErrorCorrectionProfilesAPI.DeleteSDWANErrorCorrectionProfilesByID(context.Background(), *createRes.Id)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", *createRes.Id)
	}()

	// Get by ID
	getReq := client.SDWANErrorCorrectionProfilesAPI.GetSDWANErrorCorrectionProfilesByID(context.Background(), *createRes.Id)
	getRes, httpRes, err := getReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to get SD-WAN error correction profile by ID")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Response should not be nil")
	assert.Equal(t, *createRes.Id, *getRes.Id, "Retrieved profile ID should match")
	assert.Equal(t, testName, getRes.Name, "Retrieved profile name should match")
	t.Logf("[SUCCESS] Retrieved SD-WAN error correction profile by ID: %s", *getRes.Id)
}

// Test_networkservices_SDWANErrorCorrectionProfilesAPIService_UpdateByID tests updating an SD-WAN error correction profile
func Test_networkservices_SDWANErrorCorrectionProfilesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-ecp-update-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanErrorCorrectionProfiles{
		Name:                testName,
		Folder:              common.StringPtr("All"),
		ActivationThreshold: 2,
		Mode: network_services.SdwanErrorCorrectionProfilesMode{
			ForwardErrorCorrection: &network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{
				Ratio:            "10% (20:2)",
				RecoveryDuration: 1000,
			},
		},
	}

	createReq := client.SDWANErrorCorrectionProfilesAPI.CreateSDWANErrorCorrectionProfiles(context.Background()).SdwanErrorCorrectionProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Cleanup after test
	defer func() {
		deleteReq := client.SDWANErrorCorrectionProfilesAPI.DeleteSDWANErrorCorrectionProfilesByID(context.Background(), *createRes.Id)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", *createRes.Id)
	}()

	// Update the object
	updatedObj := network_services.SdwanErrorCorrectionProfiles{
		Name:                testName,
		Folder:              common.StringPtr("All"),
		ActivationThreshold: 3, // Updated threshold (was 2)
		Mode: network_services.SdwanErrorCorrectionProfilesMode{
			ForwardErrorCorrection: &network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{
				Ratio:            "20% (20:4)", // Updated ratio
				RecoveryDuration: 2000,         // Updated duration
			},
		},
	}

	updateReq := client.SDWANErrorCorrectionProfilesAPI.UpdateSDWANErrorCorrectionProfilesByID(context.Background(), *createRes.Id).SdwanErrorCorrectionProfiles(updatedObj)
	updateRes, httpRes, err := updateReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to update SD-WAN error correction profile")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Response should not be nil")
	assert.Equal(t, int32(3), updateRes.ActivationThreshold, "Updated threshold should match")
	t.Logf("[SUCCESS] Updated SD-WAN error correction profile: %s", *updateRes.Id)
}

// Test_networkservices_SDWANErrorCorrectionProfilesAPIService_DeleteByID tests deleting an SD-WAN error correction profile
func Test_networkservices_SDWANErrorCorrectionProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-ecp-delete-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanErrorCorrectionProfiles{
		Name:                testName,
		Folder:              common.StringPtr("All"),
		ActivationThreshold: 2,
		Mode: network_services.SdwanErrorCorrectionProfilesMode{
			ForwardErrorCorrection: &network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{
				Ratio:            "10% (20:2)",
				RecoveryDuration: 1000,
			},
		},
	}

	createReq := client.SDWANErrorCorrectionProfilesAPI.CreateSDWANErrorCorrectionProfiles(context.Background()).SdwanErrorCorrectionProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Delete the object
	deleteReq := client.SDWANErrorCorrectionProfilesAPI.DeleteSDWANErrorCorrectionProfilesByID(context.Background(), *createRes.Id)
	httpRes, err := deleteReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to delete SD-WAN error correction profile")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	t.Logf("[SUCCESS] Deleted SD-WAN error correction profile: %s", *createRes.Id)
}

// Test_networkservices_SDWANErrorCorrectionProfilesAPIService_FetchSDWANErrorCorrectionProfiles tests the Fetch convenience method
func Test_networkservices_SDWANErrorCorrectionProfilesAPIService_FetchSDWANErrorCorrectionProfiles(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-ecp-fetch-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanErrorCorrectionProfiles{
		Name:                testName,
		Folder:              common.StringPtr("All"),
		ActivationThreshold: 2,
		Mode: network_services.SdwanErrorCorrectionProfilesMode{
			ForwardErrorCorrection: &network_services.SdwanErrorCorrectionProfilesModeForwardErrorCorrection{
				Ratio:            "10% (20:2)",
				RecoveryDuration: 1000,
			},
		},
	}

	createReq := client.SDWANErrorCorrectionProfilesAPI.CreateSDWANErrorCorrectionProfiles(context.Background()).SdwanErrorCorrectionProfiles(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test
	defer func() {
		if createRes.Id != nil {
			deleteReq := client.SDWANErrorCorrectionProfilesAPI.DeleteSDWANErrorCorrectionProfilesByID(context.Background(), *createRes.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *createRes.Id)
		}
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.SDWANErrorCorrectionProfilesAPI.FetchSDWANErrorCorrectionProfiles(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch SD-WAN error correction profile by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchSDWANErrorCorrectionProfiles found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.SDWANErrorCorrectionProfilesAPI.FetchSDWANErrorCorrectionProfiles(
		context.Background(),
		"non-existent-sdwan-ecp-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchSDWANErrorCorrectionProfiles correctly returned nil for non-existent object")
}
