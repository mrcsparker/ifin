package model

type ProcessorStatusSnapshotEntity struct {
	Id                      string                       `json:"id"` // The id of the processor.
	ProcessorStatusSnapshot []ProcessorStatusSnapshotDTO `json:"processorStatusSnapshot"`
	CanRead                 bool                         `json:"canRead"` // Indicates whether the user can read a given resource. *Read Only*
}
