/*
 * Config Setup Testing
 *
 * DevicesAPIService
 */
/*
Config Setup Testing DevicesAPIService

NOTE: Devices cannot be created or deleted via the API — they are onboarded
externally (e.g. via Azure vWAN or Panorama). These integration tests operate
on devices that already exist in the SCM tenant.

The tests are:
  - List:     List all devices and verify the response shape.
  - GetByID:  Retrieve a specific device by its UUID.
  - Update:   Move a device to a different folder and restore it.
*/
package config_setup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/config_setup"
)

// --- Test Cases ---

// Test_config_setup_DevicesAPIService_List tests listing devices
func Test_config_setup_DevicesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Test List operation
	reqList := client.DevicesAPI.ListDevices(context.Background()).Limit(10)
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

	t.Logf("Successfully listed devices. Total: %d, Returned: %d", listRes.Total, len(listRes.Data))

	// Verify each device in the list has required fields
	for i, device := range listRes.Data {
		assert.NotEmpty(t, device.Id, "Device %d should have an ID", i)
		assert.NotEmpty(t, device.Name, "Device %d should have a Name", i)
		assert.NotEmpty(t, device.Folder, "Device %d should have a Folder", i)
		t.Logf("  Device %d: Name=%s, Folder=%s, ID=%s", i, device.Name, device.Folder, device.Id)
	}
}

// ---

// Test_config_setup_DevicesAPIService_ListWithFolderFilter tests listing devices filtered by folder
func Test_config_setup_DevicesAPIService_ListWithFolderFilter(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// First, list all devices to discover available folders
	reqListAll := client.DevicesAPI.ListDevices(context.Background()).Limit(10)
	listAllRes, _, errListAll := reqListAll.Execute()
	if errListAll != nil {
		handleAPIError(errListAll)
	}
	require.NoError(t, errListAll, "Failed to list all devices")
	require.NotNil(t, listAllRes, "List response should not be nil")
	require.NotEmpty(t, listAllRes.Data, "Expected at least one device to exist for folder filter test")

	// Use the folder of the first device as our filter
	targetFolder := listAllRes.Data[0].Folder
	t.Logf("Testing folder filter with folder: %s", targetFolder)

	// Test List operation with folder filter
	reqList := client.DevicesAPI.ListDevices(context.Background()).Folder(targetFolder).Limit(100)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list devices with folder filter")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// All returned devices should be in the requested folder
	require.NotNil(t, listRes, "Filtered list response should not be nil")
	for i, device := range listRes.Data {
		assert.Equal(t, targetFolder, device.Folder, "Device %d should be in folder %s", i, targetFolder)
	}

	t.Logf("Successfully listed %d devices in folder: %s", len(listRes.Data), targetFolder)
}

// ---

// Test_config_setup_DevicesAPIService_GetByID tests retrieving a device by its ID
func Test_config_setup_DevicesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// First, list devices to get a valid device ID
	reqList := client.DevicesAPI.ListDevices(context.Background()).Limit(1)
	listRes, _, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list devices")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotEmpty(t, listRes.Data, "Expected at least one device to exist for GetByID test")

	// Use the first device from the list
	targetDevice := listRes.Data[0]
	t.Logf("Testing GetByID with device: Name=%s, ID=%s", targetDevice.Name, targetDevice.Id)

	// Test Get by ID operation
	reqGet := client.DevicesAPI.GetDeviceByID(context.Background(), targetDevice.Id)
	getRes, httpResGet, errGet := reqGet.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful
	require.NoError(t, errGet, "Failed to get device by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, targetDevice.Id, getRes.Id, "Device ID should match")
	assert.Equal(t, targetDevice.Name, getRes.Name, "Device name should match")
	assert.Equal(t, targetDevice.Folder, getRes.Folder, "Device folder should match")

	t.Logf("Successfully retrieved device: Name=%s, Folder=%s", getRes.Name, getRes.Folder)
}

// ---

// Test_config_setup_DevicesAPIService_UpdateByID tests updating a device (e.g., changing its description)
// This test updates the description and then restores the original value.
func Test_config_setup_DevicesAPIService_UpdateByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// 1. List devices to get a valid device ID and its current state
	reqList := client.DevicesAPI.ListDevices(context.Background()).Limit(1)
	listRes, _, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list devices")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotEmpty(t, listRes.Data, "Expected at least one device to exist for update test")

	targetDevice := listRes.Data[0]
	t.Logf("Testing UpdateByID with device: Name=%s, ID=%s, CurrentFolder=%s",
		targetDevice.Name, targetDevice.Id, targetDevice.Folder)

	// Save the original state for restoration
	originalDescription := targetDevice.Description
	originalFolder := targetDevice.Folder

	// 2. Update the device description
	testDescription := "Integration test update - " + common.GenerateRandomString(8)
	updateBody := config_setup.DevicesPut{
		Description: common.StringPtr(testDescription),
		Folder:      &originalFolder, // Keep the same folder
	}

	reqUpdate := client.DevicesAPI.UpdateDeviceByID(context.Background(), targetDevice.Id).DevicesPut(updateBody)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	require.NoError(t, errUpdate, "Failed to update device")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status for update")
	require.NotNil(t, updateRes, "Update response should not be nil")

	t.Logf("Successfully updated device description to: %s", testDescription)

	// 3. Verify the update by reading the device again
	reqGet := client.DevicesAPI.GetDeviceByID(context.Background(), targetDevice.Id)
	getRes, _, errGet := reqGet.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}
	require.NoError(t, errGet, "Failed to read device after update")
	require.NotNil(t, getRes, "Get response should not be nil after update")
	require.NotNil(t, getRes.Description, "Description should not be nil after update")
	assert.Equal(t, testDescription, *getRes.Description, "Description should match the updated value")

	// 4. Restore the original description
	restoreBody := config_setup.DevicesPut{
		Description: originalDescription,
		Folder:      &originalFolder,
	}

	reqRestore := client.DevicesAPI.UpdateDeviceByID(context.Background(), targetDevice.Id).DevicesPut(restoreBody)
	_, _, errRestore := reqRestore.Execute()
	if errRestore != nil {
		handleAPIError(errRestore)
		t.Logf("WARNING: Failed to restore original description for device %s: %v", targetDevice.Id, errRestore)
	} else {
		t.Logf("Successfully restored original description for device: %s", targetDevice.Name)
	}
}

// ---

// Test_config_setup_DevicesAPIService_GetByID_NotFound tests that getting a non-existent device returns an error
func Test_config_setup_DevicesAPIService_GetByID_NotFound(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigSvcTestClient(t)

	// Use a UUID that definitely doesn't exist
	nonExistentID := "00000000-0000-0000-0000-000000000000"

	reqGet := client.DevicesAPI.GetDeviceByID(context.Background(), nonExistentID)
	_, httpResGet, errGet := reqGet.Execute()

	// We expect an error for a device that doesn't exist
	require.Error(t, errGet, "Expected an error for non-existent device ID")
	if httpResGet != nil {
		// Expect either 404 or 400 depending on API implementation
		assert.Contains(t, []int{400, 404}, httpResGet.StatusCode,
			"Expected 400 or 404 status for non-existent device, got %d", httpResGet.StatusCode)
		t.Logf("Got expected error status %d for non-existent device", httpResGet.StatusCode)
	}
}
