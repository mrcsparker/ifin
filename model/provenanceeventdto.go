package model

type ProvenanceEventDTO struct {
	Id                              string         `json:"id"`                              // The event uuid.
	EventId                         int            `json:"eventId"`                         // The event id. This is a one up number thats unique per node.
	EventTime                       string         `json:"eventTime"`                       // The timestamp of the event.
	EventDuration                   int            `json:"eventDuration"`                   // The event duration in milliseconds.
	LineageDuration                 int            `json:"lineageDuration"`                 // The duration since the lineage began, in milliseconds.
	EventType                       string         `json:"eventType"`                       // The type of the event.
	FlowFileUuid                    string         `json:"flowFileUuid"`                    // The uuid of the flowfile for the event.
	FileSize                        string         `json:"fileSize"`                        // The size of the flowfile for the event.
	FileSizeBytes                   int            `json:"fileSizeBytes"`                   // The size of the flowfile in bytes for the event.
	ClusterNodeId                   string         `json:"clusterNodeId"`                   // The identifier for the node where the event originated.
	ClusterNodeAddress              string         `json:"clusterNodeAddress"`              // The label for the node where the event originated.
	GroupId                         string         `json:"groupId"`                         // The id of the group that the component resides in. If the component is no longer in the flow, the group id will not be set.
	ComponentId                     string         `json:"componentId"`                     // The id of the component that generated the event.
	ComponentType                   string         `json:"componentType"`                   // The type of the component that generated the event.
	ComponentName                   string         `json:"componentName"`                   // The name of the component that generated the event.
	SourceSystemFlowFileId          string         `json:"sourceSystemFlowFileId"`          // The source system flowfile id.
	AlternateIdentifierUri          string         `json:"alternateIdentifierUri"`          // The alternate identifier uri for the fileflow for the event.
	Attributes                      []AttributeDTO `json:"attributes"`                      // The attributes of the flowfile for the event.
	ParentUuids                     []string       `json:"parentUuids"`                     // The parent uuids for the event.
	ChildUuids                      []string       `json:"childUuids"`                      // The child uuids for the event.
	TransitUri                      string         `json:"transitUri"`                      // The source/destination system uri if the event was a RECEIVE/SEND.
	Relationship                    string         `json:"relationship"`                    // The relationship to which the flowfile was routed if the event is of type ROUTE.
	Details                         string         `json:"details"`                         // The event details.
	ContentEqual                    bool           `json:"contentEqual"`                    // Whether the input and output content claim is the same.
	InputContentAvailable           bool           `json:"inputContentAvailable"`           // Whether the input content is still available.
	InputContentClaimSection        string         `json:"inputContentClaimSection"`        // The section in which the input content claim lives.
	InputContentClaimContainer      string         `json:"inputContentClaimContainer"`      // The container in which the input content claim lives.
	InputContentClaimIdentifier     string         `json:"inputContentClaimIdentifier"`     // The identifier of the input content claim.
	InputContentClaimOffset         int            `json:"inputContentClaimOffset"`         // The offset into the input content claim where the flowfiles content begins.
	InputContentClaimFileSize       string         `json:"inputContentClaimFileSize"`       // The file size of the input content claim formatted.
	InputContentClaimFileSizeBytes  int            `json:"inputContentClaimFileSizeBytes"`  // The file size of the intput content claim in bytes.
	OutputContentAvailable          bool           `json:"outputContentAvailable"`          // Whether the output content is still available.
	OutputContentClaimSection       string         `json:"outputContentClaimSection"`       // The section in which the output content claim lives.
	OutputContentClaimContainer     string         `json:"outputContentClaimContainer"`     // The container in which the output content claim lives.
	OutputContentClaimIdentifier    string         `json:"outputContentClaimIdentifier"`    // The identifier of the output content claim.
	OutputContentClaimOffset        int            `json:"outputContentClaimOffset"`        // The offset into the output content claim where the flowfiles content begins.
	OutputContentClaimFileSize      string         `json:"outputContentClaimFileSize"`      // The file size of the output content claim formatted.
	OutputContentClaimFileSizeBytes int            `json:"outputContentClaimFileSizeBytes"` // The file size of the output content claim in bytes.
	ReplayAvailable                 bool           `json:"replayAvailable"`                 // Whether or not replay is available.
	ReplayExplanation               string         `json:"replayExplanation"`               // Explanation as to why replay is unavailable.
	SourceConnectionIdentifier      string         `json:"sourceConnectionIdentifier"`      // The identifier of the queue/connection from which the flowfile was pulled to genereate this event. May be null if the queue/connection is unknown or the flowfile was generated from this event.
}
