/*
 * Security Services Testing
 *
 * DecryptionRulesAPIService
 */
package security_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/security_services"
)

// Helper function to create a minimal DecryptionRules object for testing.
func createTestDecryptionRule(t *testing.T, ruleName string) security_services.DecryptionRules {
	// Required fields for DecryptionRules:
	// Action, Category, Destination, From, Name, Service, Source, SourceUser, To
	//decryptionType := security_services.DecryptionRulesType{SslInboundInspection: common.StringPtr("test_cp_all_fields-123")}
	ruleType := security_services.DecryptionRulesType{
		// Initialize SslForwardProxy as an empty map
		SslForwardProxy: make(map[string]interface{}),
	}
	return security_services.DecryptionRules{
		Action:      "no-decrypt", // Common action for a decryption policy
		Category:    []string{"any"},
		Destination: []string{"any"},
		From:        []string{"any"},
		Name:        ruleName,
		Service:     []string{"service-https"},
		Source:      []string{"any"},
		SourceUser:  []string{"any"},
		To:          []string{"any"},
		Description: common.StringPtr("Test rule for Decryption CRUD"),
		Folder:      common.StringPtr("All"),
		LogSuccess:  common.BoolPtr(true),
		Type:        &ruleType,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_DecryptionRulesAPIService_Create tests the creation of a Decryption Rule.
func Test_security_services_DecryptionRulesAPIService_Create(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-decrypt-create-" + common.GenerateRandomString(6)

	rule := createTestDecryptionRule(t, ruleName)

	fmt.Printf("Creating Decryption Rule with name: %s\n", ruleName)
	req := client.DecryptionRulesAPI.CreateDecryptionRules(context.Background()).Position("pre").DecryptionRules(rule)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Decryption Rule")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status") // 201 for POST
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created rule should have an ID")
	createdRuleID := *res.Id

	// Cleanup the created rule
	defer func() {
		t.Logf("Cleaning up Decryption Rule with ID: %s", createdRuleID)
		_, errDel := client.DecryptionRulesAPI.DeleteDecryptionRulesByID(context.Background(), createdRuleID).Execute()
		require.NoError(t, errDel, "Failed to delete Decryption Rule during cleanup")
	}()

	t.Logf("Successfully created Decryption Rule: %s with ID: %s", ruleName, createdRuleID)
	// Verify the response matches the input
	assert.Equal(t, ruleName, res.Name, "Created rule name should match")
	assert.Equal(t, "no-decrypt", res.Action, "Action should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_DecryptionRulesAPIService_GetByID tests retrieving a Decryption Rule by ID.
func Test_security_services_DecryptionRulesAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-decrypt-get-" + common.GenerateRandomString(6)
	rule := createTestDecryptionRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.DecryptionRulesAPI.CreateDecryptionRules(context.Background()).Position("pre").DecryptionRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for get test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.DecryptionRulesAPI.DeleteDecryptionRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: Retrieve the rule
	getRes, httpResGet, errGet := client.DecryptionRulesAPI.GetDecryptionRulesByID(context.Background(), createdRuleID).Execute()

	require.NoError(t, errGet, "Failed to get Decryption Rule by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, ruleName, getRes.Name, "Rule name should match")
	assert.ElementsMatch(t, []string{"service-https"}, getRes.Service, "Rule service should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_DecryptionRulesAPIService_Update tests updating a Decryption Rule.
func Test_security_services_DecryptionRulesAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-decrypt-update-" + common.GenerateRandomString(6)
	rule := createTestDecryptionRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.DecryptionRulesAPI.CreateDecryptionRules(context.Background()).Position("pre").DecryptionRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for update test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.DecryptionRulesAPI.DeleteDecryptionRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Prepare updated rule object
	updatedDescription := "Updated decryption rule description and action"
	updatedRule := createTestDecryptionRule(t, ruleName)
	updatedRule.Action = "decrypt" // Change the action
	updatedRule.Description = common.StringPtr(updatedDescription)
	updatedRule.Id = &createdRuleID

	// Test: Update the rule
	updateRes, httpResUpdate, errUpdate := client.DecryptionRulesAPI.UpdateDecryptionRulesByID(context.Background(), createdRuleID).DecryptionRules(updatedRule).Execute()

	require.NoError(t, errUpdate, "Failed to update Decryption Rule")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the changes
	assert.Equal(t, "decrypt", updateRes.Action, "Action should be updated to decrypt")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_DecryptionRulesAPIService_List tests listing Decryption Rules.
func Test_security_services_DecryptionRulesAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-decrypt-list-" + common.GenerateRandomString(6)
	rule := createTestDecryptionRule(t, ruleName)

	// Setup: Create a rule first to ensure one exists for the list
	createRes, _, err := client.DecryptionRulesAPI.CreateDecryptionRules(context.Background()).Position("pre").DecryptionRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for list test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.DecryptionRulesAPI.DeleteDecryptionRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: List the rules, filtering by the specific name to avoid issues with default rules.
	listRes, httpResList, errList := client.DecryptionRulesAPI.ListDecryptionRules(context.Background()).
		Position("pre").
		Folder("All").
		Limit(50).
		Offset(10).
		Execute()

	require.NoError(t, errList, "Failed to list Decryption Rules")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_DecryptionRulesAPIService_DeleteByID tests deleting a Decryption Rule.
func Test_security_services_DecryptionRulesAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-decrypt-delete-" + common.GenerateRandomString(6)
	rule := createTestDecryptionRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.DecryptionRulesAPI.CreateDecryptionRules(context.Background()).Position("pre").DecryptionRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for delete test setup")
	createdRuleID := *createRes.Id

	// Test: Delete the rule
	httpResDel, errDel := client.DecryptionRulesAPI.DeleteDecryptionRulesByID(context.Background(), createdRuleID).Execute()

	require.NoError(t, errDel, "Failed to delete Decryption Rule")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_DecryptionRulesAPIService_Move tests moving a Decryption Rule.
func Test_security_services_DecryptionRulesAPIService_Move(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	// Create two rules (Rule A and Rule B)
	ruleAName := "test-decrypt-move-A-" + common.GenerateRandomString(4)
	ruleBName := "test-decrypt-move-B-" + common.GenerateRandomString(4)

	ruleA := createTestDecryptionRule(t, ruleAName)
	ruleB := createTestDecryptionRule(t, ruleBName)

	// 1. Create rule B (anchor)
	resB, _, errB := client.DecryptionRulesAPI.CreateDecryptionRules(context.Background()).
		Position("pre").DecryptionRules(ruleB).Execute()
	require.NoError(t, errB, "Failed to create rule B for move test setup")
	idB := *resB.Id

	// 2. Create rule A (target, initially above B)
	resA, _, errA := client.DecryptionRulesAPI.CreateDecryptionRules(context.Background()).
		Position("pre").DecryptionRules(ruleA).Execute()
	require.NoError(t, errA, "Failed to create rule A for move test setup")
	idA := *resA.Id

	// Defer cleanup for both rules
	defer func() {
		client.DecryptionRulesAPI.DeleteDecryptionRulesByID(context.Background(), idA).Execute()
		client.DecryptionRulesAPI.DeleteDecryptionRulesByID(context.Background(), idB).Execute()
	}()

	// Define the move payload: move Rule A *after* Rule B
	movePayload := security_services.RuleBasedMove{
		Destination:     "after",
		DestinationRule: &idB, // Anchor rule ID
	}

	// Test: Execute the move operation
	httpResMove, errMove := client.DecryptionRulesAPI.
		MoveDecryptionRulesByID(context.Background(), idA). // Rule to move
		RuleBasedMove(movePayload).
		Execute()

	require.NoError(t, errMove, "Failed to execute move operation for Decryption Rule")
	assert.Equal(t, 200, httpResMove.StatusCode, "Expected 200 OK status for move operation")

	t.Logf("Successfully executed move operation for rule %s (moved after %s)", idA, idB)
}
