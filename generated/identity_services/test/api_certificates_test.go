/*
Identity Services Testing CertificatesAPIService
*/
package identity_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
)

// Test_identity_services_CertificatesAPIService_FetchCertificates tests the FetchCertificates convenience method
func Test_identity_services_CertificatesAPIService_FetchCertificates(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Test: Fetch non-existent object (should return nil, nil)
	// Cannot test positive path because Create requires existing CA infrastructure
	notFound, err := client.CertificatesAPI.FetchCertificates(
		context.Background(),
		"non-existent-cert-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchCertificates correctly returned nil for non-existent object")
}
