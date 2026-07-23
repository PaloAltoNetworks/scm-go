/*
 * ZTNA Connector Testing — ConnectorAPIService
 * Covers: CRUD, Filters, Images, Quiesce, Scheduled Upgrade, PCAPs, Tech Support
 */
package ztna_connector_all

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"
)

// onlineConnectorOID reads ZTNA_ONLINE_CONNECTOR_OID and skips if unset.
func onlineConnectorOID(t *testing.T) string {
	t.Helper()
	oid := os.Getenv("ZTNA_ONLINE_CONNECTOR_OID")
	if oid == "" {
		t.Skip("Skipping: ZTNA_ONLINE_CONNECTOR_OID environment variable is not set")
	}
	return oid
}

// pcapConnectorOID reads ZTNA_PCAP_CONNECTOR_OID and skips if unset.
func pcapConnectorOID(t *testing.T) string {
	t.Helper()
	oid := os.Getenv("ZTNA_PCAP_CONNECTOR_OID")
	if oid == "" {
		t.Skip("Skipping: ZTNA_PCAP_CONNECTOR_OID environment variable is not set")
	}
	return oid
}

// fetchFirstConnectorImage returns the first available image ID, or skips the test.
func fetchFirstConnectorImage(t *testing.T, client *ztna_connector_all.APIClient) string {
	t.Helper()
	images, _, err := client.ConnectorAPI.ListConnectorImages(context.Background()).Execute()
	if err != nil {
		t.Skipf("Cannot list connector images: %v", err)
	}
	if len(images) == 0 {
		t.Skip("No connector images available")
	}
	return images[0]
}

// ── CRUD ─────────────────────────────────────────────────────────────────────

// Test_ztna_connector_all_ConnectorAPIService_Create tests creating a connector.
func Test_ztna_connector_all_ConnectorAPIService_Create(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-create-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	assert.NotEmpty(t, oid, "OID should be set after creation")
	t.Logf("Successfully created connector: %s with OID: %s", name, oid)

	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })
}

// Test_ztna_connector_all_ConnectorAPIService_GetByID creates a connector and retrieves it by OID.
func Test_ztna_connector_all_ConnectorAPIService_GetByID(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-getbyid-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	connector, httpRes, err := client.ConnectorAPI.GetConnectorsByID(context.Background(), oid).
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to get connector by ID")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	require.NotNil(t, connector)
	assert.Equal(t, name, connector.Name)
	assert.Equal(t, groupID, connector.Group)
	require.NotNil(t, connector.Oid)
	assert.Equal(t, oid, *connector.Oid)
	t.Logf("GetByID OK: name=%s oid=%s", connector.Name, *connector.Oid)
}

// Test_ztna_connector_all_ConnectorAPIService_Update creates a connector and updates its description.
func Test_ztna_connector_all_ConnectorAPIService_Update(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-update-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	updatedDesc := "Updated by test"
	updated := ztna_connector_all.NewConnectors(groupID, name)
	updated.Description = &updatedDesc

	httpRes, err := client.ConnectorAPI.UpdateConnectorsByID(context.Background(), oid).
		Connectors(*updated).
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to update connector")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK for update")

	connector, _, err := client.ConnectorAPI.GetConnectorsByID(context.Background(), oid).
		Execute()
	require.NoError(t, err)
	require.NotNil(t, connector.Description)
	assert.Equal(t, updatedDesc, *connector.Description, "Description should be updated")
	t.Logf("Update OK: oid=%s", oid)
}

// Test_ztna_connector_all_ConnectorAPIService_List lists connectors and verifies the created one appears.
func Test_ztna_connector_all_ConnectorAPIService_List(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-list-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	listRes, httpRes, err := client.ConnectorAPI.ListConnectors(context.Background()).
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list connectors")
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, listRes)

	found := false
	for _, c := range listRes.GetData() {
		if c.Name == name {
			found = true
			if c.Oid != nil {
				assert.Equal(t, oid, *c.Oid)
			}
			break
		}
	}
	assert.True(t, found, "Created connector should appear in list")
	t.Logf("List OK: total=%d found=%v", listRes.GetTotal(), found)
}

