package model

type FlowSnippet struct {
	ProcessGroups       Set                  `json:"processGroups"`       // The process groups in this flow snippet.
	RemoteProcessGroups []RemoteProcessGroup `json:"remoteProcessGroups"` //The remote process groups in this flow snippet.
	Processors          []Processor          `json:"processors"`          // The processors in this flow snippet.
	InputPorts          []Port               `json:"inputPorts"`          // The input ports in this flow snippet.
	OutputPorts         []Port               `json:"outputPorts"`         // The output ports in this flow snippet.
	Connections         []Connection         `json:"connections"`         // The connections in this flow snippet.
	Labels              []Label              `json:"labels"`              // The labels in this flow snippet.
	Funnels             []Funnel             `json:"funnels"`             // The funnels in this flow snippet.
	ControllerServices  []ControllerService  `json:"controllerServices"`  // The controller services in this flow snippet.
}
