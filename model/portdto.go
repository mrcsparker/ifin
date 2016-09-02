package model

type PortDTO struct {
	Id                               string        `json:"id"`                               // The id of the component.
	ParentGroupId                    string        `json:"parentGroupId"`                    // The id of parent process group of this component if applicable.
	Position                         []PositionDTO `json:"position"`                         // The position of this component in the UI if applicable.
	Name                             string        `json:"name"`                             // The name of the port.
	Comments                         string        `json:"comments"`                         // The comments for the port.
	State                            string        `json:"state"`                            // The state of the port.
	Type                             string        `json:"type"`                             // The type of port.
	Transmitting                     bool          `json:"transmitting"`                     // Whether the port has incoming or output connections to a remote NiFi. This is only applicable when the port is running in the root group.
	ConcurrentlySchedulableTaskCount int           `json:"concurrentlySchedulableTaskCount"` // The number of tasks that should be concurrently scheduled for the port.
	UserAccessControl                []string      `json:"userAccessControl"`                // The users that are allowed to access the port.
	GroupAccessControl               []string      `json:"groupAccessControl"`               // The user groups that are allowed to access the port.
	ValidationErrors                 []string      `json:"validationErrors"`                 // Gets the validation errors from this port. These validation errors represent the problems with the port that must be resolved before it can be started.
}
