package model

// attribute: accessStatus
type AccessStatusDTO struct {
	Identity string `json:"identity"` // [Read Only] The user identity.
	Status   string `json:"status"`   // [Read Only] The user access status.
	Message  string `json:"message"`  // [Read Only] Additional details about the user access status.
}
