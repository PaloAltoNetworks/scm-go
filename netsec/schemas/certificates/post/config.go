package post

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/certificates-post

/*
Config object.

ShortName: mqfebzt
Parent chains:
*

Args:

Param Algorithm (AlgorithmObject, required): the Algorithm param.

Param AlternateEmails ([]string): the AlternateEmails param.

Param CertificateName (string, required): the CertificateName param. String length must exceed 1 characters.

Param CommonName (string, required): the CommonName param. String length must exceed 1 characters.

Param CountryCode (string): the CountryCode param.

Param DayTillExpiration (int64): the DayTillExpiration param.

Param Departments ([]string): the Departments param.

Param Digest (string, required): the Digest param. String must be one of these: `"sha1"`, `"sha256"`, `"sha384"`, `"sha512"`, `"md5"`.

Param Email (string): the Email param. String length must not exceed 255 characters.

Param Hostnames ([]string): the Hostnames param.

Param Ips ([]string): the Ips param.

Param IsBlockPrivateKey (bool): the IsBlockPrivateKey param.

Param IsCertificateAuthority (bool): the IsCertificateAuthority param.

Param Locality (string): the Locality param. String length must not exceed 64 characters.

Param OcspResponderUrl (string): the OcspResponderUrl param. String length must not exceed 64 characters.

Param SignedBy (string): the SignedBy param. String length must not exceed 64 characters.

Param State (string): the State param. String length must not exceed 32 characters.
*/
type Config struct {
	Algorithm              AlgorithmObject `json:"algorithm"`
	AlternateEmails        []string        `json:"alternate_email,omitempty"`
	CertificateName        string          `json:"certificate_name"`
	CommonName             string          `json:"common_name"`
	CountryCode            *string         `json:"country_code,omitempty"`
	DayTillExpiration      *int64          `json:"day_till_expiration,omitempty"`
	Departments            []string        `json:"department,omitempty"`
	Digest                 string          `json:"digest"`
	Email                  *string         `json:"email,omitempty"`
	Hostnames              []string        `json:"hostname,omitempty"`
	Ips                    []string        `json:"ip,omitempty"`
	IsBlockPrivateKey      *bool           `json:"is_block_privateKey,omitempty"`
	IsCertificateAuthority *bool           `json:"is_certificate_authority,omitempty"`
	Locality               *string         `json:"locality,omitempty"`
	OcspResponderUrl       *string         `json:"ocsp_responder_url,omitempty"`
	SignedBy               *string         `json:"signed_by,omitempty"`
	State                  *string         `json:"state,omitempty"`
}

/*
AlgorithmObject object.

ShortName:
Parent chains:
*
* algorithm

Args:

Param EcdsaNumberOfBits (int64): the EcdsaNumberOfBits param.

Param RsaNumberOfBits (int64): the RsaNumberOfBits param.

NOTE:  One of the following params should be specified:
  - RsaNumberOfBits
  - EcdsaNumberOfBits
*/
type AlgorithmObject struct {
	EcdsaNumberOfBits *int64 `json:"ecdsa_number_of_bits,omitempty"`
	RsaNumberOfBits   *int64 `json:"rsa_number_of_bits,omitempty"`
}
