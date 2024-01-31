package object

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/addresses

/*
Config object.

ShortName: aeWshcf
Parent chains:
*

Args:

Param Description (string): the Description param. String length must not exceed 1023 characters.

Param Fqdn (string): the Fqdn param. String length must be between 1 and 255 characters. String validation regex: `^[a-zA-Z0-9_]([a-zA-Z0-9._-])+[a-zA-Z0-9]$`.

Param Id (string, read-only): UUID of the resource

Param IpNetmask (string): the IpNetmask param.

Param IpRange (string): the IpRange param.

Param IpWildcard (string): the IpWildcard param.

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 63 characters.

Param Tags ([]string): Tags for address object List must contain at most 64 elements. Individual elements in this list are subject to additional validation. String length must not exceed 127 characters.

Param Type (string, read-only): the Type param.

NOTE:  One of the following params should be specified:
  - IpNetmask
  - IpRange
  - IpWildcard
  - Fqdn
*/
type Config struct {
	Description *string  `json:"description,omitempty"`
	Fqdn        *string  `json:"fqdn,omitempty"`
	Id          *string  `json:"id,omitempty"`
	IpNetmask   *string  `json:"ip_netmask,omitempty"`
	IpRange     *string  `json:"ip_range,omitempty"`
	IpWildcard  *string  `json:"ip_wildcard,omitempty"`
	Name        string   `json:"name"`
	Tags        []string `json:"tag,omitempty"`
	Type        *string  `json:"type,omitempty"`
}
