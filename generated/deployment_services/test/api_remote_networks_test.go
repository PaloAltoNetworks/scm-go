/*
 * Network Deployment Testing
 *
 * RemoteNetworksAPIService
 */

package deployment_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
	ntest "github.com/paloaltonetworks/scm-go/generated/network_services/test"
)

// createTestIPsecTunnelAndDeps creates the full chain of dependencies required for a Remote Network test.
// It returns the name of the created IPsec Tunnel and a function to clean up all created dependencies.
func createTestIPsecTunnelAndDeps(t *testing.T, suffix string) (string, func()) {
	netSvcClient := ntest.SetupNetworkSvcTestClient(t)

	// 1. Create IKE Crypto Profile
	ikeCryptoProfileName := "test-crypto-for-rn-" + suffix
	ikeCryptoProfileID := ntest.CreateTestIKECryptoProfile(t, netSvcClient, ikeCryptoProfileName)
	t.Logf("Created dependency: IKE Crypto Profile '%s' with id '%s'", ikeCryptoProfileName, ikeCryptoProfileID)

	// 2. Create IKE Gateway with IKE Crypto Profile inside
	gatewayName := "test-gw-for-rn-" + suffix
	gatewayID := ntest.CreateTestIkeGateway(t, netSvcClient, gatewayName, ikeCryptoProfileName)
	t.Logf("Created dependency: IKE Gateway '%s' with id '%s'", gatewayName, gatewayID)

	// 3. Create IPsec Tunnel with IKE Gateway inside
	ipsecTunnelName := "test-tunnel-for-rn-" + suffix
	ipsecTunnelID := ntest.CreateTestIPSecTunnel(t, netSvcClient, ipsecTunnelName, gatewayName)
	t.Logf("Created dependency: IPsec Tunnel '%s' with id '%s'", ipsecTunnelName, ipsecTunnelID)

	// Return the final tunnel name and a cleanup function that deletes all three resources.
	cleanup := func() {
		t.Logf("Cleaning up dependency: IPsec Tunnel '%s'", ipsecTunnelName)
		_, err := netSvcClient.IPsecTunnelsAPI.DeleteIPsecTunnelsByID(context.Background(), ipsecTunnelID).Execute()
		require.NoError(t, err, "Failed to delete IPsec Tunnel")

		t.Logf("Cleaning up dependency: IKE Gateway '%s'", gatewayName)
		_, err = netSvcClient.IKEGatewaysAPI.DeleteIKEGatewaysByID(context.Background(), gatewayID).Execute()
		require.NoError(t, err, "Failed to delete IKE Gateway")

		t.Logf("Cleaning up dependency: IKE Crypto Profile '%s'", ikeCryptoProfileName)
		_, err = netSvcClient.IKECryptoProfilesAPI.DeleteIKECryptoProfilesByID(context.Background(), ikeCryptoProfileID).Execute()
		require.NoError(t, err, "Failed to delete IKE Crypto Profile")
	}

	return ipsecTunnelName, cleanup
}

