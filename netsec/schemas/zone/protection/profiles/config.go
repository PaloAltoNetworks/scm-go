package profiles

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/zone-protection-profiles

/*
Config object.

ShortName: gVcGYNO
Parent chains:
*

Args:

Param AsymmetricPath (string): the AsymmetricPath param. String must be one of these: `"global"`, `"drop"`, `"bypass"`.

Param Description (string): the Description param.

Param Device (string): The device in which the resource is defined String length must not exceed 64 characters. String validation regex: `^[a-zA-Z\d-_\. ]+$`.

Param DiscardIcmpError (bool): the DiscardIcmpError param.

Param DiscardIcmpFrag (bool): the DiscardIcmpFrag param.

Param DiscardIcmpLargePacket (bool): the DiscardIcmpLargePacket param.

Param DiscardIcmpPingZeroId (bool): the DiscardIcmpPingZeroId param.

Param DiscardIpFrag (bool): the DiscardIpFrag param.

Param DiscardIpSpoof (bool): the DiscardIpSpoof param.

Param DiscardLooseSourceRouting (bool): the DiscardLooseSourceRouting param.

Param DiscardMalformedOption (bool): the DiscardMalformedOption param.

Param DiscardOverlappingTcpSegmentMismatch (bool): the DiscardOverlappingTcpSegmentMismatch param.

Param DiscardRecordRoute (bool): the DiscardRecordRoute param.

Param DiscardSecurity (bool): the DiscardSecurity param.

Param DiscardStreamId (bool): the DiscardStreamId param.

Param DiscardStrictSourceRouting (bool): the DiscardStrictSourceRouting param.

Param DiscardTcpSplitHandshake (bool): the DiscardTcpSplitHandshake param.

Param DiscardTcpSynWithData (bool): the DiscardTcpSynWithData param.

Param DiscardTcpSynackWithData (bool): the DiscardTcpSynackWithData param.

Param DiscardTimestamp (bool): the DiscardTimestamp param.

Param DiscardUnknownOption (bool): the DiscardUnknownOption param.

Param Flood (FloodObject): the Flood param.

Param Folder (string): The folder in which the resource is defined String length must not exceed 64 characters. String validation regex: `^[a-zA-Z\d-_\. ]+$`.

Param Id (string, read-only): the Id param.

Param Ipv6 (Ipv6Object): the Ipv6 param.

Param L2SecGroupTagProtection (L2SecGroupTagProtectionObject): the L2SecGroupTagProtection param.

Param Name (string, required): the Name param.

Param NonIpProtocol (NonIpProtocolObject): the NonIpProtocol param.

Param RemoveTcpTimestamp (bool): the RemoveTcpTimestamp param.

Param Scan ([]ScanObject): the Scan param.

Param ScanWhiteList ([]ScanWhiteListObject): the ScanWhiteList param.

Param Snippet (string): The snippet in which the resource is defined String length must not exceed 64 characters. String validation regex: `^[a-zA-Z\d-_\. ]+$`.

Param StrictIpCheck (bool): the StrictIpCheck param.

Param StripMptcpOption (string): the StripMptcpOption param. String must be one of these: `"global"`, `"yes"`, `"no"`.

Param StripTcpFastOpenAndData (bool): the StripTcpFastOpenAndData param.

Param SuppressIcmpNeedfrag (bool): the SuppressIcmpNeedfrag param.

Param SuppressIcmpTimeexceeded (bool): the SuppressIcmpTimeexceeded param.

Param TcpRejectNonSyn (string): the TcpRejectNonSyn param. String must be one of these: `"global"`, `"yes"`, `"no"`.

NOTE:  One of the following params should be specified:
  - Folder
  - Snippet
  - Device
*/
type Config struct {
	AsymmetricPath                       *string                        `json:"asymmetric_path,omitempty"`
	Description                          *string                        `json:"description,omitempty"`
	Device                               *string                        `json:"device,omitempty"`
	DiscardIcmpError                     *bool                          `json:"discard_icmp_error,omitempty"`
	DiscardIcmpFrag                      *bool                          `json:"discard_icmp_frag,omitempty"`
	DiscardIcmpLargePacket               *bool                          `json:"discard_icmp_large_packet,omitempty"`
	DiscardIcmpPingZeroId                *bool                          `json:"discard_icmp_ping_zero_id,omitempty"`
	DiscardIpFrag                        *bool                          `json:"discard_ip_frag,omitempty"`
	DiscardIpSpoof                       *bool                          `json:"discard_ip_spoof,omitempty"`
	DiscardLooseSourceRouting            *bool                          `json:"discard_loose_source_routing,omitempty"`
	DiscardMalformedOption               *bool                          `json:"discard_malformed_option,omitempty"`
	DiscardOverlappingTcpSegmentMismatch *bool                          `json:"discard_overlapping_tcp_segment_mismatch,omitempty"`
	DiscardRecordRoute                   *bool                          `json:"discard_record_route,omitempty"`
	DiscardSecurity                      *bool                          `json:"discard_security,omitempty"`
	DiscardStreamId                      *bool                          `json:"discard_stream_id,omitempty"`
	DiscardStrictSourceRouting           *bool                          `json:"discard_strict_source_routing,omitempty"`
	DiscardTcpSplitHandshake             *bool                          `json:"discard_tcp_split_handshake,omitempty"`
	DiscardTcpSynWithData                *bool                          `json:"discard_tcp_syn_with_data,omitempty"`
	DiscardTcpSynackWithData             *bool                          `json:"discard_tcp_synack_with_data,omitempty"`
	DiscardTimestamp                     *bool                          `json:"discard_timestamp,omitempty"`
	DiscardUnknownOption                 *bool                          `json:"discard_unknown_option,omitempty"`
	Flood                                *FloodObject                   `json:"flood,omitempty"`
	Folder                               *string                        `json:"folder,omitempty"`
	Id                                   *string                        `json:"id,omitempty"`
	Ipv6                                 *Ipv6Object                    `json:"ipv6,omitempty"`
	L2SecGroupTagProtection              *L2SecGroupTagProtectionObject `json:"l2_sec_group_tag_protection,omitempty"`
	Name                                 string                         `json:"name"`
	NonIpProtocol                        *NonIpProtocolObject           `json:"non_ip_protocol,omitempty"`
	RemoveTcpTimestamp                   *bool                          `json:"remove_tcp_timestamp,omitempty"`
	Scan                                 []ScanObject                   `json:"scan,omitempty"`
	ScanWhiteList                        []ScanWhiteListObject          `json:"scan_white_list,omitempty"`
	Snippet                              *string                        `json:"snippet,omitempty"`
	StrictIpCheck                        *bool                          `json:"strict_ip_check,omitempty"`
	StripMptcpOption                     *string                        `json:"strip_mptcp_option,omitempty"`
	StripTcpFastOpenAndData              *bool                          `json:"strip_tcp_fast_open_and_data,omitempty"`
	SuppressIcmpNeedfrag                 *bool                          `json:"suppress_icmp_needfrag,omitempty"`
	SuppressIcmpTimeexceeded             *bool                          `json:"suppress_icmp_timeexceeded,omitempty"`
	TcpRejectNonSyn                      *string                        `json:"tcp_reject_non_syn,omitempty"`
}

