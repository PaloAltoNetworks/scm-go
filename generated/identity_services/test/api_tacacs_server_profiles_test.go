/*
 * Identity Services Testing
 *
 * TACACSServerProfilesAPIService
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

// Test_identity_services_TACACSServerProfilesAPIService_Create tests the creation of a TACACS Server Profile.
func Test_identity_services_TACACSServerProfilesAPIService_Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	createdProfileName := "test-tacacs-create-" + common.GenerateRandomString(6)

	// define servers
	testServers := []identity_services.TacacsServerProfilesServerInner{
		{
			Name:    common.StringPtr("tacacs-server-1"),
			Address: common.StringPtr("200.5.5.100"),
			Port:    common.Int32Ptr(20),
			Secret:  common.StringPtr("a"),
		},
		{
			Name:    common.StringPtr("tacacs-server-2"),
			Address: common.StringPtr("100.2.120.50"),
			Port:    common.Int32Ptr(1255),
			Secret:  common.StringPtr("secret"),
		},
		{
			Name:    common.StringPtr("tacacs-server-3"),
			Address: common.StringPtr("address_3"),
			Port:    common.Int32Ptr(40000),
			Secret:  common.StringPtr("68#67p!Z7mR8*ql1XwN8@b04yV0f83sJ6hA9%uC2775&dP8xhoK4*jQ7tW0zS3rK"),
		},
	}

	// define a TACACS server profile
	profile := identity_services.TacacsServerProfiles{
		Folder:              common.StringPtr("All"),
		Name:                createdProfileName,
		Protocol:            "CHAP",
		Server:              testServers,
		Timeout:             common.Int32Ptr(15),
		UseSingleConnection: common.BoolPtr(true),
	}

	fmt.Printf("Creating TACACS Server Profile with name: %s\n", profile.Name)
	req := client.TACACSServerProfilesAPI.CreateTACACSServerProfiles(context.Background()).TacacsServerProfiles(profile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create TACACS Server Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")

	createdProfileID := res.Id

	defer func() {
		t.Logf("Cleaning up TACACS Server Profile with ID: %s", createdProfileID)
		_, errDel := client.TACACSServerProfilesAPI.DeleteTACACSServerProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete TACACS Server Profile during cleanup")
	}()

	t.Logf("Successfully created TACACS Server Profile: %s with ID: %s", profile.Name, createdProfileID)
	assert.Equal(t, "All", *res.Folder, "Folder should match")
	assert.Equal(t, "CHAP", res.Protocol, "Protocol should match")
	assert.Len(t, res.Server, 3, "Should have 3 servers configured")
}

// Test_identity_services_TACACSServerProfilesAPIService_GetByID tests retrieving a TACACS Server Profile by its ID.
func Test_identity_services_TACACSServerProfilesAPIService_GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-tacacs-get-" + common.GenerateRandomString(6)

	profile := identity_services.TacacsServerProfiles{
		Folder:   common.StringPtr("All"),
		Name:     profileName,
		Protocol: "PAP",
		Server: []identity_services.TacacsServerProfilesServerInner{
			{
				Name:    common.StringPtr("get-test-svr"),
				Address: common.StringPtr("192.168.1.10"),
				Secret:  common.StringPtr("testSecret123"),
			},
		},
	}

	createRes, _, err := client.TACACSServerProfilesAPI.CreateTACACSServerProfiles(context.Background()).TacacsServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create TACACS Server Profile for get test")
	createdProfileID := createRes.Id

	defer func() {
		t.Logf("Cleaning up TACACS Server Profile with ID: %s", createdProfileID)
		client.TACACSServerProfilesAPI.DeleteTACACSServerProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	getRes, httpResGet, errGet := client.TACACSServerProfilesAPI.GetTACACSServerProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get TACACS Server Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
}

// Test_identity_services_TACACSServerProfilesAPIService_Update tests updating an existing TACACS Server Profile.
func Test_identity_services_TACACSServerProfilesAPIService_Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-tacacs-update-" + common.GenerateRandomString(6)

	profile := identity_services.TacacsServerProfiles{
		Folder:   common.StringPtr("All"),
		Name:     profileName,
		Protocol: "CHAP",
		Server: []identity_services.TacacsServerProfilesServerInner{
			{
				Name:    common.StringPtr("orig-svr"),
				Address: common.StringPtr("1.1.1.1"),
				Secret:  common.StringPtr("origSecret"),
			},
		},
	}

	createRes, _, err := client.TACACSServerProfilesAPI.CreateTACACSServerProfiles(context.Background()).TacacsServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create TACACS Server Profile for update test")
	createdProfileID := createRes.Id

	defer func() {
		t.Logf("Cleaning up TACACS Server Profile with ID: %s", createdProfileID)
		client.TACACSServerProfilesAPI.DeleteTACACSServerProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	updatedTimeout := int32(20)
	updatedProfile := identity_services.TacacsServerProfiles{
		Name:     profileName,
		Protocol: "PAP",                           // changed protocol
		Timeout:  common.Int32Ptr(updatedTimeout), // added timeout
		Server:   profile.Server,
	}

	updateRes, httpResUpdate, errUpdate := client.TACACSServerProfilesAPI.UpdateTACACSServerProfilesByID(context.Background(), createdProfileID).TacacsServerProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update TACACS Server Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, updatedTimeout, *updateRes.Timeout, "Timeout should be updated")
}

// Test_identity_services_TACACSServerProfilesAPIService_List tests listing TACACS Server Profiles.
func Test_identity_services_TACACSServerProfilesAPIService_List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-tacacs-list-" + common.GenerateRandomString(6)

	profile := identity_services.TacacsServerProfiles{
		Folder:   common.StringPtr("Shared"),
		Name:     profileName,
		Protocol: "CHAP",
		Server: []identity_services.TacacsServerProfilesServerInner{
			{
				Name:    common.StringPtr("list-test-svr"),
				Address: common.StringPtr("2.2.2.2"),
				Secret:  common.StringPtr("listSecret"),
			},
		},
	}

	createRes, _, err := client.TACACSServerProfilesAPI.CreateTACACSServerProfiles(context.Background()).TacacsServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create TACACS Server Profile for list test")
	createdProfileID := createRes.Id

	defer func() {
		t.Logf("Cleaning up TACACS Server Profile with ID: %s", createdProfileID)
		client.TACACSServerProfilesAPI.DeleteTACACSServerProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	listRes, httpResList, errList := client.TACACSServerProfilesAPI.ListTACACSServerProfiles(context.Background()).Folder("Shared").Limit(200).Execute()
	require.NoError(t, errList, "Failed to list TACACS Server Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	foundObject := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created TACACS Server Profile should be found in the list")
}

// Test_identity_services_TACACSServerProfilesAPIService_DeleteByID tests deleting a TACACS Server Profile by its ID.
func Test_identity_services_TACACSServerProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)
	profileName := "test-tacacs-delete-" + common.GenerateRandomString(6)

	profile := identity_services.TacacsServerProfiles{
		Folder:   common.StringPtr("Shared"),
		Name:     profileName,
		Protocol: "CHAP",
		Server: []identity_services.TacacsServerProfilesServerInner{
			{
				Name:    common.StringPtr("del-test-svr"),
				Address: common.StringPtr("3.3.3.3"),
				Secret:  common.StringPtr("delSecret"),
			},
		},
	}

	createRes, _, err := client.TACACSServerProfilesAPI.CreateTACACSServerProfiles(context.Background()).TacacsServerProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create TACACS Server Profile for delete test")
	createdProfileID := createRes.Id

	httpResDel, errDel := client.TACACSServerProfilesAPI.DeleteTACACSServerProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete TACACS Server Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}
