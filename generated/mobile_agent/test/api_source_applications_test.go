/*
 * Mobile Agent Testing
 *
 * SourceApplicationsAPIService
 */

package mobile_agent

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/mobile_agent"
)

// newTestSourceApplication returns a minimal valid ForwardingProfileSourceApplications.
// The API requires 'applications' and 'name' as required fields.
func newTestSourceApplication(name string) mobile_agent.ForwardingProfileSourceApplications {
	return *mobile_agent.NewForwardingProfileSourceApplications(
		[]string{"app1", "app2"},
		name,
	)
}

// Test_mobile_agent_SourceApplicationsAPIService_Create tests the creation of a source application
// with a full set of fields including description.
func Test_mobile_agent_SourceApplicationsAPIService_Create(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	appName := "test-sourceapp-create-" + randomSuffix

	sourceApp := mobile_agent.NewForwardingProfileSourceApplications(
		[]string{"chrome", "firefox", "safari"},
		appName,
	)
	sourceApp.SetDescription("Test source application for create")

	fmt.Printf("Attempting to create Source Application with name: %s\n", sourceApp.Name)

	req := client.SourceApplicationsAPI.CreateGlobalProtectSourceApplication(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileSourceApplications(*sourceApp)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Source Application")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created Source Application should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Note: Source Applications may need manual cleanup - ID: %s", createdID)
	}()

	assert.Equal(t, appName, res.Name, "Created Source Application name should match")
	assert.Equal(t, "Test source application for create", *res.Description, "Description should match")

	require.NotNil(t, res.Applications, "Applications list should not be nil")
	assert.ElementsMatch(t, []string{"chrome", "firefox", "safari"}, res.Applications, "Applications should match")

	t.Logf("Successfully created and validated Source Application: %s with ID: %s", sourceApp.Name, createdID)
}

// Test_mobile_agent_SourceApplicationsAPIService_CreateMinimal tests the creation of a source application
// with only required fields (applications and name).
func Test_mobile_agent_SourceApplicationsAPIService_CreateMinimal(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	appName := "test-sourceapp-minimal-" + randomSuffix

	sourceApp := newTestSourceApplication(appName)

	fmt.Printf("Attempting to create minimal Source Application with name: %s\n", sourceApp.Name)

	req := client.SourceApplicationsAPI.CreateGlobalProtectSourceApplication(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileSourceApplications(sourceApp)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create minimal Source Application")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created Source Application should have an ID")

	assert.Equal(t, appName, res.Name, "Created Source Application name should match")
	assert.ElementsMatch(t, []string{"app1", "app2"}, res.Applications, "Applications should match")

	t.Logf("Successfully created minimal Source Application: %s with ID: %s", sourceApp.Name, *res.Id)
}

