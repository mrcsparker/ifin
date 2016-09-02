package model

type ConnectionStatusEntity struct {
	ConnectionStatus []ConnectionStatusDTO `json:"connectionStatus"`
	CanRead          bool                  `json:"canRead"` // Indicates whether the user can read a given resource. *Read Only*
}
