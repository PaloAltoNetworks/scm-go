/*
 * Config Setup Testing
 *
 * VariablesAPIService
 */
/*
Objects Testing VariablesAPIService
*/
package config_setup

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/config_setup"
)

// Test_config_setup_VariableAPIService_Create tests the creation of an variable
// This test creates a new variable and then deletes it to ensure proper cleanup
func Test_config_setup_VariableAPIService_Create(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create a valid variable with unique name to avoid conflicts
	createdVariableName := "$test-variable-" + common.GenerateRandomString(10)
	variable := config_setup.Variables{
		Description: common.StringPtr("Test variable for create API testing"),
		Folder:      common.StringPtr("Shared"),                  // Using Shared folder scope
		Value:       common.StringPtr("test.create.example.com"), // FQDN-based variable
		Type:        "fqdn",
		Name:        createdVariableName, // Unique test name
	}

	fmt.Printf("Creating variable with name: %s\n", variable.Name)

	// Make the create request to the API
	req := client.VariablesAPI.CreateVariable(context.Background()).Variables(variable)
	res, httpRes, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create variable")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")

	// Assert response object properties
	createdVariableValue := fmt.Sprintf("%v", res.Value.(string))

	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, createdVariableName, res.Name, "Created variable name should match")
	assert.Equal(t, common.StringPtr("Test variable for create API testing"), res.Description, "Description should match")
	assert.True(t, *res.Folder == "Shared" || *res.Folder == "Prisma Access", "Folder should be 'Shared' or 'Shared'")
	assert.Equal(t, "test.create.example.com", createdVariableValue, "FQDN should match")
	assert.NotEmpty(t, res.Id, "Created variable should have an ID")

	// Use the ID from the response object
	createdVariableID := res.Id
	t.Logf("Successfully created variable: %s with ID: %s", variable.Name, createdVariableID)

	// Cleanup: Delete the created variable to maintain test isolation
	reqDel := client.VariablesAPI.DeleteVariableByID(context.Background(), createdVariableID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete variable during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up variable: %s", createdVariableID)
}

// Test_config_setup_VariablesAPIService_GetByID tests retrieving an variable by its ID
// This test creates an variable, retrieves it by ID, then deletes it
func Test_config_setup_VariablesAPIService_GetByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create an variable first to have something to retrieve
	createdVariableName := "$test-variable-getbyid-" + common.GenerateRandomString(10)
	variable := config_setup.Variables{
		Description: common.StringPtr("Test variable for get by ID API testing"),
		Folder:      common.StringPtr("Shared"),                   // Using Shared folder scope
		Value:       common.StringPtr("test.getbyid.example.com"), // FQDN-based variable
		Type:        "fqdn",
		Name:        createdVariableName, // Unique test name
	}

	// Create the variable via API
	req := client.VariablesAPI.CreateVariable(context.Background()).Variables(variable)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create variable for get test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdVariableID := createRes.Id
	require.NotEmpty(t, createdVariableID, "Created variable should have an ID")

	// Test Get by ID operation
	reqGetById := client.VariablesAPI.GetVariableByID(context.Background(), createdVariableID)
	getRes, httpResGet, err := reqGetById.Execute()
	if err != nil {
		handleAPIError(err)
	}

	// Verify the get operation was successful
	require.NoError(t, err, "Failed to get variable by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	variableValue := fmt.Sprintf("%v", getRes.Value.(string))

	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, createdVariableName, getRes.Name, "variable name should match")
	assert.Equal(t, common.StringPtr("Test variable for get by ID API testing"), getRes.Description, "Description should match")
	assert.True(t, *getRes.Folder == "Shared" || *getRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, "test.getbyid.example.com", variableValue, "FQDN should match")
	assert.Equal(t, createdVariableID, getRes.Id, "variable ID should match")

	t.Logf("Successfully retrieved variable: %s", getRes.Name)

	// Cleanup: Delete the created variable
	reqDel := client.VariablesAPI.DeleteVariableByID(context.Background(), createdVariableID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete variable during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up variable: %s", createdVariableID)
}

