/*
 * Security Services Testing
 *
 * SecurityRulesAPIService
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

// NOTE: This file assumes the existence of the following helper functions:
// - SetupSecuritySvcTestClient(t *testing.T) *security_services.APIClient
// - handleAPIError(err error)
// These are typically located in a shared test utility file.

// Test_security_services_SecurityRulesAPIService_CreateSecurity tests the creation of a standard Security Rule.
func Test_security_services_SecurityRulesAPIService_CreateSecurity(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-security-rule-" + common.GenerateRandomString(6)

	// Use the provided working request payload
	securityRule := security_services.SecurityRules{
		Folder:      common.StringPtr("Shared"),
		Name:        common.StringPtr(ruleName),
		PolicyType:  common.StringPtr("Security"),
		From:        []string{"any"},
		To:          []string{"any"},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Application: []string{"any"},
		Action:      common.StringPtr("allow"),
		SourceUser:  []string{"casidp\\test_1", "CN=ng1,DC=casidp,DC=onmicrosoft,DC=com"},
		Category:    []string{"any"},
	}

	fmt.Printf("Creating Security Rule with name: %s\n", ruleName)
	req := client.SecurityRulesAPI.CreateSecurityRules(context.Background()).Position("pre").SecurityRules(securityRule)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Security Rule")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 200 OK status") // Note: Create returns 200 OK for rules
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, ruleName, *res.Name, "Created rule name should match")
	createdRuleID := *res.Id

	defer func() {
		t.Logf("Cleaning up Security Rule with ID: %s", createdRuleID)
		_, errDel := client.SecurityRulesAPI.DeleteSecurityRulesByID(context.Background(), createdRuleID).Execute()
		require.NoError(t, errDel, "Failed to delete Security Rule during cleanup")
	}()

	t.Logf("Successfully created Security Rule: %s with ID: %s", ruleName, createdRuleID)
	// Verify the response matches the expected output
	assert.Equal(t, "Security", *res.PolicyType, "PolicyType should match")
	assert.ElementsMatch(t, securityRule.SourceUser, res.SourceUser, "SourceUser should match")
}

// Test_security_services_SecurityRulesAPIService_CreateInternet tests the creation of an Internet Security Rule.
func Test_security_services_SecurityRulesAPIService_CreateInternet(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-internet-rule-" + common.GenerateRandomString(6)

	// Use the provided working request payload
	internetRule := security_services.SecurityRules{
		Folder:     common.StringPtr("Shared"),
		Name:       common.StringPtr(ruleName),
		PolicyType: common.StringPtr("Internet"),
		AllowWebApplication: []security_services.InternetRuleTypeAllowWebApplicationInner{
			{
				Name:                common.StringPtr("4sync"),
				Type:                common.StringPtr("application"),
				ApplicationFunction: []string{"any"},
			},
		},
		AllowUrlCategory: []security_services.InternetRuleTypeAllowUrlCategoryInner{
			{
				Name:                  common.StringPtr("AI-code-assistant"),
				Decryption:            common.StringPtr("enabled"),
				CredentialEnforcement: common.StringPtr("enabled"),
			},
		},
		From:        []string{"any"},
		To:          []string{"any"},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Action:      common.StringPtr("allow"),
		SourceUser:  []string{"casidp\\test_1", "CN=ng1,DC=casidp,DC=onmicrosoft,DC=com"},
	}

	fmt.Printf("Creating Internet Rule with name: %s\n", ruleName)
	req := client.SecurityRulesAPI.CreateSecurityRules(context.Background()).Position("pre").SecurityRules(internetRule)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Internet Rule")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, ruleName, *res.Name, "Created rule name should match")
	createdRuleID := *res.Id

	defer func() {
		t.Logf("Cleaning up Internet Rule with ID: %s", createdRuleID)
		_, errDel := client.SecurityRulesAPI.DeleteSecurityRulesByID(context.Background(), createdRuleID).Execute()
		require.NoError(t, errDel, "Failed to delete Internet Rule during cleanup")
	}()

	t.Logf("Successfully created Internet Rule: %s with ID: %s", ruleName, createdRuleID)
	// Verify the response matches the expected output
	assert.Equal(t, "Internet", *res.PolicyType, "PolicyType should match")
	require.Len(t, res.AllowWebApplication, 1, "Should have one allowed web application")
	assert.Equal(t, "4sync", *res.AllowWebApplication[0].Name, "AllowWebApplication name should match")
}

// Test_security_services_SecurityRulesAPIService_GetByID tests retrieving a Security Rule by ID.
func Test_security_services_SecurityRulesAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-get-rule-" + common.GenerateRandomString(6)
	rule := security_services.SecurityRules{
		Folder:      common.StringPtr("Shared"),
		Name:        common.StringPtr(ruleName),
		PolicyType:  common.StringPtr("Security"),
		From:        []string{"any"},
		To:          []string{"any"},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Application: []string{"any"},
		Action:      common.StringPtr("allow"),
		SourceUser:  []string{"any"},
		Category:    []string{"any"},
	}

	createRes, _, err := client.SecurityRulesAPI.CreateSecurityRules(context.Background()).Position("pre").SecurityRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for get test")
	createdRuleID := *createRes.Id

	defer func() {
		client.SecurityRulesAPI.DeleteSecurityRulesByID(context.Background(), createdRuleID).Execute()
	}()

	getRes, httpResGet, errGet := client.SecurityRulesAPI.GetSecurityRulesByID(context.Background(), createdRuleID).Execute()
	require.NoError(t, errGet, "Failed to get Security Rule by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, ruleName, *getRes.Name, "Rule name should match")
}

// Test_security_services_SecurityRulesAPIService_Update tests updating a Security Rule.
func Test_security_services_SecurityRulesAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-update-rule-" + common.GenerateRandomString(6)
	rule := security_services.SecurityRules{
		Folder:      common.StringPtr("Shared"),
		Name:        common.StringPtr(ruleName),
		PolicyType:  common.StringPtr("Security"),
		From:        []string{"any"},
		To:          []string{"any"},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Application: []string{"any"},
		Action:      common.StringPtr("allow"),
		SourceUser:  []string{"any"},
		Category:    []string{"any"},
	}

	createRes, _, err := client.SecurityRulesAPI.CreateSecurityRules(context.Background()).Position("pre").SecurityRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for update test")
	createdRuleID := *createRes.Id

	defer func() {
		client.SecurityRulesAPI.DeleteSecurityRulesByID(context.Background(), createdRuleID).Execute()
	}()

	updatedDescription := "Updated security rule description"
	updatedRule := security_services.SecurityRules{
		Name:        &ruleName,
		Description: common.StringPtr(updatedDescription),
		// All required fields must be sent on an update
		From:        []string{"any"},
		To:          []string{"any"},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Application: []string{"any"},
		Action:      common.StringPtr("allow"),
		SourceUser:  []string{"any"},
		Category:    []string{"any"},
	}

	updateRes, httpResUpdate, errUpdate := client.SecurityRulesAPI.UpdateSecurityRulesByID(context.Background(), createdRuleID).SecurityRules(updatedRule).Execute()
	require.NoError(t, errUpdate, "Failed to update Security Rule")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
}

// Test_security_services_SecurityRulesAPIService_List tests listing Security Rules.
func Test_security_services_SecurityRulesAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-list-rule-" + common.GenerateRandomString(6)
	rule := security_services.SecurityRules{
		Folder:      common.StringPtr("Shared"),
		Name:        common.StringPtr(ruleName),
		PolicyType:  common.StringPtr("Security"),
		From:        []string{"any"},
		To:          []string{"any"},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Application: []string{"any"},
		Action:      common.StringPtr("allow"),
		SourceUser:  []string{"any"},
		Category:    []string{"any"},
	}

	createRes, _, err := client.SecurityRulesAPI.CreateSecurityRules(context.Background()).Position("pre").SecurityRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for list test")
	createdRuleID := *createRes.Id

	defer func() {
		client.SecurityRulesAPI.DeleteSecurityRulesByID(context.Background(), createdRuleID).Execute()
	}()

	listRes, httpResList, errList := client.SecurityRulesAPI.ListRules(context.Background()).Position("pre").Folder("Shared").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list Security Rules")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	foundObject := false
	for _, r := range listRes.Data {
		if r.Name != nil && *r.Name == ruleName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created Security Rule should be found in the list")
}

// Test_security_services_SecurityRulesAPIService_DeleteByID tests deleting a Security Rule.
func Test_security_services_SecurityRulesAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	ruleName := "test-delete-rule-" + common.GenerateRandomString(6)
	rule := security_services.SecurityRules{
		Folder:      common.StringPtr("Shared"),
		Name:        common.StringPtr(ruleName),
		PolicyType:  common.StringPtr("Security"),
		From:        []string{"any"},
		To:          []string{"any"},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Application: []string{"any"},
		Action:      common.StringPtr("allow"),
		SourceUser:  []string{"any"},
		Category:    []string{"any"},
	}

	createRes, _, err := client.SecurityRulesAPI.CreateSecurityRules(context.Background()).Position("pre").SecurityRules(rule).Execute()
	require.NoError(t, err, "Failed to create rule for delete test")
	createdRuleID := *createRes.Id

	httpResDel, errDel := client.SecurityRulesAPI.DeleteSecurityRulesByID(context.Background(), createdRuleID).Execute()
	require.NoError(t, errDel, "Failed to delete Security Rule")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_security_services_SecurityRulesAPIService_Move tests moving a Security Rule.
func Test_security_services_SecurityRulesAPIService_Move(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	// Create two rules to test the move operation
	ruleAName := "test-move-rule-A-" + common.GenerateRandomString(4)
	ruleBName := "test-move-rule-B-" + common.GenerateRandomString(4)

	ruleA := security_services.SecurityRules{
		Folder: common.StringPtr("Shared"), Name: &ruleAName, PolicyType: common.StringPtr("Security"),
		From: []string{"any"}, To: []string{"any"}, Source: []string{"any"}, Destination: []string{"any"},
		Service: []string{"any"}, Application: []string{"any"}, Action: common.StringPtr("allow"),
		SourceUser: []string{"any"}, Category: []string{"any"},
	}
	ruleB := security_services.SecurityRules{
		Folder: common.StringPtr("Shared"), Name: &ruleBName, PolicyType: common.StringPtr("Security"),
		From: []string{"any"}, To: []string{"any"}, Source: []string{"any"}, Destination: []string{"any"},
		Service: []string{"any"}, Application: []string{"any"}, Action: common.StringPtr("allow"),
		SourceUser: []string{"any"}, Category: []string{"any"},
	}

	// Create rule A
	resA, _, errA := client.SecurityRulesAPI.CreateSecurityRules(context.Background()).Position("pre").SecurityRules(ruleA).Execute()
	require.NoError(t, errA, "Failed to create rule A for move test")
	idA := *resA.Id

	// Create rule B
	resB, _, errB := client.SecurityRulesAPI.CreateSecurityRules(context.Background()).Position("pre").SecurityRules(ruleB).Execute()
	require.NoError(t, errB, "Failed to create rule B for move test")
	idB := *resB.Id

	// Defer cleanup for both rules
	defer func() {
		client.SecurityRulesAPI.DeleteSecurityRulesByID(context.Background(), idA).Execute()
		client.SecurityRulesAPI.DeleteSecurityRulesByID(context.Background(), idB).Execute()
	}()

	// Define the move payload: move rule A after rule B
	movePayload := security_services.RuleBasedMove{
		Destination:     "after",
		DestinationRule: &idB,
		Rulebase:        "pre",
	}

	// Execute the move operation
	httpResMove, errMove := client.SecurityRulesAPI.MoveSecurityRulesByID(context.Background(), idA).RuleBasedMove(movePayload).Execute()
	require.NoError(t, errMove, "Failed to execute move operation for Security Rule")
	assert.Equal(t, 200, httpResMove.StatusCode, "Expected 200 OK status for move operation")

	t.Logf("Successfully executed move operation for rule %s", idA)
}