/*
FloodObject object.

ShortName:
Parent chains:
*
* flood

Args:

Param Icmp (IcmpObject): the Icmp param.

Param Icmpv6 (Icmpv6Object): the Icmpv6 param.

Param OtherIp (OtherIpObject): the OtherIp param.

Param SctpInit (SctpInitObject): the SctpInit param.

Param TcpSyn (TcpSynObject): the TcpSyn param.

Param Udp (UdpObject): the Udp param.
*/
type FloodObject struct {
	Icmp     *IcmpObject     `json:"icmp,omitempty"`
	Icmpv6   *Icmpv6Object   `json:"icmpv6,omitempty"`
	OtherIp  *OtherIpObject  `json:"other_ip,omitempty"`
	SctpInit *SctpInitObject `json:"sctp_init,omitempty"`
	TcpSyn   *TcpSynObject   `json:"tcp_syn,omitempty"`
	Udp      *UdpObject      `json:"udp,omitempty"`
}

/*
IcmpObject object.

ShortName:
Parent chains:
*
* flood
* icmp

Args:

Param Enable (bool): the Enable param.

Param Red (RedObject3): the Red param.
*/
type IcmpObject struct {
	Enable *bool       `json:"enable,omitempty"`
	Red    *RedObject3 `json:"red,omitempty"`
}

