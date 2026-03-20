/*
Deployment Services Testing BandwidthAllocationsAPIService
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

// Test_deployment_services_BandwidthAllocationsAPIService_Create tests creating a bandwidth allocation
func Test_deployment_services_BandwidthAllocationsAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a test object
	testName := "test-bw-alloc-" + common.GenerateRandomString(6)
	testObj := deployment_services.BandwidthAllocations{
		Name:               testName,
		AllocatedBandwidth: 100,
	}

	createReq := client.BandwidthAllocationsAPI.CreateBandwidthAllocations(context.Background()).BandwidthAllocations(testObj)
	resp, httpResp, err := createReq.Execute()

	// Cleanup after test - BandwidthAllocations uses name-based deletion
	if resp != nil && resp.Name != "" {
		defer func() {
			deleteReq := client.BandwidthAllocationsAPI.DeleteBandwidthAllocations(context.Background()).Name(resp.Name)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", resp.Name)
		}()
	}

	// Verify the response
	require.NoError(t, err, "Failed to create bandwidth allocation")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status") // BandwidthAllocations returns 200, not 201
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testName, resp.Name, "Created allocation name should match")
	assert.Equal(t, int32(100), resp.AllocatedBandwidth, "Allocated bandwidth should match")
	t.Logf("[SUCCESS] Created bandwidth allocation: %s", resp.Name)
}

// Test_deployment_services_BandwidthAllocationsAPIService_List tests listing bandwidth allocations
func Test_deployment_services_BandwidthAllocationsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a test object first
	testName := "test-bw-alloc-" + common.GenerateRandomString(6)
	testObj := deployment_services.BandwidthAllocations{
		Name:               testName,
		AllocatedBandwidth: 100,
	}

	createReq := client.BandwidthAllocationsAPI.CreateBandwidthAllocations(context.Background()).BandwidthAllocations(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test
	defer func() {
		deleteReq := client.BandwidthAllocationsAPI.DeleteBandwidthAllocations(context.Background()).Name(testName)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", testName)
	}()

	// List bandwidth allocations
	listReq := client.BandwidthAllocationsAPI.ListBandwidthAllocations(context.Background())
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list bandwidth allocations")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")

	// Verify our created object is in the list
	found := false
	if resp.Data != nil {
		for _, item := range resp.Data {
			if item.Name == testName {
				found = true
				assert.Equal(t, int32(100), item.AllocatedBandwidth, "Listed allocation bandwidth should match")
				break
			}
		}
	}
	assert.True(t, found, "Created allocation should be in the list")
	t.Logf("[SUCCESS] Listed bandwidth allocations, found test object: %s", testName)
}

// Test_deployment_services_BandwidthAllocationsAPIService_Delete tests deleting a bandwidth allocation
func Test_deployment_services_BandwidthAllocationsAPIService_Delete(t *testing.T) {
	// Note: BandwidthAllocations Delete API requires both 'name' and 'spnNameList' parameters.
	// The spnNameList requires knowledge of Service Connection Point names associated with the allocation.
	// This test is skipped as it requires complex prerequisite setup.
	t.Skip("BandwidthAllocations Delete API requires spnNameList parameter with SPN context")
}

// Test_deployment_services_BandwidthAllocationsAPIService_FetchBandwidthAllocations tests the FetchBandwidthAllocations convenience method
func Test_deployment_services_BandwidthAllocationsAPIService_FetchBandwidthAllocations(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a test object first
	testName := "fetch-bw-alloc-" + common.GenerateRandomString(6)
	testObj := deployment_services.BandwidthAllocations{
		Name:               testName,
		AllocatedBandwidth: 100, // Required field - bandwidth in Mbps
	}

	createReq := client.BandwidthAllocationsAPI.CreateBandwidthAllocations(context.Background()).BandwidthAllocations(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test - BandwidthAllocations uses name-based deletion
	defer func() {
		deleteReq := client.BandwidthAllocationsAPI.DeleteBandwidthAllocations(context.Background()).Name(testName)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", testName)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.BandwidthAllocationsAPI.FetchBandwidthAllocations(
		context.Background(),
		testName,
		nil, // folder
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch bandwidth_allocations by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchBandwidthAllocations found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.BandwidthAllocationsAPI.FetchBandwidthAllocations(
		context.Background(),
		"non-existent-bandwidth_allocations-xyz-12345",
		nil,
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchBandwidthAllocations correctly returned nil for non-existent object")
}

// Test_deployment_services_BandwidthAllocationsAPIService_Update tests updating a bandwidth allocation
func Test_deployment_services_BandwidthAllocationsAPIService_Update(t *testing.T) {
	t.Skip("API requires spn_name_list parameter when updating by region - complex prerequisite setup needed")
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a test object first
	testName := "test-bw-alloc-update-" + common.GenerateRandomString(6)
	testObj := deployment_services.BandwidthAllocations{
		Name:               testName,
		AllocatedBandwidth: 100, // Initial bandwidth in Mbps
	}

	createReq := client.BandwidthAllocationsAPI.CreateBandwidthAllocations(context.Background()).BandwidthAllocations(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for update test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test - BandwidthAllocations uses name-based deletion
	defer func() {
		deleteReq := client.BandwidthAllocationsAPI.DeleteBandwidthAllocations(context.Background()).Name(testName)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", testName)
	}()

	// Prepare the updated object with modified fields
	updatedObj := deployment_services.BandwidthAllocations{
		Name:               testName,
		AllocatedBandwidth: 200, // Updated bandwidth (doubled)
	}

	// Test: Update the bandwidth allocation
	updateReq := client.BandwidthAllocationsAPI.UpdateBandwidthAllocations(context.Background()).BandwidthAllocations(updatedObj)
	updateRes, httpResUpdate, errUpdate := updateReq.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update bandwidth allocation")
	require.NotNil(t, httpResUpdate, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, testName, updateRes.Name, "Name should remain the same")
	assert.Equal(t, int32(200), updateRes.AllocatedBandwidth, "Allocated bandwidth should be updated to 200")
	t.Logf("[SUCCESS] Updated bandwidth allocation: %s with new bandwidth: %d", updateRes.Name, updateRes.AllocatedBandwidth)
}
