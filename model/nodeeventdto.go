package model

type NodeEventDTO struct {
	Timestamp string `json:"timestamp"` // The timestamp of the node event.
	Category  string `json:"category"`  // The category of the node event.
	Message   string `json:"message"`   // The message in the node event.
}
