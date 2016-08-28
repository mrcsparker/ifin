package model

type DocumentedType struct {
	Type        string   `json:"type"`        // The fulley qualified name of the type.
	Description string   `json:"description"` // The description of the type.
	Tags        []string `json:"tags"`        // The tags associated with this type.
}
