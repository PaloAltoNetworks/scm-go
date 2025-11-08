/*
 * Network Services Testing
 *
 * VlanInterfacesAPIService
 */

package network_services

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Replace with your actual imports for common utils and generated client
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// NOTE: Placeholder functions and types are assumed to be available in the test environment.
// Replace these placeholders with your actual setup and utilities.
// - SetupNetworkSvcTestClient(t) -> returns a *network_services.APIClient
// - handleAPIError(err)
// - common.StringPtr(s) -> returns *string

// --- Helper Functions ---

// createTestVlanInterface creates a minimal VlanInterfaces object for testing.
func createTestVlanInterface(t *testing.T, interfaceName string) network_services.VlanInterfaces {
	// 'name' is the only required field.
	return *network_services.NewVlanInterfaces(interfaceName)
}

// createFullTestVlanInterface creates a more complete VlanInterfaces object for update/get testing.
func createFullTestVlanInterface(t *testing.T, interfaceName string) network_services.VlanInterfaces {
	mtu := int32(1500)
	comment := "Test VLAN Interface"

	vlan := createTestVlanInterface(t, interfaceName)
	vlan.SetMtu(mtu)
	vlan.SetComment(comment)
	vlan.SetFolder("All")
	// Note: Ip, Arp, DdnsConfig, DhcpClient are complex types and are omitted for simplicity,
	// but should be included in full-scale integration tests.
	return vlan
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_VLANInterfacesAPIService_Create tests the creation of a VLAN Interface.
func Test_network_services_VLANInterfacesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := "$scm_vlan_if-" + common.GenerateRandomString(6)

	vlanInterface := createFullTestVlanInterface(t, interfaceName)

	t.Logf("Creating VLAN Interface with name: %s", interfaceName)
	req := client.VLANInterfacesAPI.CreateVLANInterfaces(context.Background()).VlanInterfaces(vlanInterface)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create VLAN Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created interface should have an ID")
	createdID := *res.Id

	// Cleanup the created interface
	defer func() {
		t.Logf("Cleaning up VLAN Interface with ID: %s", createdID)
		_, errDel := client.VLANInterfacesAPI.DeleteVLANInterfacesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete VLAN Interface during cleanup")
	}()

	t.Logf("Successfully created VLAN Interface: %s with ID: %s", interfaceName, createdID)

	// Verify the response matches key input fields
	assert.Equal(t, interfaceName, res.Name, "Created interface name should match")
	assert.Equal(t, "1500", fmt.Sprintf("%d", res.GetMtu()), "MTU should match the set value")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_VLANInterfacesAPIService_GetByID tests retrieving a VLAN Interface by ID.
func Test_network_services_VLANInterfacesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := "$scm_vlan_if-" + common.GenerateRandomString(4)
	vlanInterface := createFullTestVlanInterface(t, interfaceName)

	// Setup: Create an interface first
	createRes, _, err := client.VLANInterfacesAPI.CreateVLANInterfaces(context.Background()).VlanInterfaces(vlanInterface).Execute()
	require.NoError(t, err, "Failed to create interface for get test setup")
	createdID := *createRes.Id

	defer func() {
		client.VLANInterfacesAPI.DeleteVLANInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the interface
	getRes, httpResGet, errGet := client.VLANInterfacesAPI.GetVLANInterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get VLAN Interface by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, interfaceName, getRes.Name, "Interface name should match")
	assert.Equal(t, "Test VLAN Interface", getRes.GetComment(), "Comment should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_VLANInterfacesAPIService_Update tests updating a VLAN Interface.
func Test_network_services_VLANInterfacesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := "$scm_vlan_if-" + common.GenerateRandomString(4)
	vlanInterface := createFullTestVlanInterface(t, interfaceName)

	// Setup: Create an interface first
	createRes, _, err := client.VLANInterfacesAPI.CreateVLANInterfaces(context.Background()).VlanInterfaces(vlanInterface).Execute()
	require.NoError(t, err, "Failed to create interface for update test setup")
	createdID := *createRes.Id

	defer func() {
		client.VLANInterfacesAPI.DeleteVLANInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Prepare updated interface object
	updatedComment := "Updated comment for VLAN 30"
	updatedMTU := int32(1400)
	updatedVlan := createTestVlanInterface(t, interfaceName)
	updatedVlan.Id = &createdID
	updatedVlan.SetComment(updatedComment)
	updatedVlan.SetMtu(updatedMTU)
	updatedVlan.SetVlanTag("300") // Also update VLAN tag

	// Test: Update the interface
	updateRes, httpResUpdate, errUpdate := client.VLANInterfacesAPI.UpdateVLANlInterfacesByID(context.Background(), createdID).VlanInterfaces(updatedVlan).Execute()

	require.NoError(t, errUpdate, "Failed to update VLAN Interface")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the changes
	assert.Equal(t, updatedComment, updateRes.GetComment(), "Comment should be updated")
	assert.Equal(t, updatedMTU, updateRes.GetMtu(), "MTU should be updated")
	assert.Equal(t, "300", updateRes.GetVlanTag(), "VLAN Tag should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_VLANInterfacesAPIService_List tests listing VLAN Interfaces.
func Test_network_services_VLANInterfacesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := "$scm_vlan_if-" + common.GenerateRandomString(6)

	vlanInterface := createFullTestVlanInterface(t, interfaceName)

	t.Logf("Creating VLAN Interface with name: %s", interfaceName)
	req := client.VLANInterfacesAPI.CreateVLANInterfaces(context.Background()).VlanInterfaces(vlanInterface)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create interface for list test setup")
	createdID := *createRes.Id

	defer func() {
		client.VLANInterfacesAPI.DeleteVLANInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: List the interfaces, using some filters
	listRes, httpResList, errList := client.VLANInterfacesAPI.ListVLANInterfaces(context.Background()).
		Folder("All").
		Limit(10).
		Offset(0).
		Execute()

	require.NoError(t, errList, "Failed to list VLAN Interfaces")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	// Since we filtered by a unique name, we expect to find at least 1 record.
	assert.Condition(t, func() bool { return len(listRes.GetData()) >= 1 }, "Expected to find at least one VLAN Interface in the list")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_VLANInterfacesAPIService_DeleteByID tests deleting a VLAN Interface.
func Test_network_services_VLANInterfacesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := "$scm_vlan_if-" + common.GenerateRandomString(6)

	vlanInterface := createFullTestVlanInterface(t, interfaceName)

	t.Logf("Creating VLAN Interface with name: %s", interfaceName)
	req := client.VLANInterfacesAPI.CreateVLANInterfaces(context.Background()).VlanInterfaces(vlanInterface)
	createRes, _, err := req.Execute()

	require.NoError(t, err, "Failed to create interface for delete test setup")
	createdID := *createRes.Id

	// Test: Delete the interface
	httpResDel, errDel := client.VLANInterfacesAPI.DeleteVLANInterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete VLAN Interface")
	// The API returns 200 OK for successful deletion.
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}
