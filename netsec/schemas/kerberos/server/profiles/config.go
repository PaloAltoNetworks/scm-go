package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/kerberos-server-profiles

/*
Config object.

ShortName: lNTtdgX
Parent chains:
*

Args:

Param Id (string, read-only): UUID of the resource

Param Servers ([]ServerObject, required): the Servers param.
*/
type Config struct {
	Id      *string        `json:"id,omitempty"`
	Servers []ServerObject `json:"server"`
}

/*
ServerObject object.

ShortName:
Parent chains:
*
* server
* _inline

Args:

Param Host (string): the Host param.

Param Name (string): the Name param.

Param Port (int64): the Port param. Value must be between 1 and 65535.
*/
type ServerObject struct {
	Host *string `json:"host,omitempty"`
	Name *string `json:"name,omitempty"`
	Port *int64  `json:"port,omitempty"`
}
