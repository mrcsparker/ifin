package model

type StatusHistoryDTO struct {
	Generated string `json:"generated"` // When the status history was generated.

	FieldDescriptors   []StatusDescriptorDTO    `json:"fieldDescriptors"`   // The Descriptors that provide information on each of the metrics provided in the status history
	AggregateSnapshots []StatusSnapshotDTO      `json:"aggregateSnapshots"` // A list of StatusSnapshotDTO objects that provide the actual metric values for the component. If the NiFi instance is clustered, this will represent the aggregate status across all nodes. If the NiFi instance is not clustered, this will represent the status of the entire NiFi instance.
	NodeSnapshots      []NodeStatusSnapshotsDTO `json:"nodeSnapshots"`      // The NodeStatusSnapshotsDTO objects that provide the actual metric values for the component, for each node. If the NiFi instance is not clustered, this value will be null.
}
