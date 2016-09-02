package model

type ClusterDTO struct {
	Nodes     []NodeDTO `json:"nodes"`     // The collection of nodes that are part of the cluster.
	Generated string    `json:"generated"` // The timestamp the report was generated.
}
