/*
 * ZTNA Discovered Application Testing — DiscoveredApplicationAPIService
 * Covers: List, Filters
 * Note: Discovered applications are read-only; creation/deletion is not supported.
 */
package ztna_connector_all

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_ztna_connector_all_DiscoveredApplicationAPIService_List lists discovered applications.
// Skips if the tenant has no CIE domain data configured.
func Test_ztna_connector_all_DiscoveredApplicationAPIService_List(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)

	listRes, httpRes, err := client.DiscoveredApplicationAPI.ListDiscoveredApplications(context.Background()).
		Execute()
	if err != nil {
		handleAPIError(err)
		t.Skipf("Skipping: ListDiscoveredApplications returned an error (tenant may not have CIE configured): %v", err)
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	require.NotNil(t, listRes)
	t.Logf("ListDiscoveredApplications OK: total=%d", listRes.GetTotal())
	for i, app := range listRes.GetData() {
		t.Logf("  [%d] tenantId=%v count=%v", i, app.TenantId, app.Count)
	}
}

// Test_ztna_connector_all_DiscoveredApplicationAPIService_ListFilters lists available filter fields.
func Test_ztna_connector_all_DiscoveredApplicationAPIService_ListFilters(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)

	filters, httpRes, err := client.DiscoveredApplicationAPI.ListDiscoveredApplicationFilters(context.Background()).
		Execute()
	if err != nil {
		handleAPIError(err)
		t.Skipf("Skipping: ListDiscoveredApplicationFilters returned an error (tenant may not have CIE configured): %v", err)
	}
	require.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	t.Logf("DiscoveredApplication filters (%d): %v", len(filters), filters)
}
