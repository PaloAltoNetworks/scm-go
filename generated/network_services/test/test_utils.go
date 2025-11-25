// Package network_services
/*
Testing utilities for network_services API
Shared utilities for testing network_services API services
*/
package network_services

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	setup "github.com/paloaltonetworks/scm-go"
	"github.com/paloaltonetworks/scm-go/common"
	"github.com/paloaltonetworks/scm-go/generated/network_services"
)

func SetupNetworkSvcTestClient(t *testing.T) *network_services.APIClient {
	configPath := common.GetConfigPath()
	setupClient := &setup.Client{
		AuthFile:         configPath,
		CheckEnvironment: false,
	}

	fmt.Printf("Using config file: %s\n", setupClient.AuthFile)

	// Setup the client configuration
	err := setupClient.Setup()
	require.NoError(t, err, "Failed to setup client")

	// Refresh JWT token
	ctx := context.Background()
	if setupClient.Jwt == "" {
		fmt.Printf("\n****************\nGetting tokens\n***********************\n")
		err = setupClient.RefreshJwt(ctx)
		if err != nil {
			// Print detailed error information
			fmt.Printf("=== JWT REFRESH ERROR ===\n")
			fmt.Printf("Error: %v\n", err)
			fmt.Printf("Error Type: %T\n", err)
			fmt.Printf("Error String: %s\n", err.Error())
			fmt.Printf("=========================\n")
		}
		// Fail the test only after all retries have been exhausted.
		require.NoError(t, err, "Failed to refresh JWT after multiple retries")
	}

	return setup.GetNetwork_servicesAPIClient(setupClient)
}

// printAPIError prints formatted API error response from error object's body
func printAPIError(err *network_services.GenericOpenAPIError) {
	if err == nil {
		return
	}
	fmt.Printf("=== API ERROR RESPONSE ===\n")
	fmt.Printf("Error: %v\n", err)
	bodyBytes := err.Body()
	if bodyBytes == nil {
		fmt.Printf("No body found in error object\n")
		fmt.Printf("===========================\n\n")
		return
	}
	if len(bodyBytes) == 0 {
		fmt.Printf("No body found in error object\n")
		fmt.Printf("===========================\n\n")
		return
	}
	// Print raw JSON response
	fmt.Printf("Raw Error Body:\n%s\n", string(bodyBytes))
	fmt.Printf("===========================\n\n")
}

// handleAPIError is a utility method to consistently handle and print API errors
func handleAPIError(err error) {
	if err == nil {
		return
	}
	// Print detailed error information if it's a GenericOpenAPIError
	var apiErr *network_services.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		printAPIError(apiErr)
	} else {
		// Print the regular error if it's not a GenericOpenAPIError
		fmt.Printf("Non-API Error: %v\n", err)
	}
}

// CreateTestIKECryptoProfile creates a test IKE Crypto Profile dependency.
// It returns the ID of the created profile.
func CreateTestIKECryptoProfile(t *testing.T, client *network_services.APIClient, name string, optionalFolder ...string) string {
	folderName := "Remote Networks"
	if len(optionalFolder) > 0 && optionalFolder[0] != "" {
		folderName = optionalFolder[0]
	}

	profile := network_services.IkeCryptoProfiles{
		Folder:     common.StringPtr(folderName), // A common folder for these tests
		Name:       name,
		Hash:       []string{"sha256"},
		DhGroup:    []string{"group14"},
		Encryption: []string{"aes-256-cbc"},
	}
	req := client.IKECryptoProfilesAPI.CreateIKECryptoProfiles(context.Background()).IkeCryptoProfiles(profile)
	res, _, err := req.Execute()
	require.NoError(t, err, "Failed to create test IKE Crypto Profile dependency")
	require.NotNil(t, res, "Test IKE Crypto Profile create response should not be nil")
	t.Logf("Created test IKE Crypto Profile '%s' with ID %s", name, *res.Id)
	return *res.Id
}

// DeleteTestIKECryptoProfile deletes a test IKE Crypto Profile dependency.
func DeleteTestIKECryptoProfile(t *testing.T, client *network_services.APIClient, id, name string) {
	req := client.IKECryptoProfilesAPI.DeleteIKECryptoProfilesByID(context.Background(), id)
	_, err := req.Execute()
	require.NoError(t, err, "Failed to delete test IKE Crypto Profile '%s'", name)
	t.Logf("Deleted test IKE Crypto Profile '%s'", name)
}

// CreateTestIPsecCryptoProfile creates a test IPsec Crypto Profile dependency.
// It returns the ID of the created profile.
func CreateTestIPsecCryptoProfile(t *testing.T, client *network_services.APIClient, name string) string {
	profile := network_services.IpsecCryptoProfiles{
		Folder:  common.StringPtr("Prisma Access"), // A common folder for these tests
		Name:    name,
		DhGroup: common.StringPtr("group14"),
		Esp: &network_services.IpsecCryptoProfilesEsp{
			Authentication: []string{"sha256"},
			Encryption:     []string{"aes-256-gcm"},
		},
		Lifetime: network_services.IpsecCryptoProfilesLifetime{
			Hours: common.Int32Ptr(8),
		},
	}
	req := client.IPsecCryptoProfilesAPI.CreateIPsecCryptoProfiles(context.Background()).IpsecCryptoProfiles(profile)
	res, _, err := req.Execute()
	require.NoError(t, err, "Failed to create test IPsec Crypto Profile dependency")
	require.NotNil(t, res, "Test IPsec Crypto Profile create response should not be nil")
	t.Logf("Created test IPsec Crypto Profile '%s' with ID %s", name, *res.Id)
	return *res.Id
}

