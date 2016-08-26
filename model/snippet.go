package model

type Snippet struct {
	Id                  string        `json:"id"`
	Uri                 string        `json:"uri"`
	ParentGroupId       string        `json:"parentGroupId"`
	Linked              bool          `json:"linked"`
	ProcessGroups       []string      `json:"processGroups"`
	RemoteProcessGroups []string      `json:"remoteProcessGroups"`
	Processors          []string      `json:"processors"`
	InputPorts          []string      `json:"inputPorts"`
	OutputPorts         []string      `json:"outputPorts"`
	Connections         []string      `json:"connections"`
	Labels              []string      `json:"labels"`
	Funnels             []string      `json:"funnels"`
	Contents            []FlowSnippet `json:"contents"`
}
