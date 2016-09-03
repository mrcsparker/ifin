package model

// attribute: entity
type BulletinEntity struct {
	Id          int64       `json:"id"`
	GroupId     string      `json:"groupId"`
	SourceId    string      `json:"sourceId"`
	Timestamp   string      `json:"timestamp"`
	NodeAddress string      `json:"nodeAddress"`
	CanRead     bool        `json:"canRead"` // [Read Only] Indicates whether the user can read a given resource.
	Bulletin    BulletinDTO `json:"bulletin"`
}
