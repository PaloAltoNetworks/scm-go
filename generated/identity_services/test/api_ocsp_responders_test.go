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
	client := SetupIdentitySvcTestClient(t)
	createdName := "test-ocsp-create-" + common.GenerateRandomString(6)

	// define the OCSP responder
	ocspResponder := identity_services.OcspResponders{
		Folder:   common.StringPtr("Prisma Access"),
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

	// Use Fetch to retrieve the created object (Create returns no model, List has deser issues)
	fetchedObj, errFetch := client.OCSPRespondersAPI.FetchOCSPResponders(
		context.Background(),
		createdName,
		common.StringPtr("Prisma Access"),
		nil, nil,
	)
	require.NoError(t, errFetch, "Failed to fetch OCSP Responder after creation")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")

	defer func() {
		t.Logf("Cleaning up OCSP Responder with ID: %s", fetchedObj.Id)
		_, errDel := client.OCSPRespondersAPI.DeleteOCSPRespondersByID(context.Background(), fetchedObj.Id).Execute()
		if errDel != nil {
			t.Logf("Cleanup failed: %v", errDel)
		}
	}()

	t.Logf("Successfully created OCSP Responder: %s with ID: %s", createdName, fetchedObj.Id)
	assert.Equal(t, createdName, fetchedObj.Name, "Created OCSP responder name should match")
	assert.Equal(t, "ocsp.example.com", fetchedObj.HostName, "Hostname should match")
}

// Test_identity_services_OCSPRespondersAPIService_GetByID tests retrieving an OCSP responder by ID.
func Test_identity_services_OCSPRespondersAPIService_GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	ocspName := "test-ocsp-get-" + common.GenerateRandomString(6)

	ocspResponder := identity_services.OcspResponders{
		Folder:   common.StringPtr("Prisma Access"),
		Name:     ocspName,
		HostName: "ocsp-get.example.com",
	}

	_, err := client.OCSPRespondersAPI.CreateOCSPResponders(context.Background()).OcspResponders(ocspResponder).Execute()
	require.NoError(t, err, "Failed to create OCSP Responder for get test")

	// Use Fetch to get the ID (Create returns no model, List has deser issues)
	fetchedObj, errFetch := client.OCSPRespondersAPI.FetchOCSPResponders(
		context.Background(), ocspName, common.StringPtr("Prisma Access"), nil, nil,
	)
	require.NoError(t, errFetch, "Failed to fetch OCSP Responder")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	createdID := fetchedObj.Id

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
	client := SetupIdentitySvcTestClient(t)
	ocspName := "test-ocsp-update-" + common.GenerateRandomString(6)

	ocspResponder := identity_services.OcspResponders{
		Folder:   common.StringPtr("Prisma Access"),
		Name:     ocspName,
		HostName: "ocsp-original.example.com",
	}

	_, err := client.OCSPRespondersAPI.CreateOCSPResponders(context.Background()).OcspResponders(ocspResponder).Execute()
	require.NoError(t, err, "Failed to create OCSP Responder for update test")

	// Use Fetch to get the ID (Create returns no model, List has deser issues)
	fetchedObj, errFetch := client.OCSPRespondersAPI.FetchOCSPResponders(
		context.Background(), ocspName, common.StringPtr("Prisma Access"), nil, nil,
	)
	require.NoError(t, errFetch, "Failed to fetch OCSP Responder")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	createdID := fetchedObj.Id

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
	client := SetupIdentitySvcTestClient(t)
	ocspName := "test-ocsp-list-" + common.GenerateRandomString(6)

	ocspResponder := identity_services.OcspResponders{
		Folder:   common.StringPtr("Shared"),
		Name:     ocspName,
		HostName: "ocsp-list.example.com",
	}

	_, err := client.OCSPRespondersAPI.CreateOCSPResponders(context.Background()).OcspResponders(ocspResponder).Execute()
	require.NoError(t, err, "Failed to create OCSP Responder for list test")

	// Get the ID via Fetch for cleanup (List without Limit returns bare array)
	fetchedObj, errFetch := client.OCSPRespondersAPI.FetchOCSPResponders(
		context.Background(), ocspName, common.StringPtr("Shared"), nil, nil,
	)
	require.NoError(t, errFetch, "Failed to fetch OCSP Responder for list test")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	createdID := fetchedObj.Id

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
	client := SetupIdentitySvcTestClient(t)
	ocspName := "test-ocsp-delete-" + common.GenerateRandomString(6)

	ocspResponder := identity_services.OcspResponders{
		Folder:   common.StringPtr("Prisma Access"),
		Name:     ocspName,
		HostName: "ocsp-delete.example.com",
	}

	_, err := client.OCSPRespondersAPI.CreateOCSPResponders(context.Background()).OcspResponders(ocspResponder).Execute()
	require.NoError(t, err, "Failed to create OCSP Responder for delete test")

	// Use Fetch to get the ID (Create returns no model, List has deser issues)
	fetchedObj, errFetch := client.OCSPRespondersAPI.FetchOCSPResponders(
		context.Background(), ocspName, common.StringPtr("Prisma Access"), nil, nil,
	)
	require.NoError(t, errFetch, "Failed to fetch OCSP Responder")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")

	httpResDel, errDel := client.OCSPRespondersAPI.DeleteOCSPRespondersByID(context.Background(), fetchedObj.Id).Execute()
	require.NoError(t, errDel, "Failed to delete OCSP Responder")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_identity_services_OCSPRespondersAPIService_FetchOCSPResponders tests the FetchOCSPResponders convenience method
func Test_identity_services_OCSPRespondersAPIService_FetchOCSPResponders(t *testing.T) {
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

	// Test 1: Fetch existing object by name (also used to get ID for cleanup)
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
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchOCSPResponders found object: %s", fetchedObj.Name)

	// Cleanup after test using ID from fetch
	defer func() {
		deleteReq := client.OCSPRespondersAPI.DeleteOCSPRespondersByID(context.Background(), fetchedObj.Id)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", fetchedObj.Id)
	}()

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
