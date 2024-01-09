package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/decryption-profiles

/*
Config object.

ShortName: wugpput
Parent chains:
*

Args:

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Must start with alphanumeric char and should contain only alphanemeric, underscore, hyphen, dot or space String validation regex: `^[A-Za-z0-9]{1}[A-Za-z0-9_\-\.\s]{0,}$`.

Param SslForwardProxy (SslForwardProxyObject): the SslForwardProxy param.

Param SslInboundProxy (SslInboundProxyObject): the SslInboundProxy param.

Param SslNoProxy (SslNoProxyObject): the SslNoProxy param.

Param SslProtocolSettings (SslProtocolSettingsObject): the SslProtocolSettings param.
*/
type Config struct {
	Id                  *string                    `json:"id,omitempty"`
	Name                string                     `json:"name"`
	SslForwardProxy     *SslForwardProxyObject     `json:"ssl_forward_proxy,omitempty"`
	SslInboundProxy     *SslInboundProxyObject     `json:"ssl_inbound_proxy,omitempty"`
	SslNoProxy          *SslNoProxyObject          `json:"ssl_no_proxy,omitempty"`
	SslProtocolSettings *SslProtocolSettingsObject `json:"ssl_protocol_settings,omitempty"`
}

/*
SslForwardProxyObject object.

ShortName:
Parent chains:
*
* ssl_forward_proxy

Args:

Param AutoIncludeAltname (bool): the AutoIncludeAltname param. Default: `false`.

Param BlockClientCert (bool): the BlockClientCert param. Default: `false`.

Param BlockExpiredCertificate (bool): the BlockExpiredCertificate param. Default: `false`.

Param BlockTimeoutCert (bool): the BlockTimeoutCert param. Default: `false`.

Param BlockTls13DowngradeNoResource (bool): the BlockTls13DowngradeNoResource param. Default: `false`.

Param BlockUnknownCert (bool): the BlockUnknownCert param. Default: `false`.

Param BlockUnsupportedCipher (bool): the BlockUnsupportedCipher param. Default: `false`.

Param BlockUnsupportedVersion (bool): the BlockUnsupportedVersion param. Default: `false`.

Param BlockUntrustedIssuer (bool): the BlockUntrustedIssuer param. Default: `false`.

Param RestrictCertExts (bool): the RestrictCertExts param. Default: `false`.

Param StripAlpn (bool): the StripAlpn param. Default: `false`.
*/
type SslForwardProxyObject struct {
	AutoIncludeAltname            *bool `json:"auto_include_altname,omitempty"`
	BlockClientCert               *bool `json:"block_client_cert,omitempty"`
	BlockExpiredCertificate       *bool `json:"block_expired_certificate,omitempty"`
	BlockTimeoutCert              *bool `json:"block_timeout_cert,omitempty"`
	BlockTls13DowngradeNoResource *bool `json:"block_tls13_downgrade_no_resource,omitempty"`
	BlockUnknownCert              *bool `json:"block_unknown_cert,omitempty"`
	BlockUnsupportedCipher        *bool `json:"block_unsupported_cipher,omitempty"`
	BlockUnsupportedVersion       *bool `json:"block_unsupported_version,omitempty"`
	BlockUntrustedIssuer          *bool `json:"block_untrusted_issuer,omitempty"`
	RestrictCertExts              *bool `json:"restrict_cert_exts,omitempty"`
	StripAlpn                     *bool `json:"strip_alpn,omitempty"`
}

/*
SslInboundProxyObject object.

ShortName:
Parent chains:
*
* ssl_inbound_proxy

Args:

Param BlockIfHsmUnavailable (bool): the BlockIfHsmUnavailable param. Default: `false`.

Param BlockIfNoResource (bool): the BlockIfNoResource param. Default: `false`.

Param BlockUnsupportedCipher (bool): the BlockUnsupportedCipher param. Default: `false`.

Param BlockUnsupportedVersion (bool): the BlockUnsupportedVersion param. Default: `false`.
*/
type SslInboundProxyObject struct {
	BlockIfHsmUnavailable   *bool `json:"block_if_hsm_unavailable,omitempty"`
	BlockIfNoResource       *bool `json:"block_if_no_resource,omitempty"`
	BlockUnsupportedCipher  *bool `json:"block_unsupported_cipher,omitempty"`
	BlockUnsupportedVersion *bool `json:"block_unsupported_version,omitempty"`
}

