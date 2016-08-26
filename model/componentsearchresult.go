package model

type ComponentSearchResult struct {
	Id      string   `json:"id"`
	GroupId string   `json:"id"`
	Name    string   `json:"name"`
	Matches []string `json:"matches"`
}
