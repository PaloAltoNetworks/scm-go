/*
Testing AutoVPNSettingsAPIService
Methods covered: Get, Update
Note: Auto VPN settings are singleton - only Get and Update are available.
*/
package network_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_network_services_AutoVPNSettingsAPIService_Get tests getting auto VPN settings
func Test_network_services_AutoVPNSettingsAPIService_Get(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	getRes, httpResGet, errGet := client.AutoVPNSettingsAPI.GetAutoVPNSettings(context.Background()).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	require.NoError(t, errGet, "Failed to get auto VPN settings")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	t.Logf("Successfully retrieved auto VPN settings")
}

// Test_network_services_AutoVPNSettingsAPIService_Update tests updating auto VPN settings
// Uses Get to retrieve existing settings, then performs a no-op update
func Test_network_services_AutoVPNSettingsAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Get existing settings
	getRes, _, errGet := client.AutoVPNSettingsAPI.GetAutoVPNSettings(context.Background()).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}
	require.NoError(t, errGet, "Failed to get auto VPN settings")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Perform no-op update with existing settings
	updateRes, httpResUpdate, errUpdate := client.AutoVPNSettingsAPI.UpdateAutoVPNSettings(context.Background()).AutoVpnSettings(*getRes).Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	require.NoError(t, errUpdate, "Failed to update auto VPN settings")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	t.Logf("Successfully updated auto VPN settings (no-op)")
}
