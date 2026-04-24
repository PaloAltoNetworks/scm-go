/*
 * Mobile Agent Testing
 *
 * RegionalAndCustomProxiesAPIService
 */

package mobile_agent

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/mobile_agent"
)

// newTestRegionalAndCustomProxy returns a minimal valid ForwardingProfileRegionalAndCustomProxies.
// The constructor sets the default type "gp-and-pac" which is required by the server.
func newTestRegionalAndCustomProxy(name string) mobile_agent.ForwardingProfileRegionalAndCustomProxies {
	proxy1 := mobile_agent.NewForwardingProfileRegionalAndCustomProxiesProxy1()
	proxy1.SetFqdn("mail.gmail.com")
	proxy1.SetPort(80)
	proxy1.SetLocation("us")

	proxy2 := mobile_agent.NewForwardingProfileRegionalAndCustomProxiesProxy2()
	proxy2.SetFqdn("www.google.com")
	proxy2.SetPort(90)
	proxy2.SetLocation("us")

	proxy := mobile_agent.NewForwardingProfileRegionalAndCustomProxies(name)
	proxy.Proxy1 = proxy1
	proxy.Proxy2 = proxy2
	return *proxy
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_Create tests the creation of a regional and
// custom proxy.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_Create(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-regcustproxy-create-" + randomSuffix

	proxy := newTestRegionalAndCustomProxy(proxyName)

	fmt.Printf("Attempting to create Regional and Custom Proxy with name: %s\n", proxy.Name)

	req := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileRegionalAndCustomProxies(proxy)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Regional and Custom Proxy")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created Regional and Custom Proxy should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up Regional and Custom Proxy with ID: %s", createdID)
		_, errDel := client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete Regional and Custom Proxy during cleanup")
	}()

	assert.Equal(t, proxyName, res.Name, "Created proxy name should match")
	assert.Equal(t, "gp-and-pac", res.GetType(), "Type should match")

	// Validate proxy_1
	require.NotNil(t, res.Proxy1, "Proxy1 should not be nil")
	assert.Equal(t, "mail.gmail.com", res.Proxy1.GetFqdn(), "Proxy1 FQDN should match")
	assert.Equal(t, "us", res.Proxy1.GetLocation(), "Proxy1 location should match")
	assert.Equal(t, int32(80), res.Proxy1.GetPort(), "Proxy1 port should match")

	// Validate proxy_2
	require.NotNil(t, res.Proxy2, "Proxy2 should not be nil")
	assert.Equal(t, "www.google.com", res.Proxy2.GetFqdn(), "Proxy2 FQDN should match")
	assert.Equal(t, "us", res.Proxy2.GetLocation(), "Proxy2 location should match")
	assert.Equal(t, int32(90), res.Proxy2.GetPort(), "Proxy2 port should match")

	t.Logf("Successfully created and validated Regional and Custom Proxy: %s with ID: %s", proxy.Name, createdID)
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_GetByID tests retrieving a regional and custom proxy by its ID.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_GetByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-regcustproxy-get-" + randomSuffix

	createRes, _, err := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileRegionalAndCustomProxies(newTestRegionalAndCustomProxy(proxyName)).
		Execute()
	require.NoError(t, err, "Failed to create Regional and Custom Proxy for get test")
	createdID := *createRes.Id

	defer func() {
		client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.RegionalAndCustomProxiesAPI.GetGlobalProtectRegionalAndCustomProxyByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get Regional and Custom Proxy by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, proxyName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_Update tests updating an existing regional and custom proxy.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_Update(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-regcustproxy-update-" + randomSuffix

	createRes, _, err := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileRegionalAndCustomProxies(newTestRegionalAndCustomProxy(proxyName)).
		Execute()
	require.NoError(t, err, "Failed to create Regional and Custom Proxy for update test")
	createdID := *createRes.Id

	defer func() {
		client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
	}()

	updatedProxy := newTestRegionalAndCustomProxy(proxyName)
	updatedProxy.Description = common.StringPtr("Updated description")

	updateRes, httpResUpdate, errUpdate := client.RegionalAndCustomProxiesAPI.UpdateGlobalProtectRegionalAndCustomProxyByID(context.Background(), createdID).
		ForwardingProfileRegionalAndCustomProxies(updatedProxy).
		Execute()
	require.NoError(t, errUpdate, "Failed to update Regional and Custom Proxy")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, proxyName, updateRes.Name, "Name should remain the same after update")
	assert.Equal(t, "Updated description", *updateRes.Description, "Description should be updated")
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_List tests listing regional and custom proxies.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_List(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-regcustproxy-list-" + randomSuffix

	createRes, _, err := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileRegionalAndCustomProxies(newTestRegionalAndCustomProxy(proxyName)).
		Execute()
	require.NoError(t, err, "Failed to create Regional and Custom Proxy for list test")
	createdID := *createRes.Id

	defer func() {
		client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.RegionalAndCustomProxiesAPI.ListGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("Mobile Users").
		Limit(10000).
		Execute()
	require.NoError(t, errList, "Failed to list Regional and Custom Proxies")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundProxy := false
	for _, p := range listRes.Data {
		if p.Name == proxyName {
			foundProxy = true
			break
		}
	}
	assert.True(t, foundProxy, "Created Regional and Custom Proxy should be found in the list")
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_DeleteByID tests deleting a regional and custom proxy by ID.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_DeleteByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-regcustproxy-delete-" + randomSuffix

	createRes, _, err := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileRegionalAndCustomProxies(newTestRegionalAndCustomProxy(proxyName)).
		Execute()
	require.NoError(t, err, "Failed to create Regional and Custom Proxy for delete test")
	createdID := *createRes.Id

	_, errDel := client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete Regional and Custom Proxy")
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_FetchRegionalAndCustomProxies tests the FetchRegionalAndCustomProxies convenience method.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_FetchRegionalAndCustomProxies(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-regcustproxy-fetch-" + randomSuffix

	testObj := newTestRegionalAndCustomProxy(testName)
	testObj.Description = common.StringPtr("Test regional and custom proxy for fetch")

	createReq := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("Mobile Users").
		ForwardingProfileRegionalAndCustomProxies(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.RegionalAndCustomProxiesAPI.FetchRegionalAndCustomProxies(
		context.Background(),
		testName,
		common.StringPtr("Mobile Users"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch regional_and_custom_proxies by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchRegionalAndCustomProxies found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.RegionalAndCustomProxiesAPI.FetchRegionalAndCustomProxies(
		context.Background(),
		"non-existent-regional-custom-proxy-xyz-12345",
		common.StringPtr("Mobile Users"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchRegionalAndCustomProxies correctly returned nil for non-existent object")
}

// newTestZtnaProxy returns a ForwardingProfileRegionalAndCustomProxies configured for ztna-agent type.
func newTestZtnaProxy(name string) mobile_agent.ForwardingProfileRegionalAndCustomProxies {
	connectivityPref := mobile_agent.NewForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner("masque")
	connectivityPref.SetEnabled(true)

	prismaLocation := mobile_agent.NewForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner("americas")
	prismaLocation.SetLocations([]string{"us-southwest"})

	proxy := mobile_agent.NewForwardingProfileRegionalAndCustomProxies(name)
	proxy.SetType("ztna-agent")
	proxy.SetConnectivityPreference([]mobile_agent.ForwardingProfileRegionalAndCustomProxiesConnectivityPreferenceInner{*connectivityPref})
	proxy.SetFallbackOption("fail-open")
	proxy.SetLocationPreference("specific-pa-location")
	proxy.SetPrismaAccessLocations([]mobile_agent.ForwardingProfileRegionalAndCustomProxiesPrismaAccessLocationsInner{*prismaLocation})
	return *proxy
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_Create tests creation of a ztna-agent proxy.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_Create(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-ztna-create-" + randomSuffix

	proxy := newTestZtnaProxy(proxyName)

	fmt.Printf("Attempting to create ZTNA Agent Proxy with name: %s\n", proxy.Name)

	req := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("All").
		ForwardingProfileRegionalAndCustomProxies(proxy)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create ZTNA Agent Proxy")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created ZTNA Agent Proxy should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up ZTNA Agent Proxy with ID: %s", createdID)
		_, errDel := client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete ZTNA Agent Proxy during cleanup")
	}()

	assert.Equal(t, proxyName, res.Name, "Created proxy name should match")
	assert.Equal(t, "ztna-agent", res.GetType(), "Type should match")
	assert.Equal(t, "fail-open", res.GetFallbackOption(), "FallbackOption should match")
	assert.Equal(t, "specific-pa-location", res.GetLocationPreference(), "LocationPreference should match")

	require.Len(t, res.ConnectivityPreference, 1, "ConnectivityPreference should have 1 entry")
	assert.Equal(t, "masque", res.ConnectivityPreference[0].Name, "ConnectivityPreference name should match")
	assert.True(t, res.ConnectivityPreference[0].GetEnabled(), "ConnectivityPreference should be enabled")

	require.Len(t, res.PrismaAccessLocations, 1, "PrismaAccessLocations should have 1 entry")
	assert.Equal(t, "americas", res.PrismaAccessLocations[0].Name, "PrismaAccessLocations name should match")
	assert.Equal(t, []string{"us-southwest"}, res.PrismaAccessLocations[0].GetLocations(), "PrismaAccessLocations locations should match")

	t.Logf("Successfully created and validated ZTNA Agent Proxy: %s with ID: %s", proxy.Name, createdID)
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_GetByID tests retrieving a ztna-agent proxy by ID.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_GetByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-ztna-get-" + randomSuffix

	createRes, _, err := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("All").
		ForwardingProfileRegionalAndCustomProxies(newTestZtnaProxy(proxyName)).
		Execute()
	require.NoError(t, err, "Failed to create ZTNA Agent Proxy for get test")
	createdID := *createRes.Id

	defer func() {
		client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.RegionalAndCustomProxiesAPI.GetGlobalProtectRegionalAndCustomProxyByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get ZTNA Agent Proxy by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, proxyName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_Update tests updating an existing ztna-agent proxy.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_Update(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-ztna-update-" + randomSuffix

	createRes, _, err := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("All").
		ForwardingProfileRegionalAndCustomProxies(newTestZtnaProxy(proxyName)).
		Execute()
	require.NoError(t, err, "Failed to create ZTNA Agent Proxy for update test")
	createdID := *createRes.Id

	defer func() {
		client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
	}()

	updatedProxy := newTestZtnaProxy(proxyName)
	updatedProxy.Description = common.StringPtr("Updated description")

	updateRes, httpResUpdate, errUpdate := client.RegionalAndCustomProxiesAPI.UpdateGlobalProtectRegionalAndCustomProxyByID(context.Background(), createdID).
		ForwardingProfileRegionalAndCustomProxies(updatedProxy).
		Execute()
	require.NoError(t, errUpdate, "Failed to update ZTNA Agent Proxy")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, proxyName, updateRes.Name, "Name should remain the same after update")
	assert.Equal(t, "Updated description", *updateRes.Description, "Description should be updated")
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_List tests listing ztna-agent proxies.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_List(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-ztna-list-" + randomSuffix

	createRes, _, err := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("All").
		ForwardingProfileRegionalAndCustomProxies(newTestZtnaProxy(proxyName)).
		Execute()
	require.NoError(t, err, "Failed to create ZTNA Agent Proxy for list test")
	createdID := *createRes.Id

	defer func() {
		client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.RegionalAndCustomProxiesAPI.ListGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("All").
		Limit(10000).
		Execute()
	require.NoError(t, errList, "Failed to list ZTNA Agent Proxies")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundProxy := false
	for _, p := range listRes.Data {
		if p.Name == proxyName {
			foundProxy = true
			break
		}
	}
	assert.True(t, foundProxy, "Created ZTNA Agent Proxy should be found in the list")
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_DeleteByID tests deleting a ztna-agent proxy by ID.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_DeleteByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	proxyName := "test-ztna-delete-" + randomSuffix

	createRes, _, err := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("All").
		ForwardingProfileRegionalAndCustomProxies(newTestZtnaProxy(proxyName)).
		Execute()
	require.NoError(t, err, "Failed to create ZTNA Agent Proxy for delete test")
	createdID := *createRes.Id

	_, errDel := client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete ZTNA Agent Proxy")
}

// Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_Fetch tests the FetchRegionalAndCustomProxies convenience method for ztna-agent proxies.
func Test_mobile_agent_RegionalAndCustomProxiesAPIService_ZtnaAgent_Fetch(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-ztna-fetch-" + randomSuffix

	testObj := newTestZtnaProxy(testName)
	testObj.Description = common.StringPtr("Test ZTNA agent proxy for fetch")

	createReq := client.RegionalAndCustomProxiesAPI.CreateGlobalProtectRegionalAndCustomProxies(context.Background()).
		Folder("All").
		ForwardingProfileRegionalAndCustomProxies(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		client.RegionalAndCustomProxiesAPI.DeleteGlobalProtectRegionalAndCustomProxies(context.Background(), createdID).Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.RegionalAndCustomProxiesAPI.FetchRegionalAndCustomProxies(
		context.Background(),
		testName,
		common.StringPtr("All"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch ztna-agent proxy by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchRegionalAndCustomProxies found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.RegionalAndCustomProxiesAPI.FetchRegionalAndCustomProxies(
		context.Background(),
		"non-existent-ztna-proxy-xyz-12345",
		common.StringPtr("All"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchRegionalAndCustomProxies correctly returned nil for non-existent object")
}
