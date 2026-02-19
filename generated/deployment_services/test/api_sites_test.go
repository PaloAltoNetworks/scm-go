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

// createTestRemoteNetworkForSite creates a remote network (and its IPsec tunnel dependencies)
// for use as a site member. Returns the remote network name and a cleanup function.
func createTestRemoteNetworkForSite(t *testing.T, depSvcClient *deployment_services.APIClient, suffix string) (string, func()) {
	// Create IPsec tunnel and its dependencies (IKE Crypto Profile, IKE Gateway)
	// Note: suffix kept short because IKE Crypto Profile names have maxLength of 31
	tunnelName, cleanupTunnelDeps := createTestIPsecTunnelAndDeps(t, suffix)

	// Create a remote network using the tunnel
	rnName := "test-rn-s-" + suffix
	rn := deployment_services.RemoteNetworks{
		Name:        rnName,
		Folder:      "Remote Networks",
		SpnName:     common.StringPtr("us-west-dakota"),
		LicenseType: "FWAAS-AGGREGATE",
		Region:      "us-west-2",
		IpsecTunnel: common.StringPtr(tunnelName),
	}

	createRes, _, err := depSvcClient.RemoteNetworksAPI.CreateRemoteNetworks(context.Background()).RemoteNetworks(rn).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create remote network for site test")
	t.Logf("Created dependency: Remote Network '%s' with id '%s'", rnName, createRes.Id)

	cleanup := func() {
		t.Logf("Cleaning up dependency: Remote Network '%s'", rnName)
		_, err := depSvcClient.RemoteNetworksAPI.DeleteRemoteNetworksByID(context.Background(), createRes.Id).Execute()
		if err != nil {
			t.Logf("Warning: failed to delete remote network %s: %v", rnName, err)
		}
		cleanupTunnelDeps()
	}

	return rnName, cleanup
}

// Test_deployment_services_SitesAPIService_Create tests the creation of a Site.
func Test_deployment_services_SitesAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create remote network prerequisite
	rnName, cleanupRN := createTestRemoteNetworkForSite(t, depSvcClient, "c"+randomSuffix)
	defer cleanupRN()

	// Create a valid Site object with a unique name.
	siteName := "test-site-create-" + randomSuffix
	site := deployment_services.Sites{
		Name:    siteName,
		City:    common.StringPtr("San Jose"),
		Country: common.StringPtr("US"),
		State:   common.StringPtr("California"),
		Members: []deployment_services.SitesMembersInner{
			{Name: rnName, Mode: "active", RemoteNetwork: common.StringPtr(rnName)},
		},
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
	require.NotNil(t, res.Id, "Created site should have an ID")

	createdSiteID := *res.Id

	// Defer the cleanup of the created site (must happen before remote network cleanup).
	defer func() {
		t.Logf("Cleaning up Site with ID: %s", createdSiteID)
		_, errDel := depSvcClient.SitesAPI.DeleteSitesByID(context.Background(), createdSiteID).Execute()
		if errDel != nil {
			t.Logf("Warning: failed to delete site %s: %v", createdSiteID, errDel)
		}
	}()

	// Assert response object properties.
	assert.Equal(t, siteName, res.Name, "Created site name should match")
	t.Logf("Successfully created and validated Site: %s with ID: %s", site.Name, createdSiteID)
}

// Test_deployment_services_SitesAPIService_GetByID tests retrieving a site by its ID.
func Test_deployment_services_SitesAPIService_GetByID(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create remote network prerequisite
	rnName, cleanupRN := createTestRemoteNetworkForSite(t, depSvcClient, "g"+randomSuffix)
	defer cleanupRN()

	// Create a site to retrieve.
	siteName := "test-site-get-" + randomSuffix
	site := deployment_services.Sites{
		Name:    siteName,
		City:    common.StringPtr("San Jose"),
		Country: common.StringPtr("US"),
		State:   common.StringPtr("California"),
		Members: []deployment_services.SitesMembersInner{
			{Name: rnName, Mode: "active", RemoteNetwork: common.StringPtr(rnName)},
		},
	}

	createRes, _, err := depSvcClient.SitesAPI.CreateSites(context.Background()).Sites(site).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create site for get test")
	createdSiteID := *createRes.Id
	defer func() {
		depSvcClient.SitesAPI.DeleteSitesByID(context.Background(), createdSiteID).Execute()
	}()

	// Test Get by ID operation.
	getRes, httpResGet, errGet := depSvcClient.SitesAPI.GetSitesByID(context.Background(), createdSiteID).Execute()
	require.NoError(t, errGet, "Failed to get site by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, siteName, getRes.Name)
	assert.Equal(t, *createRes.Id, *getRes.Id)
}

// Test_deployment_services_SitesAPIService_Update tests updating an existing site.
func Test_deployment_services_SitesAPIService_Update(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create remote network prerequisite
	rnName, cleanupRN := createTestRemoteNetworkForSite(t, depSvcClient, "u"+randomSuffix)
	defer cleanupRN()

	// Create a site to update.
	siteName := "test-site-update-" + randomSuffix
	site := deployment_services.Sites{
		Name:    siteName,
		City:    common.StringPtr("San Jose"),
		Country: common.StringPtr("US"),
		State:   common.StringPtr("California"),
		Members: []deployment_services.SitesMembersInner{
			{Name: rnName, Mode: "active", RemoteNetwork: common.StringPtr(rnName)},
		},
	}
	createRes, _, err := depSvcClient.SitesAPI.CreateSites(context.Background()).Sites(site).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create site for update test")
	createdSiteID := *createRes.Id
	defer func() {
		depSvcClient.SitesAPI.DeleteSitesByID(context.Background(), createdSiteID).Execute()
	}()

	// Define the update payload with modified address.
	updatedSite := deployment_services.Sites{
		Name:         siteName,
		Id:           createRes.Id,
		AddressLine1: common.StringPtr("123 Updated Street"),
		City:         common.StringPtr("San Jose"),
		Country:      common.StringPtr("US"),
		State:        common.StringPtr("California"),
		Members: []deployment_services.SitesMembersInner{
			{Name: rnName, Mode: "active", RemoteNetwork: common.StringPtr(rnName)},
		},
	}

	updateRes, httpResUpdate, errUpdate := depSvcClient.SitesAPI.UpdateSitesByID(context.Background(), createdSiteID).Sites(updatedSite).Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}
	require.NoError(t, errUpdate, "Failed to update site")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, "123 Updated Street", *updateRes.AddressLine1, "Address should be updated")
}

