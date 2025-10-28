/*
 * Security Services Testing
 *
 * AppOverrirdeRulesAPIService
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

// Helper function to create a minimal AppOverrideRules object for testing.
func createTestAppOverrideRule(t *testing.T, ruleName string) security_services.AppOverrideRules {
	// Required fields for AppOverrideRules:
	// Application, Destination, From, Name, Port, Protocol, Source, To
	return security_services.AppOverrideRules{
		Application:       "web-browsing",
		Destination:       []string{"any"},
		From:              []string{"any"},
		Name:              ruleName,
		Port:              "21", // Standard FTP port
		Protocol:          "tcp",
		Source:            []string{"any"},
		To:                []string{"any"},
		Description:       common.StringPtr("Test rule for AppOverride CRUD"),
		Folder:            common.StringPtr("All"),
		NegateDestination: common.BoolPtr(false),
		NegateSource:      common.BoolPtr(false),
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_ApplicationOverrideRulesAPIService_Create tests the creation of an App Override Rule.
func Test_security_services_ApplicationOverrideRulesAPIService_Create(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-app-override-create-" + common.GenerateRandomString(6)

	rule := createTestAppOverrideRule(t, ruleName)

	fmt.Printf("Creating App Override Rule with name: %s\n", ruleName)
	req := client.ApplicationOverrideRulesAPI.CreateApplicationOverrideRules(context.Background()).Position("pre").AppOverrideRules(rule)
	res, httpRes, err := req.Execute()

	if httpRes != nil {
		// Prints the HTTP status line (e.g., "201 Created")
		fmt.Printf("HTTP Status: %s\n", httpRes.Status)

		// Prints the numerical status code (e.g., 201)
		fmt.Printf("HTTP Status Code: %d\n", httpRes.StatusCode)

		// Prints a specific header (e.g., Content-Type)
		fmt.Printf("Content-Type Header: %s\n", httpRes.Header.Get("Content-Type"))
	}

	// Check for API errors and expected status code
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create App Override Rule")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status") // 201 for POST
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created rule should have an ID")
	createdRuleID := *res.Id

	// Cleanup the created rule
	defer func() {
		t.Logf("Cleaning up App Override Rule with ID: %s", createdRuleID)
		_, errDel := client.ApplicationOverrideRulesAPI.DeleteApplicationOverrideRulesByID(context.Background(), createdRuleID).Execute()
		require.NoError(t, errDel, "Failed to delete App Override Rule during cleanup")
	}()

	t.Logf("Successfully created App Override Rule: %s with ID: %s", ruleName, createdRuleID)
	// Verify the response matches the input
	assert.Equal(t, ruleName, res.Name, "Created rule name should match")
	assert.Equal(t, rule.Application, res.Application, "Application should match")
}

// ---------------------------------------------------------------------------------------------------------------------
// Test_security_services_ApplicationOverrideRulesAPIService_GetByID tests retrieving an App Override Rule by ID.
func Test_security_services_ApplicationOverrideRulesAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-app-override-get-" + common.GenerateRandomString(6)
	rule := createTestAppOverrideRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.ApplicationOverrideRulesAPI.CreateApplicationOverrideRules(context.Background()).Position("pre").AppOverrideRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for get test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.ApplicationOverrideRulesAPI.DeleteApplicationOverrideRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: Retrieve the rule
	getRes, httpResGet, errGet := client.ApplicationOverrideRulesAPI.GetApplicationOverrideRulesByID(context.Background(), createdRuleID).Execute()

	// Check for API errors and expected status code
	if errGet != nil {
		handleAPIError(errGet)
	}
	require.NoError(t, errGet, "Failed to get App Override Rule by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, ruleName, getRes.Name, "Rule name should match")
	assert.Equal(t, rule.Port, getRes.Port, "Rule port should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_ApplicationOverrideRulesAPIService_Update tests updating an App Override Rule.
func Test_security_services_ApplicationOverrideRulesAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-app-override-update-" + common.GenerateRandomString(6)
	rule := createTestAppOverrideRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.ApplicationOverrideRulesAPI.CreateApplicationOverrideRules(context.Background()).Position("pre").AppOverrideRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for update test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.ApplicationOverrideRulesAPI.DeleteApplicationOverrideRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Prepare updated rule object
	updatedPort := "8080"
	updatedDescription := "Updated app override rule description"

	updatedRule := createTestAppOverrideRule(t, ruleName)
	updatedRule.Port = updatedPort // Change a value
	updatedRule.Description = common.StringPtr(updatedDescription)
	updatedRule.Id = &createdRuleID // Include the ID for consistency

	// Test: Update the rule
	updateRes, httpResUpdate, errUpdate := client.ApplicationOverrideRulesAPI.UpdateApplicationOverrideRulesByID(context.Background(), createdRuleID).AppOverrideRules(updatedRule).Execute()

	// Check for API errors and expected status code
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}
	require.NoError(t, errUpdate, "Failed to update App Override Rule")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the changes
	assert.Equal(t, updatedPort, updateRes.Port, "Port should be updated")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_ApplicationOverrideRulesAPIService_List tests listing App Override Rules.
func Test_security_services_ApplicationOverrideRulesAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-app-override-list-" + common.GenerateRandomString(6)
	rule := createTestAppOverrideRule(t, ruleName)

	// Setup: Create a rule first to ensure one exists for the list
	createRes, _, err := client.ApplicationOverrideRulesAPI.CreateApplicationOverrideRules(context.Background()).Position("pre").AppOverrideRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for list test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.ApplicationOverrideRulesAPI.DeleteApplicationOverrideRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: List the rules
	listRes, httpResList, errList := client.ApplicationOverrideRulesAPI.ListApplicationOverrideRules(context.Background()).Position("pre").Folder("All").Limit(10).Offset(10).Execute()

	// Check for API errors and expected status code
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list App Override Rules")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.True(t, listRes.GetTotal() > 0, "Expected at least one rule in the list")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_ApplicationOverrideRulesAPIService_DeleteByID tests deleting an App Override Rule.
func Test_security_services_ApplicationOverrideRulesAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-app-override-delete-" + common.GenerateRandomString(6)
	rule := createTestAppOverrideRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.ApplicationOverrideRulesAPI.CreateApplicationOverrideRules(context.Background()).Position("pre").AppOverrideRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for delete test setup")
	createdRuleID := *createRes.Id

	// Test: Delete the rule
	httpResDel, errDel := client.ApplicationOverrideRulesAPI.DeleteApplicationOverrideRulesByID(context.Background(), createdRuleID).Execute()

	// Check for API errors and expected status code
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete App Override Rule")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_security_services_ApplicationOverrideRulesAPIService_Move tests moving an App Override Rule.
func Test_security_services_ApplicationOverrideRulesAPIService_Move(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	// Create two rules (Rule A and Rule B) to test the move operation
	ruleAName := "test-move-A-" + common.GenerateRandomString(4)
	ruleBName := "test-move-B-" + common.GenerateRandomString(4)

	ruleA := createTestAppOverrideRule(t, ruleAName)
	ruleB := createTestAppOverrideRule(t, ruleBName)

	// 1. Create rule B (will be the anchor)
	resB, _, errB := client.ApplicationOverrideRulesAPI.CreateApplicationOverrideRules(context.Background()).Position("pre").AppOverrideRules(ruleB).Execute()
	require.NoError(t, errB, "Failed to create rule B for move test setup")
	idB := *resB.Id

	// 2. Create rule A (will be moved)
	resA, _, errA := client.ApplicationOverrideRulesAPI.CreateApplicationOverrideRules(context.Background()).Position("pre").AppOverrideRules(ruleA).Execute()
	require.NoError(t, errA, "Failed to create rule A for move test setup")
	idA := *resA.Id

	// Defer cleanup for both rules
	defer func() {
		client.ApplicationOverrideRulesAPI.DeleteApplicationOverrideRulesByID(context.Background(), idA).Execute()
		client.ApplicationOverrideRulesAPI.DeleteApplicationOverrideRulesByID(context.Background(), idB).Execute()
	}()

	// Define the move payload: move rule A after rule B
	movePayload := security_services.RuleBasedMove{
		Destination:     "after",
		DestinationRule: &idB,
	}

	// Test: Execute the move operation
	httpResMove, errMove := client.ApplicationOverrideRulesAPI.MoveApplicationOverrideRulesByID(context.Background(), idA).RuleBasedMove(movePayload).Execute()

	// Check for API errors and expected status code
	if errMove != nil {
		handleAPIError(errMove)
	}
	require.NoError(t, errMove, "Failed to execute move operation for App Override Rule")
	assert.Equal(t, 200, httpResMove.StatusCode, "Expected 200 OK status for move operation")

	t.Logf("Successfully executed move operation for rule %s (moved after %s)", idA, idB)
}
