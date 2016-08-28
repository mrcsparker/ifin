package model

type SearchResult struct {
	ProcessorResults          []ComponentSearchResult `json:"processorResults"`          // The processors that matched the search.
	connectionResults         []ComponentSearchResult `json:"connectionResults"`         // The connections that matched the search.
	ProcessGroupResults       []ComponentSearchResult `json:"processGroupResults"`       // The process groups that matched the search.
	InputPortResults          []ComponentSearchResult `json:"inputPortResults"`          // The input ports that matched the search.
	OutputPortResults         []ComponentSearchResult `json:"outputPortResults"`         // The output ports that matched the search.
	RemoteProcessGroupResults []ComponentSearchResult `json:"remoteProcessGroupResults"` // The remote process groups that matched the search.
	FunnelResults             []ComponentSearchResult `json:"funnelResults"`             // The funnels that matched the search.
}
