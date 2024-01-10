package signatures

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/anti-spyware-signatures

/*
Config object.

ShortName: eumQbRC
Parent chains:
*

Args:

Param Bugtraqs ([]string): the Bugtraqs param.

Param Comment (string): the Comment param. String length must not exceed 256 characters.

Param Cves ([]string): the Cves param.

Param DefaultAction (DefaultActionObject): the DefaultAction param.

Param Direction (string): the Direction param. String must be one of these: `"client2server"`, `"server2client"`, `"both"`.

Param Id (string, read-only): UUID of the resource

Param References ([]string): the References param.

Param Severity (string): the Severity param. String must be one of these: `"critical"`, `"low"`, `"high"`, `"medium"`, `"informational"`.

Param Signature (SignatureObject): the Signature param.

Param ThreatId (int64, required): threat id range <15000-18000> and <6900001-7000000> Value must be between 15000 and 70000000.

Param Threatname (string, required): the Threatname param. String length must not exceed 1024 characters.

Param Vendors ([]string): the Vendors param.
*/
type Config struct {
	Bugtraqs      []string             `json:"bugtraq,omitempty"`
	Comment       *string              `json:"comment,omitempty"`
	Cves          []string             `json:"cve,omitempty"`
	DefaultAction *DefaultActionObject `json:"default_action,omitempty"`
	Direction     *string              `json:"direction,omitempty"`
	Id            *string              `json:"id,omitempty"`
	References    []string             `json:"reference,omitempty"`
	Severity      *string              `json:"severity,omitempty"`
	Signature     *SignatureObject     `json:"signature,omitempty"`
	ThreatId      int64                `json:"threat_id"`
	Threatname    string               `json:"threatname"`
	Vendors       []string             `json:"vendor,omitempty"`
}

/*
DefaultActionObject object.

ShortName:
Parent chains:
*
* default_action

Args:

Param Alert (any): the Alert param. Default: `false`.

Param Allow (any): the Allow param. Default: `false`.

Param BlockIp (BlockIpObject): the BlockIp param.

Param Drop (any): the Drop param. Default: `false`.

Param ResetBoth (any): the ResetBoth param. Default: `false`.

Param ResetClient (any): the ResetClient param. Default: `false`.

Param ResetServer (any): the ResetServer param. Default: `false`.

NOTE:  One of the following params should be specified:
  - Allow
  - Alert
  - Drop
  - ResetClient
  - ResetServer
  - ResetBoth
  - BlockIp
*/
type DefaultActionObject struct {
	Alert       any            `json:"alert,omitempty"`
	Allow       any            `json:"allow,omitempty"`
	BlockIp     *BlockIpObject `json:"block_ip,omitempty"`
	Drop        any            `json:"drop,omitempty"`
	ResetBoth   any            `json:"reset_both,omitempty"`
	ResetClient any            `json:"reset_client,omitempty"`
	ResetServer any            `json:"reset_server,omitempty"`
}

/*
BlockIpObject object.

ShortName:
Parent chains:
*
* default_action
* block_ip

Args:

Param Duration (int64): the Duration param. Value must be between 1 and 3600.

Param TrackBy (string): the TrackBy param. String must be one of these: `"source-and-destination"`, `"source"`.
*/
type BlockIpObject struct {
	Duration *int64  `json:"duration,omitempty"`
	TrackBy  *string `json:"track_by,omitempty"`
}

/*
SignatureObject object.

ShortName:
Parent chains:
*
* signature

Args:

Param Combination (CombinationObject): the Combination param.

Param Standards ([]StandardObject): the Standards param.

NOTE:  One of the following params should be specified:
  - Combination
  - Standards
*/
type SignatureObject struct {
	Combination *CombinationObject `json:"combination,omitempty"`
	Standards   []StandardObject   `json:"standard,omitempty"`
}

/*
CombinationObject object.

ShortName:
Parent chains:
*
* signature
* combination

Args:

Param AndConditions ([]CombinationAndConditionObject): the AndConditions param.

Param OrderFree (bool): the OrderFree param. Default: `false`.

Param TimeAttribute (TimeAttributeObject): the TimeAttribute param.
*/
type CombinationObject struct {
	AndConditions []CombinationAndConditionObject `json:"and_condition,omitempty"`
	OrderFree     *bool                           `json:"order_free,omitempty"`
	TimeAttribute *TimeAttributeObject            `json:"time_attribute,omitempty"`
}

/*
CombinationAndConditionObject object.

ShortName:
Parent chains:
*
* signature
* combination
* and_condition
* _inline

Args:

Param Name (string): the Name param.

Param OrConditions ([]CombinationOrConditionObject): the OrConditions param.
*/
type CombinationAndConditionObject struct {
	Name         *string                        `json:"name,omitempty"`
	OrConditions []CombinationOrConditionObject `json:"or_condition,omitempty"`
}

/*
CombinationOrConditionObject object.

ShortName:
Parent chains:
*
* signature
* combination
* and_condition
* _inline
* or_condition
* _inline

Args:

Param Name (string): the Name param.

Param ThreatId (string): the ThreatId param.
*/
type CombinationOrConditionObject struct {
	Name     *string `json:"name,omitempty"`
	ThreatId *string `json:"threat_id,omitempty"`
}

