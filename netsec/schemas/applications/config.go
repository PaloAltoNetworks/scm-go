package applications

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/applications

/*
Config object.

ShortName: yJkkSzS
Parent chains:
*

Args:

Param AbleToTransferFile (bool): the AbleToTransferFile param.

Param AlgDisableCapability (string): the AlgDisableCapability param. String length must not exceed 127 characters.

Param Category (string, required): the Category param.

Param ConsumeBigBandwidth (bool): the ConsumeBigBandwidth param.

Param DataIdent (bool): the DataIdent param.

Param Default (DefaultObject): the Default param.

Param Description (string): the Description param. String length must not exceed 1023 characters.

Param EvasiveBehavior (bool): the EvasiveBehavior param.

Param FileTypeIdent (bool): the FileTypeIdent param.

Param HasKnownVulnerability (bool): the HasKnownVulnerability param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param NoAppidCaching (bool): the NoAppidCaching param.

Param ParentApp (string): the ParentApp param. String length must not exceed 127 characters.

Param PervasiveUse (bool): the PervasiveUse param.

Param ProneToMisuse (bool): the ProneToMisuse param.

Param Risk (int64, required): the Risk param. Value must be between 1 and 5.

Param Signatures ([]SignatureObject): the Signatures param.

Param Subcategory (string, required): the Subcategory param. String length must not exceed 63 characters.

Param TcpHalfClosedTimeout (int64): timeout for half-close session in seconds Value must be between 1 and 604800.

Param TcpTimeWaitTimeout (int64): timeout for session in time_wait state in seconds Value must be between 1 and 600.

Param TcpTimeout (int64): timeout in seconds Value must be between 0 and 604800.

Param Technology (string, required): the Technology param. String length must not exceed 63 characters.

Param Timeout (int64): timeout in seconds Value must be between 0 and 604800.

Param TunnelApplications (bool): the TunnelApplications param.

Param TunnelOtherApplication (bool): the TunnelOtherApplication param.

Param UdpTimeout (int64): timeout in seconds Value must be between 0 and 604800.

Param UsedByMalware (bool): the UsedByMalware param.

Param VirusIdent (bool): the VirusIdent param.
*/
type Config struct {
	AbleToTransferFile     *bool             `json:"able_to_transfer_file,omitempty"`
	AlgDisableCapability   *string           `json:"alg_disable_capability,omitempty"`
	Category               string            `json:"category"`
	ConsumeBigBandwidth    *bool             `json:"consume_big_bandwidth,omitempty"`
	DataIdent              *bool             `json:"data_ident,omitempty"`
	Default                *DefaultObject    `json:"default,omitempty"`
	Description            *string           `json:"description,omitempty"`
	EvasiveBehavior        *bool             `json:"evasive_behavior,omitempty"`
	FileTypeIdent          *bool             `json:"file_type_ident,omitempty"`
	HasKnownVulnerability  *bool             `json:"has_known_vulnerability,omitempty"`
	Id                     *string           `json:"id,omitempty"`
	Name                   string            `json:"name"`
	NoAppidCaching         *bool             `json:"no_appid_caching,omitempty"`
	ParentApp              *string           `json:"parent_app,omitempty"`
	PervasiveUse           *bool             `json:"pervasive_use,omitempty"`
	ProneToMisuse          *bool             `json:"prone_to_misuse,omitempty"`
	Risk                   int64             `json:"risk"`
	Signatures             []SignatureObject `json:"signature,omitempty"`
	Subcategory            string            `json:"subcategory"`
	TcpHalfClosedTimeout   *int64            `json:"tcp_half_closed_timeout,omitempty"`
	TcpTimeWaitTimeout     *int64            `json:"tcp_time_wait_timeout,omitempty"`
	TcpTimeout             *int64            `json:"tcp_timeout,omitempty"`
	Technology             string            `json:"technology"`
	Timeout                *int64            `json:"timeout,omitempty"`
	TunnelApplications     *bool             `json:"tunnel_applications,omitempty"`
	TunnelOtherApplication *bool             `json:"tunnel_other_application,omitempty"`
	UdpTimeout             *int64            `json:"udp_timeout,omitempty"`
	UsedByMalware          *bool             `json:"used_by_malware,omitempty"`
	VirusIdent             *bool             `json:"virus_ident,omitempty"`
}

/*
DefaultObject object.

ShortName:
Parent chains:
*
* default

Args:

Param IdentByIcmp6Type (IdentByIcmp6TypeObject): the IdentByIcmp6Type param.

Param IdentByIcmpType (IdentByIcmpTypeObject): the IdentByIcmpType param.

Param IdentByIpProtocol (string): the IdentByIpProtocol param.

Param Ports ([]string): the Ports param. Individual elements in this list are subject to additional validation. String length must not exceed 63 characters.

NOTE:  One of the following params should be specified:
  - Ports
  - IdentByIpProtocol
  - IdentByIcmpType
  - IdentByIcmp6Type
*/
type DefaultObject struct {
	IdentByIcmp6Type  *IdentByIcmp6TypeObject `json:"ident_by_icmp6_type,omitempty"`
	IdentByIcmpType   *IdentByIcmpTypeObject  `json:"ident_by_icmp_type,omitempty"`
	IdentByIpProtocol *string                 `json:"ident_by_ip_protocol,omitempty"`
	Ports             []string                `json:"port,omitempty"`
}

/*
IdentByIcmp6TypeObject object.

ShortName:
Parent chains:
*
* default
* ident_by_icmp6_type

Args:

Param Code (string): the Code param.

Param Type (string, required): the Type param.
*/
type IdentByIcmp6TypeObject struct {
	Code *string `json:"code,omitempty"`
	Type string  `json:"type"`
}

