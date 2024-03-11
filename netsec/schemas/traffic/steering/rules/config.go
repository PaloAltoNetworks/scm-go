package rules

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/traffic-steering-rules

/*
Config object.

ShortName: hhIWLbI
Parent chains:
*

Args:

Param Action (ActionObject): the Action param.

Param Category ([]string): the Category param.

Param Destination ([]string): the Destination param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): the Name param.

Param Service ([]string, required): the Service param.

Param Source ([]string, required): the Source param.

Param SourceUser ([]string): the SourceUser param.
*/
type Config struct {
	Action      *ActionObject `json:"action,omitempty"`
	Category    []string      `json:"category,omitempty"`
	Destination []string      `json:"destination,omitempty"`
	Id          *string       `json:"id,omitempty"`
	Name        string        `json:"name"`
	Service     []string      `json:"service"`
	Source      []string      `json:"source"`
	SourceUser  []string      `json:"source_user,omitempty"`
}

/*
ActionObject object.

ShortName:
Parent chains:
*
* action

Args:

Param Forward (ForwardObject): the Forward param.

Param NoPbf (any): the NoPbf param.

NOTE:  One of the following params should be specified:
  - Forward
  - NoPbf
*/
type ActionObject struct {
	Forward *ForwardObject `json:"forward,omitempty"`
	NoPbf   any            `json:"no-pbf,omitempty"`
}

/*
ForwardObject object.

ShortName:
Parent chains:
*
* action
* forward

Args:

Param Target (string): the Target param.
*/
type ForwardObject struct {
	Target *string `json:"target,omitempty"`
}
