package model

type LineageRequestDTO struct {
	EventId            int    `json:"eventId"`            //
	LineageRequestType string `json:"lineageRequestType"` // The type of lineage request. PARENTS will return the lineage for the flowfiles that are parents of the specified event. CHILDREN will return the lineage for the flowfiles that are children of the specified event. FLOWFILE will return the lineage for the specified flowfile.
	Uuid               string `json:"uuid"`               // The uuid that was used to generate the lineage.
	ClusterNodeId      string `json:"clusterNodeId"`      // The id of the node where this lineage originated if clustered.
}
