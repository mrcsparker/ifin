package model

// attribute: processGroupEntity
type ProcessGroupEntity struct {
	Revision                RevisionDTO           `json:"revision"`    // The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Id                      string                `json:"id"`          // The id of the component.
	Uri                     string                `json:"uri"`         // The URI for futures requests to the component.
	Position                PositionDTO           `json:"position"`    // The position of this component in the UI if applicable.
	Permissions             PermissionsDTO        `json:"permissions"` // The permissions for this component.
	Bulletins               []BulletinEntity      `json:"bulletins"`   // The bulletins for this component.
	Component               ProcessGroupDTO       `json:"component"`
	Status                  ProcessGroupStatusDTO `json:"status"`                  // The status of the process group.
	RunningCount            int32                 `json:"runningCount"`            // The number of running componetns in this process group.
	StoppedCount            int32                 `json:"stoppedCount"`            // The number of stopped components in the process group.
	InvalidCount            int32                 `json:"invalidCount"`            // The number of invalid components in the process group.
	DisabledCount           int32                 `json:"disabledCount"`           // The number of disabled components in the process group.
	ActiveRemotePortCount   int32                 `json:"activeRemotePortCount"`   // The number of active remote ports in the process group.
	InactiveRemotePortCount int32                 `json:"inactiveRemotePortCount"` // The number of inactive remote ports in the process group.
	InputPortCount          int32                 `json:"inputPortCount"`          // The number of input ports in the process group.
	OutputPortCount         int32                 `json:"outputPortCount"`         // The number of output ports in the process group.
}
