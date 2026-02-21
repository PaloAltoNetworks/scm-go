/*
Testing SubscribedTenantsAPIService
Methods covered: Create, List, Update, Delete
Note: Subscribed tenants require specific multi-tenant snippet sharing configuration.
All CRUD operations are skipped as they require cross-tenant trust setup.
List is tested with a known snippet ID if available.
*/
package config_setup

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_config_setup_SubscribedTenantsAPIService_Create tests creating a subscribed tenant
func Test_config_setup_SubscribedTenantsAPIService_Create(t *testing.T) {
	t.Skip("Skipping - creating subscribed tenants requires cross-tenant trust configuration")
}

// Test_config_setup_SubscribedTenantsAPIService_List tests listing subscribed tenants by snippet ID
func Test_config_setup_SubscribedTenantsAPIService_List(t *testing.T) {
	client := SetupConfigSvcTestClient(t)

	// Use Fetch to find a known snippet category (List fails due to model mismatch)
	fetchRes, errFetch := client.SnippetCategoriesAPI.FetchSnippetCategories(context.Background(), "app-tagging", nil, nil, nil)
	if errFetch != nil {
		handleAPIError(errFetch)
	}
	if fetchRes == nil {
		t.Skip("Predefined snippet category 'app-tagging' not found")
	}

	snippetID := fetchRes.Id
	t.Logf("Testing List subscribed tenants with snippet ID: %s", snippetID)

	tenants, httpRes, err := client.SubscribedTenantsAPI.ListSubscribedTenantsByID(context.Background(), snippetID).Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to list subscribed tenants")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	require.NotNil(t, tenants, "List response should not be nil")
	t.Logf("Successfully listed %d subscribed tenants for snippet %s", len(tenants), snippetID)
}

// Test_config_setup_SubscribedTenantsAPIService_Update tests updating a subscribed tenant
func Test_config_setup_SubscribedTenantsAPIService_Update(t *testing.T) {
	t.Skip("Skipping - updating subscribed tenants requires existing cross-tenant subscription")
}

// Test_config_setup_SubscribedTenantsAPIService_Delete tests deleting a subscribed tenant
func Test_config_setup_SubscribedTenantsAPIService_Delete(t *testing.T) {
	t.Skip("Skipping - deleting subscribed tenants requires existing cross-tenant subscription")
}
