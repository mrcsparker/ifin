package model

type RemoteProcessGroupPort struct {
	Id                               string `json:"id"`
	GroupId                          string `json:"groupId"`
	Name                             string `json:"name"`
	Comments                         string `json:"comments"`
	ConcurrentlySchedulableTaskCount int    `json:"concurrentlySchedulableTaskCount"`
	Transmitting                     bool   `json:"transmitting"`
	UseCompression                   bool   `json:"useCompression"`
	Exists                           bool   `json:"exists"`
	TargetRunning                    bool   `json:"targetRunning"`
	Connected                        bool   `json:"connected"`
}
