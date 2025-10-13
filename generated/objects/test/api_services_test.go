/*
Objects Testing ServicesAPIService
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

// Test_objects_ServicesAPIService_Create tests the creation of a TCP service object.
// This test creates a new service and then deletes it to ensure proper cleanup.
func Test_objects_ServicesAPIService_CreateService(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create tags for the service.
	tag1Name := "tag1-" + common.GenerateRandomString(4)
	tag2Name := "tag2-" + common.GenerateRandomString(4)
	tag1 := objects.Tags{Folder: common.StringPtr("Prisma Access"), Name: tag1Name}
	tag2 := objects.Tags{Folder: common.StringPtr("Prisma Access"), Name: tag2Name}

	// Create tag1
	createTag1Req := client.TagsAPI.CreateTags(context.Background()).Tags(tag1)
	tag1Res, _, err := createTag1Req.Execute()
	require.NoError(t, err, "Failed to create tag1 for service test")
	t.Cleanup(func() {
		delReq := client.TagsAPI.DeleteTagsByID(context.Background(), *tag1Res.Id)
		_, err := delReq.Execute()
		require.NoError(t, err, "Cleanup failed for tag1")
	})

	// Create tag2
	createTag2Req := client.TagsAPI.CreateTags(context.Background()).Tags(tag2)
	tag2Res, _, err := createTag2Req.Execute()
	require.NoError(t, err, "Failed to create tag2 for service test")
	t.Cleanup(func() {
		delReq := client.TagsAPI.DeleteTagsByID(context.Background(), *tag2Res.Id)
		_, err := delReq.Execute()
		require.NoError(t, err, "Cleanup failed for tag2")
	})

	// Create a valid TCP service object with a unique name.
	createdSvcName := "test-tcp-create-" + common.GenerateRandomString(6)
	service := objects.Services{
		Folder:      common.StringPtr("Prisma Access"),
		Name:        createdSvcName,
		Description: common.StringPtr("Test TCP service for create API"),
		Protocol: &objects.ServicesProtocol{
			Tcp: &objects.ServicesProtocolTcp{
				Port:       "1024-1026",
				SourcePort: common.StringPtr("1024"),
			},
		},
		Tag: []string{tag1Name, tag2Name},
	}

	fmt.Printf("Creating service with name: %s\n", service.Name)

	// Make the create request to the API.
	req := client.ServicesAPI.CreateServices(context.Background()).Services(service)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful.
	require.NoError(t, err, "Failed to create service")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties.
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdSvcName, res.Name, "Created service name should match")
	assert.True(t, *res.Folder == "Shared" || *res.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	require.NotNil(t, res.Protocol.Tcp, "Protocol TCP data should not be nil")
	assert.Equal(t, "1024-1026", res.Protocol.Tcp.Port, "TCP port should match")
	assert.ElementsMatch(t, []string{tag1Name, tag2Name}, res.Tag, "Tags should match")
	require.NotNil(t, res.Id, "Created service ID should not be nil")
	assert.NotEmpty(t, *res.Id, "Created service ID should not be empty")

	createdSvcID := *res.Id
	t.Logf("Successfully created service: %s with ID: %s", service.Name, createdSvcID)

	// Cleanup: Delete the created service.
	reqDel := client.ServicesAPI.DeleteServicesByID(context.Background(), createdSvcID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete service during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up service: %s", createdSvcID)
}

// Test_objects_ServicesAPIService_GetByID tests retrieving a service by its ID.
func Test_objects_ServicesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a UDP service to retrieve.
	createdSvcName := "test-udp-get-" + common.GenerateRandomString(6)
	service := objects.Services{
		Folder:      common.StringPtr("Prisma Access"),
		Name:        createdSvcName,
		Description: common.StringPtr("Test UDP service for get API"),
		Protocol: &objects.ServicesProtocol{
			Udp: &objects.ServicesProtocolUdp{
				Port: "53, 55",
			},
		},
	}

	// Create the service via API.
	req := client.ServicesAPI.CreateServices(context.Background()).Services(service)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create service for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	require.NotNil(t, createRes.Id, "Created service ID should not be nil")
	createdSvcID := *createRes.Id
	require.NotEmpty(t, createdSvcID, "Created service ID should not be empty")

	// Test Get by ID operation.
	reqGetById := client.ServicesAPI.GetServicesByID(context.Background(), createdSvcID)
	getRes, httpResGet, errGet := reqGetById.Execute()
	if errGet != nil {
		handleAPIError(errGet)
	}

	// Verify the get operation was successful.
	require.NoError(t, errGet, "Failed to get service by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties.
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdSvcName, getRes.Name, "Service name should match")
	require.NotNil(t, getRes.Protocol.Udp, "Protocol UDP data should not be nil")
	assert.Equal(t, "53, 55", getRes.Protocol.Udp.Port, "UDP port should match")
	require.NotNil(t, getRes.Id, "Service ID should not be nil")
	assert.Equal(t, createdSvcID, *getRes.Id, "Service ID should match")

	t.Logf("Successfully retrieved service: %s", getRes.Name)

	// Cleanup: Delete the created service.
	reqDel := client.ServicesAPI.DeleteServicesByID(context.Background(), createdSvcID)
	_, errDel := reqDel.Execute()
	require.NoError(t, errDel, "Failed to delete service during cleanup")

	t.Logf("Successfully cleaned up service: %s", createdSvcID)
}

// Test_objects_ServicesAPIService_Update tests updating an existing service.
func Test_objects_ServicesAPIService_Update(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a service to update.
	createdSvcName := "test-svc-update-" + common.GenerateRandomString(6)
	service := objects.Services{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdSvcName,
		Protocol: &objects.ServicesProtocol{
			Tcp: &objects.ServicesProtocolTcp{
				Port: "3389",
			},
		},
	}

	// Create the service via API.
	req := client.ServicesAPI.CreateServices(context.Background()).Services(service)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create service for update test")
	require.NotNil(t, createRes.Id, "Created service ID should not be nil")
	createdSvcID := *createRes.Id
	require.NotEmpty(t, createdSvcID, "Created service ID should not be empty")

	// Create tags for the update.
	updateTag1Name := "corp-" + common.GenerateRandomString(4)
	updateTag2Name := "remote-access-" + common.GenerateRandomString(4)
	tag1 := objects.Tags{Folder: common.StringPtr("Prisma Access"), Name: updateTag1Name}
	tag2 := objects.Tags{Folder: common.StringPtr("Prisma Access"), Name: updateTag2Name}

	// Create tag1 for update
	createTag1Req := client.TagsAPI.CreateTags(context.Background()).Tags(tag1)
	tag1Res, _, err := createTag1Req.Execute()
	require.NoError(t, err, "Failed to create update tag1 for service test")
	t.Cleanup(func() {
		delReq := client.TagsAPI.DeleteTagsByID(context.Background(), *tag1Res.Id)
		_, err := delReq.Execute()
		require.NoError(t, err, "Cleanup failed for update tag1")
	})

	// Create tag2 for update
	createTag2Req := client.TagsAPI.CreateTags(context.Background()).Tags(tag2)
	tag2Res, _, err := createTag2Req.Execute()
	require.NoError(t, err, "Failed to create update tag2 for service test")
	t.Cleanup(func() {
		delReq := client.TagsAPI.DeleteTagsByID(context.Background(), *tag2Res.Id)
		_, err := delReq.Execute()
		require.NoError(t, err, "Cleanup failed for update tag2")
	})

	// Test Update operation with modified fields.
	updatedService := objects.Services{
		Name:        createdSvcName,
		Description: common.StringPtr("Updated RDP service"),
		Protocol: &objects.ServicesProtocol{
			Tcp: &objects.ServicesProtocolTcp{
				Port: "3389",
				Override: &objects.ServicesProtocolTcpOverride{
					Timeout: common.Int32Ptr(7200), // 2 hours
				},
			},
		},
		Tag: []string{updateTag1Name, updateTag2Name},
	}

	reqUpdate := client.ServicesAPI.UpdateServicesByID(context.Background(), createdSvcID).Services(updatedService)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful.
	require.NoError(t, errUpdate, "Failed to update service")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties.
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, common.StringPtr("Updated RDP service"), updateRes.Description, "Description should be updated")
	assert.ElementsMatch(t, []string{updateTag1Name, updateTag2Name}, updateRes.Tag, "Tags should be updated")
	require.NotNil(t, updateRes.Protocol.Tcp.Override, "TCP override should not be nil")
	assert.Equal(t, int32(7200), *updateRes.Protocol.Tcp.Override.Timeout, "TCP override timeout should be updated")
	require.NotNil(t, updateRes.Id, "Service ID should not be nil")
	assert.Equal(t, createdSvcID, *updateRes.Id, "Service ID should remain the same")

	t.Logf("Successfully updated service: %s", createdSvcName)

	// Cleanup: Delete the created service.
	reqDel := client.ServicesAPI.DeleteServicesByID(context.Background(), createdSvcID)
	_, errDel := reqDel.Execute()
	require.NoError(t, errDel, "Failed to delete service during cleanup")

	t.Logf("Successfully cleaned up service: %s", createdSvcID)
}

// Test_objects_ServicesAPIService_List tests listing services.
func Test_objects_ServicesAPIService_List(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a service to find in the list.
	createdSvcName := "test-svc-list-" + common.GenerateRandomString(6)
	service := objects.Services{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdSvcName,
		Protocol: &objects.ServicesProtocol{
			Tcp: &objects.ServicesProtocolTcp{Port: "8080"},
		},
	}

	// Create the service via API.
	req := client.ServicesAPI.CreateServices(context.Background()).Services(service)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create service for list test")
	require.NotNil(t, createRes.Id, "Created service ID should not be nil")
	createdSvcID := *createRes.Id
	require.NotEmpty(t, createdSvcID, "Created service ID should not be empty")

	// Test List operation.
	reqList := client.ServicesAPI.ListServices(context.Background()).Folder("Prisma Access").Limit(10000)
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful.
	require.NoError(t, errList, "Failed to list services")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties.
	require.NotNil(t, listRes, "List response should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one service in the list")

	// Verify our created service is in the list.
	foundSvc := false
	for _, svc := range listRes.Data {
		if svc.Id != nil && *svc.Id == createdSvcID {
			foundSvc = true
			assert.Equal(t, createdSvcName, svc.Name, "Service name should match in list")
			assert.True(t, *svc.Folder == "Shared" || *svc.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
			break
		}
	}
	assert.True(t, foundSvc, "Created service should be found in the list")

	t.Logf("Successfully listed and found service: %s", createdSvcName)

	// Cleanup: Delete the created service.
	reqDel := client.ServicesAPI.DeleteServicesByID(context.Background(), createdSvcID)
	_, errDel := reqDel.Execute()
	require.NoError(t, errDel, "Failed to delete service during cleanup")

	t.Logf("Successfully cleaned up service: %s", createdSvcID)
}

// Test_objects_ServicesAPIService_DeleteByID tests deleting a service.
func Test_objects_ServicesAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client.
	client := SetupObjectSvcTestClient(t)

	// Create a service to delete.
	createdSvcName := "test-svc-delete-" + common.GenerateRandomString(6)
	service := objects.Services{
		Folder: common.StringPtr("Prisma Access"),
		Name:   createdSvcName,
		Protocol: &objects.ServicesProtocol{
			Tcp: &objects.ServicesProtocolTcp{Port: "9999"},
		},
	}

	// Create the service via API.
	req := client.ServicesAPI.CreateServices(context.Background()).Services(service)
	createRes, _, err := req.Execute()
	require.NoError(t, err, "Failed to create service for delete test")
	require.NotNil(t, createRes.Id, "Created service ID should not be nil")
	createdSvcID := *createRes.Id
	require.NotEmpty(t, createdSvcID, "Created service ID should not be empty")

	// Test Delete by ID operation.
	reqDel := client.ServicesAPI.DeleteServicesByID(context.Background(), createdSvcID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful.
	require.NoError(t, errDel, "Failed to delete service")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted service: %s", createdSvcID)
}
