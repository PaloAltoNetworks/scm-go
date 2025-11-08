/*
 * Security Services Testing
 *
 * DNSSecurityProfilesAPIService
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

// Helper function to define a simple DNS Security Profile structure for reuse.
func newSimpleTestDnsSecurityProfile(name string) security_services.DnsSecurityProfiles {
	return security_services.DnsSecurityProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   common.StringPtr(name),
	}
}

// Helper function to define a simple DNS Security Profile structure for reuse.
func newCreateTestDnsSecurityProfile(name string) security_services.DnsSecurityProfiles {
	dnsSecCategory_1 := security_services.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner{
		Name:          common.StringPtr("pan-dns-sec-parked"),
		Action:        common.StringPtr("allow"),
		LogLevel:      common.StringPtr("informational"),
		PacketCapture: common.StringPtr("disable"),
	}

	dnsSecCategory_2 := security_services.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner{
		Name:          common.StringPtr("pan-dns-sec-cc"),
		Action:        common.StringPtr("block"),
		LogLevel:      common.StringPtr("low"),
		PacketCapture: common.StringPtr("extended-capture"),
	}

	listEntry := security_services.DnsSecurityProfilesBotnetDomainsListsInner{
		Name: "default-paloalto-dns",
		Action: &security_services.DnsSecurityProfilesBotnetDomainsListsInnerAction{
			Block: map[string]interface{}{},
		},
		PacketCapture: common.StringPtr("disable"),
	}

	sinkholeConfig := security_services.DnsSecurityProfilesBotnetDomainsSinkhole{
		Ipv4Address: common.StringPtr("127.0.0.1"),
		Ipv6Address: common.StringPtr("::1"),
	}

	whitelistEntry_1 := security_services.DnsSecurityProfilesBotnetDomainsWhitelistInner{
		Name:        "example.com",
		Description: common.StringPtr("creating whitelist example"),
	}

	whitelistEntry_2 := security_services.DnsSecurityProfilesBotnetDomainsWhitelistInner{
		Name:        "target.com",
		Description: common.StringPtr("creating another whitelist example"),
	}

	botnetDomains := security_services.DnsSecurityProfilesBotnetDomains{
		DnsSecurityCategories: []security_services.DnsSecurityProfilesBotnetDomainsDnsSecurityCategoriesInner{dnsSecCategory_1, dnsSecCategory_2},
		Lists:                 []security_services.DnsSecurityProfilesBotnetDomainsListsInner{listEntry},
		Sinkhole:              &sinkholeConfig,
		Whitelist:             []security_services.DnsSecurityProfilesBotnetDomainsWhitelistInner{whitelistEntry_1, whitelistEntry_2},
	}

	return security_services.DnsSecurityProfiles{
		Folder:        common.StringPtr("Shared"),
		Name:          common.StringPtr(name),
		BotnetDomains: &botnetDomains,
	}
}

// Test_security_services_DNSSecurityProfilesAPIService_Create tests the creation of a DNS Security Profile.
func Test_security_services_DNSSecurityProfilesAPIService_Create(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdProfileName := "test-dns-create-" + common.GenerateRandomString(6)

	// define a DNS security profile
	profile := newCreateTestDnsSecurityProfile(createdProfileName)

	fmt.Printf("Creating DNS Security Profile with name: %s\n", *profile.Name)
	req := client.DNSSecurityProfilesAPI.CreateDNSSecurityProfiles(context.Background()).DnsSecurityProfiles(profile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create DNS Security Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, *res.Name, "Created profile name should match")
	createdProfileID := res.Id
	require.NotNil(t, createdProfileID, "Created profile ID should not be nil")

	defer func() {
		t.Logf("Cleaning up DNS Security Profile with ID: %s", *createdProfileID)
		_, errDel := client.DNSSecurityProfilesAPI.DeleteDNSSecurityProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete DNS Security Profile during cleanup")
	}()

	t.Logf("Successfully created DNS Security Profile: %s with ID: %s", *profile.Name, *createdProfileID)
	assert.Equal(t, "Shared", *res.Folder, "Folder should match")
}

// Test_security_services_DNSSecurityProfilesAPIService_GetByID tests retrieving a DNS Security Profile by its ID.
func Test_security_services_DNSSecurityProfilesAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-dns-get-" + common.GenerateRandomString(6)
	profile := newSimpleTestDnsSecurityProfile(profileName)

	createRes, _, err := client.DNSSecurityProfilesAPI.CreateDNSSecurityProfiles(context.Background()).DnsSecurityProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create DNS Security Profile for get test")
	createdProfileID := createRes.Id
	require.NotNil(t, createdProfileID, "Created profile ID should not be nil")

	defer func() {
		t.Logf("Cleaning up DNS Security Profile with ID: %s", *createdProfileID)
		_, errDel := client.DNSSecurityProfilesAPI.DeleteDNSSecurityProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete DNS Security Profile during cleanup")
	}()

	getRes, httpResGet, errGet := client.DNSSecurityProfilesAPI.GetDNSSecurityProfilesByID(context.Background(), *createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get DNS Security Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, *getRes.Name, "Profile name should match")
}

// Test_security_services_DNSSecurityProfilesAPIService_Update tests updating an existing DNS Security Profile.
func Test_security_services_DNSSecurityProfilesAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-dns-update-" + common.GenerateRandomString(6)
	profile := newSimpleTestDnsSecurityProfile(profileName)

	createRes, _, err := client.DNSSecurityProfilesAPI.CreateDNSSecurityProfiles(context.Background()).DnsSecurityProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create DNS Security Profile for update test")
	createdProfileID := createRes.Id
	require.NotNil(t, createdProfileID, "Created profile ID should not be nil")

	defer func() {
		t.Logf("Cleaning up DNS Security Profile with ID: %s", *createdProfileID)
		_, errDel := client.DNSSecurityProfilesAPI.DeleteDNSSecurityProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete DNS Security Profile during cleanup")
	}()

	// define the update: change the description
	updatedDescription := "here is the updated description."
	updatedProfile := security_services.DnsSecurityProfiles{
		Name:        common.StringPtr(profileName), // name must match
		Folder:      common.StringPtr("Shared"),
		Description: common.StringPtr(updatedDescription),
	}

	updateRes, httpResUpdate, errUpdate := client.DNSSecurityProfilesAPI.UpdateDNSSecurityProfilesByID(context.Background(), *createdProfileID).DnsSecurityProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update DNS Security Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// the updated profile should now have the new description
	assert.True(t, updateRes.HasDescription(), "Updated profile should have a Description set")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Updated description should match")
}

// Test_security_services_DNSSecurityProfilesAPIService_List tests listing DNS Security Profiles.
func Test_security_services_DNSSecurityProfilesAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-dns-list-" + common.GenerateRandomString(6)
	profile := newSimpleTestDnsSecurityProfile(profileName)

	createRes, _, err := client.DNSSecurityProfilesAPI.CreateDNSSecurityProfiles(context.Background()).DnsSecurityProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create DNS Security Profile for list test")
	createdProfileID := createRes.Id
	require.NotNil(t, createdProfileID, "Created profile ID should not be nil")

	defer func() {
		t.Logf("Cleaning up DNS Security Profile with ID: %s", *createdProfileID)
		_, errDel := client.DNSSecurityProfilesAPI.DeleteDNSSecurityProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete DNS Security Profile during cleanup")
	}()

	// list request will use query parameters to filter (Folder and Limit)
	listRes, httpResList, errList := client.DNSSecurityProfilesAPI.ListDNSSecurityProfiles(context.Background()).Folder("Shared").Limit(100).Execute()
	require.NoError(t, errList, "Failed to list DNS Security Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.NotNil(t, listRes.Data, "List response data should not be nil")

	foundObject := false
	for _, p := range listRes.Data {
		if p.Name != nil && *p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created DNS Security Profile should be found in the list")
}

// Test_security_services_DNSSecurityProfilesAPIService_DeleteByID tests deleting a DNS Security Profile by its ID.
func Test_security_services_DNSSecurityProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-dns-delete-" + common.GenerateRandomString(6)
	profile := newSimpleTestDnsSecurityProfile(profileName)

	createRes, _, err := client.DNSSecurityProfilesAPI.CreateDNSSecurityProfiles(context.Background()).DnsSecurityProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create DNS Security Profile for delete test")
	createdProfileID := createRes.Id
	require.NotNil(t, createdProfileID, "Created profile ID should not be nil")

	// delete the profile
	httpResDel, errDel := client.DNSSecurityProfilesAPI.DeleteDNSSecurityProfilesByID(context.Background(), *createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete DNS Security Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status on successful delete")
}
