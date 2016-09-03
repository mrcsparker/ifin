package model

type ControllerStatusDTO struct {
	ActiveThreadCount       int32  `json:"activeThreadCount"`       // The number of active threads in the NiFi.
	Queued                  string `json:"queued"`                  // The number of flowfilew queued in the NiFi.
	FlowFilesQueued         int32  `json:"flowFilesQueued"`         // The number of FlowFiles queued across the entire flow
	BytesQueued             int64  `json:"bytesQueued"`             // The size of the FlowFiles queued across the entire flow
	RunningCount            int32  `json:"runningCount"`            // The number of running components in the NiFi.
	StoppedCount            int32  `json:"stoppedCount"`            // The number of stopped components in the NiFi.
	InvalidCount            int32  `json:"invalidCount"`            // The number of invalid components in the NiFi.
	DisabledCount           int32  `json:"disabledCount"`           // The number of disabled components in the NiFi.
	ActiveRemotePortCount   int32  `json:"activeRemotePortCount"`   // The number of active remote ports in the NiFi.
	InactiveRemotePortCount int32  `json:"inactiveRemotePortCount"` // The number of inactive remote ports in the NiFi.
}
