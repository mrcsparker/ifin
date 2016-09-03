package model

type ControllerDTO struct {
	Id                          string    `json:"id"`                          // The id of the NiFi.
	Name                        string    `json:"name"`                        // The name of the NiFi.
	Comments                    string    `json:"comments"`                    // The comments for the NiFi.
	RunningCount                int32     `json:"runningCount"`                // The number of running components in the NiFi.
	StoppedCount                int32     `json:"stoppedCount"`                // The number of stopped components in the NiFi.
	InvalidCount                int32     `json:"invalidCount"`                // The number of invalid components in the NiFi.
	DisabledCount               int32     `json:"disabledCount"`               // The number of disabled components in the NiFi.
	ActiveRemotePortCount       int32     `json:"activeRemotePortCount"`       // The number of active remote ports contained in the NiFi.
	InactiveRemotePortCount     int32     `json:"inactiveRemotePortCount"`     // The number of inactive remote porst contained in the NiFi.
	InputPortCount              int32     `json:"inputPortCount"`              // The number of input ports contained in the NiFi.
	OutputPortCount             int32     `json:"outputPortCount"`             // The number of output ports in the NiFi.
	RemoteSiteListeningPort     int32     `json:"remoteSiteListeningPort"`     // The Socket Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null.
	RemoteSiteHttpListeningPort int32     `json:"remoteSiteHttpListeningPort"` // The HTTP(S) Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null.
	SiteToSiteSecure            bool      `json:"siteToSiteSecure"`            // Indicates whether or not Site-to-Site communications with this instance is secure (2-way authentication).
	InstanceId                  string    `json:"instanceId"`                  // If clustered, the id of the Cluster Manager, otherwise the id of the NiFi.
	InputPorts                  []PortDTO `json:"inputPorts"`                  // The input ports available to send data to for the NiFi.
	OutputPorts                 []PortDTO `json:"outputPorts"`                 // The output ports available to received data from the NiFi.
}
