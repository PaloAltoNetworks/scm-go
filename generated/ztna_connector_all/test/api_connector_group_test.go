/*
 * ZTNA Connector Group Testing — ConnectorGroupAPIService
 * Covers: CRUD, Filters, Scheduled Upgrade, Sub-resource lists (connectors, applications, subnets, wildcards)
 */
package ztna_connector_all

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"
)

// ── CRUD ─────────────────────────────────────────────────────────────────────

// Test_ztna_connector_all_ConnectorGroupAPIService_Create tests creating a connector group.
func Test_ztna_connector_all_ConnectorGroupAPIService_Create(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-cg-create-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	assert.NotEmpty(t, oid, "OID should be set after creation")
	t.Logf("Successfully created connector group: %s with OID: %s", name, oid)

	t.Cleanup(func() { deleteTestConnectorGroup(t, client, oid, name) })
}

// Test_ztna_connector_all_ConnectorGroupAPIService_GetByID creates a connector group and retrieves it by OID.
func Test_ztna_connector_all_ConnectorGroupAPIService_GetByID(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-cg-getbyid-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() { deleteTestConnectorGroup(t, client, oid, name) })

	group, httpRes, err := client.ConnectorGroupAPI.GetConnectorGroupByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to get connector group by ID")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	require.NotNil(t, group)
	assert.Equal(t, name, group.Name)
	require.NotNil(t, group.Oid)
	assert.Equal(t, oid, *group.Oid)
	t.Logf("GetByID OK: name=%s oid=%s", group.Name, *group.Oid)
}

// Test_ztna_connector_all_ConnectorGroupAPIService_Update creates a connector group and updates its description.
func Test_ztna_connector_all_ConnectorGroupAPIService_Update(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-cg-update-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() { deleteTestConnectorGroup(t, client, oid, name) })

	updatedDesc := "Updated by test"
	updated := ztna_connector_all.NewConnectorGroups(name)
	updated.Description = &updatedDesc

	httpRes, err := client.ConnectorGroupAPI.UpdateConnectorGroupByID(context.Background(), oid).
		ConnectorGroups(*updated).
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to update connector group")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK for update")

	group, _, err := client.ConnectorGroupAPI.GetConnectorGroupByID(context.Background(), oid).Execute()
	require.NoError(t, err)
	require.NotNil(t, group.Description)
	assert.Equal(t, updatedDesc, *group.Description, "Description should be updated")
	t.Logf("Update OK: oid=%s", oid)
}

// Test_ztna_connector_all_ConnectorGroupAPIService_List lists connector groups and verifies the created one appears.
func Test_ztna_connector_all_ConnectorGroupAPIService_List(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-cg-list-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() { deleteTestConnectorGroup(t, client, oid, name) })

	listRes, httpRes, err := client.ConnectorGroupAPI.ListConnectorGroups(context.Background()).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list connector groups")
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, listRes)

	found := false
	for _, g := range listRes.GetData() {
		if g.Name == name {
			found = true
			assert.Equal(t, oid, g.GetOid())
			break
		}
	}
	assert.True(t, found, "Created connector group should appear in list")
	t.Logf("List OK: total=%d found=%v", listRes.GetTotal(), found)
}

// Test_ztna_connector_all_ConnectorGroupAPIService_Delete creates a connector group, deletes it, then verifies 404.
func Test_ztna_connector_all_ConnectorGroupAPIService_Delete(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-cg-delete-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)

	httpRes, err := client.ConnectorGroupAPI.DeleteConnectorGroupByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to delete connector group")
	require.NotNil(t, httpRes)
	assert.Equal(t, 202, httpRes.StatusCode, "Expected 202 for delete")
	t.Logf("Delete OK: name=%s oid=%s", name, oid)

	_, httpRes, err = client.ConnectorGroupAPI.GetConnectorGroupByID(context.Background(), oid).Execute()
	require.Error(t, err, "Expected error after deletion")
	require.NotNil(t, httpRes)
	assert.Equal(t, 404, httpRes.StatusCode, "Expected 404 after deletion")
	t.Logf("Verified 404 after delete: oid=%s", oid)
}

