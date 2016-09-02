package model

type UpdateControllerServiceReferenceRequestEntity struct {
	Id                            string      `json:"id"`                            // The identifier of the Controller Service.
	State                         string      `json:"state"`                         // The new state of the references for the controller service.
	ReferencingComponentRevisions RevisionDTO `json:"referencingComponentRevisions"` // The revisions for all referencing components.
}
