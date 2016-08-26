package model

type ControllerServiceType struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
