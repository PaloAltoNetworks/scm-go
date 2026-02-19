/*
Identity Services Testing

MFAServersAPIService
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

// Test_identity_services_MFAServersAPIService_Create tests the creation of an MFA server.
func Test_identity_services_MFAServersAPIService_Create(t *testing.T) {
	t.Skip("API returns 500 Internal Server Error - MFA server Create not supported in test environment")
	client := SetupIdentitySvcTestClient(t)
	createdName := "test-mfa-create-" + common.GenerateRandomString(6)

	// define the MFA server
	mfaServer := identity_services.MfaServers{
		Folder:         common.StringPtr("All"),
		Name:           createdName,
		MfaCertProfile: "Default",
	}

	fmt.Printf("Creating MFA Server with name: %s\n", mfaServer.Name)
	req := client.MFAServersAPI.CreateMFAServers(context.Background()).MfaServers(mfaServer)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create MFA Server")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdName, res.Name, "Created MFA server name should match")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up MFA Server with ID: %s", createdID)
		_, errDel := client.MFAServersAPI.DeleteMFAServersByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete MFA Server during cleanup")
	}()

	t.Logf("Successfully created MFA Server: %s with ID: %s", mfaServer.Name, createdID)
	assert.Equal(t, "All", *res.Folder, "Folder should match")
	assert.Equal(t, "Default", res.MfaCertProfile, "MFA cert profile should match")
}

// Test_identity_services_MFAServersAPIService_GetByID tests retrieving an MFA server by ID.
func Test_identity_services_MFAServersAPIService_GetByID(t *testing.T) {
	t.Skip("API returns 500 Internal Server Error - MFA server Create not supported in test environment")
	client := SetupIdentitySvcTestClient(t)
	mfaName := "test-mfa-get-" + common.GenerateRandomString(6)

	mfaServer := identity_services.MfaServers{
		Folder:         common.StringPtr("All"),
		Name:           mfaName,
		MfaCertProfile: "Default",
	}

	createRes, _, err := client.MFAServersAPI.CreateMFAServers(context.Background()).MfaServers(mfaServer).Execute()
	require.NoError(t, err, "Failed to create MFA Server for get test")
	createdID := *createRes.Id

	defer func() {
		client.MFAServersAPI.DeleteMFAServersByID(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.MFAServersAPI.GetMFAServersByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get MFA Server by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, mfaName, getRes.Name, "MFA server name should match")
}

// Test_identity_services_MFAServersAPIService_Update tests updating an existing MFA server.
func Test_identity_services_MFAServersAPIService_Update(t *testing.T) {
	t.Skip("API returns 500 Internal Server Error - MFA server Create not supported in test environment")
	client := SetupIdentitySvcTestClient(t)
	mfaName := "test-mfa-update-" + common.GenerateRandomString(6)

	mfaServer := identity_services.MfaServers{
		Folder:         common.StringPtr("All"),
		Name:           mfaName,
		MfaCertProfile: "Default",
	}

	createRes, _, err := client.MFAServersAPI.CreateMFAServers(context.Background()).MfaServers(mfaServer).Execute()
	require.NoError(t, err, "Failed to create MFA Server for update test")
	createdID := *createRes.Id

	defer func() {
		client.MFAServersAPI.DeleteMFAServersByID(context.Background(), createdID).Execute()
	}()

	// update the MFA cert profile
	updatedServer := identity_services.MfaServers{
		Name:           mfaName,
		MfaCertProfile: "UpdatedProfile",
	}

	updateRes, httpResUpdate, errUpdate := client.MFAServersAPI.UpdateMFAServersByID(context.Background(), createdID).MfaServers(updatedServer).Execute()
	require.NoError(t, errUpdate, "Failed to update MFA Server")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, "UpdatedProfile", updateRes.MfaCertProfile, "MFA cert profile should be updated")
}

// Test_identity_services_MFAServersAPIService_List tests listing MFA servers.
func Test_identity_services_MFAServersAPIService_List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Read-only test: list existing MFA servers
	listRes, httpResList, errList := client.MFAServersAPI.ListMFAServers(context.Background()).Folder("All").Limit(200).Position("pre").Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list MFA servers")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed MFA servers, total: %d", listRes.GetTotal())
}

// Test_identity_services_MFAServersAPIService_DeleteByID tests deleting an MFA server.
func Test_identity_services_MFAServersAPIService_DeleteByID(t *testing.T) {
	t.Skip("API returns 500 Internal Server Error - MFA server Create not supported in test environment")
	client := SetupIdentitySvcTestClient(t)
	mfaName := "test-mfa-delete-" + common.GenerateRandomString(6)

	mfaServer := identity_services.MfaServers{
		Folder:         common.StringPtr("Shared"),
		Name:           mfaName,
		MfaCertProfile: "Default",
	}

	createRes, _, err := client.MFAServersAPI.CreateMFAServers(context.Background()).MfaServers(mfaServer).Execute()
	require.NoError(t, err, "Failed to create MFA Server for delete test")
	createdID := *createRes.Id

	httpResDel, errDel := client.MFAServersAPI.DeleteMFAServersByID(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete MFA Server")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_identity_services_MFAServersAPIService_FetchMFAServers tests the FetchMFAServers convenience method
func Test_identity_services_MFAServersAPIService_FetchMFAServers(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Read-only not-found test (Create gives 500)
	notFound, err := client.MFAServersAPI.FetchMFAServers(
		context.Background(),
		"non-existent-mfa-server-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchMFAServers correctly returned nil for non-existent object")
}
