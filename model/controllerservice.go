package model

type ControllerService struct {
	Id                    string                                  `json:"id"`                    // The id of the component.
	Uri                   string                                  `json:"uri"`                   // The URI for futures requests to the component.
	Position              Position                                `json:"position"`              // The position of this component in the UI if applicable.
	ParentGroupId         string                                  `json:"parentGroupId"`         // The id of parent process group of this component if applicable.
	Name                  string                                  `json:"name"`                  // The name of the controller service.
	Type                  string                                  `json:"type"`                  // The type of the controller service.
	comments              string                                  `json:"comments"`              // The comments for the controller service.
	availability          string                                  `json:"availability"`          // Where the servcie is available. Allowable values: NCM, NODE
	state                 string                                  `json:"state"`                 // The state of the controller service. Allowable values: ENABLED, ENABLING, DISABLED, DISABLING
	persistsState         bool                                    `json:"persistsState"`         // Whether the controller service persists state.
	properties            interface{}                             `json:"properties"`            // The properties of the controller service.
	descriptors           interface{}                             `json:"descriptors"`           // The descriptors for the controller service properties.
	customUiUrl           string                                  `json:"customUiUrl"`           // The URL for the controller services custom configuration UI if applicable.
	annotationData        string                                  `json:"annotationData"`        // The annontation for the controller service. This is how the custom UI relays configuration to the controller service.
	referencingComponents []ControllerServiceReferencingComponent `json:"referencingComponents"` // All components referencing this controller service.
	validationErrors      []string                                `json:"validationErrors"`      // The validation errors from the controller service. These validation errors represent the problems with the controller service that must be resolved before it can be enabled.
}
