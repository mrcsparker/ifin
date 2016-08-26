package model

type Config struct {
	AutoRefreshIntervalSeconds int    `json:"autoRefreshIntervalSeconds"`
	Comments                   string `json:"comments"`
	CurrentTime                string `json:"currentTime"`
	MaxEventDrivenThreadCount  int    `json:"maxEventDrivenThreadCount"`
	MaxTimerDrivenThreadCount  int    `json:"maxTimerDrivenThreadCount"`
	Name                       string `json:"name"`
	OutputPortCount            int    `json:"outputPortCount"`
	RunningCount               int    `json:"runningCount"`
	SiteToSiteSecure           bool   `json:"siteToSiteSecure"`
	StoppedCount               int    `json:"stoppedCount"`
}
