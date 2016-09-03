package model

type NodeConnectionStatusSnapshotDTO struct {
	NodeId         string                      `json:"nodeId"`         // The unique ID that identifies the node
	Address        string                      `json:"address"`        // The API address of the node
	ApiPort        int32                       `json:"apiPort"`        // The API port used to communicate with the node
	StatusSnapshot ConnectionStatusSnapshotDTO `json:"statusSnapshot"` // The connection status snapshot from the node.
}
