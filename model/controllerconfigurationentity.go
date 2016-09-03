package model

// attribute: controllerConfigurationEntity
type ControllerConfigurationEntity struct {
	Revision    RevisionDTO                `json:"revision"`    // The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Permissions PermissionsDTO             `json:"permissions"` // The permissions for this component.
	Component   ControllerConfigurationDTO `json:"component"`   // The controller configuration.
}
