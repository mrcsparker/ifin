package model

type ProvenanceLinkDTO struct {
	SourceId     string `json:"sourceId"`     // The source node id of the link.
	TargetId     string `json:"targetId"`     // The target node id of the link.
	FlowFileUuid string `json:"flowFileUuid"` // The flowfile uuid that traversed the link.
	Timestamp    string `json:"timestamp"`    // The timestamp of the link (based on the destination).
	Millis       int    `json:"millis"`       // The timestamp of this link in milliseconds.
}
