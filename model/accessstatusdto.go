package model

type AccessStatusDTO struct {
	Identity string `json:"identity"` // The user identity. *Read Only*
	Status   string `json:"status"`   // The user access status. *Read Only*
	Message  string `json:"message"`  // Additional details about the user access status. *Read Only*
}
