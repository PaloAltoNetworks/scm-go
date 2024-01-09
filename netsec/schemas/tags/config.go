package tags

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/tags

/*
Config object.

ShortName: qFWAgJG
Parent chains:
*

Args:

Param Color (string): the Color param. String must be one of these: `"Red"`, `"Green"`, `"Blue"`, `"Yellow"`, `"Copper"`, `"Orange"`, `"Purple"`, `"Gray"`, `"Light Green"`, `"Cyan"`, `"Light Gray"`, `"Blue Gray"`, `"Lime"`, `"Black"`, `"Gold"`, `"Brown"`, `"Olive"`, `"Maroon"`, `"Red-Orange"`, `"Yellow-Orange"`, `"Forest Green"`, `"Turquoise Blue"`, `"Azure Blue"`, `"Cerulean Blue"`, `"Midnight Blue"`, `"Medium Blue"`, `"Cobalt Blue"`, `"Violet Blue"`, `"Blue Violet"`, `"Medium Violet"`, `"Medium Rose"`, `"Lavender"`, `"Orchid"`, `"Thistle"`, `"Peach"`, `"Salmon"`, `"Magenta"`, `"Red Violet"`, `"Mahogany"`, `"Burnt Sienna"`, `"Chestnut"`.

Param Comments (string): the Comments param. String length must not exceed 1023 characters.

Param Id (string, read-only): UUID of the resource

Param Name (string, required): the Name param. String length must not exceed 127 characters.
*/
type Config struct {
	Color    *string `json:"color,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Id       *string `json:"id,omitempty"`
	Name     string  `json:"name"`
}
