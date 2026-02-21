/*
Testing SaasTenantRestrictionsAPIService
Methods covered: Get, Update
Uses snippet=office365 scope for all operations.
*/
package security_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_security_services_SaasTenantRestrictionsAPIService_Get tests getting SaaS tenant restrictions
func Test_security_services_SaasTenantRestrictionsAPIService_Get(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	getRes, httpResGet, errGet := client.SaasTenantRestrictionsAPI.GetSaasTenantRestrictions(context.Background()).Snippet("office365").Limit(200).Offset(0).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	require.NoError(t, errGet, "Failed to get SaaS tenant restrictions")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	t.Logf("Successfully retrieved SaaS tenant restrictions (total: %d)", getRes.Total)
}

// Test_security_services_SaasTenantRestrictionsAPIService_Update tests updating SaaS tenant restrictions
// Gets existing restriction via Get, then performs a no-op update with same data
func Test_security_services_SaasTenantRestrictionsAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	// Get existing restrictions with office365 snippet scope
	getRes, _, errGet := client.SaasTenantRestrictionsAPI.GetSaasTenantRestrictions(context.Background()).Snippet("office365").Limit(200).Offset(0).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}
	require.NoError(t, errGet, "Failed to get SaaS tenant restrictions")
	require.NotNil(t, getRes, "Get response should not be nil")

	if len(getRes.Data) == 0 {
		t.Skip("No SaaS tenant restrictions found in office365 snippet to test Update")
	}

	// Perform no-op update with existing restriction data
	existing := getRes.Data[0]
	t.Logf("Updating existing restriction: %v", *existing.Name)

	updateRes, httpResUpdate, errUpdate := client.SaasTenantRestrictionsAPI.UpdateSaasTenantRestrictions(context.Background()).Snippet("office365").SaasTenantRestrictions(existing).Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	require.NoError(t, errUpdate, "Failed to update SaaS tenant restrictions")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	t.Logf("Successfully updated SaaS tenant restrictions (no-op)")
}
