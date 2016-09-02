package model

type AttributeDTO struct {
	Name          string `json:"name"`          // The attribute name.
	Value         string `json:"value"`         // The attribute value.
	PreviousValue string `json:"previousValue"` // The value of the attribute before the event took place.
}
