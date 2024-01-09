package servers

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/mfa-servers

/*
Config object.

ShortName: oPPPeKY
Parent chains:
*

Args:

Param Id (string, read-only): UUID of the resource

Param MfaCertProfile (string, required): the MfaCertProfile param.

Param MfaVendorType (MfaVendorTypeObject): the MfaVendorType param.

Param Name (string, required): the Name param.
*/
type Config struct {
	Id             *string              `json:"id,omitempty"`
	MfaCertProfile string               `json:"mfa_cert_profile"`
	MfaVendorType  *MfaVendorTypeObject `json:"mfa_vendor_type,omitempty"`
	Name           string               `json:"name"`
}

/*
MfaVendorTypeObject object.

ShortName:
Parent chains:
*
* mfa_vendor_type

Args:

Param DuoSecurityV2 (DuoSecurityV2Object): the DuoSecurityV2 param.

Param OktaAdaptiveV1 (OktaAdaptiveV1Object): the OktaAdaptiveV1 param.

Param PingIdentityV1 (PingIdentityV1Object): the PingIdentityV1 param.

Param RsaSecuridAccessV1 (RsaSecuridAccessV1Object): the RsaSecuridAccessV1 param.

NOTE:  One of the following params should be specified:
  - OktaAdaptiveV1
  - PingIdentityV1
  - RsaSecuridAccessV1
  - DuoSecurityV2
*/
type MfaVendorTypeObject struct {
	DuoSecurityV2      *DuoSecurityV2Object      `json:"duo_security_v2,omitempty"`
	OktaAdaptiveV1     *OktaAdaptiveV1Object     `json:"okta_adaptive_v1,omitempty"`
	PingIdentityV1     *PingIdentityV1Object     `json:"ping_identity_v1,omitempty"`
	RsaSecuridAccessV1 *RsaSecuridAccessV1Object `json:"rsa_securid_access_v1,omitempty"`
}

/*
DuoSecurityV2Object object.

ShortName:
Parent chains:
*
* mfa_vendor_type
* duo_security_v2

Args:

Param DuoApiHost (string): the DuoApiHost param.

Param DuoBaseuri (string): the DuoBaseuri param.

Param DuoIntegrationKey (string): the DuoIntegrationKey param.

Param DuoSecretKey (string): the DuoSecretKey param.

Param DuoTimeout (string): the DuoTimeout param.
*/
type DuoSecurityV2Object struct {
	DuoApiHost        *string `json:"duo_api_host,omitempty"`
	DuoBaseuri        *string `json:"duo_baseuri,omitempty"`
	DuoIntegrationKey *string `json:"duo_integration_key,omitempty"`
	DuoSecretKey      *string `json:"duo_secret_key,omitempty"`
	DuoTimeout        *string `json:"duo_timeout,omitempty"`
}

/*
OktaAdaptiveV1Object object.

ShortName:
Parent chains:
*
* mfa_vendor_type
* okta_adaptive_v1

Args:

Param OktaApiHost (string): the OktaApiHost param.

Param OktaBaseuri (string): the OktaBaseuri param.

Param OktaOrg (string): the OktaOrg param.

Param OktaTimeout (string): the OktaTimeout param.

Param OktaToken (string): the OktaToken param.
*/
type OktaAdaptiveV1Object struct {
	OktaApiHost *string `json:"okta_api_host,omitempty"`
	OktaBaseuri *string `json:"okta_baseuri,omitempty"`
	OktaOrg     *string `json:"okta_org,omitempty"`
	OktaTimeout *string `json:"okta_timeout,omitempty"`
	OktaToken   *string `json:"okta_token,omitempty"`
}

/*
PingIdentityV1Object object.

ShortName:
Parent chains:
*
* mfa_vendor_type
* ping_identity_v1

Args:

Param PingApiHost (string): the PingApiHost param.

Param PingBaseuri (string): the PingBaseuri param.

Param PingOrg (string): the PingOrg param.

Param PingOrgAlias (string): the PingOrgAlias param.

Param PingTimeout (string): the PingTimeout param.

Param PingToken (string): the PingToken param.
*/
type PingIdentityV1Object struct {
	PingApiHost  *string `json:"ping_api_host,omitempty"`
	PingBaseuri  *string `json:"ping_baseuri,omitempty"`
	PingOrg      *string `json:"ping_org,omitempty"`
	PingOrgAlias *string `json:"ping_org_alias,omitempty"`
	PingTimeout  *string `json:"ping_timeout,omitempty"`
	PingToken    *string `json:"ping_token,omitempty"`
}

/*
RsaSecuridAccessV1Object object.

ShortName:
Parent chains:
*
* mfa_vendor_type
* rsa_securid_access_v1

Args:

Param RsaAccessid (string): the RsaAccessid param.

Param RsaAccesskey (string): the RsaAccesskey param.

Param RsaApiHost (string): the RsaApiHost param.

Param RsaAssurancepolicyid (string): the RsaAssurancepolicyid param.

Param RsaBaseuri (string): the RsaBaseuri param.

Param RsaTimeout (string): the RsaTimeout param.
*/
type RsaSecuridAccessV1Object struct {
	RsaAccessid          *string `json:"rsa_accessid,omitempty"`
	RsaAccesskey         *string `json:"rsa_accesskey,omitempty"`
	RsaApiHost           *string `json:"rsa_api_host,omitempty"`
	RsaAssurancepolicyid *string `json:"rsa_assurancepolicyid,omitempty"`
	RsaBaseuri           *string `json:"rsa_baseuri,omitempty"`
	RsaTimeout           *string `json:"rsa_timeout,omitempty"`
}
