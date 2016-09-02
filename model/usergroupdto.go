package model

type UserGroupDTO struct {
	Id             string               `json:"id"`             // The id of the component.
	ParentGroupId  string               `json:"parentGroupId"`  // The id of parent process group of this component if applicable.
	Position       []PositionDTO        `json:"position"`       // The position of this component in the UI if applicable.
	Identity       string               `json:"identity"`       // The identity of the tenant.
	Users          []TenantEntity       `json:"users"`          // The users that belong to the user group.
	AccessPolicies []AccessPolicyEntity `json:"accessPolicies"` // The access policies this user group belongs to. *Read Only*
}
