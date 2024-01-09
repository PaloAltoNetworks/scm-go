package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/qos-profiles

/*
Config object.

ShortName: zGfKFAQ
Parent chains:
*

Args:

Param AggregateBandwidth (AggregateBandwidthObject): the AggregateBandwidth param.

Param ClassBandwidthType (ClassBandwidthTypeObject): the ClassBandwidthType param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string begin with letter: [0-9a-zA-Z._-] String length must not exceed 31 characters.
*/
type Config struct {
	AggregateBandwidth *AggregateBandwidthObject `json:"aggregate_bandwidth,omitempty"`
	ClassBandwidthType *ClassBandwidthTypeObject `json:"class_bandwidth_type,omitempty"`
	Id                 *string                   `json:"id,omitempty"`
	Name               string                    `json:"name"`
}

/*
AggregateBandwidthObject object.

ShortName:
Parent chains:
*
* aggregate_bandwidth

Args:

Param EgressGuaranteed (int64): guaranteed sending bandwidth in mbps Value must be between 0 and 16000.

Param EgressMax (int64): max sending bandwidth in mbps Value must be between 0 and 60000.
*/
type AggregateBandwidthObject struct {
	EgressGuaranteed *int64 `json:"egress_guaranteed,omitempty"`
	EgressMax        *int64 `json:"egress_max,omitempty"`
}

/*
ClassBandwidthTypeObject object.

ShortName:
Parent chains:
*
* class_bandwidth_type

Args:

Param Mbps (MbpsObject): the Mbps param.

Param Percentage (PercentageObject): the Percentage param.

NOTE:  One of the following params should be specified:
  - Mbps
  - Percentage
*/
type ClassBandwidthTypeObject struct {
	Mbps       *MbpsObject       `json:"mbps,omitempty"`
	Percentage *PercentageObject `json:"percentage,omitempty"`
}

/*
MbpsObject object.

ShortName:
Parent chains:
*
* class_bandwidth_type
* mbps

Args:

Param Classes ([]MbpsClassObject): QoS setting for traffic classes
*/
type MbpsObject struct {
	Classes []MbpsClassObject `json:"class,omitempty"`
}

/*
MbpsClassObject object.

ShortName:
Parent chains:
*
* class_bandwidth_type
* mbps
* class
* _inline

Args:

Param ClassBandwidth (MbpsBandwidthObject): the ClassBandwidth param.

Param Name (string): Traffic class String length must not exceed 31 characters.

Param Priority (string): traffic class priority String must be one of these: `"real-time"`, `"high"`, `"medium"`, `"low"`. Default: `"medium"`.
*/
type MbpsClassObject struct {
	ClassBandwidth *MbpsBandwidthObject `json:"class_bandwidth,omitempty"`
	Name           *string              `json:"name,omitempty"`
	Priority       *string              `json:"priority,omitempty"`
}

/*
MbpsBandwidthObject object.

ShortName:
Parent chains:
*
* class_bandwidth_type
* mbps
* class
* _inline
* class_bandwidth

Args:

Param EgressGuaranteed (int64): guaranteed sending bandwidth in mbps Value must be between 0 and 60000.

Param EgressMax (int64): max sending bandwidth in mbps Value must be between 0 and 60000.
*/
type MbpsBandwidthObject struct {
	EgressGuaranteed *int64 `json:"egress_guaranteed,omitempty"`
	EgressMax        *int64 `json:"egress_max,omitempty"`
}

/*
PercentageObject object.

ShortName:
Parent chains:
*
* class_bandwidth_type
* percentage

Args:

Param Classes ([]PercentageClassObject): QoS setting for traffic classes
*/
type PercentageObject struct {
	Classes []PercentageClassObject `json:"class,omitempty"`
}

/*
PercentageClassObject object.

ShortName:
Parent chains:
*
* class_bandwidth_type
* percentage
* class
* _inline

Args:

Param ClassBandwidth (PercentageBandwidthObject): the ClassBandwidth param.

Param Name (string): Traffic class String length must not exceed 31 characters.

Param Priority (string): traffic class priority String must be one of these: `"real-time"`, `"high"`, `"medium"`, `"low"`. Default: `"medium"`.
*/
type PercentageClassObject struct {
	ClassBandwidth *PercentageBandwidthObject `json:"class_bandwidth,omitempty"`
	Name           *string                    `json:"name,omitempty"`
	Priority       *string                    `json:"priority,omitempty"`
}

/*
PercentageBandwidthObject object.

ShortName:
Parent chains:
*
* class_bandwidth_type
* percentage
* class
* _inline
* class_bandwidth

Args:

Param EgressGuaranteed (int64): guaranteed sending bandwidth in percentage Value must be between 0 and 100.

Param EgressMax (int64): max sending bandwidth in percentage Value must be between 0 and 100.
*/
type PercentageBandwidthObject struct {
	EgressGuaranteed *int64 `json:"egress_guaranteed,omitempty"`
	EgressMax        *int64 `json:"egress_max,omitempty"`
}
