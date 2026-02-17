/*
Device Settings Testing ServiceSettingsAPIService
Note: Service settings are singleton per folder - only one exists per scope.
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

// Test_device_settings_ServiceSettingsAPIService_List tests listing service settings
func Test_device_settings_ServiceSettingsAPIService_List(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	reqList := client.ServiceSettingsAPI.ListServiceSettings(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	require.NoError(t, errList, "Failed to list service settings")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d service settings", len(listRes))
}

// Test_device_settings_ServiceSettingsAPIService_GetByID tests getting service settings by ID
func Test_device_settings_ServiceSettingsAPIService_GetByID(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	// List to get an existing ID (singleton - always exists)
	listRes, _, errList := client.ServiceSettingsAPI.ListServiceSettings(context.Background()).Folder("Prisma Access").Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list service settings")

	if len(listRes) == 0 {
		t.Skip("No service settings found to test GetByID")
	}

	existingID := *listRes[0].Id
	t.Logf("Testing GetByID with ID: %s", existingID)

	getRes, httpResGet, errGet := client.ServiceSettingsAPI.GetServiceSettingsByID(context.Background(), existingID).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	require.NoError(t, errGet, "Failed to get service settings by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, existingID, *getRes.Id, "ID should match requested ID")
	t.Logf("Successfully retrieved service settings with ID: %s", *getRes.Id)
}

// Test_device_settings_ServiceSettingsAPIService_Update tests updating service settings
// Uses list to find the existing singleton, updates it in place with a no-op update
func Test_device_settings_ServiceSettingsAPIService_Update(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	// List to get the existing singleton
	listRes, _, errList := client.ServiceSettingsAPI.ListServiceSettings(context.Background()).Folder("Prisma Access").Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list service settings")

	if len(listRes) == 0 {
		t.Skip("No service settings found to test Update")
	}

	existingID := *listRes[0].Id
	originalSettings := listRes[0]
	t.Logf("Testing Update with ID: %s", existingID)

	// Pass the original back as a no-op update to verify the API works
	updatedSettings := device_settings.ServiceSettings{
		Id:       originalSettings.Id,
		Folder:   common.StringPtr("Prisma Access"),
		Services: originalSettings.Services,
	}

	updateRes, httpResUpdate, errUpdate := client.ServiceSettingsAPI.UpdateServiceSettingsByID(context.Background(), existingID).ServiceSettings(updatedSettings).Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	require.NoError(t, errUpdate, "Failed to update service settings")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, existingID, *updateRes.Id, "ID should match")
	t.Logf("Successfully updated service settings with ID: %s", *updateRes.Id)
}
