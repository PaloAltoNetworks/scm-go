package services

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/services

/*
Config object.

ShortName: ujXZojh
Parent chains:
*

Args:

Param Description (string): the Description param. String length must not exceed 1023 characters.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 63 characters.

Param Protocol (ProtocolObject, required): the Protocol param.

Param Tags ([]string): Tags for service object List must contain at most 64 elements. Individual elements in this list are subject to additional validation. String length must not exceed 127 characters.
*/
type Config struct {
	Description *string        `json:"description,omitempty"`
	Id          *string        `json:"id,omitempty"`
	Name        string         `json:"name"`
	Protocol    ProtocolObject `json:"protocol"`
	Tags        []string       `json:"tag,omitempty"`
}

/*
ProtocolObject object.

ShortName:
Parent chains:
*
* protocol

Args:

Param Tcp (TcpObject): the Tcp param.

Param Udp (UdpObject): the Udp param.

NOTE:  One of the following params should be specified:
  - Tcp
  - Udp
*/
type ProtocolObject struct {
	Tcp *TcpObject `json:"tcp,omitempty"`
	Udp *UdpObject `json:"udp,omitempty"`
}

/*
TcpObject object.

ShortName:
Parent chains:
*
* protocol
* tcp

Args:

Param Override (TcpOverrideObject): the Override param.

Param Port (string, required): the Port param. String length must be between 1 and 1023 characters.

Param SourcePort (string): the SourcePort param. String length must be between 1 and 1023 characters.
*/
type TcpObject struct {
	Override   *TcpOverrideObject `json:"override,omitempty"`
	Port       string             `json:"port"`
	SourcePort *string            `json:"source_port,omitempty"`
}

/*
TcpOverrideObject object.

ShortName:
Parent chains:
*
* protocol
* tcp
* override

Args:

Param HalfcloseTimeout (int64): tcp session half-close timeout value (in second) Value must be between 1 and 604800. Default: `120`.

Param Timeout (int64): tcp session timeout value (in second) Value must be between 1 and 604800. Default: `3600`.

Param TimewaitTimeout (int64): tcp session time-wait timeout value (in second) Value must be between 1 and 600. Default: `15`.
*/
type TcpOverrideObject struct {
	HalfcloseTimeout *int64 `json:"halfclose_timeout,omitempty"`
	Timeout          *int64 `json:"timeout,omitempty"`
	TimewaitTimeout  *int64 `json:"timewait_timeout,omitempty"`
}

/*
UdpObject object.

ShortName:
Parent chains:
*
* protocol
* udp

Args:

Param Override (UdpOverrideObject): the Override param.

Param Port (string, required): the Port param. String length must be between 1 and 1023 characters.

Param SourcePort (string): the SourcePort param. String length must be between 1 and 1023 characters.
*/
type UdpObject struct {
	Override   *UdpOverrideObject `json:"override,omitempty"`
	Port       string             `json:"port"`
	SourcePort *string            `json:"source_port,omitempty"`
}

/*
UdpOverrideObject object.

ShortName:
Parent chains:
*
* protocol
* udp
* override

Args:

Param Timeout (int64): udp session timeout value (in second) Value must be between 1 and 604800. Default: `30`.
*/
type UdpOverrideObject struct {
	Timeout *int64 `json:"timeout,omitempty"`
}
