package rules

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/decryption-rules

/*
Config object.

ShortName: tephihM
Parent chains:
*

Args:

Param Action (string, required): the Action param. String must be one of these: `"decrypt"`, `"no-decrypt"`.

Param Categories ([]string, required): the Categories param.

Param Description (string): the Description param.

Param DestinationHips ([]string): the DestinationHips param.

Param Destinations ([]string, required): the Destinations param.

Param Disabled (bool): the Disabled param.

Param Froms ([]string, required): the Froms param.

Param Id (string, read-only): UUID of the resource

Param LogFail (bool): the LogFail param.

Param LogSetting (string): the LogSetting param.

Param LogSuccess (bool): the LogSuccess param.

Param Name (string, required): the Name param.

Param NegateDestination (bool): the NegateDestination param.

Param NegateSource (bool): the NegateSource param.

Param Profile (string): the Profile param.

Param Services ([]string, required): the Services param.

Param SourceHips ([]string): the SourceHips param.

Param SourceUsers ([]string, required): the SourceUsers param.

Param Sources ([]string, required): the Sources param.

Param Tags ([]string): the Tags param.

Param Tos ([]string, required): the Tos param.

Param Type (TypeObject): the Type param.
*/
type Config struct {
	Action            string      `json:"action"`
	Categories        []string    `json:"category"`
	Description       *string     `json:"description,omitempty"`
	DestinationHips   []string    `json:"destination_hip,omitempty"`
	Destinations      []string    `json:"destination"`
	Disabled          *bool       `json:"disabled,omitempty"`
	Froms             []string    `json:"from"`
	Id                *string     `json:"id,omitempty"`
	LogFail           *bool       `json:"log_fail,omitempty"`
	LogSetting        *string     `json:"log_setting,omitempty"`
	LogSuccess        *bool       `json:"log_success,omitempty"`
	Name              string      `json:"name"`
	NegateDestination *bool       `json:"negate_destination,omitempty"`
	NegateSource      *bool       `json:"negate_source,omitempty"`
	Profile           *string     `json:"profile,omitempty"`
	Services          []string    `json:"service"`
	SourceHips        []string    `json:"source_hip,omitempty"`
	SourceUsers       []string    `json:"source_user"`
	Sources           []string    `json:"source"`
	Tags              []string    `json:"tag,omitempty"`
	Tos               []string    `json:"to"`
	Type              *TypeObject `json:"type,omitempty"`
}

/*
TypeObject object.

ShortName:
Parent chains:
*
* type

Args:

Param SslForwardProxy (any): the SslForwardProxy param.

Param SslInboundInspection (string): add the certificate name for SSL inbound inspection

NOTE:  One of the following params should be specified:
  - SslForwardProxy
  - SslInboundInspection
*/
type TypeObject struct {
	SslForwardProxy      any     `json:"ssl_forward_proxy,omitempty"`
	SslInboundInspection *string `json:"ssl_inbound_inspection,omitempty"`
}
