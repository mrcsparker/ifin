package model

type ReportingTaskDTO struct {
	Id                      string                `json:"id"`                      // The id of the component.
	ParentGroupId           string                `json:"parentGroupId"`           // The id of parent process group of this component if applicable.
	Position                PositionDTO           `json:"position"`                // The position of this component in the UI if applicable.
	Name                    string                `json:"name"`                    // The name of the reporting task.
	Type                    string                `json:"type"`                    // The fully qualified type of the reporting task.
	State                   string                `json:"state"`                   // The state of the reporting task.
	Comments                string                `json:"comments"`                // The comments of the reporting task.
	PersistsState           bool                  `json:"persistsState"`           // Whether the reporting task persists state.
	SchedulingPeriod        string                `json:"schedulingPeriod"`        // The frequency with which to schedule the reporting task. The format of the value willd epend on the valud of the schedulingStrategy.
	SchedulingStrategy      string                `json:"schedulingStrategy"`      // The scheduling strategy that determines how the schedulingPeriod value should be interpreted.
	DefaultSchedulingPeriod string                `json:"defaultSchedulingPeriod"` // The default scheduling period for the different scheduling strategies.
	Properties              string                `json:"properties"`              // The properties of the reporting task.
	Descriptors             PropertyDescriptorDTO `json:"descriptors"`             // The descriptors for the reporting tasks properties.
	CustomUiUrl             string                `json:"customUiUrl"`             // The URL for the custom configuration UI for the reporting task.
	AnnotationData          string                `json:"annotationData"`          // The anntation data for the repoting task. This is how the custom UI relays configuration to the reporting task.
	ValidationErrors        []string              `json:"validationErrors"`        // Gets the validation errors from the reporting task. These validation errors represent the problems with the reporting task that must be resolved before it can be scheduled to run.
	ActiveThreadCount       int32                 `json:"activeThreadCount"`       // The number of active threads for the reporting task.
}
