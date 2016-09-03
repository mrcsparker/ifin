package model

type UserDTO struct {
	Id             string                      `json:"id"`             // The id of the component.
	ParentGroupId  string                      `json:"parentGroupId"`  // The id of parent process group of this component if applicable.
	Position       PositionDTO                 `json:"position"`       // The position of this component in the UI if applicable.
	Identity       string                      `json:"identity"`       // The identity of the tenant.
	UserGroups     []TenantEntity              `json:"userGroups"`     // [Read Only] The groups to which the user belongs. This field is read only and it provided for convenience.
	AccessPolicies []AccessPolicySummaryEntity `json:"accessPolicies"` // [Read Only] The access policies this user belongs to.
}