/*
IdentByIcmpTypeObject object.

ShortName:
Parent chains:
*
* default
* ident_by_icmp_type

Args:

Param Code (string): the Code param.

Param Type (string, required): the Type param.
*/
type IdentByIcmpTypeObject struct {
	Code *string `json:"code,omitempty"`
	Type string  `json:"type"`
}

/*
SignatureObject object.

ShortName:
Parent chains:
*
* signature
* _inline

Args:

Param AndConditions ([]AndConditionObject): the AndConditions param.

Param Comment (string): the Comment param. String length must not exceed 256 characters.

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param OrderFree (bool): the OrderFree param. Default: `false`.

Param Scope (string): the Scope param. String must be one of these: `"protocol-data-unit"`, `"session"`. Default: `"protocol-data-unit"`.
*/
type SignatureObject struct {
	AndConditions []AndConditionObject `json:"and_condition,omitempty"`
	Comment       *string              `json:"comment,omitempty"`
	Name          string               `json:"name"`
	OrderFree     *bool                `json:"order_free,omitempty"`
	Scope         *string              `json:"scope,omitempty"`
}

/*
AndConditionObject object.

ShortName:
Parent chains:
*
* signature
* _inline
* and_condition
* _inline

Args:

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param OrConditions ([]OrConditionObject): the OrConditions param.
*/
type AndConditionObject struct {
	Name         string              `json:"name"`
	OrConditions []OrConditionObject `json:"or_condition,omitempty"`
}

/*
OrConditionObject object.

ShortName:
Parent chains:
*
* signature
* _inline
* and_condition
* _inline
* or_condition
* _inline

Args:

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param Operator (OperatorObject, required): the Operator param.
*/
type OrConditionObject struct {
	Name     string         `json:"name"`
	Operator OperatorObject `json:"operator"`
}

/*
OperatorObject object.

ShortName:
Parent chains:
*
* signature
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator

Args:

Param EqualTo (EqualToObject): the EqualTo param.

Param GreaterThan (GreaterThanObject): the GreaterThan param.

Param LessThan (LessThanObject): the LessThan param.

Param PatternMatch (PatternMatchObject): the PatternMatch param.

NOTE:  One of the following params should be specified:
  - PatternMatch
  - GreaterThan
  - LessThan
  - EqualTo
*/
type OperatorObject struct {
	EqualTo      *EqualToObject      `json:"equal_to,omitempty"`
	GreaterThan  *GreaterThanObject  `json:"greater_than,omitempty"`
	LessThan     *LessThanObject     `json:"less_than,omitempty"`
	PatternMatch *PatternMatchObject `json:"pattern_match,omitempty"`
}

/*
EqualToObject object.

ShortName:
Parent chains:
*
* signature
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* equal_to

Args:

Param Context (string, required): the Context param.

Param Mask (string): 4-byte hex value String length must not exceed 10 characters. String validation regex: `^[0][xX][0-9A-Fa-f]{8}$`.

Param Position (string): the Position param. String length must not exceed 127 characters.

Param Value (string, required): the Value param. String length must not exceed 10 characters.
*/
type EqualToObject struct {
	Context  string  `json:"context"`
	Mask     *string `json:"mask,omitempty"`
	Position *string `json:"position,omitempty"`
	Value    string  `json:"value"`
}

/*
GreaterThanObject object.

ShortName:
Parent chains:
*
* signature
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* greater_than

Args:

Param Context (string, required): the Context param. String length must not exceed 127 characters.

Param Qualifiers ([]GreaterThanQualifierObject): the Qualifiers param.

Param Value (int64, required): the Value param. Value must be between 0 and 4294967295.
*/
type GreaterThanObject struct {
	Context    string                       `json:"context"`
	Qualifiers []GreaterThanQualifierObject `json:"qualifier,omitempty"`
	Value      int64                        `json:"value"`
}

/*
GreaterThanQualifierObject object.

ShortName:
Parent chains:
*
* signature
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* greater_than
* qualifier
* _inline

Args:

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param Value (string, required): the Value param.
*/
type GreaterThanQualifierObject struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
LessThanObject object.

ShortName:
Parent chains:
*
* signature
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* less_than

Args:

Param Context (string, required): the Context param. String length must not exceed 127 characters.

Param Qualifiers ([]LessThanQualifierObject): the Qualifiers param.

Param Value (int64, required): the Value param. Value must be between 0 and 4294967295.
*/
type LessThanObject struct {
	Context    string                    `json:"context"`
	Qualifiers []LessThanQualifierObject `json:"qualifier,omitempty"`
	Value      int64                     `json:"value"`
}

/*
LessThanQualifierObject object.

ShortName:
Parent chains:
*
* signature
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* less_than
* qualifier
* _inline

Args:

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param Value (string, required): the Value param.
*/
type LessThanQualifierObject struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
PatternMatchObject object.

ShortName:
Parent chains:
*
* signature
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* pattern_match

Args:

Param Context (string, required): the Context param. String length must not exceed 127 characters.

Param Pattern (string, required): the Pattern param. String length must not exceed 127 characters.

Param Qualifiers ([]PatternMatchQualifierObject): the Qualifiers param.
*/
type PatternMatchObject struct {
	Context    string                        `json:"context"`
	Pattern    string                        `json:"pattern"`
	Qualifiers []PatternMatchQualifierObject `json:"qualifier,omitempty"`
}

/*
PatternMatchQualifierObject object.

ShortName:
Parent chains:
*
* signature
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* pattern_match
* qualifier
* _inline

Args:

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param Value (string, required): the Value param.
*/
type PatternMatchQualifierObject struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
