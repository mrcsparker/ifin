package model

// attribute: entity
type ProcessorStatusSnapshotEntity struct {
	Id                      string                     `json:"id"` // The id of the processor.
	ProcessorStatusSnapshot ProcessorStatusSnapshotDTO `json:"processorStatusSnapshot"`
	CanRead                 bool                       `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
}
