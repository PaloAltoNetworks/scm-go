package certimport

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/certificates-import

/*
Config object.

ShortName: wugpput
Parent chains:
*

Args:

Param CertificateFile (string, required): base64 encoded content of the certificate file

Param Format (string, required): the Format param. String must be one of these: `"pem"`, `"pkcs12"`, `"der"`. Default: `"pem"`.

Param KeyFile (string): base64 encoded content of the key file

Param Name (string, required): name of the certificate String length must exceed 1 characters.

Param Passphrase (string): required when key_file is presented
*/
type Config struct {
	CertificateFile string  `json:"certificate_file"`
	Format          string  `json:"format"`
	KeyFile         *string `json:"key_file,omitempty"`
	Name            string  `json:"name"`
	Passphrase      *string `json:"passphrase,omitempty"`
}
