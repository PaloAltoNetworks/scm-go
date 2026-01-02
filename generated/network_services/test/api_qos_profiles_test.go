/*
 * Network Services Testing
 *
 * QoSProfilesAPIService
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

// Test_network_services_QoSProfilesAPIService_Create tests the creation of a QoS Profile.
func Test_network_services_QoSProfilesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	createdProfileName := "test-qos-create-" + common.GenerateRandomString(6)

	// define bandwidth classes using Mbps
	testClasses := []network_services.QosProfilesClassBandwidthTypeMbpsClassInner{
		{
			Name:     common.StringPtr("class1"),
			Priority: common.StringPtr("low"),
		},
		{
			Name:     common.StringPtr("class2"),
			Priority: common.StringPtr("real-time"),
			ClassBandwidth: &network_services.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth{
				EgressGuaranteed: common.Int32Ptr(10),
				EgressMax:        common.Int32Ptr(20),
			},
		},
		{
			Name:     common.StringPtr("class3"),
			Priority: common.StringPtr("high"),
			ClassBandwidth: &network_services.QosProfilesClassBandwidthTypeMbpsClassInnerClassBandwidth{
				EgressGuaranteed: common.Int32Ptr(1000),
				EgressMax:        common.Int32Ptr(10000),
			},
		},
	}

	// define a QoS Profile
	profile := network_services.QosProfiles{
		Folder: common.StringPtr("Service Connections"),
		Name:   createdProfileName,
		ClassBandwidthType: &network_services.QosProfilesClassBandwidthType{
			Mbps: &network_services.QosProfilesClassBandwidthTypeMbps{
				Class: testClasses,
			},
		},
		AggregateBandwidth: &network_services.QosProfilesAggregateBandwidth{
			EgressGuaranteed: common.Int32Ptr(300),
			EgressMax:        common.Int32Ptr(1000),
		},
	}

	fmt.Printf("Creating QoS Profile with name: %s\n", profile.Name)
	req := client.QoSProfilesAPI.CreateQoSProfiles(context.Background()).QosProfiles(profile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create QoS Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, res.Name, "Created profile name should match")

	createdProfileID := *res.Id

	defer func() {
		t.Logf("Cleaning up QoS Profile with ID: %s", createdProfileID)
		_, errDel := client.QoSProfilesAPI.DeleteQoSProfilesByID(context.Background(), createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete QoS Profile during cleanup")
	}()

	t.Logf("Successfully created QoS Profile: %s with ID: %s", profile.Name, createdProfileID)
	assert.Equal(t, "Service Connections", *res.Folder, "Folder should match")
	assert.NotNil(t, res.ClassBandwidthType.Mbps, "Mbps bandwidth configuration should exist")
	assert.Len(t, res.ClassBandwidthType.Mbps.Class, 3, "Should have 3 classes configured")
}

// Test_network_services_QoSProfilesAPIService_GetByID tests retrieving a QoS Profile by its ID.
func Test_network_services_QoSProfilesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-qos-get-" + common.GenerateRandomString(6)

	profile := network_services.QosProfiles{
		Folder: common.StringPtr("Service Connections"),
		Name:   profileName,
	}

	createRes, _, err := client.QoSProfilesAPI.CreateQoSProfiles(context.Background()).QosProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create QoS Profile for get test")
	createdProfileID := *createRes.Id

	defer func() {
		t.Logf("Cleaning up QoS Profile with ID: %s", createdProfileID)
		client.QoSProfilesAPI.DeleteQoSProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	getRes, httpResGet, errGet := client.QoSProfilesAPI.GetQoSProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get QoS Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name, "Profile name should match")
}

// Test_network_services_QoSProfilesAPIService_Update tests updating an existing QoS Profile.
func Test_network_services_QoSProfilesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-qos-update-" + common.GenerateRandomString(6)

	profile := network_services.QosProfiles{
		Folder: common.StringPtr("Service Connections"),
		Name:   profileName,
		AggregateBandwidth: &network_services.QosProfilesAggregateBandwidth{
			EgressMax: common.Int32Ptr(50),
		},
	}

	createRes, _, err := client.QoSProfilesAPI.CreateQoSProfiles(context.Background()).QosProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create QoS Profile for update test")
	createdProfileID := *createRes.Id

	defer func() {
		t.Logf("Cleaning up QoS Profile with ID: %s", createdProfileID)
		client.QoSProfilesAPI.DeleteQoSProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	updatedEgressMax := int32(200)
	updatedProfile := network_services.QosProfiles{
		Name: profileName,
		AggregateBandwidth: &network_services.QosProfilesAggregateBandwidth{
			EgressMax: common.Int32Ptr(updatedEgressMax),
		},
	}

	updateRes, httpResUpdate, errUpdate := client.QoSProfilesAPI.UpdateQoSProfilesByID(context.Background(), createdProfileID).QosProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update QoS Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, updatedEgressMax, *updateRes.AggregateBandwidth.EgressMax, "Aggregate bandwidth should be updated")
}

// Test_network_services_QoSProfilesAPIService_List tests listing QoS Profiles.
func Test_network_services_QoSProfilesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-qos-list-" + common.GenerateRandomString(6)

	profile := network_services.QosProfiles{
		Folder: common.StringPtr("Service Connections"),
		Name:   profileName,
	}

	createRes, _, err := client.QoSProfilesAPI.CreateQoSProfiles(context.Background()).QosProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create QoS Profile for list test")
	createdProfileID := *createRes.Id

	defer func() {
		t.Logf("Cleaning up QoS Profile with ID: %s", createdProfileID)
		client.QoSProfilesAPI.DeleteQoSProfilesByID(context.Background(), createdProfileID).Execute()
	}()

	listRes, httpResList, errList := client.QoSProfilesAPI.ListQoSProfiles(context.Background()).Folder("Service Connections").Limit(100).Execute()
	require.NoError(t, errList, "Failed to list QoS Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	foundObject := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created QoS Profile should be found in the list")
}

// Test_network_services_QoSProfilesAPIService_DeleteByID tests deleting a QoS Profile by its ID.
func Test_network_services_QoSProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	profileName := "test-qos-delete-" + common.GenerateRandomString(6)

	profile := network_services.QosProfiles{
		Folder: common.StringPtr("Service Connections"),
		Name:   profileName,
	}

	createRes, _, err := client.QoSProfilesAPI.CreateQoSProfiles(context.Background()).QosProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create QoS Profile for delete test")
	createdProfileID := *createRes.Id

	httpResDel, errDel := client.QoSProfilesAPI.DeleteQoSProfilesByID(context.Background(), createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete QoS Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}
