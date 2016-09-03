package model

type SystemDiagnosticsDTO struct {
	AggregateSnapshot SystemDiagnosticsSnapshotDTO       `json:"aggregateSnapshot"` // A systems diagnostic snapshot that represents the aggregate values of all nodes in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this represents the stats of the single instance.
	NodeSnapshots     []NodeSystemDiagnosticsSnapshotDTO `json:"nodeSnapshots"`     // A systems diagnostics snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
}
