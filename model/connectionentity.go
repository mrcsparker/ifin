package model

type ConnectionEntity struct {
	Revision           []RevisionDTO         `json:"revision"`           // The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Id                 string                `json:"id"`                 // The id of the component.
	Uri                string                `json:"uri"`                // The URI for futures requests to the component.
	Position           []PositionDTO         `json:"position"`           // The position of this component in the UI if applicable.
	Permissions        []PermissionsDTO      `json:"permissions"`        // The permissions for this component.
	Bulletins          []BulletinEntity      `json:"bulletins"`          // The bulletins for this component.
	Component          []ConnectionDTO       `json:"component"`          //
	Status             []ConnectionStatusDTO `json:"status"`             // The status of the connection.
	Bends              []PositionDTO         `json:"bends"`              // The bend points on the connection.
	LabelIndex         int                   `json:"labelIndex"`         // The index of the bend point where to place the connection label.
	GetzIndex          int                   `json:"getzIndex"`          // The z index of the connection.
	SourceId           string                `json:"sourceId"`           // The identifier of the source of this connection.
	SourceGroupId      string                `json:"sourceGroupId"`      // The identifier of the group of the source of this connection.
	SourceType         string                `json:"sourceType"`         // The type of component the source connectable is.
	DestinationId      string                `json:"destinationId"`      // The identifier of the destination of this connection.
	DestinationGroupId string                `json:"destinationGroupId"` // The identifier of the group of the destination of this connection.
	DestinationType    string                `json:"destinationType"`    // The type of component the destination connectable is.
}
