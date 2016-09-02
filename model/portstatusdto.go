package model

type PortStatusDTO struct {
	Id                 string                      `json:"id"`                 // The id of the port.
	GroupId            string                      `json:"groupId"`            // The id of the parent process group of the port.
	Name               string                      `json:"name"`               // The name of the port.
	Transmitting       bool                        `json:"transmitting"`       // Whether the port has incoming or outgoing connections to a remote NiFi.
	RunStatus          string                      `json:"runStatus"`          // The run status of the port.
	StatsLastRefreshed string                      `json:"statsLastRefreshed"` // The time the status for the process group was last refreshed.
	AggregateSnapshot  []PortStatusSnapshotDTO     `json:"aggregateSnapshot"`  // A status snapshot that represents the aggregate stats of all nodes in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this represents the stats of the single instance.
	NodeSnapshots      []NodePortStatusSnapshotDTO `json:"nodeSnapshots"`      // A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
}