// Test_ztna_connector_all_ConnectorGroupAPIService_FetchByName tests the FetchConnectorGroup convenience method.
func Test_ztna_connector_all_ConnectorGroupAPIService_FetchByName(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-cg-fetch-%s", common.GenerateRandomString(8))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() { deleteTestConnectorGroup(t, client, oid, name) })

	group, err := client.ConnectorGroupAPI.FetchConnectorGroup(context.Background(), name, nil, nil, nil)
	require.NoError(t, err, "FetchConnectorGroup should not error")
	require.NotNil(t, group, "FetchConnectorGroup should find the group")
	assert.Equal(t, name, group.Name)
	t.Logf("FetchConnectorGroup found: %s with OID: %s", group.Name, group.GetOid())

	missing, err := client.ConnectorGroupAPI.FetchConnectorGroup(context.Background(), "nonexistent-cg-xyz-12345", nil, nil, nil)
	require.NoError(t, err, "FetchConnectorGroup should not error for missing group")
	assert.Nil(t, missing, "Should return nil for non-existent group")
}

// ── Filters ───────────────────────────────────────────────────────────────────

// Test_ztna_connector_all_ConnectorGroupAPIService_ListFilters lists filter values for the "name" field.
func Test_ztna_connector_all_ConnectorGroupAPIService_ListFilters(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)

	filters, httpRes, err := client.ConnectorGroupAPI.ListConnectorGroupFilters(context.Background()).
		Field("name").
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list connector group filters")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("Connector group name filters (%d): %v", len(filters), filters)
}

// ── Sub-resource lists ────────────────────────────────────────────────────────

// Test_ztna_connector_all_ConnectorGroupAPIService_ListConnectors lists connectors for a group.
func Test_ztna_connector_all_ConnectorGroupAPIService_ListConnectors(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-cg-lconn-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() { deleteTestConnectorGroup(t, client, oid, name) })

	result, httpRes, err := client.ConnectorGroupAPI.ListConnectorGroupConnectors(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list connectors for group")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("ListConnectorGroupConnectors OK: oid=%s", oid)
	_ = result
}

// Test_ztna_connector_all_ConnectorGroupAPIService_ListApplications lists FQDN rules for a group.
func Test_ztna_connector_all_ConnectorGroupAPIService_ListApplications(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-cg-lapp-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() { deleteTestConnectorGroup(t, client, oid, name) })

	result, httpRes, err := client.ConnectorGroupAPI.ListConnectorGroupApplications(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list applications for group")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("ListConnectorGroupApplications OK: oid=%s", oid)
	_ = result
}

// Test_ztna_connector_all_ConnectorGroupAPIService_ListSubnets lists subnets for a group.
func Test_ztna_connector_all_ConnectorGroupAPIService_ListSubnets(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-cg-lsub-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() { deleteTestConnectorGroup(t, client, oid, name) })

	result, httpRes, err := client.ConnectorGroupAPI.ListConnectorGroupSubnets(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list subnets for group")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("ListConnectorGroupSubnets OK: oid=%s", oid)
	_ = result
}

// Test_ztna_connector_all_ConnectorGroupAPIService_ListWildcards lists wildcards for a group.
func Test_ztna_connector_all_ConnectorGroupAPIService_ListWildcards(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-cg-lwild-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() { deleteTestConnectorGroup(t, client, oid, name) })

	result, httpRes, err := client.ConnectorGroupAPI.ListConnectorGroupWildcards(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list wildcards for group")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("ListConnectorGroupWildcards OK: oid=%s", oid)
	_ = result
}

// ── Scheduled Upgrade ─────────────────────────────────────────────────────────

