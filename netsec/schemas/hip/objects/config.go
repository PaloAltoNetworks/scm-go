package objects

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/hip-objects

/*
Config object.

ShortName: suxdMuj
Parent chains:
*

Args:

Param AntiMalware (AntiMalwareObject): the AntiMalware param.

Param Certificate (CertificateObject): the Certificate param.

Param CustomChecks (CustomChecksObject): the CustomChecks param.

Param DataLossPrevention (DataLossPreventionObject): the DataLossPrevention param.

Param Description (string): the Description param. String length must not exceed 255 characters.

Param DiskBackup (DiskBackupObject): the DiskBackup param.

Param DiskEncryption (DiskEncryptionObject): the DiskEncryption param.

Param Firewall (FirewallObject): the Firewall param.

Param HostInfo (HostInfoObject): the HostInfo param.

Param Id (string, read-only): UUID of the resource

Param MobileDevice (MobileDeviceObject): the MobileDevice param.

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param NetworkInfo (NetworkInfoObject): the NetworkInfo param.

Param PatchManagement (PatchManagementObject): the PatchManagement param.
*/
type Config struct {
	AntiMalware        *AntiMalwareObject        `json:"anti_malware,omitempty"`
	Certificate        *CertificateObject        `json:"certificate,omitempty"`
	CustomChecks       *CustomChecksObject       `json:"custom_checks,omitempty"`
	DataLossPrevention *DataLossPreventionObject `json:"data_loss_prevention,omitempty"`
	Description        *string                   `json:"description,omitempty"`
	DiskBackup         *DiskBackupObject         `json:"disk_backup,omitempty"`
	DiskEncryption     *DiskEncryptionObject     `json:"disk_encryption,omitempty"`
	Firewall           *FirewallObject           `json:"firewall,omitempty"`
	HostInfo           *HostInfoObject           `json:"host_info,omitempty"`
	Id                 *string                   `json:"id,omitempty"`
	MobileDevice       *MobileDeviceObject       `json:"mobile_device,omitempty"`
	Name               string                    `json:"name"`
	NetworkInfo        *NetworkInfoObject        `json:"network_info,omitempty"`
	PatchManagement    *PatchManagementObject    `json:"patch_management,omitempty"`
}

/*
AntiMalwareObject object.

ShortName:
Parent chains:
*
* anti_malware

Args:

Param Criteria (AntiMalwareCriteriaObject): the Criteria param.

Param ExcludeVendor (bool): the ExcludeVendor param. Default: `false`.

Param Vendors ([]AntiMalwareVendorObject): Vendor name
*/
type AntiMalwareObject struct {
	Criteria      *AntiMalwareCriteriaObject `json:"criteria,omitempty"`
	ExcludeVendor *bool                      `json:"exclude_vendor,omitempty"`
	Vendors       []AntiMalwareVendorObject  `json:"vendor,omitempty"`
}

/*
AntiMalwareCriteriaObject object.

ShortName:
Parent chains:
*
* anti_malware
* criteria

Args:

Param IsInstalled (bool): Is Installed Default: `true`.

Param LastScanTime (LastScanTimeObject): the LastScanTime param.

Param ProductVersion (ProductVersionObject): the ProductVersion param.

Param RealTimeProtection (string): real time protection String must be one of these: `"no"`, `"yes"`, `"not-available"`.

Param VirdefVersion (VirdefVersionObject): the VirdefVersion param.
*/
type AntiMalwareCriteriaObject struct {
	IsInstalled        *bool                 `json:"is_installed,omitempty"`
	LastScanTime       *LastScanTimeObject   `json:"last_scan_time,omitempty"`
	ProductVersion     *ProductVersionObject `json:"product_version,omitempty"`
	RealTimeProtection *string               `json:"real_time_protection,omitempty"`
	VirdefVersion      *VirdefVersionObject  `json:"virdef_version,omitempty"`
}

/*
LastScanTimeObject object.

ShortName:
Parent chains:
*
* anti_malware
* criteria
* last_scan_time

Args:

Param NotAvailable (any): the NotAvailable param.

Param NotWithin (LastScanTimeNotWithinObject): the NotWithin param.

Param Within (LastScanTimeWithinObject): the Within param.

NOTE:  One of the following params should be specified:
  - NotAvailable
  - Within
  - NotWithin
*/
type LastScanTimeObject struct {
	NotAvailable any                          `json:"not_available,omitempty"`
	NotWithin    *LastScanTimeNotWithinObject `json:"not_within,omitempty"`
	Within       *LastScanTimeWithinObject    `json:"within,omitempty"`
}

/*
LastScanTimeNotWithinObject object.

ShortName:
Parent chains:
*
* anti_malware
* criteria
* last_scan_time
* not_within

Args:

Param Days (int64): specify time in days Value must be between 1 and 65535. Default: `1`.

Param Hours (int64): specify time in hours Value must be between 1 and 65535. Default: `24`.

NOTE:  One of the following params should be specified:
  - Days
  - Hours
*/
type LastScanTimeNotWithinObject struct {
	Days  *int64 `json:"days,omitempty"`
	Hours *int64 `json:"hours,omitempty"`
}

/*
LastScanTimeWithinObject object.

ShortName:
Parent chains:
*
* anti_malware
* criteria
* last_scan_time
* within

Args:

Param Days (int64): specify time in days Value must be between 1 and 65535. Default: `1`.

Param Hours (int64): specify time in hours Value must be between 1 and 65535. Default: `24`.

NOTE:  One of the following params should be specified:
  - Days
  - Hours
*/
type LastScanTimeWithinObject struct {
	Days  *int64 `json:"days,omitempty"`
	Hours *int64 `json:"hours,omitempty"`
}

