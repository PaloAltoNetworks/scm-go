package mod

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/snippets-mod

/*
Config object.

ShortName: yJkkSzS
Parent chains:
*

Args:

Param Description (string): the Description param.

Param Id (string, read-only): the Id param.

Param Labels ([]string): the Labels param.

Param Name (string, required): the Name param.

Param Type (string, read-only): the Type param. String must be one of these: `"predefined"`, `"custom"`.
*/
type Config struct {
	Description *string  `json:"description,omitempty"`
	Id          *string  `json:"id,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Name        string   `json:"name"`
	Type        *string  `json:"type,omitempty"`
}
