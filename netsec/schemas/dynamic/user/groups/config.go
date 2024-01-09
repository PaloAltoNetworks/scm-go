package groups

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/dynamic-user-groups

/*
Config object.

ShortName: ivVDSwf
Parent chains:
*

Args:

Param Description (string): the Description param. String length must not exceed 1023 characters.

Param Filter (string, required): tag-based filter String length must not exceed 2047 characters.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 63 characters.

Param Tags ([]string): Tags for dynamic user group object List must contain at most 64 elements.
*/
type Config struct {
	Description *string  `json:"description,omitempty"`
	Filter      string   `json:"filter"`
	Id          *string  `json:"id,omitempty"`
	Name        string   `json:"name"`
	Tags        []string `json:"tag,omitempty"`
}
