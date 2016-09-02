package model

type FlowDTO struct {
	ProcessGroups       []ProcessGroupEntity       `json:"processGroups"`       // The process groups in this flow.
	RemoteProcessGroups []RemoteProcessGroupEntity `json:"remoteProcessGroups"` // The remote process groups in this flow.
	Processors          []ProcessorEntity          `json:"processors"`          // The processors in this flow.
	InputPorts          []PortEntity               `json:"inputPorts"`          // The input ports in this flow.
	OutputPorts         []PortEntity               `json:"outputPorts"`         // The output ports in this flow.
	Connections         []ConnectionEntity         `json:"connections"`         // The connections in this flow.
	Labels              []LabelEntity              `json:"labels"`              // The labels in this flow.
	Funnels             []FunnelEntity             `json:"funnels"`             // The funnels in this flow.
}
