package model

type ControllerConfiguration struct {
	Name                       string `json:"name"`
	Comments                   string `json:"comments"`
	MaxTimerDrivenThreadCount  int    `json:"maxTimerDrivenThreadCount"`
	MaxEventDrivenThreadCount  int    `json:"maxEventDrivenThreadCount"`
	AutoRefreshIntervalSeconds int    `json:"autoRefreshIntervalSeconds"`
	SiteToSiteSecure           bool   `json:"siteToSiteSecure"`
	CurrentTime                string `json:"currentTime"`
	TimeOffset                 int    `json:"timeOffset"`
	ContentViewerUrl           string `json:"contentViewerUrl"`
	URI                        string `json:"uri"`
}
