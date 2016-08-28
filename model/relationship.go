package model

type Relationship struct {
	Name          string `json:"name"`          // The relationship name.
	Description   string `json:"name"`          // The relationship description.
	AutoTerminate bool   `json:"autoTerminate"` // Whether or not flowfiles sent to this relationship should auto terminate.
}
