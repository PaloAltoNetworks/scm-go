/*
 * Security Services Testing
 *
 * DataObjectsAPIService
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

// Test_security_services_DataObjectsAPIService_Create tests the creation of a Data Object with predefined pattern.
func Test_security_services_DataObjectsAPIService_Create(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdObjectName := "test-do-create-" + common.GenerateRandomString(6)

	// Define a predefined pattern (similar to Terraform example)
	predefinedPattern := security_services.DataObjectsPatternTypePredefinedPatternInner{
		Name:     common.StringPtr("ABA-Routing-Number"),
		FileType: []string{"text/html"},
	}
	predefined := security_services.DataObjectsPatternTypePredefined{
		Pattern: []security_services.DataObjectsPatternTypePredefinedPatternInner{predefinedPattern},
	}
	patternType := security_services.DataObjectsPatternType{
		Predefined: &predefined,
	}

	// Define a Data Object
	dataObject := security_services.DataObjects{
		Folder:      common.StringPtr("ngfw-shared"),
		Name:        common.StringPtr(createdObjectName),
		Description: common.StringPtr("Test Data Object for create API"),
		PatternType: &patternType,
	}

	fmt.Printf("Creating Data Object with name: %s\n", createdObjectName)
	req := client.DataObjectsAPI.CreateDataObjects(context.Background()).DataObjects(dataObject)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Data Object")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdObjectName, *res.Name, "Created object name should match")
	require.NotNil(t, res.PatternType, "PatternType should not be nil")
	require.NotNil(t, res.PatternType.Predefined, "Predefined pattern should not be nil")

	createdObjectID := res.Id

	defer func() {
		t.Logf("Cleaning up Data Object with ID: %s", *createdObjectID)
		_, errDel := client.DataObjectsAPI.DeleteDataObjectsByID(context.Background(), *createdObjectID).Execute()
		require.NoError(t, errDel, "Failed to delete Data Object during cleanup")
	}()

	t.Logf("Successfully created Data Object: %s with ID: %s", createdObjectName, *createdObjectID)
}

// Test_security_services_DataObjectsAPIService_CreateWithFileProperties tests creating a Data Object with file properties pattern.
func Test_security_services_DataObjectsAPIService_CreateWithFileProperties(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	createdObjectName := "test-do-fp-" + common.GenerateRandomString(6)

	// Define a file properties pattern (similar to Terraform example)
	filePropertiesPattern := security_services.DataObjectsPatternTypeFilePropertiesPatternInner{
		Name:          common.StringPtr("test_pdf"),
		FileType:      common.StringPtr("pdf"),
		FileProperty:  common.StringPtr("panav-rsp-pdf-dlp-author"),
		PropertyValue: common.StringPtr("test_value"),
	}
	fileProperties := security_services.DataObjectsPatternTypeFileProperties{
		Pattern: []security_services.DataObjectsPatternTypeFilePropertiesPatternInner{filePropertiesPattern},
	}
	patternType := security_services.DataObjectsPatternType{
		FileProperties: &fileProperties,
	}

	// Define a Data Object
	dataObject := security_services.DataObjects{
		Folder:      common.StringPtr("ngfw-shared"),
		Name:        common.StringPtr(createdObjectName),
		Description: common.StringPtr("Test Data Object with file properties"),
		PatternType: &patternType,
	}

	fmt.Printf("Creating Data Object with file properties: %s\n", createdObjectName)
	req := client.DataObjectsAPI.CreateDataObjects(context.Background()).DataObjects(dataObject)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Data Object with file properties")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdObjectName, *res.Name, "Created object name should match")
	require.NotNil(t, res.PatternType, "PatternType should not be nil")
	require.NotNil(t, res.PatternType.FileProperties, "FileProperties pattern should not be nil")

	createdObjectID := res.Id

	defer func() {
		t.Logf("Cleaning up Data Object with ID: %s", *createdObjectID)
		_, errDel := client.DataObjectsAPI.DeleteDataObjectsByID(context.Background(), *createdObjectID).Execute()
		require.NoError(t, errDel, "Failed to delete Data Object during cleanup")
	}()

	t.Logf("Successfully created Data Object with file properties: %s with ID: %s", createdObjectName, *createdObjectID)
}

// Test_security_services_DataObjectsAPIService_GetByID tests retrieving a Data Object by its ID.
func Test_security_services_DataObjectsAPIService_GetByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	objectName := "test-do-get-" + common.GenerateRandomString(6)

	predefinedPattern := security_services.DataObjectsPatternTypePredefinedPatternInner{
		Name:     common.StringPtr("ABA-Routing-Number"),
		FileType: []string{"text/html"},
	}
	predefined := security_services.DataObjectsPatternTypePredefined{
		Pattern: []security_services.DataObjectsPatternTypePredefinedPatternInner{predefinedPattern},
	}
	patternType := security_services.DataObjectsPatternType{
		Predefined: &predefined,
	}

	dataObject := security_services.DataObjects{
		Folder:      common.StringPtr("ngfw-shared"),
		Name:        common.StringPtr(objectName),
		PatternType: &patternType,
	}

	createRes, _, err := client.DataObjectsAPI.CreateDataObjects(context.Background()).DataObjects(dataObject).Execute()
	require.NoError(t, err, "Failed to create Data Object for get test")
	createdObjectID := createRes.Id
	require.NotEmpty(t, createdObjectID, "Created object ID should not be empty")

	defer func() {
		t.Logf("Cleaning up Data Object with ID: %s", *createdObjectID)
		_, errDel := client.DataObjectsAPI.DeleteDataObjectsByID(context.Background(), *createdObjectID).Execute()
		require.NoError(t, errDel, "Failed to delete Data Object during cleanup")
	}()

	getRes, httpResGet, errGet := client.DataObjectsAPI.GetDataObjectsByID(context.Background(), *createdObjectID).Execute()
	require.NoError(t, errGet, "Failed to get Data Object by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, objectName, *getRes.Name, "Object name should match")
}

// Test_security_services_DataObjectsAPIService_Update tests updating an existing Data Object.
func Test_security_services_DataObjectsAPIService_Update(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	objectName := "test-do-update-" + common.GenerateRandomString(6)

	predefinedPattern := security_services.DataObjectsPatternTypePredefinedPatternInner{
		Name:     common.StringPtr("ABA-Routing-Number"),
		FileType: []string{"text/html"},
	}
	predefined := security_services.DataObjectsPatternTypePredefined{
		Pattern: []security_services.DataObjectsPatternTypePredefinedPatternInner{predefinedPattern},
	}
	patternType := security_services.DataObjectsPatternType{
		Predefined: &predefined,
	}

	dataObject := security_services.DataObjects{
		Folder:      common.StringPtr("ngfw-shared"),
		Name:        common.StringPtr(objectName),
		PatternType: &patternType,
	}

	createRes, _, err := client.DataObjectsAPI.CreateDataObjects(context.Background()).DataObjects(dataObject).Execute()
	require.NoError(t, err, "Failed to create Data Object for update test")
	createdObjectID := createRes.Id
	require.NotEmpty(t, createdObjectID, "Created object ID should not be empty")

	defer func() {
		t.Logf("Cleaning up Data Object with ID: %s", *createdObjectID)
		_, errDel := client.DataObjectsAPI.DeleteDataObjectsByID(context.Background(), *createdObjectID).Execute()
		require.NoError(t, errDel, "Failed to delete Data Object during cleanup")
	}()

	updatedDescription := "Updated data object description"

	// Update with a new pattern
	updatedPredefinedPattern := security_services.DataObjectsPatternTypePredefinedPatternInner{
		Name:     common.StringPtr("ABA-Routing-Number"),
		FileType: []string{"any"},
	}
	updatedPredefined := security_services.DataObjectsPatternTypePredefined{
		Pattern: []security_services.DataObjectsPatternTypePredefinedPatternInner{updatedPredefinedPattern},
	}
	updatedPatternType := security_services.DataObjectsPatternType{
		Predefined: &updatedPredefined,
	}

	updatedObject := security_services.DataObjects{
		Name:        common.StringPtr(objectName),
		Folder:      common.StringPtr("ngfw-shared"),
		Description: common.StringPtr(updatedDescription),
		PatternType: &updatedPatternType,
	}

	updateRes, httpResUpdate, errUpdate := client.DataObjectsAPI.UpdateDataObjectsByID(context.Background(), *createdObjectID).DataObjects(updatedObject).Execute()
	require.NoError(t, errUpdate, "Failed to update Data Object")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, updatedDescription, *updateRes.Description, "Description should be updated")
}

// Test_security_services_DataObjectsAPIService_List tests listing Data Objects.
func Test_security_services_DataObjectsAPIService_List(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	objectName := "test-do-list-" + common.GenerateRandomString(6)

	predefinedPattern := security_services.DataObjectsPatternTypePredefinedPatternInner{
		Name:     common.StringPtr("ABA-Routing-Number"),
		FileType: []string{"text/html"},
	}
	predefined := security_services.DataObjectsPatternTypePredefined{
		Pattern: []security_services.DataObjectsPatternTypePredefinedPatternInner{predefinedPattern},
	}
	patternType := security_services.DataObjectsPatternType{
		Predefined: &predefined,
	}

	dataObject := security_services.DataObjects{
		Folder:      common.StringPtr("ngfw-shared"),
		Name:        common.StringPtr(objectName),
		PatternType: &patternType,
	}

	createRes, _, err := client.DataObjectsAPI.CreateDataObjects(context.Background()).DataObjects(dataObject).Execute()
	require.NoError(t, err, "Failed to create Data Object for list test")
	createdObjectID := createRes.Id
	require.NotEmpty(t, createdObjectID, "Created object ID should not be empty")

	defer func() {
		t.Logf("Cleaning up Data Object with ID: %s", *createdObjectID)
		_, errDel := client.DataObjectsAPI.DeleteDataObjectsByID(context.Background(), *createdObjectID).Execute()
		require.NoError(t, errDel, "Failed to delete Data Object during cleanup")
	}()

	listRes, httpResList, errList := client.DataObjectsAPI.ListDataObjects(context.Background()).Folder("ngfw-shared").Limit(100).Execute()
	require.NoError(t, errList, "Failed to list Data Objects")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	foundObject := false
	for _, o := range listRes.Data {
		if o.Name != nil && *o.Name == objectName {
			foundObject = true
			break
		}
	}
	assert.True(t, foundObject, "Created Data Object should be found in the list")
}

// Test_security_services_DataObjectsAPIService_DeleteByID tests deleting a Data Object by its ID.
func Test_security_services_DataObjectsAPIService_DeleteByID(t *testing.T) {
	client := SetupSecuritySvcTestClient(t)
	objectName := "test-do-delete-" + common.GenerateRandomString(6)

	predefinedPattern := security_services.DataObjectsPatternTypePredefinedPatternInner{
		Name:     common.StringPtr("ABA-Routing-Number"),
		FileType: []string{"text/html"},
	}
	predefined := security_services.DataObjectsPatternTypePredefined{
		Pattern: []security_services.DataObjectsPatternTypePredefinedPatternInner{predefinedPattern},
	}
	patternType := security_services.DataObjectsPatternType{
		Predefined: &predefined,
	}

	dataObject := security_services.DataObjects{
		Folder:      common.StringPtr("ngfw-shared"),
		Name:        common.StringPtr(objectName),
		PatternType: &patternType,
	}

	createRes, _, err := client.DataObjectsAPI.CreateDataObjects(context.Background()).DataObjects(dataObject).Execute()
	require.NoError(t, err, "Failed to create Data Object for delete test")
	createdObjectID := createRes.Id
	require.NotEmpty(t, createdObjectID, "Created object ID should not be empty")

	httpResDel, errDel := client.DataObjectsAPI.DeleteDataObjectsByID(context.Background(), *createdObjectID).Execute()
	require.NoError(t, errDel, "Failed to delete Data Object")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status on successful delete")
}
