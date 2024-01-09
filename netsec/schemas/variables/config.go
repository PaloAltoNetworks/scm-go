package variables

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/variables

/*
Config object.

ShortName: aeWshcf
Parent chains:
*

Args:

Param Description (string): the Description param.

Param Id (string, read-only): UUID of the resource

Param Name (string): Alphanumeric string begin with letter: [0-9a-zA-Z._-] String length must not exceed 63 characters.

Param Overridden (bool, read-only): the Overridden param.

Param Type (string): the Type param. String must be one of these: `"percent"`, `"count"`, `"ip-netmask"`, `"zone"`, `"ip-range"`, `"ip-wildcard"`, `"device-priority"`, `"device-id"`, `"egress-max"`, `"as-number"`, `"fqdn"`, `"port"`, `"link-tag"`, `"group-id"`, `"rate"`, `"router-id"`, `"qos-profile"`, `"timer"`.

Param Value (string): value can accept either string or integer
*/
type Config struct {
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Overridden  *bool   `json:"overridden,omitempty"`
	Type        *string `json:"type,omitempty"`
	Value       *string `json:"value,omitempty"`
}
