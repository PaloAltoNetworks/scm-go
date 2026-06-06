/*
 * ZTNA FQDN Testing — FQDNAPIService
 * Covers: Create, GetByID, Update, List, Delete, FetchByName
 */
package ztna_connector_all

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/ztna_connector_all"
)

// fqdnConnectorGroupID reads ZTNA_CONNECTOR_GROUP_ID and skips the test if unset.
func fqdnConnectorGroupID(t *testing.T) string {
	t.Helper()
	return connectorGroupID(t)
}

const testFQDN = "test.example.com"

// Test_ztna_connector_all_FQDNRulesAPIService_Create tests creating a FQDN rule.
func Test_ztna_connector_all_FQDNRulesAPIService_Create(t *testing.T) {
	groupID := fqdnConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-fqdn-create-%s", common.GenerateRandomString(6))

	createTestApplication(t, client, name, groupID, testFQDN)
	oid := fetchTestApplicationOID(t, client, name)
	assert.NotEmpty(t, oid, "OID should be set after creation")
	t.Logf("Successfully created FQDN rule: %s with OID: %s", name, oid)

	t.Cleanup(func() { deleteTestApplication(t, client, oid, name) })
}

// Test_ztna_connector_all_FQDNRulesAPIService_GetByID creates a FQDN rule and retrieves it by OID.
func Test_ztna_connector_all_FQDNRulesAPIService_GetByID(t *testing.T) {
	groupID := fqdnConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-fqdn-getbyid-%s", common.GenerateRandomString(6))

	createTestApplication(t, client, name, groupID, testFQDN)
	oid := fetchTestApplicationOID(t, client, name)
	t.Cleanup(func() { deleteTestApplication(t, client, oid, name) })

	app, httpRes, err := client.FQDNAPI.GetApplicationsByID(context.Background(), oid).
		XPanwRegion("americas").
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to get FQDN rule by ID")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	require.NotNil(t, app)
	assert.Equal(t, name, app.Name)
	assert.Equal(t, groupID, app.Group)
	assert.Equal(t, oid, app.GetOid())
	fmt.Printf("Name:  %s\nGroup: %s\nOID:   %s\n", app.Name, app.Group, app.GetOid())
	for i, s := range app.Spec {
		fmt.Printf("Spec[%d]: fqdn=%s\n", i, s.Fqdn)
	}
	t.Logf("GetByID OK: name=%s oid=%s", app.Name, app.GetOid())
}

// Test_ztna_connector_all_FQDNRulesAPIService_Update creates a FQDN rule and updates its spec.
func Test_ztna_connector_all_FQDNRulesAPIService_Update(t *testing.T) {
	groupID := fqdnConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-fqdn-update-%s", common.GenerateRandomString(6))

	createTestApplication(t, client, name, groupID, testFQDN)
	oid := fetchTestApplicationOID(t, client, name)
	t.Cleanup(func() { deleteTestApplication(t, client, oid, name) })

	// Fetch current state so immutable fields are preserved.
	current, _, err := client.FQDNAPI.GetApplicationsByID(context.Background(), oid).
		XPanwRegion("americas").Execute()
	require.NoError(t, err, "Failed to fetch current FQDN rule before update")

	// Rebuild spec with the same FQDN but updated tcp_port.
	spec := ztna_connector_all.NewApplicationsSpecInner(current.Spec[0].Fqdn)
	spec.SetTcpPort("8080")

	updated := ztna_connector_all.NewApplications(current.Group, current.Name,
		[]ztna_connector_all.ApplicationsSpecInner{*spec})
	updated.SetDescription(current.GetDescription())
	updated.SetAppEnabled(current.GetAppEnabled())
	updated.SetIcmpAllowed(current.GetIcmpAllowed())

	httpRes, err := client.FQDNAPI.UpdateApplicationsByID(context.Background(), oid).
		Applications(*updated).
		XPanwRegion("americas").
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to update FQDN rule")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK for update")
	t.Logf("Successfully updated tcp_port to 8080 for FQDN rule: %s (OID: %s)", name, oid)
}

// Test_ztna_connector_all_FQDNRulesAPIService_List lists FQDN rules and verifies the created one appears.
func Test_ztna_connector_all_FQDNRulesAPIService_List(t *testing.T) {
	groupID := fqdnConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-fqdn-list-%s", common.GenerateRandomString(6))

	createTestApplication(t, client, name, groupID, testFQDN)
	oid := fetchTestApplicationOID(t, client, name)
	t.Cleanup(func() { deleteTestApplication(t, client, oid, name) })

	listRes, httpRes, err := client.FQDNAPI.ListApplications(context.Background()).
		XPanwRegion("americas").
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list FQDN rules")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK")
	require.NotNil(t, listRes)
	fmt.Printf("List returned %d FQDN rule(s)\n", listRes.Total)

	found := false
	for _, item := range listRes.Data {
		fmt.Printf("  - %s (OID: %s)\n", item.Name, item.GetOid())
		if item.Name == name {
			found = true
			assert.Equal(t, oid, item.GetOid())
		}
	}
	assert.True(t, found, "Created FQDN rule should appear in list")
}

// Test_ztna_connector_all_FQDNRulesAPIService_Delete creates a FQDN rule, deletes it, then verifies 404.
func Test_ztna_connector_all_FQDNRulesAPIService_Delete(t *testing.T) {
	groupID := fqdnConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)
	name := fmt.Sprintf("test-fqdn-delete-%s", common.GenerateRandomString(6))

	createTestApplication(t, client, name, groupID, testFQDN)
	oid := fetchTestApplicationOID(t, client, name)

	httpRes, err := client.FQDNAPI.DeleteApplicationsByID(context.Background(), oid).
		XPanwRegion("americas").
		Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to delete FQDN rule")
	assert.Equal(t, 202, httpRes.StatusCode, "Expected 202 Accepted for delete")

	_, getRes, getErr := client.FQDNAPI.GetApplicationsByID(context.Background(), oid).
		XPanwRegion("americas").
		Execute()
	assert.Error(t, getErr, "Should error when fetching deleted rule")
	assert.Equal(t, 404, getRes.StatusCode, "Should get 404 for deleted rule")
	t.Logf("Successfully deleted FQDN rule (OID: %s)", oid)
}

// Test_ztna_connector_all_FQDNRulesAPIService_FetchByName tests the FetchFQDN convenience method.
func Test_ztna_connector_all_FQDNRulesAPIService_FetchByName(t *testing.T) {
	groupID := fqdnConnectorGroupID(t)
	client := SetupZtnaConnectorAllTestClient(t)

	name := fmt.Sprintf("test-fqdn-fetch-%s", common.GenerateRandomString(8))
	createTestApplication(t, client, name, groupID, testFQDN)
	oid := fetchTestApplicationOID(t, client, name)
	t.Cleanup(func() { deleteTestApplication(t, client, oid, name) })

	// Test: fetch existing rule by name.
	app, err := client.FQDNAPI.FetchFQDN(context.Background(), name, nil, nil, nil)
	require.NoError(t, err, "FetchFQDN should not error")
	require.NotNil(t, app, "FetchFQDN should find the rule")
	assert.Equal(t, name, app.Name)
	fmt.Printf("FetchFQDN found rule: %s with OID: %s\n", app.Name, app.GetOid())

	// Test: fetch non-existent rule returns nil, nil.
	missing, err := client.FQDNAPI.FetchFQDN(context.Background(), "nonexistent-fqdn-rule-xyz-12345", nil, nil, nil)
	require.NoError(t, err, "FetchFQDN should not error for missing rule")
	assert.Nil(t, missing, "Should return nil for non-existent rule")
}
