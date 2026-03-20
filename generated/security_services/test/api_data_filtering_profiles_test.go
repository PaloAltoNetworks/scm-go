/*
 * Security Services Testing
 *
 * DataFilteringAPIService
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

// Test_security_services_DataFilteringAPIService_Create tests the creation of a Data Filtering Profile.
func Test_security_services_DataFilteringAPIService_Create(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdProfileName := "test-df-create-" + common.GenerateRandomString(6)

	// First create a data object that will be referenced by the data filtering profile
	dataObjectName := "test-do-for-df-" + common.GenerateRandomString(6)
	predefinedPattern := security_services.DataObjectsPatternTypePredefinedPatternInner{
		Name:     common.StringPtr("ABA-Routing-Number"),
		FileType: []string{"text/html"},
	}
	predefined := security_services.DataObjectsPatternTypePredefined{
		Pattern: []security_services.DataObjectsPatternTypePredefinedPatternInner{predefinedPattern},
	}
	patternType := security_services.DataObjectsPatternType{
		Predefined: &predefined,
	}
	dataObject := security_services.DataObjects{
		Folder:      common.StringPtr("ngfw-shared"),
		Name:        common.StringPtr(dataObjectName),
		PatternType: &patternType,
	}

	doRes, _, errDO := client.DataObjectsAPI.CreateDataObjects(context.Background()).DataObjects(dataObject).Execute()
	require.NoError(t, errDO, "Failed to create Data Object for Data Filtering Profile test")
	dataObjectID := doRes.Id
	require.NotEmpty(t, dataObjectID, "Created data object ID should not be empty")

	defer func() {
		t.Logf("Cleaning up Data Object with ID: %s", *dataObjectID)
		_, errDelDO := client.DataObjectsAPI.DeleteDataObjectsByID(context.Background(), *dataObjectID).Execute()
		require.NoError(t, errDelDO, "Failed to delete Data Object during cleanup")
	}()

	// Define rules for the data filtering profile
	testRules := []security_services.DataFilteringProfilesRulesInner{
		{
			Name:           common.StringPtr("rule0"),
			DataObject:     common.StringPtr(dataObjectName),
			Application:    []string{"any"},
			FileType:       []string{"any"},
			Direction:      common.StringPtr("both"),
			AlertThreshold: common.Int32Ptr(1),
			BlockThreshold: common.Int32Ptr(1),
			LogSeverity:    common.StringPtr("informational"),
		},
	}

	// Define a Data Filtering Profile
	profile := security_services.DataFilteringProfiles{
		Folder:      common.StringPtr("ngfw-shared"),
		Name:        common.StringPtr(createdProfileName),
		Description: common.StringPtr("Test Data Filtering Profile for create API"),
		DataCapture: common.BoolPtr(false),
		Rules:       testRules,
	}

	fmt.Printf("Creating Data Filtering Profile with name: %s\n", createdProfileName)
	req := client.DataFilteringAPI.CreateDataFilteringProfiles(context.Background()).DataFilteringProfiles(profile)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Data Filtering Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdProfileName, *res.Name, "Created profile name should match")
	assert.Len(t, res.Rules, 1, "Created profile should have 1 rule")
	assert.Equal(t, "rule0", *res.Rules[0].Name, "Rule name should match")

	createdProfileID := res.Id

	defer func() {
		t.Logf("Cleaning up Data Filtering Profile with ID: %s", *createdProfileID)
		_, errDel := client.DataFilteringAPI.DeleteDataFilteringProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete Data Filtering Profile during cleanup")
	}()

	t.Logf("Successfully created Data Filtering Profile: %s with ID: %s", createdProfileName, *createdProfileID)
}

// Test_security_services_DataFilteringAPIService_GetByID tests retrieving a Data Filtering Profile by its ID.
func Test_security_services_DataFilteringAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-df-get-" + common.GenerateRandomString(6)
	profile := security_services.DataFilteringProfiles{
		Folder: common.StringPtr("ngfw-shared"),
		Name:   common.StringPtr(profileName),
	}

	createRes, _, err := client.DataFilteringAPI.CreateDataFilteringProfiles(context.Background()).DataFilteringProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Data Filtering Profile for get test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up Data Filtering Profile with ID: %s", *createdProfileID)
		_, errDel := client.DataFilteringAPI.DeleteDataFilteringProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete Data Filtering Profile during cleanup")
	}()

	getRes, httpResGet, errGet := client.DataFilteringAPI.GetDataFilteringProfilesByID(context.Background(), *createdProfileID).Execute()
	require.NoError(t, errGet, "Failed to get Data Filtering Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, *getRes.Name, "Profile name should match")
}

// Test_security_services_DataFilteringAPIService_Update tests updating an existing Data Filtering Profile.
func Test_security_services_DataFilteringAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-df-update-" + common.GenerateRandomString(6)
	profile := security_services.DataFilteringProfiles{
		Folder: common.StringPtr("ngfw-shared"),
		Name:   common.StringPtr(profileName),
	}

	createRes, _, err := client.DataFilteringAPI.CreateDataFilteringProfiles(context.Background()).DataFilteringProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Data Filtering Profile for update test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up Data Filtering Profile with ID: %s", *createdProfileID)
		_, errDel := client.DataFilteringAPI.DeleteDataFilteringProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete Data Filtering Profile during cleanup")
	}()

	updatedDescription := "Updated data filtering profile description"

	updatedProfile := security_services.DataFilteringProfiles{
		Name:        common.StringPtr(profileName),
		Folder:      common.StringPtr("ngfw-shared"),
		Description: common.StringPtr(updatedDescription),
		DataCapture: common.BoolPtr(false),
	}

	updateRes, httpResUpdate, errUpdate := client.DataFilteringAPI.UpdateDataFilteringProfilesByID(context.Background(), *createdProfileID).DataFilteringProfiles(updatedProfile).Execute()
	require.NoError(t, errUpdate, "Failed to update Data Filtering Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
}

// Test_security_services_DataFilteringAPIService_List tests listing Data Filtering Profiles.
func Test_security_services_DataFilteringAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-df-list-" + common.GenerateRandomString(6)
	profile := security_services.DataFilteringProfiles{
		Folder: common.StringPtr("ngfw-shared"),
		Name:   common.StringPtr(profileName),
	}

	createRes, _, err := client.DataFilteringAPI.CreateDataFilteringProfiles(context.Background()).DataFilteringProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Data Filtering Profile for list test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	defer func() {
		t.Logf("Cleaning up Data Filtering Profile with ID: %s", *createdProfileID)
		_, errDel := client.DataFilteringAPI.DeleteDataFilteringProfilesByID(context.Background(), *createdProfileID).Execute()
		require.NoError(t, errDel, "Failed to delete Data Filtering Profile during cleanup")
	}()

	listRes, httpResList, errList := client.DataFilteringAPI.ListDataFilteringProfiles(context.Background()).Folder("ngfw-shared").Limit(100).Execute()
	require.NoError(t, errList, "Failed to list Data Filtering Profiles")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	foundObject := false
	for _, p := range listRes.Data {
		if p.Name != nil && *p.Name == profileName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created Data Filtering Profile should be found in the list")
}

// Test_security_services_DataFilteringAPIService_DeleteByID tests deleting a Data Filtering Profile by its ID.
func Test_security_services_DataFilteringAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	profileName := "test-df-delete-" + common.GenerateRandomString(6)
	profile := security_services.DataFilteringProfiles{
		Folder: common.StringPtr("ngfw-shared"),
		Name:   common.StringPtr(profileName),
	}

	createRes, _, err := client.DataFilteringAPI.CreateDataFilteringProfiles(context.Background()).DataFilteringProfiles(profile).Execute()
	require.NoError(t, err, "Failed to create Data Filtering Profile for delete test")
	createdProfileID := createRes.Id
	require.NotEmpty(t, createdProfileID, "Created profile ID should not be empty")

	httpResDel, errDel := client.DataFilteringAPI.DeleteDataFilteringProfilesByID(context.Background(), *createdProfileID).Execute()
	require.NoError(t, errDel, "Failed to delete Data Filtering Profile")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status on successful delete")
}
