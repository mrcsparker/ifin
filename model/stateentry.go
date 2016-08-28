package model

type StateEntry struct {
	Key                string `json:"key"`                // The key for this state.
	Value              string `json:"value"`              // The value for this state.
	ClusterNodeId      string `json:"clusterNodeId"`      // The identifier for the node where the state originated.
	ClusterNodeAddress string `json:"clusterNodeAddress"` // The label for the node where the state originated.
}
