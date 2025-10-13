/*
Network Services Testing IPsecTunnelsAPIService
*/
package network_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// createTestIKEGateway creates an IKE Gateway and its dependencies.
// It returns the gateway name, gateway ID, and a cleanup function.
func createTestIKEGateway(t *testing.T, client *network_services.APIClient, suffix string) (string, string, func()) {
	// Create the IKE Crypto Profile dependency first using the shared helper.
	cryptoProfileName := "test-crypto-for-tunnel-" + suffix
	cryptoProfileID := CreateTestIKECryptoProfile(t, client, cryptoProfileName)

	// Create a valid IKE Gateway object.
	gatewayName := "test-gw-for-tunnel-" + suffix
	gateway := network_services.IkeGateways{
		Folder: common.StringPtr("Remote Networks"),
		Name:   gatewayName,
		Authentication: network_services.IkeGatewaysAuthentication{
			PreSharedKey: &network_services.IkeGatewaysAuthenticationPreSharedKey{
				Key: common.StringPtr("secret"),
			},
		},
		PeerAddress: network_services.IkeGatewaysPeerAddress{
			Ip: common.StringPtr("1.1.1.1"),
		},
		Protocol: network_services.IkeGatewaysProtocol{
			Ikev1: &network_services.IkeGatewaysProtocolIkev1{
				IkeCryptoProfile: common.StringPtr(cryptoProfileName),
			},
		},
	}

	req := client.IKEGatewaysAPI.CreateIKEGateways(context.Background()).IkeGateways(gateway)
	res, _, err := req.Execute()
	require.NoError(t, err, "Failed to create test IKE Gateway dependency")
	require.NotNil(t, res, "Test IKE Gateway create response should not be nil")
	t.Logf("Created test IKE Gateway '%s' with ID %s", gatewayName, *res.Id)

	cleanup := func() {
		delGwReq := client.IKEGatewaysAPI.DeleteIKEGatewaysByID(context.Background(), *res.Id)
		_, err := delGwReq.Execute()
		require.NoError(t, err, "Failed to delete test IKE Gateway '%s'", gatewayName)
		t.Logf("Deleted test IKE Gateway '%s'", gatewayName)

		// Call the shared cleanup helper.
		DeleteTestIKECryptoProfile(t, client, cryptoProfileID, cryptoProfileName)
	}

	return gatewayName, *res.Id, cleanup
}

// Test_networkservices_IPsecTunnelsAPIService_Create tests the creation of an IPsec tunnel object.
func Test_networkservices_IPsecTunnelsAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies.
	gatewayName, _, cleanupGw := createTestIKEGateway(t, client, randomSuffix)
	defer cleanupGw()

	// Create the IPsec tunnel object.
	createdTunnelName := "auto_ipsec_tunnel-" + randomSuffix
	tunnel := network_services.IpsecTunnels{
		Folder:                 common.StringPtr("Remote Networks"),
		Name:                   createdTunnelName,
		AntiReplay:             common.BoolPtr(true),
		CopyTos:                common.BoolPtr(false),
		EnableGreEncapsulation: common.BoolPtr(false),
		AutoKey: network_services.IpsecTunnelsAutoKey{
			IkeGateway: []network_services.IpsecTunnelsAutoKeyIkeGatewayInner{
				{
					Name: common.StringPtr(gatewayName),
				},
			},
			IpsecCryptoProfile: "PaloAlto-Networks-IPSec-Crypto",
		},
	}

	fmt.Printf("Creating IPsec tunnel with name: %s\n", tunnel.Name)
	req := client.IPsecTunnelsAPI.CreateIPsecTunnels(context.Background()).IpsecTunnels(tunnel)
	res, httpRes, err := req.Execute()

	// Defer cleanup for the IPsec tunnel.
	if res != nil && res.Id != nil {
		defer func() {
			delReq := client.IPsecTunnelsAPI.DeleteIPsecTunnelsByID(context.Background(), *res.Id)
			_, errDel := delReq.Execute()
			require.NoError(t, errDel, "Failed to delete IPsec tunnel during cleanup")
			t.Logf("Cleaned up IPsec tunnel: %s", *res.Id)
		}()
	}

	// Verify creation was successful.
	require.NoError(t, err, "Failed to create IPsec tunnel")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdTunnelName, res.Name)
}

// Test_networkservices_IPsecTunnelsAPIService_GetByID tests retrieving an IPsec tunnel by its ID.
func Test_networkservices_IPsecTunnelsAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies.
	gatewayName, _, cleanupGw := createTestIKEGateway(t, client, randomSuffix)
	defer cleanupGw()

	// Create an IPsec tunnel to retrieve.
	tunnelName := "tunnel-get-" + randomSuffix
	tunnel := network_services.IpsecTunnels{
		Folder: common.StringPtr("Remote Networks"),
		Name:   tunnelName,
		AutoKey: network_services.IpsecTunnelsAutoKey{
			IkeGateway: []network_services.IpsecTunnelsAutoKeyIkeGatewayInner{
				{
					Name: common.StringPtr(gatewayName),
				},
			},
			IpsecCryptoProfile: "PaloAlto-Networks-IPSec-Crypto",
		},
	}
	createRes, _, err := client.IPsecTunnelsAPI.CreateIPsecTunnels(context.Background()).IpsecTunnels(tunnel).Execute()
	require.NoError(t, err, "Failed to create tunnel for get test")
	createdTunnelID := *createRes.Id
	defer func() {
		client.IPsecTunnelsAPI.DeleteIPsecTunnelsByID(context.Background(), createdTunnelID).Execute()
	}()

	// Test Get by ID operation.
	reqGetById := client.IPsecTunnelsAPI.GetIPsecTunnelsByID(context.Background(), createdTunnelID)
	getRes, httpResGet, errGet := reqGetById.Execute()

	// Verify the get operation was successful.
	require.NoError(t, errGet)
	assert.Equal(t, 200, httpResGet.StatusCode)
	require.NotNil(t, getRes)
	assert.Equal(t, createdTunnelID, *getRes.Id)
}

