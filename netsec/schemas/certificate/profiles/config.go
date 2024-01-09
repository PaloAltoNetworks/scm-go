package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/certificate-profiles

/*
Config object.

ShortName: dEowMJz
Parent chains:
*

Args:

Param BlockExpiredCert (bool): the BlockExpiredCert param.

Param BlockTimeoutCert (bool): the BlockTimeoutCert param.

Param BlockUnauthenticatedCert (bool): the BlockUnauthenticatedCert param.

Param BlockUnknownCert (bool): the BlockUnknownCert param.

Param CaCertificates ([]CaCertificateObject, required): the CaCertificates param.

Param CertStatusTimeout (string): the CertStatusTimeout param.

Param CrlReceiveTimeout (string): the CrlReceiveTimeout param.

Param Domain (string): the Domain param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 63 characters.

Param OcspReceiveTimeout (string): the OcspReceiveTimeout param.

Param UseCrl (bool): the UseCrl param.

Param UseOcsp (bool): the UseOcsp param.

Param UsernameField (UsernameFieldObject): the UsernameField param.
*/
type Config struct {
	BlockExpiredCert         *bool                 `json:"block_expired_cert,omitempty"`
	BlockTimeoutCert         *bool                 `json:"block_timeout_cert,omitempty"`
	BlockUnauthenticatedCert *bool                 `json:"block_unauthenticated_cert,omitempty"`
	BlockUnknownCert         *bool                 `json:"block_unknown_cert,omitempty"`
	CaCertificates           []CaCertificateObject `json:"ca_certificates"`
	CertStatusTimeout        *string               `json:"cert_status_timeout,omitempty"`
	CrlReceiveTimeout        *string               `json:"crl_receive_timeout,omitempty"`
	Domain                   *string               `json:"domain,omitempty"`
	Id                       *string               `json:"id,omitempty"`
	Name                     string                `json:"name"`
	OcspReceiveTimeout       *string               `json:"ocsp_receive_timeout,omitempty"`
	UseCrl                   *bool                 `json:"use_crl,omitempty"`
	UseOcsp                  *bool                 `json:"use_ocsp,omitempty"`
	UsernameField            *UsernameFieldObject  `json:"username_field,omitempty"`
}

/*
CaCertificateObject object.

ShortName:
Parent chains:
*
* ca_certificates
* _inline

Args:

Param DefaultOcspUrl (string): the DefaultOcspUrl param.

Param Name (string): the Name param.

Param OcspVerifyCert (string): the OcspVerifyCert param.

Param TemplateName (string): the TemplateName param.
*/
type CaCertificateObject struct {
	DefaultOcspUrl *string `json:"default_ocsp_url,omitempty"`
	Name           *string `json:"name,omitempty"`
	OcspVerifyCert *string `json:"ocsp_verify_cert,omitempty"`
	TemplateName   *string `json:"template_name,omitempty"`
}

/*
UsernameFieldObject object.

ShortName:
Parent chains:
*
* username_field

Args:

Param Subject (string): the Subject param. String must be one of these: `"common-name"`.

Param SubjectAlt (string): the SubjectAlt param. String must be one of these: `"email"`.
*/
type UsernameFieldObject struct {
	Subject    *string `json:"subject,omitempty"`
	SubjectAlt *string `json:"subject_alt,omitempty"`
}
