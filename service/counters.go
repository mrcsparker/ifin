package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Counters struct {
}

/**
 * Gets the current counters for this NiFi
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["counters"]
 *
 * @param nodewise Whether or not to include the breakdown per node. Optional, defaults to false
 * @param clusterNodeId The id of the node where to get the status.
 * @return model.CountersEntity
 */
func (self Counters) GetCounters(nodewise bool, clusterNodeId string) model.CountersEntity {
	s := api.Setup()
	res := model.CountersEntity{}
	url := "http://localhost:8080/nifi-api/counters"
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
 * Updates the specified counter. This will reset the counter value to 0
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["counters"]
 *
 * @param id The id of the counter.
 * @return model.CounterEntity
 */
func (self Counters) UpdateCounter(id string) model.CounterEntity {
	s := api.Setup()
	res := model.CounterEntity{}
	url := "http://localhost:8080/nifi-api/counters/{id}"
	resp, err := s.Put(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
