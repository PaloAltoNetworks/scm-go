/*
 * ZTNA Wildcard Testing — WildcardAPIService
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

// wildcardFQDN derives a unique FQDN from the test name so each test owns a
// distinct FQDN — the API rejects duplicate FQDNs across all wildcards in the tenant.
func wildcardFQDN(name string) string {
	return fmt.Sprintf("*.%s.example.com", name)
}

// createTestWildcard creates a wildcard and asserts 201.
func createTestWildcard(t *testing.T, client *ztna_connector_all.APIClient, name, groupOID, fqdn string) {
	t.Helper()
	wc := ztna_connector_all.NewWildcards(fqdn, groupOID, name)
	httpRes, err := client.WildcardAPI.CreateWildcard(context.Background()).
		Wildcards(*wc).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test wildcard")
	require.NotNil(t, httpRes)
	require.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created")
	t.Logf("Created test wildcard: %s in group: %s", name, groupOID)
}

// fetchTestWildcardOID finds a wildcard by name via list and returns its OID.
// The wildcard list response populates either oid or id depending on the API version,
// so both fields are checked.
func fetchTestWildcardOID(t *testing.T, client *ztna_connector_all.APIClient, name string) string {
	t.Helper()
	listRes, httpRes, err := client.WildcardAPI.ListWildcards(context.Background()).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list wildcards")
	require.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, listRes)

	for _, w := range listRes.GetData() {
		if w.Name == name {
			oid := w.GetOid()
			if oid == "" && w.Id != nil {
				oid = *w.Id
			}
			require.NotEmpty(t, oid, "Wildcard %s should have an OID or ID", name)
			t.Logf("Found wildcard %s with OID: %s", name, oid)
			return oid
		}
	}
	t.Fatalf("Wildcard %q not found in list", name)
	return ""
}

// deleteTestWildcard deletes a wildcard by OID. Treats 404 as already deleted.
func deleteTestWildcard(t *testing.T, client *ztna_connector_all.APIClient, oid, name string) {
	t.Helper()
	httpRes, err := client.WildcardAPI.DeleteWildcardByID(context.Background(), oid).Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == 404 {
			t.Logf("Wildcard already deleted (OID: %s)", oid)
			return
		}
		handleAPIError(err)
		require.Fail(t, "Failed to delete test wildcard", "OID: %s", oid)
	}
	require.Equal(t, 202, httpRes.StatusCode, "Expected 202 for delete")
	t.Logf("Deleted test wildcard: %s (OID: %s)", name, oid)
}

// ── CRUD ─────────────────────────────────────────────────────────────────────

// Test_ztna_connector_all_WildcardAPIService_Create tests creating a wildcard.
func Test_ztna_connector_all_WildcardAPIService_Create(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-wc-create-%s", common.GenerateRandomString(6))

	createTestWildcard(t, client, name, groupID, wildcardFQDN(name))
	oid := fetchTestWildcardOID(t, client, name)
	assert.NotEmpty(t, oid, "OID should be set after creation")
	t.Logf("Successfully created wildcard: %s with OID: %s", name, oid)

	t.Cleanup(func() { deleteTestWildcard(t, client, oid, name) })
}

// Test_ztna_connector_all_WildcardAPIService_GetByID creates a wildcard and retrieves it by OID.
func Test_ztna_connector_all_WildcardAPIService_GetByID(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-wc-getbyid-%s", common.GenerateRandomString(6))

	createTestWildcard(t, client, name, groupID, wildcardFQDN(name))
	oid := fetchTestWildcardOID(t, client, name)
	t.Cleanup(func() { deleteTestWildcard(t, client, oid, name) })

	wc, httpRes, err := client.WildcardAPI.GetWildcardByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to get wildcard by ID")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	require.NotNil(t, wc)
	assert.Equal(t, name, wc.Name)
	assert.Equal(t, groupID, wc.Group)
	assert.Equal(t, wildcardFQDN(name), wc.Fqdn)
	gotOID := wc.GetOid()
	if gotOID == "" && wc.Id != nil {
		gotOID = *wc.Id
	}
	assert.Equal(t, oid, gotOID)
	t.Logf("GetByID OK: name=%s oid=%s", wc.Name, gotOID)
}

// Test_ztna_connector_all_WildcardAPIService_Update creates a wildcard and updates its description.
func Test_ztna_connector_all_WildcardAPIService_Update(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-wc-update-%s", common.GenerateRandomString(6))

	createTestWildcard(t, client, name, groupID, wildcardFQDN(name))
	oid := fetchTestWildcardOID(t, client, name)
	t.Cleanup(func() { deleteTestWildcard(t, client, oid, name) })

	updatedDesc := "Updated by test"
	updated := ztna_connector_all.NewWildcards(wildcardFQDN(name), groupID, name)
	updated.Description = &updatedDesc

	httpRes, err := client.WildcardAPI.UpdateWildcardByID(context.Background(), oid).
		Wildcards(*updated).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to update wildcard")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK for update")

	wc, _, err := client.WildcardAPI.GetWildcardByID(context.Background(), oid).Execute()
	require.NoError(t, err)
	require.NotNil(t, wc.Description)
	assert.Equal(t, updatedDesc, *wc.Description, "Description should be updated")
	t.Logf("Update OK: oid=%s", oid)
}

// Test_ztna_connector_all_WildcardAPIService_List lists wildcards and verifies the created one appears.
func Test_ztna_connector_all_WildcardAPIService_List(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-wc-list-%s", common.GenerateRandomString(6))

	createTestWildcard(t, client, name, groupID, wildcardFQDN(name))
	oid := fetchTestWildcardOID(t, client, name)
	t.Cleanup(func() { deleteTestWildcard(t, client, oid, name) })

	listRes, httpRes, err := client.WildcardAPI.ListWildcards(context.Background()).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list wildcards")
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, listRes)

	found := false
	for _, w := range listRes.GetData() {
		if w.Name == name {
			found = true
			gotOID := w.GetOid()
			if gotOID == "" && w.Id != nil {
				gotOID = *w.Id
			}
			assert.Equal(t, oid, gotOID)
			break
		}
	}
	assert.True(t, found, "Created wildcard should appear in list")
	t.Logf("List OK: total=%d found=%v", listRes.GetTotal(), found)
}

// Test_ztna_connector_all_WildcardAPIService_Delete creates a wildcard, deletes it, then verifies 404.
func Test_ztna_connector_all_WildcardAPIService_Delete(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-wc-delete-%s", common.GenerateRandomString(6))

	createTestWildcard(t, client, name, groupID, wildcardFQDN(name))
	oid := fetchTestWildcardOID(t, client, name)

	httpRes, err := client.WildcardAPI.DeleteWildcardByID(context.Background(), oid).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to delete wildcard")
	require.NotNil(t, httpRes)
	assert.Equal(t, 202, httpRes.StatusCode, "Expected 202 for delete")
	t.Logf("Delete OK: name=%s oid=%s", name, oid)

	_, httpRes, err = client.WildcardAPI.GetWildcardByID(context.Background(), oid).Execute()
	require.Error(t, err, "Expected error after deletion")
	require.NotNil(t, httpRes)
	assert.Equal(t, 404, httpRes.StatusCode, "Expected 404 after deletion")
	t.Logf("Verified 404 after delete: oid=%s", oid)
}

// ── Filters & Fetch ───────────────────────────────────────────────────────────

// Test_ztna_connector_all_WildcardAPIService_ListFilters lists filter values for the "name" field.
func Test_ztna_connector_all_WildcardAPIService_ListFilters(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)

	filters, httpRes, err := client.WildcardAPI.ListWildcardFilters(context.Background()).
		Field("name").Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list wildcard filters")
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	t.Logf("Wildcard name filters (%d): %v", len(filters), filters)
}

// Test_ztna_connector_all_WildcardAPIService_FetchByName tests the FetchWildcard convenience method.
func Test_ztna_connector_all_WildcardAPIService_FetchByName(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)
	groupID := provisionGroupID(t, client)
	name := fmt.Sprintf("test-wc-fetch-%s", common.GenerateRandomString(8))

	createTestWildcard(t, client, name, groupID, wildcardFQDN(name))
	oid := fetchTestWildcardOID(t, client, name)
	t.Cleanup(func() { deleteTestWildcard(t, client, oid, name) })

	wc, err := client.WildcardAPI.FetchWildcard(context.Background(), name, nil, nil, nil)
	require.NoError(t, err, "FetchWildcard should not error")
	require.NotNil(t, wc, "FetchWildcard should find the wildcard")
	assert.Equal(t, name, wc.Name)
	t.Logf("FetchWildcard found: %s with OID: %s", wc.Name, wc.GetOid())

	missing, err := client.WildcardAPI.FetchWildcard(context.Background(), "nonexistent-wildcard-xyz-12345", nil, nil, nil)
	require.NoError(t, err, "FetchWildcard should not error for missing wildcard")
	assert.Nil(t, missing, "Should return nil for non-existent wildcard")
}
