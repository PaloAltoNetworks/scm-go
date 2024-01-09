package actions

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/auto-tag-actions

/*
Config object.

ShortName: tephihM
Parent chains:
*

Args:

Param Actions ([]ActionsObject): the Actions param.

Param Description (string): the Description param. String length must not exceed 1023 characters.

Param Filter (string, required): Tag based filter defining group membership e.g. `tag1 AND tag2 OR tag3` String length must not exceed 2047 characters.

Param LogType (string, read-only, required): the LogType param.

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 63 characters.

Param Quarantine (bool): the Quarantine param.

Param SendToPanorama (bool): the SendToPanorama param.
*/
type Config struct {
	Actions        []ActionsObject `json:"actions,omitempty"`
	Description    *string         `json:"description,omitempty"`
	Filter         string          `json:"filter"`
	LogType        string          `json:"log_type"`
	Name           string          `json:"name"`
	Quarantine     *bool           `json:"quarantine,omitempty"`
	SendToPanorama *bool           `json:"send_to_panorama,omitempty"`
}

/*
ActionsObject object.

ShortName:
Parent chains:
*
* actions
* _inline

Args:

Param Name (string, required): the Name param.

Param Type (TypeObject, required): the Type param.
*/
type ActionsObject struct {
	Name string     `json:"name"`
	Type TypeObject `json:"type"`
}

/*
TypeObject object.

ShortName:
Parent chains:
*
* actions
* _inline
* type

Args:

Param Tagging (TaggingObject, required): the Tagging param.
*/
type TypeObject struct {
	Tagging TaggingObject `json:"tagging"`
}

/*
TaggingObject object.

ShortName:
Parent chains:
*
* actions
* _inline
* type
* tagging

Args:

Param Action (string, required): Add or Remove tag option String must be one of these: `"add-tag"`, `"remove-tag"`.

Param Tags ([]string): Tags for address object List must contain at most 64 elements.

Param Target (string, required): Source or Destination Address, User, X-Forwarded-For Address

Param Timeout (int64): the Timeout param.
*/
type TaggingObject struct {
	Action  string   `json:"action"`
	Tags    []string `json:"tags,omitempty"`
	Target  string   `json:"target"`
	Timeout *int64   `json:"timeout,omitempty"`
}
