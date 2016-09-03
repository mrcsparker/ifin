package model

// attribute: statusHistoryEntity
type StatusHistoryEntity struct {
	StatusHistory StatusHistoryDTO `json:"statusHistory"`
	CanRead       bool             `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
}
