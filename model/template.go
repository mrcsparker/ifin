package model

type Template struct {
	Uri         string      `json:"uri"`         // The URI for the template.
	Id          string      `json:"id"`          // The id of the template.
	Name        string      `json:"name"`        // The name of the template.
	Description string      `json:"description"` // The description of the template.
	Timestamp   string      `json:"timestamp"`   // The timestamp when this template was created.
	Snippet     FlowSnippet `json:"snippet"`     // The contents of the template.
}
