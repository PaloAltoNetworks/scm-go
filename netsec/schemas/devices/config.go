package devices

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/devices

/*
Config object.

ShortName: mhJDwSQ
Parent chains:
*

Args:

Param AntiVirusVersion (string): the AntiVirusVersion param.

Param AppReleaseDate (string): the AppReleaseDate param.

Param AppVersion (string): the AppVersion param.

Param AvReleaseDate (string): the AvReleaseDate param.

Param AvailableLicensess ([]AvailableLicenseObject): the AvailableLicensess param.

Param ConnectedSince (string): the ConnectedSince param.

Param DevCertDetail (string): the DevCertDetail param.

Param DevCertExpiryDate (string): the DevCertExpiryDate param.

Param Family (string): the Family param.

Param GpClientVerion (string): the GpClientVerion param.

Param GpDataVersion (string): the GpDataVersion param.

Param HaPeerSerial (string): the HaPeerSerial param.

Param HaPeerState (string): the HaPeerState param.

Param HaState (string): the HaState param.

Param Hostname (string): the Hostname param.

Param Id (string): the Id param.

Param InstalledLicenses ([]InstalledLicenseObject): the InstalledLicenses param.

Param IotReleaseDate (string): the IotReleaseDate param.

Param IotVersion (string): the IotVersion param.

Param IpAddress (string): the IpAddress param.

Param IpV6Address (string): the IpV6Address param.

Param IsConnected (bool): the IsConnected param.

Param LicenseMatch (bool): the LicenseMatch param.

Param LogDbVersion (string): the LogDbVersion param.

Param MacAddress (string): the MacAddress param.

Param Model (string): the Model param.

Param SoftwareVersion (string): the SoftwareVersion param.

Param ThreatReleaseDate (string): the ThreatReleaseDate param.

Param ThreatVersion (string): the ThreatVersion param.

Param Uptime (string): the Uptime param.

Param UrlDbType (string): the UrlDbType param.

Param UrlDbVer (string): the UrlDbVer param.

Param VmState (string): the VmState param.

Param WfReleaseDate (string): the WfReleaseDate param.

Param WfVer (string): the WfVer param.
*/
type Config struct {
	AntiVirusVersion   *string                  `json:"anti_virus_version,omitempty"`
	AppReleaseDate     *string                  `json:"app_release_date,omitempty"`
	AppVersion         *string                  `json:"app_version,omitempty"`
	AvReleaseDate      *string                  `json:"av_release_date,omitempty"`
	AvailableLicensess []AvailableLicenseObject `json:"available_licensess,omitempty"`
	ConnectedSince     *string                  `json:"connected_since,omitempty"`
	DevCertDetail      *string                  `json:"dev_cert_detail,omitempty"`
	DevCertExpiryDate  *string                  `json:"dev_cert_expiry_date,omitempty"`
	Family             *string                  `json:"family,omitempty"`
	GpClientVerion     *string                  `json:"gp_client_verion,omitempty"`
	GpDataVersion      *string                  `json:"gp_data_version,omitempty"`
	HaPeerSerial       *string                  `json:"ha_peer_serial,omitempty"`
	HaPeerState        *string                  `json:"ha_peer_state,omitempty"`
	HaState            *string                  `json:"ha_state,omitempty"`
	Hostname           *string                  `json:"hostname,omitempty"`
	Id                 *string                  `json:"id,omitempty"`
	InstalledLicenses  []InstalledLicenseObject `json:"installed_licenses,omitempty"`
	IotReleaseDate     *string                  `json:"iot_release_date,omitempty"`
	IotVersion         *string                  `json:"iot_version,omitempty"`
	IpAddress          *string                  `json:"ip_address,omitempty"`
	IpV6Address        *string                  `json:"ipV6_address,omitempty"`
	IsConnected        *bool                    `json:"is_connected,omitempty"`
	LicenseMatch       *bool                    `json:"license_match,omitempty"`
	LogDbVersion       *string                  `json:"log_db_version,omitempty"`
	MacAddress         *string                  `json:"mac_address,omitempty"`
	Model              *string                  `json:"model,omitempty"`
	SoftwareVersion    *string                  `json:"software_version,omitempty"`
	ThreatReleaseDate  *string                  `json:"threat_release_date,omitempty"`
	ThreatVersion      *string                  `json:"threat_version,omitempty"`
	Uptime             *string                  `json:"uptime,omitempty"`
	UrlDbType          *string                  `json:"url_db_type,omitempty"`
	UrlDbVer           *string                  `json:"url_db_ver,omitempty"`
	VmState            *string                  `json:"vm_state,omitempty"`
	WfReleaseDate      *string                  `json:"wf_release_date,omitempty"`
	WfVer              *string                  `json:"wf_ver,omitempty"`
}

/*
AvailableLicenseObject object.

ShortName:
Parent chains:
*
* available_licensess
* _inline

Args:

Param Authcode (string): the Authcode param.

Param Expires (string): the Expires param.

Param Feature (string): the Feature param.

Param Issued (string): the Issued param.
*/
type AvailableLicenseObject struct {
	Authcode *string `json:"authcode,omitempty"`
	Expires  *string `json:"expires,omitempty"`
	Feature  *string `json:"feature,omitempty"`
	Issued   *string `json:"issued,omitempty"`
}

/*
InstalledLicenseObject object.

ShortName:
Parent chains:
*
* installed_licenses
* _inline

Args:

Param Authcode (string): the Authcode param.

Param Expired (string): the Expired param.

Param Expires (string): the Expires param.

Param Feature (string): the Feature param.

Param Issued (string): the Issued param.
*/
type InstalledLicenseObject struct {
	Authcode *string `json:"authcode,omitempty"`
	Expired  *string `json:"expired,omitempty"`
	Expires  *string `json:"expires,omitempty"`
	Feature  *string `json:"feature,omitempty"`
	Issued   *string `json:"issued,omitempty"`
}
