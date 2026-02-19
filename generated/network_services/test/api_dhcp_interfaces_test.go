/*
 * Network Services Testing
 *
 * DHCPInterfacesAPIService
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

// Test_networkservices_DHCPInterfacesAPIService_Create tests the creation of a DHCP Interface.
func Test_networkservices_DHCPInterfacesAPIService_Create(t *testing.T) {
	t.Skip("Requires pre-existing device ethernet interface - cannot create interfaces via API")
	// Setup the authenticated client.
	client := SetupNetworkSvcTestClient(t)

	// Create a valid DHCP Interface object with a unique name.
	interfaceName := "ethernet1/1"

	dhcpInterface := network_services.DhcpInterfaces{
		Name:   interfaceName,
		Folder: common.StringPtr("All"),
	}

	fmt.Printf("Attempting to create DHCP Interface with name: %s\n", dhcpInterface.Name)

	// Make the create request to the API.
	req := client.DHCPInterfacesAPI.CreateDHCPInterfaces(context.Background()).DhcpInterfaces(dhcpInterface)
	res, httpRes, err := req.Execute()

	// Defer cleanup for the DHCP Interface.
	if res != nil && res.Id != nil {
		defer func() {
			t.Logf("Cleaning up DHCP Interface with ID: %s", *res.Id)
			delReq := client.DHCPInterfacesAPI.DeleteDHCPInterfacesByID(context.Background(), *res.Id)
			_, errDel := delReq.Execute()
			if errDel != nil {
				t.Logf("Failed to delete DHCP Interface during cleanup: %v", errDel)
			}
		}()
	}

	// Verify the request was successful.
	handleAPIError(err)
	require.NoError(t, err, "Create request should not return an error")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "The response from create should not be nil")
	assert.Equal(t, interfaceName, res.Name, "The name of the created interface should match")
	assert.NotEmpty(t, *res.Id, "The ID of the created interface should not be empty")

	t.Logf("Successfully created DHCP Interface with ID: %s", *res.Id)
}

// Test_networkservices_DHCPInterfacesAPIService_GetByID tests the retrieval of a DHCP Interface by its ID.
func Test_networkservices_DHCPInterfacesAPIService_GetByID(t *testing.T) {
	t.Skip("Requires pre-existing device ethernet interface - cannot create interfaces via API")
	client := SetupNetworkSvcTestClient(t)

	// Create an interface to retrieve.
	interfaceName := "ethernet1/2"

	dhcpInterface := network_services.DhcpInterfaces{
		Name:   interfaceName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.DHCPInterfacesAPI.CreateDHCPInterfaces(context.Background()).DhcpInterfaces(dhcpInterface).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create interface for get test")
	createdInterfaceID := *createRes.Id

	// Defer cleanup for the DHCP Interface.
	defer func() {
		t.Logf("Cleaning up DHCP Interface with ID: %s", createdInterfaceID)
		_, errDel := client.DHCPInterfacesAPI.DeleteDHCPInterfacesByID(context.Background(), createdInterfaceID).Execute()
		if errDel != nil {
			t.Logf("Failed to delete interface during cleanup: %v", errDel)
		}
	}()

	t.Logf("Created DHCP Interface for Get test with ID: %s", createdInterfaceID)

	// Test the Get by ID operation.
	fmt.Printf("Attempting to get DHCP Interface with ID: %s\n", createdInterfaceID)
	req := client.DHCPInterfacesAPI.GetDHCPInterfacesByID(context.Background(), createdInterfaceID)
	getRes, httpRes, err := req.Execute()

	// Verify the retrieval was successful.
	handleAPIError(err)
	require.NoError(t, err, "Get by ID request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "The response from get should not be nil")
	assert.Equal(t, createdInterfaceID, *getRes.Id, "The ID of the retrieved interface should match")
	assert.Equal(t, interfaceName, getRes.Name, "The name of the retrieved interface should match")
}

// Test_networkservices_DHCPInterfacesAPIService_Update tests updating a DHCP Interface.
func Test_networkservices_DHCPInterfacesAPIService_Update(t *testing.T) {
	t.Skip("Requires pre-existing device ethernet interface - cannot create interfaces via API")
	client := SetupNetworkSvcTestClient(t)

	// Create an interface to update.
	interfaceName := "ethernet1/3"

	dhcpInterface := network_services.DhcpInterfaces{
		Name:   interfaceName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.DHCPInterfacesAPI.CreateDHCPInterfaces(context.Background()).DhcpInterfaces(dhcpInterface).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create interface for update test")
	createdInterfaceID := *createRes.Id

	// Defer cleanup for the DHCP Interface.
	defer func() {
		t.Logf("Cleaning up DHCP Interface with ID: %s", createdInterfaceID)
		_, errDel := client.DHCPInterfacesAPI.DeleteDHCPInterfacesByID(context.Background(), createdInterfaceID).Execute()
		if errDel != nil {
			t.Logf("Failed to delete interface during cleanup: %v", errDel)
		}
	}()

	t.Logf("Created DHCP Interface for Update test with ID: %s", createdInterfaceID)

	// Update the interface object.
	updatedInterface := network_services.DhcpInterfaces{
		Name:   interfaceName,
		Folder: common.StringPtr("All"),
		Relay: &network_services.DhcpInterfacesRelay{
			Ip: network_services.DhcpInterfacesRelayIp{
				Enabled: true,
				Server:  []string{"10.0.0.1"},
			},
		},
	}

	// Test the Update by ID operation.
	fmt.Printf("Attempting to update DHCP Interface with ID: %s\n", createdInterfaceID)
	reqUpdate := client.DHCPInterfacesAPI.UpdateDHCPInterfacesByID(context.Background(), createdInterfaceID).DhcpInterfaces(updatedInterface)
	updateRes, httpRes, err := reqUpdate.Execute()

	// Verify the update was successful.
	handleAPIError(err)
	require.NoError(t, err, "Update request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "The response from update should not be nil")
}

// Test_networkservices_DHCPInterfacesAPIService_List tests listing DHCP Interfaces.
func Test_networkservices_DHCPInterfacesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Read-only test: list existing objects (no Create needed)
	listRes, httpResList, errList := client.DHCPInterfacesAPI.ListDHCPInterfaces(context.Background()).Folder("All").Limit(200).Offset(0).Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list DHCP interfaces")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed DHCP interfaces")
}

// Test_networkservices_DHCPInterfacesAPIService_Fetch tests the fetch convenience method.
func Test_networkservices_DHCPInterfacesAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Read-only test: Fetch non-existent object (should return nil, nil)
	notFound, err := client.DHCPInterfacesAPI.FetchDHCPInterfaces(
		context.Background(),
		"non-existent-dhcp-interface-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchDHCPInterfaces correctly returned nil for non-existent object")
}

// Test_networkservices_DHCPInterfacesAPIService_DeleteByID tests deleting a DHCP Interface.
func Test_networkservices_DHCPInterfacesAPIService_DeleteByID(t *testing.T) {
	t.Skip("Requires pre-existing device ethernet interface - cannot create interfaces via API")
	client := SetupNetworkSvcTestClient(t)

	// Create an interface to delete.
	interfaceName := "ethernet1/6"

	dhcpInterface := network_services.DhcpInterfaces{
		Name:   interfaceName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.DHCPInterfacesAPI.CreateDHCPInterfaces(context.Background()).DhcpInterfaces(dhcpInterface).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create interface for delete test")
	createdInterfaceID := *createRes.Id

	t.Logf("Created DHCP Interface for Delete test with ID: %s", createdInterfaceID)

	// Test the Delete by ID operation.
	fmt.Printf("Attempting to delete DHCP Interface with ID: %s\n", createdInterfaceID)
	reqDel := client.DHCPInterfacesAPI.DeleteDHCPInterfacesByID(context.Background(), createdInterfaceID)
	httpResDel, errDel := reqDel.Execute()

	// Verify the delete operation was successful.
	handleAPIError(errDel)
	require.NoError(t, errDel, "Failed to delete interface")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
	t.Logf("Successfully deleted DHCP Interface: %s", createdInterfaceID)
}
