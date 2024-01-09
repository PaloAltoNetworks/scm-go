package response

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/folder-get-response

/*
Config object.

ShortName: wugpput
Parent chains:
*

Args:

Param Children ([]any): children field is recursive, can contain another folder hierarchy.

Param Description (string): the Description param.

Param DisplayName (string): only available if container is of on-prem/cloud type

Param Id (string, required): the Id param.

Param Labels ([]string): the Labels param.

Param Model (string): only available if container is of on-prem type

Param Name (string, required): the Name param.

Param SerialNumber (string): only available if container is of on-prem type

Param Snippets ([]string): the Snippets param.

Param Type (string, required): the Type param. String must be one of these: `"container"`, `"on-prem"`, `"cloud"`.
*/
type Config struct {
	Children     []any    `json:"children,omitempty"`
	Description  *string  `json:"description,omitempty"`
	DisplayName  *string  `json:"display_name,omitempty"`
	Id           string   `json:"id"`
	Labels       []string `json:"labels,omitempty"`
	Model        *string  `json:"model,omitempty"`
	Name         string   `json:"name"`
	SerialNumber *string  `json:"serial_number,omitempty"`
	Snippets     []string `json:"snippets,omitempty"`
	Type         string   `json:"type"`
}
