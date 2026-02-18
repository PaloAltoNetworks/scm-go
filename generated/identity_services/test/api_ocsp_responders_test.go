/*
Identity Services Testing

OCSPRespondersAPIService
*/

package identity_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// Test_identity_services_OCSPRespondersAPIService_Create tests the creation of an OCSP responder.
func Test_identity_services_OCSPRespondersAPIService_Create(t *testing.T) {
	t.Skip("Create returns no model and List fails with deserialization error - cannot retrieve created object ID")
	client := SetupIdentitySvcTestClient(t)
	createdName := "test-ocsp-create-" + common.GenerateRandomString(6)

	// define the OCSP responder
	ocspResponder := identity_services.OcspResponders{
		Folder:   common.StringPtr("All"),
		Name:     createdName,
		HostName: "ocsp.example.com",
	}

	fmt.Printf("Creating OCSP Responder with name: %s\n", ocspResponder.Name)
	req := client.OCSPRespondersAPI.CreateOCSPResponders(context.Background()).OcspResponders(ocspResponder)
	httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create OCSP Responder")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Get the created object to retrieve the ID
	listRes, _, errList := client.OCSPRespondersAPI.ListOCSPResponders(context.Background()).Folder("All").Name(createdName).Execute()
	require.NoError(t, errList, "Failed to list OCSP Responders after creation")
	require.NotNil(t, listRes, "List response should not be nil")
	require.Greater(t, len(listRes.Data), 0, "Should have at least one OCSP responder")

	createdID := listRes.Data[0].Id

	defer func() {
		t.Logf("Cleaning up OCSP Responder with ID: %s", createdID)
		_, errDel := client.OCSPRespondersAPI.DeleteOCSPRespondersByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete OCSP Responder during cleanup")
	}()

	t.Logf("Successfully created OCSP Responder: %s with ID: %s", ocspResponder.Name, createdID)
	assert.Equal(t, createdName, listRes.Data[0].Name, "Created OCSP responder name should match")
	assert.Equal(t, "All", *listRes.Data[0].Folder, "Folder should match")
	assert.Equal(t, "ocsp.example.com", listRes.Data[0].HostName, "Hostname should match")
}

// Test_identity_services_OCSPRespondersAPIService_GetByID tests retrieving an OCSP responder by ID.
func Test_identity_services_OCSPRespondersAPIService_GetByID(t *testing.T) {
	t.Skip("Create returns no model and List fails with deserialization error - cannot retrieve created object ID")
	client := SetupIdentitySvcTestClient(t)
	ocspName := "test-ocsp-get-" + common.GenerateRandomString(6)

	ocspResponder := identity_services.OcspResponders{
		Folder:   common.StringPtr("All"),
		Name:     ocspName,
		HostName: "ocsp-get.example.com",
	}

	_, err := client.OCSPRespondersAPI.CreateOCSPResponders(context.Background()).OcspResponders(ocspResponder).Execute()
	require.NoError(t, err, "Failed to create OCSP Responder for get test")

	// Get the ID from list
	listRes, _, errList := client.OCSPRespondersAPI.ListOCSPResponders(context.Background()).Folder("All").Name(ocspName).Execute()
	require.NoError(t, errList, "Failed to list OCSP Responders")
	require.Greater(t, len(listRes.Data), 0, "Should have at least one OCSP responder")
	createdID := listRes.Data[0].Id

	defer func() {
		client.OCSPRespondersAPI.DeleteOCSPRespondersByID(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.OCSPRespondersAPI.GetOCSPRespondersByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get OCSP Responder by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, ocspName, getRes.Name, "OCSP responder name should match")
}

// Test_identity_services_OCSPRespondersAPIService_Update tests updating an existing OCSP responder.
func Test_identity_services_OCSPRespondersAPIService_Update(t *testing.T) {
	t.Skip("Create returns no model and List fails with deserialization error - cannot retrieve created object ID")
	client := SetupIdentitySvcTestClient(t)
	ocspName := "test-ocsp-update-" + common.GenerateRandomString(6)

	ocspResponder := identity_services.OcspResponders{
		Folder:   common.StringPtr("All"),
		Name:     ocspName,
		HostName: "ocsp-original.example.com",
	}

	_, err := client.OCSPRespondersAPI.CreateOCSPResponders(context.Background()).OcspResponders(ocspResponder).Execute()
	require.NoError(t, err, "Failed to create OCSP Responder for update test")

	// Get the ID from list
	listRes, _, errList := client.OCSPRespondersAPI.ListOCSPResponders(context.Background()).Folder("All").Name(ocspName).Execute()
	require.NoError(t, errList, "Failed to list OCSP Responders")
	require.Greater(t, len(listRes.Data), 0, "Should have at least one OCSP responder")
	createdID := listRes.Data[0].Id

	defer func() {
		client.OCSPRespondersAPI.DeleteOCSPRespondersByID(context.Background(), createdID).Execute()
	}()

	// update the hostname
	updatedResponder := identity_services.OcspResponders{
		Name:     ocspName,
		HostName: "ocsp-updated.example.com",
	}

	updateRes, httpResUpdate, errUpdate := client.OCSPRespondersAPI.UpdateOCSPRespondersByID(context.Background(), createdID).OcspResponders(updatedResponder).Execute()
	require.NoError(t, errUpdate, "Failed to update OCSP Responder")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, "ocsp-updated.example.com", updateRes.HostName, "Hostname should be updated")
}

