package users

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/local-users

/*
Config object.

ShortName: suxdMuj
Parent chains:
*

Args:

Param Disabled (bool): the Disabled param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): the Name param. String length must not exceed 31 characters.

Param Password (string): the Password param. String length must not exceed 63 characters.
*/
type Config struct {
	Disabled *bool   `json:"disabled,omitempty"`
	Id       *string `json:"id,omitempty"`
	Name     string  `json:"name"`
	Password *string `json:"password,omitempty"`
}