/*
SslNoProxyObject object.

ShortName:
Parent chains:
*
* ssl_no_proxy

Args:

Param BlockExpiredCertificate (bool): the BlockExpiredCertificate param. Default: `false`.

Param BlockUntrustedIssuer (bool): the BlockUntrustedIssuer param. Default: `false`.
*/
type SslNoProxyObject struct {
	BlockExpiredCertificate *bool `json:"block_expired_certificate,omitempty"`
	BlockUntrustedIssuer    *bool `json:"block_untrusted_issuer,omitempty"`
}

/*
SslProtocolSettingsObject object.

ShortName:
Parent chains:
*
* ssl_protocol_settings

Args:

Param AuthAlgoMd5 (bool): the AuthAlgoMd5 param. Default: `true`.

Param AuthAlgoSha1 (bool): the AuthAlgoSha1 param. Default: `true`.

Param AuthAlgoSha256 (bool): the AuthAlgoSha256 param. Default: `true`.

Param AuthAlgoSha384 (bool): the AuthAlgoSha384 param. Default: `true`.

Param EncAlgo3des (bool): the EncAlgo3des param. Default: `true`.

Param EncAlgoAes128Cbc (bool): the EncAlgoAes128Cbc param. Default: `true`.

Param EncAlgoAes128Gcm (bool): the EncAlgoAes128Gcm param. Default: `true`.

Param EncAlgoAes256Cbc (bool): the EncAlgoAes256Cbc param. Default: `true`.

Param EncAlgoAes256Gcm (bool): the EncAlgoAes256Gcm param. Default: `true`.

Param EncAlgoChacha20Poly1305 (bool): the EncAlgoChacha20Poly1305 param. Default: `true`.

Param EncAlgoRc4 (bool): the EncAlgoRc4 param. Default: `true`.

Param KeyxchgAlgoDhe (bool): the KeyxchgAlgoDhe param. Default: `true`.

Param KeyxchgAlgoEcdhe (bool): the KeyxchgAlgoEcdhe param. Default: `true`.

Param KeyxchgAlgoRsa (bool): the KeyxchgAlgoRsa param. Default: `true`.

Param MaxVersion (string): the MaxVersion param. String must be one of these: `"sslv3"`, `"tls1-0"`, `"tls1-1"`, `"tls1-2"`, `"tls1-3"`, `"max"`. Default: `"tls1-2"`.

Param MinVersion (string): the MinVersion param. String must be one of these: `"sslv3"`, `"tls1-0"`, `"tls1-1"`, `"tls1-2"`, `"tls1-3"`. Default: `"tls1-0"`.
*/
type SslProtocolSettingsObject struct {
	AuthAlgoMd5             *bool   `json:"auth_algo_md5,omitempty"`
	AuthAlgoSha1            *bool   `json:"auth_algo_sha1,omitempty"`
	AuthAlgoSha256          *bool   `json:"auth_algo_sha256,omitempty"`
	AuthAlgoSha384          *bool   `json:"auth_algo_sha384,omitempty"`
	EncAlgo3des             *bool   `json:"enc_algo_3des,omitempty"`
	EncAlgoAes128Cbc        *bool   `json:"enc_algo_aes_128_cbc,omitempty"`
	EncAlgoAes128Gcm        *bool   `json:"enc_algo_aes_128_gcm,omitempty"`
	EncAlgoAes256Cbc        *bool   `json:"enc_algo_aes_256_cbc,omitempty"`
	EncAlgoAes256Gcm        *bool   `json:"enc_algo_aes_256_gcm,omitempty"`
	EncAlgoChacha20Poly1305 *bool   `json:"enc_algo_chacha20_poly1305,omitempty"`
	EncAlgoRc4              *bool   `json:"enc_algo_rc4,omitempty"`
	KeyxchgAlgoDhe          *bool   `json:"keyxchg_algo_dhe,omitempty"`
	KeyxchgAlgoEcdhe        *bool   `json:"keyxchg_algo_ecdhe,omitempty"`
	KeyxchgAlgoRsa          *bool   `json:"keyxchg_algo_rsa,omitempty"`
	MaxVersion              *string `json:"max_version,omitempty"`
	MinVersion              *string `json:"min_version,omitempty"`
}
