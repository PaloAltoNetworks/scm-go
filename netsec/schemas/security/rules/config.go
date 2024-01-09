package rules

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/security-rules

/*
Config object.

ShortName: uvXdTvM
Parent chains:
*

Args:

Param Action (string, required): The action to be taken when the rule is matched String must be one of these: `"allow"`, `"deny"`, `"drop"`, `"reset-client"`, `"reset-server"`, `"reset-both"`.

Param Applications ([]string, required): The application(s) being accessed

Param Categories ([]string, required): The URL categories being accessed

Param Description (string): The description of the security rule

Param DestinationHips ([]string): The destination Host Integrity Profile(s)

Param Destinations ([]string, required): The destination address(es)

Param Disabled (bool): The state of the security rule Default: `false`.

Param Froms ([]string, required): The source security zone(s)

Param Id (string, read-only): The UUID of the security rule

Param LogSetting (string): The external log forwarding profile

Param Name (string, required): The name of the security rule

Param NegateDestination (bool): Negate the destination addresses(es) Default: `false`.

Param NegateSource (bool): Negate the source address(es) Default: `false`.

Param ProfileSetting (ProfileSettingObject): The security profile object

Param Services ([]string, required): The service(s) being accessed

Param SourceHips ([]string): The source Host Integrity Profile(s)

Param SourceUsers ([]string, required): The source user(s) or group(s)

Param Sources ([]string, required): The source address(es)

Param Tags ([]string): The tags associated with the security rule

Param Tos ([]string, required): The destination security zone(s)
*/
type Config struct {
	Action            string                `json:"action"`
	Applications      []string              `json:"application"`
	Categories        []string              `json:"category"`
	Description       *string               `json:"description,omitempty"`
	DestinationHips   []string              `json:"destination_hip,omitempty"`
	Destinations      []string              `json:"destination"`
	Disabled          *bool                 `json:"disabled,omitempty"`
	Froms             []string              `json:"from"`
	Id                *string               `json:"id,omitempty"`
	LogSetting        *string               `json:"log_setting,omitempty"`
	Name              string                `json:"name"`
	NegateDestination *bool                 `json:"negate_destination,omitempty"`
	NegateSource      *bool                 `json:"negate_source,omitempty"`
	ProfileSetting    *ProfileSettingObject `json:"profile_setting,omitempty"`
	Services          []string              `json:"service"`
	SourceHips        []string              `json:"source_hip,omitempty"`
	SourceUsers       []string              `json:"source_user"`
	Sources           []string              `json:"source"`
	Tags              []string              `json:"tag,omitempty"`
	Tos               []string              `json:"to"`
}

/*
ProfileSettingObject The security profile object

ShortName:
Parent chains:
*
* profile_setting

Args:

Param Group ([]string): The security profile group
*/
type ProfileSettingObject struct {
	Group []string `json:"group,omitempty"`
}
