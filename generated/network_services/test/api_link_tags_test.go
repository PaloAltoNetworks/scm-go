/*
 * Network Services Testing
 *
 * LinkTagsAPIService
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

// Test_network_services_LinkTagsAPIService_Create tests the creation of a Link Tag.
func Test_network_services_LinkTagsAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	tagName := "test-linktag-create-" + randomSuffix

	linkTag := network_services.LinkTags{
		Name:     tagName,
		Folder:   common.StringPtr("All"),
		Comments: common.StringPtr("Test link tag for create"),
	}

	fmt.Printf("Attempting to create Link Tag with name: %s\n", linkTag.Name)

	req := client.LinkTagsAPI.CreateLinkTags(context.Background()).LinkTags(linkTag)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Link Tag")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created link tag should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up Link Tag with ID: %s", createdID)
		_, errDel := client.LinkTagsAPI.DeleteLinkTagsByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete link tag during cleanup")
	}()

	assert.Equal(t, tagName, res.Name, "Created link tag name should match")
	t.Logf("Successfully created and validated Link Tag: %s with ID: %s", linkTag.Name, createdID)
}

// Test_network_services_LinkTagsAPIService_GetByID tests retrieving a link tag by its ID.
func Test_network_services_LinkTagsAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	tagName := "test-linktag-get-" + randomSuffix

	linkTag := network_services.LinkTags{
		Name:   tagName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.LinkTagsAPI.CreateLinkTags(context.Background()).LinkTags(linkTag).Execute()
	require.NoError(t, err, "Failed to create link tag for get test")
	createdID := *createRes.Id

	defer func() {
		client.LinkTagsAPI.DeleteLinkTagsByID(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.LinkTagsAPI.GetLinkTagsByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get link tag by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, tagName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_network_services_LinkTagsAPIService_Update tests updating an existing link tag.
func Test_network_services_LinkTagsAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	tagName := "test-linktag-update-" + randomSuffix

	linkTag := network_services.LinkTags{
		Name:   tagName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.LinkTagsAPI.CreateLinkTags(context.Background()).LinkTags(linkTag).Execute()
	require.NoError(t, err, "Failed to create link tag for update test")
	createdID := *createRes.Id

	defer func() {
		client.LinkTagsAPI.DeleteLinkTagsByID(context.Background(), createdID).Execute()
	}()

	updatedLinkTag := network_services.LinkTags{
		Name:     tagName,
		Folder:   common.StringPtr("All"),
		Comments: common.StringPtr("Updated link tag comment"),
	}

	updateRes, httpResUpdate, errUpdate := client.LinkTagsAPI.UpdateLinkTagsByID(context.Background(), createdID).LinkTags(updatedLinkTag).Execute()
	require.NoError(t, errUpdate, "Failed to update link tag")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, "Updated link tag comment", *updateRes.Comments, "Comments should be updated")
}

// Test_network_services_LinkTagsAPIService_List tests listing Link Tags.
func Test_network_services_LinkTagsAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	tagName := "test-linktag-list-" + randomSuffix

	linkTag := network_services.LinkTags{
		Name:   tagName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.LinkTagsAPI.CreateLinkTags(context.Background()).LinkTags(linkTag).Execute()
	require.NoError(t, err, "Failed to create link tag for list test")
	createdID := *createRes.Id

	defer func() {
		client.LinkTagsAPI.DeleteLinkTagsByID(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.LinkTagsAPI.ListLinkTags(context.Background()).Folder("All").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list link tags")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundTag := false
	for _, tag := range listRes.Data {
		if tag.Name == tagName {
			foundTag = true
			break
		}
	}
	assert.True(t, foundTag, "Created link tag should be found in the list")
}

// Test_network_services_LinkTagsAPIService_DeleteByID tests deleting a link tag by ID.
func Test_network_services_LinkTagsAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	tagName := "test-linktag-delete-" + randomSuffix

	linkTag := network_services.LinkTags{
		Name:   tagName,
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.LinkTagsAPI.CreateLinkTags(context.Background()).LinkTags(linkTag).Execute()
	require.NoError(t, err, "Failed to create link tag for delete test")
	createdID := *createRes.Id

	_, errDel := client.LinkTagsAPI.DeleteLinkTagsByID(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete link tag")
}

// Test_network_services_LinkTagsAPIService_FetchLinkTags tests the FetchLinkTags convenience method
func Test_network_services_LinkTagsAPIService_FetchLinkTags(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-linktag-fetch-" + randomSuffix

	testObj := network_services.LinkTags{
		Name:   testName,
		Folder: common.StringPtr("All"),
	}

	createReq := client.LinkTagsAPI.CreateLinkTags(context.Background()).LinkTags(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		deleteReq := client.LinkTagsAPI.DeleteLinkTagsByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.LinkTagsAPI.FetchLinkTags(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch link_tags by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchLinkTags found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.LinkTagsAPI.FetchLinkTags(
		context.Background(),
		"non-existent-link_tags-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchLinkTags correctly returned nil for non-existent object")
}
