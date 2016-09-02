package model

type PortStatusSnapshotEntity struct {
	Id                 string                  `json:"id"`                 // The id of the port.
	PortStatusSnapshot []PortStatusSnapshotDTO `json:"portStatusSnapshot"` //
	CanRead            bool                    `json:"canRead"`            // Indicates whether the user can read a given resource.
}
