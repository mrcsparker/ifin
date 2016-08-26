package model

type RemoteProcessGroup struct {
	Id                            string                     `json:"id"`
	Uri                           string                     `json:"uri"`
	Postition                     Position                   `json:"position"`
	ParentGroupId                 string                     `json:"parentGroupId"`
	TargetUri                     string                     `json:"targetUri"`
	TargetSecure                  bool                       `json:"targetSecure"`
	Name                          string                     `json:"name"`
	Comments                      string                     `json:"comments"`
	CommunicationsTimeout         string                     `json:"communicationsTimeout"`
	YieldDuration                 string                     `json:"yieldDuration"`
	AuthorizationIssues           []string                   `json:"authorizationIssues"`
	Transmitting                  bool                       `json:"transmitting"`
	InputPortCount                int                        `json:"inputPortCount"`
	OutputPortCount               int                        `json:"outputPortCount"`
	ActiveRemoteInputPortCount    int                        `json:"activeRemoteInputPortCount"`
	InactiveRemoteInputPortCount  int                        `json:"inactiveRemoteInputPortCount"`
	ActiveRemoteOutputPortCount   int                        `json:"activeRemoteOutputPortCount"`
	InactiveRemoteOutputPortCount int                        `json:"inactiveRemoteOutputPortCount"`
	FlowRefreshed                 string                     `json:"flowRefreshed"`
	Contents                      RemoteProcessGroupContents `json:"contents"`
}
