package model

// attribute: processGroupStatusEntity
type ProcessGroupStatusEntity struct {
	ProcessGroupStatus ProcessGroupStatusDTO `json:"processGroupStatus"`
	CanRead            bool                  `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
}
