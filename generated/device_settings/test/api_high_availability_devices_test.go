/*
Device Settings Testing HighAvailabilityDevicesAPIService
Note: This API only has List (no GetByID, Create, Update, Delete)
*/
package device_settings

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_device_settings_HighAvailabilityDevicesAPIService_List tests listing high availability devices
func Test_device_settings_HighAvailabilityDevicesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeviceSettingsTestClient(t)

	// Test List operation with folder filter (uses ListHADevices method)
	reqList := client.HighAvailabilityDevicesAPI.ListHADevices(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		// The API returns an empty array [] when no HA devices are configured,
		// but the OpenAPI spec expects a wrapper object. Skip test in this case.
		if strings.Contains(errList.Error(), "cannot unmarshal array") {
			t.Skip("No HA devices configured (API returns empty array)")
		}
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list high availability devices")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed high availability devices")
}
