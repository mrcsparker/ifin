package model

type SearchResultsDTO struct {
	ProcessorResults          []ComponentSearchResultDTO `json:"processorResults"`          // The processors that matched the search.
	ConnectionResults         []ComponentSearchResultDTO `json:"connectionResults"`         // The connections that matched the search.
	ProcessGroupResults       []ComponentSearchResultDTO `json:"processGroupResults"`       // The process groups that matched the search.
	InputPortResults          []ComponentSearchResultDTO `json:"inputPortResults"`          // The input ports that matched the search.
	OutputPortResults         []ComponentSearchResultDTO `json:"outputPortResults"`         // The output ports that matched the search.
	RemoteProcessGroupResults []ComponentSearchResultDTO `json:"remoteProcessGroupResults"` // The remote process groups that matched the search.
	FunnelResults             []ComponentSearchResultDTO `json:"funnelResults"`             // The funnels that matched the search.
}
