// Package deployment_services
/*
Testing utilities for deployment_services API
Shared utilities for testing deployment_services API services
*/
package deployment_services

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	setup "github.com/paloaltonetworks/scm-go"
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/deployment_services"
)

func SetupDeploymentSvcTestClient(t *testing.T) *deployment_services.APIClient {
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

	return setup.GetDeployment_servicesAPIClient(setupClient)
}

// printAPIError prints formatted API error response from error object's body
func printAPIError(err *deployment_services.GenericOpenAPIError) {
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
	var apiErr *deployment_services.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		printAPIError(apiErr)
	} else {
		// Print the regular error if it's not a GenericOpenAPIError
		fmt.Printf("Non-API Error: %v\n", err)
	}
}
