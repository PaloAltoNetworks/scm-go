/*
Objects Testing ServiceGroupsAPIService
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

// createTestService is a helper function to create a service for testing.
func createTestService(t *testing.T, client *objects.APIClient, name string, protocol objects.ServicesProtocol) string {
	t.Helper()

	service := objects.Services{
		Name:     name,
		Folder:   common.StringPtr("Prisma Access"),
		Protocol: &protocol,
	}

	req := client.ServicesAPI.CreateServices(context.Background()).Services(service)
	res, _, err := req.Execute()
	require.NoError(t, err, "Failed to create test service")
	require.NotNil(t, res.Id, "Created test service ID should not be nil")
	require.NotEmpty(t, *res.Id, "Created test service ID should not be empty")

	t.Logf("Successfully created test service: %s (ID: %s)", name, *res.Id)
	return *res.Id
}

// deleteTestService is a helper function to delete a service.
func deleteTestService(t *testing.T, client *objects.APIClient, id string) {
	t.Helper()

	req := client.ServicesAPI.DeleteServicesByID(context.Background(), id)
	httpRes, err := req.Execute()
	require.NoError(t, err, "Failed to delete test service")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status for service deletion")
	t.Logf("Successfully deleted test service ID: %s", id)
}

// createTestServiceGroup is a helper function to create a service group for testing.
func createTestServiceGroup(t *testing.T, client *objects.APIClient, group objects.ServiceGroups) string {
	t.Helper()

	req := client.ServiceGroupsAPI.CreateServiceGroups(context.Background()).ServiceGroups(group)
	res, _, err := req.Execute()
	require.NoError(t, err, "Failed to create service group")
	require.NotNil(t, res.Id, "Created service group ID should not be nil")
	require.NotEmpty(t, res.Id, "Created service group ID should not be empty")

	t.Logf("Successfully created test service group: %s (ID: %s)", group.Name, res.Id)
	return res.Id
}

// Test_objects_ServiceGroupsAPIService_Create tests the creation of a service group.
func Test_objects_ServiceGroupsAPIService_Create(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	service1Name := "test-svc-1-" + randomSuffix
	service2Name := "test-svc-2-" + randomSuffix
	groupName := "test-sg-create-" + randomSuffix

	// Create member services.
	protocol1 := objects.ServicesProtocol{Tcp: &objects.ServicesProtocolTcp{Port: "80"}}
	service1ID := createTestService(t, client, service1Name, protocol1)
	t.Cleanup(func() { deleteTestService(t, client, service1ID) })

	protocol2 := objects.ServicesProtocol{Udp: &objects.ServicesProtocolUdp{Port: "53"}}
	service2ID := createTestService(t, client, service2Name, protocol2)
	t.Cleanup(func() { deleteTestService(t, client, service2ID) })

	// Define the service group to be created.
	serviceGroup := objects.ServiceGroups{
		Name:    groupName,
		Folder:  common.StringPtr("Prisma Access"),
		Members: []string{service1Name, service2Name},
	}

	// Create the service group.
	createReq := client.ServiceGroupsAPI.CreateServiceGroups(context.Background()).ServiceGroups(serviceGroup)
	createRes, httpRes, err := createReq.Execute()
	require.NoError(t, err)
	assert.Equal(t, 201, httpRes.StatusCode)
	require.NotNil(t, createRes, "Create response should not be nil")
	createdGroupID := createRes.Id
	t.Cleanup(func() {
		delReq := client.ServiceGroupsAPI.DeleteServiceGroupsByID(context.Background(), createdGroupID)
		_, err := delReq.Execute()
		require.NoError(t, err)
	})

	// Verify the created group's properties.
	assert.Equal(t, groupName, createRes.Name)
	assert.True(t, *createRes.Folder == "Shared" || *createRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.ElementsMatch(t, []string{service1Name, service2Name}, createRes.Members)
	assert.NotEmpty(t, createRes.Id)
}

// Test_objects_ServiceGroupsAPIService_GetByID tests getting a service group by ID.
func Test_objects_ServiceGroupsAPIService_GetByID(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	service1Name := "test-svc-get-1-" + randomSuffix
	groupName := "test-sg-get-" + randomSuffix

	// Create member services.
	protocol1 := objects.ServicesProtocol{Tcp: &objects.ServicesProtocolTcp{Port: "8080"}}
	service1ID := createTestService(t, client, service1Name, protocol1)
	t.Cleanup(func() { deleteTestService(t, client, service1ID) })

	// Create a service group to retrieve.
	serviceGroup := objects.ServiceGroups{
		Name:    groupName,
		Folder:  common.StringPtr("Prisma Access"),
		Members: []string{service1Name},
	}
	createdGroupID := createTestServiceGroup(t, client, serviceGroup)
	t.Cleanup(func() {
		delReq := client.ServiceGroupsAPI.DeleteServiceGroupsByID(context.Background(), createdGroupID)
		_, err := delReq.Execute()
		require.NoError(t, err)
	})

	// Get the service group by its ID.
	getReq := client.ServiceGroupsAPI.GetServiceGroupsByID(context.Background(), createdGroupID)
	getRes, httpRes, err := getReq.Execute()
	require.NoError(t, err)
	assert.Equal(t, 200, httpRes.StatusCode)

	// Verify the retrieved group's properties.
	assert.Equal(t, groupName, getRes.Name)
	assert.Equal(t, createdGroupID, getRes.Id)
	assert.True(t, *getRes.Folder == "Shared" || *getRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, []string{service1Name}, getRes.Members)
}

// Test_objects_ServiceGroupsAPIService_Create tests the creation of a service group.
func Test_objects_ServiceGroupsAPIService_Update(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	service1Name := "test-svc-1-" + randomSuffix
	service2Name := "test-svc-2-" + randomSuffix
	groupName := "test-sg-create-" + randomSuffix

	// Create member services.
	protocol1 := objects.ServicesProtocol{Tcp: &objects.ServicesProtocolTcp{Port: "80"}}
	service1ID := createTestService(t, client, service1Name, protocol1)
	t.Cleanup(func() { deleteTestService(t, client, service1ID) })

	protocol2 := objects.ServicesProtocol{Udp: &objects.ServicesProtocolUdp{Port: "53"}}
	service2ID := createTestService(t, client, service2Name, protocol2)
	t.Cleanup(func() { deleteTestService(t, client, service2ID) })

	// Define the service group to be created.
	serviceGroup := objects.ServiceGroups{
		Name:    groupName,
		Folder:  common.StringPtr("Prisma Access"),
		Members: []string{service1Name, service2Name},
	}

	// Create the service group.
	createReq := client.ServiceGroupsAPI.CreateServiceGroups(context.Background()).ServiceGroups(serviceGroup)
	createRes, httpRes, err := createReq.Execute()
	require.NoError(t, err)
	assert.Equal(t, 201, httpRes.StatusCode)
	require.NotNil(t, createRes, "Create response should not be nil")
	createdGroupID := createRes.Id

	// Verify the created group's properties.
	assert.Equal(t, groupName, createRes.Name)
	assert.True(t, *createRes.Folder == "Shared" || *createRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.ElementsMatch(t, []string{service1Name, service2Name}, createRes.Members)
	assert.NotEmpty(t, createRes.Id)

	// Update the service group
	// Define the service group to be created.
	updatedServiceGroup := objects.ServiceGroups{
		Name:    groupName,
		Folder:  common.StringPtr("Prisma Access"),
		Members: []string{service1Name},
	}
	// Create the service group.
	updateReq := client.ServiceGroupsAPI.UpdateServiceGroupsByID(context.Background(), createdGroupID).ServiceGroups(updatedServiceGroup)
	updateRes, httpRes, errUpdate := updateReq.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update address group")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, groupName, updateRes.Name, "Service name should remain the same")
	assert.True(t, *updateRes.Folder == "Shared" || *updateRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")

	t.Logf("Successfully updated service: %s", updatedServiceGroup.Name)

	t.Cleanup(func() {
		delReq := client.ServiceGroupsAPI.DeleteServiceGroupsByID(context.Background(), createdGroupID)
		_, err := delReq.Execute()
		require.NoError(t, err)
	})

}

// Test_objects_ServiceGroupsAPIService_List tests listing service groups.
func Test_objects_ServiceGroupsAPIService_List(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	service1Name := "test-svc-1-" + randomSuffix
	groupName := "test-sg-create-" + randomSuffix

	// Create member services.
	protocol1 := objects.ServicesProtocol{Tcp: &objects.ServicesProtocolTcp{Port: "80"}}
	service1ID := createTestService(t, client, service1Name, protocol1)
	t.Cleanup(func() { deleteTestService(t, client, service1ID) })

	// Define the service group to be created.
	serviceGroup := objects.ServiceGroups{
		Name:    groupName,
		Folder:  common.StringPtr("Prisma Access"),
		Members: []string{service1Name},
	}

	// Create the service group.
	createReq := client.ServiceGroupsAPI.CreateServiceGroups(context.Background()).ServiceGroups(serviceGroup)
	createRes, httpRes, err := createReq.Execute()
	require.NoError(t, err)
	assert.Equal(t, 201, httpRes.StatusCode)
	require.NotNil(t, createRes, "Create response should not be nil")
	createdGroupID := createRes.Id

	// Perform the list request.
	listReq := client.ServiceGroupsAPI.ListServiceGroups(context.Background()).Folder("Prisma Access")
	listRes, httpRes, err := listReq.Execute()
	require.NoError(t, err)
	assert.Equal(t, 200, httpRes.StatusCode)
	require.NotNil(t, listRes.Data, "List data should not be nil")

	// Find the created group in the list.
	found := false
	for _, group := range listRes.GetData() {
		if group.Id == createdGroupID {
			found = true
			break
		}
	}
	assert.True(t, found, "Did not find created service group in the list")

	// Delete the service group.
	delReq := client.ServiceGroupsAPI.DeleteServiceGroupsByID(context.Background(), createdGroupID)

	httpRes, err = delReq.Execute()
	require.NoError(t, err)
	assert.Equal(t, 200, httpRes.StatusCode)
}

// Test_objects_ServiceGroupsAPIService_DeleteByID tests deleting a service group.
func Test_objects_ServiceGroupsAPIService_DeleteByID(t *testing.T) {
	client := SetupObjectSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	service1Name := "test-svc-1-" + randomSuffix
	service2Name := "test-svc-2-" + randomSuffix
	groupName := "test-sg-create-" + randomSuffix

	// Create member services.
	protocol1 := objects.ServicesProtocol{Tcp: &objects.ServicesProtocolTcp{Port: "80"}}
	service1ID := createTestService(t, client, service1Name, protocol1)
	t.Cleanup(func() { deleteTestService(t, client, service1ID) })

	protocol2 := objects.ServicesProtocol{Udp: &objects.ServicesProtocolUdp{Port: "53"}}
	service2ID := createTestService(t, client, service2Name, protocol2)
	t.Cleanup(func() { deleteTestService(t, client, service2ID) })

	// Define the service group to be created.
	serviceGroup := objects.ServiceGroups{
		Name:    groupName,
		Folder:  common.StringPtr("Prisma Access"),
		Members: []string{service1Name, service2Name},
	}

	// Create the service group.
	createReq := client.ServiceGroupsAPI.CreateServiceGroups(context.Background()).ServiceGroups(serviceGroup)
	createRes, httpRes, err := createReq.Execute()
	require.NoError(t, err)
	assert.Equal(t, 201, httpRes.StatusCode)
	require.NotNil(t, createRes, "Create response should not be nil")
	createdGroupID := createRes.Id

	// Delete the service group.
	delReq := client.ServiceGroupsAPI.DeleteServiceGroupsByID(context.Background(), createdGroupID)

	httpRes, err = delReq.Execute()
	require.NoError(t, err)
	assert.Equal(t, 200, httpRes.StatusCode)
}

// Test_objects_ServiceGroupsAPIService_FetchServiceGroups tests the FetchServiceGroups convenience method
func Test_objects_ServiceGroupsAPIService_FetchServiceGroups(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create test services first (same as Create test)
	randomSuffix := common.GenerateRandomString(6)
	service1Name := "test-svc-1-fetch-" + randomSuffix
	service2Name := "test-svc-2-fetch-" + randomSuffix
	testName := "test-sg-fetch-" + randomSuffix

	protocol1 := objects.ServicesProtocol{Tcp: &objects.ServicesProtocolTcp{Port: "80"}}
	service1ID := createTestService(t, client, service1Name, protocol1)

	protocol2 := objects.ServicesProtocol{Udp: &objects.ServicesProtocolUdp{Port: "53"}}
	service2ID := createTestService(t, client, service2Name, protocol2)

	// Create test object using same payload as Create test
	testObj := objects.ServiceGroups{
		Name:    testName,
		Folder:  common.StringPtr("Prisma Access"),
		Members: []string{service1Name, service2Name},
	}

	createReq := client.ServiceGroupsAPI.CreateServiceGroups(context.Background()).ServiceGroups(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.ServiceGroupsAPI.DeleteServiceGroupsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
		// Cleanup test services
		deleteTestService(t, client, service1ID)
		deleteTestService(t, client, service2ID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.ServiceGroupsAPI.FetchServiceGroups(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch service_groups by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchServiceGroups found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.ServiceGroupsAPI.FetchServiceGroups(
		context.Background(),
		"non-existent-service_groups-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchServiceGroups correctly returned nil for non-existent object")
}
