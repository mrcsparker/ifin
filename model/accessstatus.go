package model

type AccessStatus struct {
	Identity string `json:"identity"` // The user identity.
	Username string `json:"username"` // The username.
	Status   string `json:"status"`   // The user access status.
	Message  string `json:"message"`  // Additional details about the user access status.
}
