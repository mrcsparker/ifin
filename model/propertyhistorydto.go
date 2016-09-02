package model

type PropertyHistoryDTO struct {
	PreviousValues []PreviousValueDTO `json:"previousValues"` // Previous values for a given property.
}
