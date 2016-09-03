package model

// attribute: connectionStatusEntity
type ConnectionStatusEntity struct {
	ConnectionStatus ConnectionStatusDTO `json:"connectionStatus"`
	CanRead          bool                `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
}
