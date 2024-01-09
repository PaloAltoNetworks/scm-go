package responders

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/ocsp-responders

/*
Config object.

ShortName: uewNibC
Parent chains:
*

Args:

Param HostName (string, required): the HostName param. String length must be between 1 and 255 characters.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): alphanumeric string [:0-9a-zA-Z._-] String length must not exceed 63 characters. String validation regex: `^[a-zA-Z0-9._-]+$`.
*/
type Config struct {
	HostName string  `json:"host_name"`
	Id       *string `json:"id,omitempty"`
	Name     string  `json:"name"`
}
