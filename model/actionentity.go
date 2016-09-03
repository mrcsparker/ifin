package model

// attribute: actionEntity
type ActionEntity struct {
	Id        int32     `json:"id"`
	Timestamp string    `json:"timestamp"`
	SourceId  string    `json:"sourceId"`
	CanRead   bool      `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
	Action    ActionDTO `json:"action"`
}
