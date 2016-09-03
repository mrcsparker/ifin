package model

type RemoteProcessGroupStatusDTO struct {
	GroupId            string                                    `json:"groupId"`            // The unique ID of the process group that the Processor belongs to
	Id                 string                                    `json:"id"`                 // The unique ID of the Processor
	Name               string                                    `json:"name"`               // The name of the remote process group.
	TargetUri          string                                    `json:"targetUri"`          // The URI of the target system.
	TransmissionStatus string                                    `json:"transmissionStatus"` // The transmission status of the remote process group.
	StatsLastRefreshed string                                    `json:"statsLastRefreshed"` // The time the status for the process group was last refreshed.
	AggregateSnapshot  RemoteProcessGroupStatusSnapshotDTO       `json:"aggregateSnapshot"`  // A status snapshot that represents the aggregate stats of all nodes in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this represents the stats of the single instance.
	NodeSnapshots      []NodeRemoteProcessGroupStatusSnapshotDTO `json:"nodeSnapshots"`      // A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
}
