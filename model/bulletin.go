package model

type Bulletin struct {
	Id          int    `json:"id"`          // The id of the bulletin.
	NodeAddress string `json:"nodeAddress"` // If clustered, the address of the node from whicih the bulletin originated.
	Category    string `json:"category"`    // The catagory of this bulletin.
	GroupId     string `json:"groupId"`     // The group id of the source component.
	SourceId    string `json:"sourceId"`    // The id of the source component.
	SourceName  string `json:"sourceName"`  // The name of the source component.
	Level       string `json:"level"`       // The level of the bulletin.
	Message     string `json:"message"`     // The bulletin message.
	Timestamp   string `json:"timestamp"`   // When this bulletin was generated.
}
