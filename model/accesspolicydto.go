package model

type AccessPolicyDTO struct {
	Id            string         `json:"id"`            // The id of the component.
	ParentGroupId string         `json:"parentGroupId"` // The id of parent process group of this component if applicable.
	Position      []PositionDTO  `json:"position"`      // The position of this component in the UI if applicable.
	Resource      string         `json:"resource"`      // The resource for this access policy.
	Action        string         `json:"action"`        // The action associated with this access policy.
	Users         []TenantEntity `json:"users"`         // The set of user IDs associated with this access policy.
	UserGroups    []TenantEntity `json:"userGroups"`    // The set of user group IDs associated with this access policy.
}
