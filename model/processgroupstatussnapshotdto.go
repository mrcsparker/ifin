package model

type ProcessGroupStatusSnapshotDTO struct {
	Id                                string                                   `json:"id"`                                // The id of the process group.
	Name                              string                                   `json:"name"`                              // The name of this process group.
	ConnectionStatusSnapshots         []ConnectionStatusSnapshotEntity         `json:"connectionStatusSnapshots"`         // The status of all conenctions in the process group.
	ProcessorStatusSnapshots          []ProcessorStatusSnapshotEntity          `json:"processorStatusSnapshots"`          // The status of all processors in the process group.
	ProcessGroupStatusSnapshots       []ProcessGroupStatusSnapshotEntity       `json:"processGroupStatusSnapshots"`       // The status of all process groups in the process group.
	RemoteProcessGroupStatusSnapshots []RemoteProcessGroupStatusSnapshotEntity `json:"remoteProcessGroupStatusSnapshots"` // The status of all remote process groups in the process group.
	InputPortStatusSnapshots          []PortStatusSnapshotEntity               `json:"inputPortStatusSnapshots"`          // The status of all input ports in the process group.
	OutputPortStatusSnapshots         []PortStatusSnapshotEntity               `json:"outputPortStatusSnapshots"`         // The status of all output ports in the process group.
	FlowFilesIn                       int32                                    `json:"flowFilesIn"`                       // The number of FlowFiles that have come into this ProcessGroup in the last 5 minutes
	BytesIn                           int64                                    `json:"bytesIn"`                           // The number of bytes that have come into this ProcessGroup in the last 5 minutes
	Input                             string                                   `json:"input"`                             // The input count/size for the process group in the last 5 minutes (pretty printed).
	FlowFilesQueued                   int32                                    `json:"flowFilesQueued"`                   // The number of FlowFiles that are queued up in this ProcessGroup right now
	BytesQueued                       int64                                    `json:"bytesQueued"`                       // The number of bytes that are queued up in this ProcessGroup right now
	Queued                            string                                   `json:"queued"`                            // The count/size that is queued in the the process group.
	QueuedCount                       string                                   `json:"queuedCount"`                       // The count that is queued for the process group.
	QueuedSize                        string                                   `json:"queuedSize"`                        // The size that is queued for the process group.
	BytesRead                         int64                                    `json:"bytesRead"`                         // The number of bytes read by components in this ProcessGroup in the last 5 minutes
	Read                              string                                   `json:"read"`                              // The number of bytes read in the last 5 minutes.
	BytesWritten                      int64                                    `json:"bytesWritten"`                      // The number of bytes written by components in this ProcessGroup in the last 5 minutes
	Written                           string                                   `json:"written"`                           // The number of bytes written in the last 5 minutes.
	FlowFilesOut                      int32                                    `json:"flowFilesOut"`                      // The number of FlowFiles transferred out of this ProcessGroup in the last 5 minutes
	BytesOut                          int64                                    `json:"bytesOut"`                          // The number of bytes transferred out of this ProcessGroup in the last 5 minutes
	Output                            string                                   `json:"output"`                            // The output count/size for the process group in the last 5 minutes.
	FlowFilesTransferred              int32                                    `json:"flowFilesTransferred"`              // The number of FlowFiles transferred in this ProcessGroup in the last 5 minutes
	BytesTransferred                  int64                                    `json:"bytesTransferred"`                  // The number of bytes transferred in this ProcessGroup in the last 5 minutes
	Transferred                       string                                   `json:"transferred"`                       // The count/size transferred to/frome queues in the process group in the last 5 minutes.
	BytesReceived                     int64                                    `json:"bytesReceived"`                     // The number of bytes received from external sources by components within this ProcessGroup in the last 5 minutes
	FlowFilesReceived                 int32                                    `json:"flowFilesReceived"`                 // The number of FlowFiles received from external sources by components within this ProcessGroup in the last 5 minutes
	Received                          string                                   `json:"received"`                          // The count/size sent to the process group in the last 5 minutes.
	BytesSent                         int64                                    `json:"bytesSent"`                         // The number of bytes sent to an external sink by components within this ProcessGroup in the last 5 minutes
	FlowFilesSent                     int32                                    `json:"flowFilesSent"`                     // The number of FlowFiles sent to an external sink by components within this ProcessGroup in the last 5 minutes
	Sent                              string                                   `json:"sent"`                              // The count/size sent from this process group in the last 5 minutes.
	ActiveThreadCount                 int32                                    `json:"activeThreadCount"`                 // The active thread count for this process group.
}
