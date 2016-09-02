package model

type DropRequestDTO struct {
	Id               string `json:"id"`               // The id for this drop request.
	Uri              string `json:"uri"`              // The URI for future requests to this drop request.
	SubmissionTime   string `json:"submissionTime"`   // The timestamp when the query was submitted.
	LastUpdated      string `json:"lastUpdated"`      // The last time this drop request was updated.
	PercentCompleted int    `json:"percentCompleted"` // The current percent complete.
	Finished         bool   `json:"finished"`         // Whether the query has finished.
	FailureReason    string `json:"failureReason"`    // The reason, if any, that this drop request failed.
	CurrentCount     int    `json:"currentCount"`     // The number of flow files currently queued.
	CurrentSize      int    `json:"currentSize"`      // The size of flow files currently queued in bytes.
	Current          string `json:"current"`          // The count and size of flow files currently queued.
	OriginalCount    int    `json:"originalCount"`    // The number of flow files to be dropped as a result of this request.
	OriginalSize     int    `json:"originalSize"`     // The size of flow files to be dropped as a result of this request in bytes.
	Original         string `json:"original"`         // The count and size of flow files to be dropped as a result of this request.
	DroppedCount     int    `json:"droppedCount"`     // The number of flow files that have been dropped thus far.
	DroppedSize      int    `json:"droppedSize"`      // The size of flow files that have been dropped thus far in bytes.
	Dropped          string `json:"dropped"`          // The count and size of flow files that have been dropped thus far.
	State            string `json:"state"`            // The current state of the drop request.
}
