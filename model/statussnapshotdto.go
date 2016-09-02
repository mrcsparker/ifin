package model

type StatusSnapshotDTO struct {
	Timestamp     string `json:"timestamp"`     // The timestamp of the snapshot.
	StatusMetrics int    `json:"statusMetrics"` // The status metrics.
}
