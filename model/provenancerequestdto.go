package model

type ProvenanceRequestDTO struct {
	SearchTerms     string `json:"searchTerms"`     // The search terms used to perform the search.
	ClusterNodeId   string `json:"clusterNodeId"`   // The id of the node in the cluster where this provenance originated.
	StartDate       string `json:"startDate"`       // The earliest event time to include in the query.
	EndDate         string `json:"endDate"`         // The latest event time to include in the query.
	MinimumFileSize string `json:"minimumFileSize"` // The minimum file size to include in the query.
	MaximumFileSize string `json:"maximumFileSize"` // The maximum file size to include in the query.
	MaxResults      int32  `json:"maxResults"`      // The maximum number of results to include.
}
