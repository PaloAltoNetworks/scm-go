package get

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/certificates-get

/*
Config object.

ShortName: tephihM
Parent chains:
*

Args:

Param Algorithm (string): the Algorithm param.

Param Ca (bool): the Ca param.

Param CommonName (string): the CommonName param.

Param CommonNameInt (string): the CommonNameInt param.

Param ExpiryEpoch (string): the ExpiryEpoch param.

Param Id (string, read-only): UUID of the resource

Param Issuer (string): the Issuer param.

Param IssuerHash (string): the IssuerHash param.

Param NotValidAfter (string): the NotValidAfter param.

Param NotValidBefore (string): the NotValidBefore param.

Param PublicKey (string): the PublicKey param.

Param Subject (string): the Subject param.

Param SubjectHash (string): the SubjectHash param.

Param SubjectInt (string): the SubjectInt param.
*/
type Config struct {
	Algorithm      *string `json:"algorithm,omitempty"`
	Ca             *bool   `json:"ca,omitempty"`
	CommonName     *string `json:"common_name,omitempty"`
	CommonNameInt  *string `json:"common_name_int,omitempty"`
	ExpiryEpoch    *string `json:"expiry_epoch,omitempty"`
	Id             *string `json:"id,omitempty"`
	Issuer         *string `json:"issuer,omitempty"`
	IssuerHash     *string `json:"issuer_hash,omitempty"`
	NotValidAfter  *string `json:"not_valid_after,omitempty"`
	NotValidBefore *string `json:"not_valid_before,omitempty"`
	PublicKey      *string `json:"public_key,omitempty"`
	Subject        *string `json:"subject,omitempty"`
	SubjectHash    *string `json:"subject_hash,omitempty"`
	SubjectInt     *string `json:"subject_int,omitempty"`
}
