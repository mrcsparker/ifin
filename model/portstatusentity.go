package model

// attribute: portStatusEntity
type PortStatusEntity struct {
	PortStatus PortStatusDTO `json:"portStatus"`
	CanRead    bool          `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
}
