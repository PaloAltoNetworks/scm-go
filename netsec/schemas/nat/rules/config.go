package rules

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/nat-rules

/*
Config object.

ShortName: aeWshcf
Parent chains:
*

Args:

Param ActiveActiveDeviceBinding (string): the ActiveActiveDeviceBinding param. String must be one of these: `"primary"`, `"both"`, `"0"`, `"1"`.

Param Description (string): the Description param.

Param DestinationTranslation (DestinationTranslationObject): Static destination translation parameter.

Param Destinations ([]string, required): The destination address(es)

Param Device (string): The device in which the resource is defined String length must not exceed 64 characters. String validation regex: `^[a-zA-Z\d-_\. ]+$`.

Param Disabled (bool): the Disabled param.

Param DynamicDestinationTranslation (DynamicDestinationTranslationObject): Dynamic destination translation parameter.

Param Folder (string): The folder in which the resource is defined String length must not exceed 64 characters. String validation regex: `^[a-zA-Z\d-_\. ]+$`.

Param Froms ([]string, required): The source security zone(s)

Param GroupTag (string): the GroupTag param.

Param Id (string, read-only): the Id param.

Param Name (string, required): the Name param.

Param NatType (string): the NatType param. String must be one of these: `"ipv4"`, `"nat64"`, `"nptv6"`.

Param Service (string, required): the Service param.

Param Snippet (string): The snippet in which the resource is defined String length must not exceed 64 characters. String validation regex: `^[a-zA-Z\d-_\. ]+$`.

Param SourceTranslation (SourceTranslationObject): the SourceTranslation param.

Param Sources ([]string, required): The source address(es)

Param Tags ([]string): the Tags param.

Param Target (TargetObject): the Target param.

Param ToInterface (string): the ToInterface param.

Param Tos ([]string, required): The destination security zone(s)
*/
type Config struct {
	ActiveActiveDeviceBinding     *string                              `json:"active_active_device_binding,omitempty"`
	Description                   *string                              `json:"description,omitempty"`
	DestinationTranslation        *DestinationTranslationObject        `json:"destination_translation,omitempty"`
	Destinations                  []string                             `json:"destination"`
	Device                        *string                              `json:"device,omitempty"`
	Disabled                      *bool                                `json:"disabled,omitempty"`
	DynamicDestinationTranslation *DynamicDestinationTranslationObject `json:"dynamic_destination_translation,omitempty"`
	Folder                        *string                              `json:"folder,omitempty"`
	Froms                         []string                             `json:"from"`
	GroupTag                      *string                              `json:"group_tag,omitempty"`
	Id                            *string                              `json:"id,omitempty"`
	Name                          string                               `json:"name"`
	NatType                       *string                              `json:"nat_type,omitempty"`
	Service                       string                               `json:"service"`
	Snippet                       *string                              `json:"snippet,omitempty"`
	SourceTranslation             *SourceTranslationObject             `json:"source_translation,omitempty"`
	Sources                       []string                             `json:"source"`
	Tags                          []string                             `json:"tag,omitempty"`
	Target                        *TargetObject                        `json:"target,omitempty"`
	ToInterface                   *string                              `json:"to_interface,omitempty"`
	Tos                           []string                             `json:"to"`
}

/*
DestinationTranslationObject Static destination translation parameter.

ShortName:
Parent chains:
*
* destination_translation

Args:

Param DnsRewrite (DnsRewriteObject): the DnsRewrite param.

Param TranslatedAddressSingle (string, required): The ip address to be translated. String validation regex: `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(?:[A-Fa-f0-9]{1,4}:){7}[A-Fa-f0-9]{1,4}$`.

Param TranslatedPort (int64): the TranslatedPort param.
*/
type DestinationTranslationObject struct {
	DnsRewrite              *DnsRewriteObject `json:"dns_rewrite,omitempty"`
	TranslatedAddressSingle string            `json:"translated_address_single"`
	TranslatedPort          *int64            `json:"translated_port,omitempty"`
}

/*
DnsRewriteObject object.

ShortName:
Parent chains:
*
* destination_translation
* dns_rewrite

Args:

Param Direction (string, required): the Direction param. String must be one of these: `"reverse"`, `"forward"`. Default: `"reverse"`.
*/
type DnsRewriteObject struct {
	Direction string `json:"direction"`
}

/*
DynamicDestinationTranslationObject Dynamic destination translation parameter.

ShortName:
Parent chains:
*
* dynamic_destination_translation

Args:

Param Distribution (string, required): the Distribution param. String must be one of these: `"round-robin"`, `"source-ip-hash"`, `"ip-modulo"`, `"ip-hash"`, `"least-sessions"`. Default: `"round-robin"`.

Param TranslatedAddressSingle (string, required): The ip address to be translated. String validation regex: `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(?:[A-Fa-f0-9]{1,4}:){7}[A-Fa-f0-9]{1,4}$`.

Param TranslatedPort (int64): the TranslatedPort param.
*/
type DynamicDestinationTranslationObject struct {
	Distribution            string `json:"distribution"`
	TranslatedAddressSingle string `json:"translated_address_single"`
	TranslatedPort          *int64 `json:"translated_port,omitempty"`
}

/*
SourceTranslationObject object.

ShortName:
Parent chains:
*
* source_translation

Args:

Param BiDirectional (string): the BiDirectional param. String must be one of these: `"yes"`, `"no"`.

Param Fallback (FallbackObject): the Fallback param.

Param TranslatedAddressArray ([]string): the TranslatedAddressArray param.

Param TranslatedAddressSingle (string): the TranslatedAddressSingle param. String validation regex: `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(?:[A-Fa-f0-9]{1,4}:){7}[A-Fa-f0-9]{1,4}$`.

NOTE:  One of the following params should be specified:
  - TranslatedAddressArray
  - TranslatedAddressSingle
*/
type SourceTranslationObject struct {
	BiDirectional           *string         `json:"bi_directional,omitempty"`
	Fallback                *FallbackObject `json:"fallback,omitempty"`
	TranslatedAddressArray  []string        `json:"translated_address_array,omitempty"`
	TranslatedAddressSingle *string         `json:"translated_address_single,omitempty"`
}

/*
FallbackObject object.

ShortName:
Parent chains:
*
* source_translation
* fallback

Args:

Param Interface (string): the Interface param.

NOTE:  One of the following params should be specified:
  - Interface
*/
type FallbackObject struct {
	Interface *string `json:"interface,omitempty"`
}

/*
TargetObject object.

ShortName:
Parent chains:
*
* target

Args:

Param Devices ([]DevicesObject): the Devices param.

Param Negate (bool): the Negate param.

Param Tags ([]string): the Tags param.
*/
type TargetObject struct {
	Devices []DevicesObject `json:"devices,omitempty"`
	Negate  *bool           `json:"negate,omitempty"`
	Tags    []string        `json:"tags,omitempty"`
}

/*
DevicesObject object.

ShortName:
Parent chains:
*
* target
* devices
* _inline

Args:

Param Name (string): the Name param.
*/
type DevicesObject struct {
	Name *string `json:"name,omitempty"`
}
