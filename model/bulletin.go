package model

type Bulletin struct {
	Id          int    `json:"id"`
	NodeAddress string `json:"nodeAddress"`
	Category    string `json:"category"`
	GroupId     string `json:"groupId"`
	SourceId    string `json:"sourceId"`
	SourceName  string `json:"sourceName"`
	Level       string `json:"level"`
	Message     string `json:"message"`
	Timestamp   string `json:"timestamp"`
}
