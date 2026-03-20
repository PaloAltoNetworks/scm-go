/*
Identity Services Testing LDAPServerProfilesAPIService
*/
package identity_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// Test_identity_services_LDAPServerProfilesAPIService_Create tests the creation of a complex LDAP Server Profile.
func Test_identity_services_LDAPServerProfilesAPIService_Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	createdProfileName := "test-ldap-create-" + common.GenerateRandomString(6)

	// define multiple LDAP servers for a complex setup
	testServers := []identity_services.LdapServerProfilesServerInner{
		{
			Name:    common.StringPtr("primary-ldap"),
			Address: common.StringPtr("$tst_68081_1"),
			Port:    common.Int32Ptr(636),
		},
		{
			Name:    common.StringPtr("secondary-ldap"),
			Address: common.StringPtr("$tst_68081_2"),
			Port:    common.Int32Ptr(636),
		},
	}

	// create the LDAP server profile
	profile := identity_services.LdapServerProfiles{
		Folder:                  common.StringPtr("All"),
		Name:                    createdProfileName,
		LdapType:                common.StringPtr("active-directory"),
		Base:                    common.StringPtr("DC=internal,DC=example,DC=com"),
		BindDn:                  common.StringPtr("CN=LDAP Bind,OU=Service Accounts,DC=internal,DC=example,DC=com"),
		BindPassword:            common.StringPtr("Password123!"),
		BindTimelimit:           common.StringPtr("15"),
		RetryInterval:           common.Int32Ptr(60),
		Ssl:                     common.BoolPtr(true),
		Timelimit:               common.Int32Ptr(30),
		VerifyServerCertificate: common.BoolPtr(true),
		Server:                  testServers,
	}

	req := client.LDAPServerProfilesAPI.CreateLDAPServerProfiles(context.Background()).LdapServerProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create LDAP Server Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")

	createdProfileID := res.Id

	defer func() {
		t.Logf("Cleaning up LDAP Server Profile with ID: %s", createdProfileID)
		_, errDel := client.LDAPServerProfilesAPI.DeleteLDAPServerProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete LDAP Server Profile during cleanup")
	}()

	t.Logf("Successfully created LDAP Server Profile: %s with ID: %s", profile.Name, createdProfileID)
	assert.Equal(t, "All", *res.Folder, "Folder should match")
	assert.Equal(t, "active-directory", *res.LdapType, "LDAP type should match")
	assert.Len(t, res.Server, 2, "Should have 2 servers configured")
	assert.True(t, *res.Ssl, "SSL should be enabled")
}

// Test_identity_services_LDAPServerProfilesAPIService_GetByID tests retrieving a basic LDAP Server Profile by ID.
func Test_identity_services_LDAPServerProfilesAPIService_GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-ldap-get-" + common.GenerateRandomString(6)

	profile := identity_services.LdapServerProfiles{
		Folder: common.StringPtr("All"),
		Name:   profileName,
		Server: []identity_services.LdapServerProfilesServerInner{
			{
				Name:    common.StringPtr("get-test-svr"),
				Address: common.StringPtr("192.168.10.10"),
				Port:    common.Int32Ptr(140),
			},
		},
	}

	createRes, _, err := client.LDAPServerProfilesAPI.CreateLDAPServerProfiles(context.Background()).LdapServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create LDAP Server Profile for get test")
	createdProfileID := createRes.Id

	defer func() {
		client.LDAPServerProfilesAPI.DeleteLDAPServerProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	getRes, httpResGet, errGet := client.LDAPServerProfilesAPI.GetLDAPServerProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get LDAP Server Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
}

