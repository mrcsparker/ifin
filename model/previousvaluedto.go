package model

type PreviousValueDTO struct {
	PreviousValue string `json:"previousValue"` // The previous value.
	Timestamp     string `json:"timestamp"`     // The timestamp when the value was modified.
	UserIdentity  string `json:"userIdentity"`  // The user who changed the previous value.
}
