package group

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/address-groups

/*
Config object.

ShortName: qhaCvMo
Parent chains:
*

Args:

Param Description (string): the Description param. String length must not exceed 1023 characters.

Param DynamicValue (DynamicValue): the DynamicValue param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 63 characters.

Param StaticList ([]string): the StaticList param. Individual elements in this list are subject to additional validation. String length must not exceed 63 characters.

Param Tags ([]string): Tags for address group object List must contain at most 64 elements. Individual elements in this list are subject to additional validation. String length must not exceed 127 characters.

NOTE:  One of the following params should be specified:
  - StaticList
  - DynamicValue
*/
type Config struct {
	Description  *string       `json:"description,omitempty"`
	DynamicValue *DynamicValue `json:"dynamic,omitempty"`
	Id           *string       `json:"id,omitempty"`
	Name         string        `json:"name"`
	StaticList   []string      `json:"static,omitempty"`
	Tags         []string      `json:"tag,omitempty"`
}

/*
DynamicValue object.

ShortName:
Parent chains:
*
* dynamic

Args:

Param Filter (string, required): Tag based filter defining group membership e.g. `tag1 AND tag2 OR tag3` String length must not exceed 2047 characters.
*/
type DynamicValue struct {
	Filter string `json:"filter"`
}
