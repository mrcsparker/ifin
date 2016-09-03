package model

// attribute: entity
type PortStatusSnapshotEntity struct {
	Id                 string                `json:"id"` // The id of the port.
	PortStatusSnapshot PortStatusSnapshotDTO `json:"portStatusSnapshot"`
	CanRead            bool                  `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
}
