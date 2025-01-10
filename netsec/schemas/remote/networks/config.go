package networks

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/remote-networks

import (
	fVAkWHS "github.com/paloaltonetworks/scm-go/netsec/schemas/remote/networks/protocol/bgp"
)

/*
Config object.

ShortName: uewNibC
Parent chains:
*

Args:

Param EcmpLoadBalancing (string): the EcmpLoadBalancing param. String must be one of these: `"enable"`, `"disable"`. Default: `"disable"`.

Param EcmpTunnels ([]EcmpTunnelObject): ecmp_tunnels is required when ecmp_load_balancing is enable

Param Id (string, read-only): UUID of the resource

Param IpsecTunnel (string): ipsec_tunnel is required when ecmp_load_balancing is disable

Param LicenseType (string, required): New customer will only be on aggregate bandwidth licensing String length must exceed 1 characters. Default: `"FWAAS-AGGREGATE"`.

Param Name (string, required): Alphanumeric string begin with letter: [0-9a-zA-Z._-] String length must not exceed 63 characters.

Param Protocol (ProtocolObject): setup the protocol when ecmp_load_balancing is disable

Param Region (string, required): the Region param. String length must exceed 1 characters.

Param SecondaryIpsecTunnel (string): specify secondary ipsec_tunnel if needed

Param SpnName (string): spn-name is needed when license_type is FWAAS-AGGREGATE

Param Subnets ([]string): the Subnets param.
*/
type Config struct {
	EcmpLoadBalancing    *string            `json:"ecmp_load_balancing,omitempty"`
	EcmpTunnels          []EcmpTunnelObject `json:"ecmp_tunnels,omitempty"`
	Id                   *string            `json:"id,omitempty"`
	IpsecTunnel          *string            `json:"ipsec_tunnel,omitempty"`
	LicenseType          string             `json:"license_type"`
	Name                 string             `json:"name"`
	Protocol             *ProtocolObject    `json:"protocol,omitempty"`
	Region               string             `json:"region"`
	SecondaryIpsecTunnel *string            `json:"secondary_ipsec_tunnel,omitempty"`
	SpnName              *string            `json:"spn_name,omitempty"`
	Subnets              []string           `json:"subnets,omitempty"`
}

/*
EcmpTunnelObject object.

ShortName:
Parent chains:
*
* ecmp_tunnels
* _inline

Args:

Param IpsecTunnel (string, required): the IpsecTunnel param.

Param Name (string, required): the Name param.

Param Protocol (EcmpProtocolObject, required): the Protocol param.
*/
type EcmpTunnelObject struct {
	IpsecTunnel string             `json:"ipsec_tunnel"`
	Name        string             `json:"name"`
	Protocol    EcmpProtocolObject `json:"protocol"`
}

/*
EcmpProtocolObject object.

ShortName:
Parent chains:
*
* ecmp_tunnels
* _inline
* protocol

Args:

Param Bgp (fVAkWHS.Config): the Bgp param.
*/
type EcmpProtocolObject struct {
	Bgp *fVAkWHS.Config `json:"bgp,omitempty"`
}

/*
ProtocolObject setup the protocol when ecmp_load_balancing is disable

ShortName:
Parent chains:
*
* protocol

Args:

Param Bgp (fVAkWHS.Config): the Bgp param.

Param BgpPeer (BgpPeerObject): secondary bgp routing as bgp_peer
*/
type ProtocolObject struct {
	Bgp     *fVAkWHS.Config `json:"bgp,omitempty"`
	BgpPeer *BgpPeerObject  `json:"bgp_peer,omitempty"`
}

/*
BgpPeerObject secondary bgp routing as bgp_peer

ShortName:
Parent chains:
*
* protocol
* bgp_peer

Args:

Param LocalIpAddress (string): the LocalIpAddress param.

Param PeerIpAddress (string): the PeerIpAddress param.

Param SameAsPrimary (bool): If true, the secondary BGP peer configuration will be the same as the primary BGP peer. Default: `true`.

Param Secret (string): the Secret param.
*/
type BgpPeerObject struct {
	LocalIpAddress *string `json:"local_ip_address,omitempty"`
	PeerIpAddress  *string `json:"peer_ip_address,omitempty"`
	SameAsPrimary  *bool   `json:"same_as_primary,omitempty"`
	Secret         *string `json:"secret,omitempty"`
}
