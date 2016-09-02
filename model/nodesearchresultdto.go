package model

type NodeSearchResultDTO struct {
	Id      string `json:"id"`      // The id of the node that matched the search.
	Address string `json:"address"` // The address of the node that matched the search.
}
