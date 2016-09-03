package model

type FunnelDTO struct {
	Id            string      `json:"id"`            // The id of the component.
	ParentGroupId string      `json:"parentGroupId"` // The id of parent process group of this component if applicable.
	Position      PositionDTO `json:"position"`      // The position of this component in the UI if applicable.
}
