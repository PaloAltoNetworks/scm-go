package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/ike-crypto-profiles

/*
Config object.

ShortName: wugpput
Parent chains:
*

Args:

Param AuthenticationMultiple (int64): IKEv2 SA reauthentication interval equals authetication-multiple * rekey-lifetime; 0 means reauthentication disabled Value must be less than or equal to 50. Default: `0`.

Param DhGroups ([]string, required): the DhGroups param.

Param Encryptions ([]string, required): Encryption algorithm

Param Hashes ([]string, required): the Hashes param.

Param Id (string, read-only): UUID of the resource

Param Lifetime (LifetimeObject): the Lifetime param.

Param Name (string, required): Alphanumeric string begin with letter: [0-9a-zA-Z._-] String length must not exceed 31 characters.
*/
type Config struct {
	AuthenticationMultiple *int64          `json:"authentication_multiple,omitempty"`
	DhGroups               []string        `json:"dh_group"`
	Encryptions            []string        `json:"encryption"`
	Hashes                 []string        `json:"hash"`
	Id                     *string         `json:"id,omitempty"`
	Lifetime               *LifetimeObject `json:"lifetime,omitempty"`
	Name                   string          `json:"name"`
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
