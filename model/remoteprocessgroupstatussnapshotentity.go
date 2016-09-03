package model

// attribute: entity
type RemoteProcessGroupStatusSnapshotEntity struct {
	Id                               string                              `json:"id"` // The id of the remote processo group.
	RemoteProcessGroupStatusSnapshot RemoteProcessGroupStatusSnapshotDTO `json:"remoteProcessGroupStatusSnapshot"`
	CanRead                          bool                                `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
}
