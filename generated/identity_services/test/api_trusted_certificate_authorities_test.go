/*
Identity Services Testing TrustedCertificateAuthoritiesAPIService

Note: TrustedCertificateAuthorities is a read-only resource, so this test
only validates the Fetch method against existing CAs in the system.
*/
package identity_services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_identity_services_TrustedCertificateAuthoritiesAPIService_List tests listing trusted certificate authorities
func Test_identity_services_TrustedCertificateAuthoritiesAPIService_List(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// List all trusted certificate authorities
	listReq := client.TrustedCertificateAuthoritiesAPI.ListTrustedCertificateAuthorities(context.Background())
	resp, httpResp, err := listReq.Execute()

	// Verify the response
	require.NoError(t, err, "Failed to list trusted certificate authorities")
	require.NotNil(t, httpResp, "HTTP response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected 200 OK status")
	require.NotNil(t, resp, "Response should not be nil")

	// Note: TrustedCertificateAuthorities is a read-only resource
	// We can only verify that the list operation works
	t.Logf("[SUCCESS] Listed trusted certificate authorities (count: %d)", len(resp.Data))
}

// Test_identity_services_TrustedCertificateAuthoritiesAPIService_FetchTrustedCertificateAuthorities tests the FetchTrustedCertificateAuthorities convenience method
func Test_identity_services_TrustedCertificateAuthoritiesAPIService_FetchTrustedCertificateAuthorities(t *testing.T) {
	// Setup the authenticated client
	client := SetupIdentitySvcTestClient(t)

	// Note: TrustedCertificateAuthorities is read-only, so we can't create test objects.
	// Instead, we'll list existing CAs and test fetching one if available.

	// List existing trusted CAs
	listReq := client.TrustedCertificateAuthoritiesAPI.ListTrustedCertificateAuthorities(context.Background())
	listRes, _, err := listReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list trusted certificate authorities")

	// If there are existing CAs, test fetching one
	if listRes.Data != nil && len(listRes.Data) > 0 {
		// Get the first CA's name
		firstCA := listRes.Data[0]
		if firstCA.Name != nil && *firstCA.Name != "" {
			caName := *firstCA.Name

			// Test 1: Fetch existing CA by name
			fetchedObj, err := client.TrustedCertificateAuthoritiesAPI.FetchTrustedCertificateAuthorities(
				context.Background(),
				caName,
				nil, // folder
				nil, // snippet
				nil, // device
			)

			// Verify successful fetch
			require.NoError(t, err, "Failed to fetch trusted certificate authority by name")
			require.NotNil(t, fetchedObj, "Fetched object should not be nil")
			if fetchedObj.Name != nil {
				assert.Equal(t, caName, *fetchedObj.Name, "Fetched object name should match")
			}
			t.Logf("[SUCCESS] FetchTrustedCertificateAuthorities found object: %v", fetchedObj.Name)
		} else {
			t.Skip("No named trusted certificate authorities available for testing")
		}
	} else {
		t.Skip("No trusted certificate authorities available for testing")
	}

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.TrustedCertificateAuthoritiesAPI.FetchTrustedCertificateAuthorities(
		context.Background(),
		"non-existent-ca-xyz-12345",
		nil,
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchTrustedCertificateAuthorities correctly returned nil for non-existent object")
}
