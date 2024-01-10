package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/authentication-profiles

/*
Config object.

ShortName: eumQbRC
Parent chains:
*

Args:

Param AllowList ([]string): the AllowList param.

Param Id (string, read-only): UUID of the resource

Param Lockout (LockoutObject): the Lockout param.

Param Method (MethodObject): the Method param.

Param MultiFactorAuth (MultiFactorAuthObject): the MultiFactorAuth param.

Param Name (string, required): the Name param.

Param SingleSignOn (SingleSignOnObject): the SingleSignOn param.

Param UserDomain (string): the UserDomain param. String length must not exceed 63 characters.

Param UsernameModifier (string): the UsernameModifier param. String must be one of these: `"%USERINPUT%"`, `"%USERINPUT%@%USERDOMAIN%"`, `"%USERDOMAIN%\\%USERINPUT%"`.
*/
type Config struct {
	AllowList        []string               `json:"allow_list,omitempty"`
	Id               *string                `json:"id,omitempty"`
	Lockout          *LockoutObject         `json:"lockout,omitempty"`
	Method           *MethodObject          `json:"method,omitempty"`
	MultiFactorAuth  *MultiFactorAuthObject `json:"multi_factor_auth,omitempty"`
	Name             string                 `json:"name"`
	SingleSignOn     *SingleSignOnObject    `json:"single_sign_on,omitempty"`
	UserDomain       *string                `json:"user_domain,omitempty"`
	UsernameModifier *string                `json:"username_modifier,omitempty"`
}

/*
LockoutObject object.

ShortName:
Parent chains:
*
* lockout

Args:

Param FailedAttempts (int64): the FailedAttempts param. Value must be between 0 and 10.

Param LockoutTime (int64): the LockoutTime param. Value must be between 0 and 60.
*/
type LockoutObject struct {
	FailedAttempts *int64 `json:"failed_attempts,omitempty"`
	LockoutTime    *int64 `json:"lockout_time,omitempty"`
}

/*
MethodObject object.

ShortName:
Parent chains:
*
* method

Args:

Param Cloud (CloudObject): the Cloud param.

Param Kerberos (KerberosObject): the Kerberos param.

Param Ldap (LdapObject): the Ldap param.

Param LocalDatabase (any): the LocalDatabase param. Default: `false`.

Param Radius (RadiusObject): the Radius param.

Param SamlIdp (SamlIdpObject): the SamlIdp param.

Param Tacplus (TacplusObject): the Tacplus param.

NOTE:  One of the following params should be specified:
  - LocalDatabase
  - SamlIdp
  - Ldap
  - Radius
  - Tacplus
  - Kerberos
  - Cloud
*/
type MethodObject struct {
	Cloud         *CloudObject    `json:"cloud,omitempty"`
	Kerberos      *KerberosObject `json:"kerberos,omitempty"`
	Ldap          *LdapObject     `json:"ldap,omitempty"`
	LocalDatabase any             `json:"local_database,omitempty"`
	Radius        *RadiusObject   `json:"radius,omitempty"`
	SamlIdp       *SamlIdpObject  `json:"saml_idp,omitempty"`
	Tacplus       *TacplusObject  `json:"tacplus,omitempty"`
}

/*
CloudObject object.

ShortName:
Parent chains:
*
* method
* cloud

Args:

Param ProfileName (string): The tenant profile name
*/
type CloudObject struct {
	ProfileName *string `json:"profile_name,omitempty"`
}

/*
KerberosObject object.

ShortName:
Parent chains:
*
* method
* kerberos

Args:

Param Realm (string): the Realm param.

Param ServerProfile (string): the ServerProfile param.
*/
type KerberosObject struct {
	Realm         *string `json:"realm,omitempty"`
	ServerProfile *string `json:"server_profile,omitempty"`
}

/*
LdapObject object.

ShortName:
Parent chains:
*
* method
* ldap

Args:

Param LoginAttribute (string): the LoginAttribute param.

Param PasswdExpDays (int64): the PasswdExpDays param.

Param ServerProfile (string): the ServerProfile param.
*/
type LdapObject struct {
	LoginAttribute *string `json:"login_attribute,omitempty"`
	PasswdExpDays  *int64  `json:"passwd_exp_days,omitempty"`
	ServerProfile  *string `json:"server_profile,omitempty"`
}

/*
RadiusObject object.

ShortName:
Parent chains:
*
* method
* radius

Args:

Param Checkgroup (bool): the Checkgroup param.

Param ServerProfile (string): the ServerProfile param.
*/
type RadiusObject struct {
	Checkgroup    *bool   `json:"checkgroup,omitempty"`
	ServerProfile *string `json:"server_profile,omitempty"`
}

/*
SamlIdpObject object.

ShortName:
Parent chains:
*
* method
* saml_idp

Args:

Param AttributeNameUsergroup (string): the AttributeNameUsergroup param. String length must be between 1 and 63 characters.

Param AttributeNameUsername (string): the AttributeNameUsername param. String length must be between 1 and 63 characters.

Param CertificateProfile (string): the CertificateProfile param. String length must not exceed 31 characters.

Param EnableSingleLogout (bool): the EnableSingleLogout param.

Param RequestSigningCertificate (string): the RequestSigningCertificate param. String length must not exceed 64 characters.

Param ServerProfile (string): the ServerProfile param. String length must not exceed 63 characters.
*/
type SamlIdpObject struct {
	AttributeNameUsergroup    *string `json:"attribute_name_usergroup,omitempty"`
	AttributeNameUsername     *string `json:"attribute_name_username,omitempty"`
	CertificateProfile        *string `json:"certificate_profile,omitempty"`
	EnableSingleLogout        *bool   `json:"enable_single_logout,omitempty"`
	RequestSigningCertificate *string `json:"request_signing_certificate,omitempty"`
	ServerProfile             *string `json:"server_profile,omitempty"`
}

/*
TacplusObject object.

ShortName:
Parent chains:
*
* method
* tacplus

Args:

Param Checkgroup (bool): the Checkgroup param.

Param ServerProfile (string): the ServerProfile param.
*/
type TacplusObject struct {
	Checkgroup    *bool   `json:"checkgroup,omitempty"`
	ServerProfile *string `json:"server_profile,omitempty"`
}

/*
MultiFactorAuthObject object.

ShortName:
Parent chains:
*
* multi_factor_auth

Args:

Param Factors ([]string): the Factors param.

Param MfaEnable (bool): the MfaEnable param.
*/
type MultiFactorAuthObject struct {
	Factors   []string `json:"factors,omitempty"`
	MfaEnable *bool    `json:"mfa_enable,omitempty"`
}

/*
SingleSignOnObject object.

ShortName:
Parent chains:
*
* single_sign_on

Args:

Param KerberosKeytab (string): the KerberosKeytab param. String length must not exceed 8192 characters.

Param Realm (string): the Realm param. String length must not exceed 127 characters.
*/
type SingleSignOnObject struct {
	KerberosKeytab *string `json:"kerberos_keytab,omitempty"`
	Realm          *string `json:"realm,omitempty"`
}
