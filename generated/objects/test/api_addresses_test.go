/*
Objects Testing AddressesAPIService
*/
package objects

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)

// Test_objects_AddressesAPIService_Create tests the creation of an address object
// This test creates a new address and then deletes it to ensure proper cleanup
func Test_objects_AddressesAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a valid address object with unique name to avoid conflicts
	createdAddressName := "test-address-create-" + common.GenerateRandomString(10)
	address := objects.Addresses{
		Description: common.StringPtr("Test address for create API testing"),
		Folder:      common.StringPtr("Prisma Access"),           // Using Prisma Access folder scope
		Fqdn:        common.StringPtr("test.create.example.com"), // FQDN-based address
		Name:        createdAddressName,                          // Unique test name
	}

	fmt.Printf("Creating address with name: %s\n", address.Name)

	// Make the create request to the API
	req := client.AddressesAPI.CreateAddresses(context.Background()).Addresses(address)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create address")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdAddressName, res.Name, "Created address name should match")
	assert.Equal(t, common.StringPtr("Test address for create API testing"), res.Description, "Description should match")
	assert.True(t, *res.Folder == "Shared" || *res.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, common.StringPtr("test.create.example.com"), res.Fqdn, "FQDN should match")
	assert.NotEmpty(t, res.Id, "Created address should have an ID")

	// Use the ID from the response object
	createdAddressID := res.Id
	t.Logf("Successfully created address: %s with ID: %s", address.Name, createdAddressID)

	// Cleanup: Delete the created address to maintain test isolation
	reqDel := client.AddressesAPI.DeleteAddressesByID(context.Background(), createdAddressID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete address during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up address: %s", createdAddressID)
}

// Test_objects_AddressesAPIService_GetByID tests retrieving an address by its ID
// This test creates an address, retrieves it by ID, then deletes it
func Test_objects_AddressesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create an address first to have something to retrieve
	createdAddressName := "test-address-getbyid-" + common.GenerateRandomString(10)
	address := objects.Addresses{
		Description: common.StringPtr("Test address for get by ID API testing"),
		Folder:      common.StringPtr("Prisma Access"),            // Using Prisma Access folder scope
		Fqdn:        common.StringPtr("test.getbyid.example.com"), // FQDN-based address
		Name:        createdAddressName,                           // Unique test name
	}

	// Create the address via API
	req := client.AddressesAPI.CreateAddresses(context.Background()).Addresses(address)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create address for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAddressID := createRes.Id
	require.NotEmpty(t, createdAddressID, "Created address should have an ID")

	// Test Get by ID operation
	reqGetById := client.AddressesAPI.GetAddressesByID(context.Background(), createdAddressID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get address by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdAddressName, getRes.Name, "Address name should match")
	assert.Equal(t, common.StringPtr("Test address for get by ID API testing"), getRes.Description, "Description should match")
	assert.True(t, *getRes.Folder == "Shared" || *getRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, common.StringPtr("test.getbyid.example.com"), getRes.Fqdn, "FQDN should match")
	assert.Equal(t, createdAddressID, getRes.Id, "Address ID should match")

	t.Logf("Successfully retrieved address: %s", getRes.Name)

	// Cleanup: Delete the created address
	reqDel := client.AddressesAPI.DeleteAddressesByID(context.Background(), createdAddressID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete address during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up address: %s", createdAddressID)
}

// Test_objects_AddressesAPIService_Update tests updating an existing address
// This test creates an address, updates it, then deletes it
func Test_objects_AddressesAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create an address first to have something to update
	createdAddressName := "test-address-update-" + common.GenerateRandomString(10)
	address := objects.Addresses{
		Description: common.StringPtr("Test address for update API testing"),
		Folder:      common.StringPtr("Prisma Access"),           // Using Prisma Access folder scope
		Fqdn:        common.StringPtr("test.update.example.com"), // Initial FQDN
		Name:        createdAddressName,                          // Unique test name
	}

	// Create the address via API
	req := client.AddressesAPI.CreateAddresses(context.Background()).Addresses(address)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create address for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAddressID := createRes.Id
	require.NotEmpty(t, createdAddressID, "Created address should have an ID")

	// Test Update operation with modified fields
	updatedAddress := objects.Addresses{
		Description: common.StringPtr("Updated test address description"), // Updated description
		Folder:      common.StringPtr("Prisma Access"),                    // Keep same folder scope
		Fqdn:        common.StringPtr("updated.test.example.com"),         // Updated FQDN
		Name:        createdAddressName,                                   // Keep same name (required for update)
	}

	reqUpdate := client.AddressesAPI.UpdateAddressesByID(context.Background(), createdAddressID).Addresses(updatedAddress)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update address")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdAddressName, updateRes.Name, "Address name should remain the same")
	assert.Equal(t, common.StringPtr("Updated test address description"), updateRes.Description, "Description should be updated")
	assert.True(t, *updateRes.Folder == "Shared" || *updateRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, common.StringPtr("updated.test.example.com"), updateRes.Fqdn, "FQDN should be updated")
	assert.Equal(t, createdAddressID, updateRes.Id, "Address ID should remain the same")

	t.Logf("Successfully updated address: %s", createdAddressName)

	// Cleanup: Delete the created address
	reqDel := client.AddressesAPI.DeleteAddressesByID(context.Background(), createdAddressID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete address during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up address: %s", createdAddressID)
}

// Test_objects_AddressesAPIService_List tests listing addresses with folder filter
// This test creates an address, lists addresses to verify it's included, then deletes it
func Test_objects_AddressesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Test List operation with folder filter
	reqList := client.AddressesAPI.ListAddresses(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list addresses")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one address in the list")
}

// Test_objects_AddressesAPIService_DeleteByID tests deleting an address by its ID
// This test creates an address, deletes it, then verifies the deletion was successful
func Test_objects_AddressesAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create an address first to have something to delete
	createdAddressName := "test-address-delete-" + common.GenerateRandomString(10)
	address := objects.Addresses{
		Description: common.StringPtr("Test address for delete API testing"),
		Folder:      common.StringPtr("Prisma Access"),           // Using Prisma Access folder scope
		Fqdn:        common.StringPtr("test.delete.example.com"), // FQDN-based address
		Name:        createdAddressName,                          // Unique test name
	}

	// Create the address via API
	req := client.AddressesAPI.CreateAddresses(context.Background()).Addresses(address)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create address for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdAddressID := createRes.Id
	require.NotEmpty(t, createdAddressID, "Created address should have an ID")

	// Test Delete by ID operation
	reqDel := client.AddressesAPI.DeleteAddressesByID(context.Background(), createdAddressID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete address")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted address: %s", createdAddressID)
}
