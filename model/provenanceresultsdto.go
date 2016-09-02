package model

type ProvenanceResultsDTO struct {
	ProvenanceEvents []ProvenanceEventDTO `json:"provenanceEvents"` // The provenance events that matched the search criteria.
	Total            string               `json:"total"`            // The total number of results formatted.
	TotalCount       int                  `json:"totalCount"`       // The total number of results.
	Generated        string               `json:"generated"`        // Then the search was performed.
	OldestEvent      string               `json:"oldestEvent"`      // The oldest event available in the provenance repository.
	TimeOffset       int                  `json:"timeOffset"`       // The time offset of the server that's used for event time.
	Errors           []string             `json:"errors"`           // Any errors that occurred while performing the provenance request.
}
