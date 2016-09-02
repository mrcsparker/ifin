package model

type RelationshipDTO struct {
	Name          string `json:"name"`          // The relationship name.
	Description   string `json:"description"`   // The relationship description.
	AutoTerminate bool   `json:"autoTerminate"` // Whether or not flowfiles sent to this relationship should auto terminate.
}
