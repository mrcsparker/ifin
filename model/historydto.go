package model

type HistoryDTO struct {
	Total         int            `json:"total"`         // The number of number of actions that matched the search criteria..
	LastRefreshed string         `json:"lastRefreshed"` // The timestamp when the report was generated.
	Actions       []ActionEntity `json:"actions"`       // The actions.
}
