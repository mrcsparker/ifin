package model

type SubmitReplayRequestEntity struct {
	EventId       int    `json:"eventId"`       // The event identifier
	ClusterNodeId string `json:"clusterNodeId"` // The identifier of the node where to submit the replay request.
}
