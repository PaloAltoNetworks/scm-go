package versions

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/running-versions

/*
Config object.

ShortName: tephihM
Parent chains:
*

Args:

Param Date (string, required): The timestamp of when the configuration version was pushed to the folder or firewall

Param Device (string, required): The folder name or firewall serial number

Param Version (int64, required): The configuration version number
*/
type Config struct {
	Date    string `json:"date"`
	Device  string `json:"device"`
	Version int64  `json:"version"`
}
