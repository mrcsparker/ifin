package model

type CurrentUserEntity struct {
	Identity              string           `json:"identity"`              // The user identity being serialized.
	Anonymous             bool             `json:"anonymous"`             // Whether the current user is anonymous.
	ProvenancePermissions []PermissionsDTO `json:"provenancePermissions"` // Permissions for querying provenance.
	CountersPermissions   []PermissionsDTO `json:"countersPermissions"`   // Permissions for accessing counters.
	TenantsPermissions    []PermissionsDTO `json:"tenantsPermissions"`    // Permissions for accessing tenants.
	ControllerPermissions []PermissionsDTO `json:"controllerPermissions"` // Permissions for accessing the controller.
	PoliciesPermissions   []PermissionsDTO `json:"policiesPermissions"`   // Permissions for accessing the policies.
}
