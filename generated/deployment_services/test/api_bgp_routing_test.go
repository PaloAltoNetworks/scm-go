/*
Testing BGPRoutingAPIService
Methods covered: Get, Update
Note: BGP routing is singleton - only Get and Update are available.
*/
package deployment_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_deployment_services_BGPRoutingAPIService_Get tests getting BGP routing settings
func Test_deployment_services_BGPRoutingAPIService_Get(t *testing.T) {
	client := SetupDeploymentSvcTestClient(t)

	getRes, httpResGet, errGet := client.BGPRoutingAPI.GetBGPRouting(context.Background()).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	require.NoError(t, errGet, "Failed to get BGP routing settings")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	t.Logf("Successfully retrieved BGP routing settings")
}

// Test_deployment_services_BGPRoutingAPIService_Update tests updating BGP routing settings
// Uses Get to retrieve existing settings, then performs a no-op update
func Test_deployment_services_BGPRoutingAPIService_Update(t *testing.T) {
	client := SetupDeploymentSvcTestClient(t)

	// Get existing settings
	getRes, _, errGet := client.BGPRoutingAPI.GetBGPRouting(context.Background()).Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}
	require.NoError(t, errGet, "Failed to get BGP routing settings")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Perform no-op update with existing settings
	updateRes, httpResUpdate, errUpdate := client.BGPRoutingAPI.UpdateBGPRouting(context.Background()).BgpRouting(*getRes).Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	require.NoError(t, errUpdate, "Failed to update BGP routing settings")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	t.Logf("Successfully updated BGP routing settings (no-op)")
}
