package error

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/generic_error

import (
	zLXjrfn "github.com/paloaltonetworks/scm-go/netsec/schemas/error/detail/cause/infos"
)

/*
Config object.

ShortName: nnsRzDg
Parent chains:
*

Args:

Param Errors (zLXjrfn.Errors): the Errors param.

Param RequestId (string): the RequestId param.
*/
type Config struct {
	Errors    zLXjrfn.Errors `json:"_errors,omitempty"`
	RequestId *string        `json:"_request_id,omitempty"`
}
