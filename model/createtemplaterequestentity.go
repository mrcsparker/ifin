package model

type CreateTemplateRequestEntity struct {
	Name        string `json:"name"`        // The name of the template.
	Description string `json:"description"` // The description of the template.
	SnippetId   string `json:"snippetId"`   // The identifier of the snippet.
}
