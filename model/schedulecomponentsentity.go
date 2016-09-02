package model

type ScheduleComponentsEntity struct {
	Id    string `json:"id"`    // The id of the ProcessGroup
	State string `json:"state"` // The desired state of the descendant components

}
