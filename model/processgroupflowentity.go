package model

// attribute: processGroupFlowEntity
type ProcessGroupFlowEntity struct {
	Permissions      PermissionsDTO      `json:"permissions"` // The access policy for this process group.
	ProcessGroupFlow ProcessGroupFlowDTO `json:"processGroupFlow"`
}
