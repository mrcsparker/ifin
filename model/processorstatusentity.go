package model

// attribute: processorStatusEntity
type ProcessorStatusEntity struct {
	ProcessorStatus ProcessorStatusDTO `json:"processorStatus"`
	CanRead         bool               `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
}
