package model

type StatusHistory struct {
	Generated        string             `json:"generated"`        // The timestamp when the status history was generated.
	Details          interface{}        `json:"details"`          // The component details for the status history.
	FieldDescriptors []StatusDescriptor `json:"fieldDescriptors"` // The descriptor for each support status field.
	StatusSnapshots  []StatusSnapshot   `json:"statusSnapshots"`  // The status snapshots.
}
