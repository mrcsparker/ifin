package model

type NodeCountersSnapshotDTO struct {
	NodeId   string              `json:"nodeId"`   // The unique ID that identifies the node
	Address  string              `json:"address"`  // The API address of the node
	ApiPort  int32               `json:"apiPort"`  // The API port used to communicate with the node
	Snapshot CountersSnapshotDTO `json:"snapshot"` // The counters from the node.
}
