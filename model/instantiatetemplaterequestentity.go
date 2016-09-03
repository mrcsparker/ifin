package model

// attribute: instantiateTemplateRequestEntity
type InstantiateTemplateRequestEntity struct {
	TemplateId string  `json:"templateId"` // The identifier of the template.
	OriginX    float64 `json:"originX"`    // The x coordinate of the origin of the bounding box where the new components will be placed.
	OriginY    float64 `json:"originY"`    // The y coordinate of the origin of the bounding box where the new components will be placed.
}
