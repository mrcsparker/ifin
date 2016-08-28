package model

type StateMap struct {
	Scope           string       `json:"scope"`           // The scope of this StateMap.
	TotalEntryCount int          `json:"totalEntryCount"` // The total number of state entries. When the state map is lengthy, only of portion of the entries are returned.
	State           []StateEntry `json:"state"`           // The state.
}
