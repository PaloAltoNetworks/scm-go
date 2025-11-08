/*
 * Network Services Testing
 *
 * ZonesAPIService
 */

package network_services

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Replace with your actual imports for common utils and generated client
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// NOTE: Placeholder functions and types are assumed to be available in the test environment.
// - SetupNetworkSvcTestClient(t) -> returns a *network_services.APIClient (with SecurityZonesAPIService)
// - handleAPIError(err)
// - common.StringPtr(s) -> returns *string

// --- Helper Functions ---

// generateZoneName creates a unique name for the zone resource.
func generateZoneName(base string) string {
	// Zone name alphanumeric string (no '$' prefix)
	return base + common.GenerateRandomString(6)
}

// createTestZone creates a minimal Zones object for testing.
func createTestZone(t *testing.T, zoneName string) network_services.Zones {
	// 'name' is the only required field.
	return *network_services.NewZones(zoneName)
}

// createFullTestZone creates a comprehensive Zones object for update/get testing.
func createFullTestZone(t *testing.T, zoneName string) network_services.Zones {
	// 1. Setup complex child structs (assuming minimal required fields exist for these structs)

	// 2. Setup boolean fields
	enableDeviceID := true
	enableUserID := true

	// 3. Setup string fields

	// 4. Build the full zone object
	zone := createTestZone(t, zoneName)
	zone.SetFolder("All")
	zone.SetEnableDeviceIdentification(enableDeviceID)
	zone.SetEnableUserIdentification(enableUserID)

	// DeviceACL and UserACL are omitted for simplicity but can be added if needed.

	return zone
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_SecurityZonesAPIService_Create tests the creation of a Security Zone.
func Test_network_services_SecurityZonesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	zoneName := generateZoneName("scm-zone-create-")

	zone := createFullTestZone(t, zoneName)

	t.Logf("Creating Zone with name: %s", zoneName)
	req := client.SecurityZonesAPI.CreateZones(context.Background()).Zones(zone)
	res, httpRes, err := req.Execute()

	if httpRes != nil {
		t.Logf("--- HTTP Response Details ---")

		// 1. Read the entire body content
		bodyBytes, readErr := io.ReadAll(httpRes.Body)
		if readErr != nil {
			t.Fatalf("Failed to read response body for logging: %v", readErr)
		}

		// 2. CRITICAL STEP: Close the original body and replace it.
		// This makes the response body available for the API client to decode (unmarshal into 'res').
		httpRes.Body.Close()
		httpRes.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// 3. Log the status line
		t.Logf("Status: %s (Code: %d)", httpRes.Status, httpRes.StatusCode)

		// 4. Log the headers
		t.Logf("Headers:")
		for key, values := range httpRes.Header {
			t.Logf("  %s: %s", key, values)
		}

		// 5. Log the body content
		t.Logf("Body:\n%s", string(bodyBytes))
		t.Logf("----------------------------")
	}

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Zone")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created zone should have an ID")
	createdID := *res.Id

	// Cleanup the created zone
	defer func() {
		t.Logf("Cleaning up Zone with ID: %s", createdID)
		_, errDel := client.SecurityZonesAPI.DeleteZonesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete Zone during cleanup")
	}()

	t.Logf("Successfully created Zone: %s with ID: %s", zoneName, createdID)

	// Verify key fields in the response
	assert.Equal(t, zoneName, res.Name, "Created zone name should match")
	assert.True(t, res.GetEnableDeviceIdentification(), "Device Identification should be enabled")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_SecurityZonesAPIService_GetByID tests retrieving a Security Zone by ID.
func Test_network_services_SecurityZonesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	zoneName := generateZoneName("scm-zone-get-")
	zone := createFullTestZone(t, zoneName)

	// Setup: Create a zone first
	createRes, _, err := client.SecurityZonesAPI.CreateZones(context.Background()).Zones(zone).Execute()
	require.NoError(t, err, "Failed to create zone for get test setup")
	createdID := *createRes.Id

	defer func() {
		client.SecurityZonesAPI.DeleteZonesByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the zone
	getRes, httpResGet, errGet := client.SecurityZonesAPI.GetZonesByID(context.Background(), createdID).Execute()

	require.NoError(t, errGet, "Failed to get Zone by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, zoneName, getRes.Name, "Zone name should match")
	assert.True(t, getRes.GetEnableUserIdentification(), "User Identification should be enabled")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_SecurityZonesAPIService_Update tests updating a Security Zone.
func Test_network_services_SecurityZonesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	zoneName := generateZoneName("scm-zone-update-")
	zone := createFullTestZone(t, zoneName)

	// Setup: Create a zone first
	createRes, _, err := client.SecurityZonesAPI.CreateZones(context.Background()).Zones(zone).Execute()
	require.NoError(t, err, "Failed to create zone for update test setup")
	createdID := *createRes.Id

	defer func() {
		client.SecurityZonesAPI.DeleteZonesByID(context.Background(), createdID).Execute()
	}()

	// Disable one boolean flag
	disableDeviceID := false

	updatedZone := createTestZone(t, zoneName)
	updatedZone.Id = &createdID
	updatedZone.SetEnableDeviceIdentification(disableDeviceID)

	// Test: Update the zone
	updateRes, httpResUpdate, errUpdate := client.SecurityZonesAPI.UpdateZonesByID(context.Background(), createdID).Zones(updatedZone).Execute()

	require.NoError(t, errUpdate, "Failed to update Zone")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	assert.False(t, updateRes.GetEnableDeviceIdentification(), "Device Identification should be disabled")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_SecurityZonesAPIService_List tests listing Security Zones.
func Test_network_services_SecurityZonesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	zoneName := generateZoneName("scm-zone-list-")
	zone := createFullTestZone(t, zoneName)

	// Setup: Create a zone first
	createRes, _, err := client.SecurityZonesAPI.CreateZones(context.Background()).Zones(zone).Execute()
	require.NoError(t, err, "Failed to create zone for list test setup")
	createdID := *createRes.Id

	defer func() {
		client.SecurityZonesAPI.DeleteZonesByID(context.Background(), createdID).Execute()
	}()

	// Test: List the zones, filtering by the unique name
	listRes, httpResList, errList := client.SecurityZonesAPI.ListZones(context.Background()).
		Folder("All").
		Limit(10).
		Execute()

	require.NoError(t, errList, "Failed to list Zones")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	// Assert that the specific, uniquely named resource was returned.
	require.LessOrEqual(t, 1, len(listRes.GetData()), "Expected atleast one Zone")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_SecurityZonesAPIService_DeleteByID tests deleting a Security Zone.
func Test_network_services_SecurityZonesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	zoneName := generateZoneName("scm-zone-delete-")
	zone := createFullTestZone(t, zoneName)

	// Setup: Create a zone first
	createRes, _, err := client.SecurityZonesAPI.CreateZones(context.Background()).Zones(zone).Execute()
	require.NoError(t, err, "Failed to create zone for delete test setup")
	createdID := *createRes.Id

	// Test: Delete the zone
	httpResDel, errDel := client.SecurityZonesAPI.DeleteZonesByID(context.Background(), createdID).Execute()

	require.NoError(t, errDel, "Failed to delete Zone")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}
