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

// createSAMLServerProfile creates a simple complete SAML server profile payload
func createSAMLServerProfile(name string) identity_services.SamlServerProfiles {
	return identity_services.SamlServerProfiles{
		Name:        name,
		Folder:      common.StringPtr("All"),
		Certificate: "Global Authentication Cookie Cert",
		EntityId:    "https://idp.example.com/entity",
		SsoUrl:      "https://idp.example.com/sso",
		SsoBindings: "redirect",
	}
}

// createComplexSAMLServerProfile creates a complex complete SAML server profile payload
func createComplexSAMLServerProfile(name string) identity_services.SamlServerProfiles {
	return identity_services.SamlServerProfiles{
		Name:                   name,
		Folder:                 common.StringPtr("All"),
		Certificate:            "Global Authentication Cookie Cert",
		EntityId:               "https://idp.complex.com/entity",
		SsoUrl:                 "https://idp.complex.com/sso",
		SsoBindings:            "post",
		SloUrl:                 common.StringPtr("https://idp.complex.com/slo"),
		SloBindings:            common.StringPtr("redirect"),
		MaxClockSkew:           common.Int32Ptr(180),
		ValidateIdpCertificate: common.BoolPtr(false),
		WantAuthRequestsSigned: common.BoolPtr(true),
	}
}

// Test_identity_services_SAMLServerProfilesAPIService_Create tests the creation of a SAML Server Profile.
func Test_identity_services_SAMLServerProfilesAPIService_Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	createdProfileName := generateRandomName("saml-prof-create")
	profilePayload := createComplexSAMLServerProfile(createdProfileName)

	fmt.Printf("Creating SAML Server Profile with name: %s\n", profilePayload.Name)
	req := client.SAMLServerProfilesAPI.CreateSAMLServerProfiles(context.Background()).SamlServerProfiles(profilePayload)
	res, httpRes, err := req.Execute()

	// assertions and cleanup
	require.NoError(t, err, "Failed to create SAML Server Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")

	createdProfileID := res.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up SAML Server Profile with ID: %s", createdProfileID)
		_, errDel := client.SAMLServerProfilesAPI.DeleteSAMLServerProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete SAML Server Profile during cleanup")
	}()

	t.Logf("Successfully created SAML Server Profile: %s with ID: %s", createdProfileName, createdProfileID)
}

// Test_identity_services_SAMLServerProfilesAPIService_GetByID tests retrieving a SAML Server Profile by its ID.
func Test_identity_services_SAMLServerProfilesAPIService_GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	profileName := generateRandomName("test-saml-prof-get")
	profilePayload := createSAMLServerProfile(profileName)

	createRes, _, err := client.SAMLServerProfilesAPI.CreateSAMLServerProfiles(context.Background()).SamlServerProfiles(profilePayload).Execute()
	require.NoError(t, err, "Failed to create SAML Server Profile for get test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// cleanup
	defer func() {
		t.Logf("Cleaning up SAML Server Profile with ID: %s", createdProfileID)
		_, errDel := client.SAMLServerProfilesAPI.DeleteSAMLServerProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete SAML Server Profile during cleanup")
	}()

	getRes, httpResGet, errGet := client.SAMLServerProfilesAPI.GetSAMLServerProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get SAML Server Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
}

// Test_identity_services_SAMLServerProfilesAPIService_Update tests updating an existing SAML Server Profile.
func Test_identity_services_SAMLServerProfilesAPIService_Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	profileName := generateRandomName("test-saml-profile-update")
	profilePayload := createSAMLServerProfile(profileName)

	createRes, _, err := client.SAMLServerProfilesAPI.CreateSAMLServerProfiles(context.Background()).SamlServerProfiles(profilePayload).Execute()
	require.NoError(t, err, "Failed to create SAML Server Profile for update test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// cleanup
	defer func() {
		_, errDel := client.SAMLServerProfilesAPI.DeleteSAMLServerProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete SAML Server Profile during cleanup")
	}()

	// creating updated object
	updatedProfile := createSAMLServerProfile(profileName)
	updatedProfile.SsoUrl = "https://idp.updated.com/sso"
	updatedProfile.MaxClockSkew = common.Int32Ptr(500)

	// applying the update
	updateRes, httpResUpdate, errUpdate := client.SAMLServerProfilesAPI.UpdateSAMLServerProfilesByID(context.Background(), createdProfileID).SamlServerProfiles(updatedProfile).Execute()

	// assertions
	require.NoError(t, errUpdate, "Failed to update SAML Server Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, "https://idp.updated.com/sso", updateRes.SsoUrl, "SSO URL should be updated")
	assert.Equal(t, int32(500), *updateRes.MaxClockSkew, "Max Clock Skew should be updated")
}

// Test_identity_services_SAMLServerProfilesAPIService_List tests listing SAML Server Profiles.
func Test_identity_services_SAMLServerProfilesAPIService_List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Create a resource to ensure the list is non-empty
	profileName := generateRandomName("saml-prof-list")
	profilePayload := createSAMLServerProfile(profileName)

	createRes, _, err := client.SAMLServerProfilesAPI.CreateSAMLServerProfiles(context.Background()).SamlServerProfiles(profilePayload).Execute()
	require.NoError(t, err, "Failed to create SAML Server Profile for list test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up SAML Server Profile with ID: %s", createdProfileID)
		_, errDel := client.SAMLServerProfilesAPI.DeleteSAMLServerProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete SAML Server Profile during cleanup")
	}()

	// LIST request
	listRes, httpResList, errList := client.SAMLServerProfilesAPI.ListSAMLServerProfiles(context.Background()).Folder("All").Execute()
	require.NoError(t, errList, "Failed to list SAML Server Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	// verify the created object is in the list
	foundObject := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created SAML Server Profile should be found in the list")
}

// Test_identity_services_SAMLServerProfilesAPIService_DeleteByID tests deleting a SAML Server Profile.
func Test_identity_services_SAMLServerProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	profileName := generateRandomName("test-saml-profile-delete")
	profilePayload := createSAMLServerProfile(profileName)

	createRes, _, err := client.SAMLServerProfilesAPI.CreateSAMLServerProfiles(context.Background()).SamlServerProfiles(profilePayload).Execute()
	require.NoError(t, err, "Failed to create SAML Server Profile for delete test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// DELETE request
	httpResDel, errDel := client.SAMLServerProfilesAPI.DeleteSAMLServerProfilesByID(context.Background(), createdProfileID).Execute()

	require.NoError(t, errDel, "Failed to delete SAML Server Profile")
	assert.True(t, httpResDel.StatusCode == 200 || httpResDel.StatusCode == 204, "Expected 200 or 204 status for delete")

	t.Logf("Successfully deleted SAML Server Profile with ID: %s", createdProfileID)
}
