/*
Testing SharedSnippetsAPIService
Methods covered: List
Note: Shared snippets are read-only reference data (List only).
Convert and Load methods require specific snippet share payloads.
*/
package config_setup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_config_setup_SharedSnippetsAPIService_List tests listing shared snippets
func Test_config_setup_SharedSnippetsAPIService_List(t *testing.T) {
	client := SetupConfigSvcTestClient(t)

	reqList := client.SharedSnippetsAPI.ListSharedSnippets(context.Background())
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	require.NoError(t, errList, "Failed to list shared snippets")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed %d shared snippets", len(listRes))
}
