package network_services

/*
 * Network Services Testing
 *
 * LogicalRoutersAPIService
 */

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// Assuming these are your common utils and generated client packages
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

// --- Helper Functions (Updated) ---

// generateLogicalRouterName creates a unique name for the Logical Router.
func generateLogicalRouterName(base string) string {
	return base + common.GenerateRandomString(4)
}

// createLogicalRouterVrfInner creates the nested LogicalRoutersVrfInner structure
// with static routes using literal IP addresses and FQDNs.
func createLogicalRouterVrfInner(t *testing.T, interfaceName string) network_services.LogicalRoutersVrfInner {
	// 1. Define the Static Route components using the correct, specific model types.

	type NexthopType = network_services.LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerNexthop

	// Route 1: default-route (Literal IP address nexthop)
	route1Name := "default-route"
	route1 := *network_services.NewLogicalRoutersVrfInnerRoutingTableIpStaticRouteInner(route1Name)
	route1.SetDestination("0.0.0.0/0")
	route1.SetAdminDist(10)
	route1.SetNexthop(NexthopType{
		IpAddress: common.StringPtr("198.18.1.1"),
	})

	// Route 2: internal-route (Literal IP address nexthop with interface)
	route2Name := "internal-route"
	route2 := *network_services.NewLogicalRoutersVrfInnerRoutingTableIpStaticRouteInner(route2Name)
	route2.SetDestination("192.168.1.0/24")
	route2.SetInterface(interfaceName)
	route2.SetAdminDist(11)
	route2.SetNexthop(NexthopType{
		IpAddress: common.StringPtr("192.0.2.1"), // Use a literal IP address
	})

	// Route 3: route-with-fqdn-nh (Literal FQDN nexthop)
	route3Name := "route-with-fqdn-nh"
	route3 := *network_services.NewLogicalRoutersVrfInnerRoutingTableIpStaticRouteInner(route3Name)
	route3.SetDestination("192.168.2.0/24")
	route3.SetInterface(interfaceName)
	route3.SetAdminDist(12)
	route3.SetNexthop(NexthopType{
		Fqdn: common.StringPtr("nexthop.example.com"), // Use a literal FQDN
	})

	// 2. Define Routing Table and IP configuration
	routingTableIp := network_services.NewLogicalRoutersVrfInnerRoutingTableIp()

	// SetStaticRoute expects []LogicalRoutersVrfInnerRoutingTableIpStaticRouteInner
	// This creates three test routes:
	// - Default route via 198.18.1.1
	// - Route to 192.168.1.0/24 via interface with nexthop 192.0.2.1
	// - Route to 192.168.2.0/24 via interface with FQDN nexthop
	staticRoutes := []network_services.LogicalRoutersVrfInnerRoutingTableIpStaticRouteInner{route1, route2, route3}
	routingTableIp.SetStaticRoute(staticRoutes)

	routingTable := network_services.NewLogicalRoutersVrfInnerRoutingTable()
	routingTable.SetIp(*routingTableIp)

	// 3. Define the main VRF
	vrfName := "default"
	vrf := network_services.NewLogicalRoutersVrfInner(vrfName)
	vrf.SetInterface([]string{interfaceName})
	vrf.SetRoutingTable(*routingTable)

	return *vrf
}

// createTestLogicalRouter creates a LogicalRouters object for testing.
// Takes an actual ethernet interface name (not a variable reference).
func createTestLogicalRouter(t *testing.T, routerName string, ethInterfaceName string) network_services.LogicalRouters {
	const TargetFolder = "All"

	// Build the nested VRF structure using the actual ethernet interface
	vrfInner := createLogicalRouterVrfInner(t, ethInterfaceName)

	// Build the main Logical Router
	logicalRouter := network_services.NewLogicalRouters(routerName)

	// Set required / essential fields
	logicalRouter.SetFolder(TargetFolder)
	logicalRouter.SetRoutingStack("advanced")
	logicalRouter.SetVrf([]network_services.LogicalRoutersVrfInner{vrfInner})

	return *logicalRouter
}

