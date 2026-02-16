/*
Objects Testing HIPObjectsAPIService
*/
package objects

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)

// Test_objects_HIPObjectsAPIService_Create tests the creation of a HIP object.
func Test_objects_HIPObjectsAPIService_Create(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	createdHipObjectName := "test-hip-obj-create-" + common.GenerateRandomString(6)
	hipObject := objects.HipObjects{
		Folder:      common.StringPtr("Shared"),
		Name:        createdHipObjectName,
		Description: common.StringPtr("Test HIP object for create API"),
		HostInfo: &objects.HipObjectsHostInfo{
			Criteria: objects.HipObjectsHostInfoCriteria{
				Os: &objects.HipObjectsHostInfoCriteriaOs{
					Contains: &objects.HipObjectsHostInfoCriteriaOsContains{
						Microsoft: common.StringPtr("Microsoft Windows 10"),
					},
				},
			},
		},
		AntiMalware: &objects.HipObjectsAntiMalware{
			Criteria: &objects.HipObjectsAntiMalwareCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
		DiskBackup: &objects.HipObjectsDiskBackup{
			Criteria: &objects.HipObjectsDiskBackupCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
		DiskEncryption: &objects.HipObjectsDiskEncryption{
			Criteria: &objects.HipObjectsDiskEncryptionCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
		MobileDevice: &objects.HipObjectsMobileDevice{
			Criteria: &objects.HipObjectsMobileDeviceCriteria{
				Jailbroken: common.BoolPtr(false),
			},
		},
		PatchManagement: &objects.HipObjectsPatchManagement{
			Criteria: &objects.HipObjectsPatchManagementCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
		DataLossPrevention: &objects.HipObjectsDataLossPrevention{
			Criteria: &objects.HipObjectsDataLossPreventionCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
	}

	fmt.Printf("Creating HIP object with name: %s\n", hipObject.Name)
	req := client.HIPObjectsAPI.CreateHIPObjects(context.Background()).HipObjects(hipObject)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create HIP object")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdHipObjectName, res.Name, "Created HIP object name should match")
	createdHipObjectID := res.Id

	defer func() {
		_, errDel := client.HIPObjectsAPI.DeleteHIPObjectsByID(context.Background(), createdHipObjectID).Execute()
		require.NoError(t, errDel, "Failed to delete HIP object during cleanup")
	}()

	t.Logf("Successfully created HIP object: %s with ID: %s", hipObject.Name, createdHipObjectID)
	assert.NotNil(t, res.HostInfo, "HostInfo should be present in the response")
	assert.NotNil(t, res.AntiMalware, "AntiMalware should be present in the response")
	assert.NotNil(t, res.DiskBackup, "DiskBackup should be present in the response")
	assert.NotNil(t, res.DiskEncryption, "DiskEncryption should be present in the response")
	assert.NotNil(t, res.MobileDevice, "MobileDevice should be present in the response")
	assert.NotNil(t, res.PatchManagement, "PatchManagement should be present in the response")
	assert.NotNil(t, res.DataLossPrevention, "DataLossPrevention should be present in the response")
	assert.True(t, *res.DataLossPrevention.Criteria.IsInstalled, "DataLossPrevention criteria IsInstalled should be true")
}

// Test_objects_HIPObjectsAPIService_GetByID tests retrieving a HIP object by its ID.
func Test_objects_HIPObjectsAPIService_GetByID(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	createdHipObjectName := "test-hip-obj-get-" + common.GenerateRandomString(6)
	hipObject := objects.HipObjects{
		Folder: common.StringPtr("Shared"),
		Name:   createdHipObjectName,
		HostInfo: &objects.HipObjectsHostInfo{
			Criteria: objects.HipObjectsHostInfoCriteria{
				Os: &objects.HipObjectsHostInfoCriteriaOs{
					Contains: &objects.HipObjectsHostInfoCriteriaOsContains{
						Apple: common.StringPtr("macOS"),
					},
				},
			},
		},
	}

	createRes, _, err := client.HIPObjectsAPI.CreateHIPObjects(context.Background()).HipObjects(hipObject).Execute()
	require.NoError(t, err, "Failed to create HIP object for get test")
	createdHipObjectID := createRes.Id
	require.NotEmpty(t, createdHipObjectID, "Created HIP object ID should not be empty")

	defer func() {
		_, errDel := client.HIPObjectsAPI.DeleteHIPObjectsByID(context.Background(), createdHipObjectID).Execute()
		require.NoError(t, errDel, "Failed to delete HIP object during cleanup")
	}()

	getRes, httpResGet, errGet := client.HIPObjectsAPI.GetHIPObjectsByID(context.Background(), createdHipObjectID).Execute()
	require.NoError(t, errGet, "Failed to get HIP object by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdHipObjectName, getRes.Name, "HIP object name should match")
	assert.Equal(t, "macOS", *getRes.HostInfo.Criteria.Os.Contains.Apple, "Host info OS should match")
}

// Test_objects_HIPObjectsAPIService_Update tests updating an existing HIP object.
func Test_objects_HIPObjectsAPIService_Update(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	createdHipObjectName := "test-hip-obj-update-" + common.GenerateRandomString(6)
	hipObject := objects.HipObjects{
		Folder: common.StringPtr("Shared"),
		Name:   createdHipObjectName,
	}

	createRes, _, err := client.HIPObjectsAPI.CreateHIPObjects(context.Background()).HipObjects(hipObject).Execute()
	require.NoError(t, err, "Failed to create HIP object for update test")
	createdHipObjectID := createRes.Id
	require.NotEmpty(t, createdHipObjectID, "Created HIP object ID should not be empty")

	defer func() {
		_, errDel := client.HIPObjectsAPI.DeleteHIPObjectsByID(context.Background(), createdHipObjectID).Execute()
		require.NoError(t, errDel, "Failed to delete HIP object during cleanup")
	}()

	updatedHipObject := objects.HipObjects{
		Name:        createdHipObjectName,
		Description: common.StringPtr("Updated with all criteria"),
		HostInfo: &objects.HipObjectsHostInfo{
			Criteria: objects.HipObjectsHostInfoCriteria{
				Os: &objects.HipObjectsHostInfoCriteriaOs{
					Contains: &objects.HipObjectsHostInfoCriteriaOsContains{
						Linux: common.StringPtr("RedHat"),
					},
				},
			},
		},
		AntiMalware: &objects.HipObjectsAntiMalware{
			Criteria: &objects.HipObjectsAntiMalwareCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
		DiskBackup: &objects.HipObjectsDiskBackup{
			Criteria: &objects.HipObjectsDiskBackupCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
		DiskEncryption: &objects.HipObjectsDiskEncryption{
			Criteria: &objects.HipObjectsDiskEncryptionCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
		MobileDevice: &objects.HipObjectsMobileDevice{
			Criteria: &objects.HipObjectsMobileDeviceCriteria{
				Jailbroken: common.BoolPtr(false),
			},
		},
		PatchManagement: &objects.HipObjectsPatchManagement{
			Criteria: &objects.HipObjectsPatchManagementCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
		DataLossPrevention: &objects.HipObjectsDataLossPrevention{
			Criteria: &objects.HipObjectsDataLossPreventionCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
	}

	updateRes, httpResUpdate, errUpdate := client.HIPObjectsAPI.UpdateHIPObjectsByID(context.Background(), createdHipObjectID).HipObjects(updatedHipObject).Execute()
	require.NoError(t, errUpdate, "Failed to update HIP object")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, "Updated with all criteria", *updateRes.Description, "Description should be updated")
	assert.Equal(t, "RedHat", *updateRes.HostInfo.Criteria.Os.Contains.Linux, "Host info OS should be updated")
	require.NotNil(t, updateRes.DataLossPrevention, "DataLossPrevention should be present in the response")
	assert.Equal(t, true, *updateRes.DataLossPrevention.Criteria.IsInstalled, "DLP IsEnabled should be updated")
}

// Test_objects_HIPObjectsAPIService_List tests listing HIP objects.
func Test_objects_HIPObjectsAPIService_List(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	createdHipObjectName := "test-hip-obj-list-" + common.GenerateRandomString(6)
	hipObject := objects.HipObjects{
		Folder: common.StringPtr("Shared"),
		Name:   createdHipObjectName,
	}

	createRes, _, err := client.HIPObjectsAPI.CreateHIPObjects(context.Background()).HipObjects(hipObject).Execute()
	require.NoError(t, err, "Failed to create HIP object for list test")
	createdHipObjectID := createRes.Id
	require.NotEmpty(t, createdHipObjectID, "Created HIP object ID should not be empty")

	defer func() {
		_, errDel := client.HIPObjectsAPI.DeleteHIPObjectsByID(context.Background(), createdHipObjectID).Execute()
		require.NoError(t, errDel, "Failed to delete HIP object during cleanup")
	}()

	listRes, httpResList, errList := client.HIPObjectsAPI.ListHIPObjects(context.Background()).Folder("Shared").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list HIP objects")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	foundObject := false
	for _, obj := range listRes.Data {
		if obj.Name == createdHipObjectName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created HIP object should be found in the list")
}

// Test_objects_HIPObjectsAPIService_DeleteByID tests deleting a HIP object by its ID.
func Test_objects_HIPObjectsAPIService_DeleteByID(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	createdHipObjectName := "test-hip-obj-delete-" + common.GenerateRandomString(6)
	hipObject := objects.HipObjects{
		Folder: common.StringPtr("Shared"),
		Name:   createdHipObjectName,
	}

	createRes, _, err := client.HIPObjectsAPI.CreateHIPObjects(context.Background()).HipObjects(hipObject).Execute()
	require.NoError(t, err, "Failed to create HIP object for delete test")
	createdHipObjectID := createRes.Id
	require.NotEmpty(t, createdHipObjectID, "Created HIP object ID should not be empty")

	httpResDel, errDel := client.HIPObjectsAPI.DeleteHIPObjectsByID(context.Background(), createdHipObjectID).Execute()
	require.NoError(t, errDel, "Failed to delete HIP object")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}

// Test_objects_HIPObjectsAPIService_FetchHIPObjects tests the FetchHIPObjects convenience method
func Test_objects_HIPObjectsAPIService_FetchHIPObjects(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create test object using same payload as Create test
	testName := "test-hip-obj-fetch-" + common.GenerateRandomString(6)
	testObj := objects.HipObjects{
		Folder:      common.StringPtr("Shared"),
		Name:        testName,
		Description: common.StringPtr("Test HIP object for fetch API"),
		HostInfo: &objects.HipObjectsHostInfo{
			Criteria: objects.HipObjectsHostInfoCriteria{
				Os: &objects.HipObjectsHostInfoCriteriaOs{
					Contains: &objects.HipObjectsHostInfoCriteriaOsContains{
						Microsoft: common.StringPtr("Microsoft Windows 10"),
					},
				},
			},
		},
		AntiMalware: &objects.HipObjectsAntiMalware{
			Criteria: &objects.HipObjectsAntiMalwareCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
		DiskBackup: &objects.HipObjectsDiskBackup{
			Criteria: &objects.HipObjectsDiskBackupCriteria{
				IsInstalled: common.BoolPtr(true),
			},
		},
	}

	createReq := client.HIPObjectsAPI.CreateHIPObjects(context.Background()).HipObjects(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.HIPObjectsAPI.DeleteHIPObjectsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.HIPObjectsAPI.FetchHIPObjects(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch hip_objects by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchHIPObjects found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.HIPObjectsAPI.FetchHIPObjects(
		context.Background(),
		"non-existent-hip_objects-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchHIPObjects correctly returned nil for non-existent object")
}
