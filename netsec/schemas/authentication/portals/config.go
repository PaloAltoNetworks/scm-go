package portals

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/authentication-portals

/*
Config object.

ShortName: hhIWLbI
Parent chains:
*

Args:

Param AuthenticationProfile (string): the AuthenticationProfile param.

Param CertificateProfile (string): the CertificateProfile param.

Param GpUdpPort (int64): the GpUdpPort param. Value must be between 1 and 65535.

Param Id (string, read-only): UUID of the resource

Param IdleTimer (int64): the IdleTimer param. Value must be between 1 and 1440.

Param RedirectHost (string): the RedirectHost param.

Param Timer (int64): the Timer param. Value must be between 1 and 1440.

Param TlsServiceProfile (string): the TlsServiceProfile param.
*/
type Config struct {
	AuthenticationProfile *string `json:"authentication_profile,omitempty"`
	CertificateProfile    *string `json:"certificate_profile,omitempty"`
	GpUdpPort             *int64  `json:"gp_udp_port,omitempty"`
	Id                    *string `json:"id,omitempty"`
	IdleTimer             *int64  `json:"idle_timer,omitempty"`
	RedirectHost          *string `json:"redirect_host,omitempty"`
	Timer                 *int64  `json:"timer,omitempty"`
	TlsServiceProfile     *string `json:"tls_service_profile,omitempty"`
}
