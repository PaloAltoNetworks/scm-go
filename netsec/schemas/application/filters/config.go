package filters

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/application-filters

/*
Config object.

ShortName: wugpput
Parent chains:
*

Args:

Param Categories ([]string): the Categories param. Individual elements in this list are subject to additional validation. String length must not exceed 128 characters.

Param Evasive (bool): only True is a valid value

Param ExcessiveBandwidthUse (bool): only True is a valid value

Param Excludes ([]string): the Excludes param. Individual elements in this list are subject to additional validation. String length must not exceed 63 characters.

Param HasKnownVulnerabilities (bool): only True is a valid value

Param Id (string, read-only): UUID of the resource

Param IsSaas (bool): only True is a valid value

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param NewAppid (bool): only True is a valid value

Param Pervasive (bool): only True is a valid value

Param ProneToMisuse (bool): only True is a valid value

Param Risks ([]int64): the Risks param. Individual elements in this list are subject to additional validation. Value must be between 1 and 5.

Param SaasCertifications ([]string): the SaasCertifications param. Individual elements in this list are subject to additional validation. String length must not exceed 32 characters.

Param SaasRisks ([]string): the SaasRisks param. Individual elements in this list are subject to additional validation. String length must not exceed 32 characters.

Param Subcategories ([]string): the Subcategories param. Individual elements in this list are subject to additional validation. String length must not exceed 128 characters.

Param Tagging (TaggingObject): the Tagging param.

Param Technologies ([]string): the Technologies param. Individual elements in this list are subject to additional validation. String length must not exceed 128 characters.

Param TransfersFiles (bool): only True is a valid value

Param TunnelsOtherApps (bool): only True is a valid value

Param UsedByMalware (bool): only True is a valid value
*/
type Config struct {
	Categories              []string       `json:"category,omitempty"`
	Evasive                 *bool          `json:"evasive,omitempty"`
	ExcessiveBandwidthUse   *bool          `json:"excessive_bandwidth_use,omitempty"`
	Excludes                []string       `json:"exclude,omitempty"`
	HasKnownVulnerabilities *bool          `json:"has_known_vulnerabilities,omitempty"`
	Id                      *string        `json:"id,omitempty"`
	IsSaas                  *bool          `json:"is_saas,omitempty"`
	Name                    string         `json:"name"`
	NewAppid                *bool          `json:"new_appid,omitempty"`
	Pervasive               *bool          `json:"pervasive,omitempty"`
	ProneToMisuse           *bool          `json:"prone_to_misuse,omitempty"`
	Risks                   []int64        `json:"risk,omitempty"`
	SaasCertifications      []string       `json:"saas_certifications,omitempty"`
	SaasRisks               []string       `json:"saas_risk,omitempty"`
	Subcategories           []string       `json:"subcategory,omitempty"`
	Tagging                 *TaggingObject `json:"tagging,omitempty"`
	Technologies            []string       `json:"technology,omitempty"`
	TransfersFiles          *bool          `json:"transfers_files,omitempty"`
	TunnelsOtherApps        *bool          `json:"tunnels_other_apps,omitempty"`
	UsedByMalware           *bool          `json:"used_by_malware,omitempty"`
}

/*
TaggingObject object.

ShortName:
Parent chains:
*
* tagging

Args:

Param NoTag (bool): the NoTag param.

Param Tags ([]string): the Tags param. Individual elements in this list are subject to additional validation. String length must not exceed 127 characters.

NOTE:  One of the following params should be specified:
  - NoTag
  - Tags
*/
type TaggingObject struct {
	NoTag *bool    `json:"no_tag,omitempty"`
	Tags  []string `json:"tag,omitempty"`
}