/*
ProductVersionObject object.

ShortName:
Parent chains:
*
* anti_malware
* criteria
* product_version

Args:

Param Contains (string): the Contains param. String length must not exceed 255 characters.

Param GreaterEqual (string): the GreaterEqual param. String length must not exceed 255 characters.

Param GreaterThan (string): the GreaterThan param. String length must not exceed 255 characters.

Param Is (string): the Is param. String length must not exceed 255 characters.

Param IsNot (string): the IsNot param. String length must not exceed 255 characters.

Param LessEqual (string): the LessEqual param. String length must not exceed 255 characters.

Param LessThan (string): the LessThan param. String length must not exceed 255 characters.

Param NotWithin (ProductVersionNotWithinObject): the NotWithin param.

Param Within (ProductVersionWithinObject): the Within param.

NOTE:  One of the following params should be specified:
  - GreaterEqual
  - GreaterThan
  - Is
  - IsNot
  - LessEqual
  - LessThan
  - Contains
  - Within
  - NotWithin
*/
type ProductVersionObject struct {
	Contains     *string                        `json:"contains,omitempty"`
	GreaterEqual *string                        `json:"greater_equal,omitempty"`
	GreaterThan  *string                        `json:"greater_than,omitempty"`
	Is           *string                        `json:"is,omitempty"`
	IsNot        *string                        `json:"is_not,omitempty"`
	LessEqual    *string                        `json:"less_equal,omitempty"`
	LessThan     *string                        `json:"less_than,omitempty"`
	NotWithin    *ProductVersionNotWithinObject `json:"not_within,omitempty"`
	Within       *ProductVersionWithinObject    `json:"within,omitempty"`
}

/*
ProductVersionNotWithinObject object.

ShortName:
Parent chains:
*
* anti_malware
* criteria
* product_version
* not_within

Args:

Param Versions (int64, required): versions range Value must be between 1 and 65535. Default: `1`.
*/
type ProductVersionNotWithinObject struct {
	Versions int64 `json:"versions"`
}

/*
ProductVersionWithinObject object.

ShortName:
Parent chains:
*
* anti_malware
* criteria
* product_version
* within

Args:

Param Versions (int64, required): versions range Value must be between 1 and 65535. Default: `1`.
*/
type ProductVersionWithinObject struct {
	Versions int64 `json:"versions"`
}

/*
VirdefVersionObject object.

ShortName:
Parent chains:
*
* anti_malware
* criteria
* virdef_version

Args:

Param NotWithin (VirdefVersionNotWithinObject): the NotWithin param.

Param Within (VirdefVersionWithinObject): the Within param.

NOTE:  One of the following params should be specified:
  - Within
  - NotWithin
*/
type VirdefVersionObject struct {
	NotWithin *VirdefVersionNotWithinObject `json:"not_within,omitempty"`
	Within    *VirdefVersionWithinObject    `json:"within,omitempty"`
}

/*
VirdefVersionNotWithinObject object.

ShortName:
Parent chains:
*
* anti_malware
* criteria
* virdef_version
* not_within

Args:

Param Days (int64): specify time in days Value must be between 1 and 65535. Default: `1`.

Param Versions (int64): specify versions range Value must be between 1 and 65535. Default: `1`.

NOTE:  One of the following params should be specified:
  - Days
  - Versions
*/
type VirdefVersionNotWithinObject struct {
	Days     *int64 `json:"days,omitempty"`
	Versions *int64 `json:"versions,omitempty"`
}

/*
VirdefVersionWithinObject object.

ShortName:
Parent chains:
*
* anti_malware
* criteria
* virdef_version
* within

Args:

Param Days (int64): specify time in days Value must be between 1 and 65535. Default: `1`.

Param Versions (int64): specify versions range Value must be between 1 and 65535. Default: `1`.

NOTE:  One of the following params should be specified:
  - Days
  - Versions
*/
type VirdefVersionWithinObject struct {
	Days     *int64 `json:"days,omitempty"`
	Versions *int64 `json:"versions,omitempty"`
}

/*
AntiMalwareVendorObject Product name

ShortName:
Parent chains:
*
* anti_malware
* vendor
* _inline

Args:

Param Name (string, required): the Name param. String length must not exceed 103 characters.

Param Products ([]string): the Products param.
*/
type AntiMalwareVendorObject struct {
	Name     string   `json:"name"`
	Products []string `json:"product,omitempty"`
}

/*
CertificateObject object.

ShortName:
Parent chains:
*
* certificate

Args:

Param Criteria (CertificateCriteriaObject): the Criteria param.
*/
type CertificateObject struct {
	Criteria *CertificateCriteriaObject `json:"criteria,omitempty"`
}

/*
CertificateCriteriaObject object.

ShortName:
Parent chains:
*
* certificate
* criteria

Args:

Param CertificateAttributes ([]CertificateAttributeObject): the CertificateAttributes param.

Param CertificateProfile (string): Profile for authenticating client certificates
*/
type CertificateCriteriaObject struct {
	CertificateAttributes []CertificateAttributeObject `json:"certificate_attributes,omitempty"`
	CertificateProfile    *string                      `json:"certificate_profile,omitempty"`
}

