package payload

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/folder-post-payload

/*
Config object.

ShortName: gmuDWVF
Parent chains:
*

Args:

Param Description (string): the Description param.

Param Labels ([]string): the Labels param.

Param Name (string, required): the Name param.

Param TargetParent (string): the TargetParent param.
*/
type Config struct {
	Description  *string  `json:"description,omitempty"`
	Labels       []string `json:"labels,omitempty"`
	Name         string   `json:"name"`
	TargetParent *string  `json:"target_parent,omitempty"`
}
