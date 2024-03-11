package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/anti-spyware-profiles

/*
Config object.

ShortName: lrzxLXR
Parent chains:
*

Args:

Param CloudInlineAnalysis (bool): the CloudInlineAnalysis param. Default: `false`.

Param Description (string): the Description param.

Param Id (string, read-only): UUID of the resource

Param InlineExceptionEdlUrls ([]string): the InlineExceptionEdlUrls param.

Param InlineExceptionIpAddresses ([]string): the InlineExceptionIpAddresses param.

Param MicaEngineSpywareEnabledList ([]MicaEngineObject): the MicaEngineSpywareEnabledList param.

Param Name (string, required): the Name param.

Param Rules ([]RuleObject): the Rules param.

Param ThreatExceptions ([]ThreatExceptionObject): the ThreatExceptions param.
*/
type Config struct {
	CloudInlineAnalysis          *bool                   `json:"cloud_inline_analysis,omitempty"`
	Description                  *string                 `json:"description,omitempty"`
	Id                           *string                 `json:"id,omitempty"`
	InlineExceptionEdlUrls       []string                `json:"inline_exception_edl_url,omitempty"`
	InlineExceptionIpAddresses   []string                `json:"inline_exception_ip_address,omitempty"`
	MicaEngineSpywareEnabledList []MicaEngineObject      `json:"mica_engine_spyware_enabled,omitempty"`
	Name                         string                  `json:"name"`
	Rules                        []RuleObject            `json:"rules,omitempty"`
	ThreatExceptions             []ThreatExceptionObject `json:"threat_exception,omitempty"`
}

/*
MicaEngineObject object.

ShortName:
Parent chains:
*
* mica_engine_spyware_enabled
* _inline

Args:

Param InlinePolicyAction (string): the InlinePolicyAction param. String must be one of these: `"alert"`, `"allow"`, `"drop"`, `"reset-both"`, `"reset-client"`, `"reset-server"`. Default: `"alert"`.

Param Name (string): the Name param.
*/
type MicaEngineObject struct {
	InlinePolicyAction *string `json:"inline_policy_action,omitempty"`
	Name               *string `json:"name,omitempty"`
}

/*
RuleObject object.

ShortName:
Parent chains:
*
* rules
* _inline

Args:

Param Action (RuleActionObject): the Action param.

Param Category (string): the Category param. String must be one of these: `"dns-proxy"`, `"backdoor"`, `"data-theft"`, `"autogen"`, `"spyware"`, `"dns-security"`, `"downloader"`, `"dns-phishing"`, `"phishing-kit"`, `"cryptominer"`, `"hacktool"`, `"dns-benign"`, `"dns-wildfire"`, `"botnet"`, `"dns-grayware"`, `"inline-cloud-c2"`, `"keylogger"`, `"p2p-communication"`, `"domain-edl"`, `"webshell"`, `"command-and-control"`, `"dns-ddns"`, `"net-worm"`, `"any"`, `"tls-fingerprint"`, `"dns-new-domain"`, `"dns"`, `"fraud"`, `"dns-c2"`, `"adware"`, `"post-exploitation"`, `"dns-malware"`, `"browser-hijack"`, `"dns-parked"`.

Param Name (string): the Name param.

Param PacketCapture (string): the PacketCapture param. String must be one of these: `"disable"`, `"single-packet"`, `"extended-capture"`.

Param Severities ([]string): the Severities param.

Param ThreatName (string): the ThreatName param. String length must exceed 4 characters.
*/
type RuleObject struct {
	Action        *RuleActionObject `json:"action,omitempty"`
	Category      *string           `json:"category,omitempty"`
	Name          *string           `json:"name,omitempty"`
	PacketCapture *string           `json:"packet_capture,omitempty"`
	Severities    []string          `json:"severity,omitempty"`
	ThreatName    *string           `json:"threat_name,omitempty"`
}

