package move

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/rule-based-move

/*
Config object.

ShortName: uewNibC
Parent chains:
*

Args:

Param Destination (string, required): A destination of the rule. Valid destination values are top, bottom, before and after. String must be one of these: `"top"`, `"bottom"`, `"before"`, `"after"`.

Param DestinationRule (string): A destination_rule attribute is required only if the destination value is before or after. Valid destination_rule values are existing rule UUIDs within the same container.

Param Rulebase (string, required): A base of a rule. Valid rulebase values are pre and post. String must be one of these: `"pre"`, `"post"`.
*/
type Config struct {
	Destination     string  `json:"destination"`
	DestinationRule *string `json:"destination_rule,omitempty"`
	Rulebase        string  `json:"rulebase"`
}
