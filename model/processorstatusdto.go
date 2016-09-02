package model

type ProcessorStatusDTO struct {
	GroupId            string                           `json:"groupId"`            // The unique ID of the process group that the Processor belongs to
	Id                 string                           `json:"id"`                 // The unique ID of the Processor
	Name               string                           `json:"name"`               // The name of the Processor
	Type               string                           `json:"type"`               // The type of the Processor
	RunStatus          string                           `json:"runStatus"`          // The run status of the Processor
	StatsLastRefreshed string                           `json:"statsLastRefreshed"` // The timestamp of when the stats were last refreshed
	AggregateSnapshot  []ProcessorStatusSnapshotDTO     `json:"aggregateSnapshot"`  // A status snapshot that represents the aggregate stats of all nodes in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this represents the stats of the single instance.
	NodeSnapshots      []NodeProcessorStatusSnapshotDTO `json:"nodeSnapshots"`      // A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
}
