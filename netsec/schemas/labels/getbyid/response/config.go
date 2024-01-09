package response

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/labels-getbyid-response

/*
Config object.

ShortName: nodMrsO
Parent chains:
*

Args:

Param Description (string): the Description param.

Param Folders ([]string): the Folders param.

Param Id (string): the Id param.

Param Name (string, required): the Name param.

Param Snippets ([]string): the Snippets param.
*/
type Config struct {
	Description *string  `json:"description,omitempty"`
	Folders     []string `json:"folders,omitempty"`
	Id          *string  `json:"id,omitempty"`
	Name        string   `json:"name"`
	Snippets    []string `json:"snippets,omitempty"`
}
