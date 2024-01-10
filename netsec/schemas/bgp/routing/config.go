package routing

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/bgp-routing

/*
Config object.

ShortName: suxdMuj
Parent chains:
*

Args:

Param AcceptRouteOverSC (bool): the AcceptRouteOverSC param.

Param AddHostRouteToIkePeer (bool): the AddHostRouteToIkePeer param.

Param BackboneRouting (string): the BackboneRouting param. String must be one of these: `"no-asymmetric-routing"`, `"asymmetric-routing-only"`, `"asymmetric-routing-with-load-share"`.

Param OutboundRoutesForServices ([]string): the OutboundRoutesForServices param.

Param RoutingPreference (RoutingPreferenceObject): the RoutingPreference param.

Param WithdrawStaticRoute (bool): the WithdrawStaticRoute param.
*/
type Config struct {
	AcceptRouteOverSC         *bool                    `json:"accept_route_over_SC,omitempty"`
	AddHostRouteToIkePeer     *bool                    `json:"add_host_route_to_ike_peer,omitempty"`
	BackboneRouting           *string                  `json:"backbone_routing,omitempty"`
	OutboundRoutesForServices []string                 `json:"outbound_routes_for_services,omitempty"`
	RoutingPreference         *RoutingPreferenceObject `json:"routing_preference,omitempty"`
	WithdrawStaticRoute       *bool                    `json:"withdraw_static_route,omitempty"`
}

/*
RoutingPreferenceObject object.

ShortName:
Parent chains:
*
* routing_preference

Args:

Param Default (any): the Default param. Default: `false`.

Param HotPotatoRouting (any): the HotPotatoRouting param. Default: `false`.

NOTE:  One of the following params should be specified:
  - Default
  - HotPotatoRouting
*/
type RoutingPreferenceObject struct {
	Default          any `json:"default,omitempty"`
	HotPotatoRouting any `json:"hot_potato_routing,omitempty"`
}
