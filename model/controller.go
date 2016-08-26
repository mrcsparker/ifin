package model

type Controller struct {
	Id                      string `json:"id"`
	Name                    string `json:"name"`
	Comments                string `json:"comments"`
	RunningCount            int    `json:"runningCount"`
	StoppedCount            int    `json:"stoppedCount"`
	InvalidCount            int    `json:"invalidCount"`
	DisabledCount           int    `json:"disabledCount"`
	ActiveRemotePortCount   int    `json:"activeRemotePortCount"`
	InactiveRemotePortCount int    `json:"inactiveRemotePortCount"`
	InputPortCount          int    `json:"inputPortCount"`
	OutputPortCount         int    `json:"outputPortCount"`
	RemoteSiteListeningPort int    `json:"remoteSiteListeningPort"`
	SiteToSiteSecure        bool   `json:"siteToSiteSecure"`
	InstanceId              string `json:"instanceId"`
	InputPorts              []Port `json:"inputPorts"`
	OutputPorts             []Port `json:"outputPorts"`
}
