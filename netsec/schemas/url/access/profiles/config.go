package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/url-access-profiles

/*
Config object.

ShortName: ivVDSwf
Parent chains:
*

Args:

Param Alerts ([]string): the Alerts param.

Param Allows ([]string): the Allows param.

Param Blocks ([]string): the Blocks param.

Param CloudInlineCat (bool): the CloudInlineCat param.

Param Continues ([]string): the Continues param.

Param CredentialEnforcement (CredentialEnforcementObject): the CredentialEnforcement param.

Param Description (string): the Description param. String length must not exceed 255 characters.

Param Id (string, read-only): UUID of the resource

Param LocalInlineCat (bool): the LocalInlineCat param.

Param LogContainerPageOnly (bool): the LogContainerPageOnly param. Default: `true`.

Param LogHttpHdrReferer (bool): the LogHttpHdrReferer param. Default: `false`.

Param LogHttpHdrUserAgent (bool): the LogHttpHdrUserAgent param. Default: `false`.

Param LogHttpHdrXff (bool): the LogHttpHdrXff param. Default: `false`.

Param MlavCategoryExceptions ([]string): the MlavCategoryExceptions param.

Param Name (string, required): the Name param.

Param SafeSearchEnforcement (bool): the SafeSearchEnforcement param. Default: `false`.
*/
type Config struct {
	Alerts                 []string                     `json:"alert,omitempty"`
	Allows                 []string                     `json:"allow,omitempty"`
	Blocks                 []string                     `json:"block,omitempty"`
	CloudInlineCat         *bool                        `json:"cloud_inline_cat,omitempty"`
	Continues              []string                     `json:"continue,omitempty"`
	CredentialEnforcement  *CredentialEnforcementObject `json:"credential_enforcement,omitempty"`
	Description            *string                      `json:"description,omitempty"`
	Id                     *string                      `json:"id,omitempty"`
	LocalInlineCat         *bool                        `json:"local_inline_cat,omitempty"`
	LogContainerPageOnly   *bool                        `json:"log_container_page_only,omitempty"`
	LogHttpHdrReferer      *bool                        `json:"log_http_hdr_referer,omitempty"`
	LogHttpHdrUserAgent    *bool                        `json:"log_http_hdr_user_agent,omitempty"`
	LogHttpHdrXff          *bool                        `json:"log_http_hdr_xff,omitempty"`
	MlavCategoryExceptions []string                     `json:"mlav_category_exception,omitempty"`
	Name                   string                       `json:"name"`
	SafeSearchEnforcement  *bool                        `json:"safe_search_enforcement,omitempty"`
}

/*
CredentialEnforcementObject object.

ShortName:
Parent chains:
*
* credential_enforcement

Args:

Param Alerts ([]string): the Alerts param.

Param Allows ([]string): the Allows param.

Param Blocks ([]string): the Blocks param.

Param Continues ([]string): the Continues param.

Param LogSeverity (string): the LogSeverity param. Default: `"medium"`.

Param Mode (ModeObject): the Mode param.
*/
type CredentialEnforcementObject struct {
	Alerts      []string    `json:"alert,omitempty"`
	Allows      []string    `json:"allow,omitempty"`
	Blocks      []string    `json:"block,omitempty"`
	Continues   []string    `json:"continue,omitempty"`
	LogSeverity *string     `json:"log_severity,omitempty"`
	Mode        *ModeObject `json:"mode,omitempty"`
}

/*
ModeObject object.

ShortName:
Parent chains:
*
* credential_enforcement
* mode

Args:

Param Disabled (any): the Disabled param.

Param DomainCredentials (any): the DomainCredentials param.

Param GroupMapping (string): the GroupMapping param.

Param IpUser (any): the IpUser param.

NOTE:  One of the following params should be specified:
  - Disabled
  - DomainCredentials
  - IpUser
  - GroupMapping
*/
type ModeObject struct {
	Disabled          any     `json:"disabled,omitempty"`
	DomainCredentials any     `json:"domain_credentials,omitempty"`
	GroupMapping      *string `json:"group_mapping,omitempty"`
	IpUser            any     `json:"ip_user,omitempty"`
}
