package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Controller struct {
}

/**
 * Gets the contents of the cluster
 *
 * Returns the contents of the cluster including all nodes and their status.
 *
 * Tags: ["controller"]
 *
 * @return model.ClusterEntity
 */
func (self Controller) GetCluster() model.ClusterEntity {
	s := api.Setup()
	res := model.ClusterEntity{}
	url := "http://localhost:8080/nifi-api/controller/cluster"
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
 * Gets a node in the cluster
 *
 *
 * Tags: ["controller"]
 *
 * @param id The node id.
 * @return model.NodeEntity
 */
func (self Controller) GetNode(id string) model.NodeEntity {
	s := api.Setup()
	res := model.NodeEntity{}
	url := "http://localhost:8080/nifi-api/controller/cluster/nodes/{id}"
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
 * Updates a node in the cluster
 *
 *
 * Tags: ["controller"]
 *
 * @param id The node id.
 * @param body The node configuration. The only configuration that will be honored at this endpoint is the status or primary flag.
 * @return model.NodeEntity
 */
func (self Controller) UpdateNode(id string, body model.NodeEntity) model.NodeEntity {
	s := api.Setup()
	res := model.NodeEntity{}
	url := "http://localhost:8080/nifi-api/controller/cluster/nodes/{id}"
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
 * Removes a node from the cluster
 *
 *
 * Tags: ["controller"]
 *
 * @param id The node id.
 * @return model.NodeEntity
 */
func (self Controller) DeleteNode(id string) model.NodeEntity {
	s := api.Setup()
	res := model.NodeEntity{}
	url := "http://localhost:8080/nifi-api/controller/cluster/nodes/{id}"
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
 * Retrieves the configuration for this NiFi Controller
 *
 *
 * Tags: ["controller"]
 *
 * @return model.ControllerConfigurationEntity
 */
func (self Controller) GetControllerConfig() model.ControllerConfigurationEntity {
	s := api.Setup()
	res := model.ControllerConfigurationEntity{}
	url := "http://localhost:8080/nifi-api/controller/config"
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
 * Retrieves the configuration for this NiFi
 *
 *
 * Tags: ["controller"]
 *
 * @param body The controller configuration.
 * @return model.ControllerConfigurationEntity
 */
func (self Controller) UpdateControllerConfig(body model.ControllerConfigurationEntity) model.ControllerConfigurationEntity {
	s := api.Setup()
	res := model.ControllerConfigurationEntity{}
	url := "http://localhost:8080/nifi-api/controller/config"
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
 * Creates a new controller service
 *
 *
 * Tags: ["controller"]
 *
 * @param body The controller service configuration details.
 * @return model.ControllerServiceEntity
 */
func (self Controller) CreateControllerService(body model.ControllerServiceEntity) model.ControllerServiceEntity {
	s := api.Setup()
	res := model.ControllerServiceEntity{}
	url := "http://localhost:8080/nifi-api/controller/controller-services"
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
 * Purges history
 *
 *
 * Tags: ["controller"]
 *
 * @param endDate Purge actions before this date/time.
 * @return model.HistoryEntity
 */
func (self Controller) DeleteHistory(endDate string) model.HistoryEntity {
	s := api.Setup()
	res := model.HistoryEntity{}
	url := "http://localhost:8080/nifi-api/controller/history"
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
 * Creates a new reporting task
 *
 *
 * Tags: ["controller"]
 *
 * @param body The reporting task configuration details.
 * @return model.ReportingTaskEntity
 */
func (self Controller) CreateReportingTask(body model.ReportingTaskEntity) model.ReportingTaskEntity {
	s := api.Setup()
	res := model.ReportingTaskEntity{}
	url := "http://localhost:8080/nifi-api/controller/reporting-tasks"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
