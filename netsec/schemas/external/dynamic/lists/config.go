package lists

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/external-dynamic-lists

/*
Config object.

ShortName: hhIWLbI
Parent chains:
*

Args:

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 63 characters.

Param Type (TypeObject, required): the Type param.
*/
type Config struct {
	Id   *string    `json:"id,omitempty"`
	Name string     `json:"name"`
	Type TypeObject `json:"type"`
}

/*
TypeObject object.

ShortName:
Parent chains:
*
* type

Args:

Param Domain (DomainObject): the Domain param.

Param Imei (ImeiObject): the Imei param.

Param Imsi (ImsiObject): the Imsi param.

Param Ip (IpObject): the Ip param.

Param PredefinedIp (PredefinedIpObject): the PredefinedIp param.

Param PredefinedUrl (PredefinedUrlObject): the PredefinedUrl param.

Param Url (UrlObject): the Url param.

NOTE:  One of the following params should be specified:
  - PredefinedIp
  - PredefinedUrl
  - Ip
  - Domain
  - Url
  - Imsi
  - Imei
*/
type TypeObject struct {
	Domain        *DomainObject        `json:"domain,omitempty"`
	Imei          *ImeiObject          `json:"imei,omitempty"`
	Imsi          *ImsiObject          `json:"imsi,omitempty"`
	Ip            *IpObject            `json:"ip,omitempty"`
	PredefinedIp  *PredefinedIpObject  `json:"predefined_ip,omitempty"`
	PredefinedUrl *PredefinedUrlObject `json:"predefined_url,omitempty"`
	Url           *UrlObject           `json:"url,omitempty"`
}

/*
DomainObject object.

ShortName:
Parent chains:
*
* type
* domain

Args:

Param CertificateProfile (string): Profile for authenticating client certificates Default: `"None"`.

Param Description (string): the Description param. String length must not exceed 255 characters.

Param DomainAuth (DomainAuthObject): the DomainAuth param.

Param ExceptionList ([]string): the ExceptionList param.

Param ExpandDomain (bool): Enable/Disable expand domain Default: `false`.

Param Recurring (DomainRecurringObject, required): the Recurring param.

Param Url (string, required): the Url param. String length must not exceed 255 characters. Default: `"http://"`.
*/
type DomainObject struct {
	CertificateProfile *string               `json:"certificate_profile,omitempty"`
	Description        *string               `json:"description,omitempty"`
	DomainAuth         *DomainAuthObject     `json:"auth,omitempty"`
	ExceptionList      []string              `json:"exception_list,omitempty"`
	ExpandDomain       *bool                 `json:"expand_domain,omitempty"`
	Recurring          DomainRecurringObject `json:"recurring"`
	Url                string                `json:"url"`
}

/*
DomainAuthObject object.

ShortName:
Parent chains:
*
* type
* domain
* auth

Args:

Param Password (string, required): the Password param. String length must not exceed 255 characters.

Param Username (string, required): the Username param. String length must be between 1 and 255 characters.
*/
type DomainAuthObject struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

/*
DomainRecurringObject object.

ShortName:
Parent chains:
*
* type
* domain
* recurring

Args:

Param Daily (DomainDailyObject): the Daily param.

Param FiveMinute (any): the FiveMinute param.

Param Hourly (any): the Hourly param.

Param Monthly (DomainMonthyObject): the Monthly param.

Param Weekly (DomainWeeklyObject): the Weekly param.

NOTE:  One of the following params should be specified:
  - Hourly
  - FiveMinute
  - Daily
  - Weekly
  - Monthly
*/
type DomainRecurringObject struct {
	Daily      *DomainDailyObject  `json:"daily,omitempty"`
	FiveMinute any                 `json:"five_minute,omitempty"`
	Hourly     any                 `json:"hourly,omitempty"`
	Monthly    *DomainMonthyObject `json:"monthly,omitempty"`
	Weekly     *DomainWeeklyObject `json:"weekly,omitempty"`
}

/*
DomainDailyObject object.

ShortName:
Parent chains:
*
* type
* domain
* recurring
* daily

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
*/
type DomainDailyObject struct {
	At string `json:"at"`
}

/*
DomainMonthyObject object.

ShortName:
Parent chains:
*
* type
* domain
* recurring
* monthly

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.

Param DayOfMonth (int64, required): the DayOfMonth param. Value must be between 1 and 31.
*/
type DomainMonthyObject struct {
	At         string `json:"at"`
	DayOfMonth int64  `json:"day_of_month"`
}

