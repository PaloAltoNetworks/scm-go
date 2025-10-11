/*
 * Identity Services Testing
 *
 * CertificateProfilesAPIService
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

// Test_identityservices_CertificateProfilesAPIService_Create tests the creation of a Certificate Profile.
func Test_identityservices_CertificateProfilesAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupIdentitySvcTestClient(t)

	// Create a valid Certificate Profile object with a unique name.
	profileName := "test_cp_all_fields-" + common.GenerateRandomString(6)
	profile := identity_services.CertificateProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		CaCertificates: []identity_services.CertificateProfilesCaCertificatesInner{
			{
				Name:           "Forward-Trust-CA",
				DefaultOcspUrl: common.StringPtr("http://test.com"),
				OcspVerifyCert: common.StringPtr("Forward-Trust-CA-ECDSA"),
				TemplateName:   common.StringPtr("something"),
			},
		},
		Domain:                   common.StringPtr("test"),
		UseCrl:                   common.BoolPtr(true),
		UseOcsp:                  common.BoolPtr(true),
		BlockUnknownCert:         common.BoolPtr(true),
		BlockTimeoutCert:         common.BoolPtr(true),
		BlockUnauthenticatedCert: common.BoolPtr(true),
		BlockExpiredCert:         common.BoolPtr(true),
		UsernameField: &identity_services.CertificateProfilesUsernameField{
			Subject: common.StringPtr("common-name"),
		},
		CrlReceiveTimeout:  common.StringPtr("5"),
		OcspReceiveTimeout: common.StringPtr("5"),
		CertStatusTimeout:  common.StringPtr("5"),
	}

	fmt.Printf("Attempting to create Certificate Profile with name: %s\n", profile.Name)

	// Make the create request to the API.
	req := client.CertificateProfilesAPI.CreateCertificateProfiles(context.Background()).CertificateProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create Certificate Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotEmpty(t, res.Id, "Created profile should have an ID")

	createdProfileID := res.Id

	// Defer the cleanup of the created profile.
	defer func() {
		t.Logf("Cleaning up Certificate Profile with ID: %s", *createdProfileID)
		reqDel := client.CertificateProfilesAPI.DeleteCertificateProfilesByID(context.Background(), *createdProfileID)
		httpResDel, errDel := reqDel.Execute()
		if errDel != nil {
			handleAPIError(errDel)
		}
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
		assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK for cleanup delete")
	}()

	// Assert response object properties.
	assert.Equal(t, profileName, res.Name, "Created profile name should match")
	require.NotNil(t, res.Folder, "Folder should not be nil in response")
	assert.True(t, *res.Folder == "Shared" || *res.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	require.Len(t, res.CaCertificates, 1, "Should have one CA certificate")
	assert.Equal(t, "Forward-Trust-CA", res.CaCertificates[0].Name, "CA certificate name should match")
	assert.Equal(t, "http://test.com", *res.CaCertificates[0].DefaultOcspUrl)
	assert.Equal(t, "Forward-Trust-CA-ECDSA", *res.CaCertificates[0].OcspVerifyCert)
	assert.Equal(t, "something", *res.CaCertificates[0].TemplateName)
	assert.Equal(t, "test", *res.Domain)
	assert.True(t, *res.UseCrl)
	assert.True(t, *res.UseOcsp)
	assert.True(t, *res.BlockUnknownCert)
	assert.True(t, *res.BlockTimeoutCert)
	assert.True(t, *res.BlockUnauthenticatedCert)
	assert.True(t, *res.BlockExpiredCert)
	assert.Equal(t, "common-name", *res.UsernameField.Subject)
	assert.Equal(t, "5", *res.CrlReceiveTimeout)
	assert.Equal(t, "5", *res.OcspReceiveTimeout)
	assert.Equal(t, "5", *res.CertStatusTimeout)
	t.Logf("Successfully created and validated Certificate Profile: %s with ID: %s", profile.Name, *createdProfileID)
}

// Test_identityservices_CertificateProfilesAPIService_GetByID tests retrieving a profile by its ID.
func Test_identityservices_CertificateProfilesAPIService_GetByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Create a profile to retrieve.
	profileName := "test-cert-get-" + common.GenerateRandomString(6)
	profile := identity_services.CertificateProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		CaCertificates: []identity_services.CertificateProfilesCaCertificatesInner{
			{
				Name:           "Forward-Trust-CA",
				DefaultOcspUrl: common.StringPtr("http://test.com"),
				OcspVerifyCert: common.StringPtr("Forward-Trust-CA-ECDSA"),
				TemplateName:   common.StringPtr("something"),
			},
		},
		Domain:                   common.StringPtr("test"),
		UseCrl:                   common.BoolPtr(true),
		UseOcsp:                  common.BoolPtr(true),
		BlockUnknownCert:         common.BoolPtr(true),
		BlockTimeoutCert:         common.BoolPtr(true),
		BlockUnauthenticatedCert: common.BoolPtr(true),
		BlockExpiredCert:         common.BoolPtr(true),
		UsernameField: &identity_services.CertificateProfilesUsernameField{
			Subject: common.StringPtr("common-name"),
		},
		CrlReceiveTimeout:  common.StringPtr("5"),
		OcspReceiveTimeout: common.StringPtr("5"),
		CertStatusTimeout:  common.StringPtr("5"),
	}

	reqCreate := client.CertificateProfilesAPI.CreateCertificateProfiles(context.Background()).CertificateProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer cleanup.
	defer func() {
		t.Logf("Cleaning up Certificate Profile with ID: %s", createdProfileID)
		_, errDel := client.CertificateProfilesAPI.DeleteCertificateProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
	}()

	// Test Get by ID operation.
	fmt.Printf("Attempting to get Certificate Profile with ID: %s\n", createdProfileID)
	reqGetById := client.CertificateProfilesAPI.GetCertificateProfilesByID(context.Background(), createdProfileID)
	getRes, httpResGet, errGet := reqGetById.Execute()

	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful.
	require.NoError(t, errGet, "Failed to get profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
	assert.Equal(t, createRes.Id, getRes.Id, "Profile ID should match")
	t.Logf("Successfully retrieved Certificate Profile: %s", getRes.Name)
}

// Test_identityservices_CertificateProfilesAPIService_Update tests updating an existing profile.
func Test_identityservices_CertificateProfilesAPIService_Update(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Create a profile to update.
	profileName := "test-cert-update-" + common.GenerateRandomString(6)
	profile := identity_services.CertificateProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		CaCertificates: []identity_services.CertificateProfilesCaCertificatesInner{
			{
				Name:           "Forward-Trust-CA",
				DefaultOcspUrl: common.StringPtr("http://test.com"),
				OcspVerifyCert: common.StringPtr("Forward-Trust-CA-ECDSA"),
				TemplateName:   common.StringPtr("something"),
			},
		},
		Domain:                   common.StringPtr("test"),
		UseCrl:                   common.BoolPtr(true),
		UseOcsp:                  common.BoolPtr(true),
		BlockUnknownCert:         common.BoolPtr(true),
		BlockTimeoutCert:         common.BoolPtr(true),
		BlockUnauthenticatedCert: common.BoolPtr(true),
		BlockExpiredCert:         common.BoolPtr(true),
		UsernameField: &identity_services.CertificateProfilesUsernameField{
			Subject: common.StringPtr("common-name"),
		},
		CrlReceiveTimeout:  common.StringPtr("5"),
		OcspReceiveTimeout: common.StringPtr("5"),
		CertStatusTimeout:  common.StringPtr("5"),
	}

	reqCreate := client.CertificateProfilesAPI.CreateCertificateProfiles(context.Background()).CertificateProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for update test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer cleanup.
	defer func() {
		t.Logf("Cleaning up Certificate Profile with ID: %s", createdProfileID)
		_, errDel := client.CertificateProfilesAPI.DeleteCertificateProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
	}()

	// Define the update payload.
	updatedProfile := identity_services.CertificateProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		CaCertificates: []identity_services.CertificateProfilesCaCertificatesInner{
			{
				Name:           "Forward-Trust-CA",
				DefaultOcspUrl: common.StringPtr("http://test.com"),
				OcspVerifyCert: common.StringPtr("Forward-Trust-CA-ECDSA"),
				TemplateName:   common.StringPtr("something"),
			},
		},
		Domain:                   common.StringPtr("test-domain"),
		UseCrl:                   common.BoolPtr(true),
		UseOcsp:                  common.BoolPtr(true),
		BlockUnknownCert:         common.BoolPtr(true),
		BlockTimeoutCert:         common.BoolPtr(true),
		BlockUnauthenticatedCert: common.BoolPtr(true),
		BlockExpiredCert:         common.BoolPtr(false),
		UsernameField: &identity_services.CertificateProfilesUsernameField{
			Subject: common.StringPtr("common-name"),
		},
		CrlReceiveTimeout:  common.StringPtr("5"),
		OcspReceiveTimeout: common.StringPtr("5"),
		CertStatusTimeout:  common.StringPtr("5"),
	}

	fmt.Printf("Attempting to update Certificate Profile with ID: %s\n", createdProfileID)
	reqUpdate := client.CertificateProfilesAPI.UpdateCertificateProfilesByID(context.Background(), createdProfileID).CertificateProfiles(updatedProfile)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()

	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update was successful.
	require.NoError(t, errUpdate, "Failed to update profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.False(t, *updateRes.BlockExpiredCert, "BlockExpiredCert should be updated")
	assert.Equal(t, "test-domain", *updateRes.Domain, "Domain should be updated")
	require.Len(t, updateRes.CaCertificates, 1)

	t.Logf("Successfully updated Certificate Profile: %s", profileName)
}

// Test_identityservices_CertificateProfilesAPIService_List tests listing Certificate Profiles.
func Test_identityservices_CertificateProfilesAPIService_List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Create a profile to ensure it appears in the list.
	profileName := "test-cert-list-" + common.GenerateRandomString(6)
	profile := identity_services.CertificateProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		CaCertificates: []identity_services.CertificateProfilesCaCertificatesInner{
			{
				Name:           "Forward-Trust-CA",
				DefaultOcspUrl: common.StringPtr("http://test.com"),
				OcspVerifyCert: common.StringPtr("Forward-Trust-CA-ECDSA"),
				TemplateName:   common.StringPtr("something"),
			},
		},
		Domain:                   common.StringPtr("test-domain"),
		UseCrl:                   common.BoolPtr(true),
		UseOcsp:                  common.BoolPtr(true),
		BlockUnknownCert:         common.BoolPtr(true),
		BlockTimeoutCert:         common.BoolPtr(true),
		BlockUnauthenticatedCert: common.BoolPtr(true),
		BlockExpiredCert:         common.BoolPtr(false),
		UsernameField: &identity_services.CertificateProfilesUsernameField{
			Subject: common.StringPtr("common-name"),
		},
		CrlReceiveTimeout:  common.StringPtr("5"),
		OcspReceiveTimeout: common.StringPtr("5"),
		CertStatusTimeout:  common.StringPtr("5"),
	}

	reqCreate := client.CertificateProfilesAPI.CreateCertificateProfiles(context.Background()).CertificateProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for list test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	// Defer cleanup.
	defer func() {
		t.Logf("Cleaning up Certificate Profile with ID: %s", createdProfileID)
		_, errDel := client.CertificateProfilesAPI.DeleteCertificateProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete profile during cleanup")
	}()

	// Test List operation.
	fmt.Println("Attempting to list Certificate Profiles")
	reqList := client.CertificateProfilesAPI.ListCertificateProfiles(context.Background()).Folder("Shared").Limit(10000)
	listRes, httpResList, errList := reqList.Execute()

	// printing the list response for optional visibility
	fmt.Printf("Retrieved %d Certificate Profiles\n", len(listRes.Data))

	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful.
	require.NoError(t, errList, "Failed to list profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	assert.True(t, len(listRes.Data) > 0, "Should have at least one profile in the list")

	// Verify our created profile is in the list.
	foundProfile := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			foundProfile = true
			require.Len(t, p.CaCertificates, 1)
			break
		}
	}
	assert.True(t, foundProfile, "Created profile should be found in the list")
	t.Logf("Successfully listed Certificate Profiles, found created profile: %s", profileName)
}

// Test_identityservices_CertificateProfilesAPIService_DeleteByID tests deleting a profile by ID.
func Test_identityservices_CertificateProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Create a profile to delete.
	profileName := "test-cert-delete-" + common.GenerateRandomString(6)
	profile := identity_services.CertificateProfiles{
		Folder: common.StringPtr("Shared"),
		Name:   profileName,
		CaCertificates: []identity_services.CertificateProfilesCaCertificatesInner{
			{
				Name:           "Forward-Trust-CA",
				DefaultOcspUrl: common.StringPtr("http://test.com"),
				OcspVerifyCert: common.StringPtr("Forward-Trust-CA-ECDSA"),
				TemplateName:   common.StringPtr("something"),
			},
		},
		Domain:                   common.StringPtr("test-domain"),
		UseCrl:                   common.BoolPtr(true),
		UseOcsp:                  common.BoolPtr(true),
		BlockUnknownCert:         common.BoolPtr(true),
		BlockTimeoutCert:         common.BoolPtr(true),
		BlockUnauthenticatedCert: common.BoolPtr(true),
		BlockExpiredCert:         common.BoolPtr(false),
		UsernameField: &identity_services.CertificateProfilesUsernameField{
			Subject: common.StringPtr("common-name"),
		},
		CrlReceiveTimeout:  common.StringPtr("5"),
		OcspReceiveTimeout: common.StringPtr("5"),
		CertStatusTimeout:  common.StringPtr("5"),
	}
	reqCreate := client.CertificateProfilesAPI.CreateCertificateProfiles(context.Background()).CertificateProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	require.NoError(t, err, "Failed to create profile for delete test")
	createdProfileID := *createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	t.Logf("Created Certificate Profile for Delete test with ID: %s", createdProfileID)

	// Test Delete by ID operation.
	fmt.Printf("Attempting to delete Certificate Profile with ID: %s\n", createdProfileID)
	reqDel := client.CertificateProfilesAPI.DeleteCertificateProfilesByID(context.Background(), createdProfileID)
	httpResDel, errDel := reqDel.Execute()

	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful.
	require.NoError(t, errDel, "Failed to delete profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
	t.Logf("Successfully deleted Certificate Profile: %s", createdProfileID)
}
