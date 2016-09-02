package model

type NodeRemoteProcessGroupStatusSnapshotDTO struct {
	NodeId         string                                `json:"nodeId"`         // The unique ID that identifies the node
	Address        string                                `json:"address"`        // The API address of the node
	ApiPort        int                                   `json:"apiPort"`        // The API port used to communicate with the node
	StatusSnapshot []RemoteProcessGroupStatusSnapshotDTO `json:"statusSnapshot"` // The remote process group status snapshot from the node.
}
