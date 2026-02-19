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
)

// Test_identity_services_LocalUsersAPIService_FetchLocalUsers tests the FetchLocalUsers convenience method
// This is a read-only test (Create blocked by non-pointer Id issue)
func Test_identity_services_LocalUsersAPIService_FetchLocalUsers(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Test: Fetch non-existent object (should return nil, nil)
	// Cannot test positive path because Create is blocked by non-pointer Id field
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
