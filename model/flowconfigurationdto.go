package model

type FlowConfigurationDTO struct {
	SupportsConfigurableAuthorizer bool   `json:"supportsConfigurableAuthorizer"` // [Read Only] Whether this NiFi supports a configurable authorizer.
	AutoRefreshIntervalSeconds     int64  `json:"autoRefreshIntervalSeconds"`     // [Read Only] The interval in seconds between the automatic NiFi refresh requests.
	CurrentTime                    string `json:"currentTime"`                    // The current time on the system.
	TimeOffset                     int32  `json:"timeOffset"`                     // The time offset of the system.
}
