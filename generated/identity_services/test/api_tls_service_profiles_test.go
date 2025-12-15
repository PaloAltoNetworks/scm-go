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

func generateRandomName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, common.GenerateRandomString(6))
}

// createComplexTLSServiceProfile creates a simple complete TLS service profile payload
func createTLSServiceProfile(name string) identity_services.TlsServiceProfiles {

	protocolSettings := &identity_services.TlsServiceProfilesProtocolSettings{
		KeyxchgAlgoRsa: common.BoolPtr(true),
	}

	return identity_services.TlsServiceProfiles{
		Folder:           common.StringPtr("All"),
		Name:             name,
		Certificate:      "Forward-UnTrust-CA",
		ProtocolSettings: *protocolSettings,
	}
}

// createComplexTLSServiceProfile creates a complex complete TLS service profile payload
func createComplexTLSServiceProfile(name string) identity_services.TlsServiceProfiles {

	// define the secure protocol settings
	protocolSettings := &identity_services.TlsServiceProfilesProtocolSettings{
		MinVersion: common.StringPtr("tls1-1"),
		MaxVersion: common.StringPtr("tls1-3"),

		KeyxchgAlgoRsa:   common.BoolPtr(true),
		KeyxchgAlgoEcdhe: common.BoolPtr(true),
		KeyxchgAlgoDhe:   common.BoolPtr(true),

		EncAlgoAes128Gcm: common.BoolPtr(true),
		EncAlgoAes256Gcm: common.BoolPtr(true),
		EncAlgoAes128Cbc: common.BoolPtr(false),
		EncAlgoAes256Cbc: common.BoolPtr(true),

		AuthAlgoSha256: common.BoolPtr(true),
		AuthAlgoSha384: common.BoolPtr(false),
	}

	return identity_services.TlsServiceProfiles{
		Folder:           common.StringPtr("All"),
		Name:             name,
		Certificate:      "Forward-Trust-CA",
		ProtocolSettings: *protocolSettings,
	}
}

// Test_identity_services_TLSServiceProfilesAPIService_Create tests the creation of a TLS Service Profile.
func Test_identity_services_TLSServiceProfilesAPIService_Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	createdProfileName := generateRandomName("tls-prof-create")
	profilePayload := createComplexTLSServiceProfile(createdProfileName)

	fmt.Printf("Creating TLS Service Profile with name: %s\n", profilePayload.Name)
	req := client.TLSServiceProfilesAPI.CreateTLSServiceProfiles(context.Background()).TlsServiceProfiles(profilePayload)
	res, httpRes, err := req.Execute()

	// assertions and cleanup
	require.NoError(t, err, "Failed to create TLS Service Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")

	createdProfileID := res.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up TLS Service Profile with ID: %s", createdProfileID)
		_, errDel := client.TLSServiceProfilesAPI.DeleteTLSServiceProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete TLS Service Profile during cleanup")
	}()

	t.Logf("Successfully created TLS Service Profile: %s with ID: %s", createdProfileName, createdProfileID)
}

// Test_identity_services_TLSServiceProfilesAPIService_GetByID tests retrieving a TLS Service Profile by its ID.
func Test_identity_services_TLSServiceProfilesAPIService_GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	profileName := generateRandomName("test-tls-prof-get")
	profilePayload := createTLSServiceProfile(profileName)

	createRes, _, err := client.TLSServiceProfilesAPI.CreateTLSServiceProfiles(context.Background()).TlsServiceProfiles(profilePayload).Execute()
	require.NoError(t, err, "Failed to create TLS Service Profile for get test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// cleanup
	defer func() {
		t.Logf("Cleaning up TLS Service Profile with ID: %s", createdProfileID)
		_, errDel := client.TLSServiceProfilesAPI.DeleteTLSServiceProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete TLS Service Profile during cleanup")
	}()

	getRes, httpResGet, errGet := client.TLSServiceProfilesAPI.GetTLSServiceProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get TLS Service Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
}

