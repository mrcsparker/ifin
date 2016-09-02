package model

type StatusDescriptorDTO struct {
	Field       string `json:"field"`       // The name of the status field.
	Label       string `json:"label"`       // The label for the status field.
	Description string `json:"description"` // The description of the status field.
	Formatter   string `json:"formatter"`   // The formatter for the status descriptor.
}
