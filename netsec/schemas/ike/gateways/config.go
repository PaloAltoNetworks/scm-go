package gateways

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/ike-gateways

/*
Config object.

ShortName: yJkkSzS
Parent chains:
*

Args:

Param Authentication (AuthenticationObject, required): the Authentication param.

Param Id (string, read-only): UUID of the resource

Param LocalId (LocalIdObject): the LocalId param.

Param Name (string, required): Alphanumeric string begin with letter: [0-9a-zA-Z._-] String length must not exceed 63 characters.

Param PeerAddress (PeerAddressObject, required): the PeerAddress param.

Param PeerId (PeerIdObject): the PeerId param.

Param Protocol (ProtocolObject, required): the Protocol param.

Param ProtocolCommon (ProtocolCommonObject): the ProtocolCommon param.
*/
type Config struct {
	Authentication AuthenticationObject  `json:"authentication"`
	Id             *string               `json:"id,omitempty"`
	LocalId        *LocalIdObject        `json:"local_id,omitempty"`
	Name           string                `json:"name"`
	PeerAddress    PeerAddressObject     `json:"peer_address"`
	PeerId         *PeerIdObject         `json:"peer_id,omitempty"`
	Protocol       ProtocolObject        `json:"protocol"`
	ProtocolCommon *ProtocolCommonObject `json:"protocol_common,omitempty"`
}

/*
AuthenticationObject object.

ShortName:
Parent chains:
*
* authentication

Args:

Param Certificate (CertificateObject): the Certificate param.

Param PreSharedKey (PreSharedKeyObject): the PreSharedKey param.

NOTE:  One of the following params should be specified:
  - PreSharedKey
  - Certificate
*/
type AuthenticationObject struct {
	Certificate  *CertificateObject  `json:"certificate,omitempty"`
	PreSharedKey *PreSharedKeyObject `json:"pre_shared_key,omitempty"`
}

/*
CertificateObject object.

ShortName:
Parent chains:
*
* authentication
* certificate

Args:

Param AllowIdPayloadMismatch (bool): the AllowIdPayloadMismatch param.

Param CertificateProfile (string): the CertificateProfile param.

Param LocalCertificate (LocalCertificateObject): the LocalCertificate param.

Param StrictValidationRevocation (bool): the StrictValidationRevocation param.

Param UseManagementAsSource (bool): the UseManagementAsSource param.
*/
type CertificateObject struct {
	AllowIdPayloadMismatch     *bool                   `json:"allow_id_payload_mismatch,omitempty"`
	CertificateProfile         *string                 `json:"certificate_profile,omitempty"`
	LocalCertificate           *LocalCertificateObject `json:"local_certificate,omitempty"`
	StrictValidationRevocation *bool                   `json:"strict_validation_revocation,omitempty"`
	UseManagementAsSource      *bool                   `json:"use_management_as_source,omitempty"`
}

/*
LocalCertificateObject object.

ShortName:
Parent chains:
*
* authentication
* certificate
* local_certificate

Args:

Param LocalCertificateName (string): the LocalCertificateName param.
*/
type LocalCertificateObject struct {
	LocalCertificateName *string `json:"local_certificate_name,omitempty"`
}

/*
PreSharedKeyObject object.

ShortName:
Parent chains:
*
* authentication
* pre_shared_key

Args:

Param Key (string): the Key param.
*/
type PreSharedKeyObject struct {
	Key *string `json:"key,omitempty"`
}