/*
CertificateAttributeObject object.

ShortName:
Parent chains:
*
* certificate
* criteria
* certificate_attributes
* _inline

Args:

Param Name (string, required): Attribute Name

Param Value (string): Key value String length must not exceed 1024 characters. String validation regex: `.*`.
*/
type CertificateAttributeObject struct {
	Name  string  `json:"name"`
	Value *string `json:"value,omitempty"`
}

/*
CustomChecksObject object.

ShortName:
Parent chains:
*
* custom_checks

Args:

Param Criteria (CustomChecksCriteriaObject, required): the Criteria param.
*/
type CustomChecksObject struct {
	Criteria CustomChecksCriteriaObject `json:"criteria"`
}

/*
CustomChecksCriteriaObject object.

ShortName:
Parent chains:
*
* custom_checks
* criteria

Args:

Param Plist ([]PlistObject): the Plist param.

Param ProcessList ([]ProcessListObject): the ProcessList param.

Param RegistryKeys ([]RegistryKeyObject): the RegistryKeys param.
*/
type CustomChecksCriteriaObject struct {
	Plist        []PlistObject       `json:"plist,omitempty"`
	ProcessList  []ProcessListObject `json:"process_list,omitempty"`
	RegistryKeys []RegistryKeyObject `json:"registry_key,omitempty"`
}

/*
PlistObject object.

ShortName:
Parent chains:
*
* custom_checks
* criteria
* plist
* _inline

Args:

Param Keys ([]KeyObject): the Keys param.

Param Name (string, required): Preference list String length must not exceed 1023 characters.

Param Negate (bool): Plist does not exist Default: `false`.
*/
type PlistObject struct {
	Keys   []KeyObject `json:"key,omitempty"`
	Name   string      `json:"name"`
	Negate *bool       `json:"negate,omitempty"`
}

/*
KeyObject object.

ShortName:
Parent chains:
*
* custom_checks
* criteria
* plist
* _inline
* key
* _inline

Args:

Param Name (string, required): Key name String length must not exceed 1023 characters.

Param Negate (bool): Value does not exist or match specified value data Default: `false`.

Param Value (string): Key value String length must not exceed 1024 characters. String validation regex: `.*`.
*/
type KeyObject struct {
	Name   string  `json:"name"`
	Negate *bool   `json:"negate,omitempty"`
	Value  *string `json:"value,omitempty"`
}

/*
ProcessListObject object.

ShortName:
Parent chains:
*
* custom_checks
* criteria
* process_list
* _inline

Args:

Param Name (string, required): Process Name String length must not exceed 1023 characters.

Param Running (bool): the Running param. Default: `true`.
*/
type ProcessListObject struct {
	Name    string `json:"name"`
	Running *bool  `json:"running,omitempty"`
}

/*
RegistryKeyObject object.

ShortName:
Parent chains:
*
* custom_checks
* criteria
* registry_key
* _inline

Args:

Param DefaultValueData (string): Registry key default value data String length must not exceed 1024 characters. String validation regex: `.*`.

Param Name (string, required): Registry key String length must not exceed 1023 characters.

Param Negate (bool): Key does not exist or match specified value data Default: `false`.

Param RegistryValues ([]RegistryValueObject): the RegistryValues param.
*/
type RegistryKeyObject struct {
	DefaultValueData *string               `json:"default_value_data,omitempty"`
	Name             string                `json:"name"`
	Negate           *bool                 `json:"negate,omitempty"`
	RegistryValues   []RegistryValueObject `json:"registry_value,omitempty"`
}

/*
RegistryValueObject object.

ShortName:
Parent chains:
*
* custom_checks
* criteria
* registry_key
* _inline
* registry_value
* _inline

Args:

Param Name (string, required): Registry value name String length must not exceed 1023 characters.

Param Negate (bool): Value does not exist or match specified value data Default: `false`.

Param ValueData (string): Registry value data String length must not exceed 1024 characters. String validation regex: `.*`.
*/
type RegistryValueObject struct {
	Name      string  `json:"name"`
	Negate    *bool   `json:"negate,omitempty"`
	ValueData *string `json:"value_data,omitempty"`
}

/*
DataLossPreventionObject object.

ShortName:
Parent chains:
*
* data_loss_prevention

Args:

Param Criteria (DataLossPreventionCriteriaObject): the Criteria param.

Param ExcludeVendor (bool): the ExcludeVendor param. Default: `false`.

Param Vendors ([]DataLossPreventionVendorObject): Vendor name
*/
type DataLossPreventionObject struct {
	Criteria      *DataLossPreventionCriteriaObject `json:"criteria,omitempty"`
	ExcludeVendor *bool                             `json:"exclude_vendor,omitempty"`
	Vendors       []DataLossPreventionVendorObject  `json:"vendor,omitempty"`
}

/*
DataLossPreventionCriteriaObject object.

ShortName:
Parent chains:
*
* data_loss_prevention
* criteria

Args:

Param IsEnabled (string): is enabled String must be one of these: `"no"`, `"yes"`, `"not-available"`.

Param IsInstalled (bool): Is Installed Default: `true`.
*/
type DataLossPreventionCriteriaObject struct {
	IsEnabled   *string `json:"is_enabled,omitempty"`
	IsInstalled *bool   `json:"is_installed,omitempty"`
}

/*
DataLossPreventionVendorObject object.

ShortName:
Parent chains:
*
* data_loss_prevention
* vendor
* _inline

Args:

Param Name (string, required): the Name param. String length must not exceed 103 characters.

Param Products ([]string): Product name
*/
type DataLossPreventionVendorObject struct {
	Name     string   `json:"name"`
	Products []string `json:"product,omitempty"`
}

