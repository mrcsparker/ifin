package model

// attribute: entity
type ProcessGroupStatusSnapshotEntity struct {
	Id      string `json:"id"`      // The id of the process group.
	CanRead bool   `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
}