/*
DomainWeeklyObject object.

ShortName:
Parent chains:
*
* type
* domain
* recurring
* weekly

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.

Param DayOfWeek (string, required): the DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.
*/
type DomainWeeklyObject struct {
	At        string `json:"at"`
	DayOfWeek string `json:"day_of_week"`
}

/*
ImeiObject object.

ShortName:
Parent chains:
*
* type
* imei

Args:

Param CertificateProfile (string): Profile for authenticating client certificates Default: `"None"`.

Param Description (string): the Description param. String length must not exceed 255 characters.

Param ExceptionList ([]string): the ExceptionList param.

Param ImeiAuth (ImeiAuthObject): the ImeiAuth param.

Param Recurring (ImeiRecurringObject, required): the Recurring param.

Param Url (string, required): the Url param. String length must not exceed 255 characters. Default: `"http://"`.
*/
type ImeiObject struct {
	CertificateProfile *string             `json:"certificate_profile,omitempty"`
	Description        *string             `json:"description,omitempty"`
	ExceptionList      []string            `json:"exception_list,omitempty"`
	ImeiAuth           *ImeiAuthObject     `json:"auth,omitempty"`
	Recurring          ImeiRecurringObject `json:"recurring"`
	Url                string              `json:"url"`
}

/*
ImeiAuthObject object.

ShortName:
Parent chains:
*
* type
* imei
* auth

Args:

Param Password (string, required): the Password param. String length must not exceed 255 characters.

Param Username (string, required): the Username param. String length must be between 1 and 255 characters.
*/
type ImeiAuthObject struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

/*
ImeiRecurringObject object.

ShortName:
Parent chains:
*
* type
* imei
* recurring

Args:

Param Daily (ImeiDailyObject): the Daily param.

Param FiveMinute (any): the FiveMinute param.

Param Hourly (any): the Hourly param.

Param Monthly (ImeiMonthyObject): the Monthly param.

Param Weekly (ImeiWeeklyObject): the Weekly param.

NOTE:  One of the following params should be specified:
  - FiveMinute
  - Hourly
  - Daily
  - Weekly
  - Monthly
*/
type ImeiRecurringObject struct {
	Daily      *ImeiDailyObject  `json:"daily,omitempty"`
	FiveMinute any               `json:"five_minute,omitempty"`
	Hourly     any               `json:"hourly,omitempty"`
	Monthly    *ImeiMonthyObject `json:"monthly,omitempty"`
	Weekly     *ImeiWeeklyObject `json:"weekly,omitempty"`
}

/*
ImeiDailyObject object.

ShortName:
Parent chains:
*
* type
* imei
* recurring
* daily

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
*/
type ImeiDailyObject struct {
	At string `json:"at"`
}

/*
ImeiMonthyObject object.

ShortName:
Parent chains:
*
* type
* imei
* recurring
* monthly

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.

Param DayOfMonth (int64, required): the DayOfMonth param. Value must be between 1 and 31.
*/
type ImeiMonthyObject struct {
	At         string `json:"at"`
	DayOfMonth int64  `json:"day_of_month"`
}

/*
ImeiWeeklyObject object.

ShortName:
Parent chains:
*
* type
* imei
* recurring
* weekly

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.

Param DayOfWeek (string, required): the DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.
*/
type ImeiWeeklyObject struct {
	At        string `json:"at"`
	DayOfWeek string `json:"day_of_week"`
}

/*
ImsiObject object.

ShortName:
Parent chains:
*
* type
* imsi

Args:

Param CertificateProfile (string): Profile for authenticating client certificates Default: `"None"`.

Param Description (string): the Description param. String length must not exceed 255 characters.

Param ExceptionList ([]string): the ExceptionList param.

Param ImsiAuth (ImsiAuthObject): the ImsiAuth param.

Param Recurring (ImsiRecurringObject, required): the Recurring param.

Param Url (string, required): the Url param. String length must not exceed 255 characters. Default: `"http://"`.
*/
type ImsiObject struct {
	CertificateProfile *string             `json:"certificate_profile,omitempty"`
	Description        *string             `json:"description,omitempty"`
	ExceptionList      []string            `json:"exception_list,omitempty"`
	ImsiAuth           *ImsiAuthObject     `json:"auth,omitempty"`
	Recurring          ImsiRecurringObject `json:"recurring"`
	Url                string              `json:"url"`
}

/*
ImsiAuthObject object.

ShortName:
Parent chains:
*
* type
* imsi
* auth

Args:

Param Password (string, required): the Password param. String length must not exceed 255 characters.

Param Username (string, required): the Username param. String length must be between 1 and 255 characters.
*/
type ImsiAuthObject struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

