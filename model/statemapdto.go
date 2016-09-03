package model

type StateMapDTO struct {
	Scope           string          `json:"scope"`           // The scope of this StateMap.
	TotalEntryCount int32           `json:"totalEntryCount"` // The total number of state entries. When the state map is lengthy, only of portion of the entries are returned.
	State           []StateEntryDTO `json:"state"`           // The state.
}
