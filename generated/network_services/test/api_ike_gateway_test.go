/*
 * Network Services Testing
 *
 * IkeGatewaysAPIService
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

// Test_networkservices_IkeGatewaysAPIService_Create tests the creation of an IKE Gateway.
func Test_networkservices_IkeGatewaysAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupNetworkSvcTestClient(t)

	// Create the IKE Crypto Profile dependency first.
	randomSuffix := common.GenerateRandomString(6)
	cryptoProfileName := "test-crypto-create-" + randomSuffix
	cryptoProfileID := CreateTestIKECryptoProfile(t, client, cryptoProfileName)
	defer DeleteTestIKECryptoProfile(t, client, cryptoProfileID, cryptoProfileName)

	// Create a valid IKE Gateway object with a unique name.
	gatewayName := "auto_ike_gw_create_" + randomSuffix
	gateway := CreateIkeGatewayTestObject(gatewayName, cryptoProfileName)

	fmt.Printf("Attempting to create IKE Gateway with name: %s\n", gateway.Name)

	// Make the create request to the API.
	req := client.IKEGatewaysAPI.CreateIKEGateways(context.Background()).IkeGateways(gateway)
	res, httpRes, err := req.Execute()

	// Defer cleanup for the IKE Gateway.
	if res != nil && res.Id != nil {
		defer func() {
			t.Logf("Cleaning up IKE Gateway with ID: %s", *res.Id)
			delReq := client.IKEGatewaysAPI.DeleteIKEGatewaysByID(context.Background(), *res.Id)
			_, errDel := delReq.Execute()
			require.NoError(t, errDel, "Failed to delete IKE Gateway during cleanup")
		}()
	}

	// Verify the request was successful.
	require.NoError(t, err, "Create request should not return an error")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "The response from create should not be nil")
	assert.Equal(t, gatewayName, res.Name, "The name of the created gateway should match")
	assert.NotEmpty(t, *res.Id, "The ID of the created gateway should not be empty")

	t.Logf("Successfully created IKE Gateway with ID: %s", *res.Id)
}

// Test_networkservices_IkeGatewaysAPIService_GetByID tests the retrieval of an IKE Gateway by its ID.
func Test_networkservices_IkeGatewaysAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create the IKE Crypto Profile dependency first.
	randomSuffix := common.GenerateRandomString(6)
	cryptoProfileName := "test-crypto-get-" + randomSuffix
	cryptoProfileID := CreateTestIKECryptoProfile(t, client, cryptoProfileName)
	defer DeleteTestIKECryptoProfile(t, client, cryptoProfileID, cryptoProfileName)

	// Create a gateway to retrieve.
	gatewayName := "auto_ike_gw_get_" + randomSuffix
	gateway := CreateIkeGatewayTestObject(gatewayName, cryptoProfileName)
	createRes, _, err := client.IKEGatewaysAPI.CreateIKEGateways(context.Background()).IkeGateways(gateway).Execute()
	require.NoError(t, err, "Failed to create gateway for get test")
	createdGatewayID := *createRes.Id

	// Defer cleanup for the IKE Gateway.
	defer func() {
		t.Logf("Cleaning up IKE Gateway with ID: %s", createdGatewayID)
		_, errDel := client.IKEGatewaysAPI.DeleteIKEGatewaysByID(context.Background(), createdGatewayID).Execute()
		require.NoError(t, errDel, "Failed to delete gateway during cleanup")
	}()

	t.Logf("Created IKE Gateway for Get test with ID: %s", createdGatewayID)

	// Test the Get by ID operation.
	fmt.Printf("Attempting to get IKE Gateway with ID: %s\n", createdGatewayID)
	req := client.IKEGatewaysAPI.GetIKEGatewaysByID(context.Background(), createdGatewayID)
	getRes, httpRes, err := req.Execute()

	// Verify the retrieval was successful.
	require.NoError(t, err, "Get by ID request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "The response from get should not be nil")
	assert.Equal(t, createdGatewayID, *getRes.Id, "The ID of the retrieved gateway should match")
	assert.Equal(t, gatewayName, getRes.Name, "The name of the retrieved gateway should match")
}

// Test_networkservices_IkeGatewaysAPIService_Update tests updating an IKE Gateway.
func Test_networkservices_IkeGatewaysAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create the IKE Crypto Profile dependency first.
	randomSuffix := common.GenerateRandomString(6)
	cryptoProfileName := "test-crypto-update-" + randomSuffix
	cryptoProfileID := CreateTestIKECryptoProfile(t, client, cryptoProfileName)
	defer DeleteTestIKECryptoProfile(t, client, cryptoProfileID, cryptoProfileName)

	// Create a gateway to update.
	gatewayName := "auto_ike_gw_update_" + randomSuffix
	gateway := CreateIkeGatewayTestObject(gatewayName, cryptoProfileName)
	createRes, _, err := client.IKEGatewaysAPI.CreateIKEGateways(context.Background()).IkeGateways(gateway).Execute()
	require.NoError(t, err, "Failed to create gateway for update test")
	createdGatewayID := *createRes.Id

	// Defer cleanup for the IKE Gateway.
	defer func() {
		t.Logf("Cleaning up IKE Gateway with ID: %s", createdGatewayID)
		_, errDel := client.IKEGatewaysAPI.DeleteIKEGatewaysByID(context.Background(), createdGatewayID).Execute()
		require.NoError(t, errDel, "Failed to delete gateway during cleanup")
	}()

	t.Logf("Created IKE Gateway for Update test with ID: %s", createdGatewayID)

	// Update the gateway object.
	updatedGateway := network_services.IkeGateways{
		Name: gatewayName,
		Authentication: network_services.IkeGatewaysAuthentication{
			PreSharedKey: &network_services.IkeGatewaysAuthenticationPreSharedKey{
				Key: common.StringPtr("123456"),
			},
		},
		PeerAddress: network_services.IkeGatewaysPeerAddress{
			Ip: common.StringPtr("8.8.8.8"),
		},
	}

	// Test the Update by ID operation.
	fmt.Printf("Attempting to update IKE Gateway with ID: %s\n", createdGatewayID)
	reqUpdate := client.IKEGatewaysAPI.UpdateIKEGatewaysByID(context.Background(), createdGatewayID).IkeGateways(updatedGateway)
	updateRes, httpRes, err := reqUpdate.Execute()

	// Verify the update was successful.
	require.NoError(t, err, "Update request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "The response from update should not be nil")
}

// Test_networkservices_IkeGatewaysAPIService_List tests listing IKE Gateways.
func Test_networkservices_IkeGatewaysAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create the IKE Crypto Profile dependency first.
	randomSuffix := common.GenerateRandomString(6)
	cryptoProfileName := "test-crypto-list-" + randomSuffix
	cryptoProfileID := CreateTestIKECryptoProfile(t, client, cryptoProfileName)
	defer DeleteTestIKECryptoProfile(t, client, cryptoProfileID, cryptoProfileName)

	// Create a gateway to ensure it appears in the list.
	gatewayName := "auto_ike_gw_list-" + randomSuffix
	gateway := CreateIkeGatewayTestObject(gatewayName, cryptoProfileName)
	createRes, _, err := client.IKEGatewaysAPI.CreateIKEGateways(context.Background()).IkeGateways(gateway).Execute()
	require.NoError(t, err, "Failed to create gateway for list test")
	createdGatewayID := *createRes.Id

	// Defer cleanup for the IKE Gateway.
	defer func() {
		t.Logf("Cleaning up IKE Gateway with ID: %s", createdGatewayID)
		_, errDel := client.IKEGatewaysAPI.DeleteIKEGatewaysByID(context.Background(), createdGatewayID).Execute()
		require.NoError(t, errDel, "Failed to delete gateway during cleanup")
	}()

	// Test the List operation.
	fmt.Println("Attempting to list IKE Gateways")
	req := client.IKEGatewaysAPI.ListIKEGateways(context.Background()).Folder("Remote Networks")
	listRes, httpRes, err := req.Execute()

	// Verify the list operation was successful.
	require.NoError(t, err, "List request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "The response from list should not be nil")
	assert.True(t, len(listRes.Data) > 0, "Should have at least one gateway in the list")

	// Verify our created gateway is in the list.
	found := false
	for _, gw := range listRes.Data {
		if gw.Name == gatewayName {
			found = true
			break
		}
	}
	assert.True(t, found, "Created gateway should be found in the list")
	t.Logf("Successfully listed IKE Gateways and found created gateway: %s", gatewayName)
}

// Test_networkservices_IkeGatewaysAPIService_DeleteByID tests deleting an IKE Gateway.
func Test_networkservices_IkeGatewaysAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create the IKE Crypto Profile dependency first.
	randomSuffix := common.GenerateRandomString(6)
	cryptoProfileName := "test-crypto-delete-" + randomSuffix
	cryptoProfileID := CreateTestIKECryptoProfile(t, client, cryptoProfileName)
	defer DeleteTestIKECryptoProfile(t, client, cryptoProfileID, cryptoProfileName)

	// Create a gateway to delete.
	gatewayName := "auto_ike_gw_delete-" + randomSuffix
	gateway := CreateIkeGatewayTestObject(gatewayName, cryptoProfileName)
	createRes, _, err := client.IKEGatewaysAPI.CreateIKEGateways(context.Background()).IkeGateways(gateway).Execute()
	require.NoError(t, err, "Failed to create gateway for delete test")
	createdGatewayID := *createRes.Id

	t.Logf("Created IKE Gateway for Delete test with ID: %s", createdGatewayID)

	// Test the Delete by ID operation.
	fmt.Printf("Attempting to delete IKE Gateway with ID: %s\n", createdGatewayID)
	reqDel := client.IKEGatewaysAPI.DeleteIKEGatewaysByID(context.Background(), createdGatewayID)
	httpResDel, errDel := reqDel.Execute()

	// Verify the delete operation was successful.
	require.NoError(t, errDel, "Failed to delete gateway")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
	t.Logf("Successfully deleted IKE Gateway: %s", createdGatewayID)
}
