/*
Testing NetworkLocationsAPIService
Methods covered: List
Note: Network locations are read-only reference data.
*/
package deployment_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_deployment_services_NetworkLocationsAPIService_List tests listing network locations
func Test_deployment_services_NetworkLocationsAPIService_List(t *testing.T) {
	client := SetupDeploymentSvcTestClient(t)

	reqList := client.NetworkLocationsAPI.ListLocations(context.Background())
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	require.NoError(t, errList, "Failed to list network locations")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d network locations", len(listRes))
}
