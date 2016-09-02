package model

type ProcessGroupStatusSnapshotEntity struct {
	Id      string `json:"id"`      // The id of the process group.
	CanRead bool   `json:"canRead"` // Indicates whether the user can read a given resource.
}
