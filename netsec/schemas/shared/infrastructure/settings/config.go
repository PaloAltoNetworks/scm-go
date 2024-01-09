package settings

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/shared-infrastructure-settings

/*
Config object.

ShortName: eItpily
Parent chains:
*

Args:

Param ApiKey (string): the ApiKey param.

Param CaptivePortalRedirectIpAddress (string): the CaptivePortalRedirectIpAddress param.

Param EgressIpNotificationUrl (string): the EgressIpNotificationUrl param.

Param InfraBgpAs (string): the InfraBgpAs param.

Param InfrastructureSubnet (string): the InfrastructureSubnet param.

Param InfrastructureSubnetIpv6 (string): the InfrastructureSubnetIpv6 param.

Param Ipv6 (bool): the Ipv6 param.

Param LoopbackIps ([]string): the LoopbackIps param.

Param TunnelMonitorIpAddress (string): the TunnelMonitorIpAddress param.
*/
type Config struct {
	ApiKey                         *string  `json:"api_key,omitempty"`
	CaptivePortalRedirectIpAddress *string  `json:"captive_portal_redirect_ip_address,omitempty"`
	EgressIpNotificationUrl        *string  `json:"egress_ip_notification_url,omitempty"`
	InfraBgpAs                     *string  `json:"infra_bgp_as,omitempty"`
	InfrastructureSubnet           *string  `json:"infrastructure_subnet,omitempty"`
	InfrastructureSubnetIpv6       *string  `json:"infrastructure_subnet_ipv6,omitempty"`
	Ipv6                           *bool    `json:"ipv6,omitempty"`
	LoopbackIps                    []string `json:"loopback_ips,omitempty"`
	TunnelMonitorIpAddress         *string  `json:"tunnel_monitor_ip_address,omitempty"`
}
