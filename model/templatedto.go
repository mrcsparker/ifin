package model

type TemplateDTO struct {
	Uri             string           `json:"uri"`             // The URI for the template.
	Id              string           `json:"id"`              // The id of the template.
	GroupId         string           `json:"groupId"`         // The id of the Process Group that the template belongs to.
	Name            string           `json:"name"`            // The name of the template.
	Description     string           `json:"description"`     // The description of the template.
	Timestamp       string           `json:"timestamp"`       // The timestamp when this template was created.
	EncodingVersion string           `json:"encodingVersion"` // The encoding version of this template.
	Snippet         []FlowSnippetDTO `json:"snippet"`         // The contents of the template.
}
