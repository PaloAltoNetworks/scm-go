/*
Network Services Testing SDWANTrafficDistributionProfilesAPIService
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

// Note: SDWAN traffic distribution profiles require link tags to be associated with
// sdwan_interface_profiles before they can be referenced. The sdwan_interface_profiles
// API is not available in our OpenAPI spec, so CRUD tests cannot create the required
// prerequisite chain (link_tag -> sdwan_interface_profile -> traffic_distribution_profile).
// List test works since it only reads existing data.

// Test_networkservices_SDWANTrafficDistributionProfilesAPIService_Create tests creating an SD-WAN traffic distribution profile
func Test_networkservices_SDWANTrafficDistributionProfilesAPIService_Create(t *testing.T) {
	t.Skip("Requires sdwan_interface_profiles prerequisite which is not available in our OpenAPI spec - link tags must be associated with interface profiles first")
	client := SetupNetworkSvcTestClient(t)

	// Create prerequisite link tag
	ltName := "test-lt-tdp-" + common.GenerateRandomString(6)
	linkTag := network_services.LinkTags{
		Name:     ltName,
		Folder:   common.StringPtr("All"),
		Color:    common.StringPtr("color2"),
		Comments: common.StringPtr("Link tag for traffic distribution profile testing"),
	}
	ltReq := client.LinkTagsAPI.CreateLinkTags(context.Background()).LinkTags(linkTag)
	ltRes, _, err := ltReq.Execute()
	require.NoError(t, err, "Failed to create prerequisite link tag")
	require.NotNil(t, ltRes.Id, "Link tag should have an ID")
	defer func() {
		client.LinkTagsAPI.DeleteLinkTagsByID(context.Background(), *ltRes.Id).Execute()
	}()

	// Create a test object
	testName := "test-sdwan-tdp-" + common.GenerateRandomString(6)
	testObj := network_services.SdwanTrafficDistributionProfiles{
		Name:                testName,
		Folder:              common.StringPtr("All"),
		TrafficDistribution: common.StringPtr("Best Available Path"),
		LinkTags: []network_services.SdwanTrafficDistributionProfilesLinkTagsInner{
			{Name: ltName},
		},
	}

	createReq := client.SDWANTrafficDistributionProfilesAPI.CreateSDWANTrafficDistributionProfiles(context.Background()).SdwanTrafficDistributionProfiles(testObj)
	resp, httpResp, err := createReq.Execute()

	if resp != nil && resp.Id != nil {
		defer func() {
			client.SDWANTrafficDistributionProfilesAPI.DeleteSDWANTrafficDistributionProfilesByID(context.Background(), *resp.Id).Execute()
		}()
	}

	require.NoError(t, err, "Failed to create SD-WAN traffic distribution profile")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 201, httpResp.StatusCode, "Expected 201 Created status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testName, resp.Name, "Created profile name should match")
	t.Logf("[SUCCESS] Created SD-WAN traffic distribution profile: %s", resp.Name)
}

// Test_networkservices_SDWANTrafficDistributionProfilesAPIService_List tests listing SD-WAN traffic distribution profiles
func Test_networkservices_SDWANTrafficDistributionProfilesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// List SD-WAN traffic distribution profiles (no prerequisite needed for read-only list)
	listReq := client.SDWANTrafficDistributionProfilesAPI.ListSDWANTrafficDistributionProfiles(context.Background()).Folder("All").Limit(200).Offset(0)
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list SD-WAN traffic distribution profiles")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")
	t.Logf("[SUCCESS] Listed SD-WAN traffic distribution profiles, total: %d", len(resp.Data))
}

// Test_networkservices_SDWANTrafficDistributionProfilesAPIService_GetByID tests retrieving an SD-WAN traffic distribution profile by ID
func Test_networkservices_SDWANTrafficDistributionProfilesAPIService_GetByID(t *testing.T) {
	t.Skip("Requires sdwan_interface_profiles prerequisite which is not available in our OpenAPI spec - link tags must be associated with interface profiles first")
	client := SetupNetworkSvcTestClient(t)

	testName := "test-sdwan-tdp-get-" + common.GenerateRandomString(6)
	_ = testName
	_ = client
}

// Test_networkservices_SDWANTrafficDistributionProfilesAPIService_UpdateByID tests updating an SD-WAN traffic distribution profile
func Test_networkservices_SDWANTrafficDistributionProfilesAPIService_UpdateByID(t *testing.T) {
	t.Skip("Requires sdwan_interface_profiles prerequisite which is not available in our OpenAPI spec - link tags must be associated with interface profiles first")
	client := SetupNetworkSvcTestClient(t)

	testName := "test-sdwan-tdp-update-" + common.GenerateRandomString(6)
	_ = testName
	_ = client
}

// Test_networkservices_SDWANTrafficDistributionProfilesAPIService_DeleteByID tests deleting an SD-WAN traffic distribution profile
func Test_networkservices_SDWANTrafficDistributionProfilesAPIService_DeleteByID(t *testing.T) {
	t.Skip("Requires sdwan_interface_profiles prerequisite which is not available in our OpenAPI spec - link tags must be associated with interface profiles first")
	client := SetupNetworkSvcTestClient(t)

	testName := "test-sdwan-tdp-delete-" + common.GenerateRandomString(6)
	_ = testName
	_ = client
}

// Test_networkservices_SDWANTrafficDistributionProfilesAPIService_FetchSDWANTrafficDistributionProfiles tests the Fetch convenience method
func Test_networkservices_SDWANTrafficDistributionProfilesAPIService_FetchSDWANTrafficDistributionProfiles(t *testing.T) {
	t.Skip("Requires sdwan_interface_profiles prerequisite which is not available in our OpenAPI spec - link tags must be associated with interface profiles first")
	client := SetupNetworkSvcTestClient(t)

	testName := "test-sdwan-tdp-fetch-" + common.GenerateRandomString(6)
	_ = testName
	_ = client
}