/*
LocalIdObject object.

ShortName:
Parent chains:
*
* local_id

Args:

Param Id (string): Local ID string String length must be between 1 and 1024 characters. String validation regex: `^(.+\@[a-zA-Z0-9.-]+)$|^([$a-zA-Z0-9_:.-]+)$|^(([[:xdigit:]][[:xdigit:]])+)$|^([a-zA-Z0-9.]+=(\\,|[^,])+[, ]+)*([a-zA-Z0-9.]+=(\\,|[^,])+)$`.

Param Type (string): the Type param.
*/
type LocalIdObject struct {
	Id   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

/*
PeerAddressObject object.

ShortName:
Parent chains:
*
* peer_address

Args:

Param DynamicAddress (any): the DynamicAddress param. Default: `true`.

Param Fqdn (string): peer gateway FQDN name String length must not exceed 255 characters.

Param Ip (string): peer gateway has static IP address

NOTE:  One of the following params should be specified:
  - Ip
  - Fqdn
  - DynamicAddress
*/
type PeerAddressObject struct {
	DynamicAddress any     `json:"dynamic,omitempty"`
	Fqdn           *string `json:"fqdn,omitempty"`
	Ip             *string `json:"ip,omitempty"`
}

/*
PeerIdObject object.

ShortName:
Parent chains:
*
* peer_id

Args:

Param Id (string): Peer ID string String length must be between 1 and 1024 characters. String validation regex: `^(.+\@[\*a-zA-Z0-9.-]+)$|^([\*$a-zA-Z0-9_:.-]+)$|^(([[:xdigit:]][[:xdigit:]])+)$|^([a-zA-Z0-9.]+=(\\,|[^,])+[, ]+)*([a-zA-Z0-9.]+=(\\,|[^,])+)$`.

Param Type (string): the Type param. String must be one of these: `"ipaddr"`, `"keyid"`, `"fqdn"`, `"ufqdn"`.
*/
type PeerIdObject struct {
	Id   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

/*
ProtocolObject object.

ShortName:
Parent chains:
*
* protocol

Args:

Param Ikev1 (Ikev1Object): the Ikev1 param.

Param Ikev2 (Ikev2Object): the Ikev2 param.

Param Version (string): the Version param. String must be one of these: `"ikev2-preferred"`, `"ikev1"`, `"ikev2"`. Default: `"ikev2-preferred"`.
*/
type ProtocolObject struct {
	Ikev1   *Ikev1Object `json:"ikev1,omitempty"`
	Ikev2   *Ikev2Object `json:"ikev2,omitempty"`
	Version *string      `json:"version,omitempty"`
}

/*
Ikev1Object object.

ShortName:
Parent chains:
*
* protocol
* ikev1

Args:

Param Dpd (Ikev1DpdObject): the Dpd param.

Param IkeCryptoProfile (string): the IkeCryptoProfile param.
*/
type Ikev1Object struct {
	Dpd              *Ikev1DpdObject `json:"dpd,omitempty"`
	IkeCryptoProfile *string         `json:"ike_crypto_profile,omitempty"`
}

/*
Ikev1DpdObject object.

ShortName:
Parent chains:
*
* protocol
* ikev1
* dpd

Args:

Param Enable (bool): the Enable param.
*/
type Ikev1DpdObject struct {
	Enable *bool `json:"enable,omitempty"`
}

/*
Ikev2Object object.

ShortName:
Parent chains:
*
* protocol
* ikev2

Args:

Param Dpd (Ikev2DpdObject): the Dpd param.

Param IkeCryptoProfile (string): the IkeCryptoProfile param.
*/
type Ikev2Object struct {
	Dpd              *Ikev2DpdObject `json:"dpd,omitempty"`
	IkeCryptoProfile *string         `json:"ike_crypto_profile,omitempty"`
}

/*
Ikev2DpdObject object.

ShortName:
Parent chains:
*
* protocol
* ikev2
* dpd

Args:

Param Enable (bool): the Enable param.
*/
type Ikev2DpdObject struct {
	Enable *bool `json:"enable,omitempty"`
}

/*
ProtocolCommonObject object.

ShortName:
Parent chains:
*
* protocol_common

Args:

Param Fragmentation (FragmentationObject): the Fragmentation param.

Param NatTraversal (NatTraversalObject): the NatTraversal param.

Param PassiveMode (bool): the PassiveMode param.
*/
type ProtocolCommonObject struct {
	Fragmentation *FragmentationObject `json:"fragmentation,omitempty"`
	NatTraversal  *NatTraversalObject  `json:"nat_traversal,omitempty"`
	PassiveMode   *bool                `json:"passive_mode,omitempty"`
}

/*
FragmentationObject object.

ShortName:
Parent chains:
*
* protocol_common
* fragmentation

Args:

Param Enable (bool): the Enable param. Default: `false`.
*/
type FragmentationObject struct {
	Enable *bool `json:"enable,omitempty"`
}

/*
NatTraversalObject object.

ShortName:
Parent chains:
*
* protocol_common
* nat_traversal

Args:

Param Enable (bool): the Enable param.
*/
type NatTraversalObject struct {
	Enable *bool `json:"enable,omitempty"`
}
