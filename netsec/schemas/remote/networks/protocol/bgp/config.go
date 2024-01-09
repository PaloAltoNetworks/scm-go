package bgp

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/remote-networks-protocol-bgp

/*
Config object.

ShortName: fVAkWHS
Parent chains:
*

Args:

Param DoNotExportRoutes (bool): the DoNotExportRoutes param.

Param Enable (bool): to setup bgp protocol, enable need to set as true

Param LocalIpAddress (string): the LocalIpAddress param.

Param OriginateDefaultRoute (bool): the OriginateDefaultRoute param.

Param PeerAs (string): the PeerAs param.

Param PeerIpAddress (string): the PeerIpAddress param.

Param PeeringType (string): Exchange Routes: exchange-v4-over-v4 stands for Exchange IPv4 routes over IPv4 peering. exchange-v4-v6-over-v4 stands for Exchange both IPv4 and IPv6 routes over IPv4 peering. exchange-v4-over-v4-v6-over-v6 stands for Exchange IPv4 routes over IPv4 peer and IPv6 route over IPv6 peer. exchange-v6-over-v6 stands for Exchange IPv6 routes over IPv6 peering. String must be one of these: `"exchange-v4-over-v4"`, `"exchange-v4-v6-over-v4"`, `"exchange-v4-over-v4-v6-over-v6"`, `"exchange-v6-over-v6"`.

Param Secret (string): the Secret param.

Param SummarizeMobileUserRoutes (bool): the SummarizeMobileUserRoutes param.
*/
type Config struct {
	DoNotExportRoutes         *bool   `json:"do_not_export_routes,omitempty"`
	Enable                    *bool   `json:"enable,omitempty"`
	LocalIpAddress            *string `json:"local_ip_address,omitempty"`
	OriginateDefaultRoute     *bool   `json:"originate_default_route,omitempty"`
	PeerAs                    *string `json:"peer_as,omitempty"`
	PeerIpAddress             *string `json:"peer_ip_address,omitempty"`
	PeeringType               *string `json:"peering_type,omitempty"`
	Secret                    *string `json:"secret,omitempty"`
	SummarizeMobileUserRoutes *bool   `json:"summarize_mobile_user_routes,omitempty"`
}