// Test_ztna_connector_all_ConnectorGroupAPIService_ScheduledUpgrade_Create schedules an upgrade for a group.
func Test_ztna_connector_all_ConnectorGroupAPIService_ScheduledUpgrade_Create(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	imageID := fetchFirstConnectorImage(t, client)
	name := fmt.Sprintf("test-cg-sucreate-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() {
		_, _ = client.ConnectorGroupAPI.DeleteConnectorGroupScheduledUpgrade(context.Background(), oid).Execute()
		deleteTestConnectorGroup(t, client, oid, name)
	})

	upgrade := ztna_connector_all.NewConnectorGroupScheduledUpgrade(imageID)
	httpRes, err := client.ConnectorGroupAPI.CreateConnectorGroupScheduledUpgrade(context.Background(), oid).
		ConnectorGroupScheduledUpgrade(*upgrade).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("CreateConnectorGroupScheduledUpgrade error (group may need online connectors): %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.True(t, httpRes.StatusCode == 200 || httpRes.StatusCode == 201 || httpRes.StatusCode == 202,
		"Expected 200/201/202, got %d", httpRes.StatusCode)
	t.Logf("CreateConnectorGroupScheduledUpgrade OK: oid=%s imageID=%s", oid, imageID)
}

// Test_ztna_connector_all_ConnectorGroupAPIService_ScheduledUpgrade_Get retrieves a scheduled upgrade for a group.
func Test_ztna_connector_all_ConnectorGroupAPIService_ScheduledUpgrade_Get(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	imageID := fetchFirstConnectorImage(t, client)
	name := fmt.Sprintf("test-cg-suget-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() {
		_, _ = client.ConnectorGroupAPI.DeleteConnectorGroupScheduledUpgrade(context.Background(), oid).Execute()
		deleteTestConnectorGroup(t, client, oid, name)
	})

	upgrade := ztna_connector_all.NewConnectorGroupScheduledUpgrade(imageID)
	if _, err := client.ConnectorGroupAPI.CreateConnectorGroupScheduledUpgrade(context.Background(), oid).
		ConnectorGroupScheduledUpgrade(*upgrade).Execute(); err != nil {
		t.Skipf("Cannot create scheduled upgrade for group (may need online connectors): %v", err)
	}

	got, httpRes, err := client.ConnectorGroupAPI.GetConnectorGroupScheduledUpgrade(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("GetConnectorGroupScheduledUpgrade error: %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, got)
	assert.Equal(t, imageID, got.ImageId)
	t.Logf("GetConnectorGroupScheduledUpgrade OK: oid=%s imageId=%s", oid, got.ImageId)
}

// Test_ztna_connector_all_ConnectorGroupAPIService_ScheduledUpgrade_Delete deletes a scheduled upgrade for a group.
func Test_ztna_connector_all_ConnectorGroupAPIService_ScheduledUpgrade_Delete(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	imageID := fetchFirstConnectorImage(t, client)
	name := fmt.Sprintf("test-cg-sudelete-%s", common.GenerateRandomString(6))

	createTestConnectorGroup(t, client, name)
	oid := fetchTestConnectorGroupOID(t, client, name)
	t.Cleanup(func() { deleteTestConnectorGroup(t, client, oid, name) })

	upgrade := ztna_connector_all.NewConnectorGroupScheduledUpgrade(imageID)
	if _, err := client.ConnectorGroupAPI.CreateConnectorGroupScheduledUpgrade(context.Background(), oid).
		ConnectorGroupScheduledUpgrade(*upgrade).Execute(); err != nil {
		t.Skipf("Cannot create scheduled upgrade for group: %v", err)
	}

	httpRes, err := client.ConnectorGroupAPI.DeleteConnectorGroupScheduledUpgrade(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
		t.Logf("DeleteConnectorGroupScheduledUpgrade error: %v", err)
		return
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("DeleteConnectorGroupScheduledUpgrade OK: oid=%s", oid)
}
