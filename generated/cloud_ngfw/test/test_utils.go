// Package cloud_ngfw test utilities
package cloud_ngfw

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	setup "github.com/paloaltonetworks/scm-go"
	"github.com/paloaltonetworks/scm-go/common"
	pkg "github.com/paloaltonetworks/scm-go/generated/cloud_ngfw"
)

// setupTestClient creates an authenticated API client for the cloud_ngfw package.
// It uses GetCloudNgfwAPIClient which sets the correct host and JWT refresh transport.
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

	return setup.GetCloudNgfwAPIClient(setupClient)
}

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
