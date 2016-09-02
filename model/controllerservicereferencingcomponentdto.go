package model

type ControllerServiceReferencingComponentDTO struct {
	GroupId               string                `json:"groupId"`               // The group id for the component referencing a controller service. If this component is another controller service or a reporting task, this field is blank.
	Id                    string                `json:"id"`                    // The id of the component referencing a controller service.
	Name                  string                `json:"name"`                  // The name of the component referencing a controller service.
	Type                  string                `json:"type"`                  // The type of the component referencing a controller service.
	State                 string                `json:"state"`                 // The state of a processor or reporting task referencing a controller service. If this component is another controller service, this field is blank.
	Properties            string                `json:"properties"`            // The properties for the component.
	Descriptors           PropertyDescriptorDTO `json:"descriptors"`           // The descriptors for the componet properties.
	ValidationErrors      []string              `json:"validationErrors"`      // The validation errors for the component.
	ReferenceType         string                `json:"referenceType"`         // The type of reference this is.
	ActiveThreadCount     int                   `json:"activeThreadCount"`     // The number of active threads for the referencing component.
	ReferenceCycle        bool                  `json:"referenceCycle"`        // If the referencing component represents a controller service, this indicates whether it has already been represented in this hierarchy.
	ReferencingComponents []Set                 `json:"referencingComponents"` // If the referencing component represents a controller service, these are the components that referenc it.
}