// Test_config_setup_VariablesAPIService_Update tests updating an existing variable
// This test creates an variable, updates it, then deletes it
func Test_config_setup_VariablesAPIService_Update(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create an variable first to have something to update
	createdVariableName := "$test-variable-update-" + common.GenerateRandomString(10)
	variable := config_setup.Variables{
		Description: common.StringPtr("Test variable for update API testing"),
		Folder:      common.StringPtr("Shared"),                  // Using Shared folder scope
		Value:       common.StringPtr("test.update.example.com"), // Initial FQDN
		Type:        "fqdn",
		Name:        createdVariableName, // Unique test name
	}

	// Create the variable via API
	req := client.VariablesAPI.CreateVariable(context.Background()).Variables(variable)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create variable for update test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdVariableID := createRes.Id
	require.NotEmpty(t, createdVariableID, "Created variable should have an ID")

	// Test Update operation with modified fields
	updatedVariable := config_setup.Variables{
		Description: common.StringPtr("Updated test variable description"), // Updated description
		Folder:      common.StringPtr("Shared"),                            // Keep same folder scope
		Value:       common.StringPtr("updated.test.example.com"),          // Updated FQDN
		Type:        "fqdn",
		Name:        createdVariableName, // Keep same name (required for update)
	}

	reqUpdate := client.VariablesAPI.UpdateVariableByID(context.Background(), createdVariableID).Variables(updatedVariable)
	updateRes, httpResUpdate, errUpdate := reqUpdate.Execute()
	if errUpdate != nil {
		handleAPIError(errUpdate)
	}

	// Verify the update operation was successful
	require.NoError(t, errUpdate, "Failed to update variable")
	assert.Equal(t, 200, httpResUpdate.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	variableValue := fmt.Sprintf("%v", updateRes.Value.(string))

	require.NotNil(t, updateRes, "Update response should not be nil")
	assert.Equal(t, createdVariableName, updateRes.Name, "variable name should remain the same")
	assert.Equal(t, common.StringPtr("Updated test variable description"), updateRes.Description, "Description should be updated")
	assert.True(t, *updateRes.Folder == "Shared" || *updateRes.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
	assert.Equal(t, "updated.test.example.com", variableValue, "FQDN should be updated")
	assert.Equal(t, createdVariableID, updateRes.Id, "variable ID should remain the same")

	t.Logf("Successfully updated variable: %s", createdVariableName)

	// Cleanup: Delete the created variable
	reqDel := client.VariablesAPI.DeleteVariableByID(context.Background(), createdVariableID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete variable during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up variable: %s", createdVariableID)
}

// Test_config_setup_VariablesAPIService_List tests listing variables with folder filter
// This test creates an variable, lists variables to verify it's included, then deletes it
func Test_config_setup_VariablesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create an variable first to have something to list
	createdVariableName := "$test-variable-list-" + common.GenerateRandomString(10)
	variable := config_setup.Variables{
		Description: common.StringPtr("Test variable for list API testing"),
		Folder:      common.StringPtr("Shared"),                // Using Shared folder scope
		Value:       common.StringPtr("test.list.example.com"), // FQDN-based variable
		Type:        "fqdn",
		Name:        createdVariableName, // Unique test name
	}

	// Create the variable via API
	req := client.VariablesAPI.CreateVariable(context.Background()).Variables(variable)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create variable for list test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdVariableID := createRes.Id
	require.NotEmpty(t, createdVariableID, "Created variable should have an ID")

	// Test List operation with folder filter
	reqList := client.VariablesAPI.ListVariables(context.Background()).Folder("Shared")
	listRes, httpResList, errList := reqList.Execute()
	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list variables")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")

	// Assert response object properties
	require.NotNil(t, listRes, "List response should not be nil")
	assert.NotNil(t, listRes.Data, "List response data should not be nil")
	assert.Greater(t, len(listRes.Data), 0, "Should have at least one variable in the list")

	// Verify our created variable is in the list
	foundVariable := false
	for _, tmpVar := range listRes.Data {
		if tmpVar.Name == createdVariableName {
			variableValue := fmt.Sprintf("%v", tmpVar.Value)
			foundVariable = true
			assert.Equal(t, common.StringPtr("Test variable for list API testing"), tmpVar.Description, "Description should match")
			assert.True(t, *tmpVar.Folder == "Shared" || *tmpVar.Folder == "Prisma Access", "Folder should be 'Shared' or 'Prisma Access'")
			assert.Equal(t, "test.list.example.com", variableValue, "FQDN should match")
			break
		}
	}
	assert.True(t, foundVariable, "Created variable should be found in the list")

	t.Logf("Successfully listed variables, found created variable: %s", createdVariableName)

	// Cleanup: Delete the created variable
	reqDel := client.VariablesAPI.DeleteVariableByID(context.Background(), createdVariableID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete variable during cleanup")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status for delete")

	t.Logf("Successfully cleaned up variable: %s", createdVariableID)
}

// Test_config_setup_VariablesAPIService_DeleteByID tests deleting an variable by its ID
// This test creates an variable, deletes it, then verifies the deletion was successful
func Test_config_setup_VariablesAPIService_DeleteByID(t *testing.T) {
	// Setup the authenticated client
	client := SetupDeploymentSvcTestClient(t)

	// Create an variable first to have something to delete
	createdVariableName := "$test-variable-delete-" + common.GenerateRandomString(10)
	variable := config_setup.Variables{
		Description: common.StringPtr("Test variable for delete API testing"),
		Folder:      common.StringPtr("Shared"),                  // Using Shared folder scope
		Value:       common.StringPtr("test.delete.example.com"), // FQDN-based variable
		Type:        "fqdn",
		Name:        createdVariableName, // Unique test name
	}

	// Create the variable via API
	req := client.VariablesAPI.CreateVariable(context.Background()).Variables(variable)
	createRes, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create variable for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdVariableID := createRes.Id
	require.NotEmpty(t, createdVariableID, "Created variable should have an ID")

	// Test Delete by ID operation
	reqDel := client.VariablesAPI.DeleteVariableByID(context.Background(), createdVariableID)
	httpResDel, errDel := reqDel.Execute()
	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete variable")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")

	t.Logf("Successfully deleted variable: %s", createdVariableID)
}
