package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/saml-server-profiles

/*
Config object.

ShortName: yjUyRcw
Parent chains:
*

Args:

Param Certificate (string, required): the Certificate param. String length must not exceed 63 characters.

Param EntityId (string): the EntityId param. String length must be between 1 and 1024 characters.

Param Id (string, read-only): UUID of the resource

Param MaxClockSkew (int64): the MaxClockSkew param. Value must be between 1 and 900.

Param SloBindings (string): the SloBindings param. String must be one of these: `"post"`, `"redirect"`.

Param SsoBindings (string): the SsoBindings param. String must be one of these: `"post"`, `"redirect"`.

Param SsoUrl (string): the SsoUrl param. String length must be between 1 and 255 characters.

Param ValidateIdpCertificate (bool): the ValidateIdpCertificate param.

Param WantAuthRequestsSigned (bool): the WantAuthRequestsSigned param.
*/
type Config struct {
	Certificate            string  `json:"certificate"`
	EntityId               *string `json:"entity_id,omitempty"`
	Id                     *string `json:"id,omitempty"`
	MaxClockSkew           *int64  `json:"max_clock_skew,omitempty"`
	SloBindings            *string `json:"slo_bindings,omitempty"`
	SsoBindings            *string `json:"sso_bindings,omitempty"`
	SsoUrl                 *string `json:"sso_url,omitempty"`
	ValidateIdpCertificate *bool   `json:"validate_idp_certificate,omitempty"`
	WantAuthRequestsSigned *bool   `json:"want_auth_requests_signed,omitempty"`
}
