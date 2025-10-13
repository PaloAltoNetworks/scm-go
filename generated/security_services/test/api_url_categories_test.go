/*
 * Security Services Testing
 *
 * URLCategoriesAPIService
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

// Test_security_services_URLCategoriesAPIService_Create tests the creation of a URL Category.
func Test_security_services_URLCategoriesAPIService_Create(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdUrlCategoryName := "test-url-cat-create-" + common.GenerateRandomString(6)
	urlCategory := security_services.UrlCategories{
		Folder:      common.StringPtr("Shared"),
		Name:        createdUrlCategoryName,
		Description: common.StringPtr("Test URL Category for create API"),
		List:        []string{"example.com", "test-create.com"},
		Type:        common.StringPtr("URL List"),
	}

	fmt.Printf("Creating URL Category with name: %s\n", urlCategory.Name)
	req := client.URLCategoriesAPI.CreateURLCategories(context.Background()).UrlCategories(urlCategory)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create URL Category")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdUrlCategoryName, res.Name, "Created URL Category name should match")
	createdUrlCategoryID := *res.Id

	defer func() {
		t.Logf("Cleaning up URL Category with ID: %s", createdUrlCategoryID)
		_, errDel := client.URLCategoriesAPI.DeleteURLCategoriesByID(context.Background(), createdUrlCategoryID).Execute()
		require.NoError(t, errDel, "Failed to delete URL Category during cleanup")
	}()

	t.Logf("Successfully created URL Category: %s with ID: %s", urlCategory.Name, createdUrlCategoryID)
	assert.Equal(t, "URL List", *res.Type, "Type should be 'URL List'")
	assert.Contains(t, res.List, "example.com", "List should contain 'example.com'")
}

// Test_security_services_URLCategoriesAPIService_GetByID tests retrieving a URL Category by its ID.
func Test_security_services_URLCategoriesAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdUrlCategoryName := "test-url-cat-get-" + common.GenerateRandomString(6)
	urlCategory := security_services.UrlCategories{
		Folder: common.StringPtr("Shared"),
		Name:   createdUrlCategoryName,
		List:   []string{"get-test.com"},
		Type:   common.StringPtr("URL List"),
	}

	createRes, _, err := client.URLCategoriesAPI.CreateURLCategories(context.Background()).UrlCategories(urlCategory).Execute()
	require.NoError(t, err, "Failed to create URL Category for get test")
	createdUrlCategoryID := *createRes.Id
	require.NotEmpty(t, createdUrlCategoryID, "Created URL Category ID should not be empty")

	defer func() {
		t.Logf("Cleaning up URL Category with ID: %s", createdUrlCategoryID)
		_, errDel := client.URLCategoriesAPI.DeleteURLCategoriesByID(context.Background(), createdUrlCategoryID).Execute()
		require.NoError(t, errDel, "Failed to delete URL Category during cleanup")
	}()

	getRes, httpResGet, errGet := client.URLCategoriesAPI.GetURLCategoriesByID(context.Background(), createdUrlCategoryID).Execute()
	require.NoError(t, errGet, "Failed to get URL Category by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdUrlCategoryName, getRes.Name, "URL Category name should match")
	assert.Equal(t, []string{"get-test.com"}, getRes.List, "List should match")
}

// Test_security_services_URLCategoriesAPIService_Update tests updating an existing URL Category.
func Test_security_services_URLCategoriesAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdUrlCategoryName := "test-url-cat-update-" + common.GenerateRandomString(6)
	urlCategory := security_services.UrlCategories{
		Folder: common.StringPtr("Shared"),
		Name:   createdUrlCategoryName,
		List:   []string{"initial.com"},
		Type:   common.StringPtr("URL List"),
	}

	createRes, _, err := client.URLCategoriesAPI.CreateURLCategories(context.Background()).UrlCategories(urlCategory).Execute()
	require.NoError(t, err, "Failed to create URL Category for update test")
	createdUrlCategoryID := *createRes.Id
	require.NotEmpty(t, createdUrlCategoryID, "Created URL Category ID should not be empty")

	defer func() {
		t.Logf("Cleaning up URL Category with ID: %s", createdUrlCategoryID)
		_, errDel := client.URLCategoriesAPI.DeleteURLCategoriesByID(context.Background(), createdUrlCategoryID).Execute()
		require.NoError(t, errDel, "Failed to delete URL Category during cleanup")
	}()

	updatedUrlCategory := security_services.UrlCategories{
		Name:        createdUrlCategoryName,
		Description: common.StringPtr("Updated description"),
		List:        []string{"updated.com", "another.com"},
		Type:        common.StringPtr("URL List"),
	}

	updateRes, httpResUpdate, errUpdate := client.URLCategoriesAPI.UpdateURLCategoriesByID(context.Background(), createdUrlCategoryID).UrlCategories(updatedUrlCategory).Execute()
	require.NoError(t, errUpdate, "Failed to update URL Category")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, "Updated description", *updateRes.Description, "Description should be updated")
	assert.Equal(t, []string{"updated.com", "another.com"}, updateRes.List, "List should be updated")
}

// Test_security_services_URLCategoriesAPIService_List tests listing URL Categories.
func Test_security_services_URLCategoriesAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdUrlCategoryName := "test-url-cat-list-" + common.GenerateRandomString(6)
	urlCategory := security_services.UrlCategories{
		Folder: common.StringPtr("Shared"),
		Name:   createdUrlCategoryName,
		Type:   common.StringPtr("URL List"),
	}

	createRes, _, err := client.URLCategoriesAPI.CreateURLCategories(context.Background()).UrlCategories(urlCategory).Execute()
	require.NoError(t, err, "Failed to create URL Category for list test")
	createdUrlCategoryID := *createRes.Id
	require.NotEmpty(t, createdUrlCategoryID, "Created URL Category ID should not be empty")

	defer func() {
		t.Logf("Cleaning up URL Category with ID: %s", createdUrlCategoryID)
		_, errDel := client.URLCategoriesAPI.DeleteURLCategoriesByID(context.Background(), createdUrlCategoryID).Execute()
		require.NoError(t, errDel, "Failed to delete URL Category during cleanup")
	}()

	listRes, httpResList, errList := client.URLCategoriesAPI.ListURLCategories(context.Background()).Folder("Shared").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list URL Categories")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	foundObject := false
	for _, cat := range listRes.Data {
		if cat.Name != "" && cat.Name == createdUrlCategoryName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created URL Category should be found in the list")
}

// Test_security_services_URLCategoriesAPIService_DeleteByID tests deleting a URL Category by its ID.
func Test_security_services_URLCategoriesAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdUrlCategoryName := "test-url-cat-delete-" + common.GenerateRandomString(6)
	urlCategory := security_services.UrlCategories{
		Folder: common.StringPtr("Shared"),
		Name:   createdUrlCategoryName,
		Type:   common.StringPtr("URL List"),
	}

	createRes, _, err := client.URLCategoriesAPI.CreateURLCategories(context.Background()).UrlCategories(urlCategory).Execute()
	require.NoError(t, err, "Failed to create URL Category for delete test")
	createdUrlCategoryID := *createRes.Id
	require.NotEmpty(t, createdUrlCategoryID, "Created URL Category ID should not be empty")

	httpResDel, errDel := client.URLCategoriesAPI.DeleteURLCategoriesByID(context.Background(), createdUrlCategoryID).Execute()
	require.NoError(t, errDel, "Failed to delete URL Category")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
}
