package authorities

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/trusted-certificate-authorities

/*
Config object.

ShortName: iYmUVvF
Parent chains:
*

Args:

Param CommonName (string): the CommonName param. String length must not exceed 255 characters.

Param ExpiryEpoch (string): the ExpiryEpoch param.

Param Filename (string): the Filename param.

Param Id (string, read-only): UUID of the resource

Param Issuer (string): the Issuer param.

Param Name (string): the Name param. String length must not exceed 63 characters.

Param NotValidAfter (string): the NotValidAfter param.

Param NotValidBefore (string): the NotValidBefore param.

Param SerialNumber (string): the SerialNumber param.

Param Subject (string): the Subject param.
*/
type Config struct {
	CommonName     *string `json:"common_name,omitempty"`
	ExpiryEpoch    *string `json:"expiry_epoch,omitempty"`
	Filename       *string `json:"filename,omitempty"`
	Id             *string `json:"id,omitempty"`
	Issuer         *string `json:"issuer,omitempty"`
	Name           *string `json:"name,omitempty"`
	NotValidAfter  *string `json:"not_valid_after,omitempty"`
	NotValidBefore *string `json:"not_valid_before,omitempty"`
	SerialNumber   *string `json:"serial_number,omitempty"`
	Subject        *string `json:"subject,omitempty"`
}
