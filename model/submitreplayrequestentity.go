package model

// attribute: copySnippetRequestEntity
type SubmitReplayRequestEntity struct {
	EventId       int64  `json:"eventId"`       // The event identifier
	ClusterNodeId string `json:"clusterNodeId"` // The identifier of the node where to submit the replay request.
}
