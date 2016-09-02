package model

type PortStatusEntity struct {
	PortStatus []PortStatusDTO `json:"portStatus"`
	CanRead    bool            `json:"canRead"` // Indicates whether the user can read a given resource. *Read Only*
}
