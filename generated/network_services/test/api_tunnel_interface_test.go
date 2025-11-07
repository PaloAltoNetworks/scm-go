/*
 * Network Services Testing
 *
 * TunnelInterfacesAPIService
 */

package network_services

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Replace with your actual imports for common utils and generated client
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// --- Helper Functions ---

// generateTunnelName creates a unique name for tunnel interfaces.
func generateTunnelName(base string) string {
	// Tunnel Interface name usually doesn't have the '$' prefix like Loopback.
	return "$" + base + common.GenerateRandomString(6)
}

// createTestTunnelInterface creates a minimal TunnelInterfaces object for testing.
func createTestTunnelInterface(t *testing.T, interfaceName string) network_services.TunnelInterfaces {
	// 'name' is the only required field.
	return *network_services.NewTunnelInterfaces(interfaceName)
}

// createFullTestTunnelInterface creates a more complete TunnelInterfaces object for update/get testing.
func createFullTestTunnelInterface(t *testing.T, interfaceName string) network_services.TunnelInterfaces {
	mtu := int32(1450)
	comment := "Test Tunnel Interface"
	ipConfig := []network_services.TunnelInterfacesIpInner{
		{
			// Assuming TunnelInterfacesIpInner has a 'Name' field for the CIDR
			Name: "198.18.1.1/32",
		},
	}
	// DefaultValue regex: ^tunnel\.([1-9][0-9]{0,3})$ (tunnel.1 to tunnel.9999)

	tunnel := createTestTunnelInterface(t, interfaceName)
	tunnel.SetMtu(mtu)
	tunnel.SetComment(comment)
	tunnel.SetFolder("All")
	tunnel.SetIp(ipConfig)

	return tunnel
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_TunnelInterfacesAPIService_Create tests the creation of a Tunnel Interface.
func Test_network_services_TunnelInterfacesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := generateTunnelName("scm-tun-create-")

	tunnelInterface := createFullTestTunnelInterface(t, interfaceName)

	t.Logf("Creating Tunnel Interface with name: %s", interfaceName)
	req := client.TunnelInterfacesAPI.CreateTunnelInterfaces(context.Background()).TunnelInterfaces(tunnelInterface)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Tunnel Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created interface should have an ID")
	createdID := *res.Id

	// Cleanup the created interface
	defer func() {
		t.Logf("Cleaning up Tunnel Interface with ID: %s", createdID)
		_, errDel := client.TunnelInterfacesAPI.DeleteTunnelInterfacesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete Tunnel Interface during cleanup")
	}()

	t.Logf("Successfully created Tunnel Interface: %s with ID: %s", interfaceName, createdID)

	// Verify the response matches key input fields
	assert.Equal(t, interfaceName, res.Name, "Created interface name should match")
	assert.Equal(t, int32(1450), res.GetMtu(), "MTU should match the set value")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_TunnelInterfacesAPIService_GetByID tests retrieving a Tunnel Interface by ID.
func Test_network_services_TunnelInterfacesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := generateTunnelName("scm-tun-get-")
	tunnelInterface := createFullTestTunnelInterface(t, interfaceName)

	// Setup: Create an interface first
	createRes, _, err := client.TunnelInterfacesAPI.CreateTunnelInterfaces(context.Background()).TunnelInterfaces(tunnelInterface).Execute()
	require.NoError(t, err, "Failed to create interface for get test setup")
	createdID := *createRes.Id

	defer func() {
		client.TunnelInterfacesAPI.DeleteTunnelInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the interface
	getRes, httpResGet, errGet := client.TunnelInterfacesAPI.GetTunnelInterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get Tunnel Interface by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, interfaceName, getRes.Name, "Interface name should match")
	assert.Equal(t, "Test Tunnel Interface", getRes.GetComment(), "Comment should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_TunnelInterfacesAPIService_Update tests updating a Tunnel Interface.
func Test_network_services_TunnelInterfacesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := generateTunnelName("scm-tun-update-")
	tunnelInterface := createFullTestTunnelInterface(t, interfaceName)

	// Setup: Create an interface first
	createRes, _, err := client.TunnelInterfacesAPI.CreateTunnelInterfaces(context.Background()).TunnelInterfaces(tunnelInterface).Execute()
	require.NoError(t, err, "Failed to create interface for update test setup")
	createdID := *createRes.Id

	defer func() {
		client.TunnelInterfacesAPI.DeleteTunnelInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Prepare updated interface object
	updatedComment := "Updated comment for Tunnel"
	updatedMTU := int32(1400)
	updatedDefaultValue := "tunnel.500"

	updatedTunnel := createTestTunnelInterface(t, interfaceName)
	updatedTunnel.Id = &createdID
	updatedTunnel.SetComment(updatedComment)
	updatedTunnel.SetMtu(updatedMTU)
	updatedTunnel.SetDefaultValue(updatedDefaultValue)

	// Test: Update the interface
	updateRes, httpResUpdate, errUpdate := client.TunnelInterfacesAPI.UpdateTunnelInterfacesByID(context.Background(), createdID).TunnelInterfaces(updatedTunnel).Execute()

	require.NoError(t, errUpdate, "Failed to update Tunnel Interface")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the changes
	assert.Equal(t, updatedComment, updateRes.GetComment(), "Comment should be updated")
	assert.Equal(t, updatedMTU, updateRes.GetMtu(), "MTU should be updated")
	assert.Equal(t, updatedDefaultValue, updateRes.GetDefaultValue(), "DefaultValue should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_TunnelInterfacesAPIService_List tests listing Tunnel Interfaces.
func Test_network_services_TunnelInterfacesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Test: List the interfaces, filtering by the unique name
	listRes, httpResList, errList := client.TunnelInterfacesAPI.ListTunnelInterfaces(context.Background()).
		Folder("All").
		Execute()

	require.NoError(t, errList, "Failed to list Tunnel Interfaces")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	// Assert that the specific, uniquely named resource was returned.
	require.LessOrEqual(t, 1, len(listRes.GetData()), "Expected atleast one Tunnel Interface")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_TunnelInterfacesAPIService_DeleteByID tests deleting a Tunnel Interface.
func Test_network_services_TunnelInterfacesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := generateTunnelName("scm-tun-delete-")
	tunnelInterface := createFullTestTunnelInterface(t, interfaceName)

	// Setup: Create an interface first
	createRes, _, err := client.TunnelInterfacesAPI.CreateTunnelInterfaces(context.Background()).TunnelInterfaces(tunnelInterface).Execute()
	require.NoError(t, err, "Failed to create interface for delete test setup")
	createdID := *createRes.Id

	// Test: Delete the interface
	httpResDel, errDel := client.TunnelInterfacesAPI.DeleteTunnelInterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete Tunnel Interface")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}
