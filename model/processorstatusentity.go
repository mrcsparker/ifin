package model

type ProcessorStatusEntity struct {
	ProcessorStatus []ProcessorStatusDTO `json:"processorStatus"` //
	CanRead         bool                 `json:"canRead"`         // Indicates whether the user can read a given resource.
}