// Test_identity_services_OCSPRespondersAPIService_List tests listing OCSP responders.
func Test_identity_services_OCSPRespondersAPIService_List(t *testing.T) {
	t.Skip("List response fails with deserialization error - no value given for required property data")
	client := SetupIdentitySvcTestClient(t)
	ocspName := "test-ocsp-list-" + common.GenerateRandomString(6)

	ocspResponder := identity_services.OcspResponders{
		Folder:   common.StringPtr("Shared"),
		Name:     ocspName,
		HostName: "ocsp-list.example.com",
	}

	_, err := client.OCSPRespondersAPI.CreateOCSPResponders(context.Background()).OcspResponders(ocspResponder).Execute()
	require.NoError(t, err, "Failed to create OCSP Responder for list test")

	// Get the ID from list for cleanup
	listRes, _, errList := client.OCSPRespondersAPI.ListOCSPResponders(context.Background()).Folder("Shared").Name(ocspName).Execute()
	require.NoError(t, errList, "Failed to list OCSP Responders")
	require.Greater(t, len(listRes.Data), 0, "Should have at least one OCSP responder")
	createdID := listRes.Data[0].Id

	defer func() {
		client.OCSPRespondersAPI.DeleteOCSPRespondersByID(context.Background(), createdID).Execute()
	}()

	listRes2, httpResList, errList2 := client.OCSPRespondersAPI.ListOCSPResponders(context.Background()).Folder("Shared").Limit(200).Execute()
	require.NoError(t, errList2, "Failed to list OCSP Responders")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes2, "List response should not be nil")

	found := false
	for _, p := range listRes2.Data {
		if p.Name == ocspName {
			found = true
			break
		}
	}
	assert.True(t, found, "Created OCSP Responder should be present in the list")
}

// Test_identity_services_OCSPRespondersAPIService_DeleteByID tests deleting an OCSP responder.
func Test_identity_services_OCSPRespondersAPIService_DeleteByID(t *testing.T) {
	t.Skip("Create returns no model and List fails with deserialization error - cannot retrieve created object ID")
	client := SetupIdentitySvcTestClient(t)
	ocspName := "test-ocsp-delete-" + common.GenerateRandomString(6)

	ocspResponder := identity_services.OcspResponders{
		Folder:   common.StringPtr("Shared"),
		Name:     ocspName,
		HostName: "ocsp-delete.example.com",
	}

	_, err := client.OCSPRespondersAPI.CreateOCSPResponders(context.Background()).OcspResponders(ocspResponder).Execute()
	require.NoError(t, err, "Failed to create OCSP Responder for delete test")

	// Get the ID from list
	listRes, _, errList := client.OCSPRespondersAPI.ListOCSPResponders(context.Background()).Folder("Shared").Name(ocspName).Execute()
	require.NoError(t, errList, "Failed to list OCSP Responders")
	require.Greater(t, len(listRes.Data), 0, "Should have at least one OCSP responder")
	createdID := listRes.Data[0].Id

	httpResDel, errDel := client.OCSPRespondersAPI.DeleteOCSPRespondersByID(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete OCSP Responder")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_identity_services_OCSPRespondersAPIService_FetchOCSPResponders tests the FetchOCSPResponders convenience method
func Test_identity_services_OCSPRespondersAPIService_FetchOCSPResponders(t *testing.T) {
	t.Skip("Create returns no model and List fails with deserialization error - cannot retrieve created object ID")
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Create a test object first
	testName := "fetch-ocsp-" + common.GenerateRandomString(6)

	testObj := identity_services.OcspResponders{
		Name:     testName,
		Folder:   common.StringPtr("Prisma Access"),
		HostName: "ocsp-fetch.example.com",
	}

	createReq := client.OCSPRespondersAPI.CreateOCSPResponders(context.Background()).OcspResponders(testObj)
	_, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")

	// Get the ID from list
	listRes, _, errList := client.OCSPRespondersAPI.ListOCSPResponders(context.Background()).Folder("Prisma Access").Name(testName).Execute()
	require.NoError(t, errList, "Failed to list OCSP Responders")
	require.Greater(t, len(listRes.Data), 0, "Should have at least one OCSP responder")
	createdID := listRes.Data[0].Id

	// Cleanup after test
	defer func() {
		deleteReq := client.OCSPRespondersAPI.DeleteOCSPRespondersByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.OCSPRespondersAPI.FetchOCSPResponders(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch ocsp_responders by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchOCSPResponders found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.OCSPRespondersAPI.FetchOCSPResponders(
		context.Background(),
		"non-existent-ocsp-responders-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchOCSPResponders correctly returned nil for non-existent object")
}
