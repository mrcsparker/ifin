package model

type Connectable struct {
	Id           string `json:"id"`           // The id of the connectable component.
	Type         string `json:"type"`         // The type of component the connectable is. Allowable values: PROCESSOR, REMOTE_INPUT_PORT, REMOTE_OUTPUT_PORT, INPUT_PORT, OUTPUT_PORT, FUNNEL
	GroupId      string `json:"groupId"`      // The id of the group that the connectable component resides in
	Name         string `json:"name"`         // The name of the connectable component
	Running      bool   `json:"running"`      // Reflects the current state of the connectable component.
	Transmitting bool   `json:"transmitting"` // If the connectable component represents a remote port, indicates if the target is configured to transmit.
	Exists       bool   `json:"exists"`       // If the connectable component represents a remote port, indicates if the target exists.
	Comments     string `json:"comments"`     // The comments for the connectable component.
}
