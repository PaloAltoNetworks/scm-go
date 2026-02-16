/*
Deployment Services Testing InternalDNSServersAPIService
*/
package deployment_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
)

// Test_deployment_services_InternalDNSServersAPIService_Create tests creating an internal DNS server
func Test_deployment_services_InternalDNSServersAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a test object
	testName := "test-dns-srv-" + common.GenerateRandomString(6)
	testObj := deployment_services.InternalDnsServers{
		Name:       testName,
		DomainName: []string{"example.com"},
		Primary:    "8.8.8.8",
	}

	createReq := client.InternalDNSServersAPI.CreateInternalDNSServers(context.Background()).InternalDnsServers(testObj)
	resp, httpResp, err := createReq.Execute()

	// Cleanup after test
	if resp != nil && resp.Id != "" {
		defer func() {
			deleteReq := client.InternalDNSServersAPI.DeleteInternalDNSServersByID(context.Background(), resp.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", resp.Id)
		}()
	}

	// Verify the response
	require.NoError(t, err, "Failed to create internal DNS server")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 201, httpResp.StatusCode, "Expected 201 Created status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testName, resp.Name, "Created DNS server name should match")
	assert.NotEmpty(t, resp.Id, "Created DNS server should have an ID")
	t.Logf("[SUCCESS] Created internal DNS server: %s (ID: %s)", resp.Name, resp.Id)
}

// Test_deployment_services_InternalDNSServersAPIService_GetByID tests getting an internal DNS server by ID
func Test_deployment_services_InternalDNSServersAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a test object first
	testName := "test-dns-srv-" + common.GenerateRandomString(6)
	testObj := deployment_services.InternalDnsServers{
		Name:       testName,
		DomainName: []string{"example.com"},
		Primary:    "8.8.8.8",
	}

	createReq := client.InternalDNSServersAPI.CreateInternalDNSServers(context.Background()).InternalDnsServers(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.InternalDNSServersAPI.DeleteInternalDNSServersByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Get the object by ID
	getReq := client.InternalDNSServersAPI.GetInternalDNSServersByID(context.Background(), createdID)
	resp, httpResp, err := getReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to get internal DNS server by ID")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, createdID, resp.Id, "Retrieved DNS server ID should match")
	assert.Equal(t, testName, resp.Name, "Retrieved DNS server name should match")
	t.Logf("[SUCCESS] Retrieved internal DNS server by ID: %s", resp.Id)
}

// Test_deployment_services_InternalDNSServersAPIService_Update tests updating an internal DNS server
func Test_deployment_services_InternalDNSServersAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a test object first
	testName := "test-dns-srv-" + common.GenerateRandomString(6)
	testObj := deployment_services.InternalDnsServers{
		Name:       testName,
		DomainName: []string{"example.com"},
		Primary:    "8.8.8.8",
	}

	createReq := client.InternalDNSServersAPI.CreateInternalDNSServers(context.Background()).InternalDnsServers(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.InternalDNSServersAPI.DeleteInternalDNSServersByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Update the object with different configuration
	updateObj := deployment_services.InternalDnsServers{
		Id:         createdID,
		Name:       testName,
		DomainName: []string{"example.com", "test.com"},
		Primary:    "1.1.1.1",
		Secondary:  common.StringPtr("8.8.4.4"),
	}

	updateReq := client.InternalDNSServersAPI.UpdateInternalDNSServersByID(context.Background(), createdID).InternalDnsServers(updateObj)
	resp, httpResp, err := updateReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to update internal DNS server")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, createdID, resp.Id, "Updated DNS server ID should match")
	assert.Equal(t, "1.1.1.1", resp.Primary, "Primary DNS should be updated")
	assert.Equal(t, 2, len(resp.DomainName), "Should have 2 domains")
	t.Logf("[SUCCESS] Updated internal DNS server: %s", resp.Id)
}

// Test_deployment_services_InternalDNSServersAPIService_List tests listing internal DNS servers
func Test_deployment_services_InternalDNSServersAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a test object first
	testName := "test-dns-srv-" + common.GenerateRandomString(6)
	testObj := deployment_services.InternalDnsServers{
		Name:       testName,
		DomainName: []string{"example.com"},
		Primary:    "8.8.8.8",
	}

	createReq := client.InternalDNSServersAPI.CreateInternalDNSServers(context.Background()).InternalDnsServers(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.InternalDNSServersAPI.DeleteInternalDNSServersByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// List DNS servers
	listReq := client.InternalDNSServersAPI.ListInternalDNSServers(context.Background())
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list internal DNS servers")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")

	// Verify our created object is in the list
	found := false
	if resp.Data != nil {
		for _, item := range resp.Data {
			if item.Id == createdID {
				found = true
				assert.Equal(t, testName, item.Name, "Listed DNS server name should match")
				break
			}
		}
	}
	assert.True(t, found, "Created DNS server should be in the list")
	t.Logf("[SUCCESS] Listed internal DNS servers, found test object: %s", createdID)
}

// Test_deployment_services_InternalDNSServersAPIService_DeleteByID tests deleting an internal DNS server
func Test_deployment_services_InternalDNSServersAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a test object first
	testName := "test-dns-srv-" + common.GenerateRandomString(6)
	testObj := deployment_services.InternalDnsServers{
		Name:       testName,
		DomainName: []string{"example.com"},
		Primary:    "8.8.8.8",
	}

	createReq := client.InternalDNSServersAPI.CreateInternalDNSServers(context.Background()).InternalDnsServers(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Delete the object
	deleteReq := client.InternalDNSServersAPI.DeleteInternalDNSServersByID(context.Background(), createdID)
	httpResp, err := deleteReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to delete internal DNS server")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	t.Logf("[SUCCESS] Deleted internal DNS server: %s", createdID)
}

// Test_deployment_services_InternalDNSServersAPIService_FetchInternalDNSServers tests the FetchInternalDNSServers convenience method
func Test_deployment_services_InternalDNSServersAPIService_FetchInternalDNSServers(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a test object first
	testName := "fetch-dns-srv-" + common.GenerateRandomString(6)
	testObj := deployment_services.InternalDnsServers{
		Name:       testName,
		DomainName: []string{"example.com"}, // Required field
		Primary:    "8.8.8.8",               // Required field
	}

	createReq := client.InternalDNSServersAPI.CreateInternalDNSServers(context.Background()).InternalDnsServers(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.InternalDNSServersAPI.DeleteInternalDNSServersByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.InternalDNSServersAPI.FetchInternalDNSServers(
		context.Background(),
		testName,
		common.StringPtr("Mobile Users"), // Deployment services typically use "Mobile Users"
		nil,                              // snippet
		nil,                              // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch internal_dns_servers by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchInternalDNSServers found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.InternalDNSServersAPI.FetchInternalDNSServers(
		context.Background(),
		"non-existent-internal_dns_servers-xyz-12345",
		common.StringPtr("Mobile Users"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchInternalDNSServers correctly returned nil for non-existent object")
}
