package model

type ConnectionStatusSnapshotDTO struct {
	Id              string `json:"id"`              // The id of the connection.
	GroupId         string `json:"groupId"`         // The id of the process group the connection belongs to.
	Name            string `json:"name"`            // The name of the connection.
	SourceId        string `json:"sourceId"`        // The id of the source of the connection.
	SourceName      string `json:"sourceName"`      // The name of the source of the connection.
	DestinationId   string `json:"destinationId"`   // The id of the destination of the connection.
	DestinationName string `json:"destinationName"` // The name of the destination of the connection.
	FlowFilesIn     int    `json:"flowFilesIn"`     // The number of FlowFiles that have come into the connection in the last 5 minutes.
	BytesIn         int    `json:"bytesIn"`         // The size of the FlowFiles that have come into the connection in the last 5 minutes.
	Input           string `json:"input"`           // The input count/size for the connection in the last 5 minutes, pretty printed.
	FlowFilesOut    int    `json:"flowFilesOut"`    // The number of FlowFiles that have left the connection in the last 5 minutes.
	BytesOut        int    `json:"bytesOut"`        // The number of bytes that have left the connection in the last 5 minutes.
	Output          string `json:"output"`          // The output count/sie for the connection in the last 5 minutes, pretty printed.
	FlowFilesQueued int    `json:"flowFilesQueued"` // The number of FlowFiles that are currently queued in the connection.
	BytesQueued     int    `json:"bytesQueued"`     // The size of the FlowFiles that are currently queued in the connection.
	Queued          string `json:"queued"`          // The total count and size of queued flowfiles formatted.
	QueuedSize      string `json:"queuedSize"`      // The total size of flowfiles that are queued formatted.
	QueuedCount     string `json:"queuedCount"`     // The number of flowfiles that are queued, pretty printed.
}
