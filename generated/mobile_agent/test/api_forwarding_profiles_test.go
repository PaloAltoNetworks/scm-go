/*
 * Mobile Agent Testing
 *
 * ForwardingProfilesAPIService
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

// newTestForwardingProfile returns a minimal valid ForwardingProfiles. The API
// requires the 'type' node even though the OpenAPI schema marks only 'name' as
// required. GlobalProtectProxy has all-optional sub-fields so an empty struct
// satisfies the server.
func newTestForwardingProfile(name string) mobile_agent.ForwardingProfiles {
	return mobile_agent.ForwardingProfiles{
		Name: name,
		Type: &mobile_agent.ForwardingProfilesType{
			GlobalProtectProxy: mobile_agent.NewForwardingProfileGlobalProtectProxyGlobalProtectProxy(),
		},
	}
}

// Test_mobile_agent_ForwardingProfilesAPIService_Create tests the creation of a forwarding profile
// using a ztna_agent type with a forwarding rule and block rule fully populated.
func Test_mobile_agent_ForwardingProfilesAPIService_Create(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	profileName := "test-fwdprofile-create-" + randomSuffix

	forwardingRule := mobile_agent.NewForwardingRuleZtna("rule-1")
	forwardingRule.SetTrafficType("dns")
	forwardingRule.SetEnabled(true)
	forwardingRule.SetUserLocations("Any")
	forwardingRule.SetSourceApplications("Any")
	forwardingRule.SetDestinations("Any")
	forwardingRule.SetConnectivity("direct")

	blockRule := mobile_agent.NewBlockRuleZtna()
	blockRule.SetBlockAllOtherUnmatchedOutboundConnections(false)
	blockRule.SetBlockOutboundLanAccessWhenConnectedToTunnel(false)
	blockRule.SetBlockInboundAccessWhenConnectedToTunnel(false)
	blockRule.SetBlockNonTcpNonUdpTrafficWhenConnectedToTunnel(false)
	blockRule.SetAllowIcmpForTroubleshooting(false)
	blockRule.SetEnforcerFqdnDnsResolutionViaDnsServers(true)
	blockRule.SetResolveAllFqdnsUsingDnsServersAssignedByTheTunnel(true)

	ztnaAgent := mobile_agent.NewForwardingProfileZtnaAgentZtnaAgent()
	ztnaAgent.SetPacUpload(false)
	ztnaAgent.SetForwardingRules([]mobile_agent.ForwardingRuleZtna{*forwardingRule})
	ztnaAgent.SetBlockRule(*blockRule)

	profile := mobile_agent.ForwardingProfiles{
		Name:             profileName,
		Description:      common.StringPtr("Test forwarding profile for create"),
		DefinitionMethod: common.StringPtr("rules"),
		Type: &mobile_agent.ForwardingProfilesType{
			ZtnaAgent: ztnaAgent,
		},
	}

	fmt.Printf("Attempting to create Forwarding Profile with name: %s\n", profile.Name)

	req := client.ForwardingProfilesAPI.CreateGlobalProtectForwardingProfile(context.Background()).
		Folder("Mobile Users").
		ForwardingProfiles(profile)
	res, httpRes, err := req.Execute()

	if err != nil {
		handleAPIError(err)
	}

	require.NoError(t, err, "Failed to create Forwarding Profile")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	require.NotNil(t, res.Id, "Created Forwarding Profile should have an ID")

	createdID := *res.Id

	defer func() {
		t.Logf("Cleaning up Forwarding Profile with ID: %s", createdID)
		_, errDel := client.ForwardingProfilesAPI.DeleteGlobalProtectForwardingProfile(context.Background(), createdID).Execute()
		require.NoError(t, errDel, "Failed to delete Forwarding Profile during cleanup")
	}()

	assert.Equal(t, profileName, res.Name, "Created Forwarding Profile name should match")
	assert.Equal(t, "Test forwarding profile for create", *res.Description, "Description should match")

	require.NotNil(t, res.Type, "Type should not be nil")
	require.NotNil(t, res.Type.ZtnaAgent, "ZtnaAgent type should not be nil")

	returnedZtna := res.Type.ZtnaAgent
	assert.False(t, returnedZtna.GetPacUpload(), "PacUpload should be false")

	require.Len(t, returnedZtna.ForwardingRules, 1, "Should have exactly one forwarding rule")
	rule := returnedZtna.ForwardingRules[0]
	assert.Equal(t, "rule-1", rule.Name, "Forwarding rule name should match")
	assert.Equal(t, "dns", rule.GetTrafficType(), "Traffic type should be dns")
	assert.True(t, rule.GetEnabled(), "Rule should be enabled")
	assert.Equal(t, "Any", rule.GetUserLocations(), "UserLocations should be Any")
	assert.Equal(t, "Any", rule.GetSourceApplications(), "SourceApplications should be Any")
	assert.Equal(t, "Any", rule.GetDestinations(), "Destinations should be Any")
	assert.Equal(t, "direct", rule.GetConnectivity(), "Connectivity should be direct")

	require.NotNil(t, returnedZtna.BlockRule, "BlockRule should not be nil")
	br := returnedZtna.BlockRule
	assert.False(t, br.GetBlockAllOtherUnmatchedOutboundConnections(), "BlockAllOtherUnmatchedOutboundConnections should be false")
	assert.False(t, br.GetBlockOutboundLanAccessWhenConnectedToTunnel(), "BlockOutboundLanAccess should be false")
	assert.False(t, br.GetBlockInboundAccessWhenConnectedToTunnel(), "BlockInboundAccess should be false")
	assert.False(t, br.GetBlockNonTcpNonUdpTrafficWhenConnectedToTunnel(), "BlockNonTcpNonUdp should be false")
	assert.False(t, br.GetAllowIcmpForTroubleshooting(), "AllowIcmp should be false")
	assert.True(t, br.GetEnforcerFqdnDnsResolutionViaDnsServers(), "EnforcerFqdnDnsResolution should be true")
	assert.True(t, br.GetResolveAllFqdnsUsingDnsServersAssignedByTheTunnel(), "ResolveAllFqdns should be true")

	t.Logf("Successfully created and validated Forwarding Profile: %s with ID: %s", profile.Name, createdID)
}

// Test_mobile_agent_ForwardingProfilesAPIService_GetByID tests retrieving a forwarding profile by its ID.
func Test_mobile_agent_ForwardingProfilesAPIService_GetByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	profileName := "test-fwdprofile-get-" + randomSuffix

	createRes, _, err := client.ForwardingProfilesAPI.CreateGlobalProtectForwardingProfile(context.Background()).
		Folder("Mobile Users").
		ForwardingProfiles(newTestForwardingProfile(profileName)).
		Execute()
	require.NoError(t, err, "Failed to create Forwarding Profile for get test")
	createdID := *createRes.Id

	defer func() {
		client.ForwardingProfilesAPI.DeleteGlobalProtectForwardingProfile(context.Background(), createdID).Execute()
	}()

	getRes, httpResGet, errGet := client.ForwardingProfilesAPI.GetGlobalProtectForwardingProfileByID(context.Background(), createdID).Execute()
	require.NoError(t, errGet, "Failed to get Forwarding Profile by ID")
	assert.Equal(t, 200, httpResGet.StatusCode, "Expected 200 OK status")
	require.NotNil(t, getRes, "Get response should not be nil")
	assert.Equal(t, profileName, getRes.Name)
	assert.Equal(t, createdID, *getRes.Id)
}

// Test_mobile_agent_ForwardingProfilesAPIService_Update tests updating an existing forwarding profile.
func Test_mobile_agent_ForwardingProfilesAPIService_Update(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	profileName := "test-fwdprofile-update-" + randomSuffix

	createRes, _, err := client.ForwardingProfilesAPI.CreateGlobalProtectForwardingProfile(context.Background()).
		Folder("Mobile Users").
		ForwardingProfiles(newTestForwardingProfile(profileName)).
		Execute()
	require.NoError(t, err, "Failed to create Forwarding Profile for update test")
	createdID := *createRes.Id

	defer func() {
		client.ForwardingProfilesAPI.DeleteGlobalProtectForwardingProfile(context.Background(), createdID).Execute()
	}()

	updatedProfile := newTestForwardingProfile(profileName)
	updatedProfile.Description = common.StringPtr("Updated description")

	updateRes, httpResUpdate, errUpdate := client.ForwardingProfilesAPI.UpdateGlobalProtectForwardingProfileByID(context.Background(), createdID).
		ForwardingProfiles(updatedProfile).
		Execute()
	require.NoError(t, errUpdate, "Failed to update Forwarding Profile")
	assert.Equal(t, 200, httpResUpdate.StatusCode)
	require.NotNil(t, updateRes)
	assert.Equal(t, profileName, updateRes.Name, "Name should remain the same after update")
	assert.Equal(t, "Updated description", *updateRes.Description, "Description should be updated")
}

// Test_mobile_agent_ForwardingProfilesAPIService_List tests listing forwarding profiles.
func Test_mobile_agent_ForwardingProfilesAPIService_List(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	profileName := "test-fwdprofile-list-" + randomSuffix

	createRes, _, err := client.ForwardingProfilesAPI.CreateGlobalProtectForwardingProfile(context.Background()).
		Folder("Mobile Users").
		ForwardingProfiles(newTestForwardingProfile(profileName)).
		Execute()
	require.NoError(t, err, "Failed to create Forwarding Profile for list test")
	createdID := *createRes.Id

	defer func() {
		client.ForwardingProfilesAPI.DeleteGlobalProtectForwardingProfile(context.Background(), createdID).Execute()
	}()

	listRes, httpResList, errList := client.ForwardingProfilesAPI.ListGlobalProtectForwardingProfiles(context.Background()).
		Folder("Mobile Users").
		Limit(10000).
		Execute()
	require.NoError(t, errList, "Failed to list Forwarding Profiles")
	assert.Equal(t, 200, httpResList.StatusCode)
	require.NotNil(t, listRes)

	foundProfile := false
	for _, p := range listRes.Data {
		if p.Name == profileName {
			foundProfile = true
			break
		}
	}
	assert.True(t, foundProfile, "Created Forwarding Profile should be found in the list")
}

// Test_mobile_agent_ForwardingProfilesAPIService_DeleteByID tests deleting a forwarding profile by ID.
func Test_mobile_agent_ForwardingProfilesAPIService_DeleteByID(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	profileName := "test-fwdprofile-delete-" + randomSuffix

	createRes, _, err := client.ForwardingProfilesAPI.CreateGlobalProtectForwardingProfile(context.Background()).
		Folder("Mobile Users").
		ForwardingProfiles(newTestForwardingProfile(profileName)).
		Execute()
	require.NoError(t, err, "Failed to create Forwarding Profile for delete test")
	createdID := *createRes.Id

	_, errDel := client.ForwardingProfilesAPI.DeleteGlobalProtectForwardingProfile(context.Background(), createdID).Execute()
	require.NoError(t, errDel, "Failed to delete Forwarding Profile")
}

// Test_mobile_agent_ForwardingProfilesAPIService_FetchForwardingProfiles tests the FetchForwardingProfiles convenience method.
func Test_mobile_agent_ForwardingProfilesAPIService_FetchForwardingProfiles(t *testing.T) {
	client := SetupMobileAgentTestClient(t)
	randomSuffix := common.GenerateRandomString(6)
	testName := "test-fwdprofile-fetch-" + randomSuffix

	testObj := newTestForwardingProfile(testName)
	testObj.Description = common.StringPtr("Test forwarding profile for fetch")

	createReq := client.ForwardingProfilesAPI.CreateGlobalProtectForwardingProfile(context.Background()).
		Folder("Mobile Users").
		ForwardingProfiles(testObj)
	createRes, _, err := createReq.Execute()
	if err != nil {
		handleAPIError(err)
	}
	require.NoError(t, err, "Failed to create test object for fetch test")
	require.NotNil(t, createRes, "Create response should not be nil")
	createdID := *createRes.Id

	defer func() {
		client.ForwardingProfilesAPI.DeleteGlobalProtectForwardingProfile(context.Background(), createdID).Execute()
		t.Logf("Cleaned up test object: %s", createdID)
	}()

	// Test 1: Fetch existing object by name
	fetchedObj, err := client.ForwardingProfilesAPI.FetchForwardingProfiles(
		context.Background(),
		testName,
		common.StringPtr("Mobile Users"),
		nil, // snippet
		nil, // device
	)

	require.NoError(t, err, "Failed to fetch forwarding_profiles by name")
	require.NotNil(t, fetchedObj, "Fetched object should not be nil")
	assert.Equal(t, createdID, *fetchedObj.Id, "Fetched object ID should match")
	assert.Equal(t, testName, fetchedObj.Name, "Fetched object name should match")
	t.Logf("[SUCCESS] FetchForwardingProfiles found object: %s", fetchedObj.Name)

	// Test 2: Fetch non-existent object (should return nil, nil)
	notFound, err := client.ForwardingProfilesAPI.FetchForwardingProfiles(
		context.Background(),
		"non-existent-forwarding-profile-xyz-12345",
		common.StringPtr("Mobile Users"),
		nil,
		nil,
	)
	require.NoError(t, err, "Fetch should not error for non-existent object")
	assert.Nil(t, notFound, "Should return nil for non-existent object")
	t.Logf("[SUCCESS] FetchForwardingProfiles correctly returned nil for non-existent object")
}
