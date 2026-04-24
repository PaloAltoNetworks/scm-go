/*
Identity Services Testing CertificatesAPIService
*/
package identity_services

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/identity_services"
)

// Test_identity_services_CertificatesAPIService_List tests listing certificates
func Test_identity_services_CertificatesAPIService_List(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Test List operation
	fmt.Println("Attempting to list Certificates")
	reqList := client.CertificatesAPI.ListCertificates(context.Background()).Folder("Prisma Access").Limit(100)
	listRes, httpResList, errList := reqList.Execute()

	if errList != nil {
		handleAPIError(errList)
	}

	// Verify the list operation was successful
	require.NoError(t, errList, "Failed to list certificates")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")

	// Log the results
	fmt.Printf("Retrieved %d Certificates\n", len(listRes.Data))
	for i, cert := range listRes.Data {
		if i < 5 { // Only log first 5
			t.Logf("  Certificate: %s (ID: %s)", *cert.Name, *cert.Id)
		}
	}
	t.Logf("[SUCCESS] ListCertificates returned %d certificates", len(listRes.Data))
}

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
<<<<<<< HEAD

// Test_identity_services_CertificatesAPIService_FetchExisting tests fetching an existing certificate by name
func Test_identity_services_CertificatesAPIService_FetchExisting(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// First, list certificates to get a name to fetch
	reqList := client.CertificatesAPI.ListCertificates(context.Background()).Folder("Prisma Access").Limit(10)
	listRes, _, errList := reqList.Execute()
	require.NoError(t, errList, "Failed to list certificates for FetchExisting test")

	if len(listRes.Data) == 0 {
		t.Skip("No certificates available to test FetchExisting - skipping")
		return
	}

	// Fetch the first certificate by name
	certToFetch := listRes.Data[0]
	certName := *certToFetch.Name

	fmt.Printf("Attempting to fetch Certificate with name: %s\n", certName)
	fetchedCert, err := client.CertificatesAPI.FetchCertificates(
		context.Background(),
		certName,
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)

	require.NoError(t, err, "Failed to fetch certificate by name")
	require.NotNil(t, fetchedCert, "Fetched certificate should not be nil")
	assert.Equal(t, certName, *fetchedCert.Name, "Fetched certificate name should match")
	assert.Equal(t, certToFetch.Id, fetchedCert.Id, "Fetched certificate ID should match")
	t.Logf("[SUCCESS] FetchCertificates found certificate: %s (ID: %s)", *fetchedCert.Name, *fetchedCert.Id)
}

// Test_identity_services_CertificatesAPIService_Create tests creating a certificate
func Test_identity_services_CertificatesAPIService_Create(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Create a unique certificate name
	certName := "test-cert-create-" + common.GenerateRandomString(6)

	// Create algorithm with RSA 2048 bits
	rsaBits := float32(2048)
	algorithm := identity_services.CertificatesPostAlgorithm{
		RsaNumberOfBits: &rsaBits,
	}

	// Create the certificate payload
	cert := identity_services.CertificatesPost{
		Algorithm:       algorithm,
		CertificateName: certName,
		CommonName:      certName + ".example.com",
		Digest:          "sha256",
		SignedBy:        "Root CA",
		Folder:          common.StringPtr("Prisma Access"),
	}

	fmt.Printf("Attempting to create Certificate with name: %s\n", certName)
	req := client.CertificatesAPI.CreateCertificates(context.Background()).CertificatesPost(cert)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	// Verify the creation was successful
	require.NoError(t, err, "Failed to create Certificate")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created certificate should have an ID")

	createdCertID := *res.Id
	t.Logf("Successfully created Certificate: %s with ID: %s", certName, createdCertID)

	// Cleanup: Delete the created certificate
	defer func() {
		t.Logf("Cleaning up Certificate with ID: %s", createdCertID)
		reqDel := client.CertificatesAPI.DeleteCertificatesByID(context.Background(), createdCertID)
		httpResDel, errDel := reqDel.Execute()
		if errDel != nil {
			handleAPIError(errDel)
		}
		require.NoError(t, errDel, "Failed to delete certificate during cleanup")
		assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK for cleanup delete")
	}()

	// Assert response object properties
	assert.Equal(t, certName, *res.Name, "Created certificate name should match")
	t.Logf("[SUCCESS] CreateCertificates created certificate: %s", certName)
}

