/*
 /*
 /*
 * Deployment Service Testing
 *
 * ServiceConnectionsAPIService
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

const (
	Folder = "Service Connections"
)

// --- SERVICE CONNECTIONS TEST SUITE ---

func createSCTestIPsecTunnelAndDeps(t *testing.T, suffix string) (string, func()) {
	netSvcClient := ntest.SetupNetworkSvcTestClient(t)

	// 1. Create IKE Crypto Profile
	ikeCryptoProfileName := "test-crypto-for-" + suffix
	ikeCryptoProfileID := ntest.CreateTestIKECryptoProfile(t, netSvcClient, ikeCryptoProfileName, Folder)
	t.Logf("Created dependency: IKE Crypto Profile '%s' with id '%s'", ikeCryptoProfileName, ikeCryptoProfileID)

	// 2. Create IKE Gateway with IKE Crypto Profile inside
	gatewayName := "test-gw-for-" + suffix
	gatewayID := ntest.CreateTestIkeGateway(t, netSvcClient, gatewayName, ikeCryptoProfileName, Folder)
	t.Logf("Created dependency: IKE Gateway '%s' with id '%s'", gatewayName, gatewayID)

	// 3. Create IPsec Tunnel with IKE Gateway inside
	ipsecTunnelName := "test-tunnel-for-" + suffix
	ipsecTunnelID := ntest.CreateTestIPSecTunnel(t, netSvcClient, ipsecTunnelName, gatewayName, Folder)
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

// Test_deployment_services_ServiceConnectionsAPIService_Create tests the creation of a Service Connection.
func Test_deployment_services_ServiceConnectionsAPIService_Create(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies (IPsec Tunnel, Gateway, Crypto Profile)
	tunnelName, cleanupDeps := createSCTestIPsecTunnelAndDeps(t, "sc-"+randomSuffix)
	defer cleanupDeps()

	// Create a valid Service Connection object with unique name
	scName := "test-tunnel-for-sc-" + randomSuffix
	sc := deployment_services.ServiceConnections{
		Name:        scName,
		IpsecTunnel: tunnelName,      // Required dependency
		Region:      "us-central1-a", // Required field
	}

	// Set optional fields
	sc.SetOnboardingType("classic")
	// sc.SetFolder("Service Connections")
	sc.SetSubnets([]string{"10.0.0.0/24", "10.0.1.0/24"})
	sc.SetSourceNat(true)

	fmt.Printf("Attempting to create Service Connection with name: %s\n", sc.Name)

	// Make the create request to the API
	req := depSvcClient.ServiceConnectionsAPI.CreateServiceConnections(context.Background()).ServiceConnections(sc)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create Service Connection")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotEmpty(t, res.Id, "Created SC should have an ID")

	createdSCID := res.Id

	// Defer the cleanup of the created Service Connection.
	defer func() {
		t.Logf("Cleaning up Service Connection with ID: %s", createdSCID)
		_, errDel := depSvcClient.ServiceConnectionsAPI.DeleteServiceConnectionsByID(context.Background(), createdSCID).Execute()
		require.NoError(t, errDel, "Failed to delete Service Connection during cleanup")
	}()

	// Assert response object properties
	assert.Equal(t, scName, res.Name, "Created SC name should match")
	assert.Equal(t, tunnelName, res.IpsecTunnel, "IPSec Tunnel should match")
	assert.Equal(t, "us-central1-a", res.Region, "Region should match")
	assert.Equal(t, true, res.GetSourceNat(), "Source NAT should match input")
	assert.Contains(t, res.Subnets, "10.0.0.0/24", "Subnets should contain the added entry")

	t.Logf("Successfully created and validated Service Connection: %s with ID: %s", sc.Name, createdSCID)
}

// Test_deployment_services_ServiceConnectionsAPIService_GetByID tests retrieving a SC by its ID.
func Test_deployment_services_ServiceConnectionsAPIService_GetByID(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies
	tunnelName, cleanupDeps := createSCTestIPsecTunnelAndDeps(t, "sc-"+randomSuffix)
	defer cleanupDeps()

	// Create a Service Connection to retrieve.
	scName := "test-sc-" + randomSuffix
	sc := deployment_services.ServiceConnections{
		Name:        scName,
		IpsecTunnel: tunnelName,
		Region:      "ap-southeast-2",
	}

	createRes, _, err := depSvcClient.ServiceConnectionsAPI.CreateServiceConnections(context.Background()).ServiceConnections(sc).Execute()
	require.NoError(t, err, "Failed to create SC for get test")
	createdSCID := createRes.Id
	defer func() {
		depSvcClient.ServiceConnectionsAPI.DeleteServiceConnectionsByID(context.Background(), createdSCID).Execute()
	}()

	// Test Get by ID operation
	getRes, httpResGet, errGet := depSvcClient.ServiceConnectionsAPI.GetServiceConnectionsByID(context.Background(), createdSCID).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful
	require.NoError(t, errGet, "Failed to get Service Connection by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Assert key properties
	assert.Equal(t, scName, getRes.Name, "SC name should match")
	assert.Equal(t, createdSCID, getRes.Id, "SC ID should match")
	assert.Equal(t, tunnelName, getRes.IpsecTunnel, "IPSec Tunnel should match")

	t.Logf("Successfully retrieved SC: %s", getRes.Name)
}

// // Test_deployment_services_ServiceConnectionsAPIService_Update tests updating an existing SC.
func Test_deployment_services_ServiceConnectionsAPIService_Update(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies for initial creation and for the update (if updating the tunnel).
	tunnelName, cleanupDeps := createSCTestIPsecTunnelAndDeps(t, "sc-"+randomSuffix)
	defer cleanupDeps()

	// Create a Service Connection to update.
	scName := "test-sc-" + randomSuffix
	sc := deployment_services.ServiceConnections{
		Name:        scName,
		IpsecTunnel: tunnelName,
		Region:      "eu-west-1",
		SourceNat:   common.BoolPtr(true),
		Subnets:     []string{"10.10.0.0/24"},
	}

	createRes, _, err := depSvcClient.ServiceConnectionsAPI.CreateServiceConnections(context.Background()).ServiceConnections(sc).Execute()
	require.NoError(t, err, "Failed to create SC for update test")
	createdSCID := createRes.Id
	defer func() {
		depSvcClient.ServiceConnectionsAPI.DeleteServiceConnectionsByID(context.Background(), createdSCID).Execute()
	}()

	// Define the update payload.
	updatedSC := deployment_services.ServiceConnections{
		Name:        scName,
		IpsecTunnel: tunnelName, // Keep same tunnel for simplicity
		Region:      "eu-west-1",
	}
	updatedSC.SetSourceNat(false)                                    // Updated boolean value
	updatedSC.SetSubnets([]string{"10.10.0.0/24", "192.168.0.0/24"}) // Added a subnet

	// Execute Update
	reqUpdate := depSvcClient.ServiceConnectionsAPI.UpdateServiceConnectionsByID(context.Background(), createdSCID).ServiceConnections(updatedSC)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update Service Connection")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Assert updated properties
	assert.Equal(t, false, updateRes.GetSourceNat(), "Source NAT should be updated to false")
	assert.Len(t, updateRes.Subnets, 2, "Should now have two subnets")
	assert.Contains(t, updateRes.Subnets, "192.168.0.0/24", "Subnet should be added")

	t.Logf("Successfully updated SC: %s", scName)
}

// // Test_deployment_services_ServiceConnectionsAPIService_List tests listing SCs with folder filter.
func Test_deployment_services_ServiceConnectionsAPIService_List(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies
	tunnelName, cleanupDeps := createSCTestIPsecTunnelAndDeps(t, "sc-"+randomSuffix)
	defer cleanupDeps()

	// Create a unique SC to ensure it appears in the list.
	scName := "test-sc-" + randomSuffix
	scFolder := "Service Connections"
	sc := deployment_services.ServiceConnections{
		Name:        scName,
		IpsecTunnel: tunnelName,
		Region:      "ap-northeast-1",
	}

	createRes, _, err := depSvcClient.ServiceConnectionsAPI.CreateServiceConnections(context.Background()).ServiceConnections(sc).Execute()
	require.NoError(t, err, "Failed to create SC for list test")
	createdSCID := createRes.Id
	defer func() {
		depSvcClient.ServiceConnectionsAPI.DeleteServiceConnectionsByID(context.Background(), createdSCID).Execute()
	}()

	// Test List operation, filtering by folder (required)
	reqList := depSvcClient.ServiceConnectionsAPI.ListServiceConnections(context.Background()).Folder(scFolder).Limit(10000)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list Service Connections")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes)

	// Verify our created SC is in the list
	foundSC := false
	for _, scRes := range listRes.Data {
		if scRes.Name == scName {
			foundSC = true
			assert.Equal(t, "ap-northeast-1", scRes.Region, "Region should match the created SC")
			break
		}
	}
	assert.True(t, foundSC, "Created Service Connection should be found in the list")

	t.Logf("Successfully listed SCs, found created SC: %s", scName)
}

// // Test_deployment_services_ServiceConnectionsAPIService_DeleteByID tests deleting a SC by its ID.
func Test_deployment_services_ServiceConnectionsAPIService_DeleteByID(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create dependencies
	tunnelName, cleanupDeps := createSCTestIPsecTunnelAndDeps(t, "sc-"+randomSuffix)
	defer cleanupDeps()

	// Create a Service Connection to delete.
	scName := "test-sc-" + randomSuffix
	sc := deployment_services.ServiceConnections{
		Name:        scName,
		IpsecTunnel: tunnelName,
		Region:      "us-west-1",
	}

	createRes, _, err := depSvcClient.ServiceConnectionsAPI.CreateServiceConnections(context.Background()).ServiceConnections(sc).Execute()
	require.NoError(t, err, "Failed to create SC for delete test")
	createdSCID := createRes.Id

	// Test Delete by ID operation
	_, errDel := depSvcClient.ServiceConnectionsAPI.DeleteServiceConnectionsByID(context.Background(), createdSCID).Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete Service Connection")

	t.Logf("Successfully deleted SC: %s", createdSCID)

	t.Logf("Verified SC deletion: %s", createdSCID)
}
