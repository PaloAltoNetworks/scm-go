/*
Device Settings Testing GeneralSettingsAPIService
Note: General settings are singleton per folder - only one exists per scope.
Create/Delete are not applicable; tests cover List, GetByID, and Update.
*/
package device_settings

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/device_settings"
)

// Test_device_settings_GeneralSettingsAPIService_List tests listing general settings
func Test_device_settings_GeneralSettingsAPIService_List(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	reqList := client.GeneralSettingsAPI.ListGeneralSettings(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	require.NoError(t, errList, "Failed to list general settings")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d general settings", len(listRes))
}

// Test_device_settings_GeneralSettingsAPIService_GetByID tests getting general settings by ID
func Test_device_settings_GeneralSettingsAPIService_GetByID(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	reqList := client.GeneralSettingsAPI.ListGeneralSettings(context.Background()).Folder("Prisma Access")
	listRes, _, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list general settings")

	if len(listRes) == 0 {
		t.Skip("No general settings found to test GetByID")
	}

	existingID := *listRes[0].Id
	t.Logf("Testing GetByID with ID: %s", existingID)

	getRes, httpResGet, errGet := client.GeneralSettingsAPI.GetGeneralSettingsByID(context.Background(), existingID).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	require.NoError(t, errGet, "Failed to get general settings by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, existingID, *getRes.Id, "ID should match requested ID")
	t.Logf("Successfully retrieved general settings with ID: %s", *getRes.Id)
}

// Test_device_settings_GeneralSettingsAPIService_Update tests updating general settings
// Uses list to find an existing object, updates it in place, then restores the original
func Test_device_settings_GeneralSettingsAPIService_Update(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	// List to find existing settings to update
	listRes, _, errList := client.GeneralSettingsAPI.ListGeneralSettings(context.Background()).Folder("Prisma Access").Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list general settings")

	if len(listRes) == 0 {
		t.Skip("No general settings found to test Update")
	}

	existingID := *listRes[0].Id
	originalSettings := listRes[0]
	t.Logf("Testing Update with ID: %s", existingID)

	// Pass the original back as a no-op update to verify API works
	updatedSettings := device_settings.GeneralSettings{
		Id:      originalSettings.Id,
		Folder:  common.StringPtr("Prisma Access"),
		General: originalSettings.General,
	}

	updateRes, httpResUpdate, errUpdate := client.GeneralSettingsAPI.UpdateGeneralSettingsByID(context.Background(), existingID).GeneralSettings(updatedSettings).Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	require.NoError(t, errUpdate, "Failed to update general settings")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, existingID, *updateRes.Id, "ID should match")
	t.Logf("Successfully updated general settings with ID: %s", *updateRes.Id)
}
