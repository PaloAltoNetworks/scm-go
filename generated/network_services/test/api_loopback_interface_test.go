package network_services

/*
 * Network Services Testing
 *
 * LoopbackInterfacesAPIService
 */

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

// --- Helper Functions ---

// generateLoopbackName creates a valid name starting with $ for loopback interfaces.
func generateLoopbackName(base string) string {
	// Name must match: ^\$[a-zA-Z\d\-_.]+
	return "$" + base + common.GenerateRandomString(4)
}

// createTestLoopbackInterface creates a minimal LoopbackInterfaces object for testing.
func createTestLoopbackInterface(t *testing.T, interfaceName string) network_services.LoopbackInterfaces {
	// 'name' is the only required field.
	return *network_services.NewLoopbackInterfaces(interfaceName)
}

// createFullTestLoopbackInterface creates a more complete LoopbackInterfaces object for update/get testing.
func createFullTestLoopbackInterface(t *testing.T, interfaceName string) network_services.LoopbackInterfaces {
	mtu := int32(1500)
	comment := "Test Loopback Interface"

	loopback := createTestLoopbackInterface(t, interfaceName)
	loopback.SetMtu(mtu)
	loopback.SetComment(comment)
	loopback.SetFolder("All")

	// Note: Ip, Ipv6 are complex types and are omitted for simplicity,
	// but should be included in full-scale integration tests.
	return loopback
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_LoopbackInterfacesAPIService_Create tests the creation of a Loopback Interface.
func Test_network_services_LoopbackInterfacesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := generateLoopbackName("scm-lb-create-")

	loopbackInterface := createFullTestLoopbackInterface(t, interfaceName)

	t.Logf("Creating Loopback Interface with name: %s", interfaceName)
	req := client.LoopbackInterfacesAPI.CreateLoopbackInterfaces(context.Background()).LoopbackInterfaces(loopbackInterface)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Loopback Interface")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created interface should have an ID")
	createdID := *res.Id

	// Cleanup the created interface
	defer func() {
		t.Logf("Cleaning up Loopback Interface with ID: %s", createdID)
		_, errDel := client.LoopbackInterfacesAPI.DeleteLoopbackInterfacesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete Loopback Interface during cleanup")
	}()

	t.Logf("Successfully created Loopback Interface: %s with ID: %s", interfaceName, createdID)

	// Verify the response matches key input fields
	assert.Equal(t, interfaceName, res.Name, "Created interface name should match")
	assert.Equal(t, int32(1500), res.GetMtu(), "MTU should match the set value")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_LoopbackInterfacesAPIService_GetByID tests retrieving a Loopback Interface by ID.
