package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/dns-security-profiles

/*
Config object.

ShortName: irQawLY
Parent chains:
*

Args:

Param BotnetDomains (BotnetDomainsObject): the BotnetDomains param.

Param Description (string): the Description param.

Param Id (string, read-only): UUID of the resource

Param Name (string): the Name param.
*/
type Config struct {
	BotnetDomains *BotnetDomainsObject `json:"botnet_domains,omitempty"`
	Description   *string              `json:"description,omitempty"`
	Id            *string              `json:"id,omitempty"`
	Name          *string              `json:"name,omitempty"`
}

/*
BotnetDomainsObject object.

ShortName:
Parent chains:
*
* botnet_domains

Args:

Param DnsSecurityCategories ([]DnsSecurityCategoryObject): the DnsSecurityCategories param.

Param Lists ([]ListObject): the Lists param.

Param Sinkhole (SinkholeObject): the Sinkhole param.

Param Whitelists ([]WhitelistObject): the Whitelists param.
*/
type BotnetDomainsObject struct {
	DnsSecurityCategories []DnsSecurityCategoryObject `json:"dns_security_categories,omitempty"`
	Lists                 []ListObject                `json:"lists,omitempty"`
	Sinkhole              *SinkholeObject             `json:"sinkhole,omitempty"`
	Whitelists            []WhitelistObject           `json:"whitelist,omitempty"`
}

/*
DnsSecurityCategoryObject object.

ShortName:
Parent chains:
*
* botnet_domains
* dns_security_categories
* _inline

Args:

Param Action (string): the Action param. String must be one of these: `"default"`, `"allow"`, `"block"`, `"sinkhole"`. Default: `"default"`.

Param LogLevel (string): the LogLevel param. String must be one of these: `"default"`, `"none"`, `"low"`, `"informational"`, `"medium"`, `"high"`, `"critical"`. Default: `"default"`.

Param Name (string): the Name param.

Param PacketCapture (string): the PacketCapture param. String must be one of these: `"disable"`, `"single-packet"`, `"extended-capture"`.
*/
type DnsSecurityCategoryObject struct {
	Action        *string `json:"action,omitempty"`
	LogLevel      *string `json:"log_level,omitempty"`
	Name          *string `json:"name,omitempty"`
	PacketCapture *string `json:"packet_capture,omitempty"`
}

/*
ListObject object.

ShortName:
Parent chains:
*
* botnet_domains
* lists
* _inline

Args:

Param Action (ActionObject): the Action param.

Param Name (string, required): the Name param.

Param PacketCapture (string): the PacketCapture param. String must be one of these: `"disable"`, `"single-packet"`, `"extended-capture"`.
*/
type ListObject struct {
	Action        *ActionObject `json:"action,omitempty"`
	Name          string        `json:"name"`
	PacketCapture *string       `json:"packet_capture,omitempty"`
}

/*
ActionObject object.

ShortName:
Parent chains:
*
* botnet_domains
* lists
* _inline
* action

Args:

Param Alert (any): the Alert param.

Param Allow (any): the Allow param.

Param Block (any): the Block param.

Param Sinkhole (any): the Sinkhole param.

NOTE:  One of the following params should be specified:
  - Alert
  - Allow
  - Block
  - Sinkhole
*/
type ActionObject struct {
	Alert    any `json:"alert,omitempty"`
	Allow    any `json:"allow,omitempty"`
	Block    any `json:"block,omitempty"`
	Sinkhole any `json:"sinkhole,omitempty"`
}

/*
SinkholeObject object.

ShortName:
Parent chains:
*
* botnet_domains
* sinkhole

Args:

Param Ipv4Address (string): the Ipv4Address param. String must be one of these: `"127.0.0.1"`, `"pan-sinkhole-default-ip"`.

Param Ipv6Address (string): the Ipv6Address param. String must be one of these: `"::1"`.
*/
type SinkholeObject struct {
	Ipv4Address *string `json:"ipv4_address,omitempty"`
	Ipv6Address *string `json:"ipv6_address,omitempty"`
}

/*
WhitelistObject object.

ShortName:
Parent chains:
*
* botnet_domains
* whitelist
* _inline

Args:

Param Description (string): the Description param.

Param Name (string, required): the Name param.
*/
type WhitelistObject struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
}
