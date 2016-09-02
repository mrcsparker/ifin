package model

type ConnectionStatusSnapshotEntity struct {
	Id                       string                        `json:"id"` // The id of the connection.
	ConnectionStatusSnapshot []ConnectionStatusSnapshotDTO `json:"connectionStatusSnapshot"`
	CanRead                  bool                          `json:"canRead"` // Indicates whether the user can read a given resource. *Read Only*
}
