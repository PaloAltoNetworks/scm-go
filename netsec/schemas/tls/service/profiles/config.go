package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/tls-service-profiles

/*
Config object.

ShortName: eXzfzFj
Parent chains:
*

Args:

Param Certificate (string, required): SSL certificate file name String length must not exceed 255 characters.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): SSL TLS Service Profile name, value is muCustomDomainSSLProfile when it is used on mobile-agent infra settings. String length must not exceed 127 characters. String validation regex: `^[a-zA-Z0-9._-]+$`.

Param ProtocolSettings (ProtocolSettingsObject, required): the ProtocolSettings param.
*/
type Config struct {
	Certificate      string                 `json:"certificate"`
	Id               *string                `json:"id,omitempty"`
	Name             string                 `json:"name"`
	ProtocolSettings ProtocolSettingsObject `json:"protocol_settings"`
}

/*
ProtocolSettingsObject object.

ShortName:
Parent chains:
*
* protocol_settings

Args:

Param AuthAlgoSha1 (bool): Allow authentication SHA1

Param AuthAlgoSha256 (bool): Allow authentication SHA256

Param AuthAlgoSha384 (bool): Allow authentication SHA384

Param EncAlgo3des (bool): Allow algorithm 3DES

Param EncAlgoAes128Cbc (bool): Allow algorithm AES-128-CBC

Param EncAlgoAes128Gcm (bool): Allow algorithm AES-128-GCM

Param EncAlgoAes256Cbc (bool): Allow algorithm AES-256-CBC

Param EncAlgoAes256Gcm (bool): Allow algorithm AES-256-GCM

Param EncAlgoRc4 (bool): Allow algorithm RC4

Param KeyxchgAlgoDhe (bool): Allow algorithm DHE

Param KeyxchgAlgoEcdhe (bool): Allow algorithm ECDHE

Param KeyxchgAlgoRsa (bool): Allow algorithm RSA

Param MaxVersion (string): the MaxVersion param. String must be one of these: `"tls1-0"`, `"tls1-1"`, `"tls1-2"`, `"tls1-3"`, `"max"`. Default: `"max"`.

Param MinVersion (string): the MinVersion param. String must be one of these: `"tls1-0"`, `"tls1-1"`, `"tls1-2"`. Default: `"tls1-0"`.
*/
type ProtocolSettingsObject struct {
	AuthAlgoSha1     *bool   `json:"auth_algo_sha1,omitempty"`
	AuthAlgoSha256   *bool   `json:"auth_algo_sha256,omitempty"`
	AuthAlgoSha384   *bool   `json:"auth_algo_sha384,omitempty"`
	EncAlgo3des      *bool   `json:"enc_algo_3des,omitempty"`
	EncAlgoAes128Cbc *bool   `json:"enc_algo_aes_128_cbc,omitempty"`
	EncAlgoAes128Gcm *bool   `json:"enc_algo_aes_128_gcm,omitempty"`
	EncAlgoAes256Cbc *bool   `json:"enc_algo_aes_256_cbc,omitempty"`
	EncAlgoAes256Gcm *bool   `json:"enc_algo_aes_256_gcm,omitempty"`
	EncAlgoRc4       *bool   `json:"enc_algo_rc4,omitempty"`
	KeyxchgAlgoDhe   *bool   `json:"keyxchg_algo_dhe,omitempty"`
	KeyxchgAlgoEcdhe *bool   `json:"keyxchg_algo_ecdhe,omitempty"`
	KeyxchgAlgoRsa   *bool   `json:"keyxchg_algo_rsa,omitempty"`
	MaxVersion       *string `json:"max_version,omitempty"`
	MinVersion       *string `json:"min_version,omitempty"`
}
