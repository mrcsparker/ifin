package model

type RemoteProcessGroupContents struct {
	InputPorts  []RemoteProcessGroupPort `json:"inputPorts"`  // The input ports to which data can be sent.
	OutputPorts []RemoteProcessGroupPort `json:"outputPorts"` // The output ports from which data can be retrieved.
}
