package model

type RemoteProcessGroupStatusSnapshotEntity struct {
	Id                               string                                `json:"id"`                               // The id of the remote processo group.
	RemoteProcessGroupStatusSnapshot []RemoteProcessGroupStatusSnapshotDTO `json:"remoteProcessGroupStatusSnapshot"` //
	CanRead                          bool                                  `json:"canRead"`                          // Indicates whether the user can read a given resource.
}
