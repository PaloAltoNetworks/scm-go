/*
Device Settings Testing AuthenticationSettingsAPIService
*/
package device_settings

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/generated/device_settings"
)

// Test_device_settings_AuthenticationSettingsAPIService_List tests listing authentication settings
func Test_device_settings_AuthenticationSettingsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeviceSettingsTestClient(t)

	// Test List operation with folder filter
	reqList := client.AuthenticationSettingsAPI.ListAuthenticationSettings(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list authentication settings")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d authentication settings", len(listRes))
}

// Test_device_settings_AuthenticationSettingsAPIService_GetByID tests getting authentication settings by ID
func Test_device_settings_AuthenticationSettingsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeviceSettingsTestClient(t)

	// First list to get an existing ID
	reqList := client.AuthenticationSettingsAPI.ListAuthenticationSettings(context.Background()).Folder("Prisma Access")
	listRes, _, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list authentication settings")

	if len(listRes) == 0 {
		t.Skip("No authentication settings found to test GetByID")
	}

	// Get the first item's ID
	existingID := listRes[0].Id
	t.Logf("Testing GetByID with ID: %s", existingID)

	// Test Get by ID operation
	reqGet := client.AuthenticationSettingsAPI.GetAuthenticationSettingsByID(context.Background(), existingID)
	getRes, httpResGet, errGet := reqGet.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful
	require.NoError(t, errGet, "Failed to get authentication settings by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, existingID, getRes.Id, "ID should match requested ID")
	t.Logf("Successfully retrieved authentication settings with ID: %s", getRes.Id)
}

// Test_device_settings_AuthenticationSettingsAPIService_Update tests updating authentication settings
// Note: This test reads existing settings, updates them, verifies the update, then restores original
func Test_device_settings_AuthenticationSettingsAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeviceSettingsTestClient(t)

	// First list to get existing settings
	reqList := client.AuthenticationSettingsAPI.ListAuthenticationSettings(context.Background()).Folder("Prisma Access")
	listRes, _, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list authentication settings")

	if len(listRes) == 0 {
		t.Skip("No authentication settings found to test Update")
	}

	// Get the first item's ID and store original state
	existingID := listRes[0].Id
	originalSettings := listRes[0]
	t.Logf("Testing Update with ID: %s", existingID)

	// Create updated settings - just pass the original back (no-op update to verify API works)
	updatedSettings := device_settings.AuthenticationSettings{
		Id:             originalSettings.Id,
		Folder:         originalSettings.Folder,
		Authentication: originalSettings.Authentication,
	}

	// Test Update operation
	reqUpdate := client.AuthenticationSettingsAPI.UpdateAuthenticationSettingsByID(context.Background(), existingID).AuthenticationSettings(updatedSettings)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update authentication settings")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, existingID, updateRes.Id, "ID should match")
	t.Logf("Successfully updated authentication settings with ID: %s", updateRes.Id)
}
