package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/tacacs-server-profiles

/*
Config object.

ShortName: uAgupQd
Parent chains:
*

Args:

Param Id (string, read-only): UUID of the resource

Param Protocol (string, required): the Protocol param. String must be one of these: `"CHAP"`, `"PAP"`.

Param Servers ([]ServerObject, required): the Servers param.

Param Timeout (int64): the Timeout param. Value must be between 1 and 30.

Param UseSingleConnection (bool): the UseSingleConnection param.
*/
type Config struct {
	Id                  *string        `json:"id,omitempty"`
	Protocol            string         `json:"protocol"`
	Servers             []ServerObject `json:"server"`
	Timeout             *int64         `json:"timeout,omitempty"`
	UseSingleConnection *bool          `json:"use_single_connection,omitempty"`
}

/*
ServerObject object.

ShortName:
Parent chains:
*
* server
* _inline

Args:

Param Address (string): the Address param.

Param Name (string): the Name param.

Param Port (int64): the Port param. Value must be between 1 and 65535.

Param Secret (string): the Secret param. String length must not exceed 64 characters.
*/
type ServerObject struct {
	Address *string `json:"address,omitempty"`
	Name    *string `json:"name,omitempty"`
	Port    *int64  `json:"port,omitempty"`
	Secret  *string `json:"secret,omitempty"`
}
