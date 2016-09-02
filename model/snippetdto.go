package model

type SnippetDTO struct {
	Id                  string      `json:"id"`                  // The id of the snippet.
	Uri                 string      `json:"uri"`                 // The URI of the snippet.
	ParentGroupId       string      `json:"parentGroupId"`       // The group id for the components in the snippet.
	ProcessGroups       RevisionDTO `json:"processGroups"`       // The ids of the process groups in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	RemoteProcessGroups RevisionDTO `json:"remoteProcessGroups"` // The ids of the remote process groups in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	Processors          RevisionDTO `json:"processors"`          // The ids of the processors in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	InputPorts          RevisionDTO `json:"inputPorts"`          // The ids of the input ports in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	OutputPorts         RevisionDTO `json:"outputPorts"`         // The ids of the output ports in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	Connections         RevisionDTO `json:"connections"`         // The ids of the connections in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	Labels              RevisionDTO `json:"labels"`              // The ids of the labels in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	Funnels             RevisionDTO `json:"funnels"`             // The ids of the funnels in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
}
