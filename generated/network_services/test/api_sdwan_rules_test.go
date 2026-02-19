/*
Network Services Testing SDWANRulesAPIService
*/
package network_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
)

// Note: SDWAN rules require traffic_distribution_profiles as a prerequisite,
// which in turn require sdwan_interface_profiles. The sdwan_interface_profiles
// API is not available in our OpenAPI spec, so CRUD tests cannot create the
// full prerequisite chain. List and Fetch tests work since they only read existing data.

// Test_networkservices_SDWANRulesAPIService_Create tests creating an SD-WAN rule
func Test_networkservices_SDWANRulesAPIService_Create(t *testing.T) {
	t.Skip("Requires sdwan_interface_profiles prerequisite (for traffic distribution profiles) which is not available in our OpenAPI spec")
	_ = SetupNetworkSvcTestClient(t)
	_ = common.GenerateRandomString(6)
}

// Test_networkservices_SDWANRulesAPIService_List tests listing SD-WAN rules
func Test_networkservices_SDWANRulesAPIService_List(t *testing.T) {
	t.Skip("API returns 500 Internal Server Error on list request")
	client := SetupNetworkSvcTestClient(t)

	// Read-only test: list existing objects (no Create needed)
	listRes, httpResList, errList := client.SDWANRulesAPI.ListSDWANRules(context.Background()).Folder("All").Limit(200).Offset(0).Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list SDWAN rules")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed SDWAN rules")
}

// Test_networkservices_SDWANRulesAPIService_GetByID tests retrieving an SD-WAN rule by ID
func Test_networkservices_SDWANRulesAPIService_GetByID(t *testing.T) {
	t.Skip("Requires sdwan_interface_profiles prerequisite (for traffic distribution profiles) which is not available in our OpenAPI spec")
	_ = SetupNetworkSvcTestClient(t)
	_ = common.GenerateRandomString(6)
}

// Test_networkservices_SDWANRulesAPIService_UpdateByID tests updating an SD-WAN rule
func Test_networkservices_SDWANRulesAPIService_UpdateByID(t *testing.T) {
	t.Skip("Requires sdwan_interface_profiles prerequisite (for traffic distribution profiles) which is not available in our OpenAPI spec")
	_ = SetupNetworkSvcTestClient(t)
	_ = common.GenerateRandomString(6)
}

// Test_networkservices_SDWANRulesAPIService_DeleteByID tests deleting an SD-WAN rule
func Test_networkservices_SDWANRulesAPIService_DeleteByID(t *testing.T) {
	t.Skip("Requires sdwan_interface_profiles prerequisite (for traffic distribution profiles) which is not available in our OpenAPI spec")
	_ = SetupNetworkSvcTestClient(t)
	_ = common.GenerateRandomString(6)
}

// Test_networkservices_SDWANRulesAPIService_FetchSDWANRules tests the Fetch convenience method
func Test_networkservices_SDWANRulesAPIService_FetchSDWANRules(t *testing.T) {
	t.Skip("API returns 500 Internal Server Error on list/fetch request")
	client := SetupNetworkSvcTestClient(t)

	// Read-only test: Fetch non-existent object (should return nil, nil)
	notFound, err := client.SDWANRulesAPI.FetchSDWANRules(
		context.Background(),
		"non-existent-sdwan-rule-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchSDWANRules correctly returned nil for non-existent object")
}
