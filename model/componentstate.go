package model

type ComponentState struct {
	ComponentId      string   `json:"componentId"`      // The component identifier.
	StateDescription string   `json:"stateDescription"` // Description of the state this component persists.
	ClusterState     StateMap `json:"clusterState"`     // The cluster state for this component, or null if this NiFi is a standalone instance.
	LocalState       StateMap `json:"localState"`       // The local state for this component.
}
