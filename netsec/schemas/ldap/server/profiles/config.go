package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/ldap-server-profiles

/*
Config object.

ShortName: quvWIWu
Parent chains:
*

Args:

Param Base (string): the Base param. String length must not exceed 255 characters.

Param BindDn (string): the BindDn param. String length must not exceed 255 characters.

Param BindPassword (string): the BindPassword param. String length must not exceed 121 characters.

Param BindTimelimit (string): the BindTimelimit param.

Param Id (string, read-only): UUID of the resource

Param LdapType (string): the LdapType param. String must be one of these: `"active-directory"`, `"e-directory"`, `"sun"`, `"other"`.

Param RetryInterval (int64): the RetryInterval param.

Param Servers ([]ServerObject, required): the Servers param.

Param Ssl (bool): the Ssl param.

Param Timelimit (int64): the Timelimit param.

Param VerifyServerCertificate (bool): the VerifyServerCertificate param.
*/
type Config struct {
	Base                    *string        `json:"base,omitempty"`
	BindDn                  *string        `json:"bind_dn,omitempty"`
	BindPassword            *string        `json:"bind_password,omitempty"`
	BindTimelimit           *string        `json:"bind_timelimit,omitempty"`
	Id                      *string        `json:"id,omitempty"`
	LdapType                *string        `json:"ldap_type,omitempty"`
	RetryInterval           *int64         `json:"retry_interval,omitempty"`
	Servers                 []ServerObject `json:"server"`
	Ssl                     *bool          `json:"ssl,omitempty"`
	Timelimit               *int64         `json:"timelimit,omitempty"`
	VerifyServerCertificate *bool          `json:"verify_server_certificate,omitempty"`
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
*/
type ServerObject struct {
	Address *string `json:"address,omitempty"`
	Name    *string `json:"name,omitempty"`
	Port    *int64  `json:"port,omitempty"`
}
