package sequences

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/authentication-sequences

/*
Config object.

ShortName: zLXjrfn
Parent chains:
*

Args:

Param AuthenticationProfiles ([]string): the AuthenticationProfiles param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): the Name param.

Param UseDomainFindProfile (bool): the UseDomainFindProfile param. Default: `true`.
*/
type Config struct {
	AuthenticationProfiles []string `json:"authentication_profiles,omitempty"`
	Id                     *string  `json:"id,omitempty"`
	Name                   string   `json:"name"`
	UseDomainFindProfile   *bool    `json:"use_domain_find_profile,omitempty"`
}
