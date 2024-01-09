package devices

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/quarantined-devices

/*
Config object.

ShortName: gmuDWVF
Parent chains:
*

Args:

Param HostId (string, required): Device host id

Param SerialNumber (string): Device serial number
*/
type Config struct {
	HostId       string  `json:"host_id"`
	SerialNumber *string `json:"serial_number,omitempty"`
}
