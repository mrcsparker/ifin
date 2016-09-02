package model

type ProcessorDTO struct {
	Id                         string               `json:"id"`                         // The id of the component.
	ParentGroupId              string               `json:"parentGroupId"`              // The id of parent process group of this component if applicable.
	Position                   []PositionDTO        `json:"position"`                   // The position of this component in the UI if applicable.
	Name                       string               `json:"name"`                       // The name of the processor.
	Type                       string               `json:"type"`                       // The type of the processor.
	State                      string               `json:"state"`                      // The state of the processor
	Style                      string               `json:"style"`                      // Styles for the processor (background-color => #eee).
	Relationships              []RelationshipDTO    `json:"relationships"`              // The available relationships that the processor currently supports. *Read Only*
	Description                string               `json:"description"`                // The description of the processor.
	SupportsParallelProcessing bool                 `json:"supportsParallelProcessing"` // Whether the processor supports parallel processing.
	SupportsEventDriven        bool                 `json:"supportsEventDriven"`        // Whether the processor supports event driven scheduling.
	SupportsBatching           bool                 `json:"supportsBatching"`           // Whether the processor supports batching. This makes the run duration settings available.
	PersistsState              bool                 `json:"persistsState"`              // Whether the processor persists state.
	InputRequirement           string               `json:"inputRequirement"`           // The input requirement for this processor.
	Config                     []ProcessorConfigDTO `json:"config"`                     // The configuration details for the processor. These details will be included in a resopnse if the verbose flag is included in a request.
	ValidationErrors           []string             `json:"validationErrors"`           // The validation errors for the processor. These validation errors represent the problems with the processor that must be resolved before it can be started.
}
