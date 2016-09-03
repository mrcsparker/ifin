package model

// attribute: templatesEntity
type TemplatesEntity struct {
	Templates []TemplateEntity `json:"templates"`
	Generated string           `json:"generated"`
}
