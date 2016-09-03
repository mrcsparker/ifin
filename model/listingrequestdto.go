package model

type ListingRequestDTO struct {
	Id                 string               `json:"id"`                 // The id for this listing request.
	Uri                string               `json:"uri"`                // The URI for future requests to this listing request.
	SubmissionTime     string               `json:"submissionTime"`     // The timestamp when the query was submitted.
	LastUpdated        string               `json:"lastUpdated"`        // The last time this listing request was updated.
	PercentCompleted   int32                `json:"percentCompleted"`   // The current percent complete.
	Finished           bool                 `json:"finished"`           // Whether the query has finished.
	FailureReason      string               `json:"failureReason"`      // The reason, if any, that this listing request failed.
	MaxResults         int32                `json:"maxResults"`         // The maximum number of FlowFileSummary objects to return
	State              string               `json:"state"`              // The current state of the listing request.
	QueueSize          QueueSizeDTO         `json:"queueSize"`          // The size of the queue
	FlowFileSummaries  []FlowFileSummaryDTO `json:"flowFileSummaries"`  // The FlowFile summaries. The summaries will be populated once the request has completed.
	SourceRunning      bool                 `json:"sourceRunning"`      // Whether the source of the connection is running
	DestinationRunning bool                 `json:"destinationRunning"` // Whether the destination of the connection is running
}
