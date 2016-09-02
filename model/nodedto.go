package model

type NodeDTO struct {
	NodeId              string         `json:"nodeId"`              // The id of the node. *Read Only*
	Address             string         `json:"address"`             // The node's host/ip address. *Read Only*
	ApiPort             int            `json:"apiPort"`             // The port the node is listening for API requests. *Read Only*
	Status              string         `json:"status"`              // The node's status.
	Heartbeat           string         `json:"heartbeat"`           // the time of the nodes's last heartbeat. *Read Only*
	ConnectionRequested string         `json:"connectionRequested"` // The time of the node's last connection request. *Read Only*
	Roles               []string       `json:"roles"`               // The roles of this node. *Read Only*
	ActiveThreadCount   int            `json:"activeThreadCount"`   // The active threads for the NiFi on the node. *Read Only*
	Queued              string         `json:"queued"`              // The queue the NiFi on the node. *Read Only*
	Events              []NodeEventDTO `json:"events"`              // The node's events. *Read Only*
	NodeStartTime       string         `json:"nodeStartTime"`       // The time at which this Node was last refreshed. *Read Only*
}
