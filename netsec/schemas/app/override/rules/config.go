package rules

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/app-override-rules

/*
Config object.

ShortName: qhDZEMT
Parent chains:
*

Args:

Param Application (string, required): the Application param.

Param Description (string): the Description param. String length must not exceed 1024 characters.

Param Destinations ([]string, required): the Destinations param.

Param Disabled (bool): the Disabled param. Default: `false`.

Param Froms ([]string, required): the Froms param.

Param GroupTag (string): the GroupTag param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): the Name param. String length must not exceed 63 characters. String validation regex: `^[a-zA-Z0-9._-]+$`.

Param NegateDestination (bool): the NegateDestination param. Default: `false`.

Param NegateSource (bool): the NegateSource param. Default: `false`.

Param Port (int64, required): the Port param. Value must be between 0 and 65535.

Param Protocol (string, required): the Protocol param. String must be one of these: `"tcp"`, `"udp"`.

Param Sources ([]string, required): the Sources param.

Param Tags ([]string): the Tags param.

Param Tos ([]string, required): the Tos param.
*/
type Config struct {
	Application       string   `json:"application"`
	Description       *string  `json:"description,omitempty"`
	Destinations      []string `json:"destination"`
	Disabled          *bool    `json:"disabled,omitempty"`
	Froms             []string `json:"from"`
	GroupTag          *string  `json:"group_tag,omitempty"`
	Id                *string  `json:"id,omitempty"`
	Name              string   `json:"name"`
	NegateDestination *bool    `json:"negate_destination,omitempty"`
	NegateSource      *bool    `json:"negate_source,omitempty"`
	Port              int64    `json:"port"`
	Protocol          string   `json:"protocol"`
	Sources           []string `json:"source"`
	Tags              []string `json:"tag,omitempty"`
	Tos               []string `json:"to"`
}