/*
DiskBackupObject object.

ShortName:
Parent chains:
*
* disk_backup

Args:

Param Criteria (DiskBackupCriteriaObject): the Criteria param.

Param ExcludeVendor (bool): the ExcludeVendor param. Default: `false`.

Param Vendors ([]DiskBackupVendorObject): Vendor name
*/
type DiskBackupObject struct {
	Criteria      *DiskBackupCriteriaObject `json:"criteria,omitempty"`
	ExcludeVendor *bool                     `json:"exclude_vendor,omitempty"`
	Vendors       []DiskBackupVendorObject  `json:"vendor,omitempty"`
}

/*
DiskBackupCriteriaObject object.

ShortName:
Parent chains:
*
* disk_backup
* criteria

Args:

Param IsInstalled (bool): Is Installed Default: `true`.

Param LastBackupTime (LastBackupTimeObject): the LastBackupTime param.
*/
type DiskBackupCriteriaObject struct {
	IsInstalled    *bool                 `json:"is_installed,omitempty"`
	LastBackupTime *LastBackupTimeObject `json:"last_backup_time,omitempty"`
}

/*
LastBackupTimeObject object.

ShortName:
Parent chains:
*
* disk_backup
* criteria
* last_backup_time

Args:

Param NotAvailable (any): the NotAvailable param.

Param NotWithin (LastBackupTimeNotWithinObject): the NotWithin param.

Param Within (LastBackupTimeWithinObject): the Within param.

NOTE:  One of the following params should be specified:
  - NotAvailable
  - Within
  - NotWithin
*/
type LastBackupTimeObject struct {
	NotAvailable any                            `json:"not_available,omitempty"`
	NotWithin    *LastBackupTimeNotWithinObject `json:"not_within,omitempty"`
	Within       *LastBackupTimeWithinObject    `json:"within,omitempty"`
}

/*
LastBackupTimeNotWithinObject object.

ShortName:
Parent chains:
*
* disk_backup
* criteria
* last_backup_time
* not_within

Args:

Param Days (int64): specify time in days Value must be between 1 and 65535. Default: `1`.

Param Hours (int64): specify time in hours Value must be between 1 and 65535. Default: `24`.

NOTE:  One of the following params should be specified:
  - Days
  - Hours
*/
type LastBackupTimeNotWithinObject struct {
	Days  *int64 `json:"days,omitempty"`
	Hours *int64 `json:"hours,omitempty"`
}

/*
LastBackupTimeWithinObject object.

ShortName:
Parent chains:
*
* disk_backup
* criteria
* last_backup_time
* within

Args:

Param Days (int64): specify time in days Value must be between 1 and 65535. Default: `1`.

Param Hours (int64): specify time in hours Value must be between 1 and 65535. Default: `24`.

NOTE:  One of the following params should be specified:
  - Days
  - Hours
*/
type LastBackupTimeWithinObject struct {
	Days  *int64 `json:"days,omitempty"`
	Hours *int64 `json:"hours,omitempty"`
}

/*
DiskBackupVendorObject Product name

ShortName:
Parent chains:
*
* disk_backup
* vendor
* _inline

Args:

Param Name (string, required): the Name param. String length must not exceed 103 characters.

Param Products ([]string): the Products param.
*/
type DiskBackupVendorObject struct {
	Name     string   `json:"name"`
	Products []string `json:"product,omitempty"`
}

/*
DiskEncryptionObject object.

ShortName:
Parent chains:
*
* disk_encryption

Args:

Param Criteria (DiskEncryptionCriteriaObject): Encryption locations

Param ExcludeVendor (bool): the ExcludeVendor param. Default: `false`.

Param Vendors ([]DiskEncryptionVendorObject): Vendor name
*/
type DiskEncryptionObject struct {
	Criteria      *DiskEncryptionCriteriaObject `json:"criteria,omitempty"`
	ExcludeVendor *bool                         `json:"exclude_vendor,omitempty"`
	Vendors       []DiskEncryptionVendorObject  `json:"vendor,omitempty"`
}

/*
DiskEncryptionCriteriaObject Encryption locations

ShortName:
Parent chains:
*
* disk_encryption
* criteria

Args:

Param EncryptedLocations ([]EncryptedLocationObject): the EncryptedLocations param.

Param IsInstalled (bool): Is Installed Default: `true`.
*/
type DiskEncryptionCriteriaObject struct {
	EncryptedLocations []EncryptedLocationObject `json:"encrypted_locations,omitempty"`
	IsInstalled        *bool                     `json:"is_installed,omitempty"`
}

/*
EncryptedLocationObject object.

ShortName:
Parent chains:
*
* disk_encryption
* criteria
* encrypted_locations
* _inline

Args:

Param EncryptionState (EncryptionStateObject): the EncryptionState param.

Param Name (string, required): Encryption location String length must not exceed 1023 characters.
*/
type EncryptedLocationObject struct {
	EncryptionState *EncryptionStateObject `json:"encryption_state,omitempty"`
	Name            string                 `json:"name"`
}

/*
EncryptionStateObject object.

ShortName:
Parent chains:
*
* disk_encryption
* criteria
* encrypted_locations
* _inline
* encryption_state

Args:

Param Is (string): the Is param. String must be one of these: `"encrypted"`, `"unencrypted"`, `"partial"`, `"unknown"`. Default: `"encrypted"`.

Param IsNot (string): the IsNot param. String must be one of these: `"encrypted"`, `"unencrypted"`, `"partial"`, `"unknown"`. Default: `"encrypted"`.

NOTE:  One of the following params should be specified:
  - Is
  - IsNot
*/
type EncryptionStateObject struct {
	Is    *string `json:"is,omitempty"`
	IsNot *string `json:"is_not,omitempty"`
}

