/*
Device Settings Testing VPNSettingsAPIService
Note: VPN settings are singleton per folder - only one exists per scope.
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

// Test_device_settings_VPNSettingsAPIService_List tests listing VPN settings
func Test_device_settings_VPNSettingsAPIService_List(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	reqList := client.VPNSettingsAPI.ListVPNSettings(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	require.NoError(t, errList, "Failed to list VPN settings")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d VPN settings", len(listRes))
}

// Test_device_settings_VPNSettingsAPIService_GetByID tests getting VPN settings by ID
func Test_device_settings_VPNSettingsAPIService_GetByID(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	// List to get an existing ID (singleton - always exists)
	listRes, _, errList := client.VPNSettingsAPI.ListVPNSettings(context.Background()).Folder("Prisma Access").Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list VPN settings")

	if len(listRes) == 0 {
		t.Skip("No VPN settings found to test GetByID")
	}

	existingID := *listRes[0].Id
	t.Logf("Testing GetByID with ID: %s", existingID)

	getRes, httpResGet, errGet := client.VPNSettingsAPI.GetVPNSettingsByID(context.Background(), existingID).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	require.NoError(t, errGet, "Failed to get VPN settings by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, existingID, *getRes.Id, "ID should match requested ID")
	t.Logf("Successfully retrieved VPN settings with ID: %s", *getRes.Id)
}

// Test_device_settings_VPNSettingsAPIService_Update tests updating VPN settings
// Uses list to find the existing singleton, updates it in place with a no-op update
func Test_device_settings_VPNSettingsAPIService_Update(t *testing.T) {
	client := SetupDeviceSettingsTestClient(t)

	// List to get the existing singleton
	listRes, _, errList := client.VPNSettingsAPI.ListVPNSettings(context.Background()).Folder("Prisma Access").Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list VPN settings")

	if len(listRes) == 0 {
		t.Skip("No VPN settings found to test Update")
	}

	existingID := *listRes[0].Id
	originalSettings := listRes[0]
	t.Logf("Testing Update with ID: %s", existingID)

	// Pass the original back as a no-op update to verify the API works
	updatedSettings := device_settings.VpnSettings{
		Id:     originalSettings.Id,
		Folder: common.StringPtr("Prisma Access"),
		Vpn:    originalSettings.Vpn,
	}

	updateRes, httpResUpdate, errUpdate := client.VPNSettingsAPI.UpdateVPNSettingsByID(context.Background(), existingID).VpnSettings(updatedSettings).Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	require.NoError(t, errUpdate, "Failed to update VPN settings")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, existingID, *updateRes.Id, "ID should match")
	t.Logf("Successfully updated VPN settings with ID: %s", *updateRes.Id)
}
