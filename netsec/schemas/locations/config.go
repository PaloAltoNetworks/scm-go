package locations

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/locations

/*
Config object.

ShortName: aeWshcf
Parent chains:
*

Args:

Param AggregateRegion (string): the AggregateRegion param.

Param Continent (string): the Continent param.

Param Display (string): the Display param.

Param Latitude (float64): the Latitude param. Value must be between -90 and 90.

Param Longitude (float64): the Longitude param. Value must be between -180 and 180.

Param Region (string): the Region param.

Param Value (string): the Value param.
*/
type Config struct {
	AggregateRegion *string  `json:"aggregate_region,omitempty"`
	Continent       *string  `json:"continent,omitempty"`
	Display         *string  `json:"display,omitempty"`
	Latitude        *float64 `json:"latitude,omitempty"`
	Longitude       *float64 `json:"longitude,omitempty"`
	Region          *string  `json:"region,omitempty"`
	Value           *string  `json:"value,omitempty"`
}
