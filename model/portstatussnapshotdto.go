package model

type PortStatusSnapshotDTO struct {
	Id                string `json:"id"`                // The id of the port.
	GroupId           string `json:"groupId"`           // The id of the parent process group of the port.
	Name              string `json:"name"`              // The name of the port.
	ActiveThreadCount int    `json:"activeThreadCount"` // The active thread count for the port.
	FlowFilesIn       int    `json:"flowFilesIn"`       // The number of FlowFiles that have been accepted in the last 5 minutes.
	BytesIn           int    `json:"bytesIn"`           // The size of hte FlowFiles that have been accepted in the last 5 minutes.
	Input             string `json:"input"`             // The count/size of flowfiles that have been accepted in the last 5 minutes.
	FlowFilesOut      int    `json:"flowFilesOut"`      // The number of FlowFiles that have been processed in the last 5 minutes.
	BytesOut          int    `json:"bytesOut"`          // The number of bytes that have been processed in the last 5 minutes.
	Output            string `json:"output"`            // The count/size of flowfiles that have been processed in the last 5 minutes.
	Transmitting      bool   `json:"transmitting"`      // Whether the port has incoming or outgoing connections to a remote NiFi.
	RunStatus         string `json:"runStatus"`         // The run status of the port.
}
