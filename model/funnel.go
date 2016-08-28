package model

type Funnel struct {
	Id            string   `json:"id"`            // The id of the component.
	Uri           string   `json:"uri"`           // The URI for futures requests to the component.
	Position      Position `json:"position"`      // The position of this component in the UI if applicable.
	ParentGroupId string   `json:"parentGroupId"` // The id of parent process group of this component if applicable.
}
