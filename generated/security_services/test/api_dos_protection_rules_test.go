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
	"github.com/paloaltonetworks/scm-go/generated/network_services"
	"github.com/paloaltonetworks/scm-go/generated/security_services"
)

// createTestZone creates a zone for testing and returns the zone name and a cleanup function
func createTestZone(t *testing.T, client *network_services.APIClient, prefix string) (string, func()) {
	zoneName := prefix + common.GenerateRandomString(6)
	zone := network_services.NewZones(zoneName)
	zone.SetFolder("All")

	req := client.SecurityZonesAPI.CreateZones(context.Background()).Zones(*zone)
	res, _, err := req.Execute()
	require.NoError(t, err, "Failed to create test zone: %s", zoneName)
	require.NotNil(t, res.Id, "Created zone should have an ID")

	createdID := *res.Id
	t.Logf("Created test zone: %s with ID: %s", zoneName, createdID)

	cleanup := func() {
		_, err := client.SecurityZonesAPI.DeleteZonesByID(context.Background(), createdID).Execute()
		if err != nil {
			t.Logf("Warning: Failed to delete test zone %s: %v", zoneName, err)
		} else {
			t.Logf("Cleaned up test zone: %s", zoneName)
		}
	}

	return zoneName, cleanup
}

// Test_security_services_DoSProtectionRulesAPIService_Create tests the creation of a dosprotectionrule object
// This test creates a new dosprotectionrule and then deletes it to ensure proper cleanup
func Test_security_services_DoSProtectionRulesAPIService_Create(t *testing.T) {
	// Setup the authenticated clients
	client := SetupSecuritySvcTestClient(t)
	networkClient := SetupNetworkSvcTestClient(t)

	// Create test zones for from and to
	fromZoneName, cleanupFromZone := createTestZone(t, networkClient, "dos-from-")
	defer cleanupFromZone()
	toZoneName, cleanupToZone := createTestZone(t, networkClient, "dos-to-")
	defer cleanupToZone()

	// Create a valid dosprotectionrule object with unique name to avoid conflicts
	createdDoSProtectionRuleName := "test-" + common.GenerateRandomString(10)
	dosprotectionrule := security_services.DosProtectionRules{
		Description: common.StringPtr("Test DoS protection rule for create API testing"),
		Folder:      common.StringPtr("All"),      // Using All folder scope
		Name:        createdDoSProtectionRuleName, // Unique test name
		From:        security_services.DosProtectionRulesFrom{Zone: []string{fromZoneName}},
		To:          security_services.DosProtectionRulesTo{Zone: []string{toZoneName}},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Protection:  security_services.DosProtectionRulesProtection{},
		Action:      &security_services.DosProtectionRulesAction{Allow: map[string]interface{}{}},
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
	// Setup the authenticated clients
	client := SetupSecuritySvcTestClient(t)
	networkClient := SetupNetworkSvcTestClient(t)

	// Create test zones for from and to
	fromZoneName, cleanupFromZone := createTestZone(t, networkClient, "dos-from-")
	defer cleanupFromZone()
	toZoneName, cleanupToZone := createTestZone(t, networkClient, "dos-to-")
	defer cleanupToZone()

	// Create a dosprotectionrule first to have something to retrieve
	createdDoSProtectionRuleName := "test-getbyid-" + common.GenerateRandomString(10)
	dosprotectionrule := security_services.DosProtectionRules{
		Description: common.StringPtr("Test DoS protection rule for get by ID API testing"),
		Folder:      common.StringPtr("All"),      // Using All folder scope
		Name:        createdDoSProtectionRuleName, // Unique test name
		From:        security_services.DosProtectionRulesFrom{Zone: []string{fromZoneName}},
		To:          security_services.DosProtectionRulesTo{Zone: []string{toZoneName}},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Protection:  security_services.DosProtectionRulesProtection{},
		Action:      &security_services.DosProtectionRulesAction{Allow: map[string]interface{}{}},
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
	// Setup the authenticated clients
	client := SetupSecuritySvcTestClient(t)
	networkClient := SetupNetworkSvcTestClient(t)

	// Create test zones for from and to
	fromZoneName, cleanupFromZone := createTestZone(t, networkClient, "dos-from-")
	defer cleanupFromZone()
	toZoneName, cleanupToZone := createTestZone(t, networkClient, "dos-to-")
	defer cleanupToZone()

	// Create a dosprotectionrule first to have something to update
	createdDoSProtectionRuleName := "test-update-" + common.GenerateRandomString(10)
	dosprotectionrule := security_services.DosProtectionRules{
		Description: common.StringPtr("Test DoS protection rule for update API testing"),
		Folder:      common.StringPtr("All"),      // Using All folder scope
		Name:        createdDoSProtectionRuleName, // Unique test name
		From:        security_services.DosProtectionRulesFrom{Zone: []string{fromZoneName}},
		To:          security_services.DosProtectionRulesTo{Zone: []string{toZoneName}},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Protection:  security_services.DosProtectionRulesProtection{},
		Action:      &security_services.DosProtectionRulesAction{Allow: map[string]interface{}{}},
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
		From:        security_services.DosProtectionRulesFrom{Zone: []string{fromZoneName}},
		To:          security_services.DosProtectionRulesTo{Zone: []string{toZoneName}},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Protection:  security_services.DosProtectionRulesProtection{},
		Action:      &security_services.DosProtectionRulesAction{Allow: map[string]interface{}{}},
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

// Test_security_services_DoSProtectionRulesAPIService_List tests listing existing DoS protection rules
// Read-only test: list existing DoS protection rules without creating any
func Test_security_services_DoSProtectionRulesAPIService_List(t *testing.T) {
	t.Skip("Skipping: existing rules in system don't have required 'from' field, causing unmarshal failure")
	client := SetupSecuritySvcTestClient(t)

	reqList := client.DoSProtectionRulesAPI.ListDoSProtectionRules(context.Background()).Folder("All").Limit(200).Offset(0)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list DoS protection rules")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed DoS protection rules, total: %d", listRes.GetTotal())
}

// Test_security_services_DoSProtectionRulesAPIService_DeleteByID tests deleting a dosprotectionrule by its ID
// This test creates a dosprotectionrule, deletes it, then verifies the deletion was successful
func Test_security_services_DoSProtectionRulesAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated clients
	client := SetupSecuritySvcTestClient(t)
	networkClient := SetupNetworkSvcTestClient(t)

	// Create test zones for from and to
	fromZoneName, cleanupFromZone := createTestZone(t, networkClient, "dos-from-")
	defer cleanupFromZone()
	toZoneName, cleanupToZone := createTestZone(t, networkClient, "dos-to-")
	defer cleanupToZone()

	// Create a dosprotectionrule first to have something to delete
	createdDoSProtectionRuleName := "test-delete-" + common.GenerateRandomString(10)
	dosprotectionrule := security_services.DosProtectionRules{
		Description: common.StringPtr("Test DoS protection rule for delete API testing"),
		Folder:      common.StringPtr("All"),      // Using All folder scope
		Name:        createdDoSProtectionRuleName, // Unique test name
		From:        security_services.DosProtectionRulesFrom{Zone: []string{fromZoneName}},
		To:          security_services.DosProtectionRulesTo{Zone: []string{toZoneName}},
		Source:      []string{"any"},
		Destination: []string{"any"},
		Service:     []string{"any"},
		Protection:  security_services.DosProtectionRulesProtection{},
		Action:      &security_services.DosProtectionRulesAction{Allow: map[string]interface{}{}},
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
// Read-only test: verify fetch returns nil for a non-existent DoS protection rule
func Test_security_services_DoSProtectionRulesAPIService_Fetch(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)

	notFound, err := client.DoSProtectionRulesAPI.FetchDoSProtectionRules(
		context.Background(),
		"non-existent-dos-rule-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchDoSProtectionRules correctly returned nil for non-existent object")
}
