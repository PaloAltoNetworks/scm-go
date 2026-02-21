/*
Testing ApplicationDefaultsAPIService
Methods covered: Create
Note: Application defaults Create is a configuration-setting operation.
*/
package deployment_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_deployment_services_ApplicationDefaultsAPIService_Create tests creating application defaults
func Test_deployment_services_ApplicationDefaultsAPIService_Create(t *testing.T) {
	t.Skip("Skipping - creating application defaults requires specific deployment configuration")

	client := SetupDeploymentSvcTestClient(t)

	httpRes, err := client.ApplicationDefaultsAPI.CreateApplicationDefaults(context.Background()).Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create application defaults")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	t.Logf("Successfully created application defaults")
}