/*
DiskEncryptionVendorObject Product name

ShortName:
Parent chains:
*
* disk_encryption
* vendor
* _inline

Args:

Param Name (string, required): the Name param. String length must not exceed 103 characters.

Param Products ([]string): the Products param.
*/
type DiskEncryptionVendorObject struct {
	Name     string   `json:"name"`
	Products []string `json:"product,omitempty"`
}

/*
FirewallObject object.

ShortName:
Parent chains:
*
* firewall

Args:

Param Criteria (FirewallCriteriaObject): the Criteria param.

Param ExcludeVendor (bool): the ExcludeVendor param. Default: `false`.

Param Vendors ([]FirewallVendorObject): Vendor name
*/
type FirewallObject struct {
	Criteria      *FirewallCriteriaObject `json:"criteria,omitempty"`
	ExcludeVendor *bool                   `json:"exclude_vendor,omitempty"`
	Vendors       []FirewallVendorObject  `json:"vendor,omitempty"`
}

/*
FirewallCriteriaObject object.

ShortName:
Parent chains:
*
* firewall
* criteria

Args:

Param IsEnabled (string): is enabled String must be one of these: `"no"`, `"yes"`, `"not-available"`.

Param IsInstalled (bool): Is Installed Default: `true`.
*/
type FirewallCriteriaObject struct {
	IsEnabled   *string `json:"is_enabled,omitempty"`
	IsInstalled *bool   `json:"is_installed,omitempty"`
}

/*
FirewallVendorObject Product name

ShortName:
Parent chains:
*
* firewall
* vendor
* _inline

Args:

Param Name (string, required): the Name param. String length must not exceed 103 characters.

Param Products ([]string): the Products param.
*/
type FirewallVendorObject struct {
	Name     string   `json:"name"`
	Products []string `json:"product,omitempty"`
}

/*
HostInfoObject object.

ShortName:
Parent chains:
*
* host_info

Args:

Param Criteria (HostInfoCriteriaObject, required): the Criteria param.
*/
type HostInfoObject struct {
	Criteria HostInfoCriteriaObject `json:"criteria"`
}

/*
HostInfoCriteriaObject object.

ShortName:
Parent chains:
*
* host_info
* criteria

Args:

Param ClientVersion (ClientVersionObject): the ClientVersion param.

Param Domain (DomainObject): the Domain param.

Param HostId (HostIdObject): the HostId param.

Param HostName (HostNameObject): the HostName param.

Param Managed (bool): If device is managed

Param Os (OsObject): the Os param.

Param SerialNumber (SerialNumberObject): the SerialNumber param.
*/
type HostInfoCriteriaObject struct {
	ClientVersion *ClientVersionObject `json:"client_version,omitempty"`
	Domain        *DomainObject        `json:"domain,omitempty"`
	HostId        *HostIdObject        `json:"host_id,omitempty"`
	HostName      *HostNameObject      `json:"host_name,omitempty"`
	Managed       *bool                `json:"managed,omitempty"`
	Os            *OsObject            `json:"os,omitempty"`
	SerialNumber  *SerialNumberObject  `json:"serial_number,omitempty"`
}

/*
ClientVersionObject object.

ShortName:
Parent chains:
*
* host_info
* criteria
* client_version

Args:

Param Contains (string): the Contains param. String length must not exceed 255 characters.

Param Is (string): the Is param. String length must not exceed 255 characters.

Param IsNot (string): the IsNot param. String length must not exceed 255 characters.

NOTE:  One of the following params should be specified:
  - Contains
  - Is
  - IsNot
*/
type ClientVersionObject struct {
	Contains *string `json:"contains,omitempty"`
	Is       *string `json:"is,omitempty"`
	IsNot    *string `json:"is_not,omitempty"`
}

/*
DomainObject object.

ShortName:
Parent chains:
*
* host_info
* criteria
* domain

Args:

Param Contains (string): the Contains param. String length must not exceed 255 characters.

Param Is (string): the Is param. String length must not exceed 255 characters.

Param IsNot (string): the IsNot param. String length must not exceed 255 characters.

NOTE:  One of the following params should be specified:
  - Contains
  - Is
  - IsNot
*/
type DomainObject struct {
	Contains *string `json:"contains,omitempty"`
	Is       *string `json:"is,omitempty"`
	IsNot    *string `json:"is_not,omitempty"`
}

/*
HostIdObject object.

ShortName:
Parent chains:
*
* host_info
* criteria
* host_id

Args:

Param Contains (string): the Contains param. String length must not exceed 255 characters.

Param Is (string): the Is param. String length must not exceed 255 characters.

Param IsNot (string): the IsNot param. String length must not exceed 255 characters.

NOTE:  One of the following params should be specified:
  - Contains
  - Is
  - IsNot
*/
type HostIdObject struct {
	Contains *string `json:"contains,omitempty"`
	Is       *string `json:"is,omitempty"`
	IsNot    *string `json:"is_not,omitempty"`
}

