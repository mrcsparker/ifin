package model

type RemoteProcessGroupStatusSnapshotDTO struct {
	Id                 string `json:"id"`                 // The id of the remote process group.
	GroupId            string `json:"groupId"`            // The id of the parent process group the remote process group resides in.
	Name               string `json:"name"`               // The name of the remote process group.
	TargetUri          string `json:"targetUri"`          // The URI of the target system.
	TransmissionStatus string `json:"transmissionStatus"` // The transmission status of the remote process group.
	ActiveThreadCount  int32  `json:"activeThreadCount"`  // The number of active threads for the remote process group.
	FlowFilesSent      int32  `json:"flowFilesSent"`      // The number of FlowFiles sent to the remote process group in the last 5 minutes.
	BytesSent          int64  `json:"bytesSent"`          // The size of the FlowFiles sent to the remote process group in the last 5 minutes.
	Sent               string `json:"sent"`               // The count/size of the flowfiles sent to the remote process group in the last 5 minutes.
	FlowFilesReceived  int32  `json:"flowFilesReceived"`  // The number of FlowFiles received from the remote process group in the last 5 minutes.
	BytesReceived      int64  `json:"bytesReceived"`      // The size of the FlowFiles received from the remote process group in the last 5 minutes.
	Received           string `json:"received"`           // The count/size of the flowfiles received from the remote process group in the last 5 minutes.
}
