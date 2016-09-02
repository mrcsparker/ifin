package model

type ActionEntity struct {
	Id        int         `json:"id"`        //
	Timestamp string      `json:"timestamp"` //
	SourceId  string      `json:"sourceId"`  //
	CanRead   bool        `json:"canRead"`   // Indicates whether the user can read a given resource.
	Action    []ActionDTO `json:"action"`    //
}
