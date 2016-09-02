package model

type ControllerServicesEntity struct {
	CurrentTime        string                    `json:"currentTime"` // The current time on the system.
	ControllerServices []ControllerServiceEntity `json:"controllerServices"`
}
