package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/scep-profiles

/*
Config object.

ShortName: nnsRzDg
Parent chains:
*

Args:

Param Algorithm (AlgorithmObject): the Algorithm param.

Param CaIdentityName (string, required): the CaIdentityName param.

Param CertificateAttributes (CertificateAttributesObject): the CertificateAttributes param.

Param Digest (string, required): the Digest param.

Param Fingerprint (string): the Fingerprint param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param ScepCaCert (string): the ScepCaCert param.

Param ScepChallenge (ScepChallengeObject): the ScepChallenge param.

Param ScepClientCert (string): the ScepClientCert param.

Param ScepUrl (string, required): the ScepUrl param.

Param Subject (string): the Subject param.

Param UseAsDigitalSignature (bool): the UseAsDigitalSignature param.

Param UseForKeyEncipherment (bool): the UseForKeyEncipherment param.
*/
type Config struct {
	Algorithm             *AlgorithmObject             `json:"algorithm,omitempty"`
	CaIdentityName        string                       `json:"ca_identity_name"`
	CertificateAttributes *CertificateAttributesObject `json:"certificate_attributes,omitempty"`
	Digest                string                       `json:"digest"`
	Fingerprint           *string                      `json:"fingerprint,omitempty"`
	Id                    *string                      `json:"id,omitempty"`
	Name                  string                       `json:"name"`
	ScepCaCert            *string                      `json:"scep_ca_cert,omitempty"`
	ScepChallenge         *ScepChallengeObject         `json:"scep_challenge,omitempty"`
	ScepClientCert        *string                      `json:"scep_client_cert,omitempty"`
	ScepUrl               string                       `json:"scep_url"`
	Subject               *string                      `json:"subject,omitempty"`
	UseAsDigitalSignature *bool                        `json:"use_as_digital_signature,omitempty"`
	UseForKeyEncipherment *bool                        `json:"use_for_key_encipherment,omitempty"`
}

/*
AlgorithmObject object.

ShortName:
Parent chains:
*
* algorithm

Args:

Param Rsa (RsaObject): the Rsa param.
*/
type AlgorithmObject struct {
	Rsa *RsaObject `json:"rsa,omitempty"`
}

/*
RsaObject object.

ShortName:
Parent chains:
*
* algorithm
* rsa

Args:

Param RsaNbits (string): the RsaNbits param.
*/
type RsaObject struct {
	RsaNbits *string `json:"rsa_nbits,omitempty"`
}

/*
CertificateAttributesObject object.

ShortName:
Parent chains:
*
* certificate_attributes

Args:

Param Dnsname (string): the Dnsname param.

Param Rfc822name (string): the Rfc822name param.

Param UniformResourceIdentifier (string): the UniformResourceIdentifier param.

NOTE:  One of the following params should be specified:
  - Rfc822name
  - Dnsname
  - UniformResourceIdentifier
*/
type CertificateAttributesObject struct {
	Dnsname                   *string `json:"dnsname,omitempty"`
	Rfc822name                *string `json:"rfc822name,omitempty"`
	UniformResourceIdentifier *string `json:"uniform_resource_identifier,omitempty"`
}

/*
ScepChallengeObject object.

ShortName:
Parent chains:
*
* scep_challenge

Args:

Param DynamicChallenge (DynamicChallengeObject): the DynamicChallenge param.

Param Fixed (string): Challenge to use for SCEP server on mobile clients String length must not exceed 1024 characters.

Param None (string): the None param. String must be one of these: `""`.

NOTE:  One of the following params should be specified:
  - None
  - Fixed
  - DynamicChallenge
*/
type ScepChallengeObject struct {
	DynamicChallenge *DynamicChallengeObject `json:"dynamic,omitempty"`
	Fixed            *string                 `json:"fixed,omitempty"`
	None             *string                 `json:"none,omitempty"`
}

/*
DynamicChallengeObject object.

ShortName:
Parent chains:
*
* scep_challenge
* dynamic

Args:

Param OtpServerUrl (string): the OtpServerUrl param. String length must not exceed 255 characters.

Param Password (string): the Password param. String length must not exceed 255 characters.

Param Username (string): the Username param. String length must not exceed 255 characters.
*/
type DynamicChallengeObject struct {
	OtpServerUrl *string `json:"otp_server_url,omitempty"`
	Password     *string `json:"password,omitempty"`
	Username     *string `json:"username,omitempty"`
}
