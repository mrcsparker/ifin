package model

type Config struct {
	AutoRefreshIntervalSeconds int    `json:"autoRefreshIntervalSeconds"`
	Comments                   string `json:"comments"`
	CurrentTime                string `json:"currentTime"`
	MaxEventDrivenThreadCount  int    `json:"maxEventDrivenThreadCount"`
	MaxTimerDrivenThreadCount  int    `json:"maxTimerDrivenThreadCount"`
	Name                       string `json:"name"`
}
