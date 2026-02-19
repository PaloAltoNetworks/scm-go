/*
Objects Testing QuarantinedDevicesAPIService
Note: This API has limited functionality - Create, Delete (by HostId), List only
No GetByID, Update, or Fetch methods available
*/
package objects

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)

// Test_objects_QuarantinedDevicesAPIService_Create tests the creation of a quarantined device
// Skip: Requires real device serial numbers connected to the system
func Test_objects_QuarantinedDevicesAPIService_Create(t *testing.T) {
	t.Skip("Quarantined devices API requires real connected device serial numbers")
	client := SetupObjectSvcTestClient(t)

	// Generate a unique host ID for testing
	testHostId := "test-host-create-" + common.GenerateRandomString(10)
	device := objects.QuarantinedDevices{
		HostId:       testHostId,
		SerialNumber: common.StringPtr("SN-" + common.GenerateRandomString(8)),
	}

	req := client.QuarantinedDevicesAPI.CreateQuarantinedDevices(context.Background()).QuarantinedDevices(device)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create quarantined device")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, testHostId, res.HostId, "Created device host ID should match")

	t.Logf("Successfully created quarantined device with host ID: %s", testHostId)

	// Cleanup - uses HostId-based delete
	reqDel := client.QuarantinedDevicesAPI.DeleteQuarantinedDevices(context.Background()).HostId(testHostId)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete quarantined device during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
	t.Logf("Successfully cleaned up quarantined device: %s", testHostId)
}

// Test_objects_QuarantinedDevicesAPIService_List tests listing quarantined devices
// Skip: Requires real device serial numbers connected to the system
func Test_objects_QuarantinedDevicesAPIService_List(t *testing.T) {
	t.Skip("Quarantined devices API requires real connected device serial numbers")
	client := SetupObjectSvcTestClient(t)

	testHostId := "test-host-list-" + common.GenerateRandomString(10)
	device := objects.QuarantinedDevices{
		HostId:       testHostId,
		SerialNumber: common.StringPtr("SN-" + common.GenerateRandomString(8)),
	}

	req := client.QuarantinedDevicesAPI.CreateQuarantinedDevices(context.Background()).QuarantinedDevices(device)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create quarantined device for list test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Test List
	reqList := client.QuarantinedDevicesAPI.ListQuarantinedDevices(context.Background())
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	require.NoError(t, errList, "Failed to list quarantined devices")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	// Note: ListQuarantinedDevices returns []QuarantinedDevices directly, not a wrapper
	assert.GreaterOrEqual(t, len(listRes), 0, "Should return a list (may be empty)")

	t.Logf("Successfully listed quarantined devices, count: %d", len(listRes))

	// Cleanup
	reqDel := client.QuarantinedDevicesAPI.DeleteQuarantinedDevices(context.Background()).HostId(testHostId)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete quarantined device during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
	t.Logf("Successfully cleaned up quarantined device: %s", testHostId)
}

// Test_objects_QuarantinedDevicesAPIService_Delete tests deleting a quarantined device by host ID
// Skip: Requires real device serial numbers connected to the system
func Test_objects_QuarantinedDevicesAPIService_Delete(t *testing.T) {
	t.Skip("Quarantined devices API requires real connected device serial numbers")
	client := SetupObjectSvcTestClient(t)

	testHostId := "test-host-delete-" + common.GenerateRandomString(10)
	device := objects.QuarantinedDevices{
		HostId:       testHostId,
		SerialNumber: common.StringPtr("SN-" + common.GenerateRandomString(8)),
	}

	req := client.QuarantinedDevicesAPI.CreateQuarantinedDevices(context.Background()).QuarantinedDevices(device)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create quarantined device for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Test Delete by HostId
	reqDel := client.QuarantinedDevicesAPI.DeleteQuarantinedDevices(context.Background()).HostId(testHostId)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	require.NoError(t, errDel, "Failed to delete quarantined device")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted quarantined device: %s", testHostId)
}