// Test_deployment_services_RemoteNetworksAPIService_Create tests the creation of a Remote Network.
func Test_deployment_services_RemoteNetworksAPIService_Create(t *testing.T) {
	// Setup the authenticated clients for both services.
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies.
	tunnelName, cleanupDeps := createTestIPsecTunnelAndDeps(t, randomSuffix)
	defer cleanupDeps()

	// Create a valid Remote Network object with a unique name.
	networkName := "test-rn-create-" + randomSuffix
	network := deployment_services.RemoteNetworks{
		Name:        networkName,
		Folder:      "Remote Networks",
		SpnName:     common.StringPtr("us-west-dakota"),
		LicenseType: "FWAAS-AGGREGATE",
		Region:      "us-west-2",
		IpsecTunnel: common.StringPtr(tunnelName),
		Subnets:     []string{"192.168.1.0/24"},
	}

	fmt.Printf("Attempting to create Remote Network with name: %s\n", network.Name)

	// Make the create request to the API.
	req := depSvcClient.RemoteNetworksAPI.CreateRemoteNetworks(context.Background()).RemoteNetworks(network)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create Remote Network")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotEmpty(t, res.Id, "Created remote network should have an ID")

	createdNetworkID := res.Id

	// Defer the cleanup of the created remote network.
	defer func() {
		t.Logf("Cleaning up Remote Network with ID: %s", createdNetworkID)
		_, errDel := depSvcClient.RemoteNetworksAPI.DeleteRemoteNetworksByID(context.Background(), createdNetworkID).Execute()
		require.NoError(t, errDel, "Failed to delete remote network during cleanup")
	}()

	// Assert response object properties.
	assert.Equal(t, networkName, res.Name, "Created remote network name should match")
	assert.Equal(t, "us-west-2", res.Region, "Region should match")
	require.NotNil(t, res.IpsecTunnel, "IPsec tunnel should not be nil")
	assert.Equal(t, tunnelName, *res.IpsecTunnel, "IPsec tunnel name should match")
	t.Logf("Successfully created and validated Remote Network: %s with ID: %s", network.Name, createdNetworkID)
}

// Test_deployment_services_RemoteNetworksAPIService_GetByID tests retrieving a remote network by its ID.
func Test_deployment_services_RemoteNetworksAPIService_GetByID(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies.
	tunnelName, cleanupDeps := createTestIPsecTunnelAndDeps(t, randomSuffix)
	defer cleanupDeps()

	// Create a remote network to retrieve.
	networkName := "test-rn-get-" + randomSuffix
	network := deployment_services.RemoteNetworks{
		Name:        networkName,
		Folder:      "Remote Networks",
		SpnName:     common.StringPtr("us-west-dakota"),
		LicenseType: "FWAAS-AGGREGATE",
		Region:      "us-east-1",
		IpsecTunnel: common.StringPtr(tunnelName),
		Subnets:     []string{"192.168.2.0/24"},
	}

	createRes, _, err := depSvcClient.RemoteNetworksAPI.CreateRemoteNetworks(context.Background()).RemoteNetworks(network).Execute()
	require.NoError(t, err, "Failed to create remote network for get test")
	createdNetworkID := createRes.Id
	defer func() {
		depSvcClient.RemoteNetworksAPI.DeleteRemoteNetworksByID(context.Background(), createdNetworkID).Execute()
	}()

	// Test Get by ID operation.
	getRes, httpResGet, errGet := depSvcClient.RemoteNetworksAPI.GetRemoteNetworksByID(context.Background(), createdNetworkID).Execute()
	require.NoError(t, errGet, "Failed to get remote network by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, networkName, getRes.Name)
	assert.Equal(t, createRes.Id, getRes.Id)
}

// Test_deployment_services_RemoteNetworksAPIService_Update tests updating an existing remote network.
func Test_deployment_services_RemoteNetworksAPIService_Update(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies for initial creation and for the update.
	initialTunnelName, cleanupInitialDeps := createTestIPsecTunnelAndDeps(t, "init-"+randomSuffix)
	defer cleanupInitialDeps()
	updatedTunnelName, cleanupUpdatedDeps := createTestIPsecTunnelAndDeps(t, "upd-"+randomSuffix)
	defer cleanupUpdatedDeps()

	// Create a remote network to update.
	networkName := "test-rn-update-" + randomSuffix
	network := deployment_services.RemoteNetworks{
		Name:        networkName,
		Folder:      "Remote Networks",
		SpnName:     common.StringPtr("us-west-dakota"),
		LicenseType: "FWAAS-AGGREGATE",
		Region:      "us-central1-a",
		IpsecTunnel: common.StringPtr(initialTunnelName),
	}
	createRes, _, err := depSvcClient.RemoteNetworksAPI.CreateRemoteNetworks(context.Background()).RemoteNetworks(network).Execute()
	require.NoError(t, err, "Failed to create remote network for update test")
	createdNetworkID := createRes.Id
	defer func() {
		depSvcClient.RemoteNetworksAPI.DeleteRemoteNetworksByID(context.Background(), createdNetworkID).Execute()
	}()

	// Define the update payload.
	updatedNetwork := deployment_services.RemoteNetworks{
		Name:        networkName,
		Folder:      "Remote Networks",
		SpnName:     common.StringPtr("us-west-dakota"),
		LicenseType: "FWAAS-AGGREGATE",
		Region:      "us-central1-a",
		IpsecTunnel: common.StringPtr(updatedTunnelName),
	}

	updateRes, httpResUpdate, errUpdate := depSvcClient.RemoteNetworksAPI.UpdateRemoteNetworksByID(context.Background(), createdNetworkID).RemoteNetworks(updatedNetwork).Execute()
	require.NoError(t, errUpdate, "Failed to update remote network")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, updatedTunnelName, *updateRes.IpsecTunnel, "IPsec tunnel should be updated")
}

