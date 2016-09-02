package model

type ComponentHistoryDTO struct {
	ComponentId     string             `json:"componentId"`     // The component id.
	PropertyHistory PropertyHistoryDTO `json:"propertyHistory"` // The history for the properties of the component.
}
