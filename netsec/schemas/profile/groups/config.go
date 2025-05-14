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

Param AiSecurities ([]string): List of AI security profiles.

Param DnsSecurities ([]string): List of DNS security profiles.

Param FileBlockings ([]string): List of file blocking profiles.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): The name of the profile group.

Param SaasSecurities ([]string): List of HTTP header insertion profiles.

Param Spywares ([]string): List of anti-spyware profiles.

Param UrlFilterings ([]string): List of URL filtering profiles.

Param VirusAndWildfireAnalyses ([]string): List of anti-virus and Wildfire analysis profiles.

Param Vulnerabilities ([]string): List of vulnerability protection profiles.
*/
type Config struct {
	AiSecurities             []string `json:"ai_security,omitempty"`
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