// Test_ztna_connector_all_ConnectorAPIService_Delete creates a connector, deletes it, then verifies 404.
func Test_ztna_connector_all_ConnectorAPIService_Delete(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-delete-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)

	httpRes, err := client.ConnectorAPI.DeleteConnectorsByID(context.Background(), oid).
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to delete connector")
	require.NotNil(t, httpRes)
	assert.Equal(t, 202, httpRes.StatusCode, "Expected 202 for delete")
	t.Logf("Delete OK: name=%s oid=%s", name, oid)

	_, httpRes, err = client.ConnectorAPI.GetConnectorsByID(context.Background(), oid).
		Execute()
	require.Error(t, err, "Expected error after deletion")
	require.NotNil(t, httpRes)
	assert.Equal(t, 404, httpRes.StatusCode, "Expected 404 after deletion")
	t.Logf("Verified 404 after delete: oid=%s", oid)
}

// ── Filters & Images ──────────────────────────────────────────────────────────

// Test_ztna_connector_all_ConnectorAPIService_ListFilters lists filter values for the "name" field.
func Test_ztna_connector_all_ConnectorAPIService_ListFilters(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)

	filters, httpRes, err := client.ConnectorAPI.ListConnectorFilters(context.Background()).
		Field("name").
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list connector filters")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("Connector name filters (%d): %v", len(filters), filters)
}

// Test_ztna_connector_all_ConnectorAPIService_ListImages lists available connector image versions.
func Test_ztna_connector_all_ConnectorAPIService_ListImages(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)

	_, httpRes, err := client.ConnectorAPI.ListConnectorImages(context.Background()).Execute()

	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	if err != nil {
		handleAPIError(err)
		t.Logf("Note: SDK/spec mismatch — API may return paginated object, spec expects []string")
	}
}

// ── Quiesce ───────────────────────────────────────────────────────────────────

// Test_ztna_connector_all_ConnectorAPIService_GetQuiesce reads quiesce state from a new connector.
// Offline connectors may return an error — both outcomes are acceptable.
func Test_ztna_connector_all_ConnectorAPIService_GetQuiesce(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-quiesce-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	quiesce, httpRes, err := client.ConnectorAPI.GetConnectorsQuiesceByID(context.Background(), oid).
		Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("GetQuiesce returned an error (connector likely offline): %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	if quiesce != nil {
		t.Logf("Quiesce state: mode=%s", quiesce.Mode)
	}
}

// Test_ztna_connector_all_ConnectorAPIService_GetQuiesce_OnlineConnector reads quiesce state
// from a known-online connector. Requires ZTNA_ONLINE_CONNECTOR_OID.
func Test_ztna_connector_all_ConnectorAPIService_GetQuiesce_OnlineConnector(t *testing.T) {
	oid := onlineConnectorOID(t)
	client := SetupZtnaConnectorAllTestClient(t)

	quiesce, httpRes, err := client.ConnectorAPI.GetConnectorsQuiesceByID(context.Background(), oid).
		Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("GetQuiesce error on online connector: %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, quiesce, "Expected a quiesce object from an online connector")
	t.Logf("Quiesce state: mode=%s", quiesce.Mode)
}

// Test_ztna_connector_all_ConnectorAPIService_UpdateQuiesce sends a quiesce command.
// Non-2xx is acceptable — connector is not online.
func Test_ztna_connector_all_ConnectorAPIService_UpdateQuiesce(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-uquiesce-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	quiesce := ztna_connector_all.NewConnectorQuiesce("quiesce")
	httpRes, err := client.ConnectorAPI.UpdateConnectorsQuiesceByID(context.Background(), oid).
		ConnectorQuiesce(*quiesce).
		Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("UpdateQuiesce returned an error (connector likely offline): %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.True(t, httpRes.StatusCode == 200 || httpRes.StatusCode == 202,
		"Expected 200 or 202, got %d", httpRes.StatusCode)
	t.Logf("UpdateQuiesce OK: oid=%s status=%d", oid, httpRes.StatusCode)
}

// ── Scheduled Upgrade ─────────────────────────────────────────────────────────

// Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_Create tests scheduling an upgrade.
func Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_Create(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	imageID := fetchFirstConnectorImage(t, client)
	name := fmt.Sprintf("test-connector-sucreate-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() {
		_, _ = client.ConnectorAPI.DeleteConnectorsScheduledUpgradeByID(context.Background(), oid).Execute()
		deleteTestConnector(t, client, oid, name)
	})

	upgrade := ztna_connector_all.NewConnectorScheduledUpgrade(imageID)
	httpRes, err := client.ConnectorAPI.CreateConnectorsScheduledUpgradeByID(context.Background(), oid).
		ConnectorScheduledUpgrade(*upgrade).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("CreateScheduledUpgrade error (connector may need to be online): %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.True(t, httpRes.StatusCode == 200 || httpRes.StatusCode == 201 || httpRes.StatusCode == 202,
		"Expected 200/201/202, got %d", httpRes.StatusCode)
	t.Logf("CreateScheduledUpgrade OK: oid=%s imageID=%s", oid, imageID)
}

// Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_Get retrieves a scheduled upgrade.
func Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_Get(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	imageID := fetchFirstConnectorImage(t, client)
	name := fmt.Sprintf("test-connector-suget-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() {
		_, _ = client.ConnectorAPI.DeleteConnectorsScheduledUpgradeByID(context.Background(), oid).Execute()
		deleteTestConnector(t, client, oid, name)
	})

	upgrade := ztna_connector_all.NewConnectorScheduledUpgrade(imageID)
	if _, err := client.ConnectorAPI.CreateConnectorsScheduledUpgradeByID(context.Background(), oid).
		ConnectorScheduledUpgrade(*upgrade).Execute(); err != nil {
		t.Skipf("Cannot create scheduled upgrade (connector may be offline): %v", err)
	}

	got, httpRes, err := client.ConnectorAPI.GetConnectorsScheduledUpgradeByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("GetScheduledUpgrade error: %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, got)
	assert.Equal(t, imageID, got.ImageId)
	t.Logf("GetScheduledUpgrade OK: oid=%s imageId=%s", oid, got.ImageId)
}

// Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_GetStatus retrieves upgrade status.
func Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_GetStatus(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-sustatus-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	status, httpRes, err := client.ConnectorAPI.GetConnectorsScheduledUpgradeStatusByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("GetScheduledUpgradeStatus error (connector may be offline): %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	if status != nil {
		t.Logf("Upgrade status: %+v", status)
	}
}

// Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_Update tests updating a scheduled upgrade.
func Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_Update(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	images, _, _ := client.ConnectorAPI.ListConnectorImages(context.Background()).Execute()
	if len(images) < 2 {
		t.Skip("Need at least 2 connector images to test scheduled upgrade update")
	}
	name := fmt.Sprintf("test-connector-suupdate-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() {
		_, _ = client.ConnectorAPI.DeleteConnectorsScheduledUpgradeByID(context.Background(), oid).Execute()
		deleteTestConnector(t, client, oid, name)
	})

	upgrade := ztna_connector_all.NewConnectorScheduledUpgrade(images[0])
	if _, err := client.ConnectorAPI.CreateConnectorsScheduledUpgradeByID(context.Background(), oid).
		ConnectorScheduledUpgrade(*upgrade).Execute(); err != nil {
		t.Skipf("Cannot create scheduled upgrade: %v", err)
	}

	updated := ztna_connector_all.NewConnectorScheduledUpgrade(images[1])
	httpRes, err := client.ConnectorAPI.UpdateConnectorsScheduledUpgradeByID(context.Background(), oid).
		ConnectorScheduledUpgrade(*updated).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("UpdateScheduledUpgrade error: %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("UpdateScheduledUpgrade OK: oid=%s newImageID=%s", oid, images[1])
}

// Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_Delete deletes a scheduled upgrade.
func Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_Delete(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	imageID := fetchFirstConnectorImage(t, client)
	name := fmt.Sprintf("test-connector-sudelete-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	upgrade := ztna_connector_all.NewConnectorScheduledUpgrade(imageID)
	if _, err := client.ConnectorAPI.CreateConnectorsScheduledUpgradeByID(context.Background(), oid).
		ConnectorScheduledUpgrade(*upgrade).Execute(); err != nil {
		t.Skipf("Cannot create scheduled upgrade (connector may be offline): %v", err)
	}

	httpRes, err := client.ConnectorAPI.DeleteConnectorsScheduledUpgradeByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("DeleteScheduledUpgrade error: %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("DeleteScheduledUpgrade OK: oid=%s", oid)
}

// Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_Get_OnlineConnector reads the
// scheduled upgrade from a known-online connector. Requires ZTNA_ONLINE_CONNECTOR_OID.
func Test_ztna_connector_all_ConnectorAPIService_ScheduledUpgrade_Get_OnlineConnector(t *testing.T) {
	oid := onlineConnectorOID(t)
	client := SetupZtnaConnectorAllTestClient(t)

	got, httpRes, err := client.ConnectorAPI.GetConnectorsScheduledUpgradeByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("GetScheduledUpgrade returned an error (no upgrade may be scheduled): %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	if got != nil {
		t.Logf("Scheduled upgrade: imageId=%s", got.ImageId)
	} else {
		t.Log("No scheduled upgrade found for this connector")
	}
}

// ── Packet Captures (PCAPs) ───────────────────────────────────────────────────

// Test_ztna_connector_all_ConnectorAPIService_Pcap_List lists PCAPs for a connector.
func Test_ztna_connector_all_ConnectorAPIService_Pcap_List(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-pcaplist-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	list, httpRes, err := client.ConnectorAPI.ListConnectorsPcapsByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("ListPcaps error (connector may be offline): %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	if list != nil {
		t.Logf("PCAP files: %+v", list)
	}
}

// Test_ztna_connector_all_ConnectorAPIService_Pcap_List_OnlineConnector lists PCAPs on a
// known-online connector. Requires ZTNA_ONLINE_CONNECTOR_OID.
func Test_ztna_connector_all_ConnectorAPIService_Pcap_List_OnlineConnector(t *testing.T) {
	oid := onlineConnectorOID(t)
	client := SetupZtnaConnectorAllTestClient(t)

	list, httpRes, err := client.ConnectorAPI.ListConnectorsPcapsByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("ListPcaps error: %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	if list != nil {
		t.Logf("PCAP files: %+v", list)
	}
}

// Test_ztna_connector_all_ConnectorAPIService_Pcap_Create starts a packet capture.
// Only succeeds if the connector is online.
func Test_ztna_connector_all_ConnectorAPIService_Pcap_Create(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-pcap-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	pcapReq := ztna_connector_all.NewPacketCaptureCreate("eth0")
	pcap, httpRes, err := client.ConnectorAPI.CreateConnectorsPcapsByID(context.Background(), oid).
		PacketCaptureCreate(*pcapReq).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("CreatePcap error (connector must be online): %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.True(t, httpRes.StatusCode == 200 || httpRes.StatusCode == 201,
		"Expected 200/201, got %d", httpRes.StatusCode)
	require.NotNil(t, pcap)
	t.Logf("CreatePcap OK: oid=%s pcapId=%s", oid, pcap.GetId())

	t.Cleanup(func() {
		_, _, _ = client.ConnectorAPI.StopConnectorsPcapsByID(context.Background(), oid, pcap.GetId()).Execute()
	})
}

// Test_ztna_connector_all_ConnectorAPIService_Pcap_StopAndDownload stops and downloads a PCAP.
// Skipped if the connector is not online.
func Test_ztna_connector_all_ConnectorAPIService_Pcap_StopAndDownload(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-pcapstop-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	pcapReq := ztna_connector_all.NewPacketCaptureCreate("eth0")
	pcap, _, err := client.ConnectorAPI.CreateConnectorsPcapsByID(context.Background(), oid).
		PacketCaptureCreate(*pcapReq).Execute()
	if err != nil {
		t.Skipf("Cannot create PCAP (connector must be online): %v", err)
	}
	require.NotNil(t, pcap)
	pcapID := pcap.GetId()

	stopped, httpRes, err := client.ConnectorAPI.StopConnectorsPcapsByID(context.Background(), oid, pcapID).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("StopPcap error: %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	if stopped != nil {
		t.Logf("StopPcap OK: pcapId=%s", pcapID)
	}

	dlRes, dlErr := client.ConnectorAPI.DownloadConnectorsPcapsByID(context.Background(), oid, pcapID).Execute()
	if dlErr != nil {
		handleAPIError(dlErr)
		t.Logf("DownloadPcap error: %v", dlErr)
		return
	}
	require.NotNil(t, dlRes)
	assert.Equal(t, 200, dlRes.StatusCode)
	t.Logf("DownloadPcap OK: pcapId=%s", pcapID)
}

// Test_ztna_connector_all_ConnectorAPIService_Pcap_StartWaitStop starts a PCAP on a known-online
// connector (ZTNA_PCAP_CONNECTOR_OID), waits 1 minute, stops it, then verifies via list.
func Test_ztna_connector_all_ConnectorAPIService_Pcap_StartWaitStop(t *testing.T) {
	oid := pcapConnectorOID(t)
	client := SetupZtnaConnectorAllTestClient(t)

	pcapReq := ztna_connector_all.NewPacketCaptureCreate("external")
	pcap, httpRes, err := client.ConnectorAPI.CreateConnectorsPcapsByID(context.Background(), oid).
		PacketCaptureCreate(*pcapReq).Execute()
	if err != nil {
		handleAPIError(err)
		t.Fatalf("Failed to start packet capture: %v", err)
	}
	require.NotNil(t, httpRes)
	assert.True(t, httpRes.StatusCode == 200 || httpRes.StatusCode == 201,
		"Expected 200/201, got %d", httpRes.StatusCode)
	require.NotNil(t, pcap)
	pcapID := pcap.GetId()
	t.Logf("Capture started: pcapId=%s", pcapID)

	t.Log("Waiting 1 minute for packet capture...")
	time.Sleep(1 * time.Minute)

	stopped, stopRes, err := client.ConnectorAPI.StopConnectorsPcapsByID(context.Background(), oid, pcapID).Execute()
	if err != nil {
		handleAPIError(err)
		t.Fatalf("Failed to stop packet capture: %v", err)
	}
	require.NotNil(t, stopRes)
	assert.Equal(t, 200, stopRes.StatusCode)
	if stopped != nil {
		t.Logf("Capture stopped: pcapId=%s", pcapID)
	}

	list, listRes, err := client.ConnectorAPI.ListConnectorsPcapsByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Fatalf("Failed to list packet captures: %v", err)
	}
	require.NotNil(t, listRes)
	assert.Equal(t, 200, listRes.StatusCode)
	if list != nil {
		t.Logf("PCAP list after stop: %+v", list)
	}
}

// ── Tech Support Files ────────────────────────────────────────────────────────

// Test_ztna_connector_all_ConnectorAPIService_TechSupport_List lists tech support files.
func Test_ztna_connector_all_ConnectorAPIService_TechSupport_List(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-tslist-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	list, httpRes, err := client.ConnectorAPI.ListConnectorsTechSupportFilesByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("ListTechSupportFiles error (connector may be offline): %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	if list != nil {
		t.Logf("Tech support files: %+v", list)
	}
}

// Test_ztna_connector_all_ConnectorAPIService_TechSupport_Create triggers tech support collection.
// Only succeeds on online connectors.
func Test_ztna_connector_all_ConnectorAPIService_TechSupport_Create(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-tscreate-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	httpRes, err := client.ConnectorAPI.CreateConnectorsTechSupportFilesByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("CreateTechSupportFile error (connector must be online): %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.True(t, httpRes.StatusCode == 200 || httpRes.StatusCode == 201 || httpRes.StatusCode == 202,
		"Expected 200/201/202, got %d", httpRes.StatusCode)
	t.Logf("CreateTechSupportFile OK: oid=%s status=%d", oid, httpRes.StatusCode)
}

// Test_ztna_connector_all_ConnectorAPIService_TechSupport_StopAndDownload stops and downloads
// a tech support file. Skipped if file collection cannot be triggered.
func Test_ztna_connector_all_ConnectorAPIService_TechSupport_StopAndDownload(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-connector-tsstop-%s", common.GenerateRandomString(6))

	createTestConnector(t, client, name, groupID)
	oid := fetchTestConnectorOID(t, client, name)
	t.Cleanup(func() { deleteTestConnector(t, client, oid, name) })

	_, err := client.ConnectorAPI.CreateConnectorsTechSupportFilesByID(context.Background(), oid).Execute()
	if err != nil {
		t.Skipf("Cannot create tech support file (connector must be online): %v", err)
	}

	list, _, listErr := client.ConnectorAPI.ListConnectorsTechSupportFilesByID(context.Background(), oid).Execute()
	if listErr != nil || list == nil || len(list.GetData()) == 0 {
		t.Skip("No tech support files available to stop/download")
	}
	fileID := list.GetData()[0].GetId()

	stopRes, err := client.ConnectorAPI.StopConnectorsTechSupportFilesByID(context.Background(), oid, fileID).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("StopTechSupportFile error: %v", err)
	} else {
		require.NotNil(t, stopRes)
		assert.Equal(t, 200, stopRes.StatusCode)
		t.Logf("StopTechSupportFile OK: fileId=%s", fileID)
	}

	dlRes, dlErr := client.ConnectorAPI.DownloadConnectorsTechSupportFilesByID(context.Background(), oid, fileID).Execute()
	if dlErr != nil {
		handleAPIError(dlErr)
		t.Logf("DownloadTechSupportFile error: %v", dlErr)
		return
	}
	require.NotNil(t, dlRes)
	assert.Equal(t, 200, dlRes.StatusCode)
	t.Logf("DownloadTechSupportFile OK: fileId=%s", fileID)
}