func Test_network_services_LoopbackInterfacesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := generateLoopbackName("scm-lb-get-")
	loopbackInterface := createFullTestLoopbackInterface(t, interfaceName)

	// Setup: Create an interface first
	createRes, _, err := client.LoopbackInterfacesAPI.CreateLoopbackInterfaces(context.Background()).LoopbackInterfaces(loopbackInterface).Execute()
	require.NoError(t, err, "Failed to create interface for get test setup")
	createdID := *createRes.Id

	defer func() {
		client.LoopbackInterfacesAPI.DeleteLoopbackInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the interface
	getRes, httpResGet, errGet := client.LoopbackInterfacesAPI.GetLoopbackInterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get Loopback Interface by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, interfaceName, getRes.Name, "Interface name should match")
	assert.Equal(t, "Test Loopback Interface", getRes.GetComment(), "Comment should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_LoopbackInterfacesAPIService_Update tests updating a Loopback Interface.
func Test_network_services_LoopbackInterfacesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := generateLoopbackName("scm-lb-update-")
	loopbackInterface := createFullTestLoopbackInterface(t, interfaceName)

	// Setup: Create an interface first
	createRes, _, err := client.LoopbackInterfacesAPI.CreateLoopbackInterfaces(context.Background()).LoopbackInterfaces(loopbackInterface).Execute()
	require.NoError(t, err, "Failed to create interface for update test setup")
	createdID := *createRes.Id

	defer func() {
		client.LoopbackInterfacesAPI.DeleteLoopbackInterfacesByID(context.Background(), createdID).Execute()
	}()

	// Prepare updated interface object
	updatedComment := "Updated comment for Loopback"
	updatedMTU := int32(1450)
	updatedDefaultValue := "loopback.2000"

	updatedLoopback := createTestLoopbackInterface(t, interfaceName)
	updatedLoopback.Id = &createdID
	updatedLoopback.SetComment(updatedComment)
	updatedLoopback.SetMtu(updatedMTU)
	updatedLoopback.SetDefaultValue(updatedDefaultValue)

	// Test: Update the interface
	updateRes, httpResUpdate, errUpdate := client.LoopbackInterfacesAPI.UpdateLoopbackInterfacesByID(context.Background(), createdID).LoopbackInterfaces(updatedLoopback).Execute()

	require.NoError(t, errUpdate, "Failed to update Loopback Interface")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the changes
	assert.Equal(t, updatedComment, updateRes.GetComment(), "Comment should be updated")
	assert.Equal(t, updatedMTU, updateRes.GetMtu(), "MTU should be updated")
	assert.Equal(t, updatedDefaultValue, updateRes.GetDefaultValue(), "DefaultValue should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_LoopbackInterfacesAPIService_List tests listing Loopback Interfaces.
func Test_network_services_LoopbackInterfacesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Test: List the interfaces, filtering by the unique name
	listRes, httpResList, errList := client.LoopbackInterfacesAPI.ListLoopbackInterfaces(context.Background()).
		Folder("All").
		Execute()

	require.NoError(t, errList, "Failed to list Loopback Interfaces")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_LoopbackInterfacesAPIService_DeleteByID tests deleting a Loopback Interface.
func Test_network_services_LoopbackInterfacesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	interfaceName := generateLoopbackName("scm-lb-delete-")
	loopbackInterface := createFullTestLoopbackInterface(t, interfaceName)

	// Setup: Create an interface first
	createRes, _, err := client.LoopbackInterfacesAPI.CreateLoopbackInterfaces(context.Background()).LoopbackInterfaces(loopbackInterface).Execute()
	require.NoError(t, err, "Failed to create interface for delete test setup")
	createdID := *createRes.Id

	// Test: Delete the interface
	httpResDel, errDel := client.LoopbackInterfacesAPI.DeleteLoopbackInterfacesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete Loopback Interface")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}

// Test_network_services_LoopbackInterfacesAPIService_Fetch tests the fetch convenience method.
func Test_network_services_LoopbackInterfacesAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a Loopback Interface to fetch
	ifName := generateLoopbackName("loopback-fetch")
	loopIf := createFullTestLoopbackInterface(t, ifName)

	reqCreate := client.LoopbackInterfacesAPI.CreateLoopbackInterfaces(context.Background()).
		LoopbackInterfaces(loopIf)

	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create Loopback Interface for fetch test")
	createdID := *createRes.Id
	createdFolder := createRes.Folder
	require.NotEmpty(t, createdID, "Created loopback interface ID should not be empty")

	// Defer cleanup
	defer func() {
		t.Logf("Cleaning up Loopback Interface with ID: %s", createdID)
		_, errDel := client.LoopbackInterfacesAPI.DeleteLoopbackInterfacesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete loopback interface during cleanup")
	}()

	// Test Fetch by name operation
	fmt.Printf("Attempting to fetch Loopback Interface with name: %s\n", ifName)
	fetchedIf, errFetch := client.LoopbackInterfacesAPI.FetchLoopbackInterfaces(context.Background(), ifName, createdFolder, nil, nil)

	// Verify the fetch operation was successful
	require.NoError(t, errFetch, "Failed to fetch loopback interface by name")
	require.NotNil(t, fetchedIf, "Fetched loopback interface should not be nil")
	assert.Equal(t, ifName, fetchedIf.GetName(), "Loopback interface name should match")
	assert.Equal(t, createdID, *fetchedIf.Id, "Loopback interface ID should match")
	assert.Equal(t, *createdFolder, fetchedIf.GetFolder(), "Folder should match")
	t.Logf("Successfully fetched Loopback Interface: %s", ifName)

	// Test fetching non-existent loopback interface (should return nil)
	nonExistentName := "loopback.99999"
	notFoundIf, errNotFound := client.LoopbackInterfacesAPI.FetchLoopbackInterfaces(context.Background(), nonExistentName, createdFolder, nil, nil)
	require.NoError(t, errNotFound, "Fetch for non-existent loopback interface should not error")
	assert.Nil(t, notFoundIf, "Non-existent loopback interface should return nil")
	t.Logf("Successfully verified fetch returns nil for non-existent loopback interface")
}
