package model

type ProvenanceDTO struct {
	Id               string               `json:"id"`               // The id of the provenance query.
	Uri              string               `json:"uri"`              // The URI for this query. Used for obtaining/deleting the request at a later time
	SubmissionTime   string               `json:"submissionTime"`   // The timestamp when the query was submitted.
	Expiration       string               `json:"expiration"`       // The timestamp when the query will expire.
	PercentCompleted int32                `json:"percentCompleted"` // The current percent complete.
	Finished         bool                 `json:"finished"`         // Whether the query has finished.
	Request          ProvenanceRequestDTO `json:"request"`          // The provenance request.
	Results          ProvenanceResultsDTO `json:"results"`          // The provenance results.
}
