package regions

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/regions

/*
Config object.

ShortName: mhJDwSQ
Parent chains:
*

Args:

Param Addresses ([]string): the Addresses param.

Param GeoLocation (GeoLocationObject): the GeoLocation param.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.
*/
type Config struct {
	Addresses   []string           `json:"address,omitempty"`
	GeoLocation *GeoLocationObject `json:"geo_location,omitempty"`
	Id          *string            `json:"id,omitempty"`
	Name        string             `json:"name"`
}

/*
GeoLocationObject object.

ShortName:
Parent chains:
*
* geo_location

Args:

Param Latitude (float64, required): latitude coordinate Value must be between -90 and 90.

Param Longitude (float64, required): longitude coordinate Value must be between -180 and 180.
*/
type GeoLocationObject struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
