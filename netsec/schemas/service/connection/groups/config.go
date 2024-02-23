package groups

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/service-connection-groups

/*
Config object.

ShortName: sPPuDTU
Parent chains:
*

Args:

Param DisableSnat (bool): the DisableSnat param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): the Name param.

Param PbfOnly (bool): the PbfOnly param.

Param Targets ([]string, required): the Targets param.
*/
type Config struct {
	DisableSnat *bool    `json:"disable_snat,omitempty"`
	Id          *string  `json:"id,omitempty"`
	Name        string   `json:"name"`
	PbfOnly     *bool    `json:"pbf_only,omitempty"`
	Targets     []string `json:"target"`
}
