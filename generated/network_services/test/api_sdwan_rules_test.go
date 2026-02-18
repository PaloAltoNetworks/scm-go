/*
Network Services Testing SDWANRulesAPIService
*/
package network_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Test_networkservices_SDWANRulesAPIService_Create tests creating an SD-WAN rule
func Test_networkservices_SDWANRulesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object
	testName := "test-sdwan-rule-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanRules{
		Name:     testName,
		Folder:   common.StringPtr("All"),
		Position: "pre",
		Action: network_services.SdwanRulesAction{
			TrafficDistributionProfile: "Best Available Path",
		},
		Application:        []string{"any"},
		Destination:        []string{"any"},
		From:               []string{"any"},
		Service:            []string{"any"},
		Source:             []string{"any"},
		SourceUser:         []string{"any"},
		To:                 []string{"any"},
		PathQualityProfile: "Acceptable Performance",
	}

	createReq := client.SDWANRulesAPI.CreateSDWANRules(context.Background()).SdwanRules(testObj)
	resp, httpResp, err := createReq.Execute()

	// Cleanup after test
	if resp != nil && resp.Id != nil {
		defer func() {
			deleteReq := client.SDWANRulesAPI.DeleteSDWANRulesByID(context.Background(), *resp.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *resp.Id)
		}()
	}

	// Verify the response
	require.NoError(t, err, "Failed to create SD-WAN rule")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 201, httpResp.StatusCode, "Expected 201 Created status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testName, resp.Name, "Created rule name should match")
	t.Logf("[SUCCESS] Created SD-WAN rule: %s", resp.Name)
}

// Test_networkservices_SDWANRulesAPIService_List tests listing SD-WAN rules
func Test_networkservices_SDWANRulesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-rule-list-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanRules{
		Name:     testName,
		Folder:   common.StringPtr("All"),
		Position: "pre",
		Action: network_services.SdwanRulesAction{
			TrafficDistributionProfile: "Best Available Path",
		},
		Application:        []string{"any"},
		Destination:        []string{"any"},
		From:               []string{"any"},
		Service:            []string{"any"},
		Source:             []string{"any"},
		SourceUser:         []string{"any"},
		To:                 []string{"any"},
		PathQualityProfile: "Acceptable Performance",
	}

	createReq := client.SDWANRulesAPI.CreateSDWANRules(context.Background()).SdwanRules(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test
	defer func() {
		if createRes.Id != nil {
			deleteReq := client.SDWANRulesAPI.DeleteSDWANRulesByID(context.Background(), *createRes.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *createRes.Id)
		}
	}()

	// List SD-WAN rules
	listReq := client.SDWANRulesAPI.ListSDWANRules(context.Background()).Folder("All").Limit(200).Offset(0)
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list SD-WAN rules")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")

	// Verify our created object is in the list
	found := false
	if resp.Data != nil {
		for _, item := range resp.Data {
			if item.Name == testName {
				found = true
				break
			}
		}
	}
	assert.True(t, found, "Created rule should be in the list")
	t.Logf("[SUCCESS] Listed SD-WAN rules, found test object: %s", testName)
}

// Test_networkservices_SDWANRulesAPIService_GetByID tests retrieving an SD-WAN rule by ID
func Test_networkservices_SDWANRulesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-rule-get-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanRules{
		Name:     testName,
		Folder:   common.StringPtr("All"),
		Position: "pre",
		Action: network_services.SdwanRulesAction{
			TrafficDistributionProfile: "Best Available Path",
		},
		Application:        []string{"any"},
		Destination:        []string{"any"},
		From:               []string{"any"},
		Service:            []string{"any"},
		Source:             []string{"any"},
		SourceUser:         []string{"any"},
		To:                 []string{"any"},
		PathQualityProfile: "Acceptable Performance",
	}

	createReq := client.SDWANRulesAPI.CreateSDWANRules(context.Background()).SdwanRules(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Cleanup after test
	defer func() {
		deleteReq := client.SDWANRulesAPI.DeleteSDWANRulesByID(context.Background(), *createRes.Id)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", *createRes.Id)
	}()

	// Get by ID
	getReq := client.SDWANRulesAPI.GetSDWANRulesByID(context.Background(), *createRes.Id)
	getRes, httpRes, err := getReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to get SD-WAN rule by ID")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Response should not be nil")
	assert.Equal(t, *createRes.Id, *getRes.Id, "Retrieved rule ID should match")
	assert.Equal(t, testName, getRes.Name, "Retrieved rule name should match")
	t.Logf("[SUCCESS] Retrieved SD-WAN rule by ID: %s", *getRes.Id)
}

