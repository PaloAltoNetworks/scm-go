package servers

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/internal-dns-servers

/*
Config object.

ShortName: ljnPEAA
Parent chains:
*

Args:

Param DomainNames ([]string, required): the DomainNames param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): the Name param.

Param Primary (string, required): the Primary param.

Param Secondary (string): the Secondary param.
*/
type Config struct {
	DomainNames []string `json:"domain_name"`
	Id          *string  `json:"id,omitempty"`
	Name        string   `json:"name"`
	Primary     string   `json:"primary"`
	Secondary   *string  `json:"secondary,omitempty"`
}
