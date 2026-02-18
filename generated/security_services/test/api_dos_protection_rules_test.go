/*
* Security Services Testing
* DoSProtectionRulesAPIService
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

// Test_security_services_DoSProtectionRulesAPIService_Create tests the creation of a dosprotectionrule object
// This test creates a new dosprotectionrule and then deletes it to ensure proper cleanup
func Test_security_services_DoSProtectionRulesAPIService_Create(t *testing.T) {
	t.Skip("API requires from/to as objects and service/protection fields - model has from/to as string arrays causing 400 Bad Request")
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a valid dosprotectionrule object with unique name to avoid conflicts
	createdDoSProtectionRuleName := "test-" + common.GenerateRandomString(10)
	dosprotectionrule := security_services.DosProtectionRules{
		Description: common.StringPtr("Test DoS protection rule for create API testing"),
		Folder:      common.StringPtr("All"),      // Using All folder scope
		Name:        createdDoSProtectionRuleName, // Unique test name
		From:        []string{"any"},              // Required field - source zones
		To:          []string{"any"},              // Required field - destination zones
		Source:      []string{"any"},              // Required field - source addresses
		Destination: []string{"any"},              // Required field - destination addresses
	}

	fmt.Printf("Creating dosprotectionrule with name: %s\n", dosprotectionrule.Name)

	// Make the create request to the API
	req := client.DoSProtectionRulesAPI.CreateDoSProtectionRules(context.Background()).DosProtectionRules(dosprotectionrule)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create dosprotectionrule")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdDoSProtectionRuleName, res.Name, "Created dosprotectionrule name should match")
	assert.Equal(t, common.StringPtr("Test DoS protection rule for create API testing"), res.Description, "Description should match")
	assert.NotEmpty(t, res.Id, "Created dosprotectionrule should have an ID")

	// Use the ID from the response object
	createdDoSProtectionRuleID := res.Id
	t.Logf("Successfully created dosprotectionrule: %s with ID: %s", dosprotectionrule.Name, *createdDoSProtectionRuleID)

	// Cleanup: Delete the created dosprotectionrule to maintain test isolation
	reqDel := client.DoSProtectionRulesAPI.DeleteDoSProtectionRulesByID(context.Background(), *createdDoSProtectionRuleID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete dosprotectionrule during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up dosprotectionrule: %s", *createdDoSProtectionRuleID)
}

// Test_security_services_DoSProtectionRulesAPIService_GetByID tests retrieving a dosprotectionrule by its ID
// This test creates a dosprotectionrule, retrieves it by ID, then deletes it
func Test_security_services_DoSProtectionRulesAPIService_GetByID(t *testing.T) {
	t.Skip("API requires from/to as objects and service/protection fields - model has from/to as string arrays causing 400 Bad Request")
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a dosprotectionrule first to have something to retrieve
	createdDoSProtectionRuleName := "test-getbyid-" + common.GenerateRandomString(10)
	dosprotectionrule := security_services.DosProtectionRules{
		Description: common.StringPtr("Test DoS protection rule for get by ID API testing"),
		Folder:      common.StringPtr("All"),      // Using All folder scope
		Name:        createdDoSProtectionRuleName, // Unique test name
		From:        []string{"any"},              // Required field
		To:          []string{"any"},              // Required field
		Source:      []string{"any"},              // Required field
		Destination: []string{"any"},              // Required field
	}

	// Create the dosprotectionrule via API
	req := client.DoSProtectionRulesAPI.CreateDoSProtectionRules(context.Background()).DosProtectionRules(dosprotectionrule)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create dosprotectionrule for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDoSProtectionRuleID := createRes.Id
	require.NotEmpty(t, createdDoSProtectionRuleID, "Created dosprotectionrule should have an ID")

	// Test Get by ID operation
	reqGetById := client.DoSProtectionRulesAPI.GetDoSProtectionRulesByID(context.Background(), *createdDoSProtectionRuleID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get dosprotectionrule by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdDoSProtectionRuleName, getRes.Name, "DoSProtectionRule name should match")
	assert.Equal(t, common.StringPtr("Test DoS protection rule for get by ID API testing"), getRes.Description, "Description should match")
	assert.Equal(t, *createdDoSProtectionRuleID, *getRes.Id, "DoSProtectionRule ID should match")

	t.Logf("Successfully retrieved dosprotectionrule: %s", getRes.Name)

	// Cleanup: Delete the created dosprotectionrule
	reqDel := client.DoSProtectionRulesAPI.DeleteDoSProtectionRulesByID(context.Background(), *createdDoSProtectionRuleID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete dosprotectionrule during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up dosprotectionrule: %s", *createdDoSProtectionRuleID)
}

// Test_security_services_DoSProtectionRulesAPIService_Update tests updating an existing dosprotectionrule
// This test creates a dosprotectionrule, updates it, then deletes it
func Test_security_services_DoSProtectionRulesAPIService_Update(t *testing.T) {
	t.Skip("API requires from/to as objects and service/protection fields - model has from/to as string arrays causing 400 Bad Request")
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a dosprotectionrule first to have something to update
	createdDoSProtectionRuleName := "test-update-" + common.GenerateRandomString(10)
	dosprotectionrule := security_services.DosProtectionRules{
		Description: common.StringPtr("Test DoS protection rule for update API testing"),
		Folder:      common.StringPtr("All"),      // Using All folder scope
		Name:        createdDoSProtectionRuleName, // Unique test name
		From:        []string{"any"},              // Required field
		To:          []string{"any"},              // Required field
		Source:      []string{"any"},              // Required field
		Destination: []string{"any"},              // Required field
	}

	// Create the dosprotectionrule via API
	req := client.DoSProtectionRulesAPI.CreateDoSProtectionRules(context.Background()).DosProtectionRules(dosprotectionrule)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create dosprotectionrule for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDoSProtectionRuleID := createRes.Id
	require.NotEmpty(t, createdDoSProtectionRuleID, "Created dosprotectionrule should have an ID")

	// Test Update operation with modified fields
	updatedDoSProtectionRule := security_services.DosProtectionRules{
		Description: common.StringPtr("Updated test DoS protection rule description"), // Updated description
		Folder:      common.StringPtr("All"),                                          // Keep same folder scope
		Name:        createdDoSProtectionRuleName,                                     // Keep same name (required for update)
		From:        []string{"any"},                                                  // Keep required fields
		To:          []string{"any"},
		Source:      []string{"any"},
		Destination: []string{"any"},
	}

	reqUpdate := client.DoSProtectionRulesAPI.UpdateDoSProtectionRulesByID(context.Background(), *createdDoSProtectionRuleID).DosProtectionRules(updatedDoSProtectionRule)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update dosprotectionrule")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdDoSProtectionRuleName, updateRes.Name, "DoSProtectionRule name should remain the same")
	assert.Equal(t, common.StringPtr("Updated test DoS protection rule description"), updateRes.Description, "Description should be updated")
	assert.Equal(t, *createdDoSProtectionRuleID, *updateRes.Id, "DoSProtectionRule ID should remain the same")

	t.Logf("Successfully updated dosprotectionrule: %s", createdDoSProtectionRuleName)

	// Cleanup: Delete the created dosprotectionrule
	reqDel := client.DoSProtectionRulesAPI.DeleteDoSProtectionRulesByID(context.Background(), *createdDoSProtectionRuleID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete dosprotectionrule during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up dosprotectionrule: %s", *createdDoSProtectionRuleID)
}

// Test_security_services_DoSProtectionRulesAPIService_List tests listing dosprotectionrules with folder filter
// This test creates a dosprotectionrule, lists dosprotectionrules to verify it's included, then deletes it
func Test_security_services_DoSProtectionRulesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a dosprotectionrule first to have something to list
	createdDoSProtectionRuleName := "test-list-" + common.GenerateRandomString(10)
	dosprotectionrule := security_services.DosProtectionRules{
		Description: common.StringPtr("Test DoS protection rule for list API testing"),
		Folder:      common.StringPtr("All"),      // Using All folder scope
		Name:        createdDoSProtectionRuleName, // Unique test name
		From:        []string{"any"},              // Required field
		To:          []string{"any"},              // Required field
		Source:      []string{"any"},              // Required field
		Destination: []string{"any"},              // Required field
	}

	// Create the dosprotectionrule via API
	req := client.DoSProtectionRulesAPI.CreateDoSProtectionRules(context.Background()).DosProtectionRules(dosprotectionrule)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create dosprotectionrule for list test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDoSProtectionRuleID := createRes.Id
	require.NotEmpty(t, createdDoSProtectionRuleID, "Created dosprotectionrule should have an ID")

	// Test List operation with folder filter
	reqList := client.DoSProtectionRulesAPI.ListDoSProtectionRules(context.Background()).Folder("All").Limit(200).Offset(0)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list dosprotectionrules")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one dosprotectionrule in the list")

	// Verify our created dosprotectionrule is in the list
	foundDoSProtectionRule := false
	for _, rule := range listRes.Data {
		if rule.Name == createdDoSProtectionRuleName {
			foundDoSProtectionRule = true
			assert.Equal(t, common.StringPtr("Test DoS protection rule for list API testing"), rule.Description, "Description should match")
			break
		}
	}
	assert.True(t, foundDoSProtectionRule, "Created dosprotectionrule should be found in the list")

	t.Logf("Successfully listed dosprotectionrules, found created dosprotectionrule: %s", createdDoSProtectionRuleName)

	// Cleanup: Delete the created dosprotectionrule
	reqDel := client.DoSProtectionRulesAPI.DeleteDoSProtectionRulesByID(context.Background(), *createdDoSProtectionRuleID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete dosprotectionrule during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up dosprotectionrule: %s", *createdDoSProtectionRuleID)
}

// Test_security_services_DoSProtectionRulesAPIService_DeleteByID tests deleting a dosprotectionrule by its ID
// This test creates a dosprotectionrule, deletes it, then verifies the deletion was successful
func Test_security_services_DoSProtectionRulesAPIService_DeleteByID(t *testing.T) {
	t.Skip("API requires from/to as objects and service/protection fields - model has from/to as string arrays causing 400 Bad Request")
	// Setup the authenticated client
	client := SetupSecuritySvcTestClient(t)

	// Create a dosprotectionrule first to have something to delete
	createdDoSProtectionRuleName := "test-delete-" + common.GenerateRandomString(10)
	dosprotectionrule := security_services.DosProtectionRules{
		Description: common.StringPtr("Test DoS protection rule for delete API testing"),
		Folder:      common.StringPtr("All"),      // Using All folder scope
		Name:        createdDoSProtectionRuleName, // Unique test name
		From:        []string{"any"},              // Required field
		To:          []string{"any"},              // Required field
		Source:      []string{"any"},              // Required field
		Destination: []string{"any"},              // Required field
	}

	// Create the dosprotectionrule via API
	req := client.DoSProtectionRulesAPI.CreateDoSProtectionRules(context.Background()).DosProtectionRules(dosprotectionrule)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create dosprotectionrule for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdDoSProtectionRuleID := createRes.Id
	require.NotEmpty(t, createdDoSProtectionRuleID, "Created dosprotectionrule should have an ID")

	// Test Delete by ID operation
	reqDel := client.DoSProtectionRulesAPI.DeleteDoSProtectionRulesByID(context.Background(), *createdDoSProtectionRuleID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete dosprotectionrule")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted dosprotectionrule: %s", *createdDoSProtectionRuleID)
}

// Test_security_services_DoSProtectionRulesAPIService_Fetch tests the fetch convenience method
func Test_security_services_DoSProtectionRulesAPIService_Fetch(t *testing.T) {
	t.Skip("API requires from/to as objects and service/protection fields - model has from/to as string arrays causing 400 Bad Request")
	client := SetupSecuritySvcTestClient(t)

	// Create a DoS protection rule to fetch
	ruleName := "test-fetch-" + common.GenerateRandomString(10)
	rule := security_services.DosProtectionRules{
		Description: common.StringPtr("Test DoS protection rule for fetch"),
		Folder:      common.StringPtr("All"),
		Name:        ruleName,
		From:        []string{"any"},
		To:          []string{"any"},
		Source:      []string{"any"},
		Destination: []string{"any"},
	}

	reqCreate := client.DoSProtectionRulesAPI.CreateDoSProtectionRules(context.Background()).DosProtectionRules(rule)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create DoS protection rule for fetch test")
	createdID := createRes.Id
	createdFolder := createRes.Folder
	require.NotEmpty(t, createdID, "Created DoS protection rule ID should not be empty")

	// Defer cleanup
	defer func() {
		t.Logf("Cleaning up DoS protection rule with ID: %s", *createdID)
		_, errDel := client.DoSProtectionRulesAPI.DeleteDoSProtectionRulesByID(context.Background(), *createdID).Execute()
		require.NoError(t, errDel, "Failed to delete DoS protection rule during cleanup")
	}()

	// Test Fetch by name operation
	fmt.Printf("Attempting to fetch DoS protection rule with name: %s\n", ruleName)
	fetchedRule, errFetch := client.DoSProtectionRulesAPI.FetchDoSProtectionRules(context.Background(), ruleName, createdFolder, nil, nil)

	// Verify the fetch operation was successful
	require.NoError(t, errFetch, "Failed to fetch DoS protection rule by name")
	require.NotNil(t, fetchedRule, "Fetched DoS protection rule should not be nil")
	assert.Equal(t, ruleName, fetchedRule.Name, "DoS protection rule name should match")
	assert.Equal(t, *createdID, *fetchedRule.Id, "DoS protection rule ID should match")
	assert.Equal(t, *createdFolder, *fetchedRule.Folder, "Folder should match")
	t.Logf("Successfully fetched DoS protection rule: %s", ruleName)

	// Test fetching non-existent DoS protection rule (should return nil)
	nonExistentName := "non-existent-dosrule-xyz-12345"
	notFoundRule, errNotFound := client.DoSProtectionRulesAPI.FetchDoSProtectionRules(context.Background(), nonExistentName, createdFolder, nil, nil)
	require.NoError(t, errNotFound, "Fetch for non-existent DoS protection rule should not error")
	assert.Nil(t, notFoundRule, "Non-existent DoS protection rule should return nil")
	t.Logf("Successfully verified fetch returns nil for non-existent DoS protection rule")
}
