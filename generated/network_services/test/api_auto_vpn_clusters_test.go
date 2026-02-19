/*
 * Network Services Testing
 *
 * AutoVPNClustersAPIService
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

// Test_networkservices_AutoVPNClustersAPIService_Create tests the creation of an Auto VPN Cluster.
func Test_networkservices_AutoVPNClustersAPIService_Create(t *testing.T) {
	t.Skip("Auto VPN Clusters require special infrastructure (gateways, branches) that may not be available in test environment")

	// Setup the authenticated client.
	client := SetupNetworkSvcTestClient(t)

	// Create a valid Auto VPN Cluster object with a unique name.
	randomSuffix := common.GenerateRandomString(6)
	clusterName := "auto_vpn_cluster_create_" + randomSuffix

	cluster := network_services.AutoVpnClusters{
		Name: common.StringPtr(clusterName),
		Type: common.StringPtr("hub-spoke"),
	}

	fmt.Printf("Attempting to create Auto VPN Cluster with name: %s\n", *cluster.Name)

	// Make the create request to the API.
	req := client.AutoVPNClustersAPI.CreateAutoVPNClusters(context.Background()).AutoVpnClusters(cluster)
	res, httpRes, err := req.Execute()

	// Defer cleanup for the Auto VPN Cluster.
	if res != nil && res.Id != nil {
		defer func() {
			t.Logf("Cleaning up Auto VPN Cluster with ID: %s", *res.Id)
			delReq := client.AutoVPNClustersAPI.DeleteAutoVPNClustersByID(context.Background(), *res.Id)
			_, errDel := delReq.Execute()
			if errDel != nil {
				t.Logf("Failed to delete Auto VPN Cluster during cleanup: %v", errDel)
			}
		}()
	}

	// Verify the request was successful.
	handleAPIError(err)
	require.NoError(t, err, "Create request should not return an error")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "The response from create should not be nil")
	assert.Equal(t, clusterName, *res.Name, "The name of the created cluster should match")
	assert.NotEmpty(t, *res.Id, "The ID of the created cluster should not be empty")

	t.Logf("Successfully created Auto VPN Cluster with ID: %s", *res.Id)
}

// Test_networkservices_AutoVPNClustersAPIService_GetByID tests the retrieval of an Auto VPN Cluster by its ID.
func Test_networkservices_AutoVPNClustersAPIService_GetByID(t *testing.T) {
	t.Skip("Auto VPN Clusters require special infrastructure (gateways, branches) that may not be available in test environment")

	client := SetupNetworkSvcTestClient(t)

	// Create a cluster to retrieve.
	randomSuffix := common.GenerateRandomString(6)
	clusterName := "auto_vpn_cluster_get_" + randomSuffix

	cluster := network_services.AutoVpnClusters{
		Name: common.StringPtr(clusterName),
		Type: common.StringPtr("hub-spoke"),
	}

	createRes, _, err := client.AutoVPNClustersAPI.CreateAutoVPNClusters(context.Background()).AutoVpnClusters(cluster).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create cluster for get test")
	createdClusterID := *createRes.Id

	// Defer cleanup for the Auto VPN Cluster.
	defer func() {
		t.Logf("Cleaning up Auto VPN Cluster with ID: %s", createdClusterID)
		_, errDel := client.AutoVPNClustersAPI.DeleteAutoVPNClustersByID(context.Background(), createdClusterID).Execute()
		if errDel != nil {
			t.Logf("Failed to delete cluster during cleanup: %v", errDel)
		}
	}()

	t.Logf("Created Auto VPN Cluster for Get test with ID: %s", createdClusterID)

	// Test the Get by ID operation.
	fmt.Printf("Attempting to get Auto VPN Cluster with ID: %s\n", createdClusterID)
	req := client.AutoVPNClustersAPI.GetAutoVPNClustersByID(context.Background(), createdClusterID)
	getRes, httpRes, err := req.Execute()

	// Verify the retrieval was successful.
	handleAPIError(err)
	require.NoError(t, err, "Get by ID request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "The response from get should not be nil")
	assert.Equal(t, createdClusterID, *getRes.Id, "The ID of the retrieved cluster should match")
	assert.Equal(t, clusterName, *getRes.Name, "The name of the retrieved cluster should match")
}

// Test_networkservices_AutoVPNClustersAPIService_Update tests updating an Auto VPN Cluster.
func Test_networkservices_AutoVPNClustersAPIService_Update(t *testing.T) {
	t.Skip("Auto VPN Clusters require special infrastructure (gateways, branches) that may not be available in test environment")

	client := SetupNetworkSvcTestClient(t)

	// Create a cluster to update.
	randomSuffix := common.GenerateRandomString(6)
	clusterName := "auto_vpn_cluster_update_" + randomSuffix

	cluster := network_services.AutoVpnClusters{
		Name: common.StringPtr(clusterName),
		Type: common.StringPtr("hub-spoke"),
	}

	createRes, _, err := client.AutoVPNClustersAPI.CreateAutoVPNClusters(context.Background()).AutoVpnClusters(cluster).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create cluster for update test")
	createdClusterID := *createRes.Id

	// Defer cleanup for the Auto VPN Cluster.
	defer func() {
		t.Logf("Cleaning up Auto VPN Cluster with ID: %s", createdClusterID)
		_, errDel := client.AutoVPNClustersAPI.DeleteAutoVPNClustersByID(context.Background(), createdClusterID).Execute()
		if errDel != nil {
			t.Logf("Failed to delete cluster during cleanup: %v", errDel)
		}
	}()

	t.Logf("Created Auto VPN Cluster for Update test with ID: %s", createdClusterID)

	// Update the cluster object.
	updatedCluster := network_services.AutoVpnClusters{
		Name:                  common.StringPtr(clusterName),
		Type:                  common.StringPtr("hub-spoke"),
		EnableMeshBetweenHubs: common.BoolPtr(true),
	}

	// Test the Update by ID operation.
	fmt.Printf("Attempting to update Auto VPN Cluster with ID: %s\n", createdClusterID)
	reqUpdate := client.AutoVPNClustersAPI.UpdateAutoVPNClustersByID(context.Background(), createdClusterID).AutoVpnClusters(updatedCluster)
	updateRes, httpRes, err := reqUpdate.Execute()

	// Verify the update was successful.
	handleAPIError(err)
	require.NoError(t, err, "Update request should not return an error")
	assert.Equal(t, 200, httpRes.StatusCode, "Expected 200 OK status")
	require.NotNil(t, updateRes, "The response from update should not be nil")
}

// Test_networkservices_AutoVPNClustersAPIService_List tests listing Auto VPN Clusters.
func Test_networkservices_AutoVPNClustersAPIService_List(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Read-only test: list existing objects (no Create needed)
	listRes, httpResList, errList := client.AutoVPNClustersAPI.ListAutoVPNClusters(context.Background()).Limit(200).Offset(0).Execute()
	if errList != nil {
		handleAPIError(errList)
	}
	require.NoError(t, errList, "Failed to list auto VPN clusters")
	assert.Equal(t, 200, httpResList.StatusCode, "Expected 200 OK status")
	require.NotNil(t, listRes, "List response should not be nil")
	t.Logf("Successfully listed auto VPN clusters")
}

// Test_networkservices_AutoVPNClustersAPIService_Fetch tests the fetch convenience method.
func Test_networkservices_AutoVPNClustersAPIService_Fetch(t *testing.T) {
	client := SetupNetworkSvcTestClient(t)

	// Read-only test: Fetch non-existent object (should return nil, nil)
	notFound, err := client.AutoVPNClustersAPI.FetchAutoVPNClusters(
		context.Background(),
		"non-existent-auto-vpn-cluster-xyz-12345",
		nil,
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchAutoVPNClusters correctly returned nil for non-existent object")
}

// Test_networkservices_AutoVPNClustersAPIService_DeleteByID tests deleting an Auto VPN Cluster.
func Test_networkservices_AutoVPNClustersAPIService_DeleteByID(t *testing.T) {
	t.Skip("Auto VPN Clusters require special infrastructure (gateways, branches) that may not be available in test environment")

	client := SetupNetworkSvcTestClient(t)

	// Create a cluster to delete.
	randomSuffix := common.GenerateRandomString(6)
	clusterName := "auto_vpn_cluster_delete_" + randomSuffix

	cluster := network_services.AutoVpnClusters{
		Name: common.StringPtr(clusterName),
		Type: common.StringPtr("hub-spoke"),
	}

	createRes, _, err := client.AutoVPNClustersAPI.CreateAutoVPNClusters(context.Background()).AutoVpnClusters(cluster).Execute()
	handleAPIError(err)
	require.NoError(t, err, "Failed to create cluster for delete test")
	createdClusterID := *createRes.Id

	t.Logf("Created Auto VPN Cluster for Delete test with ID: %s", createdClusterID)

	// Test the Delete by ID operation.
	fmt.Printf("Attempting to delete Auto VPN Cluster with ID: %s\n", createdClusterID)
	reqDel := client.AutoVPNClustersAPI.DeleteAutoVPNClustersByID(context.Background(), createdClusterID)
	httpResDel, errDel := reqDel.Execute()

	// Verify the delete operation was successful.
	handleAPIError(errDel)
	require.NoError(t, errDel, "Failed to delete cluster")
	assert.Equal(t, 200, httpResDel.StatusCode, "Expected 200 OK status")
	t.Logf("Successfully deleted Auto VPN Cluster: %s", createdClusterID)
}
