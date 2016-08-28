package model

type StatusSnapshot struct {
	Timestamp     string      `json:"timestamp"` // The timestamp of the snapshot.
	StatusMetrics interface{} `json:"timestamp"` // The status metrics.
}