// Test_identity_services_TLSServiceProfilesAPIService_Update tests updating an existing TLS Service Profile.
func Test_identity_services_TLSServiceProfilesAPIService_Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	profileName := generateRandomName("test-tls-profile-update")
	profilePayload := createTLSServiceProfile(profileName)

	createRes, _, err := client.TLSServiceProfilesAPI.CreateTLSServiceProfiles(context.Background()).TlsServiceProfiles(profilePayload).Execute()
	require.NoError(t, err, "Failed to create TLS Service Profile for update test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// cleanup
	defer func() {
		_, errDel := client.TLSServiceProfilesAPI.DeleteTLSServiceProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete TLS Service Profile during cleanup")
	}()

	// creating updated object
	updatedProtocolSettings := &identity_services.TlsServiceProfilesProtocolSettings{
		MinVersion: common.StringPtr("tls1-0"),
		MaxVersion: common.StringPtr("tls1-2"),

		KeyxchgAlgoRsa:   common.BoolPtr(true),
		KeyxchgAlgoEcdhe: common.BoolPtr(true),
		KeyxchgAlgoDhe:   common.BoolPtr(true),

		AuthAlgoSha256: common.BoolPtr(true),
		AuthAlgoSha384: common.BoolPtr(false),
	}

	updatedProfile := identity_services.TlsServiceProfiles{
		Folder:           common.StringPtr("All"),
		Name:             createRes.Name,
		Certificate:      "Forward-UnTrust-CA",
		ProtocolSettings: *updatedProtocolSettings,
	}

	// applying the update
	updateRes, httpResUpdate, errUpdate := client.TLSServiceProfilesAPI.UpdateTLSServiceProfilesByID(context.Background(), createdProfileID).TlsServiceProfiles(updatedProfile).Execute()

	// assertions
	require.NoError(t, errUpdate, "Failed to update TLS Service Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// assert the specific field that was updated.
	require.NotNil(t, updateRes.ProtocolSettings, "Updated ProtocolSettings should not be nil")
	assert.Equal(t, updatedProtocolSettings.MinVersion, common.StringPtr(*updateRes.ProtocolSettings.MinVersion), "Min Version should be updated to tls1-0")
	assert.Equal(t, updatedProtocolSettings.MaxVersion, common.StringPtr(*updateRes.ProtocolSettings.MaxVersion), "Min Version should be updated to tls1-3")
}

// Test_identity_services_TLSServiceProfilesAPIService_List tests listing TLS Service Profiles.
func Test_identity_services_TLSServiceProfilesAPIService_List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Create a resource to ensure the list is non-empty
	profileName := generateRandomName("test-tls-profile-list-item")
	profilePayload := createTLSServiceProfile(profileName)

	createRes, _, err := client.TLSServiceProfilesAPI.CreateTLSServiceProfiles(context.Background()).TlsServiceProfiles(profilePayload).Execute()
	require.NoError(t, err, "Failed to create TLS Service Profile for list test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up TLS Service Profile with ID: %s", createdProfileID)
		_, errDel := client.TLSServiceProfilesAPI.DeleteTLSServiceProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete TLS Service Profile during cleanup")
	}()

	// LIST Request, typically filtered to limit results
	listRes, httpResList, errList := client.TLSServiceProfilesAPI.ListTLSServiceProfiles(context.Background()).Folder("Shared").Execute()
	require.NoError(t, errList, "Failed to list TLS Service Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	// Verify the created object is in the list
	foundObject := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created TLS Service Profile should be found in the list")
}

// Test_identity_services_TLSServiceProfilesAPIService_DeleteByID tests deleting a TLS Service Profile.
func Test_identity_services_TLSServiceProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	profileName := generateRandomName("test-tls-profile-delete")
	profilePayload := createTLSServiceProfile(profileName)

	createRes, _, err := client.TLSServiceProfilesAPI.CreateTLSServiceProfiles(context.Background()).TlsServiceProfiles(profilePayload).Execute()
	require.NoError(t, err, "Failed to create TLS Service Profile for delete test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// DELETE Request
	httpResDel, errDel := client.TLSServiceProfilesAPI.DeleteTLSServiceProfilesByID(context.Background(), createdProfileID).Execute()

	require.NoError(t, errDel, "Failed to delete TLS Service Profile")
	// The API typically returns 200 OK or 204 No Content for successful deletion
	assert.True(t, httpResDel.StatusCode == 200 || httpResDel.StatusCode == 204, "Expected 200 or 204 status for delete")

	t.Logf("Successfully deleted TLS Service Profile with ID: %s", createdProfileID)
}
