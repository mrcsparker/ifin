package model

type ComponentHistory struct {
	ComponentId     string      `json:"componentId"`     // The component id.
	PropertyHistory interface{} `json:"propertyHistory"` // The history for the properties of the component.
}
