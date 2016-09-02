package model

type ActionDTO struct {
	Id               int                   `json:"id"`               // The action id.
	UserIdentity     string                `json:"userIdentity"`     // The identity of the user that performed the action.
	Timestamp        string                `json:"timestamp"`        // The timestamp of the action.
	SourceId         string                `json:"sourceId"`         // The id of the source component.
	SourceName       string                `json:"sourceName"`       // The name of the source component.
	SourceType       string                `json:"sourceType"`       // The type of the source component.
	ComponentDetails []ComponentDetailsDTO `json:"componentDetails"` // The details of the source component.
	Operation        string                `json:"operation"`        // The operation that was performed.
	ActionDetails    []ActionDetailsDTO    `json:"actionDetails"`    // The details of the action.
}