/*
TimeAttributeObject object.

ShortName:
Parent chains:
*
* signature
* combination
* time_attribute

Args:

Param Interval (int64): the Interval param. Value must be between 1 and 3600.

Param Threshold (int64): the Threshold param. Value must be between 1 and 255.

Param TrackBy (string): the TrackBy param. String must be one of these: `"source-and-destination"`, `"source"`, `"destination"`.
*/
type TimeAttributeObject struct {
	Interval  *int64  `json:"interval,omitempty"`
	Threshold *int64  `json:"threshold,omitempty"`
	TrackBy   *string `json:"track_by,omitempty"`
}

/*
StandardObject object.

ShortName:
Parent chains:
*
* signature
* standard
* _inline

Args:

Param AndConditions ([]StandardAndConditionObject): the AndConditions param.

Param Comment (string): the Comment param. String length must not exceed 256 characters.

Param Name (string, required): the Name param.

Param OrderFree (bool): the OrderFree param. Default: `false`.

Param Scope (string): the Scope param. String must be one of these: `"protocol-data-unit"`, `"session"`.
*/
type StandardObject struct {
	AndConditions []StandardAndConditionObject `json:"and_condition,omitempty"`
	Comment       *string                      `json:"comment,omitempty"`
	Name          string                       `json:"name"`
	OrderFree     *bool                        `json:"order_free,omitempty"`
	Scope         *string                      `json:"scope,omitempty"`
}

/*
StandardAndConditionObject object.

ShortName:
Parent chains:
*
* signature
* standard
* _inline
* and_condition
* _inline

Args:

Param Name (string): the Name param.

Param OrConditions ([]StandardOrConditionObject): the OrConditions param.
*/
type StandardAndConditionObject struct {
	Name         *string                     `json:"name,omitempty"`
	OrConditions []StandardOrConditionObject `json:"or_condition,omitempty"`
}

/*
StandardOrConditionObject object.

ShortName:
Parent chains:
*
* signature
* standard
* _inline
* and_condition
* _inline
* or_condition
* _inline

Args:

Param Name (string): the Name param.

Param Operator (OperatorObject): the Operator param.
*/
type StandardOrConditionObject struct {
	Name     *string         `json:"name,omitempty"`
	Operator *OperatorObject `json:"operator,omitempty"`
}

/*
OperatorObject object.

ShortName:
Parent chains:
*
* signature
* standard
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
* standard
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* equal_to

Args:

Param Context (string): the Context param.

Param Negate (bool): the Negate param. Default: `false`.

Param Qualifiers ([]EqualToQualifierObject): the Qualifiers param.

Param Value (int64): the Value param. Value must be between 0 and 4294967295.
*/
type EqualToObject struct {
	Context    *string                  `json:"context,omitempty"`
	Negate     *bool                    `json:"negate,omitempty"`
	Qualifiers []EqualToQualifierObject `json:"qualifier,omitempty"`
	Value      *int64                   `json:"value,omitempty"`
}

/*
EqualToQualifierObject object.

ShortName:
Parent chains:
*
* signature
* standard
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* equal_to
* qualifier
* _inline

Args:

Param Name (string): the Name param.

Param Value (string): the Value param.
*/
type EqualToQualifierObject struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

/*
GreaterThanObject object.

ShortName:
Parent chains:
*
* signature
* standard
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* greater_than

Args:

Param Context (string): the Context param.

Param Qualifiers ([]GreaterThanQualifierObject): the Qualifiers param.

Param Value (int64): the Value param. Value must be between 0 and 4294967295.
*/
type GreaterThanObject struct {
	Context    *string                      `json:"context,omitempty"`
	Qualifiers []GreaterThanQualifierObject `json:"qualifier,omitempty"`
	Value      *int64                       `json:"value,omitempty"`
}

/*
GreaterThanQualifierObject object.

ShortName:
Parent chains:
*
* signature
* standard
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

Param Name (string): the Name param.

Param Value (string): the Value param.
*/
type GreaterThanQualifierObject struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

/*
LessThanObject object.

ShortName:
Parent chains:
*
* signature
* standard
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* less_than

Args:

Param Context (string): the Context param.

Param Qualifiers ([]LessThanQualifierObject): the Qualifiers param.

Param Value (int64): the Value param. Value must be between 0 and 4294967295.
*/
type LessThanObject struct {
	Context    *string                   `json:"context,omitempty"`
	Qualifiers []LessThanQualifierObject `json:"qualifier,omitempty"`
	Value      *int64                    `json:"value,omitempty"`
}

/*
LessThanQualifierObject object.

ShortName:
Parent chains:
*
* signature
* standard
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

Param Name (string): the Name param.

Param Value (string): the Value param.
*/
type LessThanQualifierObject struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

/*
PatternMatchObject object.

ShortName:
Parent chains:
*
* signature
* standard
* _inline
* and_condition
* _inline
* or_condition
* _inline
* operator
* pattern_match

Args:

Param Context (string): the Context param.

Param Negate (bool): the Negate param. Default: `false`.

Param Pattern (string): the Pattern param.

Param Qualifiers ([]PatternMatchQualifierObject): the Qualifiers param.
*/
type PatternMatchObject struct {
	Context    *string                       `json:"context,omitempty"`
	Negate     *bool                         `json:"negate,omitempty"`
	Pattern    *string                       `json:"pattern,omitempty"`
	Qualifiers []PatternMatchQualifierObject `json:"qualifier,omitempty"`
}

/*
PatternMatchQualifierObject object.

ShortName:
Parent chains:
*
* signature
* standard
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

Param Name (string): the Name param.

Param Value (string): the Value param.
*/
type PatternMatchQualifierObject struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
