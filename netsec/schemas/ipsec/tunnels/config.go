package tunnels

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/ipsec-tunnels

/*
Config object.

ShortName: nnsRzDg
Parent chains:
*

Args:

Param AntiReplay (bool): Enable Anti-Replay check on this tunnel

Param AutoKey (AutoKeyObject, required): the AutoKey param.

Param CopyTos (bool): Copy IP TOS bits from inner packet to IPSec packet (not recommended) Default: `false`.

Param EnableGreEncapsulation (bool): allow GRE over IPSec Default: `false`.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string begin with letter: [0-9a-zA-Z._-] String length must not exceed 63 characters.

Param TunnelMonitor (TunnelMonitorObject): the TunnelMonitor param.
*/
type Config struct {
	AntiReplay             *bool                `json:"anti_replay,omitempty"`
	AutoKey                AutoKeyObject        `json:"auto_key"`
	CopyTos                *bool                `json:"copy_tos,omitempty"`
	EnableGreEncapsulation *bool                `json:"enable_gre_encapsulation,omitempty"`
	Id                     *string              `json:"id,omitempty"`
	Name                   string               `json:"name"`
	TunnelMonitor          *TunnelMonitorObject `json:"tunnel_monitor,omitempty"`
}

/*
AutoKeyObject object.

ShortName:
Parent chains:
*
* auto_key

Args:

Param IkeGateways ([]IkeGatewayObject, required): the IkeGateways param.

Param IpsecCryptoProfile (string, required): the IpsecCryptoProfile param.

Param ProxyIdV6s ([]ProxyIdV6Object): IPv6 type of proxy_id values

Param ProxyIds ([]ProxyIdObject): IPv4 type of proxy_id values
*/
type AutoKeyObject struct {
	IkeGateways        []IkeGatewayObject `json:"ike_gateway"`
	IpsecCryptoProfile string             `json:"ipsec_crypto_profile"`
	ProxyIdV6s         []ProxyIdV6Object  `json:"proxy_id_v6,omitempty"`
	ProxyIds           []ProxyIdObject    `json:"proxy_id,omitempty"`
}

/*
IkeGatewayObject object.

ShortName:
Parent chains:
*
* auto_key
* ike_gateway
* _inline

Args:

Param Name (string): the Name param.
*/
type IkeGatewayObject struct {
	Name *string `json:"name,omitempty"`
}

/*
ProxyIdV6Object object.

ShortName:
Parent chains:
*
* auto_key
* proxy_id_v6
* _inline

Args:

Param Local (string): the Local param.

Param Name (string, required): the Name param.

Param Protocol (ProxyIdV6ProtocolObject): the Protocol param.

Param Remote (string): the Remote param.
*/
type ProxyIdV6Object struct {
	Local    *string                  `json:"local,omitempty"`
	Name     string                   `json:"name"`
	Protocol *ProxyIdV6ProtocolObject `json:"protocol,omitempty"`
	Remote   *string                  `json:"remote,omitempty"`
}

/*
ProxyIdV6ProtocolObject object.

ShortName:
Parent chains:
*
* auto_key
* proxy_id_v6
* _inline
* protocol

Args:

Param Number (int64): IP protocol number Value must be between 1 and 254.

Param Tcp (ProxyIdV6TcpProtocolObject): the Tcp param.

Param Udp (ProxyIdV6UdpProtocolObject): the Udp param.

NOTE:  One of the following params should be specified:
  - Number
  - Tcp
  - Udp
*/
type ProxyIdV6ProtocolObject struct {
	Number *int64                      `json:"number,omitempty"`
	Tcp    *ProxyIdV6TcpProtocolObject `json:"tcp,omitempty"`
	Udp    *ProxyIdV6UdpProtocolObject `json:"udp,omitempty"`
}

