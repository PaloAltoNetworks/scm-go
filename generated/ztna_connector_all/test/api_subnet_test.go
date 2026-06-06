/*
 * ZTNA Subnet Testing — SubnetAPIService
 * Covers: CRUD, Filters, FetchByName
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

// subnetConnectorGroupID reads ZTNA_CONNECTOR_GROUP_ID and skips if unset.
func subnetConnectorGroupID(t *testing.T) string {
	t.Helper()
	return connectorGroupID(t)
}

const testIPSubnet = "10.100.0.0/24"

// createTestSubnet creates a subnet and asserts 201.
func createTestSubnet(t *testing.T, client *ztna_connector_all.APIClient, name, groupOID, ipSubnets string) {
	t.Helper()
	subnet := ztna_connector_all.NewSubnets(groupOID, ipSubnets, name)
	httpRes, err := client.SubnetAPI.CreateSubnet(context.Background()).
		Subnets(*subnet).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test subnet")
	require.NotNil(t, httpRes)
	require.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created")
	t.Logf("Created test subnet: %s in group: %s", name, groupOID)
}

// fetchTestSubnetOID finds a subnet by name via list and returns its OID.
func fetchTestSubnetOID(t *testing.T, client *ztna_connector_all.APIClient, name string) string {
	t.Helper()
	listRes, httpRes, err := client.SubnetAPI.ListSubnets(context.Background()).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list subnets")
	require.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, listRes)

	for _, s := range listRes.GetData() {
		if s.Name == name {
			oid := s.GetOid()
			require.NotEmpty(t, oid, "Subnet %s should have an OID", name)
			t.Logf("Found subnet %s with OID: %s", name, oid)
			return oid
		}
	}
	t.Fatalf("Subnet %q not found in list", name)
	return ""
}

// deleteTestSubnet deletes a subnet by OID. Treats 404 as already deleted.
func deleteTestSubnet(t *testing.T, client *ztna_connector_all.APIClient, oid, name string) {
	t.Helper()
	httpRes, err := client.SubnetAPI.DeleteSubnetByID(context.Background(), oid).Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == 404 {
			t.Logf("Subnet already deleted (OID: %s)", oid)
			return
		}
		handleAPIError(err)
		require.Fail(t, "Failed to delete test subnet", "OID: %s", oid)
	}
	require.Equal(t, 202, httpRes.StatusCode, "Expected 202 for delete")
	t.Logf("Deleted test subnet: %s (OID: %s)", name, oid)
}

// ── CRUD ─────────────────────────────────────────────────────────────────────

// Test_ztna_connector_all_SubnetAPIService_Create tests creating a subnet.
func Test_ztna_connector_all_SubnetAPIService_Create(t *testing.T) {
	groupID := subnetConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-subnet-create-%s", common.GenerateRandomString(6))

	createTestSubnet(t, client, name, groupID, testIPSubnet)
	oid := fetchTestSubnetOID(t, client, name)
	assert.NotEmpty(t, oid, "OID should be set after creation")
	t.Logf("Successfully created subnet: %s with OID: %s", name, oid)

	t.Cleanup(func() { deleteTestSubnet(t, client, oid, name) })
}

// Test_ztna_connector_all_SubnetAPIService_GetByID creates a subnet and retrieves it by OID.
func Test_ztna_connector_all_SubnetAPIService_GetByID(t *testing.T) {
	groupID := subnetConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-subnet-getbyid-%s", common.GenerateRandomString(6))

	createTestSubnet(t, client, name, groupID, testIPSubnet)
	oid := fetchTestSubnetOID(t, client, name)
	t.Cleanup(func() { deleteTestSubnet(t, client, oid, name) })

	subnet, httpRes, err := client.SubnetAPI.GetSubnetByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to get subnet by ID")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	require.NotNil(t, subnet)
	assert.Equal(t, name, subnet.Name)
	assert.Equal(t, groupID, subnet.Group)
	assert.Equal(t, testIPSubnet, subnet.IpSubnets)
	require.NotNil(t, subnet.Oid)
	assert.Equal(t, oid, *subnet.Oid)
	t.Logf("GetByID OK: name=%s oid=%s", subnet.Name, *subnet.Oid)
}

// Test_ztna_connector_all_SubnetAPIService_Update creates a subnet and updates its description.
func Test_ztna_connector_all_SubnetAPIService_Update(t *testing.T) {
	groupID := subnetConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-subnet-update-%s", common.GenerateRandomString(6))

	createTestSubnet(t, client, name, groupID, testIPSubnet)
	oid := fetchTestSubnetOID(t, client, name)
	t.Cleanup(func() { deleteTestSubnet(t, client, oid, name) })

	updatedDesc := "Updated by test"
	updated := ztna_connector_all.NewSubnets(groupID, testIPSubnet, name)
	updated.Description = &updatedDesc

	httpRes, err := client.SubnetAPI.UpdateSubnetByID(context.Background(), oid).
		Subnets(*updated).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to update subnet")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK for update")

	subnet, _, err := client.SubnetAPI.GetSubnetByID(context.Background(), oid).Execute()
	require.NoError(t, err)
	require.NotNil(t, subnet.Description)
	assert.Equal(t, updatedDesc, *subnet.Description, "Description should be updated")
	t.Logf("Update OK: oid=%s", oid)
}

// Test_ztna_connector_all_SubnetAPIService_List lists subnets and verifies the created one appears.
func Test_ztna_connector_all_SubnetAPIService_List(t *testing.T) {
	groupID := subnetConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-subnet-list-%s", common.GenerateRandomString(6))

	createTestSubnet(t, client, name, groupID, testIPSubnet)
	oid := fetchTestSubnetOID(t, client, name)
	t.Cleanup(func() { deleteTestSubnet(t, client, oid, name) })

	listRes, httpRes, err := client.SubnetAPI.ListSubnets(context.Background()).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list subnets")
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, listRes)

	found := false
	for _, s := range listRes.GetData() {
		if s.Name == name {
			found = true
			assert.Equal(t, oid, s.GetOid())
			break
		}
	}
	assert.True(t, found, "Created subnet should appear in list")
	t.Logf("List OK: total=%d found=%v", listRes.GetTotal(), found)
}

// Test_ztna_connector_all_SubnetAPIService_Delete creates a subnet, deletes it, then verifies 404.
func Test_ztna_connector_all_SubnetAPIService_Delete(t *testing.T) {
	groupID := subnetConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-subnet-delete-%s", common.GenerateRandomString(6))

	createTestSubnet(t, client, name, groupID, testIPSubnet)
	oid := fetchTestSubnetOID(t, client, name)

	httpRes, err := client.SubnetAPI.DeleteSubnetByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to delete subnet")
	require.NotNil(t, httpRes)
	assert.Equal(t, 202, httpRes.StatusCode, "Expected 202 for delete")
	t.Logf("Delete OK: name=%s oid=%s", name, oid)

	_, httpRes, err = client.SubnetAPI.GetSubnetByID(context.Background(), oid).Execute()
	require.Error(t, err, "Expected error after deletion")
	require.NotNil(t, httpRes)
	assert.Equal(t, 404, httpRes.StatusCode, "Expected 404 after deletion")
	t.Logf("Verified 404 after delete: oid=%s", oid)
}

// ── Filters & Fetch ───────────────────────────────────────────────────────────

// Test_ztna_connector_all_SubnetAPIService_ListFilters lists filter values for the "name" field.
func Test_ztna_connector_all_SubnetAPIService_ListFilters(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)

	filters, httpRes, err := client.SubnetAPI.ListSubnetFilters(context.Background()).
		Field("name").Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list subnet filters")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("Subnet name filters (%d): %v", len(filters), filters)
}

// Test_ztna_connector_all_SubnetAPIService_FetchByName tests the FetchSubnet convenience method.
func Test_ztna_connector_all_SubnetAPIService_FetchByName(t *testing.T) {
	groupID := subnetConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-subnet-fetch-%s", common.GenerateRandomString(8))

	createTestSubnet(t, client, name, groupID, testIPSubnet)
	oid := fetchTestSubnetOID(t, client, name)
	t.Cleanup(func() { deleteTestSubnet(t, client, oid, name) })

	subnet, err := client.SubnetAPI.FetchSubnet(context.Background(), name, nil, nil, nil)
	require.NoError(t, err, "FetchSubnet should not error")
	require.NotNil(t, subnet, "FetchSubnet should find the subnet")
	assert.Equal(t, name, subnet.Name)
	t.Logf("FetchSubnet found: %s with OID: %s", subnet.Name, subnet.GetOid())

	missing, err := client.SubnetAPI.FetchSubnet(context.Background(), "nonexistent-subnet-xyz-12345", nil, nil, nil)
	require.NoError(t, err, "FetchSubnet should not error for missing subnet")
	assert.Nil(t, missing, "Should return nil for non-existent subnet")
}
