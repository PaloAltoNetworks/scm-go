/*
 * ZTNA Tenant API Testing — TenantAPIService
 * Covers: GetLicense, GetTenantStatus
 */
package ztna_connector_all

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_ztna_connector_all_TenantAPIService_GetLicense retrieves license information for the tenant.
// The endpoint returns a 200 with the license_info payload regardless of whether any connectors
// or applications are provisioned, so this test is safe to run against any live tenant.
func Test_ztna_connector_all_TenantAPIService_GetLicense(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)

	result, httpRes, err := client.TenantAPI.GetLicense(context.Background()).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "GetLicense should not return an error")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	require.NotNil(t, result, "License info result should not be nil")

	t.Logf("License info:")
	t.Logf("  license_name     = %v", result.GetLicenseName())
	t.Logf("  expiry           = %v", result.GetExpiry())
	t.Logf("  connectors       = %v", result.GetConnectors())
	t.Logf("  max_connectors   = %v", result.GetMaxConnectors())
	t.Logf("  applications     = %v", result.GetApplications())
	t.Logf("  max_applications = %v", result.GetMaxApplications())
}

// Test_ztna_connector_all_TenantAPIService_GetTenantStatus retrieves the current status of the tenant.
// Returns one of: ok, not_found, delete_in_progress
func Test_ztna_connector_all_TenantAPIService_GetTenantStatus(t *testing.T) {
	client := SetupZtnaConnectorAllTestClient(t)

	result, httpRes, err := client.TenantAPI.GetTenantStatus(context.Background()).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "GetTenantStatus should not return an error")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	require.NotNil(t, result, "Tenant status result should not be nil")

	t.Logf("Tenant status:")
	t.Logf("  status = %v", result.GetStatus())
	assert.NotEmpty(t, result.GetStatus(), "Tenant status should not be empty")
}