/*
RedObject3 object.

ShortName:
Parent chains:
*
* flood
* icmp
* red

Args:

Param ActivateRate (int64, required): the ActivateRate param.

Param AlarmRate (int64, required): the AlarmRate param.

Param MaximalRate (int64, required): the MaximalRate param.
*/
type RedObject3 struct {
	ActivateRate int64 `json:"activate_rate"`
	AlarmRate    int64 `json:"alarm_rate"`
	MaximalRate  int64 `json:"maximal_rate"`
}

/*
Icmpv6Object object.

ShortName:
Parent chains:
*
* flood
* icmpv6

Args:

Param Enable (bool): the Enable param.

Param Red (RedObject4): the Red param.
*/
type Icmpv6Object struct {
	Enable *bool       `json:"enable,omitempty"`
	Red    *RedObject4 `json:"red,omitempty"`
}

/*
RedObject4 object.

ShortName:
Parent chains:
*
* flood
* icmpv6
* red

Args:

Param ActivateRate (int64, required): the ActivateRate param.

Param AlarmRate (int64, required): the AlarmRate param.

Param MaximalRate (int64, required): the MaximalRate param.
*/
type RedObject4 struct {
	ActivateRate int64 `json:"activate_rate"`
	AlarmRate    int64 `json:"alarm_rate"`
	MaximalRate  int64 `json:"maximal_rate"`
}

/*
OtherIpObject object.

ShortName:
Parent chains:
*
* flood
* other_ip

Args:

Param Enable (bool): the Enable param.

Param Red (RedObject): the Red param.
*/
type OtherIpObject struct {
	Enable *bool      `json:"enable,omitempty"`
	Red    *RedObject `json:"red,omitempty"`
}

/*
RedObject object.

ShortName:
Parent chains:
*
* flood
* other_ip
* red

Args:

Param ActivateRate (int64, required): the ActivateRate param.

Param AlarmRate (int64, required): the AlarmRate param.

Param MaximalRate (int64, required): the MaximalRate param.
*/
type RedObject struct {
	ActivateRate int64 `json:"activate_rate"`
	AlarmRate    int64 `json:"alarm_rate"`
	MaximalRate  int64 `json:"maximal_rate"`
}

/*
SctpInitObject object.

ShortName:
Parent chains:
*
* flood
* sctp_init

Args:

Param Enable (bool): the Enable param.

Param Red (RedObject1): the Red param.
*/
type SctpInitObject struct {
	Enable *bool       `json:"enable,omitempty"`
	Red    *RedObject1 `json:"red,omitempty"`
}

/*
RedObject1 object.

ShortName:
Parent chains:
*
* flood
* sctp_init
* red

Args:

Param ActivateRate (int64, required): the ActivateRate param.

Param AlarmRate (int64, required): the AlarmRate param.

Param MaximalRate (int64, required): the MaximalRate param.
*/
type RedObject1 struct {
	ActivateRate int64 `json:"activate_rate"`
	AlarmRate    int64 `json:"alarm_rate"`
	MaximalRate  int64 `json:"maximal_rate"`
}

/*
TcpSynObject object.

ShortName:
Parent chains:
*
* flood
* tcp_syn

Args:

Param ActivateRate (int64): the ActivateRate param.

Param AlarmRate (int64): the AlarmRate param.

Param Enable (bool): the Enable param.

Param MaximalRate (int64): the MaximalRate param.

NOTE:  One of the following params should be specified:
  - AlarmRate
  - ActivateRate
  - MaximalRate
  - AlarmRate
  - ActivateRate
  - MaximalRate
*/
type TcpSynObject struct {
	ActivateRate *int64 `json:"activate_rate,omitempty"`
	AlarmRate    *int64 `json:"alarm_rate,omitempty"`
	Enable       *bool  `json:"enable,omitempty"`
	MaximalRate  *int64 `json:"maximal_rate,omitempty"`
}

