package model

type ProcessorConfig struct {
	Properties                       interface{} `json:"properties"`                       // The properties for the processor. Properties whose value is not set will only contain the property name.
	Descriptors                      interface{} `json:"descriptors"`                      // Descriptors for the processor's properties.
	SchedulingPeriod                 string      `json:"schedulingPeriod"`                 // The frequency with which to schedule the processor. The format of the value will depend on th value of schedulingStrategy.
	SchedulingStrategy               string      `json:"schedulingStrategy"`               // Indcates whether the prcessor should be scheduled to run in event or timer driven mode.
	PenaltyDuration                  string      `json:"penaltyDuration"`                  // The amout of time that is used when the process penalizes a flowfile.
	YieldDuration                    string      `json:"yieldDuration"`                    // The amount of time that must elapse before this processor is scheduled again after yielding.
	BulletinLevel                    string      `json:"bulletinLevel"`                    // The level at which the processor will report bulletins.
	RunDurationMillis                int         `json:"runDurationMillis"`                // The run duration for the processor in milliseconds.
	ConcurrentlySchedulableTaskCount int         `json:"concurrentlySchedulableTaskCount"` // The number of tasks that should be concurrently schedule for the processor. If the processor doesn't allow parallol processing then any positive input will be ignored.
	AutoTerminatedRelationships      []string    `json:"autoTerminatedRelationships"`      // The names of all relationships that cause a flow file to be terminated if the relationship is not connected elsewhere.
	Comments                         string      `json:"comments"`                         // The comments for the processor.
	CustomUiUrl                      string      `json:"customUiUrl"`                      // The URL for the processor's custom configuration UI if applicable.
	LossTolerant                     bool        `json:"lossTolerant"`                     // Whether the processor is loss tolerant.
	AnnotationData                   string      `json:"annotationData"`                   // The annotation data for the processor used to relay configuration between a custom UI and the procesosr.
	DefaultConcurrentTasks           interface{} `json:"defaultConcurrentTasks"`           // Maps default values for concurrent tasks for each applicable scheduling strategy.
	DefaultSchedulingPeriod          interface{} `json:"defaultSchedulingPeriod"`          // Maps default values for scheduling period for each applicable scheduling strategy.
}
