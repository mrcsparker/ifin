package model

type CountersDTO struct {
	AggregateSnapshot CountersSnapshotDTO       `json:"aggregateSnapshot"` // A Counters snapshot that represents the aggregate values of all nodes in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this represents the stats of the single instance.
	NodeSnapshots     []NodeCountersSnapshotDTO `json:"nodeSnapshots"`     // A Counters snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
}