/*
UdpObject object.

ShortName:
Parent chains:
*
* flood
* udp

Args:

Param Enable (bool): the Enable param.

Param Red (RedObject2): the Red param.
*/
type UdpObject struct {
	Enable *bool       `json:"enable,omitempty"`
	Red    *RedObject2 `json:"red,omitempty"`
}

/*
RedObject2 object.

ShortName:
Parent chains:
*
* flood
* udp
* red

Args:

Param ActivateRate (int64, required): the ActivateRate param.

Param AlarmRate (int64, required): the AlarmRate param.

Param MaximalRate (int64, required): the MaximalRate param.
*/
type RedObject2 struct {
	ActivateRate int64 `json:"activate_rate"`
	AlarmRate    int64 `json:"alarm_rate"`
	MaximalRate  int64 `json:"maximal_rate"`
}

/*
Ipv6Object object.

ShortName:
Parent chains:
*
* ipv6

Args:

Param AnycastSource (bool): the AnycastSource param.

Param FilterExtHdr (FilterExtHdrObject): the FilterExtHdr param.

Param Icmpv6TooBigSmallMtuDiscard (bool): the Icmpv6TooBigSmallMtuDiscard param.

Param IgnoreInvPkt (IgnoreInvPktObject): the IgnoreInvPkt param.

Param Ipv4CompatibleAddress (bool): the Ipv4CompatibleAddress param.

Param MulticastSource (bool): the MulticastSource param.

Param NeedlessFragmentHdr (bool): the NeedlessFragmentHdr param.

Param OptionsInvalidIpv6Discard (bool): the OptionsInvalidIpv6Discard param.

Param ReservedFieldSetDiscard (bool): the ReservedFieldSetDiscard param.

Param RoutingHeader0 (bool): the RoutingHeader0 param.

Param RoutingHeader1 (bool): the RoutingHeader1 param.

Param RoutingHeader253 (bool): the RoutingHeader253 param.

Param RoutingHeader254 (bool): the RoutingHeader254 param.

Param RoutingHeader255 (bool): the RoutingHeader255 param.

Param RoutingHeader3 (bool): the RoutingHeader3 param.

Param RoutingHeader4252 (bool): the RoutingHeader4252 param.
*/
type Ipv6Object struct {
	AnycastSource               *bool               `json:"anycast_source,omitempty"`
	FilterExtHdr                *FilterExtHdrObject `json:"filter_ext_hdr,omitempty"`
	Icmpv6TooBigSmallMtuDiscard *bool               `json:"icmpv6_too_big_small_mtu_discard,omitempty"`
	IgnoreInvPkt                *IgnoreInvPktObject `json:"ignore_inv_pkt,omitempty"`
	Ipv4CompatibleAddress       *bool               `json:"ipv4_compatible_address,omitempty"`
	MulticastSource             *bool               `json:"multicast_source,omitempty"`
	NeedlessFragmentHdr         *bool               `json:"needless_fragment_hdr,omitempty"`
	OptionsInvalidIpv6Discard   *bool               `json:"options_invalid_ipv6_discard,omitempty"`
	ReservedFieldSetDiscard     *bool               `json:"reserved_field_set_discard,omitempty"`
	RoutingHeader0              *bool               `json:"routing_header_0,omitempty"`
	RoutingHeader1              *bool               `json:"routing_header_1,omitempty"`
	RoutingHeader253            *bool               `json:"routing_header_253,omitempty"`
	RoutingHeader254            *bool               `json:"routing_header_254,omitempty"`
	RoutingHeader255            *bool               `json:"routing_header_255,omitempty"`
	RoutingHeader3              *bool               `json:"routing_header_3,omitempty"`
	RoutingHeader4252           *bool               `json:"routing_header_4_252,omitempty"`
}