/*
ImsiRecurringObject object.

ShortName:
Parent chains:
*
* type
* imsi
* recurring

Args:

Param Daily (ImsiDailyObject): the Daily param.

Param FiveMinute (any): the FiveMinute param.

Param Hourly (any): the Hourly param.

Param Monthly (ImsiMonthyObject): the Monthly param.

Param Weekly (ImsiWeeklyObject): the Weekly param.

NOTE:  One of the following params should be specified:
  - FiveMinute
  - Hourly
  - Daily
  - Weekly
  - Monthly
*/
type ImsiRecurringObject struct {
	Daily      *ImsiDailyObject  `json:"daily,omitempty"`
	FiveMinute any               `json:"five_minute,omitempty"`
	Hourly     any               `json:"hourly,omitempty"`
	Monthly    *ImsiMonthyObject `json:"monthly,omitempty"`
	Weekly     *ImsiWeeklyObject `json:"weekly,omitempty"`
}

/*
ImsiDailyObject object.

ShortName:
Parent chains:
*
* type
* imsi
* recurring
* daily

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
*/
type ImsiDailyObject struct {
	At string `json:"at"`
}

/*
ImsiMonthyObject object.

ShortName:
Parent chains:
*
* type
* imsi
* recurring
* monthly

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.

Param DayOfMonth (int64, required): the DayOfMonth param. Value must be between 1 and 31.
*/
type ImsiMonthyObject struct {
	At         string `json:"at"`
	DayOfMonth int64  `json:"day_of_month"`
}

/*
ImsiWeeklyObject object.

ShortName:
Parent chains:
*
* type
* imsi
* recurring
* weekly

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.

Param DayOfWeek (string, required): the DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.
*/
type ImsiWeeklyObject struct {
	At        string `json:"at"`
	DayOfWeek string `json:"day_of_week"`
}

/*
IpObject object.

ShortName:
Parent chains:
*
* type
* ip

Args:

Param CertificateProfile (string): Profile for authenticating client certificates Default: `"None"`.

Param Description (string): the Description param. String length must not exceed 255 characters.

Param ExceptionList ([]string): the ExceptionList param.

Param IpAuth (IpAuthObject): the IpAuth param.

Param Recurring (IpRecurringObject, required): the Recurring param.

Param Url (string, required): the Url param. String length must not exceed 255 characters. Default: `"http://"`.
*/
type IpObject struct {
	CertificateProfile *string           `json:"certificate_profile,omitempty"`
	Description        *string           `json:"description,omitempty"`
	ExceptionList      []string          `json:"exception_list,omitempty"`
	IpAuth             *IpAuthObject     `json:"auth,omitempty"`
	Recurring          IpRecurringObject `json:"recurring"`
	Url                string            `json:"url"`
}

/*
IpAuthObject object.

ShortName:
Parent chains:
*
* type
* ip
* auth

Args:

Param Password (string, required): the Password param. String length must not exceed 255 characters.

Param Username (string, required): the Username param. String length must be between 1 and 255 characters.
*/
type IpAuthObject struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

/*
IpRecurringObject object.

ShortName:
Parent chains:
*
* type
* ip
* recurring

Args:

Param Daily (IpDailyObject): the Daily param.

Param FiveMinute (any): the FiveMinute param.

Param Hourly (any): the Hourly param.

Param Monthly (IpMonthyObject): the Monthly param.

Param Weekly (IpWeeklyObject): the Weekly param.

NOTE:  One of the following params should be specified:
  - FiveMinute
  - Hourly
  - Daily
  - Weekly
  - Monthly
*/
type IpRecurringObject struct {
	Daily      *IpDailyObject  `json:"daily,omitempty"`
	FiveMinute any             `json:"five_minute,omitempty"`
	Hourly     any             `json:"hourly,omitempty"`
	Monthly    *IpMonthyObject `json:"monthly,omitempty"`
	Weekly     *IpWeeklyObject `json:"weekly,omitempty"`
}

/*
IpDailyObject object.

ShortName:
Parent chains:
*
* type
* ip
* recurring
* daily

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
*/
type IpDailyObject struct {
	At string `json:"at"`
}

/*
IpMonthyObject object.

ShortName:
Parent chains:
*
* type
* ip
* recurring
* monthly

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.

Param DayOfMonth (int64, required): the DayOfMonth param. Value must be between 1 and 31.
*/
type IpMonthyObject struct {
	At         string `json:"at"`
	DayOfMonth int64  `json:"day_of_month"`
}

