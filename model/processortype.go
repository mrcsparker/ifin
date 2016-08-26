package model

type ProcessorType struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
