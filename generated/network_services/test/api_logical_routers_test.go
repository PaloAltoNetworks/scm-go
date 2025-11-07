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
// using the specific resource names from the updated Terraform example.
func createLogicalRouterVrfInner(t *testing.T, interfaceName string, nextHopVarName string) network_services.LogicalRoutersVrfInner {
	// 1. Define the Static Route components using the correct, specific model types.

	// Define the type for Nexthop (assuming it must be defined elsewhere, but using inferred name)
	// NOTE: If the client did not generate 'LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerNexthop',
	// you must replace this type with the actual generated Nexthop struct.
	type NexthopType = network_services.LogicalRoutersVrfInnerRoutingTableIpStaticRouteInnerNexthop

	// Route 1: default-route (Literal IP address nexthop)
	route1Name := "default-route"
	route1 := *network_services.NewLogicalRoutersVrfInnerRoutingTableIpStaticRouteInner(route1Name)
	route1.SetDestination("0.0.0.0/0")
	route1.SetAdminDist(10) // Assuming Preference maps to AdminDist or Metric based on standard routing practice
	route1.SetNexthop(NexthopType{
		IpAddress: common.StringPtr("198.18.1.1"),
	})

	// Route 2: internal-route (Variable IP address nexthop)
	route2Name := "internal-route"
	route2 := *network_services.NewLogicalRoutersVrfInnerRoutingTableIpStaticRouteInner(route2Name)
	route2.SetDestination("192.168.1.0/24")
	route2.SetInterface(interfaceName)
	route2.SetAdminDist(11)
	route2.SetNexthop(NexthopType{
		IpAddress: common.StringPtr(nextHopVarName), // Use the variable name
	})

	// Route 3: route-with-fqdn-nh (Variable FQDN nexthop)
	route3Name := "route-with-fqdn-nh"
	route3 := *network_services.NewLogicalRoutersVrfInnerRoutingTableIpStaticRouteInner(route3Name)
	route3.SetDestination("192.168.2.0/24")
	route3.SetInterface(interfaceName)
	route3.SetAdminDist(12)
	route3.SetNexthop(NexthopType{
		Fqdn: common.StringPtr(nextHopVarName), // Use the FQDN variable
	})
	// If the client generated a Bfd struct:
	// route3.SetBfd(network_services.LogicalRoutersVrfInnerBgpGlobalBfd{})

	// 2. Define Routing Table and IP configuration
	routingTableIp := network_services.NewLogicalRoutersVrfInnerRoutingTableIp()

	// SetStaticRoute expects []LogicalRoutersVrfInnerRoutingTableIpStaticRouteInner
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
func createTestLogicalRouter(t *testing.T, routerName string) network_services.LogicalRouters {
	// UPDATED CONSTANTS based on your new resource definitions
	const (
		EthInterfaceName = "$scm_ethernet_interface_test1"
		NextHopVarName   = "$scm_next_hop_test"
		TargetFolder     = "All"
	)

	// Build the nested VRF structure
	vrfInner := createLogicalRouterVrfInner(t, EthInterfaceName, NextHopVarName)

	// Build the main Logical Router
	logicalRouter := network_services.NewLogicalRouters(routerName)

	// Set required / essential fields
	logicalRouter.SetFolder(TargetFolder)
	logicalRouter.SetRoutingStack("advanced")
	logicalRouter.SetVrf([]network_services.LogicalRoutersVrfInner{vrfInner})

	return *logicalRouter
}

// ---------------------------------------------------------------------------------------------------------------------

// Test_network_services_LogicalRoutersAPIService_Create tests the creation of a Logical Router.
func Test_network_services_LogicalRoutersAPIService_Create(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)
	// Use the name from the example, and let the UUID be unique
	routerName := "scm_logical_router_test123"

	// To ensure the test is isolated, append a random string to the name.
	uniqueRouterName := generateLogicalRouterName(routerName + "-")

	logicalRouter := createTestLogicalRouter(t, uniqueRouterName)

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
	routerName := "scm_logical_router_test"

	// To ensure the test is isolated, append a random string to the name.
	uniqueRouterName := generateLogicalRouterName(routerName + "-")

	logicalRouter := createTestLogicalRouter(t, uniqueRouterName)
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
	routerName := "scm_logical_router_test"

	// To ensure the test is isolated, append a random string to the name.
	uniqueRouterName := generateLogicalRouterName(routerName + "-")

	logicalRouter := createTestLogicalRouter(t, uniqueRouterName)

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

	// Test: List the routers, filtering by name and folder
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
	routerName := "scm_lr_root_update_test"
	uniqueRouterName := generateLogicalRouterName(routerName + "-")
	targetFolder := "All"

	// --- 1. SETUP: Create the resource ---
	initialRouter := createTestLogicalRouter(t, uniqueRouterName)
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

	updatedRouter := createTestLogicalRouter(t, uniqueRouterName)

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
