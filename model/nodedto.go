package model
type NodeDTO struct {
NodeId string `json:"nodeId"` // The id of the node.
Address string `json:"address"` // The node's host/ip address.
ApiPort int `json:"apiPort"` // The port the node is listening for API requests.
Status string `json:"status"` // The node's status.
Heartbeat string `json:"heartbeat"` // the time of the nodes's last heartbeat.
ConnectionRequested string `json:"connectionRequested"` // The time of the node's last connection request.
Roles [] `json:"roles"` // The roles of this node.
ActiveThreadCount int `json:"activeThreadCount"` // The active threads for the NiFi on the node.
Queued string `json:"queued"` // The queue the NiFi on the node.
Events []NodeEventDTO `json:"events"` // The node's events.
NodeStartTime string `json:"nodeStartTime"` // The time at which this Node was last refreshed.
}
