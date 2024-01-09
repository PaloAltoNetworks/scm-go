package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/file-blocking-profiles

/*
Config object.

ShortName: hhIWLbI
Parent chains:
*

Args:

Param Description (string): the Description param.

Param Id (string, read-only): the Id param.

Param Name (string, required): the Name param.

Param Rules ([]RuleObject): the Rules param.
*/
type Config struct {
	Description *string      `json:"description,omitempty"`
	Id          *string      `json:"id,omitempty"`
	Name        string       `json:"name"`
	Rules       []RuleObject `json:"rules,omitempty"`
}

/*
RuleObject object.

ShortName:
Parent chains:
*
* rules
* _inline

Args:

Param Action (string, required): the Action param. String must be one of these: `"alert"`, `"block"`, `"continue"`. Default: `"alert"`.

Param Applications ([]string, required): the Applications param. List must contain at least 1 elements.

Param Direction (string, required): the Direction param. String must be one of these: `"download"`, `"upload"`, `"both"`. Default: `"both"`.

Param FileTypes ([]string, required): the FileTypes param. List must contain at least 1 elements.

Param Name (string, required): the Name param.
*/
type RuleObject struct {
	Action       string   `json:"action"`
	Applications []string `json:"application"`
	Direction    string   `json:"direction"`
	FileTypes    []string `json:"file_type"`
	Name         string   `json:"name"`
}
