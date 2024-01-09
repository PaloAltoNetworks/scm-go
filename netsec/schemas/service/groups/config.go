package groups

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/service-groups

/*
Config object.

ShortName: uvXdTvM
Parent chains:
*

Args:

Param Id (string, read-only): UUID of the resource

Param Members ([]string, required): the Members param.

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 63 characters.

Param Tags ([]string): Tags for service group object List must contain at most 64 elements.
*/
type Config struct {
	Id      *string  `json:"id,omitempty"`
	Members []string `json:"members"`
	Name    string   `json:"name"`
	Tags    []string `json:"tag,omitempty"`
}
