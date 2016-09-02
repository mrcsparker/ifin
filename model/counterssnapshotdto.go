package model

type CountersSnapshotDTO struct {
	Generated string       `json:"generated"` // The timestamp when the report was generated.
	Counters  []CounterDTO `json:"counters"`  // All counters in the NiFi.
}
