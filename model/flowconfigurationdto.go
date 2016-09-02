package model

type FlowConfigurationDTO struct {
	SupportsConfigurableAuthorizer bool   `json:"supportsConfigurableAuthorizer"` // Whether this NiFi supports a configurable authorizer.
	AutoRefreshIntervalSeconds     int    `json:"autoRefreshIntervalSeconds"`     // The interval in seconds between the automatic NiFi refresh requests.
	CurrentTime                    string `json:"currentTime"`                    // The current time on the system.
	TimeOffset                     int    `json:"timeOffset"`                     // The time offset of the system.
}
