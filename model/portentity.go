package model

// attribute: portEntity
type PortEntity struct {
	Revision    RevisionDTO      `json:"revision"`    // The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Id          string           `json:"id"`          // The id of the component.
	Uri         string           `json:"uri"`         // The URI for futures requests to the component.
	Position    PositionDTO      `json:"position"`    // The position of this component in the UI if applicable.
	Permissions PermissionsDTO   `json:"permissions"` // The permissions for this component.
	Bulletins   []BulletinEntity `json:"bulletins"`   // The bulletins for this component.
	Component   PortDTO          `json:"component"`
	Status      PortStatusDTO    `json:"status"` // The status of the port.
	PortType    string           `json:"portType"`
}
