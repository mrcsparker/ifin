package model

type ProcessorStatusSnapshotDTO struct {
	Id                 string `json:"id"`                 // The id of the processor.
	GroupId            string `json:"groupId"`            // The id of the parent process group to which the processor belongs.
	Name               string `json:"name"`               // The name of the prcessor.
	Type               string `json:"type"`               // The type of the processor.
	RunStatus          string `json:"runStatus"`          // The state of the processor.
	BytesRead          int    `json:"bytesRead"`          // The number of bytes read by this Processor in the last 5 mintues
	BytesWritten       int    `json:"bytesWritten"`       // The number of bytes written by this Processor in the last 5 minutes
	Read               string `json:"read"`               // The number of bytes read in the last 5 minutes.
	Written            string `json:"written"`            // The number of bytes written in the last 5 minutes.
	FlowFilesIn        int    `json:"flowFilesIn"`        // The number of FlowFiles that have been accepted in the last 5 minutes
	BytesIn            int    `json:"bytesIn"`            // The size of the FlowFiles that have been accepted in the last 5 minutes
	Input              string `json:"input"`              // The count/size of flowfiles that have been accepted in the last 5 minutes.
	FlowFilesOut       int    `json:"flowFilesOut"`       // The number of FlowFiles transferred to a Connection in the last 5 minutes
	BytesOut           int    `json:"bytesOut"`           // The size of the FlowFiles transferred to a Connection in the last 5 minutes
	Output             string `json:"output"`             // The count/size of flowfiles that have been processed in the last 5 minutes.
	TaskCount          int    `json:"taskCount"`          // The number of times this Processor has run in the last 5 minutes
	TasksDurationNanos int    `json:"tasksDurationNanos"` // The number of nanoseconds that this Processor has spent running in the last 5 minutes
	Tasks              string `json:"tasks"`              // The total number of task this connectable has completed over the last 5 minutes.
	TasksDuration      string `json:"tasksDuration"`      // The total duration of all tasks for this connectable over the last 5 minutes.
	ActiveThreadCount  int    `json:"activeThreadCount"`  // The number of threads currently executing in the processor.
}
