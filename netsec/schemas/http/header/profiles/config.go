package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/http-header-profiles

/*
Config object.

ShortName: ljnPEAA
Parent chains:
*

Args:

Param Description (string): the Description param.

Param HttpHeaderInsertions ([]HttpHeaderInsertionObject): the HttpHeaderInsertions param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): the Name param.
*/
type Config struct {
	Description          *string                     `json:"description,omitempty"`
	HttpHeaderInsertions []HttpHeaderInsertionObject `json:"http_header_insertion,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	Name                 string                      `json:"name"`
}

/*
HttpHeaderInsertionObject object.

ShortName:
Parent chains:
*
* http_header_insertion
* _inline

Args:

Param Name (string, required): the Name param.

Param Types ([]TypeObject, required): the Types param.
*/
type HttpHeaderInsertionObject struct {
	Name  string       `json:"name"`
	Types []TypeObject `json:"type"`
}

/*
TypeObject object.

ShortName:
Parent chains:
*
* http_header_insertion
* _inline
* type
* _inline

Args:

Param Domains ([]string, required): the Domains param.

Param Headers ([]HeaderObject, required): the Headers param.

Param Name (string, required): the Name param.
*/
type TypeObject struct {
	Domains []string       `json:"domains"`
	Headers []HeaderObject `json:"headers"`
	Name    string         `json:"name"`
}

/*
HeaderObject object.

ShortName:
Parent chains:
*
* http_header_insertion
* _inline
* type
* _inline
* headers
* _inline

Args:

Param Header (string, required): the Header param.

Param Log (bool): the Log param. Default: `false`.

Param Name (string, required): the Name param.

Param Value (string, required): the Value param.
*/
type HeaderObject struct {
	Header string `json:"header"`
	Log    *bool  `json:"log,omitempty"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}
