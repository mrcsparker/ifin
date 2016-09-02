package model

type FlowBreadcrumbEntity struct {
	Id          string              `json:"id"`          // The id of this ancestor ProcessGroup.
	Permissions []PermissionsDTO    `json:"permissions"` // The permissions for this ancestor ProcessGroup.
	Breadcrumb  []FlowBreadcrumbDTO `json:"breadcrumb"`  // This breadcrumb.
}
