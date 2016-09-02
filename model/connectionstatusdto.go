package model

type ConnectionStatusDTO struct {
	Id                 string                            `json:"id"`                 // The ID of the connection
	GroupId            string                            `json:"groupId"`            // The ID of the Process Group that the connection belongs to
	Name               string                            `json:"name"`               // The name of the connection
	StatsLastRefreshed string                            `json:"statsLastRefreshed"` // The timestamp of when the stats were last refreshed
	SourceId           string                            `json:"sourceId"`           // The ID of the source component
	SourceName         string                            `json:"sourceName"`         // The name of the source component
	DestinationId      string                            `json:"destinationId"`      // The ID of the destination component
	DestinationName    string                            `json:"destinationName"`    // The name of the destination component
	AggregateSnapshot  []ConnectionStatusSnapshotDTO     `json:"aggregateSnapshot"`  // The status snapshot that represents the aggregate stats of the cluster
	NodeSnapshots      []NodeConnectionStatusSnapshotDTO `json:"nodeSnapshots"`      // A list of status snapshots for each node
}
