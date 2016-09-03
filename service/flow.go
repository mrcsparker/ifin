package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Flow struct {
}

/**
 * Retrieves details about this NiFi to put in the About dialog
 *
 *
 * Tags: ["flow"]
 *
 * @return model.AboutEntity
 */
func (self Flow) GetAboutInfo() model.AboutEntity {
	s := api.Setup()
	res := model.AboutEntity{}
	url := "http://localhost:8080/nifi-api/flow/about"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Retrieves the banners for this NiFi
 *
 *
 * Tags: ["flow"]
 *
 * @return model.BannerEntity
 */
func (self Flow) GetBanners() model.BannerEntity {
	s := api.Setup()
	res := model.BannerEntity{}
	url := "http://localhost:8080/nifi-api/flow/banners"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets current bulletins
 *
 *
 * Tags: ["flow"]
 *
 * @param after Includes bulletins with an id after this value.
 * @param sourceName Includes bulletins originating from this sources whose name match this regular expression.
 * @param message Includes bulletins whose message that match this regular expression.
 * @param sourceId Includes bulletins originating from this sources whose id match this regular expression.
 * @param groupId Includes bulletins originating from this sources whose group id match this regular expression.
 * @param limit The number of bulletins to limit the response to.
 * @return model.BulletinBoardEntity
 */
func (self Flow) GetBulletinBoard(after string, sourceName string, message string, sourceId string, groupId string, limit string) model.BulletinBoardEntity {
	s := api.Setup()
	res := model.BulletinBoardEntity{}
	url := "http://localhost:8080/nifi-api/flow/bulletin-board"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Generates a client id.
 *
 *
 * Tags: ["flow"]
 *
 * @return string
 */
func (self Flow) GenerateClientId() string {
	s := api.Setup()
	res := ""
	url := "http://localhost:8080/nifi-api/flow/client-id"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Searches the cluster for a node with the specified address
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["flow"]
 *
 * @param q Node address to search for.
 * @return model.ClusterSearchResultsEntity
 */
func (self Flow) SearchCluster(q string) model.ClusterSearchResultsEntity {
	s := api.Setup()
	res := model.ClusterSearchResultsEntity{}
	url := "http://localhost:8080/nifi-api/flow/cluster/search-results"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets the current status of this NiFi
 *
 *
 * Tags: ["flow"]
 *
 * @return model.ControllerStatusEntity
 */
func (self Flow) GetClusterSummary() model.ControllerStatusEntity {
	s := api.Setup()
	res := model.ControllerStatusEntity{}
	url := "http://localhost:8080/nifi-api/flow/cluster/summary"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Retrieves the configuration for this NiFi flow
 *
 *
 * Tags: ["flow"]
 *
 * @return model.FlowConfigurationEntity
 */
func (self Flow) GetFlowConfig() model.FlowConfigurationEntity {
	s := api.Setup()
	res := model.FlowConfigurationEntity{}
	url := "http://localhost:8080/nifi-api/flow/config"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets status for a connection
 *
 *
 * Tags: ["flow"]
 *
 * @param nodewise Whether or not to include the breakdown per node. Optional, defaults to false
 * @param clusterNodeId The id of the node where to get the status.
 * @param id The connection id.
 * @return model.ConnectionStatusEntity
 */
func (self Flow) GetConnectionStatus(nodewise bool, clusterNodeId string, id string) model.ConnectionStatusEntity {
	s := api.Setup()
	res := model.ConnectionStatusEntity{}
	url := "http://localhost:8080/nifi-api/flow/connections/{id}/status"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets the status history for a connection
 *
 *
 * Tags: ["flow"]
 *
 * @param id The connection id.
 * @return model.StatusHistoryEntity
 */
func (self Flow) GetConnectionStatusHistory(id string) model.StatusHistoryEntity {
	s := api.Setup()
	res := model.StatusHistoryEntity{}
	url := "http://localhost:8080/nifi-api/flow/connections/{id}/status/history"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Retrieves the types of controller services that this NiFi supports
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["flow"]
 *
 * @param serviceType If specified, will only return controller services of this type.
 * @return model.ControllerServiceTypesEntity
 */
func (self Flow) GetControllerServiceTypes(serviceType string) model.ControllerServiceTypesEntity {
	s := api.Setup()
	res := model.ControllerServiceTypesEntity{}
	url := "http://localhost:8080/nifi-api/flow/controller-service-types"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Retrieves Controller level bulletins
 *
 *
 * Tags: ["flow"]
 *
 * @return model.ControllerBulletinsEntity
 */
func (self Flow) GetBulletins() model.ControllerBulletinsEntity {
	s := api.Setup()
	res := model.ControllerBulletinsEntity{}
	url := "http://localhost:8080/nifi-api/flow/controller/bulletins"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all controller services
 *
 *
 * Tags: ["flow"]
 *
 * @return model.ControllerServicesEntity
 */
func (self Flow) GetControllerServicesFromController() model.ControllerServicesEntity {
	s := api.Setup()
	res := model.ControllerServicesEntity{}
	url := "http://localhost:8080/nifi-api/flow/controller/controller-services"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Retrieves the user identity of the user making the request
 *
 *
 * Tags: ["flow"]
 *
 * @return model.CurrentUserEntity
 */
func (self Flow) GetCurrentUser() model.CurrentUserEntity {
	s := api.Setup()
	res := model.CurrentUserEntity{}
	url := "http://localhost:8080/nifi-api/flow/current-user"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets configuration history
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["flow"]
 *
 * @param offset The offset into the result set.
 * @param count The number of actions to return.
 * @param sortColumn The field to sort on.
 * @param sortOrder The direction to sort.
 * @param startDate Include actions after this date.
 * @param endDate Include actions before this date.
 * @param userIdentity Include actions performed by this user.
 * @param sourceId Include actions on this component.
 * @return model.HistoryEntity
 */
func (self Flow) QueryHistory(offset string, count string, sortColumn string, sortOrder string, startDate string, endDate string, userIdentity string, sourceId string) model.HistoryEntity {
	s := api.Setup()
	res := model.HistoryEntity{}
	url := "http://localhost:8080/nifi-api/flow/history"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets configuration history for a component
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["flow"]
 *
 * @param componentId The component id.
 * @return model.ComponentHistoryEntity
 */
func (self Flow) GetComponentHistory(componentId string) model.ComponentHistoryEntity {
	s := api.Setup()
	res := model.ComponentHistoryEntity{}
	url := "http://localhost:8080/nifi-api/flow/history/components/{componentId}"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets an action
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["flow"]
 *
 * @param id The action id.
 * @return model.ActionEntity
 */
func (self Flow) GetAction(id string) model.ActionEntity {
	s := api.Setup()
	res := model.ActionEntity{}
	url := "http://localhost:8080/nifi-api/flow/history/{id}"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets status for an input port
 *
 *
 * Tags: ["flow"]
 *
 * @param nodewise Whether or not to include the breakdown per node. Optional, defaults to false
 * @param clusterNodeId The id of the node where to get the status.
 * @param id The input port id.
 * @return model.PortStatusEntity
 */
func (self Flow) GetInputPortStatus(nodewise bool, clusterNodeId string, id string) model.PortStatusEntity {
	s := api.Setup()
	res := model.PortStatusEntity{}
	url := "http://localhost:8080/nifi-api/flow/input-ports/{id}/status"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets status for an output port
 *
 *
 * Tags: ["flow"]
 *
 * @param nodewise Whether or not to include the breakdown per node. Optional, defaults to false
 * @param clusterNodeId The id of the node where to get the status.
 * @param id The output port id.
 * @return model.PortStatusEntity
 */
func (self Flow) GetOutputPortStatus(nodewise bool, clusterNodeId string, id string) model.PortStatusEntity {
	s := api.Setup()
	res := model.PortStatusEntity{}
	url := "http://localhost:8080/nifi-api/flow/output-ports/{id}/status"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Retrieves the types of prioritizers that this NiFi supports
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["flow"]
 *
 * @return model.PrioritizerTypesEntity
 */
func (self Flow) GetPrioritizers() model.PrioritizerTypesEntity {
	s := api.Setup()
	res := model.PrioritizerTypesEntity{}
	url := "http://localhost:8080/nifi-api/flow/prioritizers"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets a process group
 *
 *
 * Tags: ["flow"]
 *
 * @param id The process group id.
 * @return model.ProcessGroupFlowEntity
 */
func (self Flow) GetFlow(id string) model.ProcessGroupFlowEntity {
	s := api.Setup()
	res := model.ProcessGroupFlowEntity{}
	url := "http://localhost:8080/nifi-api/flow/process-groups/{id}"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Schedule or unschedule comopnents in the specified Process Group.
 *
 *
 * Tags: ["flow"]
 *
 * @param id The process group id.
 * @param body The request to schedule or unschedule. If the comopnents in the request are not specified, all authorized components will be considered.
 * @return model.ScheduleComponentsEntity
 */
func (self Flow) ScheduleComponents(id string, body model.ScheduleComponentsEntity) model.ScheduleComponentsEntity {
	s := api.Setup()
	res := model.ScheduleComponentsEntity{}
	url := "http://localhost:8080/nifi-api/flow/process-groups/{id}"
	resp, err := s.Put(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all controller services
 *
 *
 * Tags: ["flow"]
 *
 * @param id The process group id.
 * @return model.ControllerServicesEntity
 */
func (self Flow) GetControllerServicesFromGroup(id string) model.ControllerServicesEntity {
	s := api.Setup()
	res := model.ControllerServicesEntity{}
	url := "http://localhost:8080/nifi-api/flow/process-groups/{id}/controller-services"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets the status for a process group
 *
 * The status for a process group includes status for all descendent components. When invoked on the root group with recursive set to true, it will return the current status of every component in the flow.
 *
 * Tags: ["flow"]
 *
 * @param recursive Whether all descendant groups and the status of their content will be included. Optional, defaults to false
 * @param nodewise Whether or not to include the breakdown per node. Optional, defaults to false
 * @param clusterNodeId The id of the node where to get the status.
 * @param id The process group id.
 * @return model.ProcessGroupStatusEntity
 */
func (self Flow) GetProcessGroupStatus(recursive bool, nodewise bool, clusterNodeId string, id string) model.ProcessGroupStatusEntity {
	s := api.Setup()
	res := model.ProcessGroupStatusEntity{}
	url := "http://localhost:8080/nifi-api/flow/process-groups/{id}/status"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets status history for a remote process group
 *
 *
 * Tags: ["flow"]
 *
 * @param id The process group id.
 * @return model.StatusHistoryEntity
 */
func (self Flow) GetProcessGroupStatusHistory(id string) model.StatusHistoryEntity {
	s := api.Setup()
	res := model.StatusHistoryEntity{}
	url := "http://localhost:8080/nifi-api/flow/process-groups/{id}/status/history"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Retrieves the types of processors that this NiFi supports
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["flow"]
 *
 * @return model.ProcessorTypesEntity
 */
func (self Flow) GetProcessorTypes() model.ProcessorTypesEntity {
	s := api.Setup()
	res := model.ProcessorTypesEntity{}
	url := "http://localhost:8080/nifi-api/flow/processor-types"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets status for a processor
 *
 *
 * Tags: ["flow"]
 *
 * @param nodewise Whether or not to include the breakdown per node. Optional, defaults to false
 * @param clusterNodeId The id of the node where to get the status.
 * @param id The processor id.
 * @return model.ProcessorStatusEntity
 */
func (self Flow) GetProcessorStatus(nodewise bool, clusterNodeId string, id string) model.ProcessorStatusEntity {
	s := api.Setup()
	res := model.ProcessorStatusEntity{}
	url := "http://localhost:8080/nifi-api/flow/processors/{id}/status"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets status history for a processor
 *
 *
 * Tags: ["flow"]
 *
 * @param id The processor id.
 * @return model.StatusHistoryEntity
 */
func (self Flow) GetProcessorStatusHistory(id string) model.StatusHistoryEntity {
	s := api.Setup()
	res := model.StatusHistoryEntity{}
	url := "http://localhost:8080/nifi-api/flow/processors/{id}/status/history"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets status for a remote process group
 *
 *
 * Tags: ["flow"]
 *
 * @param nodewise Whether or not to include the breakdown per node. Optional, defaults to false
 * @param clusterNodeId The id of the node where to get the status.
 * @param id The remote process group id.
 * @return model.ProcessorStatusEntity
 */
func (self Flow) GetRemoteProcessGroupStatus(nodewise bool, clusterNodeId string, id string) model.ProcessorStatusEntity {
	s := api.Setup()
	res := model.ProcessorStatusEntity{}
	url := "http://localhost:8080/nifi-api/flow/remote-process-groups/{id}/status"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets the status history
 *
 *
 * Tags: ["flow"]
 *
 * @param id The remote process group id.
 * @return model.StatusHistoryEntity
 */
func (self Flow) GetRemoteProcessGroupStatusHistory(id string) model.StatusHistoryEntity {
	s := api.Setup()
	res := model.StatusHistoryEntity{}
	url := "http://localhost:8080/nifi-api/flow/remote-process-groups/{id}/status/history"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Retrieves the types of reporting tasks that this NiFi supports
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["flow"]
 *
 * @return model.ReportingTaskTypesEntity
 */
func (self Flow) GetReportingTaskTypes() model.ReportingTaskTypesEntity {
	s := api.Setup()
	res := model.ReportingTaskTypesEntity{}
	url := "http://localhost:8080/nifi-api/flow/reporting-task-types"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all reporting tasks
 *
 *
 * Tags: ["flow"]
 *
 * @return model.ReportingTasksEntity
 */
func (self Flow) GetReportingTasks() model.ReportingTasksEntity {
	s := api.Setup()
	res := model.ReportingTasksEntity{}
	url := "http://localhost:8080/nifi-api/flow/reporting-tasks"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Performs a search against this NiFi using the specified search term
 *
 * Only search results from authorized components will be returned.
 *
 * Tags: ["flow"]
 *
 * @param q
 * @return model.SearchResultsEntity
 */
func (self Flow) SearchFlow(q string) model.SearchResultsEntity {
	s := api.Setup()
	res := model.SearchResultsEntity{}
	url := "http://localhost:8080/nifi-api/flow/search-results"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets the current status of this NiFi
 *
 *
 * Tags: ["flow"]
 *
 * @return model.ControllerStatusEntity
 */
func (self Flow) GetControllerStatus() model.ControllerStatusEntity {
	s := api.Setup()
	res := model.ControllerStatusEntity{}
	url := "http://localhost:8080/nifi-api/flow/status"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all templates
 *
 *
 * Tags: ["flow"]
 *
 * @return model.TemplatesEntity
 */
func (self Flow) GetTemplates() model.TemplatesEntity {
	s := api.Setup()
	res := model.TemplatesEntity{}
	url := "http://localhost:8080/nifi-api/flow/templates"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
