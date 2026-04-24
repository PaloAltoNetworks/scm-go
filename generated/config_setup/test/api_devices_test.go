/*
Config Setup Testing DevicesAPIService

Tests for the Devices API - read-only operations (List, Get, Fetch)
Devices are managed externally, so only read operations are tested.
*/
package config_setup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_config_setup_DevicesAPIService_List tests listing devices
// This is a read-only test that lists existing devices
func Test_config_setup_DevicesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Test List operation
	reqList := client.DevicesAPI.ListDevices(context.Background()).Limit(200).Offset(0)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list devices")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")

	t.Logf("Successfully listed devices, count: %d", len(listRes.Data))

	// Log first few devices for debugging
	for i, device := range listRes.Data {
		if i >= 3 {
			break
		}
		t.Logf("  Device %d: ID=%s, Name=%s", i+1, device.Id, device.Name)
	}
}

// Test_config_setup_DevicesAPIService_GetByID tests retrieving a device by its ID
// This test first lists devices to get a valid ID, then retrieves that device
func Test_config_setup_DevicesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// First, list devices to get a valid device ID
	reqList := client.DevicesAPI.ListDevices(context.Background()).Limit(10).Offset(0)
	listRes, _, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list devices")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")

	// Skip test if no devices exist
	if len(listRes.Data) == 0 {
		t.Skip("No devices found in the system, skipping GetByID test")
	}

	// Get the first device's ID
	firstDevice := listRes.Data[0]
	deviceID := firstDevice.Id
	deviceName := firstDevice.Name
	t.Logf("Using device from list: ID=%s, Name=%s", deviceID, deviceName)

	// Test Get by ID operation
	reqGetById := client.DevicesAPI.GetDeviceByID(context.Background(), deviceID)
	getRes, httpResGet, errGet := reqGetById.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful
	require.NoError(t, errGet, "Failed to get device by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, deviceID, getRes.Id, "Device ID should match")
	assert.Equal(t, deviceName, getRes.Name, "Device name should match")

	t.Logf("Successfully retrieved device: ID=%s, Name=%s", getRes.Id, getRes.Name)
}

// Test_config_setup_DevicesAPIService_FetchDevices tests the FetchDevices convenience method
// This test first lists devices to get a valid name, then fetches that device by name
func Test_config_setup_DevicesAPIService_FetchDevices(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// First, list devices to get a valid device name
	reqList := client.DevicesAPI.ListDevices(context.Background()).Limit(10).Offset(0)
	listRes, _, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list devices")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")

	// Skip test if no devices exist
	if len(listRes.Data) == 0 {
		t.Skip("No devices found in the system, skipping FetchDevices test")
	}

	// Get the first device's name and ID for verification
	firstDevice := listRes.Data[0]
	deviceName := firstDevice.Name
	deviceID := firstDevice.Id
	t.Logf("Using device from list: ID=%s, Name=%s", deviceID, deviceName)

	// Test 1: Fetch existing device by name
	fetchedDevice, err := client.DevicesAPI.FetchDevices(
		context.Background(),
		deviceName,
		nil, // folder
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch device by name")
	require.NotNil(t, fetchedDevice, "Fetched device should not be nil")
	assert.Equal(t, deviceID, fetchedDevice.Id, "Fetched device ID should match")
	assert.Equal(t, deviceName, fetchedDevice.Name, "Fetched device name should match")
	t.Logf("[SUCCESS] FetchDevices found device: ID=%s, Name=%s", fetchedDevice.Id, fetchedDevice.Name)

	// Test 2: Fetch non-existent device (should return nil, nil)
	notFound, err := client.DevicesAPI.FetchDevices(
		context.Background(),
		"non-existent-device-xyz-12345",
		nil, // folder
		nil, // snippet
		nil, // device
	)
	require.NoError(t, err, "Fetch should not error for non-existent device")
	assert.Nil(t, notFound, "Should return nil for non-existent device")
	t.Logf("[SUCCESS] FetchDevices correctly returned nil for non-existent device")
}