// createPrerequisiteEthernetInterface creates an ethernet interface for use in logical router tests.
func createPrerequisiteEthernetInterface(t *testing.T, client *network_services.APIClient, baseName string) (string, string) {
	// Generate unique name (ethernet interfaces require $ prefix)
	intfName := "$" + baseName + common.GenerateRandomString(4)

	// Create a simple Layer 3 ethernet interface
	intf := *network_services.NewEthernetInterfacesWithDefaults()
	intf.SetName(intfName)
	intf.SetComment("Prerequisite for Logical Router Test")
	intf.SetFolder("All")
	intf.SetLinkDuplex("auto")
	intf.SetLinkSpeed("auto")
	intf.SetLinkState("up")

	// Set Layer 3 mode with a basic IP
	layer3Config := network_services.NewEthernetInterfacesLayer3WithDefaults()
	layer3Config.SetIp([]network_services.EthernetInterfacesLayer3IpInner{
		{Name: "192.0.2.1/24"},
	})
	intf.SetLayer3(*layer3Config)

	res, _, err := client.EthernetInterfacesAPI.CreateEthernetInterfaces(context.Background()).
		EthernetInterfaces(intf).Execute()
	require.NoError(t, err, "Failed to create prerequisite ethernet interface")

	return res.GetId(), intfName
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_LogicalRoutersAPIService_Create tests the creation of a Logical Router.
func Test_network_services_LogicalRoutersAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Step 1: Create prerequisite ethernet interface
	ethIntfID, ethIntfName := createPrerequisiteEthernetInterface(t, client, "eth-lr-create-")
	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), ethIntfID).Execute()
	}()

	// Step 2: Create logical router using the actual interface
	routerName := "scm_logical_router_test123"
	uniqueRouterName := generateLogicalRouterName(routerName + "-")

	logicalRouter := createTestLogicalRouter(t, uniqueRouterName, ethIntfName)

	t.Logf("Creating Logical Router with name: %s in folder: %s", uniqueRouterName, logicalRouter.GetFolder())
	req := client.LogicalRoutersAPI.CreateLogicalRouters(context.Background()).LogicalRouters(logicalRouter)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create Logical Router")
	assert.Equal(t, http.StatusCreated, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response body should not be nil")

	createdID := *res.Id

	// Cleanup the created router
	defer func() {
		t.Logf("Cleaning up Logical Router with ID: %s", createdID)
		_, errDel := client.LogicalRoutersAPI.DeleteLogicalRoutersByID(context.Background(), createdID).Execute()
		// If cleanup fails, log the error but don't fail the primary test result
		if errDel != nil {
			t.Logf("Cleanup failed for ID %s: %v", createdID, errDel)
		}
	}()

	t.Logf("Successfully created Logical Router: %s with ID: %s", uniqueRouterName, createdID)

	// Verify the response matches key input fields
	assert.Equal(t, uniqueRouterName, res.Name, "Created router name should match")
	assert.Equal(t, "All", res.GetFolder(), "Folder should be 'All'")
	assert.Equal(t, "advanced", res.GetRoutingStack(), "Routing stack should match")
	require.Len(t, res.Vrf, 1, "Vrf list should contain one element")
}

