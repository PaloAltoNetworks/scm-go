/*
Device Settings Testing GeneralSettingsAPIService
*/
package device_settings

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_device_settings_GeneralSettingsAPIService_List tests listing general settings
func Test_device_settings_GeneralSettingsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeviceSettingsTestClient(t)

	// Test List operation with folder filter
	reqList := client.GeneralSettingsAPI.ListGeneralSettings(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list general settings")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d general settings", len(listRes))
}

// Test_device_settings_GeneralSettingsAPIService_GetByID tests getting general settings by ID
func Test_device_settings_GeneralSettingsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeviceSettingsTestClient(t)

	// First list to get an existing ID
	reqList := client.GeneralSettingsAPI.ListGeneralSettings(context.Background()).Folder("Prisma Access")
	listRes, _, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list general settings")

	if len(listRes) == 0 {
		t.Skip("No general settings found to test GetByID")
	}

	// Get the first item's ID
	existingID := *listRes[0].Id
	t.Logf("Testing GetByID with ID: %s", existingID)

	// Test Get by ID operation
	reqGet := client.GeneralSettingsAPI.GetGeneralSettingsByID(context.Background(), existingID)
	getRes, httpResGet, errGet := reqGet.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful
	require.NoError(t, errGet, "Failed to get general settings by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, existingID, *getRes.Id, "ID should match requested ID")
	t.Logf("Successfully retrieved general settings with ID: %s", *getRes.Id)
}
