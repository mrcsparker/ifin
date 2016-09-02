package model

type AllowableValueEntity struct {
	AllowableValue []AllowableValueDTO `json:"allowableValue"`
	CanRead        bool                `json:"canRead"` // Indicates whether the user can read a given resource. *Read Only*
}