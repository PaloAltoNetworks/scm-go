/*
Identity Services Testing

KerberosServerProfilesAPIService
*/

package identity_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// Test_identity_services_KerberosServerProfilesAPIService_Create tests the creation of a Kerberos Server Profile.
func Test_identity_services_KerberosServerProfilesAPIService_Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	createdProfileName := "test-kerberos-create-" + common.GenerateRandomString(6)

	// define kerberos servers
	testServers := []identity_services.KerberosServerProfilesServerInner{
		{
			Name: "kerb-server-1",
			Host: "10.0.1.50",
			Port: common.Int32Ptr(88),
		},
		{
			Name: "kerb-server-2",
			Host: "kerberos.example.com",
			Port: common.Int32Ptr(88),
		},
	}

	// define the kerberos server profile
	profile := identity_services.KerberosServerProfiles{
		Folder: common.StringPtr("All"),
		Name:   createdProfileName,
		Server: testServers,
	}

	fmt.Printf("Creating Kerberos Server Profile with name: %s\n", profile.Name)
	req := client.KerberosServerProfilesAPI.CreateKerberosServerProfiles(context.Background()).KerberosServerProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Kerberos Server Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")

	createdProfileID := res.Id

	defer func() {
		t.Logf("Cleaning up Kerberos Server Profile with ID: %s", createdProfileID)
		_, errDel := client.KerberosServerProfilesAPI.DeleteKerberosServerProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete Kerberos Server Profile during cleanup")
	}()

	t.Logf("Successfully created Kerberos Server Profile: %s with ID: %s", profile.Name, createdProfileID)
	assert.Equal(t, "All", *res.Folder, "Folder should match")
	assert.Len(t, res.Server, 2, "Should have 2 servers configured")
}

// Test_identity_services_KerberosServerProfilesAPIService_GetByID tests retrieving a Kerberos Server Profile by ID.
func Test_identity_services_KerberosServerProfilesAPIService_GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-kerberos-get-" + common.GenerateRandomString(6)

	profile := identity_services.KerberosServerProfiles{
		Folder: common.StringPtr("All"),
		Name:   profileName,
		Server: []identity_services.KerberosServerProfilesServerInner{
			{
				Name: "get-test-svr",
				Host: "192.168.10.10",
				Port: common.Int32Ptr(88),
			},
		},
	}

	createRes, _, err := client.KerberosServerProfilesAPI.CreateKerberosServerProfiles(context.Background()).KerberosServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Kerberos Server Profile for get test")
	createdProfileID := createRes.Id

	defer func() {
		client.KerberosServerProfilesAPI.DeleteKerberosServerProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	getRes, httpResGet, errGet := client.KerberosServerProfilesAPI.GetKerberosServerProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get Kerberos Server Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
}

// Test_identity_services_KerberosServerProfilesAPIService_Update tests updating an existing Kerberos Server Profile.
func Test_identity_services_KerberosServerProfilesAPIService_Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-kerberos-update-" + common.GenerateRandomString(6)

	profile := identity_services.KerberosServerProfiles{
		Folder: common.StringPtr("All"),
		Name:   profileName,
		Server: []identity_services.KerberosServerProfilesServerInner{
			{
				Name: "orig-svr",
				Host: "1.1.1.1",
				Port: common.Int32Ptr(88),
			},
		},
	}

	createRes, _, err := client.KerberosServerProfilesAPI.CreateKerberosServerProfiles(context.Background()).KerberosServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Kerberos Server Profile for update test")
	createdProfileID := createRes.Id

	defer func() {
		client.KerberosServerProfilesAPI.DeleteKerberosServerProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	// update the server list
	updatedServers := []identity_services.KerberosServerProfilesServerInner{
		{
			Name: "updated-test-svr-1",
			Host: "2.2.2.2",
			Port: common.Int32Ptr(8888),
		},
		{
			Name: "updated-test-svr-2",
			Host: "192.10.20.115",
			Port: common.Int32Ptr(10),
		},
	}

	updatedProfile := identity_services.KerberosServerProfiles{
		Name:   profileName,
		Server: updatedServers,
	}

	updateRes, httpResUpdate, errUpdate := client.KerberosServerProfilesAPI.UpdateKerberosServerProfilesByID(context.Background(), createdProfileID).KerberosServerProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update Kerberos Server Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, "updated-test-svr-1", updateRes.Server[0].Name, "Server name should be updated")
	assert.Equal(t, "updated-test-svr-2", updateRes.Server[1].Name, "Server name should be added")
	assert.Equal(t, int32(8888), *updateRes.Server[0].Port, "Port should be updated")
}

// Test_identity_services_KerberosServerProfilesAPIService_List tests listing Kerberos Server Profiles.
func Test_identity_services_KerberosServerProfilesAPIService_List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-kerberos-list-" + common.GenerateRandomString(6)

	profile := identity_services.KerberosServerProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		Server: []identity_services.KerberosServerProfilesServerInner{
			{
				Name: "list-svr",
				Host: "3.3.3.3",
			},
		},
	}

	createRes, _, err := client.KerberosServerProfilesAPI.CreateKerberosServerProfiles(context.Background()).KerberosServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Kerberos Server Profile for list test")
	createdProfileID := createRes.Id

	defer func() {
		client.KerberosServerProfilesAPI.DeleteKerberosServerProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	listRes, httpResList, errList := client.KerberosServerProfilesAPI.ListKerberosServerProfiles(context.Background()).Folder("Shared").Limit(200).Execute()
	require.NoError(t, errList, "Failed to list Kerberos Server Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	found := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			found = true
			break
		}
	}
	assert.True(t, found, "Created Kerberos Server Profile should be present in the list")
}

// Test_identity_services_KerberosServerProfilesAPIService_DeleteByID tests deleting a Kerberos Server Profile.
func Test_identity_services_KerberosServerProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-kerberos-delete-" + common.GenerateRandomString(6)

	profile := identity_services.KerberosServerProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		Server: []identity_services.KerberosServerProfilesServerInner{
			{
				Name: "del-svr",
				Host: "4.4.4.4",
			},
		},
	}

	createRes, _, err := client.KerberosServerProfilesAPI.CreateKerberosServerProfiles(context.Background()).KerberosServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Kerberos Server Profile for delete test")
	createdProfileID := createRes.Id

	httpResDel, errDel := client.KerberosServerProfilesAPI.DeleteKerberosServerProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete Kerberos Server Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_identity_services_KerberosServerProfilesAPIService_FetchKerberosServerProfiles tests the FetchKerberosServerProfiles convenience method
func Test_identity_services_KerberosServerProfilesAPIService_FetchKerberosServerProfiles(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Create a test object first
	testName := "fetch-kerb-" + common.GenerateRandomString(6)

	server := identity_services.KerberosServerProfilesServerInner{
		Name: "kerb-server-fetch",
		Host: "kdc.example.com",
		Port: common.Int32Ptr(88),
	}

	testObj := identity_services.KerberosServerProfiles{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
		Server: []identity_services.KerberosServerProfilesServerInner{server}, // Required field
	}

	createReq := client.KerberosServerProfilesAPI.CreateKerberosServerProfiles(context.Background()).KerberosServerProfiles(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.KerberosServerProfilesAPI.DeleteKerberosServerProfilesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.KerberosServerProfilesAPI.FetchKerberosServerProfiles(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch kerberos_server_profiles by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchKerberosServerProfiles found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.KerberosServerProfilesAPI.FetchKerberosServerProfiles(
		context.Background(),
		"non-existent-kerberos_server_profiles-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchKerberosServerProfiles correctly returned nil for non-existent object")
}
