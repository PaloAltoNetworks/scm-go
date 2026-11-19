// Package cloud_ngfw_gcp test utilities
package cloud_ngfw_gcp

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	setup "github.com/paloaltonetworks/scm-go"
	"github.com/paloaltonetworks/scm-go/common"
	pkg "github.com/paloaltonetworks/scm-go/generated/cloud_ngfw_gcp"
)

// setupTestClient creates an authenticated API client for the cloud_ngfw_gcp package.
// It uses GetCloudNgfwGcpAPIClient which sets the correct host and JWT refresh transport.
func setupTestClient(t *testing.T) *pkg.APIClient {
	t.Helper()
	configPath := common.GetConfigPath()
	setupClient := &setup.Client{
		AuthFile:         configPath,
		CheckEnvironment: false,
	}
	fmt.Printf("Using config file: %s\n", setupClient.AuthFile)

	err := setupClient.Setup()
	require.NoError(t, err, "Failed to setup SCM client")

	ctx := context.Background()
	if setupClient.Jwt == "" {
		err = setupClient.RefreshJwt(ctx)
		require.NoError(t, err, "Failed to refresh JWT token")
	}

	return setup.GetCloudNgfwGcpAPIClient(setupClient)
}

// handleAPIError prints a failed call. Error bodies never carry a link code, so printing them is safe.
func handleAPIError(err error) {
	if err == nil {
		return
	}
	var apiErr *pkg.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		fmt.Printf("API Error: %v\nBody: %s\n", err, string(apiErr.Body()))
	} else {
		fmt.Printf("API Error: %v\n", err)
	}
}

// linkErrorCode returns ResponseStatus.Code from the link service's error envelope, or "" when
// the error did not come from the service (an API gateway error, for example).
func linkErrorCode(err error) string {
	var apiErr *pkg.GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		return ""
	}
	if m, ok := apiErr.Model().(pkg.LinkErrorResponse); ok && m.ResponseStatus != nil {
		return m.ResponseStatus.GetCode()
	}
	return ""
}
