package model

type ProcessGroupStatusDTO struct {
	Id                 string                              `json:"id"`                 // The ID of the Process Group
	Name               string                              `json:"name"`               // The name of the Process Group
	StatsLastRefreshed string                              `json:"statsLastRefreshed"` // The time the status for the process group was last refreshed.
	AggregateSnapshot  []ProcessGroupStatusSnapshotDTO     `json:"aggregateSnapshot"`  // The aggregate status of all nodes in the cluster
	NodeSnapshots      []NodeProcessGroupStatusSnapshotDTO `json:"nodeSnapshots"`      // The status reported by each node in the cluster. If the NiFi instance is a standalone instance, rather than a clustered instance, this value may be null.
}
