package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type ReportingTasks struct {
}

/**
 * Gets a reporting task
 *
 *
 * Tags: ["reporting-tasks"]
 *
 * @param id The reporting task id.
 * @return model.ReportingTaskEntity
 */
func (self ReportingTasks) GetReportingTask(id string) model.ReportingTaskEntity {
	s := api.Setup()
	res := model.ReportingTaskEntity{}
	url := "http://localhost:8080/nifi-api/reporting-tasks/{id}"
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
 * Updates a reporting task
 *
 *
 * Tags: ["reporting-tasks"]
 *
 * @param id The reporting task id.
 * @param body The reporting task configuration details.
 * @return model.ReportingTaskEntity
 */
func (self ReportingTasks) UpdateReportingTask(id string, body model.ReportingTaskEntity) model.ReportingTaskEntity {
	s := api.Setup()
	res := model.ReportingTaskEntity{}
	url := "http://localhost:8080/nifi-api/reporting-tasks/{id}"
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
 * Deletes a reporting task
 *
 *
 * Tags: ["reporting-tasks"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The reporting task id.
 * @return model.ReportingTaskEntity
 */
func (self ReportingTasks) RemoveReportingTask(version string, clientId string, id string) model.ReportingTaskEntity {
	s := api.Setup()
	res := model.ReportingTaskEntity{}
	url := "http://localhost:8080/nifi-api/reporting-tasks/{id}"
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
 * Gets a reporting task property descriptor
 *
 *
 * Tags: ["reporting-tasks"]
 *
 * @param id The reporting task id.
 * @param propertyName The property name.
 * @return model.PropertyDescriptorEntity
 */
func (self ReportingTasks) GetPropertyDescriptor(id string, propertyName string) model.PropertyDescriptorEntity {
	s := api.Setup()
	res := model.PropertyDescriptorEntity{}
	url := "http://localhost:8080/nifi-api/reporting-tasks/{id}/descriptors"
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
 * Gets the state for a reporting task
 *
 *
 * Tags: ["reporting-tasks"]
 *
 * @param id The reporting task id.
 * @return model.ComponentStateDTO
 */
func (self ReportingTasks) GetState(id string) model.ComponentStateDTO {
	s := api.Setup()
	res := model.ComponentStateDTO{}
	url := "http://localhost:8080/nifi-api/reporting-tasks/{id}/state"
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
 * Clears the state for a reporting task
 *
 *
 * Tags: ["reporting-tasks"]
 *
 * @param id The reporting task id.
 * @return model.ComponentStateDTO
 */
func (self ReportingTasks) ClearState(id string) model.ComponentStateDTO {
	s := api.Setup()
	res := model.ComponentStateDTO{}
	url := "http://localhost:8080/nifi-api/reporting-tasks/{id}/state/clear-requests"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
