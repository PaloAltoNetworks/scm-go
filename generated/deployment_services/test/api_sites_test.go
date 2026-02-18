/*
 * Network Deployment Testing
 *
 * SitesAPIService
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

// Test_deployment_services_SitesAPIService_Create tests the creation of a Site.
func Test_deployment_services_SitesAPIService_Create(t *testing.T) {
	t.Skip("Skipping test: Sites model has non-pointer Id field (Id string), which causes issues with object creation as the API populates this field")

	// Setup the authenticated client.
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create a valid Site object with a unique name.
	siteName := "test-site-create-" + randomSuffix
	site := deployment_services.Sites{
		Name: siteName,
		Id:   "", // Will be populated by API
	}

	fmt.Printf("Attempting to create Site with name: %s\n", site.Name)

	// Make the create request to the API.
	req := depSvcClient.SitesAPI.CreateSites(context.Background()).Sites(site)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create Site")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotEmpty(t, res.Id, "Created site should have an ID")

	createdSiteID := res.Id

	// Defer the cleanup of the created site.
	defer func() {
		t.Logf("Cleaning up Site with ID: %s", createdSiteID)
		_, errDel := depSvcClient.SitesAPI.DeleteSitesByID(context.Background(), createdSiteID).Execute()
		require.NoError(t, errDel, "Failed to delete site during cleanup")
	}()

	// Assert response object properties.
	assert.Equal(t, siteName, res.Name, "Created site name should match")
	t.Logf("Successfully created and validated Site: %s with ID: %s", site.Name, createdSiteID)
}

// Test_deployment_services_SitesAPIService_GetByID tests retrieving a site by its ID.
func Test_deployment_services_SitesAPIService_GetByID(t *testing.T) {
	t.Skip("Skipping test: Sites model has non-pointer Id field (Id string), which causes issues with object creation as the API populates this field")

	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create a site to retrieve.
	siteName := "test-site-get-" + randomSuffix
	site := deployment_services.Sites{
		Name: siteName,
		Id:   "", // Will be populated by API
	}

	createRes, _, err := depSvcClient.SitesAPI.CreateSites(context.Background()).Sites(site).Execute()
	require.NoError(t, err, "Failed to create site for get test")
	createdSiteID := createRes.Id
	defer func() {
		depSvcClient.SitesAPI.DeleteSitesByID(context.Background(), createdSiteID).Execute()
	}()

	// Test Get by ID operation.
	getRes, httpResGet, errGet := depSvcClient.SitesAPI.GetSitesByID(context.Background(), createdSiteID).Execute()
	require.NoError(t, errGet, "Failed to get site by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, siteName, getRes.Name)
	assert.Equal(t, createRes.Id, getRes.Id)
}

// Test_deployment_services_SitesAPIService_Update tests updating an existing site.
func Test_deployment_services_SitesAPIService_Update(t *testing.T) {
	t.Skip("Skipping test: Sites model has non-pointer Id field (Id string), which causes issues with object creation as the API populates this field")

	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create a site to update.
	siteName := "test-site-update-" + randomSuffix
	site := deployment_services.Sites{
		Name: siteName,
		Id:   "", // Will be populated by API
	}
	createRes, _, err := depSvcClient.SitesAPI.CreateSites(context.Background()).Sites(site).Execute()
	require.NoError(t, err, "Failed to create site for update test")
	createdSiteID := createRes.Id
	defer func() {
		depSvcClient.SitesAPI.DeleteSitesByID(context.Background(), createdSiteID).Execute()
	}()

	// Define the update payload with modified address.
	updatedSite := deployment_services.Sites{
		Name:         siteName,
		Id:           createdSiteID,
		AddressLine1: common.StringPtr("123 Updated Street"),
		City:         common.StringPtr("UpdatedCity"),
	}

	updateRes, httpResUpdate, errUpdate := depSvcClient.SitesAPI.UpdateSitesByID(context.Background(), createdSiteID).Sites(updatedSite).Execute()
	require.NoError(t, errUpdate, "Failed to update site")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, "123 Updated Street", *updateRes.AddressLine1, "Address should be updated")
}

// Test_deployment_services_SitesAPIService_List tests listing Sites.
func Test_deployment_services_SitesAPIService_List(t *testing.T) {
	t.Skip("Skipping test: Sites model has non-pointer Id field (Id string), which causes issues with object creation as the API populates this field")

	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create a site to ensure it appears in the list.
	siteName := "test-site-list-" + randomSuffix
	site := deployment_services.Sites{
		Name: siteName,
		Id:   "", // Will be populated by API
	}
	createRes, _, err := depSvcClient.SitesAPI.CreateSites(context.Background()).Sites(site).Execute()
	require.NoError(t, err, "Failed to create site for list test")
	createdSiteID := createRes.Id
	defer func() {
		depSvcClient.SitesAPI.DeleteSitesByID(context.Background(), createdSiteID).Execute()
	}()

	// Test List operation.
	// Note: Sites API requires folder parameter
	listRes, httpResList, errList := depSvcClient.SitesAPI.ListSites(context.Background()).Folder("All").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list sites")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	// Verify our created site is in the list.
	foundSite := false
	for _, s := range listRes.Data {
		if s.Name == siteName {
			foundSite = true
			break
		}
	}
	assert.True(t, foundSite, "Created site should be found in the list")
}

// Test_deployment_services_SitesAPIService_DeleteByID tests deleting a site by ID.
func Test_deployment_services_SitesAPIService_DeleteByID(t *testing.T) {
	t.Skip("Skipping test: Sites model has non-pointer Id field (Id string), which causes issues with object creation as the API populates this field")

	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create a site to delete.
	siteName := "test-site-delete-" + randomSuffix
	site := deployment_services.Sites{
		Name: siteName,
		Id:   "", // Will be populated by API
	}
	createRes, _, err := depSvcClient.SitesAPI.CreateSites(context.Background()).Sites(site).Execute()
	require.NoError(t, err, "Failed to create site for delete test")
	createdSiteID := createRes.Id

	// Test Delete by ID operation.
	_, errDel := depSvcClient.SitesAPI.DeleteSitesByID(context.Background(), createdSiteID).Execute()
	require.NoError(t, errDel, "Failed to delete site")
}

// Test_deployment_services_SitesAPIService_FetchSites tests the FetchSites convenience method
func Test_deployment_services_SitesAPIService_FetchSites(t *testing.T) {
	t.Skip("Skipping test: Sites model has non-pointer Id field (Id string), which causes issues with object creation as the API populates this field")

	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create test object using same payload as Create test
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-site-fetch-" + randomSuffix

	testObj := deployment_services.Sites{
		Name: testName,
		Id:   "", // Will be populated by API
	}

	createReq := client.SitesAPI.CreateSites(context.Background()).Sites(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.SitesAPI.DeleteSitesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.SitesAPI.FetchSites(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch sites by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchSites found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.SitesAPI.FetchSites(
		context.Background(),
		"non-existent-sites-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchSites correctly returned nil for non-existent object")
}
