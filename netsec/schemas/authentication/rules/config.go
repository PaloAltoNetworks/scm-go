package rules

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/authentication-rules

/*
Config object.

ShortName: ljnPEAA
Parent chains:
*

Args:

Param AuthenticationEnforcement (string): the authentication profile name to apply to authentication rule

Param Categories ([]string): the Categories param.

Param Description (string): the Description param.

Param DestinationHips ([]string): the DestinationHips param.

Param Destinations ([]string): the Destinations param.

Param Disabled (bool): the Disabled param. Default: `false`.

Param Froms ([]string): the Froms param.

Param GroupTag (string): the GroupTag param.

Param HipProfiles ([]string): the HipProfiles param.

Param Id (string, read-only): UUID of the resource

Param LogAuthenticationTimeout (bool): the LogAuthenticationTimeout param. Default: `false`.

Param LogSetting (string): the LogSetting param.

Param Name (string): the Name param.

Param NegateDestination (bool): the NegateDestination param. Default: `false`.

Param NegateSource (bool): the NegateSource param. Default: `false`.

Param Services ([]string): the Services param.

Param SourceHips ([]string): the SourceHips param.

Param SourceUsers ([]string): the SourceUsers param.

Param Sources ([]string): the Sources param.

Param Tags ([]string): the Tags param.

Param Timeout (int64): the Timeout param. Value must be between 1 and 1440.

Param Tos ([]string): the Tos param.
*/
type Config struct {
	AuthenticationEnforcement *string  `json:"authentication_enforcement,omitempty"`
	Categories                []string `json:"category,omitempty"`
	Description               *string  `json:"description,omitempty"`
	DestinationHips           []string `json:"destination_hip,omitempty"`
	Destinations              []string `json:"destination,omitempty"`
	Disabled                  *bool    `json:"disabled,omitempty"`
	Froms                     []string `json:"from,omitempty"`
	GroupTag                  *string  `json:"group_tag,omitempty"`
	HipProfiles               []string `json:"hip_profiles,omitempty"`
	Id                        *string  `json:"id,omitempty"`
	LogAuthenticationTimeout  *bool    `json:"log_authentication_timeout,omitempty"`
	LogSetting                *string  `json:"log_setting,omitempty"`
	Name                      *string  `json:"name,omitempty"`
	NegateDestination         *bool    `json:"negate_destination,omitempty"`
	NegateSource              *bool    `json:"negate_source,omitempty"`
	Services                  []string `json:"service,omitempty"`
	SourceHips                []string `json:"source_hip,omitempty"`
	SourceUsers               []string `json:"source_user,omitempty"`
	Sources                   []string `json:"source,omitempty"`
	Tags                      []string `json:"tag,omitempty"`
	Timeout                   *int64   `json:"timeout,omitempty"`
	Tos                       []string `json:"to,omitempty"`
}
