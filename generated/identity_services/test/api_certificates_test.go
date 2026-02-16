/*
Identity Services Testing CertificatesAPIService
*/
package identity_services

import (
	"testing"
)

// Test_identity_services_CertificatesAPIService_FetchCertificates tests the FetchCertificates convenience method
func Test_identity_services_CertificatesAPIService_FetchCertificates(t *testing.T) {
	// Note: Certificates API requires an existing CA (Certificate Authority) to sign new certificates.
	// The 'signed_by' field must reference an existing CA name from TrustedCertificateAuthorities.
	// Without a standardized/guaranteed CA name across environments, this test cannot reliably run.
	// Python SDK (scm-python) also does not include tests for this API for the same reason.
	t.Skip("Certificates API requires existing CA infrastructure - no standard CA name available across environments")
}
