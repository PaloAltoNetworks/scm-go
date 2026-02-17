/*
Identity Services Testing LocalUsersAPIService
*/
package identity_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// Test_identity_services_LocalUsersAPIService_FetchLocalUsers tests the FetchLocalUsers convenience method
func Test_identity_services_LocalUsersAPIService_FetchLocalUsers(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Create a test object first
	testName := "fetch-user-" + common.GenerateRandomString(6)

	// Note: LocalUsers API rejects empty Id field, so we skip this test
	// This is an API limitation, not a template issue
	t.Skip("LocalUsers API does not support standard create pattern with empty Id field")

	testObj := identity_services.LocalUsers{
		Name:     testName,
		Folder:   common.StringPtr("Prisma Access"),
		Password: "TestPass123!",
	}

	createReq := client.LocalUsersAPI.CreateLocalUsers(context.Background()).LocalUsers(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.LocalUsersAPI.DeleteLocalUsersByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.LocalUsersAPI.FetchLocalUsers(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch local users by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchLocalUsers found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.LocalUsersAPI.FetchLocalUsers(
		context.Background(),
		"non-existent-user-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchLocalUsers correctly returned nil for non-existent object")
}
