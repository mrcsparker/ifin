package model

type NodeStatusSnapshotsDTO struct {
	NodeId          string              `json:"nodeId"`          // The id of the node.
	Address         string              `json:"address"`         // The node's host/ip address.
	ApiPort         int32               `json:"apiPort"`         // The port the node is listening for API requests.
	StatusSnapshots []StatusSnapshotDTO `json:"statusSnapshots"` // A list of StatusSnapshotDTO objects that provide the actual metric values for the component for this node.
}