/*
RuleActionObject object.

ShortName:
Parent chains:
*
* rules
* _inline
* action

Args:

Param Alert (any): the Alert param.

Param Allow (any): the Allow param.

Param BlockIp (RuleBlockIpObject): the BlockIp param.

Param Drop (any): the Drop param.

Param ResetBoth (any): the ResetBoth param.

Param ResetClient (any): the ResetClient param.

Param ResetServer (any): the ResetServer param.

NOTE:  One of the following params should be specified:
  - Allow
  - Alert
  - Drop
  - ResetClient
  - ResetServer
  - ResetBoth
  - BlockIp
*/
type RuleActionObject struct {
	Alert       any                `json:"alert,omitempty"`
	Allow       any                `json:"allow,omitempty"`
	BlockIp     *RuleBlockIpObject `json:"block_ip,omitempty"`
	Drop        any                `json:"drop,omitempty"`
	ResetBoth   any                `json:"reset_both,omitempty"`
	ResetClient any                `json:"reset_client,omitempty"`
	ResetServer any                `json:"reset_server,omitempty"`
}

/*
RuleBlockIpObject object.

ShortName:
Parent chains:
*
* rules
* _inline
* action
* block_ip

Args:

Param Duration (int64): the Duration param. Value must be between 1 and 3600.

Param TrackBy (string): the TrackBy param. String must be one of these: `"source-and-destination"`, `"source"`.
*/
type RuleBlockIpObject struct {
	Duration *int64  `json:"duration,omitempty"`
	TrackBy  *string `json:"track_by,omitempty"`
}

/*
ThreatExceptionObject object.

ShortName:
Parent chains:
*
* threat_exception
* _inline

Args:

Param Action (ThreatExceptionActionObject): the Action param.

Param ExemptIps ([]ExemptIpObject): the ExemptIps param.

Param Name (string): the Name param.

Param Notes (string): the Notes param.

Param PacketCapture (string): the PacketCapture param. String must be one of these: `"disable"`, `"single-packet"`, `"extended-capture"`.
*/
type ThreatExceptionObject struct {
	Action        *ThreatExceptionActionObject `json:"action,omitempty"`
	ExemptIps     []ExemptIpObject             `json:"exempt_ip,omitempty"`
	Name          *string                      `json:"name,omitempty"`
	Notes         *string                      `json:"notes,omitempty"`
	PacketCapture *string                      `json:"packet_capture,omitempty"`
}

/*
ThreatExceptionActionObject object.

ShortName:
Parent chains:
*
* threat_exception
* _inline
* action

Args:

Param Alert (any): the Alert param.

Param Allow (any): the Allow param.

Param BlockIp (ThreatExceptionBlockIpObject): the BlockIp param.

Param Default (any): the Default param.

Param Drop (any): the Drop param.

Param ResetBoth (any): the ResetBoth param.

Param ResetClient (any): the ResetClient param.

Param ResetServer (any): the ResetServer param.

NOTE:  One of the following params should be specified:
  - Default
  - Allow
  - Alert
  - Drop
  - ResetClient
  - ResetServer
  - ResetBoth
  - BlockIp
*/
type ThreatExceptionActionObject struct {
	Alert       any                           `json:"alert,omitempty"`
	Allow       any                           `json:"allow,omitempty"`
	BlockIp     *ThreatExceptionBlockIpObject `json:"block_ip,omitempty"`
	Default     any                           `json:"default,omitempty"`
	Drop        any                           `json:"drop,omitempty"`
	ResetBoth   any                           `json:"reset_both,omitempty"`
	ResetClient any                           `json:"reset_client,omitempty"`
	ResetServer any                           `json:"reset_server,omitempty"`
}

/*
ThreatExceptionBlockIpObject object.

ShortName:
Parent chains:
*
* threat_exception
* _inline
* action
* block_ip

Args:

Param Duration (int64): the Duration param. Value must be between 1 and 3600.

Param TrackBy (string): the TrackBy param. String must be one of these: `"source-and-destination"`, `"source"`.
*/
type ThreatExceptionBlockIpObject struct {
	Duration *int64  `json:"duration,omitempty"`
	TrackBy  *string `json:"track_by,omitempty"`
}

/*
ExemptIpObject object.

ShortName:
Parent chains:
*
* threat_exception
* _inline
* exempt_ip
* _inline

Args:

Param Name (string, required): the Name param.
*/
type ExemptIpObject struct {
	Name string `json:"name"`
}
