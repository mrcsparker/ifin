package model

type RemoteProcessGroupContents struct {
	InputPorts  []RemoteProcessGroupPort `json:"inputPorts"`
	OutputPorts []RemoteProcessGroupPort `json:"outputPorts"`
}
