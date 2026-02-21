/*
Testing AutoVPNConfigPushAPIService
Methods covered: Create
Note: Config push triggers a push operation - test only verifies API call works.
*/
package network_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_network_services_AutoVPNConfigPushAPIService_Create tests creating an auto VPN config push
func Test_network_services_AutoVPNConfigPushAPIService_Create(t *testing.T) {
	t.Skip("Skipping - config push triggers an actual push operation and requires VPN infrastructure")

	client := SetupNetworkSvcTestClient(t)

	pushRes, httpRes, err := client.AutoVPNConfigPushAPI.CreateAutoVPNPushConfigs(context.Background()).Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create auto VPN config push")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	require.NotNil(t, pushRes, "Push response should not be nil")
	t.Logf("Successfully triggered auto VPN config push")
}
