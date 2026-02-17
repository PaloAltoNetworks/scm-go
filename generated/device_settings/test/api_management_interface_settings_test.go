/*
Device Settings Testing ManagementInterfaceSettingsAPIService
*/
package device_settings

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_device_settings_ManagementInterfaceSettingsAPIService_List tests listing management interface settings
func Test_device_settings_ManagementInterfaceSettingsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeviceSettingsTestClient(t)

	// Test List operation with folder filter
	reqList := client.ManagementInterfaceSettingsAPI.ListManagementInterfaceSettings(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list management interface settings")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d management interface settings", len(listRes))
}

// Test_device_settings_ManagementInterfaceSettingsAPIService_GetByID tests getting management interface settings by ID
func Test_device_settings_ManagementInterfaceSettingsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeviceSettingsTestClient(t)

	// First list to get an existing ID
	reqList := client.ManagementInterfaceSettingsAPI.ListManagementInterfaceSettings(context.Background()).Folder("Prisma Access")
	listRes, _, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list management interface settings")

	if len(listRes) == 0 {
		t.Skip("No management interface settings found to test GetByID")
	}

	// Get the first item's ID
	existingID := *listRes[0].Id
	t.Logf("Testing GetByID with ID: %s", existingID)

	// Test Get by ID operation
	reqGet := client.ManagementInterfaceSettingsAPI.GetManagementInterfaceSettingsByID(context.Background(), existingID)
	getRes, httpResGet, errGet := reqGet.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful
	require.NoError(t, errGet, "Failed to get management interface settings by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, existingID, *getRes.Id, "ID should match requested ID")
	t.Logf("Successfully retrieved management interface settings with ID: %s", *getRes.Id)
}