/*
FilterExtHdrObject object.

ShortName:
Parent chains:
*
* ipv6
* filter_ext_hdr

Args:

Param DestOptionHdr (bool): the DestOptionHdr param.

Param HopByHopHdr (bool): the HopByHopHdr param.

Param RoutingHdr (bool): the RoutingHdr param.
*/
type FilterExtHdrObject struct {
	DestOptionHdr *bool `json:"dest_option_hdr,omitempty"`
	HopByHopHdr   *bool `json:"hop_by_hop_hdr,omitempty"`
	RoutingHdr    *bool `json:"routing_hdr,omitempty"`
}

/*
IgnoreInvPktObject object.

ShortName:
Parent chains:
*
* ipv6
* ignore_inv_pkt

Args:

Param DestUnreach (bool): the DestUnreach param.

Param ParamProblem (bool): the ParamProblem param.

Param PktTooBig (bool): the PktTooBig param.

Param Redirect (bool): the Redirect param.

Param TimeExceeded (bool): the TimeExceeded param.
*/
type IgnoreInvPktObject struct {
	DestUnreach  *bool `json:"dest_unreach,omitempty"`
	ParamProblem *bool `json:"param_problem,omitempty"`
	PktTooBig    *bool `json:"pkt_too_big,omitempty"`
	Redirect     *bool `json:"redirect,omitempty"`
	TimeExceeded *bool `json:"time_exceeded,omitempty"`
}

/*
L2SecGroupTagProtectionObject object.

ShortName:
Parent chains:
*
* l2_sec_group_tag_protection

Args:

Param Tags ([]TagsObject): the Tags param.
*/
type L2SecGroupTagProtectionObject struct {
	Tags []TagsObject `json:"tags,omitempty"`
}

/*
TagsObject object.

ShortName:
Parent chains:
*
* l2_sec_group_tag_protection
* tags
* _inline

Args:

Param Enable (bool): the Enable param.

Param Name (string, required): the Name param.

Param Tag (string, required): the Tag param.
*/
type TagsObject struct {
	Enable *bool  `json:"enable,omitempty"`
	Name   string `json:"name"`
	Tag    string `json:"tag"`
}

/*
NonIpProtocolObject object.

ShortName:
Parent chains:
*
* non_ip_protocol

Args:

Param ListType (string): the ListType param. String must be one of these: `"exclude"`, `"include"`.

Param Protocol ([]ProtocolObject): the Protocol param.
*/
type NonIpProtocolObject struct {
	ListType *string          `json:"list_type,omitempty"`
	Protocol []ProtocolObject `json:"protocol,omitempty"`
}

/*
ProtocolObject object.

ShortName:
Parent chains:
*
* non_ip_protocol
* protocol
* _inline

Args:

Param Enable (bool): the Enable param.

Param EtherType (string, required): the EtherType param.

Param Name (string, required): the Name param.
*/
type ProtocolObject struct {
	Enable    *bool  `json:"enable,omitempty"`
	EtherType string `json:"ether_type"`
	Name      string `json:"name"`
}

/*
ScanObject object.

ShortName:
Parent chains:
*
* scan
* _inline

Args:

Param Action (ActionObject): the Action param.

Param Interval (int64): the Interval param.

Param Name (string, required): the Name param.

Param Threshold (int64): the Threshold param.
*/
type ScanObject struct {
	Action    *ActionObject `json:"action,omitempty"`
	Interval  *int64        `json:"interval,omitempty"`
	Name      string        `json:"name"`
	Threshold *int64        `json:"threshold,omitempty"`
}

/*
ActionObject object.

ShortName:
Parent chains:
*
* scan
* _inline
* action

Args:

Param Duration (int64): the Duration param.

Param TrackBy (string): the TrackBy param. String must be one of these: `"source-and-destination"`, `"source"`.

NOTE:  One of the following params should be specified:
  - TrackBy
  - Duration
*/
type ActionObject struct {
	Duration *int64  `json:"duration,omitempty"`
	TrackBy  *string `json:"track_by,omitempty"`
}

/*
ScanWhiteListObject object.

ShortName:
Parent chains:
*
* scan_white_list
* _inline

Args:

Param Name (string, required): the Name param.
*/
type ScanWhiteListObject struct {
	Name string `json:"name"`
}
