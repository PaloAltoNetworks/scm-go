package rules

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/qos-policy-rules

/*
Config object.

ShortName: tephihM
Parent chains:
*

Args:

Param Action (ActionObject, required): the Action param.

Param Description (string): the Description param.

Param DscpTos (DscpTosObject): the DscpTos param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): the Name param.

Param Schedule (string): the Schedule param.
*/
type Config struct {
	Action      ActionObject   `json:"action"`
	Description *string        `json:"description,omitempty"`
	DscpTos     *DscpTosObject `json:"dscp_tos,omitempty"`
	Id          *string        `json:"id,omitempty"`
	Name        string         `json:"name"`
	Schedule    *string        `json:"schedule,omitempty"`
}

/*
ActionObject object.

ShortName:
Parent chains:
*
* action

Args:

Param Class (string): the Class param.
*/
type ActionObject struct {
	Class *string `json:"class,omitempty"`
}

/*
DscpTosObject object.

ShortName:
Parent chains:
*
* dscp_tos

Args:

Param Codepoints ([]CodepointObject): the Codepoints param.
*/
type DscpTosObject struct {
	Codepoints []CodepointObject `json:"codepoints,omitempty"`
}

/*
CodepointObject object.

ShortName:
Parent chains:
*
* dscp_tos
* codepoints
* _inline

Args:

Param Name (string): the Name param.

Param Type (TypeObject): the Type param.
*/
type CodepointObject struct {
	Name *string     `json:"name,omitempty"`
	Type *TypeObject `json:"type,omitempty"`
}

/*
TypeObject object.

ShortName:
Parent chains:
*
* dscp_tos
* codepoints
* _inline
* type

Args:

Param Af (AfObject): the Af param.

Param Cs (CsObject): the Cs param.

Param Custom (CustomObject): the Custom param.

Param Ef (any): the Ef param.

Param Tos (TosObject): the Tos param.

NOTE:  One of the following params should be specified:
  - Ef
  - Af
  - Cs
  - Tos
  - Custom
*/
type TypeObject struct {
	Af     *AfObject     `json:"af,omitempty"`
	Cs     *CsObject     `json:"cs,omitempty"`
	Custom *CustomObject `json:"custom,omitempty"`
	Ef     any           `json:"ef,omitempty"`
	Tos    *TosObject    `json:"tos,omitempty"`
}

/*
AfObject object.

ShortName:
Parent chains:
*
* dscp_tos
* codepoints
* _inline
* type
* af

Args:

Param Codepoint (string): the Codepoint param.
*/
type AfObject struct {
	Codepoint *string `json:"codepoint,omitempty"`
}

/*
CsObject object.

ShortName:
Parent chains:
*
* dscp_tos
* codepoints
* _inline
* type
* cs

Args:

Param Codepoint (string): the Codepoint param.
*/
type CsObject struct {
	Codepoint *string `json:"codepoint,omitempty"`
}

/*
CustomObject object.

ShortName:
Parent chains:
*
* dscp_tos
* codepoints
* _inline
* type
* custom

Args:

Param Codepoint (CodepointObject1): the Codepoint param.
*/
type CustomObject struct {
	Codepoint *CodepointObject1 `json:"codepoint,omitempty"`
}

/*
CodepointObject1 object.

ShortName:
Parent chains:
*
* dscp_tos
* codepoints
* _inline
* type
* custom
* codepoint

Args:

Param BinaryValue (string): the BinaryValue param.

Param CodepointName (string): the CodepointName param.
*/
type CodepointObject1 struct {
	BinaryValue   *string `json:"binary_value,omitempty"`
	CodepointName *string `json:"codepoint_name,omitempty"`
}

/*
TosObject object.

ShortName:
Parent chains:
*
* dscp_tos
* codepoints
* _inline
* type
* tos

Args:

Param Codepoint (string): the Codepoint param.
*/
type TosObject struct {
	Codepoint *string `json:"codepoint,omitempty"`
}
