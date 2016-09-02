package model
type ConnectionDTO struct {
Id string `json:"id"` // The id of the component.
ParentGroupId string `json:"parentGroupId"` // The id of parent process group of this component if applicable.
Position []PositionDTO `json:"position"` // The position of this component in the UI if applicable.
Source []ConnectableDTO `json:"source"` // The source of the connection.
Destination []ConnectableDTO `json:"destination"` // The destination of the connection.
Name string `json:"name"` // The name of the connection.
LabelIndex int `json:"labelIndex"` // The index of the bend point where to place the connection label.
GetzIndex int `json:"getzIndex"` // The z index of the connection.
SelectedRelationships [] `json:"selectedRelationships"` // The selected relationship that comprise the connection.
AvailableRelationships [] `json:"availableRelationships"` // The relationships that the source of the connection currently supports.
BackPressureObjectThreshold int `json:"backPressureObjectThreshold"` // The object count threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
BackPressureDataSizeThreshold string `json:"backPressureDataSizeThreshold"` // The object data size threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
FlowFileExpiration string `json:"flowFileExpiration"` // The amount of time a flow file may be in the flow before it will be automatically aged out of the flow. Once a flow file reaches this age it will be terminated from the flow the next time a processor attempts to start work on it.
Prioritizers [] `json:"prioritizers"` // The comparators used to prioritize the queue.
Bends []PositionDTO `json:"bends"` // The bend points on the connection.
}
