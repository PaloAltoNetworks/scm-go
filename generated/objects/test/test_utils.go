// Package objects
/*
Testing utilities for objects API
Shared utilities for testing objects API services
*/
package objects

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	setup "github.com/paloaltonetworks/scm-go"
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/objects"
)

func SetupObjectSvcTestClient(t *testing.T) *objects.APIClient {
	configPath := common.GetConfigPath()
	setupClient := &setup.Client{
		AuthFile:         configPath,
		CheckEnvironment: false,
	}

	fmt.Printf("Using config file: %s\n", setupClient.AuthFile)

	// Setup the client configuration
	err := setupClient.Setup()
	require.NoError(t, err, "Failed to setup client")

	// Refresh JWT token
	ctx := context.Background()
	if setupClient.Jwt == "" {
		fmt.Printf("\n****************\nGetting tokens\n***********************\n")
		err = setupClient.RefreshJwt(ctx)
		if err != nil {
			// Print detailed error information
			fmt.Printf("=== JWT REFRESH ERROR ===\n")
			fmt.Printf("Error: %v\n", err)
			fmt.Printf("Error Type: %T\n", err)
			fmt.Printf("Error String: %s\n", err.Error())
			fmt.Printf("=========================\n")
		}
		// Fail the test only after all retries have been exhausted.
		require.NoError(t, err, "Failed to refresh JWT after multiple retries")
	}

	return setup.GetObjectsAPIClient(setupClient)
}

// printAPIError prints formatted API error response from error object's body
func printAPIError(err *objects.GenericOpenAPIError) {
	if err == nil {
		return
	}
	fmt.Printf("=== API ERROR RESPONSE ===\n")
	fmt.Printf("Error: %v\n", err)
	bodyBytes := err.Body()
	if bodyBytes == nil {
		fmt.Printf("No body found in error object\n")
		fmt.Printf("===========================\n\n")
		return
	}
	if len(bodyBytes) == 0 {
		fmt.Printf("No body found in error object\n")
		fmt.Printf("===========================\n\n")
		return
	}
	// Print raw JSON response
	fmt.Printf("Raw Error Body:\n%s\n", string(bodyBytes))
	fmt.Printf("===========================\n\n")
}

// handleAPIError is a utility method to consistently handle and print API errors
func handleAPIError(err error) {
	if err == nil {
		return
	}
	// Print detailed error information if it's a GenericOpenAPIError
	var apiErr *objects.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		printAPIError(apiErr)
	} else {
		// Print the regular error if it's not a GenericOpenAPIError
		fmt.Printf("Non-API Error: %v\n", err)
	}
}

// createTestAddress creates a test address and returns its ID
func createTestAddress(t *testing.T, client *objects.APIClient, name string, ipNetmask string) string {
	address := objects.Addresses{
		Description: common.StringPtr("Test address for testing purposes"),
		Folder:      common.StringPtr("Shared"),
		IpNetmask:   common.StringPtr(ipNetmask),
		Name:        name,
	}
	req := client.AddressesAPI.CreateAddresses(context.Background()).Addresses(address)
	res, _, err := req.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test address")
	require.NotNil(t, res, "Create address response should not be nil")
	require.NotEmpty(t, res.Id, "Created address should have an ID")
	t.Logf("Created test address: %s with ID: %s", name, res.Id)
	return res.Id
}

// deleteTestAddress deletes a test address by ID.
func deleteTestAddress(t *testing.T, client *objects.APIClient, addressID string, addressName string) {
	reqDel := client.AddressesAPI.DeleteAddressesByID(context.Background(), addressID)
	httpRes, err := reqDel.Execute()

	// FIX: Don't fail the test if the object is already gone (404 Not Found).
	if err != nil {
		if httpRes != nil && httpRes.StatusCode != 404 {
			handleAPIError(err)
			require.Fail(t, "Failed to delete test address", "ID: %s", addressID)
		}
		t.Logf("Test address already deleted (ID: %s)", addressID)
		return
	}

	require.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status for delete")
	if addressName != "" {
		t.Logf("Deleted test address: %s (ID: %s)", addressName, addressID)
	} else {
		t.Logf("Deleted test address: (ID: %s)", addressID)
	}
}

// deleteTestAddressGroup deletes a test address group by ID.
func deleteTestAddressGroup(t *testing.T, client *objects.APIClient, addressGroupID string, addressGroupName string) {
	reqDel := client.AddressGroupsAPI.DeleteAddressGroupsByID(context.Background(), addressGroupID)
	httpResDel, errDel := reqDel.Execute()

	// FIX: Don't fail the test if the object is already gone (404 Not Found).
	if errDel != nil {
		if httpResDel != nil && httpResDel.StatusCode != 404 {
			handleAPIError(errDel)
			require.Fail(t, "Failed to delete test address group", "ID: %s", addressGroupID)
		}
		t.Logf("Test address group already deleted (ID: %s)", addressGroupID)
		return
	}

	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
	if addressGroupName != "" {
		t.Logf("Deleted test address group: %s (ID: %s)", addressGroupName, addressGroupID)
	} else {
		t.Logf("Deleted test address group: (ID: %s)", addressGroupID)
	}
}

// deleteTestTag deletes a test tag by ID.
func deleteTestTag(t *testing.T, client *objects.APIClient, tagID string, tagName string) {
	reqDel := client.TagsAPI.DeleteTagsByID(context.Background(), tagID)
	httpRes, err := reqDel.Execute()

	// Don't fail the test if the object is already gone (404 Not Found).
	if err != nil {
		if httpRes != nil && httpRes.StatusCode != 404 {
			handleAPIError(err)
			require.Fail(t, "Failed to delete test tag", "ID: %s", tagID)
		}
		t.Logf("Test tag already deleted (ID: %s)", tagID)
		return
	}

	require.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status for delete")
	if tagName != "" {
		t.Logf("Deleted test tag: %s (ID: %s)", tagName, tagID)
	} else {
		t.Logf("Deleted test tag: (ID: %s)", tagID)
	}
}
