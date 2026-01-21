package identity_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// createSimpleScepProfile creates a basic SCEP profile payload with "none" challenge
func createSimpleScepProfile(name string) identity_services.ScepProfiles {
	algorithm := identity_services.ScepProfilesAlgorithm{
		Rsa: identity_services.ScepProfilesAlgorithmRsa{
			RsaNbits: "2048",
		},
	}

	challenge := identity_services.ScepProfilesScepChallenge{
		Fixed: common.StringPtr("mypassword123"),
	}

	return identity_services.ScepProfiles{
		Folder:         common.StringPtr("All"),
		Name:           name,
		ScepUrl:        "https://scep.example.com/",
		CaIdentityName: "Default",
		Digest:         "sha256",
		Subject:        "CN=$USERNAME",
		Algorithm:      algorithm,
		ScepChallenge:  challenge,
	}
}

// createComplexScepProfile creates a complete SCEP profile with a fixed challenge and attributes
func createComplexScepProfile(name string) identity_services.ScepProfiles {
	algorithm := identity_services.ScepProfilesAlgorithm{
		Rsa: identity_services.ScepProfilesAlgorithmRsa{
			RsaNbits: "2048",
		},
	}

	dynamicSettings := &identity_services.ScepProfilesScepChallengeDynamic{
		Username:     common.StringPtr("scep-admin"),
		Password:     common.StringPtr("mypassword123"),
		OtpServerUrl: common.StringPtr("https://otp.example.com/api/v1/generate"),
	}

	challenge := identity_services.ScepProfilesScepChallenge{
		Dynamic: dynamicSettings,
	}

	attributes := &identity_services.ScepProfilesCertificateAttributes{
		Dnsname: common.StringPtr("device.example.com"),
	}

	return identity_services.ScepProfiles{
		Folder:                common.StringPtr("All"),
		Name:                  name,
		ScepUrl:               "https://scep.example.com/certsrv/mscep/mscep.dll",
		CaIdentityName:        "Example-Name",
		Digest:                "sha256",
		Subject:               "CN=$USERNAME",
		Fingerprint:           common.StringPtr("D14A028C2A3A2BC9476102BB288234C415A2B01F"),
		Algorithm:             algorithm,
		ScepChallenge:         challenge,
		ScepCaCert:            common.StringPtr("Forward-Trust-CA"),
		ScepClientCert:        common.StringPtr("Forward-UnTrust-CA"),
		CertificateAttributes: attributes,
		UseAsDigitalSignature: common.BoolPtr(true),
		UseForKeyEncipherment: common.BoolPtr(true),
	}
}

// Test_ScepProfilesAPIService_Create tests the creation of a SCEP Profile.
func Test_ScepProfilesAPIService_Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	createdName := generateRandomName("scep-create")
	payload := createComplexScepProfile(createdName)

	t.Logf("Creating SCEP Profile: %s", createdName)
	req := client.SCEPProfilesAPI.CreateSCEPProfiles(context.Background()).ScepProfiles(payload)
	res, httpRes, err := req.Execute()

	require.NoError(t, err, "Failed to create SCEP Profile")
	assert.Equal(t, 201, httpRes.StatusCode)
	require.NotNil(t, res)
	assert.Equal(t, createdName, res.Name)

	createdID := res.Id
	defer func() {
		client.SCEPProfilesAPI.DeleteSCEPProfilesByID(context.Background(), createdID).Execute()
	}()
}

// Test_ScepProfilesAPIService_GetByID tests retrieving a SCEP Profile by ID.
func Test_ScepProfilesAPIService_GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	name := generateRandomName("scep-get")
	payload := createSimpleScepProfile(name)

	createRes, _, err := client.SCEPProfilesAPI.CreateSCEPProfiles(context.Background()).ScepProfiles(payload).Execute()
	require.NoError(t, err)
	createdID := createRes.Id

	defer func() {
		client.SCEPProfilesAPI.DeleteSCEPProfilesByID(context.Background(), createdID).Execute()
	}()

	getRes, httpRes, err := client.SCEPProfilesAPI.GetSCEPProfilesByID(context.Background(), createdID).Execute()
	require.NoError(t, err)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, name, getRes.Name)
	assert.Equal(t, "sha256", getRes.Digest)
}

// Test_ScepProfilesAPIService_Update tests updating a SCEP Profile.
func Test_ScepProfilesAPIService_Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	name := generateRandomName("scep-update")
	payload := createSimpleScepProfile(name)

	createRes, _, err := client.SCEPProfilesAPI.CreateSCEPProfiles(context.Background()).ScepProfiles(payload).Execute()
	require.NoError(t, err)
	createdID := createRes.Id

	defer func() {
		client.SCEPProfilesAPI.DeleteSCEPProfilesByID(context.Background(), createdID).Execute()
	}()

	// Update the profile to use a different digest and identity
	updatedPayload := createRes
	updatedPayload.Digest = "sha512"
	updatedPayload.CaIdentityName = "Updated-CA"

	updateRes, httpRes, err := client.SCEPProfilesAPI.UpdateSCEPProfilesByID(context.Background(), createdID).ScepProfiles(*updatedPayload).Execute()
	require.NoError(t, err)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, "sha512", updateRes.Digest)
	assert.Equal(t, "Updated-CA", updateRes.CaIdentityName)
}

// Test_ScepProfilesAPIService_List tests listing SCEP Profiles.
func Test_ScepProfilesAPIService_List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	name := generateRandomName("scep-list")
	payload := createSimpleScepProfile(name)

	createRes, _, err := client.SCEPProfilesAPI.CreateSCEPProfiles(context.Background()).ScepProfiles(payload).Execute()
	require.NoError(t, err)
	createdID := createRes.Id

	defer func() {
		client.SCEPProfilesAPI.DeleteSCEPProfilesByID(context.Background(), createdID).Execute()
	}()

	listRes, httpRes, err := client.SCEPProfilesAPI.ListSCEPProfiles(context.Background()).Folder("All").Execute()
	require.NoError(t, err)
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, listRes)

	found := false
	for _, item := range listRes.Data {
		if item.Name == name {
			found = true
			break
		}
	}
	assert.True(t, found, "Created SCEP Profile should be in the list")
}

// Test_ScepProfilesAPIService_DeleteByID tests deleting a SCEP Profile.
func Test_ScepProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	name := generateRandomName("scep-delete")
	payload := createSimpleScepProfile(name)

	createRes, _, err := client.SCEPProfilesAPI.CreateSCEPProfiles(context.Background()).ScepProfiles(payload).Execute()
	require.NoError(t, err)
	createdID := createRes.Id

	httpRes, err := client.SCEPProfilesAPI.DeleteSCEPProfilesByID(context.Background(), createdID).Execute()
	require.NoError(t, err)
	assert.True(t, httpRes.StatusCode == 200 || httpRes.StatusCode == 204)
}
