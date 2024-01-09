package groups

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/local-user-groups

/*
Config object.

ShortName: tOChZgD
Parent chains:
*

Args:

Param Name (string, required): the Name param. String length must not exceed 31 characters. String validation regex: `^[a-zA-Z0-9._-]+$`.

Param Users ([]string): the Users param.
*/
type Config struct {
	Name  string   `json:"name"`
	Users []string `json:"user,omitempty"`
}
