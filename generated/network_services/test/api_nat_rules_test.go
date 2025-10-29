package network_services

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Helper function to create a minimal NatRules object for testing.
func createTestNatRule(t *testing.T, ruleName string) network_services.NatRules {
	// Required fields for NatRules: Destination, From, Id, Name, Service, Source, To

	dns := network_services.NatRulesDestinationTranslationDnsRewrite{
		Direction: common.StringPtr("reverse"),
	}
	dynamicIpPort := network_services.NatRulesSourceTranslationDynamicIpAndPort{
		TranslatedAddress: []string{"10.1.1.20", "10.2.2.23"},
	}
	destinationTranslation := network_services.NatRulesDestinationTranslation{
		TranslatedAddress: common.StringPtr("10.1.1.10"),
		TranslatedPort:    common.Int32Ptr(443), // Use Int32Ptr for port
		DnsRewrite:        &dns,
	}
	sourceTranslation := network_services.NatRulesSourceTranslation{
		DynamicIpAndPort: &dynamicIpPort,
	}
	return network_services.NatRules{
		Name:        ruleName,
		Description: common.StringPtr("Test NAT rule for CRUD"),
		From:        []string{"any"},
		To:          []string{"untrust"},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     "service-https",
		Folder:      common.StringPtr("All"), // Common default folder
		// Optional fields
		NatType:                   common.StringPtr("ipv4"),
		DestinationTranslation:    &destinationTranslation,
		SourceTranslation:         &sourceTranslation,
		ActiveActiveDeviceBinding: common.StringPtr("1"),
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_NATRulesAPIService_Create tests the creation of a NAT Rule.
func Test_network_services_NATRulesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-nat-create-" + common.GenerateRandomString(6)
	rule := createTestNatRule(t, ruleName)

	t.Logf("Creating NAT Rule with name: %s", ruleName)
	// 'position' is a required query parameter
	req := client.NATRulesAPI.CreateNatRules(context.Background()).Position("pre").NatRules(rule)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create NAT Rule")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	createdRuleID := res.Id

	// Cleanup the created rule
	defer func() {
		t.Logf("Cleaning up NAT Rule with ID: %s", createdRuleID)
		_, errDel := client.NATRulesAPI.DeleteNatRulesByID(context.Background(), createdRuleID).Execute()
		if errDel != nil {
			t.Logf("Cleanup failed: %v", errDel)
		}
	}()

	// Verify the response matches key input fields
	assert.Equal(t, ruleName, res.Name, "Created rule name should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_NATRulesAPIService_GetByID tests retrieving a NAT Rule by ID.
func Test_network_services_NATRulesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-nat-get-" + common.GenerateRandomString(6)
	rule := createTestNatRule(t, ruleName)

	// Setup: Create a rule first (Position is required for creation)
	createRes, _, err := client.NATRulesAPI.CreateNatRules(context.Background()).Position("pre").NatRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for get test setup")
	createdRuleID := createRes.Id

	defer func() {
		client.NATRulesAPI.DeleteNatRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: Retrieve the rule
	getRes, httpResGet, errGet := client.NATRulesAPI.GetNatRulesByID(context.Background(), createdRuleID).Execute()

	require.NoError(t, errGet, "Failed to get NAT Rule by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, ruleName, getRes.Name, "Rule name should match")
	assert.Equal(t, createdRuleID, getRes.Id, "Rule ID should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_NATRulesAPIService_UpdateByID tests updating a NAT Rule.
func Test_network_services_NATRulesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-nat-update-" + common.GenerateRandomString(6)
	rule := createTestNatRule(t, ruleName)

	// Setup: Create a rule first (Position is required for creation)
	createRes, _, err := client.NATRulesAPI.CreateNatRules(context.Background()).Position("pre").NatRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for update test setup")
	createdRuleID := createRes.Id

	defer func() {
		client.NATRulesAPI.DeleteNatRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Prepare updated rule object
	updatedDescription := "Updated NAT rule description"
	updatedRuleName := ruleName + "-v2"

	// Create a copy of the original rule for the update payload
	updatedRule := *createRes
	updatedRule.Name = updatedRuleName
	updatedRule.Description = common.StringPtr(updatedDescription)
	updatedRule.Destination = []string{"10.0.0.0/8"} // Change a key field

	// Test: Update the rule
	// 'position' is a required query parameter for UpdateNatRulesByID
	updateRes, httpResUpdate, errUpdate := client.NATRulesAPI.UpdateNatRulesByID(context.Background(), createdRuleID).Position("pre").NatRules(updatedRule).Execute()

	require.NoError(t, errUpdate, "Failed to update NAT Rule")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the changes
	assert.Equal(t, updatedRuleName, updateRes.Name, "Name should be updated")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_NATRulesAPIService_List tests listing NAT Rules.
func Test_network_services_NATRulesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-nat-list-" + common.GenerateRandomString(6)
	rule := createTestNatRule(t, ruleName)

	// Setup: Create a rule first to ensure at least one result
	createRes, _, err := client.NATRulesAPI.CreateNatRules(context.Background()).Position("pre").NatRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for list test setup")
	createdRuleID := createRes.Id

	defer func() {
		client.NATRulesAPI.DeleteNatRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: List the rules
	// 'position' is a required query parameter
	listReq := client.NATRulesAPI.ListNatRules(context.Background()).
		Position("pre").
		Offset(10).
		Folder("All")

	listRes, httpResList, errList := listReq.Execute()

	require.NoError(t, errList, "Failed to list NAT Rules")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_NATRulesAPIService_DeleteByID tests deleting a NAT Rule.
func Test_network_services_NATRulesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-nat-delete-" + common.GenerateRandomString(6)
	rule := createTestNatRule(t, ruleName)

	// Setup: Create a rule first (Position is required for creation)
	createRes, _, err := client.NATRulesAPI.CreateNatRules(context.Background()).Position("pre").NatRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for delete test setup")
	createdRuleID := createRes.Id

	// Test: Delete the rule
	httpResDel, errDel := client.NATRulesAPI.DeleteNatRulesByID(context.Background(), createdRuleID).Execute()

	require.NoError(t, errDel, "Failed to delete NAT Rule")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status")
}
