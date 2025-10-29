/*
 * Security Services Testing
 *
 * DecryptionProfilesAPIService
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

// Test_security_services_DecryptionProfilesAPIService_Create tests the creation of a Decryption Profile.
func Test_security_services_DecryptionProfilesAPIService_Create(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdProfileName := "test-decrypt-create-" + common.GenerateRandomString(6)

	sslForwardProxySettings := security_services.DecryptionProfilesSslForwardProxy{
		BlockClientCert:               common.BoolPtr(true),
		BlockExpiredCertificate:       common.BoolPtr(true),
		BlockTimeoutCert:              common.BoolPtr(false),
		BlockTls13DowngradeNoResource: common.BoolPtr(true),
		BlockUnsupportedCipher:        common.BoolPtr(true),
		BlockUnsupportedVersion:       common.BoolPtr(true),
		StripAlpn:                     common.BoolPtr(true),
	}

	sslInboundProxySettings := security_services.DecryptionProfilesSslInboundProxy{
		BlockIfHsmUnavailable:   common.BoolPtr(true),
		BlockIfNoResource:       common.BoolPtr(false),
		BlockUnsupportedCipher:  common.BoolPtr(true),
		BlockUnsupportedVersion: common.BoolPtr(true),
	}

	sslNoProxySettings := security_services.DecryptionProfilesSslNoProxy{
		BlockExpiredCertificate: common.BoolPtr(true),
		BlockUntrustedIssuer:    common.BoolPtr(true),
	}

	sslProtocolSettings := security_services.DecryptionProfilesSslProtocolSettings{
		AuthAlgoSha256:          common.BoolPtr(true),
		AuthAlgoSha384:          common.BoolPtr(true),
		EncAlgoAes256Gcm:        common.BoolPtr(true),
		EncAlgoChacha20Poly1305: common.BoolPtr(true),
		KeyxchgAlgoEcdhe:        common.BoolPtr(true),
		MaxVersion:              common.StringPtr("tls1-3"),
		MinVersion:              common.StringPtr("tls1-2"),
	}

	// define a Decryption Profile
	profile := security_services.DecryptionProfiles{
		Folder:              common.StringPtr("Shared"),
		Name:                createdProfileName,
		SslForwardProxy:     &sslForwardProxySettings,
		SslInboundProxy:     &sslInboundProxySettings,
		SslNoProxy:          &sslNoProxySettings,
		SslProtocolSettings: &sslProtocolSettings,
	}

	fmt.Printf("Creating Decryption Profile with name: %s\n", profile.Name)
	req := client.DecryptionProfilesAPI.CreateDecryptionProfiles(context.Background()).DecryptionProfiles(profile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Decryption Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")
	createdProfileID := res.Id

	defer func() {
		t.Logf("Cleaning up Decryption Profile with ID: %s", *createdProfileID)
		_, errDel := client.DecryptionProfilesAPI.DeleteDecryptionProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete Decryption Profile during cleanup")
	}()

	t.Logf("Successfully created Decryption Profile: %s with ID: %s", profile.Name, *createdProfileID)
	assert.Equal(t, "Shared", *res.Folder, "Folder should match")
}

// Test_security_services_DecryptionProfilesAPIService_GetByID tests retrieving a Decryption Profile by its ID.
func Test_security_services_DecryptionProfilesAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-decrypt-get-" + common.GenerateRandomString(6)
	profile := security_services.DecryptionProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
	}

	createRes, _, err := client.DecryptionProfilesAPI.CreateDecryptionProfiles(context.Background()).DecryptionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Decryption Profile for get test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up Decryption Profile with ID: %s", *createdProfileID)
		_, errDel := client.DecryptionProfilesAPI.DeleteDecryptionProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete Decryption Profile during cleanup")
	}()

	getRes, httpResGet, errGet := client.DecryptionProfilesAPI.GetDecryptionProfilesByID(context.Background(), *createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get Decryption Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
}

// Test_security_services_DecryptionProfilesAPIService_Update tests updating an existing Decryption Profile.
func Test_security_services_DecryptionProfilesAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-decrypt-update-" + common.GenerateRandomString(6)
	profile := security_services.DecryptionProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
	}

	createRes, _, err := client.DecryptionProfilesAPI.CreateDecryptionProfiles(context.Background()).DecryptionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Decryption Profile for update test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up Decryption Profile with ID: %s", *createdProfileID)
		_, errDel := client.DecryptionProfilesAPI.DeleteDecryptionProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete Decryption Profile during cleanup")
	}()

	// test update: change proxy and add proxy
	updatedSslInboundProxy := security_services.DecryptionProfilesSslInboundProxy{}

	updatedProfile := security_services.DecryptionProfiles{
		Name:            profileName, // Name must match
		Folder:          common.StringPtr("Shared"),
		SslInboundProxy: &updatedSslInboundProxy,
	}

	updateRes, httpResUpdate, errUpdate := client.DecryptionProfilesAPI.UpdateDecryptionProfilesByID(context.Background(), *createdProfileID).DecryptionProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update Decryption Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")

	// the updated one should now have the new setting
	assert.True(t, updateRes.HasSslInboundProxy(), "Updated profile should have SslInboundProxy set")
}

// Test_security_services_DecryptionProfilesAPIService_List tests listing Decryption Profiles.
func Test_security_services_DecryptionProfilesAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-decrypt-list-" + common.GenerateRandomString(6)
	profile := security_services.DecryptionProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
	}

	createRes, _, err := client.DecryptionProfilesAPI.CreateDecryptionProfiles(context.Background()).DecryptionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Decryption Profile for list test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up Decryption Profile with ID: %s", *createdProfileID)
		_, errDel := client.DecryptionProfilesAPI.DeleteDecryptionProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete Decryption Profile during cleanup")
	}()

	// list request will use the 'Folder' query parameter to filter
	listRes, httpResList, errList := client.DecryptionProfilesAPI.ListDecryptionProfiles(context.Background()).Folder("Shared").Limit(100).Execute()
	require.NoError(t, errList, "Failed to list Decryption Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	foundObject := false
	for _, p := range listRes.Data {
		if p.Name != "" && p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created Decryption Profile should be found in the list")
}

// Test_security_services_DecryptionProfilesAPIService_DeleteByID tests deleting a Decryption Profile by its ID.
func Test_security_services_DecryptionProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-decrypt-delete-" + common.GenerateRandomString(6)
	profile := security_services.DecryptionProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
	}

	createRes, _, err := client.DecryptionProfilesAPI.CreateDecryptionProfiles(context.Background()).DecryptionProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Decryption Profile for delete test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// delete the profile
	httpResDel, errDel := client.DecryptionProfilesAPI.DeleteDecryptionProfilesByID(context.Background(), *createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete Decryption Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status on successful delete")
}
