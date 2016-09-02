package model

type ProcessGroupDTO struct {
	Id                      string        `json:"id"`                      // The id of the component.
	ParentGroupId           string        `json:"parentGroupId"`           // The id of parent process group of this component if applicable.
	Position                []PositionDTO `json:"position"`                // The position of this component in the UI if applicable.
	Name                    string        `json:"name"`                    // The name of the process group.
	Comments                string        `json:"comments"`                // The comments for the process group.
	RunningCount            int           `json:"runningCount"`            // The number of running componetns in this process group.
	StoppedCount            int           `json:"stoppedCount"`            // The number of stopped components in the process group.
	InvalidCount            int           `json:"invalidCount"`            // The number of invalid components in the process group.
	DisabledCount           int           `json:"disabledCount"`           // The number of disabled components in the process group.
	ActiveRemotePortCount   int           `json:"activeRemotePortCount"`   // The number of active remote ports in the process group.
	InactiveRemotePortCount int           `json:"inactiveRemotePortCount"` // The number of inactive remote ports in the process group.
	InputPortCount          int           `json:"inputPortCount"`          // The number of input ports in the process group.
	OutputPortCount         int           `json:"outputPortCount"`         // The number of output ports in the process group.
}
