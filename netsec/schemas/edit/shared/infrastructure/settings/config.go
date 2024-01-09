package settings

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/edit-shared-infrastructure-settings

/*
Config object.

ShortName: ctlHcHg
Parent chains:
*

Args:

Param EgressIpNotificationUrl (string): the EgressIpNotificationUrl param.

Param InfraBgpAs (string): the InfraBgpAs param.

Param InfrastructureSubnet (string): the InfrastructureSubnet param.

Param InfrastructureSubnetIpv6 (string): the InfrastructureSubnetIpv6 param.
*/
type Config struct {
	EgressIpNotificationUrl  *string `json:"egress_ip_notification_url,omitempty"`
	InfraBgpAs               *string `json:"infra_bgp_as,omitempty"`
	InfrastructureSubnet     *string `json:"infrastructure_subnet,omitempty"`
	InfrastructureSubnetIpv6 *string `json:"infrastructure_subnet_ipv6,omitempty"`
}
