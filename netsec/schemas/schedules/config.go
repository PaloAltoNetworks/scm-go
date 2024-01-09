package schedules

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/schedules

/*
Config object.

ShortName: mCcgFKg
Parent chains:
*

Args:

Param Id (string, read-only): UUID of the resource

Param Name (string, required): Alphanumeric string [ 0-9a-zA-Z._-] String length must not exceed 31 characters.

Param ScheduleType (ScheduleTypeObject, required): the ScheduleType param.
*/
type Config struct {
	Id           *string            `json:"id,omitempty"`
	Name         string             `json:"name"`
	ScheduleType ScheduleTypeObject `json:"schedule_type"`
}

/*
ScheduleTypeObject object.

ShortName:
Parent chains:
*
* schedule_type

Args:

Param NonRecurringList ([]string): the NonRecurringList param.

Param Recurring (RecurringObject): the Recurring param.

NOTE:  One of the following params should be specified:
  - Recurring
  - NonRecurringList
*/
type ScheduleTypeObject struct {
	NonRecurringList []string         `json:"non_recurring,omitempty"`
	Recurring        *RecurringObject `json:"recurring,omitempty"`
}

/*
RecurringObject object.

ShortName:
Parent chains:
*
* schedule_type
* recurring

Args:

Param DailyList ([]string): the DailyList param.

Param Weekly (WeeklyObject): the Weekly param.

NOTE:  One of the following params should be specified:
  - Weekly
  - DailyList
*/
type RecurringObject struct {
	DailyList []string      `json:"daily,omitempty"`
	Weekly    *WeeklyObject `json:"weekly,omitempty"`
}

/*
WeeklyObject object.

ShortName:
Parent chains:
*
* schedule_type
* recurring
* weekly

Args:

Param FridayList ([]string): the FridayList param.

Param MondayList ([]string): the MondayList param.

Param SaturdayList ([]string): the SaturdayList param.

Param SundayList ([]string): the SundayList param.

Param ThursdayList ([]string): the ThursdayList param.

Param TuesdayList ([]string): the TuesdayList param.

Param WednesdayList ([]string): the WednesdayList param.
*/
type WeeklyObject struct {
	FridayList    []string `json:"friday,omitempty"`
	MondayList    []string `json:"monday,omitempty"`
	SaturdayList  []string `json:"saturday,omitempty"`
	SundayList    []string `json:"sunday,omitempty"`
	ThursdayList  []string `json:"thursday,omitempty"`
	TuesdayList   []string `json:"tuesday,omitempty"`
	WednesdayList []string `json:"wednesday,omitempty"`
}