/*
HostNameObject object.

ShortName:
Parent chains:
*
* host_info
* criteria
* host_name

Args:

Param Contains (string): the Contains param. String length must not exceed 255 characters.

Param Is (string): the Is param. String length must not exceed 255 characters.

Param IsNot (string): the IsNot param. String length must not exceed 255 characters.

NOTE:  One of the following params should be specified:
  - Contains
  - Is
  - IsNot
*/
type HostNameObject struct {
	Contains *string `json:"contains,omitempty"`
	Is       *string `json:"is,omitempty"`
	IsNot    *string `json:"is_not,omitempty"`
}

/*
OsObject object.

ShortName:
Parent chains:
*
* host_info
* criteria
* os

Args:

Param Contains (ContainsObject): the Contains param.

NOTE:  One of the following params should be specified:
  - Contains
*/
type OsObject struct {
	Contains *ContainsObject `json:"contains,omitempty"`
}

/*
ContainsObject object.

ShortName:
Parent chains:
*
* host_info
* criteria
* os
* contains

Args:

Param Apple (string): Apple vendor String length must not exceed 255 characters. Default: `"All"`.

Param Google (string): Google vendor String length must not exceed 255 characters. Default: `"All"`.

Param Linux (string): Linux vendor String length must not exceed 255 characters. Default: `"All"`.

Param Microsoft (string): Microsoft vendor String length must not exceed 255 characters. Default: `"All"`.

Param Other (string): Other vendor String length must not exceed 255 characters.

NOTE:  One of the following params should be specified:
  - Microsoft
  - Apple
  - Google
  - Linux
  - Other
*/
type ContainsObject struct {
	Apple     *string `json:"Apple,omitempty"`
	Google    *string `json:"Google,omitempty"`
	Linux     *string `json:"Linux,omitempty"`
	Microsoft *string `json:"Microsoft,omitempty"`
	Other     *string `json:"Other,omitempty"`
}

/*
SerialNumberObject object.

ShortName:
Parent chains:
*
* host_info
* criteria
* serial_number

Args:

Param Contains (string): the Contains param. String length must not exceed 255 characters.

Param Is (string): the Is param. String length must not exceed 255 characters.

Param IsNot (string): the IsNot param. String length must not exceed 255 characters.

NOTE:  One of the following params should be specified:
  - Contains
  - Is
  - IsNot
*/
type SerialNumberObject struct {
	Contains *string `json:"contains,omitempty"`
	Is       *string `json:"is,omitempty"`
	IsNot    *string `json:"is_not,omitempty"`
}

/*
MobileDeviceObject object.

ShortName:
Parent chains:
*
* mobile_device

Args:

Param Criteria (MobileDeviceCriteriaObject): the Criteria param.
*/
type MobileDeviceObject struct {
	Criteria *MobileDeviceCriteriaObject `json:"criteria,omitempty"`
}

/*
MobileDeviceCriteriaObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria

Args:

Param Applications (ApplicationsObject): the Applications param.

Param DiskEncrypted (bool): If device's disk is encrypted

Param Imei (ImeiObject): the Imei param.

Param Jailbroken (bool): If device is by rooted/jailbroken

Param LastCheckinTime (LastCheckinTimeObject): the LastCheckinTime param.

Param Model (ModelObject): the Model param.

Param PasscodeSet (bool): If device's passcode is present

Param PhoneNumber (PhoneNumberObject): the PhoneNumber param.

Param Tag (TagObject): the Tag param.
*/
type MobileDeviceCriteriaObject struct {
	Applications    *ApplicationsObject    `json:"applications,omitempty"`
	DiskEncrypted   *bool                  `json:"disk_encrypted,omitempty"`
	Imei            *ImeiObject            `json:"imei,omitempty"`
	Jailbroken      *bool                  `json:"jailbroken,omitempty"`
	LastCheckinTime *LastCheckinTimeObject `json:"last_checkin_time,omitempty"`
	Model           *ModelObject           `json:"model,omitempty"`
	PasscodeSet     *bool                  `json:"passcode_set,omitempty"`
	PhoneNumber     *PhoneNumberObject     `json:"phone_number,omitempty"`
	Tag             *TagObject             `json:"tag,omitempty"`
}

/*
ApplicationsObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* applications

Args:

Param HasMalware (HasMalwareObject): the HasMalware param.

Param HasUnmanagedApp (bool): Has apps that are not managed

Param Includes ([]IncludeObject): the Includes param.
*/
type ApplicationsObject struct {
	HasMalware      *HasMalwareObject `json:"has_malware,omitempty"`
	HasUnmanagedApp *bool             `json:"has_unmanaged_app,omitempty"`
	Includes        []IncludeObject   `json:"includes,omitempty"`
}

/*
HasMalwareObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* applications
* has_malware

Args:

Param No (any): the No param.

Param Yes (YesObject): the Yes param.

NOTE:  One of the following params should be specified:
  - No
  - Yes
*/
type HasMalwareObject struct {
	No  any        `json:"no,omitempty"`
	Yes *YesObject `json:"yes,omitempty"`
}

/*
YesObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* applications
* has_malware
* yes

Args:

Param Excludes ([]ExcludeObject): the Excludes param.
*/
type YesObject struct {
	Excludes []ExcludeObject `json:"excludes,omitempty"`
}

/*
ExcludeObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* applications
* has_malware
* yes
* excludes
* _inline

Args:

Param Hash (string): application hash String length must not exceed 1024 characters. String validation regex: `.*`.

Param Name (string, required): the Name param. String length must not exceed 31 characters.

Param Package (string): application package name String length must not exceed 1024 characters. String validation regex: `.*`.
*/
type ExcludeObject struct {
	Hash    *string `json:"hash,omitempty"`
	Name    string  `json:"name"`
	Package *string `json:"package,omitempty"`
}

