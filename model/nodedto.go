package model

type NodeDTO struct {
	NodeId              string         `json:"nodeId"`              // [Read Only] The id of the node.
	Address             string         `json:"address"`             // [Read Only] The node's host/ip address.
	ApiPort             int32          `json:"apiPort"`             // [Read Only] The port the node is listening for API requests.
	Status              string         `json:"status"`              // The node's status.
	Heartbeat           string         `json:"heartbeat"`           // [Read Only] the time of the nodes's last heartbeat.
	ConnectionRequested string         `json:"connectionRequested"` // [Read Only] The time of the node's last connection request.
	Roles               []string       `json:"roles"`               // [Read Only] The roles of this node.
	ActiveThreadCount   int32          `json:"activeThreadCount"`   // [Read Only] The active threads for the NiFi on the node.
	Queued              string         `json:"queued"`              // [Read Only] The queue the NiFi on the node.
	Events              []NodeEventDTO `json:"events"`              // [Read Only] The node's events.
	NodeStartTime       string         `json:"nodeStartTime"`       // [Read Only] The time at which this Node was last refreshed.
}
