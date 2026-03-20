/*
Testing ApplicationDefaultsAPIService
Methods covered: Create
Note: Application defaults Create is a configuration-setting operation.
*/
package deployment_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_deployment_services_ApplicationDefaultsAPIService_Create tests creating/enabling application defaults
// This is an idempotent enable operation (POST /enable) - safe to call repeatedly
func Test_deployment_services_ApplicationDefaultsAPIService_Create(t *testing.T) {
	client := SetupDeploymentSvcTestClient(t)

	httpRes, err := client.ApplicationDefaultsAPI.CreateApplicationDefaults(context.Background()).Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create application defaults")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Contains(t, []int{200, 201}, httpRes.StatusCode, "Expected 200 or 201 status")
	t.Logf("Successfully created/enabled application defaults (status: %d)", httpRes.StatusCode)
}
