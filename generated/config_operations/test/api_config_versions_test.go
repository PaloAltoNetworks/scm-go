/*
Config Operations Testing ConfigVersionsAPIService
*/
package config_operations

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_config_operations_ConfigVersionsAPIService_List tests listing config versions
// This is a read-only operation that retrieves the list of configuration versions
func Test_config_operations_ConfigVersionsAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigOperationsTestClient(t)

	fmt.Printf("Listing configuration versions\n")

	// Make the list request to the API
	req := client.ConfigVersionsAPI.ListConfigVersions(context.Background())
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the request was successful
	require.NoError(t, err, "Failed to list config versions")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")

	// Assert response properties
	require.NotNil(t, res, "Response should not be nil")

	// Get the data from the paginated response
	versions := res.Data
	require.NotNil(t, versions, "Data array should not be nil")

	fmt.Printf("Successfully retrieved %d config versions (limit: %d, offset: %d, total: %d)\n",
		len(versions), res.Limit, res.Offset, res.Total)

	// If there are versions, verify the structure
	if len(versions) > 0 {
		firstVersion := versions[0]
		assert.NotZero(t, firstVersion.Id, "Version should have an ID")
		assert.NotEmpty(t, firstVersion.Version, "Version should have a version string")

		fmt.Printf("Sample version - ID: %d, Version: %s, Date: %s\n",
			firstVersion.Id, firstVersion.Version, firstVersion.Date)
	} else {
		fmt.Printf("No config versions found in the system\n")
	}

	t.Logf("Successfully listed config versions")
}

// Test_config_operations_ConfigVersionsAPIService_GetByID tests retrieving a specific config version by ID
// Note: This test uses a hardcoded version number since we can't easily extract from the list response
func Test_config_operations_ConfigVersionsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigOperationsTestClient(t)

	// Use version 1 as a test (this might not exist in all environments)
	versionID := int32(1)
	fmt.Printf("Testing GetByID with version: %d\n", versionID)

	// Retrieve the specific version by ID
	req := client.ConfigVersionsAPI.GetConfigVersionsByID(context.Background(), versionID)
	res, httpRes, err := req.Execute()

	// This test may fail if version 1 doesn't exist, which is acceptable
	if err != nil {
		handleAPIError(err)
		t.Skipf("Version %d not found - this is expected if no configs exist", versionID)
		return
	}

	// Verify the request was successful
	require.NoError(t, err, "Failed to get config version by ID")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")

	// Assert response properties
	require.NotNil(t, res, "Response should not be nil")

	// The API returns an array with config versions (even for single ID lookup)
	require.Greater(t, len(res), 0, "Should have at least one config version in response")

	// Get the first version from the array
	version := res[0]
	assert.NotZero(t, version.Id, "Version should have an ID")
	assert.NotEmpty(t, version.Version, "Version should have a version string")

	fmt.Printf("Retrieved config version - ID: %d, Version: %s, Date: %s, Admin: %s\n",
		version.Id, version.Version, version.Date, version.Admin)

	t.Logf("Successfully retrieved config version by ID: %d", versionID)
}

// Test_config_operations_ConfigVersionsAPIService_GetRunning tests retrieving the running configuration version
// This is a read-only operation that retrieves the currently active configuration
func Test_config_operations_ConfigVersionsAPIService_GetRunning(t *testing.T) {
	// Setup the authenticated client
	client := SetupConfigOperationsTestClient(t)

	fmt.Printf("Retrieving running configuration version\n")

	// Make the request to get running config
	req := client.ConfigVersionsAPI.GetRunningConfigVersions(context.Background())
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the request was successful
	require.NoError(t, err, "Failed to get running config version")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")

	// Assert response properties
	require.NotNil(t, res, "Response should not be nil")

	// Get the data from the paginated response
	runningVersions := res.Data
	require.NotNil(t, runningVersions, "Data array should not be nil")
	require.Greater(t, len(runningVersions), 0, "Should have at least one running version")

	fmt.Printf("Retrieved %d running config versions (limit: %d, offset: %d, total: %d)\n",
		len(runningVersions), res.Limit, res.Offset, res.Total)

	// Verify the structure of the first running version
	firstRunning := runningVersions[0]
	assert.NotEmpty(t, firstRunning.Device, "Running version should have a device")
	assert.NotZero(t, firstRunning.Version, "Running version should have a version number")

	fmt.Printf("Sample running version - Device: %s, Version: %d, Date: %s\n",
		firstRunning.Device, firstRunning.Version, firstRunning.Date)

	t.Logf("Successfully retrieved %d running config version(s)", len(runningVersions))
}

// NOTE: The following operations are NOT tested as they are destructive/action operations:
// - LoadConfigVersions: This loads a candidate config (action operation)
// - PushCandidateConfigVersions: This pushes config to devices (action operation)
// - DeleteCandidateConfigVersions: This deletes the candidate config (destructive operation)
//
// These operations should be tested in integration tests or manually in a controlled environment.
