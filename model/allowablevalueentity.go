package model

// attribute: entity
type AllowableValueEntity struct {
	AllowableValue AllowableValueDTO `json:"allowableValue"`
	CanRead        bool              `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
}
