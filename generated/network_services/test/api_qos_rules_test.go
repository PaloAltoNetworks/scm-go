/*
 * Network Deployment Testing
 *
 * QosRulesAPIService
 */

package network_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Helper function to create a minimal QosPolicyRules object for testing.
func createTestQosRule(t *testing.T, ruleName string) network_services.QosPolicyRules {
	// Required fields for QosPolicyRules: Action, Name

	// NOTE: Replace this placeholder with a correct, valid instantiation of QosPolicyRulesAction
	action := network_services.QosPolicyRulesAction{
		// ASSUMPTION: Setting Class defines the action. Using a common default class.
		Class: common.StringPtr("1"),
	}

	return network_services.QosPolicyRules{
		Name:        ruleName,
		Action:      action,
		Schedule:    common.StringPtr("Test"),
		Description: common.StringPtr("Test rule for QoS Policy CRUD"),
		Folder:      common.StringPtr("All"),
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_QoSRulesAPIService_Create tests the creation of a QoS Rule.
func Test_network_services_QoSRulesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-qos-create-" + common.GenerateRandomString(6)

	rule := createTestQosRule(t, ruleName)

	t.Logf("Creating QoS Rule with name: %s", ruleName)
	req := client.QoSRulesAPI.CreateQoSPolicyRules(context.Background()).Position("pre").QosPolicyRules(rule)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create QoS Rule")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created rule should have an ID")
	createdRuleID := *res.Id

	// Cleanup the created rule
	defer func() {
		t.Logf("Cleaning up QoS Rule with ID: %s", createdRuleID)
		_, errDel := client.QoSRulesAPI.DeleteQoSPolicyRulesByID(context.Background(), createdRuleID).Execute()
		require.NoError(t, errDel, "Failed to delete QoS Rule during cleanup")
	}()

	t.Logf("Successfully created QoS Rule: %s with ID: %s", ruleName, createdRuleID)

	// Verify the response matches key input fields
	assert.Equal(t, ruleName, res.Name, "Created rule name should match")
	//assert.Equal(t, "EF", res.GetDscpTos().GetCodepoint(), "DscpTos codepoint should match the expected placeholder value")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_QoSRulesAPIService_GetByID tests retrieving a QoS Rule by ID.
func Test_network_services_QoSRulesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-qos-get-" + common.GenerateRandomString(6)
	rule := createTestQosRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.QoSRulesAPI.CreateQoSPolicyRules(context.Background()).Position("pre").QosPolicyRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for get test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.QoSRulesAPI.DeleteQoSPolicyRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: Retrieve the rule
	getRes, httpResGet, errGet := client.QoSRulesAPI.GetQoSPolicyRulesByID(context.Background(), createdRuleID).Execute()

	require.NoError(t, errGet, "Failed to get QoS Rule by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, ruleName, getRes.Name, "Rule name should match")
	assert.Equal(t, "Test", getRes.GetSchedule(), "Schedule should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_QoSRulesAPIService_Update tests updating a QoS Rule.
func Test_network_services_QoSRulesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-qos-update-" + common.GenerateRandomString(6)
	rule := createTestQosRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.QoSRulesAPI.CreateQoSPolicyRules(context.Background()).Position("pre").QosPolicyRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for update test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.QoSRulesAPI.DeleteQoSPolicyRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Prepare updated rule object
	updatedDescription := "Updated QoS rule description"
	updatedRule := createTestQosRule(t, ruleName)
	updatedRule.Schedule = common.StringPtr("non-work-hours") // Change the schedule
	updatedRule.Description = common.StringPtr(updatedDescription)
	updatedRule.Id = &createdRuleID

	// Test: Update the rule
	updateRes, httpResUpdate, errUpdate := client.QoSRulesAPI.UpdateQoSPolicyRulesByID(context.Background(), createdRuleID).QosPolicyRules(updatedRule).Execute()

	require.NoError(t, errUpdate, "Failed to update QoS Rule")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the changes
	assert.Equal(t, "non-work-hours", *updateRes.Schedule, "Schedule should be updated")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_QoSRulesAPIService_List tests listing QoS Rules.
func Test_network_services_QoSRulesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-qos-list-" + common.GenerateRandomString(6)
	rule := createTestQosRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.QoSRulesAPI.CreateQoSPolicyRules(context.Background()).Position("pre").QosPolicyRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for list test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.QoSRulesAPI.DeleteQoSPolicyRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: List the rules, filtering by the specific name to avoid issues with default rules.
	listRes, httpResList, errList := client.QoSRulesAPI.ListQoSPolicyRules(context.Background()).
		Position("pre").
		Folder("All").
		Limit(50).
		Offset(10).
		Execute()
	fmt.Printf("Successfully retrieved %d QoS Rule(s).\n", len(listRes.Data))
	require.NoError(t, errList, "Failed to list QoS Rules")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_QoSRulesAPIService_DeleteByID tests deleting a QoS Rule.
func Test_network_services_QoSRulesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	ruleName := "test-qos-delete-" + common.GenerateRandomString(6)
	rule := createTestQosRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.QoSRulesAPI.CreateQoSPolicyRules(context.Background()).Position("pre").QosPolicyRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for delete test setup")
	createdRuleID := *createRes.Id

	// Test: Delete the rule
	httpResDel, errDel := client.QoSRulesAPI.DeleteQoSPolicyRulesByID(context.Background(), createdRuleID).Execute()

	require.NoError(t, errDel, "Failed to delete QoS Rule")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_network_services_QoSRulesAPIService_Move tests moving a QoS Rule.
func Test_network_services_QoSRulesAPIService_Move(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create two rules (Rule A and Rule B)
	ruleAName := "test-qos-move-A-" + common.GenerateRandomString(4)
	ruleBName := "test-qos-move-B-" + common.GenerateRandomString(4)

	ruleA := createTestQosRule(t, ruleAName)
	ruleB := createTestQosRule(t, ruleBName)

	// 1. Create rule B (anchor)
	resB, _, errB := client.QoSRulesAPI.CreateQoSPolicyRules(context.Background()).
		Position("pre").QosPolicyRules(ruleB).Execute()
	require.NoError(t, errB, "Failed to create rule B for move test setup")
	idB := *resB.Id

	// 2. Create rule A (target)
	resA, _, errA := client.QoSRulesAPI.CreateQoSPolicyRules(context.Background()).
		Position("pre").QosPolicyRules(ruleA).Execute()
	require.NoError(t, errA, "Failed to create rule A for move test setup")
	idA := *resA.Id

	// Defer cleanup for both rules
	defer func() {
		client.QoSRulesAPI.DeleteQoSPolicyRulesByID(context.Background(), idA).Execute()
		client.QoSRulesAPI.DeleteQoSPolicyRulesByID(context.Background(), idB).Execute()
	}()

	// Define the move payload: move Rule A *after* Rule B
	movePayload := network_services.RuleBasedMove{ // Assuming RuleBasedMove struct exists in this scope
		Destination:     "after",
		DestinationRule: &idB,
	}

	// Test: Execute the move operation
	httpResMove, errMove := client.QoSRulesAPI.
		MoveQoSPolicyRulesByID(context.Background(), idA).
		RuleBasedMove(movePayload).
		Execute()

	require.NoError(t, errMove, "Failed to execute move operation for QoS Rule")
	assert.Equal(t, 200, httpResMove.StatusCode, "Expected 200 OK status for move operation")

	t.Logf("Successfully executed move operation for rule %s (moved after %s)", idA, idB)
}