// DeleteTestIPsecCryptoProfile deletes a test IPsec Crypto Profile dependency.
func DeleteTestIPsecCryptoProfile(t *testing.T, client *network_services.APIClient, id, name string) {
	req := client.IPsecCryptoProfilesAPI.DeleteIPsecCryptoProfilesByID(context.Background(), id)
	_, err := req.Execute()
	require.NoError(t, err, "Failed to delete test IPsec Crypto Profile '%s'", name)
	t.Logf("Deleted test IPsec Crypto Profile '%s'", name)
}

// CreateIkeGatewayTestObject Helper function to create the standard test object.
func CreateIkeGatewayTestObject(name string, cryptoProfileName string, optionalFolder ...string) network_services.IkeGateways {
	folderName := "Remote Networks"
	if len(optionalFolder) > 0 && optionalFolder[0] != "" {
		folderName = optionalFolder[0]
	}
	return network_services.IkeGateways{
		Folder: common.StringPtr(folderName),
		Name:   name,
		Authentication: network_services.IkeGatewaysAuthentication{
			PreSharedKey: &network_services.IkeGatewaysAuthenticationPreSharedKey{
				Key: common.StringPtr("123456"),
			},
		},
		PeerAddress: network_services.IkeGatewaysPeerAddress{
			Ip: common.StringPtr("2.2.2.4"),
		},
		PeerId: &network_services.IkeGatewaysPeerId{
			Type: common.StringPtr("ipaddr"),
			Id:   common.StringPtr("10.3.3.4"),
		},
		LocalId: &network_services.IkeGatewaysLocalId{
			Type: common.StringPtr("ipaddr"),
			Id:   common.StringPtr("10.3.4.4"),
		},
		Protocol: network_services.IkeGatewaysProtocol{
			Ikev1: &network_services.IkeGatewaysProtocolIkev1{
				IkeCryptoProfile: common.StringPtr(cryptoProfileName), // Use the created profile
				Dpd: &network_services.IkeGatewaysProtocolIkev1Dpd{
					Enable: common.BoolPtr(true),
				},
			},
			Ikev2: &network_services.IkeGatewaysProtocolIkev1{
				IkeCryptoProfile: common.StringPtr(cryptoProfileName), // Use the created profile
				Dpd: &network_services.IkeGatewaysProtocolIkev1Dpd{
					Enable: common.BoolPtr(true),
				},
			},
		},
	}
}

// CreateTestIkeGateway Creates ike gateway given a gateway name and ikeCryptoProfileName
func CreateTestIkeGateway(t *testing.T, client *network_services.APIClient, gatewayName string, ikeCryptoProfileName string, optionalFolder ...string) string {
	gateway := CreateIkeGatewayTestObject(gatewayName, ikeCryptoProfileName, optionalFolder...)

	fmt.Printf("Attempting to create IKE Gateway with name: %s\n", gateway.Name)

	// Make the create request to the API.
	req := client.IKEGatewaysAPI.CreateIKEGateways(context.Background()).IkeGateways(gateway)
	res, httpRes, err := req.Execute()

	// Verify the request was successful.
	require.NoError(t, err, "Create request should not return an error")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "The response from create should not be nil")
	assert.Equal(t, gatewayName, res.Name, "The name of the created gateway should match")
	assert.NotEmpty(t, *res.Id, "The ID of the created gateway should not be empty")

	t.Logf("Successfully created IKE Gateway with ID: %s", *res.Id)
	return *res.Id
}

// CreateTestIPSecTunnel Creates ipsec tunnel given a ipsec tunnel name and gateway name
func CreateTestIPSecTunnel(t *testing.T, client *network_services.APIClient, ipsecTunnelName string, gatewayName string, optionalFolder ...string) string {
	folderName := "Remote Networks"
	if len(optionalFolder) > 0 && optionalFolder[0] != "" {
		folderName = optionalFolder[0]
	}
	tunnel := network_services.IpsecTunnels{
		Folder:                 common.StringPtr(folderName),
		Name:                   ipsecTunnelName,
		AntiReplay:             common.BoolPtr(true),
		CopyTos:                common.BoolPtr(false),
		EnableGreEncapsulation: common.BoolPtr(false),
		AutoKey: network_services.IpsecTunnelsAutoKey{
			IkeGateway: []network_services.IpsecTunnelsAutoKeyIkeGatewayInner{
				{
					Name: common.StringPtr(gatewayName),
				},
			},
			IpsecCryptoProfile: "PaloAlto-Networks-IPSec-Crypto",
		},
	}

	fmt.Printf("Creating IPsec tunnel with name: %s\n", tunnel.Name)
	req := client.IPsecTunnelsAPI.CreateIPsecTunnels(context.Background()).IpsecTunnels(tunnel)
	res, httpRes, err := req.Execute()

	// Verify creation was successful.
	require.NoError(t, err, "Failed to create IPsec tunnel")
	assert.Equal(t, 201, httpRes.StatusCode, "Expected 201 Created status")
	require.NotNil(t, res, "Response should not be nil")
	assert.Equal(t, ipsecTunnelName, res.Name)

	return *res.Id
}
