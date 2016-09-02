package model

type ControllerStatusDTO struct {
	ActiveThreadCount       int    `json:"activeThreadCount"`       // The number of active threads in the NiFi.
	Queued                  string `json:"queued"`                  // The number of flowfilew queued in the NiFi.
	FlowFilesQueued         int    `json:"flowFilesQueued"`         // The number of FlowFiles queued across the entire flow
	BytesQueued             int    `json:"bytesQueued"`             // The size of the FlowFiles queued across the entire flow
	RunningCount            int    `json:"runningCount"`            // The number of running components in the NiFi.
	StoppedCount            int    `json:"stoppedCount"`            // The number of stopped components in the NiFi.
	InvalidCount            int    `json:"invalidCount"`            // The number of invalid components in the NiFi.
	DisabledCount           int    `json:"disabledCount"`           // The number of disabled components in the NiFi.
	ActiveRemotePortCount   int    `json:"activeRemotePortCount"`   // The number of active remote ports in the NiFi.
	InactiveRemotePortCount int    `json:"inactiveRemotePortCount"` // The number of inactive remote ports in the NiFi.
}
