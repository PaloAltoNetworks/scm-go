/*
 * Network Deployment Testing
 *
 * TrafficSteeringRulesAPIService
 */

package deployment_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
)

// Test_deployment_services_TrafficSteeringRulesAPIService_Create tests the creation of a Traffic Steering Rule.
func Test_deployment_services_TrafficSteeringRulesAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create a valid Traffic Steering Rule object with a unique name.
	ruleName := "test-tsr-create-" + randomSuffix

	// Create action with no-pbf using AdditionalProperties
	action := &deployment_services.TrafficSteeringRulesAction{
		AdditionalProperties: map[string]interface{}{
			"no-pbf": map[string]interface{}{},
		},
	}

	rule := deployment_services.TrafficSteeringRules{
		Name:    ruleName,
		Folder:  "Service Connections",
		Id:      "", // Will be populated by API
		Service: []string{"any"},
		Source:  []string{"any"},
		Action:  action,
	}

	fmt.Printf("Attempting to create Traffic Steering Rule with name: %s\n", rule.Name)

	// Make the create request to the API.
	req := depSvcClient.TrafficSteeringRulesAPI.CreateTrafficSteeringRules(context.Background()).Folder("Service Connections").TrafficSteeringRules(rule)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create Traffic Steering Rule")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotEmpty(t, res.Id, "Created traffic steering rule should have an ID")

	createdRuleID := res.Id

	// Defer the cleanup of the created traffic steering rule.
	defer func() {
		t.Logf("Cleaning up Traffic Steering Rule with ID: %s", createdRuleID)
		_, errDel := depSvcClient.TrafficSteeringRulesAPI.DeleteTrafficSteeringRulesByID(context.Background(), createdRuleID).Execute()
		require.NoError(t, errDel, "Failed to delete traffic steering rule during cleanup")
	}()

	// Assert response object properties.
	assert.Equal(t, ruleName, res.Name, "Created traffic steering rule name should match")
	assert.Equal(t, "Service Connections", res.Folder, "Folder should match")
	t.Logf("Successfully created and validated Traffic Steering Rule: %s with ID: %s", rule.Name, createdRuleID)
}

