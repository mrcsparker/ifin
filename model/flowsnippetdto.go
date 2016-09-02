package model

type FlowSnippetDTO struct {
	ProcessGroups       []ProcessGroupDTO       `json:"processGroups"`       // The process groups in this flow snippet.
	RemoteProcessGroups []RemoteProcessGroupDTO `json:"remoteProcessGroups"` // The remote process groups in this flow snippet.
	Processors          []ProcessorDTO          `json:"processors"`          // The processors in this flow snippet.
	InputPorts          []PortDTO               `json:"inputPorts"`          // The input ports in this flow snippet.
	OutputPorts         []PortDTO               `json:"outputPorts"`         // The output ports in this flow snippet.
	Connections         []ConnectionDTO         `json:"connections"`         // The connections in this flow snippet.
	Labels              []LabelDTO              `json:"labels"`              // The labels in this flow snippet.
	Funnels             []FunnelDTO             `json:"funnels"`             // The funnels in this flow snippet.
	ControllerServices  []ControllerServiceDTO  `json:"controllerServices"`  // The controller services in this flow snippet.
}
