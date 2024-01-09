package groups

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/profile-groups

/*
Config object.

ShortName: qhaCvMo
Parent chains:
*

Args:

Param DnsSecurities ([]string): the DnsSecurities param.

Param FileBlockings ([]string): the FileBlockings param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): the Name param.

Param SaasSecurities ([]string): the SaasSecurities param.

Param Spywares ([]string): the Spywares param.

Param UrlFilterings ([]string): the UrlFilterings param.

Param VirusAndWildfireAnalyses ([]string): the VirusAndWildfireAnalyses param.

Param Vulnerabilities ([]string): the Vulnerabilities param.
*/
type Config struct {
	DnsSecurities            []string `json:"dns_security,omitempty"`
	FileBlockings            []string `json:"file_blocking,omitempty"`
	Id                       *string  `json:"id,omitempty"`
	Name                     string   `json:"name"`
	SaasSecurities           []string `json:"saas_security,omitempty"`
	Spywares                 []string `json:"spyware,omitempty"`
	UrlFilterings            []string `json:"url_filtering,omitempty"`
	VirusAndWildfireAnalyses []string `json:"virus_and_wildfire_analysis,omitempty"`
	Vulnerabilities          []string `json:"vulnerability,omitempty"`
}
