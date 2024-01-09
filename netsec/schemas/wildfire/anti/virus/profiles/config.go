package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/wildfire-anti-virus-profiles

/*
Config object.

ShortName: fVAkWHS
Parent chains:
*

Args:

Param Description (string): the Description param.

Param Id (string, read-only): UUID of the resource

Param MlavExceptions ([]MlavExceptionObject): the MlavExceptions param.

Param Name (string, required): the Name param. String validation regex: `^[a-zA-Z0-9._-]+$`.

Param PacketCapture (bool): the PacketCapture param.

Param Rules ([]RuleObject): the Rules param.

Param ThreatExceptions ([]ThreatExceptionObject): the ThreatExceptions param.
*/
type Config struct {
	Description      *string                 `json:"description,omitempty"`
	Id               *string                 `json:"id,omitempty"`
	MlavExceptions   []MlavExceptionObject   `json:"mlav_exception,omitempty"`
	Name             string                  `json:"name"`
	PacketCapture    *bool                   `json:"packet_capture,omitempty"`
	Rules            []RuleObject            `json:"rules,omitempty"`
	ThreatExceptions []ThreatExceptionObject `json:"threat_exception,omitempty"`
}

/*
MlavExceptionObject object.

ShortName:
Parent chains:
*
* mlav_exception
* _inline

Args:

Param Description (string): the Description param.

Param Filename (string): the Filename param.

Param Name (string): the Name param.
*/
type MlavExceptionObject struct {
	Description *string `json:"description,omitempty"`
	Filename    *string `json:"filename,omitempty"`
	Name        *string `json:"name,omitempty"`
}

/*
RuleObject object.

ShortName:
Parent chains:
*
* rules
* _inline

Args:

Param Analysis (string): the Analysis param. String must be one of these: `"public-cloud"`, `"private-cloud"`.

Param Applications ([]string): the Applications param.

Param Direction (string): the Direction param. String must be one of these: `"download"`, `"upload"`, `"both"`.

Param FileTypes ([]string): the FileTypes param.

Param Name (string): the Name param.
*/
type RuleObject struct {
	Analysis     *string  `json:"analysis,omitempty"`
	Applications []string `json:"application,omitempty"`
	Direction    *string  `json:"direction,omitempty"`
	FileTypes    []string `json:"file_type,omitempty"`
	Name         *string  `json:"name,omitempty"`
}

/*
ThreatExceptionObject object.

ShortName:
Parent chains:
*
* threat_exception
* _inline

Args:

Param Name (string): the Name param.

Param Notes (string): the Notes param.
*/
type ThreatExceptionObject struct {
	Name  *string `json:"name,omitempty"`
	Notes *string `json:"notes,omitempty"`
}
