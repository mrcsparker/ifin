package model

type AccessStatusDTO struct {
	Identity string `json:"identity"` // The user identity.
	Status   string `json:"status"`   // The user access status.
	Message  string `json:"message"`  // Additional details about the user access status.
}
