/*
Testing TrustInformationAPIService
Methods covered: List
Note: Trust information is read-only reference data.
*/
package config_setup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_config_setup_TrustInformationAPIService_List tests listing trusted tenants with snippets
func Test_config_setup_TrustInformationAPIService_List(t *testing.T) {
	client := SetupConfigSvcTestClient(t)

	reqList := client.TrustInformationAPI.ListTrustedTenantsWithSnippets(context.Background()).Type_("subscriber")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	require.NoError(t, errList, "Failed to list trusted tenants with snippets")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d trusted tenants with snippets", len(listRes))
}
