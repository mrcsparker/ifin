package model

type FlowFileSummaryDTO struct {
	Uri                string `json:"uri"`                // The URI that can be used to access this FlowFile.
	Uuid               string `json:"uuid"`               // The FlowFile UUID.
	Filename           string `json:"filename"`           // The FlowFile filename.
	Position           int32  `json:"position"`           // The FlowFile's position in the queue.
	Size               int64  `json:"size"`               // The FlowFile file size.
	QueuedDuration     int64  `json:"queuedDuration"`     // How long this FlowFile has been enqueued.
	LineageDuration    int64  `json:"lineageDuration"`    // Duration since the FlowFile's greatest ancestor entered the flow.
	ClusterNodeId      string `json:"clusterNodeId"`      // The id of the node where this FlowFile resides.
	ClusterNodeAddress string `json:"clusterNodeAddress"` // The label for the node where this FlowFile resides.
	Penalized          bool   `json:"penalized"`          // If the FlowFile is penalized.
}
