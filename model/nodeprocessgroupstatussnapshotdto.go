package model

type NodeProcessGroupStatusSnapshotDTO struct {
	NodeId         string                        `json:"nodeId"`         // The unique ID that identifies the node
	Address        string                        `json:"address"`        // The API address of the node
	ApiPort        int32                         `json:"apiPort"`        // The API port used to communicate with the node
	StatusSnapshot ProcessGroupStatusSnapshotDTO `json:"statusSnapshot"` // The process group status snapshot from the node.
}
