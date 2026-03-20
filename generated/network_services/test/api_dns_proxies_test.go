/*
 * Network Services Testing
 *
 * DNSProxiesAPIService
 */

package network_services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// Test_network_services_DNSProxiesAPIService_Create tests the creation of a DNS Proxy.
func Test_network_services_DNSProxiesAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-dnsproxy-create-" + randomSuffix

	dnsProxy := network_services.DnsProxies{
		Name: proxyName,
		Default: network_services.DnsProxiesDefault{
			Primary:   "8.8.8.8",
			Secondary: common.StringPtr("8.8.4.4"),
		},
		Folder:  common.StringPtr("All"),
		Enabled: common.BoolPtr(true),
	}

	fmt.Printf("Attempting to create DNS Proxy with name: %s\n", dnsProxy.Name)

	req := client.DNSProxiesAPI.CreateDNSProxies(context.Background()).DnsProxies(dnsProxy)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create DNS Proxy")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created DNS proxy should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up DNS Proxy with ID: %s", createdID)
		_, errDel := client.DNSProxiesAPI.DeleteDNSProxiesByID(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete DNS proxy during cleanup")
	}()

	assert.Equal(t, proxyName, res.Name, "Created DNS proxy name should match")
	assert.Equal(t, "8.8.8.8", res.Default.Primary, "Primary DNS should match")
	t.Logf("Successfully created and validated DNS Proxy: %s with ID: %s", dnsProxy.Name, createdID)
}

// Test_network_services_DNSProxiesAPIService_GetByID tests retrieving a DNS proxy by its ID.
func Test_network_services_DNSProxiesAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-dnsproxy-get-" + randomSuffix

	dnsProxy := network_services.DnsProxies{
		Name: proxyName,
		Default: network_services.DnsProxiesDefault{
			Primary: "1.1.1.1",
		},
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.DNSProxiesAPI.CreateDNSProxies(context.Background()).DnsProxies(dnsProxy).Execute()
	require.NoError(t, err, "Failed to create DNS proxy for get test")
	createdID := *createRes.Id

	defer func() {
		client.DNSProxiesAPI.DeleteDNSProxiesByID(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.DNSProxiesAPI.GetDNSProxiesByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get DNS proxy by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, proxyName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_network_services_DNSProxiesAPIService_Update tests updating an existing DNS proxy.
func Test_network_services_DNSProxiesAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-dnsproxy-update-" + randomSuffix

	dnsProxy := network_services.DnsProxies{
		Name: proxyName,
		Default: network_services.DnsProxiesDefault{
			Primary: "8.8.8.8",
		},
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.DNSProxiesAPI.CreateDNSProxies(context.Background()).DnsProxies(dnsProxy).Execute()
	require.NoError(t, err, "Failed to create DNS proxy for update test")
	createdID := *createRes.Id

	defer func() {
		client.DNSProxiesAPI.DeleteDNSProxiesByID(context.Background(), createdID).Execute()
	}()

	updatedDNSProxy := network_services.DnsProxies{
		Name: proxyName,
		Default: network_services.DnsProxiesDefault{
			Primary:   "1.1.1.1",
			Secondary: common.StringPtr("1.0.0.1"),
		},
		Folder:  common.StringPtr("All"),
		Enabled: common.BoolPtr(false),
	}

	updateRes, httpResUpdate, errUpdate := client.DNSProxiesAPI.UpdateDNSProxiesByID(context.Background(), createdID).DnsProxies(updatedDNSProxy).Execute()
	require.NoError(t, errUpdate, "Failed to update DNS proxy")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, "1.1.1.1", updateRes.Default.Primary, "Primary DNS should be updated")
	assert.Equal(t, false, *updateRes.Enabled, "Enabled should be updated")
}

// Test_network_services_DNSProxiesAPIService_List tests listing DNS Proxies.
func Test_network_services_DNSProxiesAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-dnsproxy-list-" + randomSuffix

	dnsProxy := network_services.DnsProxies{
		Name: proxyName,
		Default: network_services.DnsProxiesDefault{
			Primary: "8.8.8.8",
		},
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.DNSProxiesAPI.CreateDNSProxies(context.Background()).DnsProxies(dnsProxy).Execute()
	require.NoError(t, err, "Failed to create DNS proxy for list test")
	createdID := *createRes.Id

	defer func() {
		client.DNSProxiesAPI.DeleteDNSProxiesByID(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.DNSProxiesAPI.ListDNSProxies(context.Background()).Folder("All").Limit(10000).Execute()
	require.NoError(t, errList, "Failed to list DNS proxies")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundProxy := false
	for _, proxy := range listRes.Data {
		if proxy.Name == proxyName {
			foundProxy = true
			break
		}
	}
	assert.True(t, foundProxy, "Created DNS proxy should be found in the list")
}

// Test_network_services_DNSProxiesAPIService_DeleteByID tests deleting a DNS proxy by ID.
func Test_network_services_DNSProxiesAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-dnsproxy-delete-" + randomSuffix

	dnsProxy := network_services.DnsProxies{
		Name: proxyName,
		Default: network_services.DnsProxiesDefault{
			Primary: "8.8.8.8",
		},
		Folder: common.StringPtr("All"),
	}

	createRes, _, err := client.DNSProxiesAPI.CreateDNSProxies(context.Background()).DnsProxies(dnsProxy).Execute()
	require.NoError(t, err, "Failed to create DNS proxy for delete test")
	createdID := *createRes.Id

	_, errDel := client.DNSProxiesAPI.DeleteDNSProxiesByID(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete DNS proxy")
}

// Test_network_services_DNSProxiesAPIService_FetchDNSProxies tests the FetchDNSProxies convenience method
func Test_network_services_DNSProxiesAPIService_FetchDNSProxies(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-dnsproxy-fetch-" + randomSuffix

	testObj := network_services.DnsProxies{
		Name: testName,
		Default: network_services.DnsProxiesDefault{
			Primary: "8.8.8.8",
		},
		Folder: common.StringPtr("All"),
	}

	createReq := client.DNSProxiesAPI.CreateDNSProxies(context.Background()).DnsProxies(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		deleteReq := client.DNSProxiesAPI.DeleteDNSProxiesByID(context.Background(), createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.DNSProxiesAPI.FetchDNSProxies(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch dns_proxies by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchDNSProxies found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.DNSProxiesAPI.FetchDNSProxies(
		context.Background(),
		"non-existent-dns_proxies-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchDNSProxies correctly returned nil for non-existent object")
}
