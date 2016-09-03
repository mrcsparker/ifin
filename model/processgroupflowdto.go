package model

type ProcessGroupFlowDTO struct {
	Id            string               `json:"id"`            // The id of the component.
	Uri           string               `json:"uri"`           // The URI for futures requests to the component.
	ParentGroupId string               `json:"parentGroupId"` // The id of parent process group of this component if applicable.
	Breadcrumb    FlowBreadcrumbEntity `json:"breadcrumb"`    // The breadcrumb of the process group.
	Flow          FlowDTO              `json:"flow"`          // The flow structure starting at this Process Group.
	LastRefreshed string               `json:"lastRefreshed"` // The time the flow for the process group was last refreshed.
}
