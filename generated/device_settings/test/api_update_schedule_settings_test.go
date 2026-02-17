/*
Device Settings Testing UpdateScheduleSettingsAPIService
*/
package device_settings

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_device_settings_UpdateScheduleSettingsAPIService_List tests listing update schedule settings
func Test_device_settings_UpdateScheduleSettingsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeviceSettingsTestClient(t)

	// Test List operation with folder filter
	reqList := client.UpdateScheduleSettingsAPI.ListUpdateScheduleSettings(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list update schedule settings")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d update schedule settings", len(listRes))
}

// Test_device_settings_UpdateScheduleSettingsAPIService_GetByID tests getting update schedule settings by ID
func Test_device_settings_UpdateScheduleSettingsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeviceSettingsTestClient(t)

	// First list to get an existing ID
	reqList := client.UpdateScheduleSettingsAPI.ListUpdateScheduleSettings(context.Background()).Folder("Prisma Access")
	listRes, _, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list update schedule settings")

	if len(listRes) == 0 {
		t.Skip("No update schedule settings found to test GetByID")
	}

	// Get the first item's ID
	existingID := *listRes[0].Id
	t.Logf("Testing GetByID with ID: %s", existingID)

	// Test Get by ID operation
	reqGet := client.UpdateScheduleSettingsAPI.GetUpdateScheduleSettingsByID(context.Background(), existingID)
	getRes, httpResGet, errGet := reqGet.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful
	require.NoError(t, errGet, "Failed to get update schedule settings by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, existingID, *getRes.Id, "ID should match requested ID")
	t.Logf("Successfully retrieved update schedule settings with ID: %s", *getRes.Id)
}