/*
IncludeObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* applications
* includes
* _inline

Args:

Param Hash (string): application hash String length must not exceed 1024 characters. String validation regex: `.*`.

Param Name (string, required): the Name param. String length must not exceed 31 characters.

Param Package (string): application package name String length must not exceed 1024 characters. String validation regex: `.*`.
*/
type IncludeObject struct {
	Hash    *string `json:"hash,omitempty"`
	Name    string  `json:"name"`
	Package *string `json:"package,omitempty"`
}

/*
ImeiObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* imei

Args:

Param Contains (string): the Contains param. String length must not exceed 255 characters.

Param Is (string): the Is param. String length must not exceed 255 characters.

Param IsNot (string): the IsNot param. String length must not exceed 255 characters.

NOTE:  One of the following params should be specified:
  - Contains
  - Is
  - IsNot
*/
type ImeiObject struct {
	Contains *string `json:"contains,omitempty"`
	Is       *string `json:"is,omitempty"`
	IsNot    *string `json:"is_not,omitempty"`
}

/*
LastCheckinTimeObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* last_checkin_time

Args:

Param NotWithin (LastCheckinTimeNotWithinObject): the NotWithin param.

Param Within (LastCheckinTimeWithinObject): the Within param.

NOTE:  One of the following params should be specified:
  - Within
  - NotWithin
*/
type LastCheckinTimeObject struct {
	NotWithin *LastCheckinTimeNotWithinObject `json:"not_within,omitempty"`
	Within    *LastCheckinTimeWithinObject    `json:"within,omitempty"`
}

/*
LastCheckinTimeNotWithinObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* last_checkin_time
* not_within

Args:

Param Days (int64, required): specify time in days Value must be between 1 and 365. Default: `30`.
*/
type LastCheckinTimeNotWithinObject struct {
	Days int64 `json:"days"`
}

/*
LastCheckinTimeWithinObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* last_checkin_time
* within

Args:

Param Days (int64, required): specify time in days Value must be between 1 and 365. Default: `30`.
*/
type LastCheckinTimeWithinObject struct {
	Days int64 `json:"days"`
}

/*
ModelObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* model

Args:

Param Contains (string): the Contains param. String length must not exceed 255 characters.

Param Is (string): the Is param. String length must not exceed 255 characters.

Param IsNot (string): the IsNot param. String length must not exceed 255 characters.

NOTE:  One of the following params should be specified:
  - Contains
  - Is
  - IsNot
*/
type ModelObject struct {
	Contains *string `json:"contains,omitempty"`
	Is       *string `json:"is,omitempty"`
	IsNot    *string `json:"is_not,omitempty"`
}

/*
PhoneNumberObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* phone_number

Args:

Param Contains (string): the Contains param. String length must not exceed 255 characters.

Param Is (string): the Is param. String length must not exceed 255 characters.

Param IsNot (string): the IsNot param. String length must not exceed 255 characters.

NOTE:  One of the following params should be specified:
  - Contains
  - Is
  - IsNot
*/
type PhoneNumberObject struct {
	Contains *string `json:"contains,omitempty"`
	Is       *string `json:"is,omitempty"`
	IsNot    *string `json:"is_not,omitempty"`
}

/*
TagObject object.

ShortName:
Parent chains:
*
* mobile_device
* criteria
* tag

Args:

Param Contains (string): the Contains param. String length must not exceed 255 characters.

Param Is (string): the Is param. String length must not exceed 255 characters.

Param IsNot (string): the IsNot param. String length must not exceed 255 characters.

NOTE:  One of the following params should be specified:
  - Contains
  - Is
  - IsNot
*/
type TagObject struct {
	Contains *string `json:"contains,omitempty"`
	Is       *string `json:"is,omitempty"`
	IsNot    *string `json:"is_not,omitempty"`
}

/*
NetworkInfoObject object.

ShortName:
Parent chains:
*
* network_info

Args:

Param Criteria (NetworkInfoCriteriaObject): the Criteria param.
*/
type NetworkInfoObject struct {
	Criteria *NetworkInfoCriteriaObject `json:"criteria,omitempty"`
}

/*
NetworkInfoCriteriaObject object.

ShortName:
Parent chains:
*
* network_info
* criteria

Args:

Param Network (NetworkObject): the Network param.
*/
type NetworkInfoCriteriaObject struct {
	Network *NetworkObject `json:"network,omitempty"`
}

/*
NetworkObject object.

ShortName:
Parent chains:
*
* network_info
* criteria
* network

Args:

Param Is (IsObject): the Is param.

Param IsNot (IsNotObject): the IsNot param.

NOTE:  One of the following params should be specified:
  - Is
  - IsNot
*/
type NetworkObject struct {
	Is    *IsObject    `json:"is,omitempty"`
	IsNot *IsNotObject `json:"is_not,omitempty"`
}

/*
IsObject object.

ShortName:
Parent chains:
*
* network_info
* criteria
* network
* is

Args:

Param Mobile (IsMobileObject): the Mobile param.

Param Unknown (any): the Unknown param.

Param Wifi (IsWifiObject): the Wifi param.

NOTE:  One of the following params should be specified:
  - Wifi
  - Mobile
  - Unknown
*/
type IsObject struct {
	Mobile  *IsMobileObject `json:"mobile,omitempty"`
	Unknown any             `json:"unknown,omitempty"`
	Wifi    *IsWifiObject   `json:"wifi,omitempty"`
}

