package manual

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/folder-manual

/*
Config object.

ShortName: vnMMXJg
Parent chains:
*

Args:

Param Description (string): the Description param.

Param Id (string, read-only): the Id param.

Param Labels ([]string): the Labels param.

Param Name (string, required): the Name param.

Param Parent (string, required): the Parent param.

Param Snippets ([]string): the Snippets param.
*/
type Config struct {
	Description *string  `json:"description,omitempty"`
	Id          *string  `json:"id,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Name        string   `json:"name"`
	Parent      string   `json:"parent"`
	Snippets    []string `json:"snippets,omitempty"`
}
