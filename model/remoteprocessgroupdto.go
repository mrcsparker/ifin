package model

type RemoteProcessGroupDTO struct {
	Id                            string                          `json:"id"`                    // The id of the component.
	ParentGroupId                 string                          `json:"parentGroupId"`         // The id of parent process group of this component if applicable.
	Position                      []PositionDTO                   `json:"position"`              // The position of this component in the UI if applicable.
	TargetUri                     string                          `json:"targetUri"`             // The target URI of the remote process group.
	TargetSecure                  bool                            `json:"targetSecure"`          // Whether the target is running securely.
	Name                          string                          `json:"name"`                  // The name of the remote process group.
	Comments                      string                          `json:"comments"`              // The comments for the remote process group.
	CommunicationsTimeout         string                          `json:"communicationsTimeout"` // The time period used for the timeout when commicating with the target.
	YieldDuration                 string                          `json:"yieldDuration"`         // When yielding, this amount of time must elapse before the remote process group is scheduled again.
	TransportProtocol             string                          `json:"transportProtocol"`
	ProxyHost                     string                          `json:"proxyHost"`
	ProxyPort                     int                             `json:"proxyPort"`
	ProxyUser                     string                          `json:"proxyUser"`
	ProxyPassword                 string                          `json:"proxyPassword"`
	AuthorizationIssues           []string                        `json:"authorizationIssues"`           // Any remote authorization issues for the remote process group.
	Transmitting                  bool                            `json:"transmitting"`                  // Whether the remote process group is actively transmitting.
	InputPortCount                int                             `json:"inputPortCount"`                // The number of remote input ports currently available on the target.
	OutputPortCount               int                             `json:"outputPortCount"`               // The number of remote output ports currently available on the target.
	ActiveRemoteInputPortCount    int                             `json:"activeRemoteInputPortCount"`    // The number of active remote input ports.
	InactiveRemoteInputPortCount  int                             `json:"inactiveRemoteInputPortCount"`  // The number of inactive remote input ports.
	ActiveRemoteOutputPortCount   int                             `json:"activeRemoteOutputPortCount"`   // The number of acitve remote output ports.
	InactiveRemoteOutputPortCount int                             `json:"inactiveRemoteOutputPortCount"` // The number of inactive remote output ports.
	FlowRefreshed                 string                          `json:"flowRefreshed"`                 // The timestamp when this remote process group was last refreshed.
	Contents                      []RemoteProcessGroupContentsDTO `json:"contents"`                      // The contents of the remote process group. Will contain available input/output ports.
}
