package connections

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/service-connections

/*
Config object.

ShortName: wugpput
Parent chains:
*

Args:

Param BackupSC (string): the BackupSC param.

Param BgpPeer (BgpPeerObject): the BgpPeer param.

Param Id (string, read-only): UUID of the resource

Param IpsecTunnel (string, required): the IpsecTunnel param.

Param Name (string, required): the Name param.

Param NatPool (string): the NatPool param.

Param NoExportCommunity (string): the NoExportCommunity param. String must be one of these: `"Disabled"`, `"Enabled-In"`, `"Enabled-Out"`, `"Enabled-Both"`.

Param OnboardingType (string): the OnboardingType param. String must be one of these: `"classic"`. Default: `"classic"`.

Param Protocol (ProtocolObject): the Protocol param.

Param Qos (QosObject): the Qos param.

Param Region (string, required): the Region param.

Param SecondaryIpsecTunnel (string): the SecondaryIpsecTunnel param.

Param SourceNat (bool): the SourceNat param.

Param Subnets ([]string): the Subnets param.
*/
type Config struct {
	BackupSC             *string         `json:"backup_SC,omitempty"`
	BgpPeer              *BgpPeerObject  `json:"bgp_peer,omitempty"`
	Id                   *string         `json:"id,omitempty"`
	IpsecTunnel          string          `json:"ipsec_tunnel"`
	Name                 string          `json:"name"`
	NatPool              *string         `json:"nat_pool,omitempty"`
	NoExportCommunity    *string         `json:"no_export_community,omitempty"`
	OnboardingType       *string         `json:"onboarding_type,omitempty"`
	Protocol             *ProtocolObject `json:"protocol,omitempty"`
	Qos                  *QosObject      `json:"qos,omitempty"`
	Region               string          `json:"region"`
	SecondaryIpsecTunnel *string         `json:"secondary_ipsec_tunnel,omitempty"`
	SourceNat            *bool           `json:"source_nat,omitempty"`
	Subnets              []string        `json:"subnets,omitempty"`
}

/*
BgpPeerObject object.

ShortName:
Parent chains:
*
* bgp_peer

Args:

Param LocalIpAddress (string): the LocalIpAddress param.

Param LocalIpv6Address (string): the LocalIpv6Address param.

Param PeerIpAddress (string): the PeerIpAddress param.

Param PeerIpv6Address (string): the PeerIpv6Address param.

Param SameAsPrimary (bool): the SameAsPrimary param.

Param Secret (string): the Secret param.
*/
type BgpPeerObject struct {
	LocalIpAddress   *string `json:"local_ip_address,omitempty"`
	LocalIpv6Address *string `json:"local_ipv6_address,omitempty"`
	PeerIpAddress    *string `json:"peer_ip_address,omitempty"`
	PeerIpv6Address  *string `json:"peer_ipv6_address,omitempty"`
	SameAsPrimary    *bool   `json:"same_as_primary,omitempty"`
	Secret           *string `json:"secret,omitempty"`
}

/*
ProtocolObject object.

ShortName:
Parent chains:
*
* protocol

Args:

Param Bgp (BgpObject): the Bgp param.
*/
type ProtocolObject struct {
	Bgp *BgpObject `json:"bgp,omitempty"`
}

/*
BgpObject object.

ShortName:
Parent chains:
*
* protocol
* bgp

Args:

Param DoNotExportRoutes (bool): the DoNotExportRoutes param.

Param Enable (bool): the Enable param.

Param FastFailover (bool): the FastFailover param.

Param LocalIpAddress (string): the LocalIpAddress param.

Param OriginateDefaultRoute (bool): the OriginateDefaultRoute param.

Param PeerAs (string): the PeerAs param.

Param PeerIpAddress (string): the PeerIpAddress param.

Param Secret (string): the Secret param.

Param SummarizeMobileUserRoutes (bool): the SummarizeMobileUserRoutes param.
*/
type BgpObject struct {
	DoNotExportRoutes         *bool   `json:"do_not_export_routes,omitempty"`
	Enable                    *bool   `json:"enable,omitempty"`
	FastFailover              *bool   `json:"fast_failover,omitempty"`
	LocalIpAddress            *string `json:"local_ip_address,omitempty"`
	OriginateDefaultRoute     *bool   `json:"originate_default_route,omitempty"`
	PeerAs                    *string `json:"peer_as,omitempty"`
	PeerIpAddress             *string `json:"peer_ip_address,omitempty"`
	Secret                    *string `json:"secret,omitempty"`
	SummarizeMobileUserRoutes *bool   `json:"summarize_mobile_user_routes,omitempty"`
}

/*
QosObject object.

ShortName:
Parent chains:
*
* qos

Args:

Param Enable (bool): the Enable param.

Param QosProfile (string): the QosProfile param.
*/
type QosObject struct {
	Enable     *bool   `json:"enable,omitempty"`
	QosProfile *string `json:"qos_profile,omitempty"`
}
