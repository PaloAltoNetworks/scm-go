package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/ipsec-crypto-profiles

/*
Config object.

ShortName: lrzxLXR
Parent chains:
*

Args:

Param Ah (AhObject): the Ah param.

Param DhGroup (string): phase-2 DH group (PFS DH group) String must be one of these: `"no-pfs"`, `"group1"`, `"group2"`, `"group5"`, `"group14"`, `"group19"`, `"group20"`. Default: `"group2"`.

Param Esp (EspObject): the Esp param.

Param Id (string, read-only): UUID of the resource

Param Lifesize (LifesizeObject): the Lifesize param.

Param Lifetime (LifetimeObject, required): the Lifetime param.

Param Name (string, required): Alphanumeric string begin with letter: [0-9a-zA-Z._-] String length must not exceed 31 characters.

NOTE:  One of the following params should be specified:
  - Esp
  - Ah
*/
type Config struct {
	Ah       *AhObject       `json:"ah,omitempty"`
	DhGroup  *string         `json:"dh_group,omitempty"`
	Esp      *EspObject      `json:"esp,omitempty"`
	Id       *string         `json:"id,omitempty"`
	Lifesize *LifesizeObject `json:"lifesize,omitempty"`
	Lifetime LifetimeObject  `json:"lifetime"`
	Name     string          `json:"name"`
}

/*
AhObject object.

ShortName:
Parent chains:
*
* ah

Args:

Param Authentications ([]string, required): the Authentications param.
*/
type AhObject struct {
	Authentications []string `json:"authentication"`
}

/*
EspObject object.

ShortName:
Parent chains:
*
* esp

Args:

Param Authentications ([]string, required): Authentication algorithm

Param Encryptions ([]string, required): Encryption algorithm
*/
type EspObject struct {
	Authentications []string `json:"authentication"`
	Encryptions     []string `json:"encryption"`
}

/*
LifesizeObject object.

ShortName:
Parent chains:
*
* lifesize

Args:

Param Gb (int64): specify lifesize in gigabytes(GB) Value must be between 1 and 65535.

Param Kb (int64): specify lifesize in kilobytes(KB) Value must be between 1 and 65535.

Param Mb (int64): specify lifesize in megabytes(MB) Value must be between 1 and 65535.

Param Tb (int64): specify lifesize in terabytes(TB) Value must be between 1 and 65535.

NOTE:  One of the following params should be specified:
  - Kb
  - Mb
  - Gb
  - Tb
*/
type LifesizeObject struct {
	Gb *int64 `json:"gb,omitempty"`
	Kb *int64 `json:"kb,omitempty"`
	Mb *int64 `json:"mb,omitempty"`
	Tb *int64 `json:"tb,omitempty"`
}

/*
LifetimeObject object.

ShortName:
Parent chains:
*
* lifetime

Args:

Param Days (int64): specify lifetime in days Value must be between 1 and 365.

Param Hours (int64): specify lifetime in hours Value must be between 1 and 65535.

Param Minutes (int64): specify lifetime in minutes Value must be between 3 and 65535.

Param Seconds (int64): specify lifetime in seconds Value must be between 180 and 65535.

NOTE:  One of the following params should be specified:
  - Seconds
  - Minutes
  - Hours
  - Days
*/
type LifetimeObject struct {
	Days    *int64 `json:"days,omitempty"`
	Hours   *int64 `json:"hours,omitempty"`
	Minutes *int64 `json:"minutes,omitempty"`
	Seconds *int64 `json:"seconds,omitempty"`
}