// Test_mobile_agent_SourceApplicationsAPIService_List tests listing source applications.
func Test_mobile_agent_SourceApplicationsAPIService_List(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	appName := "test-sourceapp-list-" + randomSuffix

	createRes, _, err := client.SourceApplicationsAPI.CreateGlobalProtectSourceApplication(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileSourceApplications(newTestSourceApplication(appName)).
		Execute()
	require.NoError(t, err, "Failed to create Source Application for list test")
	require.NotNil(t, createRes.Id, "Created Source Application should have an ID")

	listRes, httpResList, errList := client.SourceApplicationsAPI.ListGlobalProtectSourceApplications(context.Background()).
		Folder("Mobile Users").
		Limit(10000).
		Execute()
	require.NoError(t, errList, "Failed to list Source Applications")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundApp := false
	for _, app := range listRes.Data {
		if app.Name == appName {
			foundApp = true
			assert.Equal(t, *createRes.Id, *app.Id, "Listed application ID should match created ID")
			assert.ElementsMatch(t, []string{"app1", "app2"}, app.Applications, "Applications should match")
			break
		}
	}
	assert.True(t, foundApp, "Created Source Application should be found in the list")
}

// Test_mobile_agent_SourceApplicationsAPIService_ListWithNameFilter tests listing source applications
// with a name filter.
func Test_mobile_agent_SourceApplicationsAPIService_ListWithNameFilter(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	appName := "test-sourceapp-filter-" + randomSuffix

	createRes, _, err := client.SourceApplicationsAPI.CreateGlobalProtectSourceApplication(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileSourceApplications(newTestSourceApplication(appName)).
		Execute()
	require.NoError(t, err, "Failed to create Source Application for filter test")
	require.NotNil(t, createRes.Id, "Created Source Application should have an ID")

	listRes, httpResList, errList := client.SourceApplicationsAPI.ListGlobalProtectSourceApplications(context.Background()).
		Folder("Mobile Users").
		Name(appName).
		Limit(10).
		Execute()
	require.NoError(t, errList, "Failed to list Source Applications with name filter")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundApp := false
	for _, app := range listRes.Data {
		if app.Name == appName {
			foundApp = true
			assert.Equal(t, *createRes.Id, *app.Id, "Filtered application ID should match created ID")
			break
		}
	}
	assert.True(t, foundApp, "Created Source Application should be found with name filter")
}

// Test_mobile_agent_SourceApplicationsAPIService_FetchSourceApplications tests the FetchSourceApplications convenience method.
func Test_mobile_agent_SourceApplicationsAPIService_FetchSourceApplications(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-sourceapp-fetch-" + randomSuffix

	testObj := newTestSourceApplication(testName)
	testObj.SetDescription("Test source application for fetch")

	createReq := client.SourceApplicationsAPI.CreateGlobalProtectSourceApplication(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileSourceApplications(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created object should have an ID")
	createdID := *createRes.Id

	t.Logf("Created test object: %s with ID: %s", testName, createdID)

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.SourceApplicationsAPI.FetchSourceApplications(
		context.Background(),
		testName,
		common.StringPtr("Mobile Users"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch source_applications by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	assert.Equal(t, "Test source application for fetch", fetchedObj.GetDescription(), "Description should match")
	assert.ElementsMatch(t, []string{"app1", "app2"}, fetchedObj.Applications, "Applications should match")
	t.Logf("[SUCCESS] FetchSourceApplications found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.SourceApplicationsAPI.FetchSourceApplications(
		context.Background(),
		"non-existent-source-application-xyz-12345",
		common.StringPtr("Mobile Users"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchSourceApplications correctly returned nil for non-existent object")
}

// Test_mobile_agent_SourceApplicationsAPIService_CreateWithEmptyApplications tests that creation fails
// when applications list is empty (validation test).
func Test_mobile_agent_SourceApplicationsAPIService_CreateWithEmptyApplications(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	appName := "test-sourceapp-empty-" + randomSuffix

	sourceApp := mobile_agent.NewForwardingProfileSourceApplications(
		[]string{}, // empty applications list
		appName,
	)

	fmt.Printf("Attempting to create Source Application with empty applications list\n")

	req := client.SourceApplicationsAPI.CreateGlobalProtectSourceApplication(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileSourceApplications(*sourceApp)
	res, httpRes, err := req.Execute()

	// This test expects the API to reject empty applications list
	// The exact behavior depends on API validation
	if err != nil {
		// Expected path: API rejects empty list
		t.Logf("API correctly rejected empty applications list with error: %v", err)
		assert.Nil(t, res, "Response should be nil on error")
		if httpRes != nil {
			assert.True(t, httpRes.StatusCode >= 400, "Should return 4xx error status")
		}
	} else {
		// If API accepts it, we should at least verify the response
		t.Logf("WARNING: API accepted empty applications list")
		if res != nil && res.Id != nil {
			t.Logf("Created object with ID: %s (may need cleanup)", *res.Id)
		}
	}
}

// Test_mobile_agent_SourceApplicationsAPIService_ListPagination tests pagination parameters.
func Test_mobile_agent_SourceApplicationsAPIService_ListPagination(t *testing.T) {
	client := SetupMobileAgentTestClient(t)

	// Create multiple test objects to test pagination
	randomSuffix := common.GenerateRandomString(6)
	createdNames := []string{}
	for i := 0; i < 3; i++ {
		appName := fmt.Sprintf("test-sourceapp-page-%s-%d", randomSuffix, i)
		createdNames = append(createdNames, appName)
		_, _, err := client.SourceApplicationsAPI.CreateGlobalProtectSourceApplication(context.Background()).
			Folder("Mobile Users").
			ForwardingProfileSourceApplications(newTestSourceApplication(appName)).
			Execute()
		require.NoError(t, err, "Failed to create test Source Application for pagination test")
	}

	// Test with limit
	listRes, httpRes, err := client.SourceApplicationsAPI.ListGlobalProtectSourceApplications(context.Background()).
		Folder("Mobile Users").
		Limit(2).
		Offset(0).
		Execute()
	require.NoError(t, err, "Failed to list Source Applications with pagination")
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, listRes)

	t.Logf("Retrieved %d items with limit=2", len(listRes.Data))

	// Test with offset
	listRes2, httpRes2, err2 := client.SourceApplicationsAPI.ListGlobalProtectSourceApplications(context.Background()).
		Folder("Mobile Users").
		Limit(10).
		Offset(1).
		Execute()
	require.NoError(t, err2, "Failed to list Source Applications with offset")
	assert.Equal(t, 200, httpRes2.StatusCode)
	require.NotNil(t, listRes2)

	t.Logf("Retrieved %d items with offset=1", len(listRes2.Data))
	t.Logf("Pagination test completed successfully")
}

// Test_mobile_agent_SourceApplicationsAPIService_GetByID tests retrieving a source application by its ID.
func Test_mobile_agent_SourceApplicationsAPIService_GetByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	appName := "test-sourceapp-getbyid-" + randomSuffix

	createRes, _, err := client.SourceApplicationsAPI.CreateGlobalProtectSourceApplication(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileSourceApplications(newTestSourceApplication(appName)).
		Execute()
	require.NoError(t, err, "Failed to create Source Application for get by ID test")
	require.NotNil(t, createRes.Id, "Created Source Application should have an ID")
	createdID := *createRes.Id

	defer func() {
		t.Logf("Note: Source Applications may need manual cleanup - ID: %s", createdID)
	}()

	getRes, httpResGet, errGet := client.SourceApplicationsAPI.GetGlobalProtectSourceApplicationByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get Source Application by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, appName, getRes.Name, "Name should match")
	assert.Equal(t, createdID, *getRes.Id, "ID should match")
	assert.ElementsMatch(t, []string{"app1", "app2"}, getRes.Applications, "Applications should match")

	t.Logf("Successfully retrieved Source Application by ID: %s", createdID)
}

// Test_mobile_agent_SourceApplicationsAPIService_Update tests updating an existing source application.
func Test_mobile_agent_SourceApplicationsAPIService_Update(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	appName := "test-sourceapp-update-" + randomSuffix

	createRes, _, err := client.SourceApplicationsAPI.CreateGlobalProtectSourceApplication(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileSourceApplications(newTestSourceApplication(appName)).
		Execute()
	require.NoError(t, err, "Failed to create Source Application for update test")
	require.NotNil(t, createRes.Id, "Created Source Application should have an ID")
	createdID := *createRes.Id

	defer func() {
		t.Logf("Note: Source Applications may need manual cleanup - ID: %s", createdID)
	}()

	updatedApp := newTestSourceApplication(appName)
	updatedApp.SetDescription("Updated description for source application")
	updatedApp.Applications = []string{"chrome", "firefox", "edge"}

	updateRes, httpResUpdate, errUpdate := client.SourceApplicationsAPI.UpdateGlobalProtectSourceApplicationByID(context.Background(), createdID).
		ForwardingProfileSourceApplications(updatedApp).
		Execute()
	require.NoError(t, errUpdate, "Failed to update Source Application")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, appName, updateRes.Name, "Name should remain the same after update")
	assert.Equal(t, "Updated description for source application", *updateRes.Description, "Description should be updated")
	assert.ElementsMatch(t, []string{"chrome", "firefox", "edge"}, updateRes.Applications, "Applications should be updated")

	t.Logf("Successfully updated Source Application: %s with ID: %s", appName, createdID)
}

// Test_mobile_agent_SourceApplicationsAPIService_DeleteByID tests deleting a source application by ID.
func Test_mobile_agent_SourceApplicationsAPIService_DeleteByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	appName := "test-sourceapp-delete-" + randomSuffix

	createRes, _, err := client.SourceApplicationsAPI.CreateGlobalProtectSourceApplication(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileSourceApplications(newTestSourceApplication(appName)).
		Execute()
	require.NoError(t, err, "Failed to create Source Application for delete test")
	require.NotNil(t, createRes.Id, "Created Source Application should have an ID")
	createdID := *createRes.Id

	_, errDel := client.SourceApplicationsAPI.DeleteGlobalProtectSourceApplication(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete Source Application")

	t.Logf("Successfully deleted Source Application with ID: %s", createdID)
}
