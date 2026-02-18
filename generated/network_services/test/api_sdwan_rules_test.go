/*
Network Services Testing SDWANRulesAPIService
*/
package network_services

import (
	"testing"

	"github.com/paloaltonetworks/scm-go/common"
)

// Note: SDWAN rules require traffic_distribution_profiles as a prerequisite,
// which in turn require sdwan_interface_profiles. The sdwan_interface_profiles
// API is not available in our OpenAPI spec, so CRUD tests cannot create the
// full prerequisite chain. List test works since it only reads existing data.

// Test_networkservices_SDWANRulesAPIService_Create tests creating an SD-WAN rule
func Test_networkservices_SDWANRulesAPIService_Create(t *testing.T) {
	t.Skip("Requires sdwan_interface_profiles prerequisite (for traffic distribution profiles) which is not available in our OpenAPI spec")
	_ = SetupNetworkSvcTestClient(t)
	_ = common.GenerateRandomString(6)
}

// Test_networkservices_SDWANRulesAPIService_List tests listing SD-WAN rules
func Test_networkservices_SDWANRulesAPIService_List(t *testing.T) {
	t.Skip("SDWAN rules List API returns 500 Internal Server Error - server-side issue, requires sdwan_interface_profiles setup")
	_ = SetupNetworkSvcTestClient(t)
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
	t.Skip("Requires sdwan_interface_profiles prerequisite (for traffic distribution profiles) which is not available in our OpenAPI spec")
	_ = SetupNetworkSvcTestClient(t)
	_ = common.GenerateRandomString(6)
}
