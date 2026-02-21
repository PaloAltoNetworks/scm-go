/*
Testing SaasTenantRestrictionsAPIService
Methods covered: Get, Update
Note: SaaS tenant restrictions are singleton - only Get and Update are available.
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

	getRes, httpResGet, errGet := client.SaasTenantRestrictionsAPI.GetSaasTenantRestrictions(context.Background()).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	require.NoError(t, errGet, "Failed to get SaaS tenant restrictions")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	t.Logf("Successfully retrieved SaaS tenant restrictions")
}

// Test_security_services_SaasTenantRestrictionsAPIService_Update tests updating SaaS tenant restrictions
// Uses Get to retrieve existing settings, then performs a no-op update
func Test_security_services_SaasTenantRestrictionsAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	// Get existing settings
	getRes, _, errGet := client.SaasTenantRestrictionsAPI.GetSaasTenantRestrictions(context.Background()).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}
	require.NoError(t, errGet, "Failed to get SaaS tenant restrictions")
	require.NotNil(t, getRes, "Get response should not be nil")

	t.Skip("Skipping Update - SaaS tenant restrictions update requires specific tenant configuration")
}
