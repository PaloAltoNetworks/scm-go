package labels

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/labels

/*
Config object.

ShortName: fOGwhfV
Parent chains:
*

Args:

Param Description (string): the Description param.

Param Id (string, read-only): the Id param.

Param Name (string, required): the Name param.
*/
type Config struct {
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
}
