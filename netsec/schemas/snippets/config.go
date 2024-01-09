package snippets

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/snippets

/*
Config object.

ShortName: ujXZojh
Parent chains:
*

Args:

Param CreatedIn (string, read-only): the CreatedIn param.

Param Description (string): the Description param.

Param DisplayName (string, read-only): the DisplayName param.

Param Folders ([]FolderObject, read-only): the Folders param.

Param Id (string, read-only): the Id param.

Param Labels ([]string): the Labels param.

Param LastUpdate (string, read-only): the LastUpdate param.

Param Name (string, required): the Name param.

Param SharedIn (string, read-only): the SharedIn param.

Param Type (string, read-only): the Type param. String must be one of these: `"predefined"`.
*/
type Config struct {
	CreatedIn   *string        `json:"created_in,omitempty"`
	Description *string        `json:"description,omitempty"`
	DisplayName *string        `json:"display_name,omitempty"`
	Folders     []FolderObject `json:"folders,omitempty"`
	Id          *string        `json:"id,omitempty"`
	Labels      []string       `json:"labels,omitempty"`
	LastUpdate  *string        `json:"last_update,omitempty"`
	Name        string         `json:"name"`
	SharedIn    *string        `json:"shared_in,omitempty"`
	Type        *string        `json:"type,omitempty"`
}

/*
FolderObject object.

ShortName:
Parent chains:
*
* folders
* _inline

Args:

Param Id (string): the Id param.

Param Name (string): the Name param.
*/
type FolderObject struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