// Test_deployment_services_SitesAPIService_List tests listing Sites (read-only).
func Test_deployment_services_SitesAPIService_List(t *testing.T) {
	client := SetupDeploymentSvcTestClient(t)

	listRes, httpResList, errList := client.SitesAPI.ListSites(context.Background()).Folder("All").Limit(200).Offset(0).Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list sites")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed sites, total: %d", listRes.GetTotal())
}

// Test_deployment_services_SitesAPIService_DeleteByID tests deleting a site by ID.
func Test_deployment_services_SitesAPIService_DeleteByID(t *testing.T) {
	depSvcClient := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create remote network prerequisite
	rnName, cleanupRN := createTestRemoteNetworkForSite(t, depSvcClient, "d"+randomSuffix)
	defer cleanupRN()

	// Create a site to delete.
	siteName := "test-site-delete-" + randomSuffix
	site := deployment_services.Sites{
		Name:    siteName,
		City:    common.StringPtr("San Jose"),
		Country: common.StringPtr("US"),
		State:   common.StringPtr("California"),
		Members: []deployment_services.SitesMembersInner{
			{Name: rnName, Mode: "active", RemoteNetwork: common.StringPtr(rnName)},
		},
	}
	createRes, _, err := depSvcClient.SitesAPI.CreateSites(context.Background()).Sites(site).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create site for delete test")
	createdSiteID := *createRes.Id

	// Test Delete by ID operation.
	_, errDel := depSvcClient.SitesAPI.DeleteSitesByID(context.Background(), createdSiteID).Execute()
	require.NoError(t, errDel, "Failed to delete site")
}

// Test_deployment_services_SitesAPIService_FetchSites tests the FetchSites convenience method.
func Test_deployment_services_SitesAPIService_FetchSites(t *testing.T) {
	client := SetupDeploymentSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)

	// Create remote network prerequisite
	rnName, cleanupRN := createTestRemoteNetworkForSite(t, client, "f"+randomSuffix)

	// Create a site to fetch
	siteName := "test-site-fetch-" + randomSuffix
	site := deployment_services.Sites{
		Name:    siteName,
		City:    common.StringPtr("San Jose"),
		Country: common.StringPtr("US"),
		State:   common.StringPtr("California"),
		Members: []deployment_services.SitesMembersInner{
			{Name: rnName, Mode: "active", RemoteNetwork: common.StringPtr(rnName)},
		},
	}

	createRes, _, err := client.SitesAPI.CreateSites(context.Background()).Sites(site).Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create site for fetch test")
	createdSiteID := *createRes.Id

	// Cleanup after test
	defer func() {
		_, _ = client.SitesAPI.DeleteSitesByID(context.Background(), createdSiteID).Execute()
		t.Logf("Cleaned up site: %s", createdSiteID)
		cleanupRN()
	}()

	// Test 1: Fetch existing site by name
	fetchedSite, err := client.SitesAPI.FetchSites(
		context.Background(),
		siteName,
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Failed to fetch site by name")
	require.NotNil(t, fetchedSite, "Fetched site should not be nil")
	assert.Equal(t, createdSiteID, *fetchedSite.Id, "Fetched site ID should match")
	assert.Equal(t, siteName, fetchedSite.Name, "Fetched site name should match")
	t.Logf("[SUCCESS] FetchSites found object: %s", fetchedSite.Name)

	// Test 2: Fetch non-existent site (should return nil, nil)
	notFound, err := client.SitesAPI.FetchSites(
		context.Background(),
		"non-existent-site-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchSites correctly returned nil for non-existent object")
}
