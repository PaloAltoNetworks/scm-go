/*
Device Settings Testing LoginBannerSettingsAPIService
Note: Login banner settings are singleton per folder - only one exists per scope.
Create/Delete are not applicable; tests cover List, GetByID, and Update.
Note: LoginBannerSettings uses MotdBannerSettings as the model type.
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

// Test_device_settings_LoginBannerSettingsAPIService_List tests listing login banner settings
func Test_device_settings_LoginBannerSettingsAPIService_List(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	reqList := client.LoginBannerSettingsAPI.ListLoginBannerSettings(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	require.NoError(t, errList, "Failed to list login banner settings")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d login banner settings", len(listRes))
}

// Test_device_settings_LoginBannerSettingsAPIService_GetByID tests getting login banner settings by ID
func Test_device_settings_LoginBannerSettingsAPIService_GetByID(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	// List to get an existing ID (singleton - always exists)
	listRes, _, errList := client.LoginBannerSettingsAPI.ListLoginBannerSettings(context.Background()).Folder("Prisma Access").Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list login banner settings")

	if len(listRes) == 0 {
		t.Skip("No login banner settings found to test GetByID")
	}

	existingID := *listRes[0].Id
	t.Logf("Testing GetByID with ID: %s", existingID)

	getRes, httpResGet, errGet := client.LoginBannerSettingsAPI.GetLoginBannerSettingsByID(context.Background(), existingID).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	require.NoError(t, errGet, "Failed to get login banner settings by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, existingID, *getRes.Id, "ID should match requested ID")
	t.Logf("Successfully retrieved login banner settings with ID: %s", *getRes.Id)
}

// Test_device_settings_LoginBannerSettingsAPIService_Update tests updating login banner settings
// Uses list to find the existing singleton, updates it in place with a no-op update
func Test_device_settings_LoginBannerSettingsAPIService_Update(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	// List to get the existing singleton
	listRes, _, errList := client.LoginBannerSettingsAPI.ListLoginBannerSettings(context.Background()).Folder("Prisma Access").Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list login banner settings")

	if len(listRes) == 0 {
		t.Skip("No login banner settings found to test Update")
	}

	existingID := *listRes[0].Id
	originalSettings := listRes[0]
	t.Logf("Testing Update with ID: %s", existingID)

	// Pass the original back as a no-op update to verify the API works
	updatedSettings := device_settings.MotdBannerSettings{
		Id:            originalSettings.Id,
		Folder:        common.StringPtr("Prisma Access"),
		MotdAndBanner: originalSettings.MotdAndBanner,
	}

	updateRes, httpResUpdate, errUpdate := client.LoginBannerSettingsAPI.UpdateLoginBannerSettingsByID(context.Background(), existingID).MotdBannerSettings(updatedSettings).Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	require.NoError(t, errUpdate, "Failed to update login banner settings")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, existingID, *updateRes.Id, "ID should match")
	t.Logf("Successfully updated login banner settings with ID: %s", *updateRes.Id)
}
