// Package config_setup
/*
Testing utilities for config_setup API
Shared utilities for testing config_setup API services
*/
package config_setup

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	setup "github.com/paloaltonetworks/scm-go"
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/config_setup"
)

func SetupConfigSvcTestClient(t *testing.T) *config_setup.APIClient {
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
		maxRetries := 3
		retryDelay := 2 * time.Second
		for i := 0; i < maxRetries; i++ {
			err = setupClient.RefreshJwt(ctx)
			if err == nil {
				break // Success, exit the loop
			}
			t.Logf("Failed to refresh JWT (attempt %d/%d), retrying in %v... Error: %v", i+1, maxRetries, retryDelay, err)
			time.Sleep(retryDelay)
		}
		// Fail the test only after all retries have been exhausted.
		require.NoError(t, err, "Failed to refresh JWT after multiple retries")
	}

	// Create the config_setup API client
	config := config_setup.NewConfiguration()
	config.Host = setupClient.GetHost()
	config.Scheme = "https"

	// Create a custom HTTP client that includes the JWT token and logging
	if setupClient.HttpClient == nil {
		setupClient.HttpClient = &http.Client{}
	}

	// Wrap the transport with our logging transport
	if setupClient.HttpClient.Transport == nil {
		setupClient.HttpClient.Transport = http.DefaultTransport
	}
	setupClient.HttpClient.Transport = &common.LoggingRoundTripper{
		Wrapped: setupClient.HttpClient.Transport,
	}

	config.HTTPClient = setupClient.HttpClient

	// Set up the default header with JWT
	config.DefaultHeader = make(map[string]string)
	config.DefaultHeader["Authorization"] = "Bearer " + setupClient.Jwt
	config.DefaultHeader["x-auth-jwt"] = setupClient.Jwt

	apiClient := config_setup.NewAPIClient(config)
	return apiClient
}

// printAPIError prints formatted API error response from error object's body
func printAPIError(err *config_setup.GenericOpenAPIError) {
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
	var apiErr *config_setup.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		printAPIError(apiErr)
	} else {
		// Print the regular error if it's not a GenericOpenAPIError
		fmt.Printf("Non-API Error: %v\n", err)
	}
}
