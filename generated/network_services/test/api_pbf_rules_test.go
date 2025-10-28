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

type PbfRulesFrom struct {
	// A placeholder field. Real 'From' might include an interface name or zone name.
	Zone []string `json:"zone,omitempty"`
}

// Helper function to create a minimal PbfRules object for testing.
func createTestPbfRule(t *testing.T, ruleName string) network_services.PbfRules {

	// Using concrete, safe test values instead of "any" or default keywords
	// to prevent API validation errors like the previous NAT rules issue.
	validZone := "zone-trust"
	validAddress := "192.168.10.0/24"

	ruleAction := network_services.PbfRulesAction{
		// Initialize SslForwardProxy as an empty map
		Discard: make(map[string]interface{}),
	}

	return network_services.PbfRules{
		Name:        &ruleName,
		Description: common.StringPtr("Test PBF rule for CRUD"),
		Folder:      common.StringPtr("All"),

		// Required logic fields (assuming these fields are critical for validity)
		From:        &network_services.PbfRulesFrom{Zone: []string{validZone}},
		Source:      []string{validAddress},
		Destination: []string{validAddress},
		Application: []string{"web-browsing"},
		Service:     []string{"service-http"},
		Schedule:    common.StringPtr("non-work-hours"),
		Action:      &ruleAction,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_PBFRulesAPIService_Create tests the creation of a PBF Rule.
func Test_network_services_PBFRulesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-pbf-create-" + common.GenerateRandomString(6)
	rule := createTestPbfRule(t, ruleName)

	t.Logf("Creating PBF Rule with name: %s", ruleName)
	// CreatePBFRules does not require a 'position' query parameter in the API snippet
	req := client.PBFRulesAPI.CreatePBFRules(context.Background()).PbfRules(rule)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create PBF Rule")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created rule should have an ID")

	createdRuleID := *res.Id

	// Cleanup the created rule
	defer func() {
		t.Logf("Cleaning up PBF Rule with ID: %s", createdRuleID)
		_, errDel := client.PBFRulesAPI.DeletePBFRulesByID(context.Background(), createdRuleID).Execute()
		if errDel != nil {
			t.Logf("Cleanup failed: %v", errDel)
		}
	}()

	// Verify the response matches key input fields
	assert.Equal(t, ruleName, *res.Name, "Created rule name should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_PBFRulesAPIService_GetByID tests retrieving a PBF Rule by ID.
func Test_network_services_PBFRulesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-pbf-get-" + common.GenerateRandomString(6)
	rule := createTestPbfRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.PBFRulesAPI.CreatePBFRules(context.Background()).PbfRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for get test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.PBFRulesAPI.DeletePBFRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: Retrieve the rule
	getRes, httpResGet, errGet := client.PBFRulesAPI.GetPBFRulesByID(context.Background(), createdRuleID).Execute()

	require.NoError(t, errGet, "Failed to get PBF Rule by ID")
	assert.Equal(t, http.StatusOK, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, ruleName, *getRes.Name, "Rule name should match")
	assert.Equal(t, createdRuleID, *getRes.Id, "Rule ID should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_PBFRulesAPIService_UpdateByID tests updating a PBF Rule.
func Test_network_services_PBFRulesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-pbf-update-" + common.GenerateRandomString(6)
	rule := createTestPbfRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.PBFRulesAPI.CreatePBFRules(context.Background()).PbfRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for update test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.PBFRulesAPI.DeletePBFRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Prepare updated rule object
	updatedDescription := "Updated PBF rule description"
	updatedRuleName := ruleName + "-v2"

	// Use the retrieved object to ensure all required fields are present
	updatedRule := *createRes
	updatedRule.Id = nil
	updatedRule.Name = &updatedRuleName
	updatedRule.Description = common.StringPtr(updatedDescription)
	updatedRule.Source = []string{"10.1.1.1/32"} // Change a key field

	// Test: Update the rule
	updateRes, httpResUpdate, errUpdate := client.PBFRulesAPI.UpdatePBFRulesByID(context.Background(), createdRuleID).PbfRules(updatedRule).Execute()

	require.NoError(t, errUpdate, "Failed to update PBF Rule")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the changes
	assert.Equal(t, updatedRuleName, *updateRes.Name, "Name should be updated")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
	assert.ElementsMatch(t, []string{"10.1.1.1/32"}, updateRes.Source, "Source should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_PBFRulesAPIService_List tests listing PBF Rules.
func Test_network_services_PBFRulesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-pbf-list-" + common.GenerateRandomString(6)
	rule := createTestPbfRule(t, ruleName)

	// Setup: Create a rule first to ensure at least one result
	createRes, _, err := client.PBFRulesAPI.CreatePBFRules(context.Background()).PbfRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for list test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.PBFRulesAPI.DeletePBFRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: List the rules
	listReq := client.PBFRulesAPI.ListPBFRules(context.Background()).
		Limit(50).Folder("All").Offset(10)

	listRes, httpResList, errList := listReq.Execute()

	require.NoError(t, errList, "Failed to list PBF Rules")
	assert.Equal(t, http.StatusOK, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_PBFRulesAPIService_DeleteByID tests deleting a PBF Rule.
func Test_network_services_PBFRulesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-pbf-delete-" + common.GenerateRandomString(6)
	rule := createTestPbfRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.PBFRulesAPI.CreatePBFRules(context.Background()).PbfRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for delete test setup")
	createdRuleID := *createRes.Id

	// Test: Delete the rule
	httpResDel, errDel := client.PBFRulesAPI.DeletePBFRulesByID(context.Background(), createdRuleID).Execute()

	require.NoError(t, errDel, "Failed to delete PBF Rule")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status")

}
