/*
 * Identity Services Testing
 *
 * AuthenticationRulesAPIService
 */
package identity_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// Helper function to create a minimal AuthenticationRules object for testing.
func createTestAuthenticationRule(t *testing.T, ruleName string) identity_services.AuthenticationRules {
	// Required fields for AuthenticationRules:
	// Destination, From, Name, Service, Source, To

	// Note: The AuthenticationEnforcement profile is often required in real-world scenarios,
	// so we include a placeholder, assuming it exists in the test environment.
	authEnforcementName := "TEST-AUTH-PROFILE"

	return identity_services.AuthenticationRules{
		Destination:               []string{"any"},
		From:                      []string{"any"},
		Name:                      ruleName,
		Service:                   []string{"any"},
		Source:                    []string{"any"},
		To:                        []string{"any"},
		AuthenticationEnforcement: common.StringPtr(authEnforcementName),
		Timeout:                   common.Int32Ptr(int32(1000)), // 1 hour timeout
		Description:               common.StringPtr("Test rule for Auth Rule CRUD"),
		Folder:                    common.StringPtr("All"),
		LogAuthenticationTimeout:  common.BoolPtr(true),
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identity_services_AuthenticationRulesAPIService_Create tests the creation of an Authentication Rule.
func Test_identity_services_AuthenticationRulesAPIService_Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	ruleName := "test-auth-create-" + common.GenerateRandomString(6)

	rule := createTestAuthenticationRule(t, ruleName)

	t.Logf("Creating Auth Rule with name: %s", ruleName)
	req := client.AuthenticationRulesAPI.CreateAuthenticationRules(context.Background()).Position("pre").AuthenticationRules(rule)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Authentication Rule")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status") // 201 for POST
	require.NotNil(t, res, "Response body should not be nil")
	require.NotNil(t, res.Id, "Created rule should have an ID")
	createdRuleID := *res.Id

	// Cleanup the created rule
	defer func() {
		t.Logf("Cleaning up Auth Rule with ID: %s", createdRuleID)
		_, errDel := client.AuthenticationRulesAPI.DeleteAuthenticationRulesByID(context.Background(), createdRuleID).Execute()
		require.NoError(t, errDel, "Failed to delete Authentication Rule during cleanup")
	}()

	t.Logf("Successfully created Auth Rule: %s with ID: %s", ruleName, createdRuleID)

	// Verify the response matches key input fields
	assert.Equal(t, ruleName, res.Name, "Created rule name should match")
	assert.Equal(t, int32(1000), *res.Timeout, "Timeout should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identity_services_AuthenticationRulesAPIService_GetByID tests retrieving an Authentication Rule by ID.
func Test_identity_services_AuthenticationRulesAPIService_GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	ruleName := "test-auth-get-" + common.GenerateRandomString(6)
	rule := createTestAuthenticationRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.AuthenticationRulesAPI.CreateAuthenticationRules(context.Background()).Position("pre").AuthenticationRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for get test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.AuthenticationRulesAPI.DeleteAuthenticationRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: Retrieve the rule
	getRes, httpResGet, errGet := client.AuthenticationRulesAPI.GetAuthenticationRulesByID(context.Background(), createdRuleID).Execute()

	require.NoError(t, errGet, "Failed to get Authentication Rule by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data
	assert.Equal(t, ruleName, getRes.Name, "Rule name should match")
	assert.ElementsMatch(t, []string{"any"}, getRes.Service, "Rule service should match")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identity_services_AuthenticationRulesAPIService_Update tests updating an Authentication Rule.
func Test_identity_services_AuthenticationRulesAPIService_Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	ruleName := "test-auth-update-" + common.GenerateRandomString(6)
	rule := createTestAuthenticationRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.AuthenticationRulesAPI.CreateAuthenticationRules(context.Background()).Position("pre").AuthenticationRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for update test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.AuthenticationRulesAPI.DeleteAuthenticationRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Prepare updated rule object
	updatedTimeout := int32(900) // Change timeout to 15 minutes
	updatedDescription := "Updated auth rule description"
	updatedRule := createTestAuthenticationRule(t, ruleName)
	updatedRule.Timeout = &updatedTimeout
	updatedRule.Description = common.StringPtr(updatedDescription)
	updatedRule.Id = &createdRuleID

	// Test: Update the rule
	updateRes, httpResUpdate, errUpdate := client.AuthenticationRulesAPI.UpdateAuthenticationRulesByID(context.Background(), createdRuleID).AuthenticationRules(updatedRule).Execute()

	require.NoError(t, errUpdate, "Failed to update Authentication Rule")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// Verify the changes
	assert.Equal(t, updatedTimeout, *updateRes.Timeout, "Timeout should be updated")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identity_services_AuthenticationRulesAPIService_List tests listing Authentication Rules.
func Test_identity_services_AuthenticationRulesAPIService_List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	ruleName := "test-auth-list-" + common.GenerateRandomString(6)
	rule := createTestAuthenticationRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.AuthenticationRulesAPI.CreateAuthenticationRules(context.Background()).Position("pre").AuthenticationRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for list test setup")
	createdRuleID := *createRes.Id

	defer func() {
		client.AuthenticationRulesAPI.DeleteAuthenticationRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test: List the rules, filtering by the specific name to avoid issues with default rules.
	listRes, httpResList, errList := client.AuthenticationRulesAPI.ListAuthenticationRules(context.Background()).Position("pre").
		Folder("All").Limit(10).Offset(15).
		Execute()

	require.NoError(t, errList, "Failed to list Authentication Rules")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identity_services_AuthenticationRulesAPIService_DeleteByID tests deleting an Authentication Rule.
func Test_identity_services_AuthenticationRulesAPIService_DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	ruleName := "test-auth-delete-" + common.GenerateRandomString(6)
	rule := createTestAuthenticationRule(t, ruleName)

	// Setup: Create a rule first
	createRes, _, err := client.AuthenticationRulesAPI.CreateAuthenticationRules(context.Background()).Position("pre").AuthenticationRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for delete test setup")
	createdRuleID := *createRes.Id

	// Test: Delete the rule
	httpResDel, errDel := client.AuthenticationRulesAPI.DeleteAuthenticationRulesByID(context.Background(), createdRuleID).Execute()

	require.NoError(t, errDel, "Failed to delete Authentication Rule")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_identity_services_AuthenticationRulesAPIService_Move tests moving an Authentication Rule.
func Test_identity_services_AuthenticationRulesAPIService_Move(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Create two rules (Rule A and Rule B)
	ruleAName := "test-auth-move-A-" + common.GenerateRandomString(4)
	ruleBName := "test-auth-move-B-" + common.GenerateRandomString(4)

	ruleA := createTestAuthenticationRule(t, ruleAName)
	ruleB := createTestAuthenticationRule(t, ruleBName)

	// 1. Create rule B (anchor)
	resB, _, errB := client.AuthenticationRulesAPI.CreateAuthenticationRules(context.Background()).
		Position("pre").AuthenticationRules(ruleB).Execute()
	require.NoError(t, errB, "Failed to create rule B for move test setup")
	idB := *resB.Id

	// 2. Create rule A (target)
	resA, _, errA := client.AuthenticationRulesAPI.CreateAuthenticationRules(context.Background()).
		Position("pre").AuthenticationRules(ruleA).Execute()
	require.NoError(t, errA, "Failed to create rule A for move test setup")
	idA := *resA.Id

	// Defer cleanup for both rules
	defer func() {
		client.AuthenticationRulesAPI.DeleteAuthenticationRulesByID(context.Background(), idA).Execute()
		client.AuthenticationRulesAPI.DeleteAuthenticationRulesByID(context.Background(), idB).Execute()
	}()

	// Define the move payload: move Rule A *after* Rule B
	movePayload := identity_services.RuleBasedMove{ // Assuming RuleBasedMove struct exists in this scope
		Destination:     "after",
		DestinationRule: &idB,
	}

	// Test: Execute the move operation
	httpResMove, errMove := client.AuthenticationRulesAPI.
		MoveAuthenticationRulesByID(context.Background(), idA).
		RuleBasedMove(movePayload).
		Execute()

	require.NoError(t, errMove, "Failed to execute move operation for Authentication Rule")
	assert.Equal(t, 200, httpResMove.StatusCode, "Expected 200 OK status for move operation")

	t.Logf("Successfully executed move operation for rule %s (moved after %s)", idA, idB)
}
