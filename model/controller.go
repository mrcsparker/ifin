package model

type Controller struct {
	Id                      string `json:"id"`                      // The id of the NiFi.
	Name                    string `json:"name"`                    // The name of the NiFi.
	Comments                string `json:"comments"`                // The comments for the NiFi.
	RunningCount            int    `json:"runningCount"`            // The number of running components in the NiFi.
	StoppedCount            int    `json:"stoppedCount"`            // The number of stopped components in the NiFi.
	InvalidCount            int    `json:"invalidCount"`            // The number of invalid components in the NiFi.
	DisabledCount           int    `json:"disabledCount"`           // The number of disabled components in the NiFi.
	ActiveRemotePortCount   int    `json:"activeRemotePortCount"`   // The number of active remote ports contained in the NiFi.
	InactiveRemotePortCount int    `json:"inactiveRemotePortCount"` // The number of inactive remote porst contained in the NiFi.
	InputPortCount          int    `json:"inputPortCount"`          // The number of input ports contained in the NiFi.
	OutputPortCount         int    `json:"outputPortCount"`         // The number of output ports in the NiFi.
	RemoteSiteListeningPort int    `json:"remoteSiteListeningPort"` // The Socket Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null.
	SiteToSiteSecure        bool   `json:"siteToSiteSecure"`        // Indicates whether or not Site-to-Site communications with this instance is secure (2-way authentication).
	InstanceId              string `json:"instanceId"`              // If clustered, the id of the Cluster Manager, otherwise the id of the NiFi.
	InputPorts              []Port `json:"inputPorts"`              // The input ports available to send data to for the NiFi.
	OutputPorts             []Port `json:"outputPorts"`             // The output ports available to received data from the NiFi.
}
