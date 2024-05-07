package categories

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/url-categories

/*
Config object.

ShortName: cSTLpMW
Parent chains:
*

Args:

Param Description (string): the Description param.

Param Id (string, read-only): UUID of the resource

Param List ([]string): the List param.

Param Name (string, required): the Name param.

Param Type (string): the Type param. String must be one of these: `"URL List"`, `"Category Match"`. Default: `"URL List"`.
*/
type Config struct {
	Description *string  `json:"description,omitempty"`
	Id          *string  `json:"id,omitempty"`
	List        []string `json:"list,omitempty"`
	Name        string   `json:"name"`
	Type        *string  `json:"type,omitempty"`
}
