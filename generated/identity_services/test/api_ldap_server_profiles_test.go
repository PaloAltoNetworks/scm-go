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
