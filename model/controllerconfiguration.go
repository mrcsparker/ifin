package model

type ControllerConfiguration struct {
	Name                       string `json:"name"`                       // The name of this NiFi.
	Comments                   string `json:"comments"`                   // The comments for this NiFi.
	MaxTimerDrivenThreadCount  int32  `json:"maxTimerDrivenThreadCount"`  // The maimum number of timer driven threads the NiFi has available.
	MaxEventDrivenThreadCount  int32  `json:"maxEventDrivenThreadCount"`  // The maximum number of event driven threads the NiFi has avaiable.
	AutoRefreshIntervalSeconds int64  `json:"autoRefreshIntervalSeconds"` // The interval in seconds between the automatic NiFi refresh requests.
	SiteToSiteSecure           bool   `json:"siteToSiteSecure"`           // Indicates whether site to site communication with the NiFi is secure (requires 2-way authenticiation).
	CurrentTime                string `json:"currentTime"`                // The current time on the system.
	TimeOffset                 int32  `json:"timeOffset"`                 // The time offset of the system.
	ContentViewerUrl           string `json:"contentViewerUrl"`           // The URL for the content viewer if configured.
	URI                        string `json:"uri"`                        // The URI for the NiFi.
}
