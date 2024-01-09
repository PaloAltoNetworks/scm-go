package info

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/error_detail_cause_info

/*
Config object.

ShortName: eumQbRC
Parent chains:
*

Args:

Param Code (string): the Code param.

Param Details (any): the Details param.

Param Help (string): the Help param.

Param Message (string): the Message param.
*/
type Config struct {
	Code    *string `json:"code,omitempty"`
	Details any     `json:"details,omitempty"`
	Help    *string `json:"help,omitempty"`
	Message *string `json:"message,omitempty"`
}
