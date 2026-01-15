/*
 * Network Services Testing
 *
 * LLDPProfilesAPIService
 */

package network_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Test_network_services_LLDPProfilesAPIService_Create tests the creation of an LLDP Profile.
func Test_network_services_LLDPProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	createdProfileName := "test-lldp-" + common.GenerateRandomString(6)

	// Define an LLDP Profile based on the provided JSON sample
	profile := network_services.LldpProfiles{
		Name:                   createdProfileName,
		Folder:                 common.StringPtr("All"),
		Mode:                   common.StringPtr("transmit-receive"),
		SnmpSyslogNotification: common.BoolPtr(true),
		OptionTlvs: &network_services.LldpProfilesOptionTlvs{
			PortDescription:    common.BoolPtr(false),
			SystemName:         common.BoolPtr(true),
			SystemDescription:  common.BoolPtr(false),
			SystemCapabilities: common.BoolPtr(true),
			ManagementAddress: &network_services.LldpProfilesOptionTlvsManagementAddress{
				Enabled: common.BoolPtr(false),
			},
		},
	}

	fmt.Printf("Creating LLDP Profile with name: %s\n", profile.Name)
	req := client.LLDPProfilesAPI.CreateLLDPProfiles(context.Background()).LldpProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create LLDP Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")

	createdProfileID := *res.Id

	defer func() {
		t.Logf("Cleaning up LLDP Profile with ID: %s", createdProfileID)
		_, errDel := client.LLDPProfilesAPI.DeleteLLDPProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete LLDP Profile during cleanup")
	}()

	t.Logf("Successfully created LLDP Profile: %s with ID: %s", profile.Name, createdProfileID)
	assert.Equal(t, "All", *res.Folder, "Folder should match")
	assert.Equal(t, "transmit-receive", *res.Mode, "Mode should match")
	assert.True(t, *res.SnmpSyslogNotification, "SnmpSyslogNotification should be true")
}

// Test_network_services_LLDPProfilesAPIService_GetByID tests retrieving an LLDP Profile by its ID.
func Test_network_services_LLDPProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-lldp-get-" + common.GenerateRandomString(6)

	// Define an LLDP Profile based on the provided JSON sample
	profile := network_services.LldpProfiles{
		Name:                   profileName,
		Folder:                 common.StringPtr("All"),
		Mode:                   common.StringPtr("transmit-receive"),
		SnmpSyslogNotification: common.BoolPtr(true),
		OptionTlvs: &network_services.LldpProfilesOptionTlvs{
			PortDescription:    common.BoolPtr(false),
			SystemName:         common.BoolPtr(true),
			SystemDescription:  common.BoolPtr(false),
			SystemCapabilities: common.BoolPtr(true),
			ManagementAddress: &network_services.LldpProfilesOptionTlvsManagementAddress{
				Enabled: common.BoolPtr(false),
			},
		},
	}

	createRes, _, err := client.LLDPProfilesAPI.CreateLLDPProfiles(context.Background()).LldpProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create LLDP Profile for get test")
	createdProfileID := *createRes.Id

	defer func() {
		client.LLDPProfilesAPI.DeleteLLDPProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	getRes, httpResGet, errGet := client.LLDPProfilesAPI.GetLLDPProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get LLDP Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
}

// Test_network_services_LLDPProfilesAPIService_Update tests updating an existing LLDP Profile.
func Test_network_services_LLDPProfilesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-lldp-upd-" + common.GenerateRandomString(6)

	profile := network_services.LldpProfiles{
		Name:                   profileName,
		Folder:                 common.StringPtr("All"),
		Mode:                   common.StringPtr("transmit-receive"),
		SnmpSyslogNotification: common.BoolPtr(true),
		OptionTlvs: &network_services.LldpProfilesOptionTlvs{
			PortDescription:    common.BoolPtr(false),
			SystemName:         common.BoolPtr(true),
			SystemDescription:  common.BoolPtr(false),
			SystemCapabilities: common.BoolPtr(true),
			ManagementAddress: &network_services.LldpProfilesOptionTlvsManagementAddress{
				Enabled: common.BoolPtr(false),
			},
		},
	}

	createRes, _, err := client.LLDPProfilesAPI.CreateLLDPProfiles(context.Background()).LldpProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create LLDP Profile for update test")
	createdProfileID := *createRes.Id

	defer func() {
		client.LLDPProfilesAPI.DeleteLLDPProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	updatedProfile := network_services.LldpProfiles{
		Name:                   profileName,
		Folder:                 common.StringPtr("All"),
		Mode:                   common.StringPtr("transmit-receive"),
		SnmpSyslogNotification: common.BoolPtr(true),
		OptionTlvs: &network_services.LldpProfilesOptionTlvs{
			PortDescription:    common.BoolPtr(true),  // Updated
			SystemName:         common.BoolPtr(false), // Updated
			SystemDescription:  common.BoolPtr(false),
			SystemCapabilities: common.BoolPtr(true),
			ManagementAddress: &network_services.LldpProfilesOptionTlvsManagementAddress{
				Enabled: common.BoolPtr(false),
			},
		},
	}

	updateRes, httpResUpdate, errUpdate := client.LLDPProfilesAPI.UpdateLLDPProfilesByID(context.Background(), createdProfileID).LldpProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update LLDP Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.False(t, *updateRes.OptionTlvs.SystemName, "SystemName TLV should be updated to false")
	assert.True(t, *updateRes.OptionTlvs.PortDescription, "PortDescription TLV should be updated to false")
}

// Test_network_services_LLDPProfilesAPIService_List tests listing LLDP Profiles.
func Test_network_services_LLDPProfilesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-lldp-list-" + common.GenerateRandomString(6)

	profile := network_services.LldpProfiles{
		Name:   profileName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.LLDPProfilesAPI.CreateLLDPProfiles(context.Background()).LldpProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create LLDP Profile for list test")
	createdProfileID := *createRes.Id

	defer func() {
		client.LLDPProfilesAPI.DeleteLLDPProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	listRes, httpResList, errList := client.LLDPProfilesAPI.ListLLDPProfiles(context.Background()).Folder("All").Execute()
	require.NoError(t, errList, "Failed to list LLDP Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
}

// Test_network_services_LLDPProfilesAPIService_DeleteByID tests deleting an LLDP Profile.
func Test_network_services_LLDPProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-lldp-del-" + common.GenerateRandomString(6)

	profile := network_services.LldpProfiles{
		Name:   profileName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.LLDPProfilesAPI.CreateLLDPProfiles(context.Background()).LldpProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create LLDP Profile for delete test")
	createdProfileID := *createRes.Id

	httpResDel, errDel := client.LLDPProfilesAPI.DeleteLLDPProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete LLDP Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}
