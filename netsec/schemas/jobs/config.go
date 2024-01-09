package jobs

// This code is automatically generated.
// Manual changes made will be overwritten upon SDK generation.

// Schema: #/components/schemas/jobs

/*
Config object.

ShortName: qFWAgJG
Parent chains:
*

Args:

Param Description (string, required): A description provided by the administrator or service account

Param DeviceName (string, required): The name of the device

Param EndTs (string, required): The timestamp indicating when the job was finished

Param Id (int64, required): The job ID

Param JobResult (int64, required): The job result

Param JobStatus (int64, required): The current status of the job

Param JobType (int64, required): The job type

Param ParentId (int64, required): The parent job ID

Param Percent (int64, required): Job completion percentage Value must be less than or equal to 100.

Param ResultStr (string, required): The result of the job String must be one of these: `"OK"`, `"FAIL"`, `"PEND"`, `"WAIT"`, `"CANCELLED"`.

Param StartTs (string, required): The timestamp indicating when the job was created

Param StatusStr (string, required): The current status of the job String must be one of these: `"ACT"`, `"FIN"`, `"PEND"`, `"PUSHSENT"`, `"PUSHFAIL"`.

Param Summary (string, required): The completion summary of the job

Param TypeStr (string, required): The job type String must be one of these: `"CommitAll"`, `"CommitAndPush"`, `"NGFW-Bootstrap-Push"`, `"Validate"`.

Param Uname (string, required): The administrator or service account that created the job
*/
type Config struct {
	Description string `json:"description"`
	DeviceName  string `json:"device_name"`
	EndTs       string `json:"end_ts"`
	Id          int64  `json:"id"`
	JobResult   int64  `json:"job_result"`
	JobStatus   int64  `json:"job_status"`
	JobType     int64  `json:"job_type"`
	ParentId    int64  `json:"parent_id"`
	Percent     int64  `json:"percent"`
	ResultStr   string `json:"result_str"`
	StartTs     string `json:"start_ts"`
	StatusStr   string `json:"status_str"`
	Summary     string `json:"summary"`
	TypeStr     string `json:"type_str"`
	Uname       string `json:"uname"`
}