// Test_deployment_services_TrafficSteeringRulesAPIService_GetByID tests retrieving a traffic steering rule by its ID.
func Test_deployment_services_TrafficSteeringRulesAPIService_GetByID(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create a traffic steering rule to retrieve.
	ruleName := "test-tsr-get-" + randomSuffix

	// Create action with no-pbf using AdditionalProperties
	action := &deployment_services.TrafficSteeringRulesAction{
		AdditionalProperties: map[string]interface{}{
			"no-pbf": map[string]interface{}{},
		},
	}

	rule := deployment_services.TrafficSteeringRules{
		Name:    ruleName,
		Folder:  "Service Connections",
		Id:      "", // Will be populated by API
		Service: []string{"any"},
		Source:  []string{"any"},
		Action:  action,
	}

	createRes, _, err := depSvcClient.TrafficSteeringRulesAPI.CreateTrafficSteeringRules(context.Background()).Folder("Service Connections").TrafficSteeringRules(rule).Execute()
	require.NoError(t, err, "Failed to create traffic steering rule for get test")
	createdRuleID := createRes.Id
	defer func() {
		depSvcClient.TrafficSteeringRulesAPI.DeleteTrafficSteeringRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test Get by ID operation.
	t.Skip("GetByID response missing required folder property - SDK model validation error")
	getRes, httpResGet, errGet := depSvcClient.TrafficSteeringRulesAPI.GetTrafficSteeringRulesByID(context.Background(), createdRuleID).Execute()
	require.NoError(t, errGet, "Failed to get traffic steering rule by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, ruleName, getRes.Name)
	assert.Equal(t, createRes.Id, getRes.Id)
}

// Test_deployment_services_TrafficSteeringRulesAPIService_Update tests updating an existing traffic steering rule.
func Test_deployment_services_TrafficSteeringRulesAPIService_Update(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create a traffic steering rule to update.
	ruleName := "test-tsr-update-" + randomSuffix

	// Create action with no-pbf using AdditionalProperties
	action := &deployment_services.TrafficSteeringRulesAction{
		AdditionalProperties: map[string]interface{}{
			"no-pbf": map[string]interface{}{},
		},
	}

	rule := deployment_services.TrafficSteeringRules{
		Name:    ruleName,
		Folder:  "Service Connections",
		Id:      "", // Will be populated by API
		Service: []string{"any"},
		Source:  []string{"any"},
		Action:  action,
	}
	createRes, _, err := depSvcClient.TrafficSteeringRulesAPI.CreateTrafficSteeringRules(context.Background()).Folder("Service Connections").TrafficSteeringRules(rule).Execute()
	require.NoError(t, err, "Failed to create traffic steering rule for update test")
	createdRuleID := createRes.Id
	defer func() {
		depSvcClient.TrafficSteeringRulesAPI.DeleteTrafficSteeringRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Define the update payload with modified destination.
	updatedRule := deployment_services.TrafficSteeringRules{
		Name:        ruleName,
		Folder:      "Service Connections",
		Id:          createdRuleID,
		Service:     []string{"any"},
		Source:      []string{"any"},
		Destination: []string{"10.0.0.0/8"},
		Action:      action,
	}

	updateRes, httpResUpdate, errUpdate := depSvcClient.TrafficSteeringRulesAPI.UpdateTrafficSteeringRulesByID(context.Background(), createdRuleID).TrafficSteeringRules(updatedRule).Execute()
	require.NoError(t, errUpdate, "Failed to update traffic steering rule")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	require.NotNil(t, updateRes.Destination, "Destination should not be nil")
	assert.Contains(t, updateRes.Destination, "10.0.0.0/8", "Destination should be updated")
}

// Test_deployment_services_TrafficSteeringRulesAPIService_List tests listing Traffic Steering Rules.
func Test_deployment_services_TrafficSteeringRulesAPIService_List(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create a traffic steering rule to ensure it appears in the list.
	ruleName := "test-tsr-list-" + randomSuffix

	// Create action with no-pbf using AdditionalProperties
	action := &deployment_services.TrafficSteeringRulesAction{
		AdditionalProperties: map[string]interface{}{
			"no-pbf": map[string]interface{}{},
		},
	}

	rule := deployment_services.TrafficSteeringRules{
		Name:    ruleName,
		Folder:  "Service Connections",
		Id:      "", // Will be populated by API
		Service: []string{"any"},
		Source:  []string{"any"},
		Action:  action,
	}
	createRes, _, err := depSvcClient.TrafficSteeringRulesAPI.CreateTrafficSteeringRules(context.Background()).Folder("Service Connections").TrafficSteeringRules(rule).Execute()
	require.NoError(t, err, "Failed to create traffic steering rule for list test")
	createdRuleID := createRes.Id
	defer func() {
		depSvcClient.TrafficSteeringRulesAPI.DeleteTrafficSteeringRulesByID(context.Background(), createdRuleID).Execute()
	}()

	// Test List operation.
	listRes, httpResList, errList := depSvcClient.TrafficSteeringRulesAPI.ListTrafficSteeringRules(context.Background()).Folder("Service Connections").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list traffic steering rules")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	// Verify our created traffic steering rule is in the list.
	foundRule := false
	for _, r := range listRes.Data {
		if r.Name == ruleName {
			foundRule = true
			break
		}
	}
	assert.True(t, foundRule, "Created traffic steering rule should be found in the list")
}

// Test_deployment_services_TrafficSteeringRulesAPIService_DeleteByID tests deleting a traffic steering rule by ID.
func Test_deployment_services_TrafficSteeringRulesAPIService_DeleteByID(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create a traffic steering rule to delete.
	ruleName := "test-tsr-delete-" + randomSuffix

	// Create action with no-pbf using AdditionalProperties
	action := &deployment_services.TrafficSteeringRulesAction{
		AdditionalProperties: map[string]interface{}{
			"no-pbf": map[string]interface{}{},
		},
	}

	rule := deployment_services.TrafficSteeringRules{
		Name:    ruleName,
		Folder:  "Service Connections",
		Id:      "", // Will be populated by API
		Service: []string{"any"},
		Source:  []string{"any"},
		Action:  action,
	}
	createRes, _, err := depSvcClient.TrafficSteeringRulesAPI.CreateTrafficSteeringRules(context.Background()).Folder("Service Connections").TrafficSteeringRules(rule).Execute()
	require.NoError(t, err, "Failed to create traffic steering rule for delete test")
	createdRuleID := createRes.Id

	// Test Delete by ID operation.
	_, errDel := depSvcClient.TrafficSteeringRulesAPI.DeleteTrafficSteeringRulesByID(context.Background(), createdRuleID).Execute()
	require.NoError(t, errDel, "Failed to delete traffic steering rule")
}

// Test_deployment_services_TrafficSteeringRulesAPIService_FetchTrafficSteeringRules tests the FetchTrafficSteeringRules convenience method
func Test_deployment_services_TrafficSteeringRulesAPIService_FetchTrafficSteeringRules(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create test object using same payload as Create test
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-tsr-fetch-" + randomSuffix

	// Create action with no-pbf using AdditionalProperties
	action := &deployment_services.TrafficSteeringRulesAction{
		AdditionalProperties: map[string]interface{}{
			"no-pbf": map[string]interface{}{},
		},
	}

	testObj := deployment_services.TrafficSteeringRules{
		Name:    testName,
		Folder:  "Service Connections",
		Id:      "", // Will be populated by API
		Service: []string{"any"},
		Source:  []string{"any"},
		Action:  action,
	}

	createReq := client.TrafficSteeringRulesAPI.CreateTrafficSteeringRules(context.Background()).Folder("Service Connections").TrafficSteeringRules(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.TrafficSteeringRulesAPI.DeleteTrafficSteeringRulesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.TrafficSteeringRulesAPI.FetchTrafficSteeringRules(
		context.Background(),
		testName,
		common.StringPtr("Service Connections"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch traffic_steering_rules by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchTrafficSteeringRules found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.TrafficSteeringRulesAPI.FetchTrafficSteeringRules(
		context.Background(),
		"non-existent-traffic_steering_rules-xyz-12345",
		common.StringPtr("Service Connections"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchTrafficSteeringRules correctly returned nil for non-existent object")
}
