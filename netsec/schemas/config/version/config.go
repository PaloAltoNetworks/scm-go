package version

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/config-version

/*
Config object.

ShortName: qhaCvMo
Parent chains:
*

Args:

Param Admin (string, required): The administrator or service account that created this configuration version

Param Created (int64, required): the Created param.

Param Date (string, required): the Date param.

Param Deleted (int64, required): the Deleted param.

Param Description (string, required): the Description param.

Param EditedBy (string, required): the EditedBy param.

Param Id (int64, required): The configuration version

Param ImpactedDevices (string, required): the ImpactedDevices param.

Param NgfwScope (string, required): A comma separated list of firewall serial numbers

Param Scope (string, required): the Scope param.

Param SwgConfig (string, required): the SwgConfig param.

Param Types (string, required): the Types param.

Param Updated (int64, required): the Updated param.

Param Version (string, required): The configuration version name
*/
type Config struct {
	Admin           string `json:"admin"`
	Created         int64  `json:"created"`
	Date            string `json:"date"`
	Deleted         int64  `json:"deleted"`
	Description     string `json:"description"`
	EditedBy        string `json:"edited_by"`
	Id              int64  `json:"id"`
	ImpactedDevices string `json:"impacted_devices"`
	NgfwScope       string `json:"ngfw_scope"`
	Scope           string `json:"scope"`
	SwgConfig       string `json:"swg_config"`
	Types           string `json:"types"`
	Updated         int64  `json:"updated"`
	Version         string `json:"version"`
}
