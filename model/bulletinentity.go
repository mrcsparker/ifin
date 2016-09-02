package model

type BulletinEntity struct {
	Id          int           `json:"id"`          //
	GroupId     string        `json:"groupId"`     //
	SourceId    string        `json:"sourceId"`    //
	Timestamp   string        `json:"timestamp"`   //
	NodeAddress string        `json:"nodeAddress"` //
	CanRead     bool          `json:"canRead"`     // Indicates whether the user can read a given resource.
	Bulletin    []BulletinDTO `json:"bulletin"`    //
}
