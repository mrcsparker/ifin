package model

type StatusHistoryEntity struct {
	StatusHistory []StatusHistoryDTO `json:"statusHistory"` //
	CanRead       bool               `json:"canRead"`       // Indicates whether the user can read a given resource.
}
