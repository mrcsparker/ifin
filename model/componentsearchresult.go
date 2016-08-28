package model

type ComponentSearchResult struct {
	Id      string   `json:"id"`      // The id of the component that matched the search.
	GroupId string   `json:"id"`      // The group id of the component that matched the search.
	Name    string   `json:"name"`    // The name of the component that matched the search.
	Matches []string `json:"matches"` // What matched the search from the component.
}
