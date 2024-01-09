package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/hip-profiles

/*
Config object.

ShortName: zGfKFAQ
Parent chains:
*

Args:

Param Description (string): the Description param. String length must not exceed 255 characters.

Param Id (string, read-only): UUID of the resource

Param Match (string, required): the Match param. String length must not exceed 2048 characters.

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.
*/
type Config struct {
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	Match       string  `json:"match"`
	Name        string  `json:"name"`
}