/*
IsMobileObject object.

ShortName:
Parent chains:
*
* network_info
* criteria
* network
* is
* mobile

Args:

Param Carrier (string): the Carrier param. String length must not exceed 1023 characters. String validation regex: `.*`.
*/
type IsMobileObject struct {
	Carrier *string `json:"carrier,omitempty"`
}

/*
IsWifiObject object.

ShortName:
Parent chains:
*
* network_info
* criteria
* network
* is
* wifi

Args:

Param Ssid (string): SSID String length must not exceed 1023 characters. String validation regex: `.*`.
*/
type IsWifiObject struct {
	Ssid *string `json:"ssid,omitempty"`
}

/*
IsNotObject object.

ShortName:
Parent chains:
*
* network_info
* criteria
* network
* is_not

Args:

Param Ethernet (any): the Ethernet param.

Param Mobile (IsNotMobileObject): the Mobile param.

Param Unknown (any): the Unknown param.

Param Wifi (IsNotWifiObject): the Wifi param.

NOTE:  One of the following params should be specified:
  - Wifi
  - Mobile
  - Ethernet
  - Unknown
*/
type IsNotObject struct {
	Ethernet any                `json:"ethernet,omitempty"`
	Mobile   *IsNotMobileObject `json:"mobile,omitempty"`
	Unknown  any                `json:"unknown,omitempty"`
	Wifi     *IsNotWifiObject   `json:"wifi,omitempty"`
}

/*
IsNotMobileObject object.

ShortName:
Parent chains:
*
* network_info
* criteria
* network
* is_not
* mobile

Args:

Param Carrier (string): the Carrier param. String length must not exceed 1023 characters. String validation regex: `.*`.
*/
type IsNotMobileObject struct {
	Carrier *string `json:"carrier,omitempty"`
}

/*
IsNotWifiObject object.

ShortName:
Parent chains:
*
* network_info
* criteria
* network
* is_not
* wifi

Args:

Param Ssid (string): SSID String length must not exceed 1023 characters. String validation regex: `.*`.
*/
type IsNotWifiObject struct {
	Ssid *string `json:"ssid,omitempty"`
}

/*
PatchManagementObject object.

ShortName:
Parent chains:
*
* patch_management

Args:

Param Criteria (CriteriaObject): the Criteria param.

Param ExcludeVendor (bool): the ExcludeVendor param. Default: `false`.

Param Vendors ([]PatchManagementVendorObject): Vendor name
*/
type PatchManagementObject struct {
	Criteria      *CriteriaObject               `json:"criteria,omitempty"`
	ExcludeVendor *bool                         `json:"exclude_vendor,omitempty"`
	Vendors       []PatchManagementVendorObject `json:"vendor,omitempty"`
}

/*
CriteriaObject object.

ShortName:
Parent chains:
*
* patch_management
* criteria

Args:

Param IsEnabled (string): is enabled String must be one of these: `"no"`, `"yes"`, `"not-available"`.

Param IsInstalled (bool): Is Installed Default: `true`.

Param MissingPatches (MissingPatchesObject): the MissingPatches param.
*/
type CriteriaObject struct {
	IsEnabled      *string               `json:"is_enabled,omitempty"`
	IsInstalled    *bool                 `json:"is_installed,omitempty"`
	MissingPatches *MissingPatchesObject `json:"missing_patches,omitempty"`
}

/*
MissingPatchesObject object.

ShortName:
Parent chains:
*
* patch_management
* criteria
* missing_patches

Args:

Param Check (string, required): the Check param. String must be one of these: `"has-any"`, `"has-none"`, `"has-all"`. Default: `"has-any"`.

Param Patches ([]string): the Patches param.

Param Severity (SeverityObject): the Severity param.
*/
type MissingPatchesObject struct {
	Check    string          `json:"check"`
	Patches  []string        `json:"patches,omitempty"`
	Severity *SeverityObject `json:"severity,omitempty"`
}

/*
SeverityObject object.

ShortName:
Parent chains:
*
* patch_management
* criteria
* missing_patches
* severity

Args:

Param GreaterEqual (int64): the GreaterEqual param. Value must be between 0 and 100000.

Param GreaterThan (int64): the GreaterThan param. Value must be between 0 and 100000.

Param Is (int64): the Is param. Value must be between 0 and 100000.

Param IsNot (int64): the IsNot param. Value must be between 0 and 100000.

Param LessEqual (int64): the LessEqual param. Value must be between 0 and 100000.

Param LessThan (int64): the LessThan param. Value must be between 0 and 100000.

NOTE:  One of the following params should be specified:
  - GreaterEqual
  - GreaterThan
  - Is
  - IsNot
  - LessEqual
  - LessThan
*/
type SeverityObject struct {
	GreaterEqual *int64 `json:"greater_equal,omitempty"`
	GreaterThan  *int64 `json:"greater_than,omitempty"`
	Is           *int64 `json:"is,omitempty"`
	IsNot        *int64 `json:"is_not,omitempty"`
	LessEqual    *int64 `json:"less_equal,omitempty"`
	LessThan     *int64 `json:"less_than,omitempty"`
}

/*
PatchManagementVendorObject object.

ShortName:
Parent chains:
*
* patch_management
* vendor
* _inline

Args:

Param Name (string, required): the Name param. String length must not exceed 103 characters.

Param Products ([]string): Product name
*/
type PatchManagementVendorObject struct {
	Name     string   `json:"name"`
	Products []string `json:"product,omitempty"`
}