// Test_network_services_LogicalRoutersAPIService_GetByID tests retrieving a Logical Router by ID.
func Test_network_services_LogicalRoutersAPIService_GetByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Step 1: Create prerequisite ethernet interface
	ethIntfID, ethIntfName := createPrerequisiteEthernetInterface(t, client, "eth-lr-get-")
	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), ethIntfID).Execute()
	}()

	// Step 2: Create logical router
	routerName := "scm_logical_router_test"
	uniqueRouterName := generateLogicalRouterName(routerName + "-")

	logicalRouter := createTestLogicalRouter(t, uniqueRouterName, ethIntfName)
	createRes, _, err := client.LogicalRoutersAPI.CreateLogicalRouters(context.Background()).LogicalRouters(logicalRouter).Execute()
	require.NoError(t, err, "Failed to create router for get test setup")
	createdID := *createRes.Id

	defer func() {
		// Cleanup the created router after the test runs
		client.LogicalRoutersAPI.DeleteLogicalRoutersByID(context.Background(), createdID).Execute()
	}()

	// Test: Retrieve the router
	getRes, httpRes, err := client.LogicalRoutersAPI.GetLogicalRoutersByID(context.Background(), createdID).Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to get Logical Router by ID")
	assert.Equal(t, http.StatusOK, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")

	// Verify the retrieved data matches the creation input
	assert.Equal(t, createdID, getRes.GetId(), "Retrieved ID should match the created ID")
	assert.Equal(t, uniqueRouterName, getRes.Name, "Retrieved name should match")
	assert.Equal(t, "advanced", getRes.GetRoutingStack(), "Routing stack should match")
}

// Test_network_services_LogicalRoutersAPIService_DeleteByID tests deleting a Logical Router.
func Test_network_services_LogicalRoutersAPIService_DeleteByID(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Step 1: Create prerequisite ethernet interface
	ethIntfID, ethIntfName := createPrerequisiteEthernetInterface(t, client, "eth-lr-del-")
	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), ethIntfID).Execute()
	}()

	// Step 2: Create logical router
	routerName := "scm_logical_router_test"
	uniqueRouterName := generateLogicalRouterName(routerName + "-")

	logicalRouter := createTestLogicalRouter(t, uniqueRouterName, ethIntfName)

	createRes, _, err := client.LogicalRoutersAPI.CreateLogicalRouters(context.Background()).LogicalRouters(logicalRouter).Execute()
	require.NoError(t, err, "Failed to create router for delete test setup")
	createdID := *createRes.Id

	// Test: Delete the router
	httpResDel, errDel := client.LogicalRoutersAPI.DeleteLogicalRoutersByID(context.Background(), createdID).Execute()

	if errDel != nil {
		handleAPIError(errDel)
	}
	require.NoError(t, errDel, "Failed to delete Logical Router")
	assert.Equal(t, http.StatusOK, httpResDel.StatusCode, "Expected 200 OK status for deletion")
}

// Test_network_services_LogicalRoutersAPIService_List tests listing Logical Routers.
func Test_network_services_LogicalRoutersAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	targetFolder := "All"

	// Step 1: Create prerequisite ethernet interface
	ethIntfID, ethIntfName := createPrerequisiteEthernetInterface(t, client, "eth-lr-list-")
	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), ethIntfID).Execute()
	}()

	// Step 2: Create a logical router to ensure list is not empty
	routerName := "scm_logical_router_list"
	uniqueRouterName := generateLogicalRouterName(routerName + "-")
	logicalRouter := createTestLogicalRouter(t, uniqueRouterName, ethIntfName)

	createRes, _, err := client.LogicalRoutersAPI.CreateLogicalRouters(context.Background()).LogicalRouters(logicalRouter).Execute()
	require.NoError(t, err, "Failed to create router for list test setup")
	createdID := *createRes.Id

	defer func() {
		client.LogicalRoutersAPI.DeleteLogicalRoutersByID(context.Background(), createdID).Execute()
	}()

	// Test: List the routers, filtering by folder
	listRes, httpRes, err := client.LogicalRoutersAPI.ListLogicalRouters(context.Background()).
		Folder(targetFolder).
		Limit(100).
		Execute()

	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to list Logical Routers")
	assert.Equal(t, http.StatusOK, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	require.GreaterOrEqual(t, len(listRes.Data), 1, "Expected at least one router in the filtered list")
}

