package model

type ProcessGroupStatusEntity struct {
	ProcessGroupStatus []ProcessGroupStatusDTO `json:"processGroupStatus"`
	CanRead            bool                    `json:"canRead"` // Indicates whether the user can read a given resource. *Read Only*
}
