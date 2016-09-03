package model

type ControllerConfigurationDTO struct {
	MaxTimerDrivenThreadCount int32 `json:"maxTimerDrivenThreadCount"` // The maimum number of timer driven threads the NiFi has available.
	MaxEventDrivenThreadCount int32 `json:"maxEventDrivenThreadCount"` // The maximum number of event driven threads the NiFi has avaiable.
}
