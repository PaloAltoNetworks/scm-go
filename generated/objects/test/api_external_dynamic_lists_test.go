/*
Objects Testing ExternalDynamicListsAPIService
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

// Test_objects_ExternalDynamicListsAPIService_Create tests the creation of an EDL object.
// This test creates a new EDL and then deletes it to ensure proper cleanup.
func Test_objects_ExternalDynamicListsAPIService_Create(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a valid EDL object with a unique name to avoid conflicts.
	createdEdlName := "test-edl-create-" + common.GenerateRandomString(6)
	edl := objects.ExternalDynamicLists{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdEdlName,
		Type: &objects.ExternalDynamicListsType{
			Domain: &objects.ExternalDynamicListsTypeDomain{
				Url:         "http://example.com/domains.txt",
				Description: common.StringPtr("Test EDL for create API"),
				Recurring: objects.ExternalDynamicListsTypeDomainRecurring{
					Daily: &objects.ExternalDynamicListsTypeDomainRecurringDaily{
						At: "03",
					},
				},
			},
		},
	}

	t.Logf("Creating EDL with name: %s", edl.Name)

	// Make the create request to the API.
	req := client.ExternalDynamicListsAPI.CreateExternalDynamicLists(context.Background()).ExternalDynamicLists(edl)
	res, httpRes, err := req.Execute()
	if err == nil {
		// Use the dereference operator (*) since Id is now a pointer.
		t.Logf("Successfully created EDL: %s with ID: %s", edl.Name, *res.Id)
	}
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create EDL")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties.
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdEdlName, res.Name, "Created EDL name should match")
	assert.True(t, *res.Folder == "Shared" || *res.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	require.NotNil(t, res.Type.Domain, "EDL type should be domain")
	assert.Equal(t, "http://example.com/domains.txt", res.Type.Domain.Url)
	require.NotNil(t, res.Id, "Created EDL should have an ID")
	assert.NotEmpty(t, *res.Id, "Created EDL ID should not be empty")

	// Use the ID from the response object.
	createdEdlID := *res.Id

	// Cleanup: Delete the created EDL to maintain test isolation.
	reqDel := client.ExternalDynamicListsAPI.DeleteExternalDynamicListsByID(context.Background(), createdEdlID)
	httpResDel, errDel := reqDel.Execute()
	if errDel == nil {
		t.Logf("Successfully cleaned up EDL: %s", createdEdlID)
	}
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete EDL during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")
}

// Test_objects_ExternalDynamicListsAPIService_GetByID tests retrieving an EDL by its ID.
// This test creates an EDL, retrieves it by ID, then deletes it.
func Test_objects_ExternalDynamicListsAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create an EDL first to have something to retrieve.
	createdEdlName := "test-edl-get-" + common.GenerateRandomString(6)
	edl := objects.ExternalDynamicLists{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdEdlName,
		Type: &objects.ExternalDynamicListsType{
			Ip: &objects.ExternalDynamicListsTypeIp{
				Url: "http://example.com/ips.txt",
				Recurring: objects.ExternalDynamicListsTypeIpRecurring{
					FiveMinute: make(map[string]interface{}),
				},
			},
		},
	}

	// Create the EDL via API.
	req := client.ExternalDynamicListsAPI.CreateExternalDynamicLists(context.Background()).ExternalDynamicLists(edl)
	createRes, _, err := req.Execute()
	if err == nil {
		t.Logf("Successfully created EDL for get test: %s", createdEdlName)
	}
	require.NoError(t, err, "Failed to create EDL for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created EDL should have an ID")
	createdEdlID := *createRes.Id
	require.NotEmpty(t, createdEdlID, "Created EDL ID should not be empty")

	// Test Get by ID operation.
	reqGetById := client.ExternalDynamicListsAPI.GetExternalDynamicListsByID(context.Background(), createdEdlID)
	getRes, httpResGet, errGet := reqGetById.Execute()
	if errGet == nil {
		t.Logf("Successfully retrieved EDL: %s", getRes.Name)
	}
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful.
	require.NoError(t, errGet, "Failed to get EDL by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties.
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdEdlName, getRes.Name, "EDL name should match")
	assert.True(t, *getRes.Folder == "Shared" || *getRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	require.NotNil(t, getRes.Type.Ip, "EDL type should be IP")
	assert.Equal(t, "http://example.com/ips.txt", getRes.Type.Ip.Url)
	require.NotNil(t, getRes.Id, "EDL ID should not be nil")
	assert.Equal(t, createdEdlID, *getRes.Id, "EDL ID should match")

	// Cleanup: Delete the created EDL.
	reqDel := client.ExternalDynamicListsAPI.DeleteExternalDynamicListsByID(context.Background(), createdEdlID)
	_, errDel := reqDel.Execute()
	if errDel == nil {
		t.Logf("Successfully cleaned up EDL: %s", createdEdlID)
	}
	require.NoError(t, errDel, "Failed to delete EDL during cleanup")
}

// Test_objects_ExternalDynamicListsAPIService_Update tests updating an existing EDL.
// This test creates an EDL, updates it, then deletes it.
func Test_objects_ExternalDynamicListsAPIService_Update(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create an EDL first to have something to update.
	createdEdlName := "test-edl-update-" + common.GenerateRandomString(6)
	edl := objects.ExternalDynamicLists{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdEdlName,
		Type: &objects.ExternalDynamicListsType{
			Url: &objects.ExternalDynamicListsTypeUrl{
				Url: "http://example.com/initial-urls.txt",
				Recurring: objects.ExternalDynamicListsTypeUrlRecurring{
					Hourly: make(map[string]interface{}),
				},
			},
		},
	}

	// Create the EDL via API.
	req := client.ExternalDynamicListsAPI.CreateExternalDynamicLists(context.Background()).ExternalDynamicLists(edl)
	createRes, _, err := req.Execute()
	if err == nil {
		t.Logf("Successfully created EDL for update test: %s", createdEdlName)
	}
	require.NoError(t, err, "Failed to create EDL for update test")
	require.NotNil(t, createRes.Id, "Created EDL should have an ID")
	createdEdlID := *createRes.Id
	require.NotEmpty(t, createdEdlID, "Created EDL ID should not be empty")

	// Test Update operation with modified fields.
	updatedEdl := objects.ExternalDynamicLists{
		Name: createdEdlName, // Name must be included in update payload.
		Type: &objects.ExternalDynamicListsType{
			Url: &objects.ExternalDynamicListsTypeUrl{
				Description: common.StringPtr("Updated URL description"),
				Url:         "http://example.com/updated-urls.txt",
				Recurring: objects.ExternalDynamicListsTypeUrlRecurring{
					Weekly: &objects.ExternalDynamicListsTypeUrlRecurringWeekly{
						DayOfWeek: "sunday",
						At:        "14",
					},
				},
			},
		},
	}

	reqUpdate := client.ExternalDynamicListsAPI.UpdateExternalDynamicListsByID(context.Background(), createdEdlID).ExternalDynamicLists(updatedEdl)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate == nil {
		t.Logf("Successfully updated EDL: %s", createdEdlName)
	}
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful.
	require.NoError(t, errUpdate, "Failed to update EDL")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties.
	require.NotNil(t, updateRes, "Update response should not be nil")
	require.NotNil(t, updateRes.Type.Url, "EDL type should be URL")
	assert.Equal(t, common.StringPtr("Updated URL description"), updateRes.Type.Url.Description, "Description should be updated")
	assert.Equal(t, "http://example.com/updated-urls.txt", updateRes.Type.Url.Url, "URL should be updated")
	require.NotNil(t, updateRes.Type.Url.Recurring.Weekly, "Recurring schedule should be updated to weekly")
	assert.Equal(t, "sunday", updateRes.Type.Url.Recurring.Weekly.DayOfWeek, "Day of week should be updated")
	assert.Equal(t, "14", updateRes.Type.Url.Recurring.Weekly.At, "Time should be updated")
	require.NotNil(t, updateRes.Id, "EDL ID should not be nil")
	assert.Equal(t, createdEdlID, *updateRes.Id, "EDL ID should remain the same")

	// Cleanup: Delete the created EDL.
	reqDel := client.ExternalDynamicListsAPI.DeleteExternalDynamicListsByID(context.Background(), createdEdlID)
	_, errDel := reqDel.Execute()
	if errDel == nil {
		t.Logf("Successfully cleaned up EDL: %s", createdEdlID)
	}
	require.NoError(t, errDel, "Failed to delete EDL during cleanup")
}

// Test_objects_ExternalDynamicListsAPIService_List tests listing EDLs.
// This test creates an EDL, lists EDLs to verify it's included, then deletes it.
func Test_objects_ExternalDynamicListsAPIService_List(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create an EDL first to have something to list.
	createdEdlName := "test-edl-list-" + common.GenerateRandomString(6)
	edl := objects.ExternalDynamicLists{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdEdlName,
		Type: &objects.ExternalDynamicListsType{
			Domain: &objects.ExternalDynamicListsTypeDomain{
				Url: "http://example.com/list-test.txt",
				Recurring: objects.ExternalDynamicListsTypeDomainRecurring{
					Hourly: make(map[string]interface{}),
				},
			},
		},
	}

	// Create the EDL via API.
	req := client.ExternalDynamicListsAPI.CreateExternalDynamicLists(context.Background()).ExternalDynamicLists(edl)
	createRes, _, err := req.Execute()
	if err == nil {
		t.Logf("Successfully created EDL for list test: %s", createdEdlName)
	}
	require.NoError(t, err, "Failed to create EDL for list test")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created EDL should have an ID")
	createdEdlID := *createRes.Id
	require.NotEmpty(t, createdEdlID, "Created EDL ID should not be empty")

	// Test List operation.
	reqList := client.ExternalDynamicListsAPI.ListExternalDynamicLists(context.Background()).Folder("Prisma Access")
	listRes, httpResList, errList := reqList.Execute()
	if errList == nil {
		t.Logf("Successfully listed EDLs")
	}
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful.
	require.NoError(t, errList, "Failed to list EDLs")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one EDL in the list")

	// Verify our created EDL is in the list by iterating through the results.
	foundEdl := false
	for _, edlFromList := range listRes.Data {
		if edlFromList.Id != nil && *edlFromList.Id == createdEdlID {
			foundEdl = true
			assert.Equal(t, "http://example.com/list-test.txt", edlFromList.Type.Domain.Url)
			break
		}
	}
	assert.True(t, foundEdl, "Created EDL should be found in the list")

	// Cleanup: Delete the created EDL.
	reqDel := client.ExternalDynamicListsAPI.DeleteExternalDynamicListsByID(context.Background(), createdEdlID)
	_, errDel := reqDel.Execute()
	if errDel == nil {
		t.Logf("Successfully cleaned up EDL: %s", createdEdlID)
	}
	require.NoError(t, errDel, "Failed to delete EDL during cleanup")
}

// Test_objects_ExternalDynamicListsAPIService_DeleteByID tests deleting an EDL by its ID.
// This test creates an EDL, deletes it, then verifies the deletion was successful.
func Test_objects_ExternalDynamicListsAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create an EDL first to have something to delete.
	createdEdlName := "test-edl-delete-" + common.GenerateRandomString(6)
	edl := objects.ExternalDynamicLists{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdEdlName,
		Type: &objects.ExternalDynamicListsType{
			Domain: &objects.ExternalDynamicListsTypeDomain{
				Url: "http://example.com/delete-me.txt",
				Recurring: objects.ExternalDynamicListsTypeDomainRecurring{
					Hourly: make(map[string]interface{}),
				},
			},
		},
	}

	// Create the EDL via API.
	req := client.ExternalDynamicListsAPI.CreateExternalDynamicLists(context.Background()).ExternalDynamicLists(edl)
	createRes, _, err := req.Execute()
	if err == nil {
		t.Logf("Successfully created EDL for delete test: %s", createdEdlName)
	}
	require.NoError(t, err, "Failed to create EDL for delete test")
	require.NotNil(t, createRes.Id, "Created EDL should have an ID")
	createdEdlID := *createRes.Id
	require.NotEmpty(t, createdEdlID, "Created EDL ID should not be empty")

	// Test Delete by ID operation.
	reqDel := client.ExternalDynamicListsAPI.DeleteExternalDynamicListsByID(context.Background(), createdEdlID)
	httpResDel, errDel := reqDel.Execute()
	if errDel == nil {
		t.Logf("Successfully deleted EDL: %s", createdEdlID)
	}
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful.
	require.NoError(t, errDel, "Failed to delete EDL")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}
