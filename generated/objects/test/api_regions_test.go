/*
Objects Testing RegionsAPIService
*/
package objects

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)

// Test_objects_RegionsAPIService_Create tests the creation of a region object
func Test_objects_RegionsAPIService_Create(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	testName := "test-region-create-" + common.GenerateRandomString(10)
	region := objects.Regions{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Address: []string{"10.0.0.0/8"},
	}

	req := client.RegionsAPI.CreateRegions(context.Background()).Regions(region)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create region")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, testName, res.Name, "Created region name should match")
	assert.NotEmpty(t, res.Id, "Created region should have an ID")

	createdID := res.Id
	t.Logf("Successfully created region: %s with ID: %s", testName, createdID)

	// Cleanup
	reqDel := client.RegionsAPI.DeleteRegionsByID(context.Background(), createdID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete region during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
	t.Logf("Successfully cleaned up region: %s", createdID)
}

// Test_objects_RegionsAPIService_GetByID tests retrieving a region by its ID
func Test_objects_RegionsAPIService_GetByID(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	testName := "test-region-getbyid-" + common.GenerateRandomString(10)
	region := objects.Regions{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Address: []string{"172.16.0.0/12"},
	}

	req := client.RegionsAPI.CreateRegions(context.Background()).Regions(region)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create region for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id
	require.NotEmpty(t, createdID, "Created region ID should not be empty")

	// Test Get by ID
	reqGet := client.RegionsAPI.GetRegionsByID(context.Background(), createdID)
	getRes, httpResGet, errGet := reqGet.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	require.NoError(t, errGet, "Failed to get region by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, testName, getRes.Name, "Region name should match")
	assert.Equal(t, createdID, getRes.Id, "Region ID should match")

	t.Logf("Successfully retrieved region: %s", getRes.Name)

	// Cleanup
	reqDel := client.RegionsAPI.DeleteRegionsByID(context.Background(), createdID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete region during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
	t.Logf("Successfully cleaned up region: %s", createdID)
}

// Test_objects_RegionsAPIService_Update tests updating an existing region
func Test_objects_RegionsAPIService_Update(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	testName := "test-region-update-" + common.GenerateRandomString(10)
	region := objects.Regions{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Address: []string{"192.168.0.0/16"},
	}

	req := client.RegionsAPI.CreateRegions(context.Background()).Regions(region)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create region for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id
	require.NotEmpty(t, createdID, "Created region ID should not be empty")

	// Test Update
	updatedRegion := objects.Regions{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Address: []string{"192.168.0.0/16", "10.10.0.0/16"},
	}

	reqUpdate := client.RegionsAPI.UpdateRegionsByID(context.Background(), createdID).Regions(updatedRegion)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	require.NoError(t, errUpdate, "Failed to update region")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, testName, updateRes.Name, "Region name should remain the same")
	assert.Equal(t, 2, len(updateRes.Address), "Region should have 2 addresses after update")

	t.Logf("Successfully updated region: %s", testName)

	// Cleanup
	reqDel := client.RegionsAPI.DeleteRegionsByID(context.Background(), createdID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete region during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
	t.Logf("Successfully cleaned up region: %s", createdID)
}

// Test_objects_RegionsAPIService_DeleteByID tests deleting a region by its ID
func Test_objects_RegionsAPIService_DeleteByID(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	testName := "test-region-delete-" + common.GenerateRandomString(10)
	region := objects.Regions{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Address: []string{"10.200.0.0/16"},
	}

	req := client.RegionsAPI.CreateRegions(context.Background()).Regions(region)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create region for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id
	require.NotEmpty(t, createdID, "Created region ID should not be empty")

	// Test Delete by ID
	reqDel := client.RegionsAPI.DeleteRegionsByID(context.Background(), createdID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	require.NoError(t, errDel, "Failed to delete region")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted region: %s", createdID)
}
