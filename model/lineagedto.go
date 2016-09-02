package model

type LineageDTO struct {
	Id               string              `json:"id"`               // The id of this lineage query.
	Uri              string              `json:"uri"`              // The URI for this lineage query for later retrieval and deletion.
	SubmissionTime   string              `json:"submissionTime"`   // When the lineage query was submitted.
	Expiration       string              `json:"expiration"`       // When the lineage query will expire.
	PercentCompleted int                 `json:"percentCompleted"` // The percent complete for the lineage query.
	Finished         bool                `json:"finished"`         // Whether the lineage query has finished.
	Request          []LineageRequestDTO `json:"request"`          // The initial lineage result.
	Results          []LineageResultsDTO `json:"results"`          // The results of the lineage query.
}