// Test_network_services_LogicalRoutersAPIService_Update tests updating a Logical Router's root-level property.
func Test_network_services_LogicalRoutersAPIService_Update(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Step 1: Create prerequisite ethernet interface
	ethIntfID, ethIntfName := createPrerequisiteEthernetInterface(t, client, "eth-lr-upd-")
	defer func() {
		client.EthernetInterfacesAPI.DeleteEthernetInterfacesByID(context.Background(), ethIntfID).Execute()
	}()

	// Step 2: Create logical router
	routerName := "scm_lr_root_update_test"
	uniqueRouterName := generateLogicalRouterName(routerName + "-")
	targetFolder := "All"

	// --- 1. SETUP: Create the resource ---
	initialRouter := createTestLogicalRouter(t, uniqueRouterName, ethIntfName)
	initialRouter.SetFolder(targetFolder)

	originalVrf := initialRouter.GetVrf()

	createRes, _, err := client.LogicalRoutersAPI.CreateLogicalRouters(context.Background()).LogicalRouters(initialRouter).Execute()
	require.NoError(t, err, "Failed to create router for update test setup")
	createdID := *createRes.Id

	defer func() {
		// Cleanup the created resource
		client.LogicalRoutersAPI.DeleteLogicalRoutersByID(context.Background(), createdID).Execute()
	}()

	// --- 2. UPDATE: Prepare payload and execute PUT ---

	updatedRoutingStack := "legacy"

	updatedRouter := createTestLogicalRouter(t, uniqueRouterName, ethIntfName)

	// 2.1. Apply the ONLY change: Update the root-level field
	updatedRouter.SetRoutingStack(updatedRoutingStack)

	// 2.2. Critical: Set metadata required for the PUT request
	updatedRouter.SetId(createdID)
	updatedRouter.SetFolder(targetFolder)
	updatedRouter.SetVrf(originalVrf) // Must send back the original Vrf list intact

	// Execute the update
	updateRes, httpResUpdate, errUpdate := client.LogicalRoutersAPI.UpdateLogicalRoutersByID(context.Background(), createdID).
		LogicalRouters(updatedRouter).
		Execute()

	// --- 3. VERIFICATION ---

	require.NoError(t, errUpdate, "Failed to update Logical Router")
	assert.Equal(t, http.StatusOK, httpResUpdate.StatusCode, "Expected 200 OK status")
	assert.NotNil(t, updateRes, "Update response should not be nil")

	// Verify updated root field
	assert.Equal(t, updatedRoutingStack, updateRes.GetRoutingStack(), "Routing Stack should be updated to 'legacy'")

	// Verify non-updated identifying and nested fields are preserved
	assert.Equal(t, uniqueRouterName, updateRes.Name, "Name must remain unchanged")
	require.Equal(t, len(originalVrf), len(updateRes.Vrf), "VRF list size must be preserved")
}

// Test_network_services_LogicalRoutersAPIService_FetchLogicalRouters tests the FetchLogicalRouters convenience method
func Test_network_services_LogicalRoutersAPIService_FetchLogicalRouters(t *testing.T) {
	// Setup the authenticated client
	client := SetupNetworkSvcTestClient(t)

	// Create a test object first (inline creation like other tests)
	testName := "fetch-lr-" + common.GenerateRandomString(6)
	testObj := network_services.LogicalRouters{
		Name:   testName,
		Folder: common.StringPtr("Prisma Access"),
	}

	createReq := client.LogicalRoutersAPI.CreateLogicalRouters(context.Background()).LogicalRouters(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := createRes.Id

	// Cleanup after test
	defer func() {
		deleteReq := client.LogicalRoutersAPI.DeleteLogicalRoutersByID(context.Background(), *createdID)
		_, _ = deleteReq.Execute()
		t.Logf("Cleaned up test object: %s", *createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.LogicalRoutersAPI.FetchLogicalRouters(
		context.Background(),
		testName,
		common.StringPtr("Prisma Access"),
		nil, // snippet
		nil, // device
	)

	// Verify successful fetch
	require.NoError(t, err, "Failed to fetch logical_routers by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchLogicalRouters found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.LogicalRoutersAPI.FetchLogicalRouters(
		context.Background(),
		"non-existent-logical_routers-xyz-12345",
		common.StringPtr("Prisma Access"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchLogicalRouters correctly returned nil for non-existent object")
}
