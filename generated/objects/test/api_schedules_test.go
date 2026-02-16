/*
Objects Testing SchedulesAPIService
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

// Test_objects_SchedulesAPIService_Create tests creating a schedule
func Test_objects_SchedulesAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object
	testName := "test-schedule-" + common.GenerateRandomString(6)

	// Create schedule_type with recurring daily option
	scheduleType := objects.SchedulesScheduleType{
		Recurring: &objects.SchedulesScheduleTypeRecurring{
			Daily: []string{"00:00-23:59"},
		},
	}

	testObj := objects.Schedules{
		Name:         testName,
		Folder:       common.StringPtr("Prisma Access"),
		ScheduleType: scheduleType,
	}

	createReq := client.SchedulesAPI.CreateSchedules(context.Background()).Schedules(testObj)
	resp, httpResp, err := createReq.Execute()

	// Cleanup after test
	if resp != nil && resp.Id != "" {
		defer func() {
			deleteReq := client.SchedulesAPI.DeleteSchedulesByID(context.Background(), resp.Id)
			_, _ = deleteReq.Execute()
			t.Logf("Cleaned up test object: %s", resp.Id)
		}()
	}

	// Verify the response
	require.NoError(t, err, "Failed to create schedule")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 201, httpResp.StatusCode, "Expected 201 Created status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, testName, resp.Name, "Created schedule name should match")
	assert.NotEmpty(t, resp.Id, "Created schedule should have an ID")
	t.Logf("[SUCCESS] Created schedule: %s (ID: %s)", resp.Name, resp.Id)
}

// Test_objects_SchedulesAPIService_GetByID tests getting a schedule by ID
func Test_objects_SchedulesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first
	testName := "test-schedule-" + common.GenerateRandomString(6)

	scheduleType := objects.SchedulesScheduleType{
		Recurring: &objects.SchedulesScheduleTypeRecurring{
			Daily: []string{"00:00-23:59"},
		},
	}

	testObj := objects.Schedules{
		Name:         testName,
		Folder:       common.StringPtr("Prisma Access"),
		ScheduleType: scheduleType,
	}

	createReq := client.SchedulesAPI.CreateSchedules(context.Background()).Schedules(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.SchedulesAPI.DeleteSchedulesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Get the object by ID
	getReq := client.SchedulesAPI.GetSchedulesByID(context.Background(), createdID)
	resp, httpResp, err := getReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to get schedule by ID")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, createdID, resp.Id, "Retrieved schedule ID should match")
	assert.Equal(t, testName, resp.Name, "Retrieved schedule name should match")
	t.Logf("[SUCCESS] Retrieved schedule by ID: %s", resp.Id)
}

// Test_objects_SchedulesAPIService_Update tests updating a schedule
func Test_objects_SchedulesAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first
	testName := "test-schedule-" + common.GenerateRandomString(6)

	scheduleType := objects.SchedulesScheduleType{
		Recurring: &objects.SchedulesScheduleTypeRecurring{
			Daily: []string{"00:00-23:59"},
		},
	}

	testObj := objects.Schedules{
		Name:         testName,
		Folder:       common.StringPtr("Prisma Access"),
		ScheduleType: scheduleType,
	}

	createReq := client.SchedulesAPI.CreateSchedules(context.Background()).Schedules(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.SchedulesAPI.DeleteSchedulesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Update the object with different schedule
	updatedScheduleType := objects.SchedulesScheduleType{
		Recurring: &objects.SchedulesScheduleTypeRecurring{
			Daily: []string{"08:00-17:00"},
		},
	}

	updateObj := objects.Schedules{
		Id:           createdID,
		Name:         testName,
		Folder:       common.StringPtr("Prisma Access"),
		ScheduleType: updatedScheduleType,
	}

	updateReq := client.SchedulesAPI.UpdateSchedulesByID(context.Background(), createdID).Schedules(updateObj)
	resp, httpResp, err := updateReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to update schedule")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, createdID, resp.Id, "Updated schedule ID should match")

	// Verify the schedule was updated
	if resp.ScheduleType.Recurring != nil && len(resp.ScheduleType.Recurring.Daily) > 0 {
		assert.Equal(t, "08:00-17:00", resp.ScheduleType.Recurring.Daily[0], "Schedule should be updated")
	}
	t.Logf("[SUCCESS] Updated schedule: %s", resp.Id)
}

// Test_objects_SchedulesAPIService_List tests listing schedules
func Test_objects_SchedulesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first
	testName := "test-schedule-" + common.GenerateRandomString(6)

	scheduleType := objects.SchedulesScheduleType{
		Recurring: &objects.SchedulesScheduleTypeRecurring{
			Daily: []string{"00:00-23:59"},
		},
	}

	testObj := objects.Schedules{
		Name:         testName,
		Folder:       common.StringPtr("Prisma Access"),
		ScheduleType: scheduleType,
	}

	createReq := client.SchedulesAPI.CreateSchedules(context.Background()).Schedules(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.SchedulesAPI.DeleteSchedulesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// List schedules with folder filter
	listReq := client.SchedulesAPI.ListSchedules(context.Background()).Folder("Prisma Access")
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list schedules")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")

	// Verify our created object is in the list
	found := false
	if resp.Data != nil {
		for _, item := range resp.Data {
			if item.Id == createdID {
				found = true
				assert.Equal(t, testName, item.Name, "Listed schedule name should match")
				break
			}
		}
	}
	assert.True(t, found, "Created schedule should be in the list")
	t.Logf("[SUCCESS] Listed schedules, found test object: %s", createdID)
}

// Test_objects_SchedulesAPIService_DeleteByID tests deleting a schedule
func Test_objects_SchedulesAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first
	testName := "test-schedule-" + common.GenerateRandomString(6)

	scheduleType := objects.SchedulesScheduleType{
		Recurring: &objects.SchedulesScheduleTypeRecurring{
			Daily: []string{"00:00-23:59"},
		},
	}

	testObj := objects.Schedules{
		Name:         testName,
		Folder:       common.StringPtr("Prisma Access"),
		ScheduleType: scheduleType,
	}

	createReq := client.SchedulesAPI.CreateSchedules(context.Background()).Schedules(testObj)
	createRes, _, err := createReq.Execute()
	require.NoError(t, err, "Failed to create test object")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Delete the object
	deleteReq := client.SchedulesAPI.DeleteSchedulesByID(context.Background(), createdID)
	httpResp, err := deleteReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to delete schedule")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	t.Logf("[SUCCESS] Deleted schedule: %s", createdID)
}

// Test_objects_SchedulesAPIService_FetchSchedules tests the FetchSchedules convenience method
func Test_objects_SchedulesAPIService_FetchSchedules(t *testing.T) {
	// Setup the authenticated client
	client := SetupObjectSvcTestClient(t)

	// Create a test object first
	testName := "fetch-schedule-" + common.GenerateRandomString(6)

	// Create schedule_type with recurring daily option
	scheduleType := objects.SchedulesScheduleType{
		Recurring: &objects.SchedulesScheduleTypeRecurring{
			Daily: []string{"00:00-23:59"},
		},
	}

	testObj := objects.Schedules{
		Name:         testName,
		Folder:       common.StringPtr("Prisma Access"),
		ScheduleType: scheduleType, // Required field
	}

	createReq := client.SchedulesAPI.CreateSchedules(context.Background()).Schedules(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.SchedulesAPI.DeleteSchedulesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.SchedulesAPI.FetchSchedules(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch schedules by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchSchedules found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.SchedulesAPI.FetchSchedules(
		context.Background(),
		"non-existent-schedules-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchSchedules correctly returned nil for non-existent object")
}