// Test_identity_services_LDAPServerProfilesAPIService_Update tests updating a basic LDAP Server Profile.
func Test_identity_services_LDAPServerProfilesAPIService_Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-ldap-update-" + common.GenerateRandomString(6)

	profile := identity_services.LdapServerProfiles{
		Folder: common.StringPtr("All"),
		Name:   profileName,
		Server: []identity_services.LdapServerProfilesServerInner{
			{
				Name:    common.StringPtr("orig-svr"),
				Address: common.StringPtr("1.1.1.1"),
				Port:    common.Int32Ptr(200),
			},
		},
	}

	createRes, _, err := client.LDAPServerProfilesAPI.CreateLDAPServerProfiles(context.Background()).LdapServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create LDAP Server Profile for update test")
	createdProfileID := createRes.Id

	defer func() {
		client.LDAPServerProfilesAPI.DeleteLDAPServerProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	// update the server list and add a timeout
	updatedServers := []identity_services.LdapServerProfilesServerInner{
		{
			Name:    common.StringPtr("updated-test-svr-1"),
			Address: common.StringPtr("2.2.2.2"),
			Port:    common.Int32Ptr(636),
		},
		{
			Name:    common.StringPtr("updated-test-svr-2"),
			Address: common.StringPtr("192.10.20.115"),
			Port:    common.Int32Ptr(6360),
		},
	}

	updatedProfile := identity_services.LdapServerProfiles{
		Id:        createdProfileID,
		Name:      profileName,
		Folder:    common.StringPtr("All"),
		Timelimit: common.Int32Ptr(30),
		Server:    updatedServers,
	}

	updateRes, httpResUpdate, errUpdate := client.LDAPServerProfilesAPI.UpdateLDAPServerProfiles(context.Background(), createdProfileID).LdapServerProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update LDAP Server Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, "updated-test-svr-1", *updateRes.Server[0].Name, "Server name should be updated")
	assert.Equal(t, "updated-test-svr-2", *updateRes.Server[1].Name, "Server name should be added")
	assert.Equal(t, int32(636), *updateRes.Server[0].Port, "Port should be updated")
	assert.Equal(t, int32(30), *updateRes.Timelimit, "Timelimit should be updated")
}

// Test_identity_services_LDAPServerProfilesAPIService_List tests listing LDAP Server Profiles.
func Test_identity_services_LDAPServerProfilesAPIService_List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-ldap-list-" + common.GenerateRandomString(6)

	profile := identity_services.LdapServerProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		Server: []identity_services.LdapServerProfilesServerInner{
			{
				Name:    common.StringPtr("list-svr"),
				Address: common.StringPtr("3.3.3.3"),
				Port:    common.Int32Ptr(39900),
			},
		},
	}

	createRes, _, err := client.LDAPServerProfilesAPI.CreateLDAPServerProfiles(context.Background()).LdapServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create LDAP Server Profile for list test")
	createdProfileID := createRes.Id

	defer func() {
		client.LDAPServerProfilesAPI.DeleteLDAPServerProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	listRes, httpResList, errList := client.LDAPServerProfilesAPI.ListLDAPServerProfiles(context.Background()).Folder("Shared").Limit(200).Execute()
	require.NoError(t, errList, "Failed to list LDAP Server Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	found := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			found = true
			break
		}
	}
	assert.True(t, found, "Created LDAP Server Profile should be present in the list")
}

// Test_identity_services_LDAPServerProfilesAPIService_DeleteByID tests deleting an LDAP Server Profile.
func Test_identity_services_LDAPServerProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-ldap-delete-" + common.GenerateRandomString(6)

	profile := identity_services.LdapServerProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		Server: []identity_services.LdapServerProfilesServerInner{
			{
				Name:    common.StringPtr("del-svr"),
				Address: common.StringPtr("4.4.4.4"),
				Port:    common.Int32Ptr(24000),
			},
		},
	}

	createRes, _, err := client.LDAPServerProfilesAPI.CreateLDAPServerProfiles(context.Background()).LdapServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create LDAP Server Profile for delete test")
	createdProfileID := createRes.Id

	httpResDel, errDel := client.LDAPServerProfilesAPI.DeleteLDAPServerProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete LDAP Server Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_identity_services_LDAPServerProfilesAPIService_FetchLDAPServerProfiles tests the FetchLDAPServerProfiles convenience method
func Test_identity_services_LDAPServerProfilesAPIService_FetchLDAPServerProfiles(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Create a test object first
	testName := "fetch-ldap-" + common.GenerateRandomString(6)
	testObj := identity_services.LdapServerProfiles{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
		Server: []identity_services.LdapServerProfilesServerInner{
			{
				Name:    common.StringPtr("ldap-server-1"),
				Address: common.StringPtr("192.168.1.1"),
				Port:    common.Int32Ptr(389),
			},
		},
	}

	createReq := client.LDAPServerProfilesAPI.CreateLDAPServerProfiles(context.Background()).LdapServerProfiles(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.LDAPServerProfilesAPI.DeleteLDAPServerProfilesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.LDAPServerProfilesAPI.FetchLDAPServerProfiles(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch ldap server profiles by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchLDAPServerProfiles found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.LDAPServerProfilesAPI.FetchLDAPServerProfiles(
		context.Background(),
		"non-existent-ldap-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchLDAPServerProfiles correctly returned nil for non-existent object")
}
