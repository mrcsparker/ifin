package model
type DocumentedTypeDTO struct {
Type string `json:"type"` // The fulley qualified name of the type.
Description string `json:"description"` // The description of the type.
Tags [] `json:"tags"` // The tags associated with this type.
}