// Test_networkservices_IPsecTunnelsAPIService_Update tests updating an existing IPsec tunnel.
func Test_networkservices_IPsecTunnelsAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies.
	gatewayName, _, cleanupGw := createTestIKEGateway(t, client, randomSuffix)
	defer cleanupGw()

	// Create an IPsec tunnel to update.
	tunnelName := "tunnel-update-" + randomSuffix
	tunnel := network_services.IpsecTunnels{
		Folder: common.StringPtr("Remote Networks"),
		Name:   tunnelName,
		AutoKey: network_services.IpsecTunnelsAutoKey{
			IkeGateway: []network_services.IpsecTunnelsAutoKeyIkeGatewayInner{
				{
					Name: common.StringPtr(gatewayName),
				},
			},
			IpsecCryptoProfile: "PaloAlto-Networks-IPSec-Crypto",
		},
	}
	createRes, _, err := client.IPsecTunnelsAPI.CreateIPsecTunnels(context.Background()).IpsecTunnels(tunnel).Execute()
	require.NoError(t, err, "Failed to create tunnel for update test")
	createdTunnelID := *createRes.Id
	defer func() {
		client.IPsecTunnelsAPI.DeleteIPsecTunnelsByID(context.Background(), createdTunnelID).Execute()
	}()

	// Test Update operation.
	updatedTunnel := network_services.IpsecTunnels{
		Name:       tunnelName,
		CopyTos:    common.BoolPtr(true),
		AntiReplay: common.BoolPtr(false),
		AutoKey:    tunnel.AutoKey, // Must include auto_key config
	}
	reqUpdate := client.IPsecTunnelsAPI.UpdateIPsecTunnelsByID(context.Background(), createdTunnelID).IpsecTunnels(updatedTunnel)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()

	// Verify the update operation was successful.
	require.NoError(t, errUpdate)
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, true, *updateRes.CopyTos)
	assert.Equal(t, false, *updateRes.AntiReplay)
}

// Test_networkservices_IPsecTunnelsAPIService_List tests listing IPsec tunnels.
func Test_networkservices_IPsecTunnelsAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies.
	gatewayName, _, cleanupGw := createTestIKEGateway(t, client, randomSuffix)
	defer cleanupGw()

	// Create an IPsec tunnel to list.
	tunnelName := "tunnel-list-" + randomSuffix
	tunnel := network_services.IpsecTunnels{
		Folder: common.StringPtr("Remote Networks"),
		Name:   tunnelName,
		AutoKey: network_services.IpsecTunnelsAutoKey{
			IkeGateway: []network_services.IpsecTunnelsAutoKeyIkeGatewayInner{
				{
					Name: common.StringPtr(gatewayName),
				},
			},
			IpsecCryptoProfile: "PaloAlto-Networks-IPSec-Crypto",
		},
	}
	createRes, _, err := client.IPsecTunnelsAPI.CreateIPsecTunnels(context.Background()).IpsecTunnels(tunnel).Execute()
	require.NoError(t, err, "Failed to create tunnel for list test")
	createdTunnelID := *createRes.Id
	defer func() {
		client.IPsecTunnelsAPI.DeleteIPsecTunnelsByID(context.Background(), createdTunnelID).Execute()
	}()

	// Test List operation.
	req := client.IPsecTunnelsAPI.ListIPsecTunnels(context.Background()).Folder("Remote Networks")
	listRes, httpRes, err := req.Execute()

	// Verify the list operation was successful.
	require.NoError(t, err, "List request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "The response from list should not be nil")
	assert.True(t, len(listRes.Data) > 0, "Should have at least one tunnel in the list")

	// Verify our created gateway is in the list.
	found := false
	for _, ipsecTunnel := range listRes.Data {
		if ipsecTunnel.Name == tunnelName {
			found = true
			break
		}
	}
	assert.True(t, found, "Created tunnel should be found in the list")
	t.Logf("Successfully listed IPSec Tunnels and found created tunnel: %s", tunnelName)
}

// Test_networkservices_IPsecTunnelsAPIService_DeleteByID tests deleting an IPsec tunnel.
func Test_networkservices_IPsecTunnelsAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies.
	gatewayName, _, cleanupGw := createTestIKEGateway(t, client, randomSuffix)
	defer cleanupGw()

	// Create an IPsec tunnel to delete.
	tunnelName := "tunnel-delete-" + randomSuffix
	tunnel := network_services.IpsecTunnels{
		Folder: common.StringPtr("Remote Networks"),
		Name:   tunnelName,
		AutoKey: network_services.IpsecTunnelsAutoKey{
			IkeGateway: []network_services.IpsecTunnelsAutoKeyIkeGatewayInner{
				{
					Name: common.StringPtr(gatewayName),
				},
			},
			IpsecCryptoProfile: "PaloAlto-Networks-IPSec-Crypto",
		},
	}
	createRes, _, err := client.IPsecTunnelsAPI.CreateIPsecTunnels(context.Background()).IpsecTunnels(tunnel).Execute()
	require.NoError(t, err, "Failed to create tunnel for delete test")
	createdTunnelID := *createRes.Id

	// Test Delete operation.
	reqDel := client.IPsecTunnelsAPI.DeleteIPsecTunnelsByID(context.Background(), createdTunnelID)
	httpResDel, errDel := reqDel.Execute()

	// Verify the delete was successful.
	require.NoError(t, errDel)
	assert.Equal(t, 200, httpResDel.StatusCode)
}