// Test_deployment_services_RemoteNetworksAPIService_List tests listing Remote Networks.
func Test_deployment_services_RemoteNetworksAPIService_List(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies.
	tunnelName, cleanupDeps := createTestIPsecTunnelAndDeps(t, randomSuffix)
	defer cleanupDeps()

	// Create a remote network to ensure it appears in the list.
	networkName := "test-rn-list-" + randomSuffix
	network := deployment_services.RemoteNetworks{
		Name:        networkName,
		Folder:      "Remote Networks",
		SpnName:     common.StringPtr("us-west-dakota"),
		LicenseType: "FWAAS-AGGREGATE",
		Region:      "us-west-1",
		IpsecTunnel: common.StringPtr(tunnelName),
	}
	createRes, _, err := depSvcClient.RemoteNetworksAPI.CreateRemoteNetworks(context.Background()).RemoteNetworks(network).Execute()
	require.NoError(t, err, "Failed to create remote network for list test")
	createdNetworkID := createRes.Id
	defer func() {
		depSvcClient.RemoteNetworksAPI.DeleteRemoteNetworksByID(context.Background(), createdNetworkID).Execute()
	}()

	// Test List operation.
	listRes, httpResList, errList := depSvcClient.RemoteNetworksAPI.ListRemoteNetworks(context.Background()).Folder("Remote Networks").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list remote networks")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	// Verify our created remote network is in the list.
	foundNetwork := false
	for _, rn := range listRes.Data {
		if rn.Name == networkName {
			foundNetwork = true
			break
		}
	}
	assert.True(t, foundNetwork, "Created remote network should be found in the list")
}

// Test_deployment_services_RemoteNetworksAPIService_DeleteByID tests deleting a remote network by ID.
func Test_deployment_services_RemoteNetworksAPIService_DeleteByID(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies.
	tunnelName, cleanupDeps := createTestIPsecTunnelAndDeps(t, randomSuffix)
	defer cleanupDeps()

	// Create a remote network to delete.
	networkName := "test-rn-delete-" + randomSuffix
	network := deployment_services.RemoteNetworks{
		Name:        networkName,
		Folder:      "Remote Networks",
		SpnName:     common.StringPtr("us-west-dakota"),
		LicenseType: "FWAAS-AGGREGATE",
		Region:      "us-west-2",
		IpsecTunnel: common.StringPtr(tunnelName),
	}
	createRes, _, err := depSvcClient.RemoteNetworksAPI.CreateRemoteNetworks(context.Background()).RemoteNetworks(network).Execute()
	require.NoError(t, err, "Failed to create remote network for delete test")
	createdNetworkID := createRes.Id

	// Test Delete by ID operation.
	_, errDel := depSvcClient.RemoteNetworksAPI.DeleteRemoteNetworksByID(context.Background(), createdNetworkID).Execute()
	require.NoError(t, errDel, "Failed to delete remote network")
}