/*
ProxyIdV6TcpProtocolObject object.

ShortName:
Parent chains:
*
* auto_key
* proxy_id_v6
* _inline
* protocol
* tcp

Args:

Param LocalPort (int64): the LocalPort param. Value must be between 0 and 65535. Default: `0`.

Param RemotePort (int64): the RemotePort param. Value must be between 0 and 65535. Default: `0`.
*/
type ProxyIdV6TcpProtocolObject struct {
	LocalPort  *int64 `json:"local_port,omitempty"`
	RemotePort *int64 `json:"remote_port,omitempty"`
}

/*
ProxyIdV6UdpProtocolObject object.

ShortName:
Parent chains:
*
* auto_key
* proxy_id_v6
* _inline
* protocol
* udp

Args:

Param LocalPort (int64): the LocalPort param. Value must be between 0 and 65535. Default: `0`.

Param RemotePort (int64): the RemotePort param. Value must be between 0 and 65535. Default: `0`.
*/
type ProxyIdV6UdpProtocolObject struct {
	LocalPort  *int64 `json:"local_port,omitempty"`
	RemotePort *int64 `json:"remote_port,omitempty"`
}

/*
ProxyIdObject object.

ShortName:
Parent chains:
*
* auto_key
* proxy_id
* _inline

Args:

Param Local (string): the Local param.

Param Name (string, required): the Name param.

Param Protocol (ProxyIdProtocolObject): the Protocol param.

Param Remote (string): the Remote param.
*/
type ProxyIdObject struct {
	Local    *string                `json:"local,omitempty"`
	Name     string                 `json:"name"`
	Protocol *ProxyIdProtocolObject `json:"protocol,omitempty"`
	Remote   *string                `json:"remote,omitempty"`
}

/*
ProxyIdProtocolObject object.

ShortName:
Parent chains:
*
* auto_key
* proxy_id
* _inline
* protocol

Args:

Param Number (int64): IP protocol number Value must be between 1 and 254.

Param Tcp (ProxyIdTcpProtocolObject): the Tcp param.

Param Udp (ProxyIdUdpProtocolObject): the Udp param.

NOTE:  One of the following params should be specified:
  - Number
  - Tcp
  - Udp
*/
type ProxyIdProtocolObject struct {
	Number *int64                    `json:"number,omitempty"`
	Tcp    *ProxyIdTcpProtocolObject `json:"tcp,omitempty"`
	Udp    *ProxyIdUdpProtocolObject `json:"udp,omitempty"`
}

/*
ProxyIdTcpProtocolObject object.

ShortName:
Parent chains:
*
* auto_key
* proxy_id
* _inline
* protocol
* tcp

Args:

Param LocalPort (int64): the LocalPort param. Value must be between 0 and 65535. Default: `0`.

Param RemotePort (int64): the RemotePort param. Value must be between 0 and 65535. Default: `0`.
*/
type ProxyIdTcpProtocolObject struct {
	LocalPort  *int64 `json:"local_port,omitempty"`
	RemotePort *int64 `json:"remote_port,omitempty"`
}

/*
ProxyIdUdpProtocolObject object.

ShortName:
Parent chains:
*
* auto_key
* proxy_id
* _inline
* protocol
* udp

Args:

Param LocalPort (int64): the LocalPort param. Value must be between 0 and 65535. Default: `0`.

Param RemotePort (int64): the RemotePort param. Value must be between 0 and 65535. Default: `0`.
*/
type ProxyIdUdpProtocolObject struct {
	LocalPort  *int64 `json:"local_port,omitempty"`
	RemotePort *int64 `json:"remote_port,omitempty"`
}

/*
TunnelMonitorObject object.

ShortName:
Parent chains:
*
* tunnel_monitor

Args:

Param DestinationIp (string, required): Destination IP to send ICMP probe

Param Enable (bool): Enable tunnel monitoring on this tunnel Default: `true`.

Param ProxyId (string): Which proxy-id (or proxy-id-v6) the monitoring traffic will use
*/
type TunnelMonitorObject struct {
	DestinationIp string  `json:"destination_ip"`
	Enable        *bool   `json:"enable,omitempty"`
	ProxyId       *string `json:"proxy_id,omitempty"`
}
