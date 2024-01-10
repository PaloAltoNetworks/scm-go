package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/radius-server-profiles

/*
Config object.

ShortName: xlJKoIe
Parent chains:
*

Args:

Param Id (string, read-only): UUID of the resource

Param Protocol (ProtocolObject): the Protocol param.

Param Retries (int64): the Retries param. Value must be between 1 and 5.

Param Servers ([]ServerObject, required): the Servers param.

Param Timeout (int64): the Timeout param. Value must be between 1 and 120.
*/
type Config struct {
	Id       *string         `json:"id,omitempty"`
	Protocol *ProtocolObject `json:"protocol,omitempty"`
	Retries  *int64          `json:"retries,omitempty"`
	Servers  []ServerObject  `json:"server"`
	Timeout  *int64          `json:"timeout,omitempty"`
}

/*
ProtocolObject object.

ShortName:
Parent chains:
*
* protocol

Args:

Param Chap (any): the Chap param. Default: `false`.

Param EapTtlsWithPap (EapTtlsWithPapObject): the EapTtlsWithPap param.

Param Pap (any): the Pap param. Default: `false`.

Param PeapMschapV2 (PeapMschapV2Object): the PeapMschapV2 param.

Param PeapWithGtc (PeapWithGtcObject): the PeapWithGtc param.

NOTE:  One of the following params should be specified:
  - Chap
  - EapTtlsWithPap
  - Pap
  - PeapMschapV2
  - PeapWithGtc
*/
type ProtocolObject struct {
	Chap           any                   `json:"CHAP,omitempty"`
	EapTtlsWithPap *EapTtlsWithPapObject `json:"EAP_TTLS_with_PAP,omitempty"`
	Pap            any                   `json:"PAP,omitempty"`
	PeapMschapV2   *PeapMschapV2Object   `json:"PEAP_MSCHAPv2,omitempty"`
	PeapWithGtc    *PeapWithGtcObject    `json:"PEAP_with_GTC,omitempty"`
}

/*
EapTtlsWithPapObject object.

ShortName:
Parent chains:
*
* protocol
* EAP_TTLS_with_PAP

Args:

Param AnonOuterId (bool): the AnonOuterId param.

Param RadiusCertProfile (string): the RadiusCertProfile param.
*/
type EapTtlsWithPapObject struct {
	AnonOuterId       *bool   `json:"anon_outer_id,omitempty"`
	RadiusCertProfile *string `json:"radius_cert_profile,omitempty"`
}

/*
PeapMschapV2Object object.

ShortName:
Parent chains:
*
* protocol
* PEAP_MSCHAPv2

Args:

Param AllowPwdChange (bool): the AllowPwdChange param.

Param AnonOuterId (bool): the AnonOuterId param.

Param RadiusCertProfile (string): the RadiusCertProfile param.
*/
type PeapMschapV2Object struct {
	AllowPwdChange    *bool   `json:"allow_pwd_change,omitempty"`
	AnonOuterId       *bool   `json:"anon_outer_id,omitempty"`
	RadiusCertProfile *string `json:"radius_cert_profile,omitempty"`
}

/*
PeapWithGtcObject object.

ShortName:
Parent chains:
*
* protocol
* PEAP_with_GTC

Args:

Param AnonOuterId (bool): the AnonOuterId param.

Param RadiusCertProfile (string): the RadiusCertProfile param.
*/
type PeapWithGtcObject struct {
	AnonOuterId       *bool   `json:"anon_outer_id,omitempty"`
	RadiusCertProfile *string `json:"radius_cert_profile,omitempty"`
}

/*
ServerObject object.

ShortName:
Parent chains:
*
* server
* _inline

Args:

Param IpAddress (string): the IpAddress param.

Param Name (string): the Name param.

Param Port (int64): the Port param. Value must be between 1 and 65535.

Param Secret (string): the Secret param. String length must not exceed 64 characters.
*/
type ServerObject struct {
	IpAddress *string `json:"ip_address,omitempty"`
	Name      *string `json:"name,omitempty"`
	Port      *int64  `json:"port,omitempty"`
	Secret    *string `json:"secret,omitempty"`
}
