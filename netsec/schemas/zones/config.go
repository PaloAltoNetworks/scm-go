package zones

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/zones

/*
Config object.

ShortName: impVGiJ
Parent chains:
*

Args:

Param Device (string): The device in which the resource is defined String length must not exceed 64 characters. String validation regex: `^[a-zA-Z\d-_\. ]+$`.

Param DeviceAcl (DeviceAclObject): the DeviceAcl param.

Param DosLogSetting (string): the DosLogSetting param.

Param DosProfile (string): the DosProfile param.

Param EnableDeviceIdentification (bool): the EnableDeviceIdentification param.

Param EnableUserIdentification (bool): the EnableUserIdentification param.

Param Folder (string): The folder in which the resource is defined String length must not exceed 64 characters. String validation regex: `^[a-zA-Z\d-_\. ]+$`.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string begin with letter: [0-9a-zA-Z._-] String length must not exceed 63 characters.

Param Network (NetworkObject): the Network param.

Param Snippet (string): The snippet in which the resource is defined String length must not exceed 64 characters. String validation regex: `^[a-zA-Z\d-_\. ]+$`.

Param UserAcl (UserAclObject): the UserAcl param.

NOTE:  One of the following params should be specified:
  - Folder
  - Snippet
  - Device
*/
type Config struct {
	Device                     *string          `json:"device,omitempty"`
	DeviceAcl                  *DeviceAclObject `json:"device_acl,omitempty"`
	DosLogSetting              *string          `json:"dos_log_setting,omitempty"`
	DosProfile                 *string          `json:"dos_profile,omitempty"`
	EnableDeviceIdentification *bool            `json:"enable_device_identification,omitempty"`
	EnableUserIdentification   *bool            `json:"enable_user_identification,omitempty"`
	Folder                     *string          `json:"folder,omitempty"`
	Id                         *string          `json:"id,omitempty"`
	Name                       string           `json:"name"`
	Network                    *NetworkObject   `json:"network,omitempty"`
	Snippet                    *string          `json:"snippet,omitempty"`
	UserAcl                    *UserAclObject   `json:"user_acl,omitempty"`
}

/*
DeviceAclObject object.

ShortName:
Parent chains:
*
* device_acl

Args:

Param ExcludeList ([]string): the ExcludeList param.

Param IncludeList ([]string): the IncludeList param.
*/
type DeviceAclObject struct {
	ExcludeList []string `json:"exclude_list,omitempty"`
	IncludeList []string `json:"include_list,omitempty"`
}

/*
NetworkObject object.

ShortName:
Parent chains:
*
* network

Args:

Param EnablePacketBufferProtection (bool): the EnablePacketBufferProtection param.

Param LogSetting (string): the LogSetting param.

Param ZoneProtectionProfile (string): the ZoneProtectionProfile param.
*/
type NetworkObject struct {
	EnablePacketBufferProtection *bool   `json:"enable_packet_buffer_protection,omitempty"`
	LogSetting                   *string `json:"log_setting,omitempty"`
	ZoneProtectionProfile        *string `json:"zone_protection_profile,omitempty"`
}

/*
UserAclObject object.

ShortName:
Parent chains:
*
* user_acl

Args:

Param ExcludeList ([]string): the ExcludeList param.

Param IncludeList ([]string): the IncludeList param.
*/
type UserAclObject struct {
	ExcludeList []string `json:"exclude_list,omitempty"`
	IncludeList []string `json:"include_list,omitempty"`
}
