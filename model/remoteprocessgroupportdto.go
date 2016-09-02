package model

type RemoteProcessGroupPortDTO struct {
	Id                               string `json:"id"`                               // The id of the target port.
	GroupId                          string `json:"groupId"`                          // The id of the remote process group that the port resides in.
	Name                             string `json:"name"`                             // The name of the target port.
	Comments                         string `json:"comments"`                         // The comments as configured on the target port.
	ConcurrentlySchedulableTaskCount int    `json:"concurrentlySchedulableTaskCount"` // The number of task that may transmit flowfiles to the target port concurrently.
	Transmitting                     bool   `json:"transmitting"`                     // Whether the remote port is configured for transmission.
	UseCompression                   bool   `json:"useCompression"`                   // Whether the flowfiles are compressed when sent to the target port.
	Exists                           bool   `json:"exists"`                           // Whether the target port exists.
	TargetRunning                    bool   `json:"targetRunning"`                    // Whether the target port is running.
	Connected                        bool   `json:"connected"`                        // Whether the port has either an incoming or outgoing connection.
}
