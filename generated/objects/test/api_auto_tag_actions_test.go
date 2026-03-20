/*
Objects Testing AutoTagActionsAPIService
Note: This API uses name-based operations for delete (not ID-based)
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

// Test_objects_AutoTagActionsAPIService_Create tests the creation of an auto tag action
func Test_objects_AutoTagActionsAPIService_Create(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	testName := "test-autotag-create-" + common.GenerateRandomString(10)
	autoTag := objects.AutoTagActions{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Filter:  "addr.src in 10.0.0.0/8",
		LogType: "traffic",
		Actions: []objects.AutoTagActionsActionsInner{
			{
				Name: "test-action",
				Type: objects.AutoTagActionsActionsInnerType{
					Tagging: objects.AutoTagActionsActionsInnerTypeTagging{
						Action: "add-tag",
						Tags:   []string{"test-tag"},
						Target: "source-address",
					},
				},
			},
		},
	}

	req := client.AutoTagActionsAPI.CreateAutoTagActions(context.Background()).AutoTagActions(autoTag)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create auto tag action")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, testName, res.Name, "Created auto tag action name should match")

	t.Logf("Successfully created auto tag action: %s", testName)

	// Cleanup - uses name-based delete
	reqDel := client.AutoTagActionsAPI.DeleteAutoTagActions(context.Background()).Name(testName)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete auto tag action during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
	t.Logf("Successfully cleaned up auto tag action: %s", testName)
}

// Test_objects_AutoTagActionsAPIService_List tests listing auto tag actions
func Test_objects_AutoTagActionsAPIService_List(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	testName := "test-autotag-list-" + common.GenerateRandomString(10)
	autoTag := objects.AutoTagActions{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Filter:  "addr.src in 172.16.0.0/12",
		LogType: "traffic",
		Actions: []objects.AutoTagActionsActionsInner{
			{
				Name: "list-action",
				Type: objects.AutoTagActionsActionsInnerType{
					Tagging: objects.AutoTagActionsActionsInnerTypeTagging{
						Action: "add-tag",
						Tags:   []string{"list-tag"},
						Target: "source-address",
					},
				},
			},
		},
	}

	req := client.AutoTagActionsAPI.CreateAutoTagActions(context.Background()).AutoTagActions(autoTag)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create auto tag action for list test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Test List
	reqList := client.AutoTagActionsAPI.ListAutoTagActions(context.Background())
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	require.NoError(t, errList, "Failed to list auto tag actions")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one auto tag action in the list")

	t.Logf("Successfully listed auto tag actions, count: %d", len(listRes.Data))

	// Cleanup
	reqDel := client.AutoTagActionsAPI.DeleteAutoTagActions(context.Background()).Name(testName)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete auto tag action during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
	t.Logf("Successfully cleaned up auto tag action: %s", testName)
}

// Test_objects_AutoTagActionsAPIService_Update tests updating an existing auto tag action
func Test_objects_AutoTagActionsAPIService_Update(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	testName := "test-autotag-update-" + common.GenerateRandomString(10)
	autoTag := objects.AutoTagActions{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Filter:  "addr.src in 192.168.0.0/16",
		LogType: "traffic",
		Actions: []objects.AutoTagActionsActionsInner{
			{
				Name: "update-action",
				Type: objects.AutoTagActionsActionsInnerType{
					Tagging: objects.AutoTagActionsActionsInnerTypeTagging{
						Action: "add-tag",
						Tags:   []string{"update-tag"},
						Target: "source-address",
					},
				},
			},
		},
	}

	req := client.AutoTagActionsAPI.CreateAutoTagActions(context.Background()).AutoTagActions(autoTag)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create auto tag action for update test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Test Update
	updatedAutoTag := objects.AutoTagActions{
		Name:        testName,
		Folder:      common.StringPtr("Prisma Access"),
		Filter:      "addr.src in 192.168.0.0/16",
		LogType:     "traffic",
		Description: common.StringPtr("Updated description"),
		Actions: []objects.AutoTagActionsActionsInner{
			{
				Name: "updated-action",
				Type: objects.AutoTagActionsActionsInnerType{
					Tagging: objects.AutoTagActionsActionsInnerTypeTagging{
						Action: "add-tag",
						Tags:   []string{"updated-tag"},
						Target: "source-address",
					},
				},
			},
		},
	}

	reqUpdate := client.AutoTagActionsAPI.UpdateAutoTagActions(context.Background()).AutoTagActions(updatedAutoTag)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	require.NoError(t, errUpdate, "Failed to update auto tag action")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, testName, updateRes.Name, "Auto tag action name should remain the same")
	assert.Equal(t, common.StringPtr("Updated description"), updateRes.Description, "Description should be updated")

	t.Logf("Successfully updated auto tag action: %s", testName)

	// Cleanup
	reqDel := client.AutoTagActionsAPI.DeleteAutoTagActions(context.Background()).Name(testName)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete auto tag action during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
	t.Logf("Successfully cleaned up auto tag action: %s", testName)
}

// Test_objects_AutoTagActionsAPIService_Delete tests deleting an auto tag action by name
func Test_objects_AutoTagActionsAPIService_Delete(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	testName := "test-autotag-delete-" + common.GenerateRandomString(10)
	autoTag := objects.AutoTagActions{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Filter:  "addr.src in 10.10.0.0/16",
		LogType: "traffic",
		Actions: []objects.AutoTagActionsActionsInner{
			{
				Name: "delete-action",
				Type: objects.AutoTagActionsActionsInnerType{
					Tagging: objects.AutoTagActionsActionsInnerTypeTagging{
						Action: "add-tag",
						Tags:   []string{"delete-tag"},
						Target: "source-address",
					},
				},
			},
		},
	}

	req := client.AutoTagActionsAPI.CreateAutoTagActions(context.Background()).AutoTagActions(autoTag)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create auto tag action for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Test Delete by name
	reqDel := client.AutoTagActionsAPI.DeleteAutoTagActions(context.Background()).Name(testName)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	require.NoError(t, errDel, "Failed to delete auto tag action")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted auto tag action: %s", testName)
}

// Test_objects_AutoTagActionsAPIService_Fetch tests the Fetch convenience method
func Test_objects_AutoTagActionsAPIService_Fetch(t *testing.T) {
	client := SetupObjectSvcTestClient(t)

	testName := "test-autotag-fetch-" + common.GenerateRandomString(10)
	autoTag := objects.AutoTagActions{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Filter:  "addr.src in 10.20.0.0/16",
		LogType: "traffic",
		Actions: []objects.AutoTagActionsActionsInner{
			{
				Name: "fetch-action",
				Type: objects.AutoTagActionsActionsInnerType{
					Tagging: objects.AutoTagActionsActionsInnerTypeTagging{
						Action: "add-tag",
						Tags:   []string{"fetch-tag"},
						Target: "source-address",
					},
				},
			},
		},
	}

	req := client.AutoTagActionsAPI.CreateAutoTagActions(context.Background()).AutoTagActions(autoTag)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create auto tag action for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")

	// Cleanup after test
	defer func() {
		deleteReq := client.AutoTagActionsAPI.DeleteAutoTagActions(context.Background()).Name(testName)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", testName)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.AutoTagActionsAPI.FetchAutoTagActions(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch auto tag action by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchAutoTagActions found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.AutoTagActionsAPI.FetchAutoTagActions(
		context.Background(),
		"non-existent-autotag-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchAutoTagActions correctly returned nil for non-existent object")
}
