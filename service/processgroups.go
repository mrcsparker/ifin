package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type ProcessGroups struct {
}

/**
 * Gets a process group
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @return model.ProcessGroupEntity
 */
func (self ProcessGroups) GetProcessGroup(id string) model.ProcessGroupEntity {
	s := api.Setup()
	res := model.ProcessGroupEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}"
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
 * Updates a process group
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The process group configuration details.
 * @return model.ProcessGroupEntity
 */
func (self ProcessGroups) UpdateProcessGroup(id string, body model.ProcessGroupEntity) model.ProcessGroupEntity {
	s := api.Setup()
	res := model.ProcessGroupEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}"
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
 * Deletes a process group
 *
 *
 * Tags: ["process-groups"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The process group id.
 * @return model.ProcessGroupEntity
 */
func (self ProcessGroups) RemoveProcessGroup(version string, clientId string, id string) model.ProcessGroupEntity {
	s := api.Setup()
	res := model.ProcessGroupEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all connections
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @return model.ConnectionsEntity
 */
func (self ProcessGroups) GetConnections(id string) model.ConnectionsEntity {
	s := api.Setup()
	res := model.ConnectionsEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/connections"
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
 * Creates a connection
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The connection configuration details.
 * @return model.ConnectionEntity
 */
func (self ProcessGroups) CreateConnection(id string, body model.ConnectionEntity) model.ConnectionEntity {
	s := api.Setup()
	res := model.ConnectionEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/connections"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Creates a new controller service
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The controller service configuration details.
 * @return model.ControllerServiceEntity
 */
func (self ProcessGroups) CreateControllerService(id string, body model.ControllerServiceEntity) model.ControllerServiceEntity {
	s := api.Setup()
	res := model.ControllerServiceEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/controller-services"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all funnels
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @return model.FunnelsEntity
 */
func (self ProcessGroups) GetFunnels(id string) model.FunnelsEntity {
	s := api.Setup()
	res := model.FunnelsEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/funnels"
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
 * Creates a funnel
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The funnel configuration details.
 * @return model.FunnelEntity
 */
func (self ProcessGroups) CreateFunnel(id string, body model.FunnelEntity) model.FunnelEntity {
	s := api.Setup()
	res := model.FunnelEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/funnels"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all input ports
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @return model.InputPortsEntity
 */
func (self ProcessGroups) GetInputPorts(id string) model.InputPortsEntity {
	s := api.Setup()
	res := model.InputPortsEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/input-ports"
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
 * Creates an input port
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The input port configuration details.
 * @return model.PortEntity
 */
func (self ProcessGroups) CreateInputPort(id string, body model.PortEntity) model.PortEntity {
	s := api.Setup()
	res := model.PortEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/input-ports"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all labels
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @return model.LabelsEntity
 */
func (self ProcessGroups) GetLabels(id string) model.LabelsEntity {
	s := api.Setup()
	res := model.LabelsEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/labels"
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
 * Creates a label
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The label configuration details.
 * @return model.LabelEntity
 */
func (self ProcessGroups) CreateLabel(id string, body model.LabelEntity) model.LabelEntity {
	s := api.Setup()
	res := model.LabelEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/labels"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all output ports
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @return model.OutputPortsEntity
 */
func (self ProcessGroups) GetOutputPorts(id string) model.OutputPortsEntity {
	s := api.Setup()
	res := model.OutputPortsEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/output-ports"
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
 * Creates an output port
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The output port configuration.
 * @return model.PortEntity
 */
func (self ProcessGroups) CreateOutputPort(id string, body model.PortEntity) model.PortEntity {
	s := api.Setup()
	res := model.PortEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/output-ports"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all process groups
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @return model.ProcessorsEntity
 */
func (self ProcessGroups) GetProcessGroups(id string) model.ProcessorsEntity {
	s := api.Setup()
	res := model.ProcessorsEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/process-groups"
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
 * Creates a process group
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The process group configuration details.
 * @return model.ProcessGroupEntity
 */
func (self ProcessGroups) CreateProcessGroup(id string, body model.ProcessGroupEntity) model.ProcessGroupEntity {
	s := api.Setup()
	res := model.ProcessGroupEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/process-groups"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all processors
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @return model.ProcessorsEntity
 */
func (self ProcessGroups) GetProcessors(id string) model.ProcessorsEntity {
	s := api.Setup()
	res := model.ProcessorsEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/processors"
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
 * Creates a new processor
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The processor configuration details.
 * @return model.ProcessorEntity
 */
func (self ProcessGroups) CreateProcessor(id string, body model.ProcessorEntity) model.ProcessorEntity {
	s := api.Setup()
	res := model.ProcessorEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/processors"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Gets all remote process groups
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @return model.RemoteProcessGroupsEntity
 */
func (self ProcessGroups) GetRemoteProcessGroups(id string) model.RemoteProcessGroupsEntity {
	s := api.Setup()
	res := model.RemoteProcessGroupsEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/remote-process-groups"
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
 * Creates a new process group
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The remote process group configuration details.
 * @return model.RemoteProcessGroupEntity
 */
func (self ProcessGroups) CreateRemoteProcessGroup(id string, body model.RemoteProcessGroupEntity) model.RemoteProcessGroupEntity {
	s := api.Setup()
	res := model.RemoteProcessGroupEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/remote-process-groups"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Copies a snippet
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The copy snippet request.
 * @return model.FlowSnippetEntity
 */
func (self ProcessGroups) CopySnippet(id string, body model.CopySnippetRequestEntity) model.FlowSnippetEntity {
	s := api.Setup()
	res := model.FlowSnippetEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/snippet-instance"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Instantiates a template
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The instantiate template request.
 * @return model.FlowEntity
 */
func (self ProcessGroups) InstantiateTemplate(id string, body model.InstantiateTemplateRequestEntity) model.FlowEntity {
	s := api.Setup()
	res := model.FlowEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/template-instance"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Creates a template
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @param body The create template request.
 * @return model.TemplateEntity
 */
func (self ProcessGroups) CreateTemplate(id string, body model.CreateTemplateRequestEntity) model.TemplateEntity {
	s := api.Setup()
	res := model.TemplateEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/templates"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Imports a template
 *
 *
 * Tags: ["process-groups"]
 *
 * @param id The process group id.
 * @return model.TemplateEntity
 */
func (self ProcessGroups) ImportTemplate(id string) model.TemplateEntity {
	s := api.Setup()
	res := model.TemplateEntity{}
	url := "http://localhost:8080/nifi-api/process-groups/{id}/templates/import"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