/*
IpWeeklyObject object.

ShortName:
Parent chains:
*
* type
* ip
* recurring
* weekly

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.

Param DayOfWeek (string, required): the DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.
*/
type IpWeeklyObject struct {
	At        string `json:"at"`
	DayOfWeek string `json:"day_of_week"`
}

/*
PredefinedIpObject object.

ShortName:
Parent chains:
*
* type
* predefined_ip

Args:

Param Description (string): the Description param. String length must not exceed 255 characters.

Param ExceptionList ([]string): the ExceptionList param.

Param Url (string, required): the Url param.
*/
type PredefinedIpObject struct {
	Description   *string  `json:"description,omitempty"`
	ExceptionList []string `json:"exception_list,omitempty"`
	Url           string   `json:"url"`
}

/*
PredefinedUrlObject object.

ShortName:
Parent chains:
*
* type
* predefined_url

Args:

Param Description (string): the Description param. String length must not exceed 255 characters.

Param ExceptionList ([]string): the ExceptionList param.

Param Url (string, required): the Url param.
*/
type PredefinedUrlObject struct {
	Description   *string  `json:"description,omitempty"`
	ExceptionList []string `json:"exception_list,omitempty"`
	Url           string   `json:"url"`
}

/*
UrlObject object.

ShortName:
Parent chains:
*
* type
* url

Args:

Param CertificateProfile (string): Profile for authenticating client certificates Default: `"None"`.

Param Description (string): the Description param. String length must not exceed 255 characters.

Param ExceptionList ([]string): the ExceptionList param.

Param Recurring (UrlRecurringObject, required): the Recurring param.

Param Url (string, required): the Url param. String length must not exceed 255 characters. Default: `"http://"`.

Param UrlAuth (UrlAuthObject): the UrlAuth param.
*/
type UrlObject struct {
	CertificateProfile *string            `json:"certificate_profile,omitempty"`
	Description        *string            `json:"description,omitempty"`
	ExceptionList      []string           `json:"exception_list,omitempty"`
	Recurring          UrlRecurringObject `json:"recurring"`
	Url                string             `json:"url"`
	UrlAuth            *UrlAuthObject     `json:"auth,omitempty"`
}

/*
UrlRecurringObject object.

ShortName:
Parent chains:
*
* type
* url
* recurring

Args:

Param Daily (UrlDailyObject): the Daily param.

Param FiveMinute (any): the FiveMinute param.

Param Hourly (any): the Hourly param.

Param Monthly (UrlMonthyObject): the Monthly param.

Param Weekly (UrlWeeklyObject): the Weekly param.

NOTE:  One of the following params should be specified:
  - Hourly
  - FiveMinute
  - Daily
  - Weekly
  - Monthly
*/
type UrlRecurringObject struct {
	Daily      *UrlDailyObject  `json:"daily,omitempty"`
	FiveMinute any              `json:"five_minute,omitempty"`
	Hourly     any              `json:"hourly,omitempty"`
	Monthly    *UrlMonthyObject `json:"monthly,omitempty"`
	Weekly     *UrlWeeklyObject `json:"weekly,omitempty"`
}

/*
UrlDailyObject object.

ShortName:
Parent chains:
*
* type
* url
* recurring
* daily

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.
*/
type UrlDailyObject struct {
	At string `json:"at"`
}

/*
UrlMonthyObject object.

ShortName:
Parent chains:
*
* type
* url
* recurring
* monthly

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.

Param DayOfMonth (int64, required): the DayOfMonth param. Value must be between 1 and 31.
*/
type UrlMonthyObject struct {
	At         string `json:"at"`
	DayOfMonth int64  `json:"day_of_month"`
}

/*
UrlWeeklyObject object.

ShortName:
Parent chains:
*
* type
* url
* recurring
* weekly

Args:

Param At (string, required): Time specification hh (e.g. 20) String length must be between 2 and 2 characters. String validation regex: `([01][0-9]|[2][0-3])`. Default: `"00"`.

Param DayOfWeek (string, required): the DayOfWeek param. String must be one of these: `"sunday"`, `"monday"`, `"tuesday"`, `"wednesday"`, `"thursday"`, `"friday"`, `"saturday"`.
*/
type UrlWeeklyObject struct {
	At        string `json:"at"`
	DayOfWeek string `json:"day_of_week"`
}

/*
UrlAuthObject object.

ShortName:
Parent chains:
*
* type
* url
* auth

Args:

Param Password (string, required): the Password param. String length must not exceed 255 characters.

Param Username (string, required): the Username param. String length must be between 1 and 255 characters.
*/
type UrlAuthObject struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
