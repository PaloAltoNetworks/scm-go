package groups

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/application-groups

/*
Config object.

ShortName: qhDZEMT
Parent chains:
*

Args:

Param Id (string, read-only): UUID of the resource

Param Members ([]string, required): the Members param. Individual elements in this list are subject to additional validation. String length must not exceed 63 characters.

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.
*/
type Config struct {
	Id      *string  `json:"id,omitempty"`
	Members []string `json:"members"`
	Name    string   `json:"name"`
}
