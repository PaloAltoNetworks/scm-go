package allocations

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/bandwidth-allocations

/*
Config object.

ShortName: lrzxLXR
Parent chains:
*

Args:

Param AllocatedBandwidth (int64, required): bandwidth to allocate in Mbps

Param Name (string, required): name of the aggregated bandwidth region

Param Qos (QosObject): the Qos param.

Param SpnNameList ([]string): the SpnNameList param.
*/
type Config struct {
	AllocatedBandwidth int64      `json:"allocated_bandwidth"`
	Name               string     `json:"name"`
	Qos                *QosObject `json:"qos,omitempty"`
	SpnNameList        []string   `json:"spn_name_list,omitempty"`
}

/*
QosObject object.

ShortName:
Parent chains:
*
* qos

Args:

Param Customized (bool): the Customized param.

Param Enabled (bool): the Enabled param.

Param GuaranteedRatio (int64): the GuaranteedRatio param.

Param Profile (string): the Profile param.
*/
type QosObject struct {
	Customized      *bool   `json:"customized,omitempty"`
	Enabled         *bool   `json:"enabled,omitempty"`
	GuaranteedRatio *int64  `json:"guaranteed_ratio,omitempty"`
	Profile         *string `json:"profile,omitempty"`
}
