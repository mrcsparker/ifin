package model
type ControllerServiceDTO struct {
Id string `json:"id"` // The id of the component.
ParentGroupId string `json:"parentGroupId"` // The id of parent process group of this component if applicable.
Position []PositionDTO `json:"position"` // The position of this component in the UI if applicable.
Name string `json:"name"` // The name of the controller service.
Type string `json:"type"` // The type of the controller service.
Comments string `json:"comments"` // The comments for the controller service.
State string `json:"state"` // The state of the controller service.
PersistsState bool `json:"persistsState"` // Whether the controller service persists state.


CustomUiUrl string `json:"customUiUrl"` // The URL for the controller services custom configuration UI if applicable.
AnnotationData string `json:"annotationData"` // The annontation for the controller service. This is how the custom UI relays configuration to the controller service.
ReferencingComponents []ControllerServiceReferencingComponentEntity `json:"referencingComponents"` // All components referencing this controller service.
ValidationErrors [] `json:"validationErrors"` // The validation errors from the controller service. These validation errors represent the problems with the controller service that must be resolved before it can be enabled.
}
