package model

type ProvenanceNodeDTO struct {
	Id                    string   `json:"id"`                    // The id of the node.
	FlowFileUuid          string   `json:"flowFileUuid"`          // The uuid of the flowfile associated with the provenance event.
	ParentUuids           []string `json:"parentUuids"`           // The uuid of the parent flowfiles of the provenance event.
	ChildUuids            []string `json:"childUuids"`            // The uuid of the childrent flowfiles of the provenance event.
	ClusterNodeIdentifier string   `json:"clusterNodeIdentifier"` // The identifier of the node that this event/flowfile originated from.
	Type                  string   `json:"type"`                  // The type of the node.
	EventType             string   `json:"eventType"`             // If the type is EVENT, this is the type of event.
	Millis                int64    `json:"millis"`                // The timestamp of the node in milliseconds.
	Timestamp             string   `json:"timestamp"`             // The timestamp of the node formatted.
}