// Test_identity_services_CertificatesAPIService_Delete tests deleting a certificate
func Test_identity_services_CertificatesAPIService_Delete(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// First, create a certificate to delete
	certName := "test-cert-delete-" + common.GenerateRandomString(6)

	// Create algorithm with RSA 2048 bits
	rsaBits := float32(2048)
	algorithm := identity_services.CertificatesPostAlgorithm{
		RsaNumberOfBits: &rsaBits,
	}

	// Create the certificate payload
	cert := identity_services.CertificatesPost{
		Algorithm:       algorithm,
		CertificateName: certName,
		CommonName:      certName + ".example.com",
		Digest:          "sha256",
		SignedBy:        "Root CA",
		Folder:          common.StringPtr("Prisma Access"),
	}

	fmt.Printf("Creating Certificate for delete test with name: %s\n", certName)
	reqCreate := client.CertificatesAPI.CreateCertificates(context.Background()).CertificatesPost(cert)
	createRes, _, errCreate := reqCreate.Execute()

	if errCreate != nil {
		handleAPIError(errCreate)
	}
	require.NoError(t, errCreate, "Failed to create certificate for delete test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdCertID := *createRes.Id
	require.NotEmpty(t, createdCertID, "Created certificate ID should not be empty")

	t.Logf("Created Certificate for Delete test with ID: %s", createdCertID)

	// Test Delete by ID operation
	fmt.Printf("Attempting to delete Certificate with ID: %s\n", createdCertID)
	reqDel := client.CertificatesAPI.DeleteCertificatesByID(context.Background(), createdCertID)
	httpResDel, errDel := reqDel.Execute()

	if errDel != nil {
		handleAPIError(errDel)
	}

	// Verify the delete operation was successful
	require.NoError(t, errDel, "Failed to delete certificate")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
	t.Logf("[SUCCESS] DeleteCertificatesByID deleted certificate: %s", createdCertID)

	// Verify the certificate no longer exists
	fetchedCert, errFetch := client.CertificatesAPI.FetchCertificates(
		context.Background(),
		certName,
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, errFetch, "Fetch after delete should not error")
	assert.Nil(t, fetchedCert, "Certificate should not exist after deletion")
	t.Logf("[SUCCESS] Verified certificate no longer exists after deletion")
}

// Test_identity_services_CertificatesAPIService_Export tests exporting a certificate
func Test_identity_services_CertificatesAPIService_Export(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// First, create a certificate to export
	certName := "test-cert-export-" + common.GenerateRandomString(6)

	// Create algorithm with RSA 2048 bits
	rsaBits := float32(2048)
	algorithm := identity_services.CertificatesPostAlgorithm{
		RsaNumberOfBits: &rsaBits,
	}

	// Create the certificate payload
	cert := identity_services.CertificatesPost{
		Algorithm:       algorithm,
		CertificateName: certName,
		CommonName:      certName + ".example.com",
		Digest:          "sha256",
		SignedBy:        "Root CA",
		Folder:          common.StringPtr("Prisma Access"),
	}

	fmt.Printf("Creating Certificate for export test with name: %s\n", certName)
	reqCreate := client.CertificatesAPI.CreateCertificates(context.Background()).CertificatesPost(cert)
	createRes, _, errCreate := reqCreate.Execute()

	if errCreate != nil {
		handleAPIError(errCreate)
	}
	require.NoError(t, errCreate, "Failed to create certificate for export test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdCertID := *createRes.Id
	require.NotEmpty(t, createdCertID, "Created certificate ID should not be empty")

	t.Logf("Created Certificate for Export test with ID: %s", createdCertID)

	// Cleanup: Delete the created certificate after test
	defer func() {
		t.Logf("Cleaning up Certificate with ID: %s", createdCertID)
		reqDel := client.CertificatesAPI.DeleteCertificatesByID(context.Background(), createdCertID)
		httpResDel, errDel := reqDel.Execute()
		if errDel != nil {
			handleAPIError(errDel)
		}
		require.NoError(t, errDel, "Failed to delete certificate during cleanup")
		assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK for cleanup delete")
	}()

	// Test Export operation
	fmt.Printf("Attempting to export Certificate with ID: %s\n", createdCertID)
	exportPayload := identity_services.ExportCertificatePayload{
		Format:     "pem",
		Passphrase: common.StringPtr("Test@Passphrase123"),
	}
	reqExport := client.CertificatesAPI.ExportCertificateByID(context.Background(), createdCertID).ExportCertificatePayload(exportPayload)
	exportRes, httpResExport, errExport := reqExport.Execute()

	if errExport != nil {
		handleAPIError(errExport)
	}

	// Verify the export operation was successful
	require.NoError(t, errExport, "Failed to export certificate")
	assert.Equal(t, 200, httpResExport.StatusCode, "Expected 200 OK status")
	require.NotNil(t, exportRes, "Export response should not be nil")
	require.NotNil(t, exportRes.Certificate, "Exported certificate data should not be nil")
	assert.Contains(t, *exportRes.Certificate, "-----BEGIN CERTIFICATE-----", "Exported certificate should be in PEM format")
	t.Logf("[SUCCESS] ExportCertificateByID exported certificate: %s", createdCertID)
}

// Test_identity_services_CertificatesAPIService_Import tests importing a certificate
func Test_identity_services_CertificatesAPIService_Import(t *testing.T) {
	client := SetupIdentitySvcTestClient(t)

	// Step 1: Create a certificate to export (so we have valid PEM content)
	certName := "test-cert-for-import-" + common.GenerateRandomString(6)
	rsaBits := float32(2048)
	algorithm := identity_services.CertificatesPostAlgorithm{
		RsaNumberOfBits: &rsaBits,
	}

	cert := identity_services.CertificatesPost{
		Algorithm:       algorithm,
		CertificateName: certName,
		CommonName:      certName + ".example.com",
		Digest:          "sha256",
		SignedBy:        "Root CA",
		Folder:          common.StringPtr("Prisma Access"),
	}

	fmt.Printf("Creating source certificate for import test: %s\n", certName)
	reqCreate := client.CertificatesAPI.CreateCertificates(context.Background()).CertificatesPost(cert)
	createRes, _, errCreate := reqCreate.Execute()

	if errCreate != nil {
		handleAPIError(errCreate)
	}
	require.NoError(t, errCreate, "Failed to create source certificate for import test")
	require.NotNil(t, createRes, "Create response should not be nil")
	sourceCertID := *createRes.Id
	t.Logf("Created source certificate with ID: %s", sourceCertID)

	// Cleanup source certificate at the end
	defer func() {
		t.Logf("Cleaning up source certificate with ID: %s", sourceCertID)
		reqDel := client.CertificatesAPI.DeleteCertificatesByID(context.Background(), sourceCertID)
		_, _ = reqDel.Execute()
	}()

	// Step 2: Export the certificate to get the PEM content
	exportPayload := identity_services.ExportCertificatePayload{
		Format:     "pem",
		Passphrase: common.StringPtr("Test@Passphrase123"),
	}
	reqExport := client.CertificatesAPI.ExportCertificateByID(context.Background(), sourceCertID).ExportCertificatePayload(exportPayload)
	exportRes, _, errExport := reqExport.Execute()

	if errExport != nil {
		handleAPIError(errExport)
	}
	require.NoError(t, errExport, "Failed to export source certificate")
	require.NotNil(t, exportRes, "Export response should not be nil")
	require.NotNil(t, exportRes.Certificate, "Exported certificate should not be nil")

	exportedPEM := *exportRes.Certificate
	t.Logf("Exported certificate PEM (length: %d bytes)", len(exportedPEM))

	// Step 3: Parse the exported PEM to separate certificate and key
	// The exported PEM contains both certificate and encrypted private key
	var certPEM, keyPEM string

	// Find and extract the certificate block
	certStart := strings.Index(exportedPEM, "-----BEGIN CERTIFICATE-----")
	certEnd := strings.Index(exportedPEM, "-----END CERTIFICATE-----")
	if certStart >= 0 && certEnd > certStart {
		certPEM = exportedPEM[certStart : certEnd+len("-----END CERTIFICATE-----")]
	}

	// Find and extract the encrypted private key block
	keyStart := strings.Index(exportedPEM, "-----BEGIN ENCRYPTED PRIVATE KEY-----")
	keyEnd := strings.Index(exportedPEM, "-----END ENCRYPTED PRIVATE KEY-----")
	if keyStart >= 0 && keyEnd > keyStart {
		keyPEM = exportedPEM[keyStart : keyEnd+len("-----END ENCRYPTED PRIVATE KEY-----")]
	}

	require.NotEmpty(t, certPEM, "Failed to extract certificate from exported PEM")
	require.NotEmpty(t, keyPEM, "Failed to extract private key from exported PEM")
	t.Logf("Extracted certificate (%d bytes) and key (%d bytes)", len(certPEM), len(keyPEM))

	// Base64 encode the certificate and key for import
	certBase64 := base64.StdEncoding.EncodeToString([]byte(certPEM))
	keyBase64 := base64.StdEncoding.EncodeToString([]byte(keyPEM))

	// Step 4: Import the certificate with a new name
	importCertName := "test-cert-imported-" + common.GenerateRandomString(6)

	importPayload := identity_services.CertificatesImport{
		Name:            importCertName,
		CertificateFile: certBase64,
		Format:          "pem",
		Folder:          common.StringPtr("Prisma Access"),
		KeyFile:         common.StringPtr(keyBase64),
		Passphrase:      common.StringPtr("Test@Passphrase123"),
	}

	fmt.Printf("Attempting to import certificate with name: %s\n", importCertName)
	reqImport := client.CertificatesAPI.ImportCertificates(context.Background()).CertificatesImport(importPayload)
	importRes, httpResImport, errImport := reqImport.Execute()

	if errImport != nil {
		handleAPIError(errImport)
	}

	// Verify the import operation was successful
	require.NoError(t, errImport, "Failed to import certificate")
	assert.Equal(t, 200, httpResImport.StatusCode, "Expected 200 OK status")
	require.NotNil(t, importRes, "Import response should not be nil")
	assert.Equal(t, importCertName, *importRes.Name, "Imported certificate name should match")
	t.Logf("[SUCCESS] ImportCertificates imported certificate: %s", importCertName)

	// Cleanup imported certificate
	defer func() {
		// Fetch the imported certificate to get its ID for deletion
		fetchedCert, errFetch := client.CertificatesAPI.FetchCertificates(
			context.Background(),
			importCertName,
			common.StringPtr("Prisma Access"),
			nil,
			nil,
		)
		if errFetch == nil && fetchedCert != nil {
			t.Logf("Cleaning up imported certificate with ID: %s", *fetchedCert.Id)
			reqDel := client.CertificatesAPI.DeleteCertificatesByID(context.Background(), *fetchedCert.Id)
			_, _ = reqDel.Execute()
		}
	}()
}
