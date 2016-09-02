package model

type RemoteProcessGroupContentsDTO struct {
	InputPorts  []RemoteProcessGroupPortDTO `json:"inputPorts"`  // The input ports to which data can be sent.
	OutputPorts []RemoteProcessGroupPortDTO `json:"outputPorts"` // The output ports from which data can be retrieved.
}