// Test_networkservices_SDWANRulesAPIService_UpdateByID tests updating an SD-WAN rule
func Test_networkservices_SDWANRulesAPIService_UpdateByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-rule-update-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanRules{
		Name:     testName,
		Folder:   common.StringPtr("All"),
		Position: "pre",
		Action: network_services.SdwanRulesAction{
			TrafficDistributionProfile: "Best Available Path",
		},
		Application:        []string{"any"},
		Destination:        []string{"any"},
		From:               []string{"any"},
		Service:            []string{"any"},
		Source:             []string{"any"},
		SourceUser:         []string{"any"},
		To:                 []string{"any"},
		PathQualityProfile: "Acceptable Performance",
	}

	createReq := client.SDWANRulesAPI.CreateSDWANRules(context.Background()).SdwanRules(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Cleanup after test
	defer func() {
		deleteReq := client.SDWANRulesAPI.DeleteSDWANRulesByID(context.Background(), *createRes.Id)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", *createRes.Id)
	}()

	// Update the object
	updatedObj := network_services.SdwanRules{
		Name:     testName,
		Position: "pre",
		Action: network_services.SdwanRulesAction{
			TrafficDistributionProfile: "Best Available Path",
		},
		Application:        []string{"any"},
		Destination:        []string{"any"},
		From:               []string{"any"},
		Service:            []string{"any"},
		Source:             []string{"any"},
		SourceUser:         []string{"any"},
		To:                 []string{"any"},
		PathQualityProfile: "Acceptable Performance",
		Description:        common.StringPtr("Updated description"), // Update field
	}

	updateReq := client.SDWANRulesAPI.UpdateSDWANRulesByID(context.Background(), *createRes.Id).SdwanRules(updatedObj)
	updateRes, httpRes, err := updateReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to update SD-WAN rule")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Response should not be nil")
	assert.Equal(t, "Updated description", *updateRes.Description, "Updated description should match")
	t.Logf("[SUCCESS] Updated SD-WAN rule: %s", *updateRes.Id)
}

// Test_networkservices_SDWANRulesAPIService_DeleteByID tests deleting an SD-WAN rule
func Test_networkservices_SDWANRulesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-rule-delete-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanRules{
		Name:     testName,
		Folder:   common.StringPtr("All"),
		Position: "pre",
		Action: network_services.SdwanRulesAction{
			TrafficDistributionProfile: "Best Available Path",
		},
		Application:        []string{"any"},
		Destination:        []string{"any"},
		From:               []string{"any"},
		Service:            []string{"any"},
		Source:             []string{"any"},
		SourceUser:         []string{"any"},
		To:                 []string{"any"},
		PathQualityProfile: "Acceptable Performance",
	}

	createReq := client.SDWANRulesAPI.CreateSDWANRules(context.Background()).SdwanRules(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object ID should not be nil")

	// Delete the object
	deleteReq := client.SDWANRulesAPI.DeleteSDWANRulesByID(context.Background(), *createRes.Id)
	httpRes, err := deleteReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to delete SD-WAN rule")
	require.NotNil(t, httpRes, "HTTP response should not be nil")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	t.Logf("[SUCCESS] Deleted SD-WAN rule: %s", *createRes.Id)
}

// Test_networkservices_SDWANRulesAPIService_FetchSDWANRules tests the Fetch convenience method
func Test_networkservices_SDWANRulesAPIService_FetchSDWANRules(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first
	testName := "test-sdwan-rule-fetch-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanRules{
		Name:     testName,
		Folder:   common.StringPtr("All"),
		Position: "pre",
		Action: network_services.SdwanRulesAction{
			TrafficDistributionProfile: "Best Available Path",
		},
		Application:        []string{"any"},
		Destination:        []string{"any"},
		From:               []string{"any"},
		Service:            []string{"any"},
		Source:             []string{"any"},
		SourceUser:         []string{"any"},
		To:                 []string{"any"},
		PathQualityProfile: "Acceptable Performance",
	}

	createReq := client.SDWANRulesAPI.CreateSDWANRules(context.Background()).SdwanRules(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test
	defer func() {
		if createRes.Id != nil {
			deleteReq := client.SDWANRulesAPI.DeleteSDWANRulesByID(context.Background(), *createRes.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", *createRes.Id)
		}
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.SDWANRulesAPI.FetchSDWANRules(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch SD-WAN rule by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchSDWANRules found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.SDWANRulesAPI.FetchSDWANRules(
		context.Background(),
		"non-existent-sdwan-rule-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchSDWANRules correctly returned nil for non-existent object")
}
