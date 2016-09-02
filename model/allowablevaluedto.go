package model

type AllowableValueDTO struct {
	DisplayName string `json:"displayName"` // A human readable value that is allowed for the property descriptor.
	Value       string `json:"value"`       // A value that is allowed for the property descriptor.
	Description string `json:"description"` // A description for this allowable value.
}
