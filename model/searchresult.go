package model

type SearchResult struct {
	ProcessorResults          []ComponentSearchResult `json:"processorResults"`
	connectionResults         []ComponentSearchResult `json:"connectionResults"`
	ProcessGroupResults       []ComponentSearchResult `json:"processGroupResults"`
	InputPortResults          []ComponentSearchResult `json:"inputPortResults"`
	OutputPortResults         []ComponentSearchResult `json:"outputPortResults"`
	RemoteProcessGroupResults []ComponentSearchResult `json:"remoteProcessGroupResults"`
	FunnelResults             []ComponentSearchResult `json:"funnelResults"`
}
