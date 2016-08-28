package model

type Connection struct {
	Id                            string      `json:"id"`                            // The id of the component.
	Uri                           string      `json:"uri"`                           // The URI for futures requests to the component.
	Position                      Position    `json:"position"`                      // The position of this component in the UI if applicable.
	ParentGroupId                 string      `json:"parentGroupId"`                 // The id of parent process group of this component if applicable.
	Source                        Connectable `json:"source"`                        // The source of the connection.
	Destination                   Connectable `json:"destination"`                   // The destination of the connection.
	Name                          string      `json:"name"`                          // The name of the connection.
	LabelIndex                    int32       `json:"labelIndex"`                    // The index of the bend point where to place the connection label.
	GetzIndex                     int64       `json:"getzIndex"`                     // The z index of the connection.
	SelectedRelationships         []string    `json:"selectedRelationships"`         // The selected relationship that comprise the connection.
	AvailableRelationships        []string    `json:"availableRelationships"`        // The relationships that the source of the connection currently supports. This property is read only.
	BackPressureObjectThreshold   int64       `json:"backPressureObjectThreshold"`   // The object count threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
	BackPressureDataSizeThreshold string      `json:"backPressureDataSizeThreshold"` // The object data size threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
	FlowFileExpiration            string      `json:"flowFileExpiration"`            // The amount of time a flow file may be in the flow before it will be automatically aged out of the flow. Once a flow file reaches this age it will be terminated from the flow the next time a processor attempts to start work on it.
	Prioritizers                  []string    `json:"prioritizers"`                  // The comparators used to prioritize the queue.
	Bends                         []Position  `json:"bends"`                         // The bend points on the connection.
}
