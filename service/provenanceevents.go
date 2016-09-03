package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type ProvenanceEvents struct {
}

/**
 * Replays content from a provenance event
 *
 *
 * Tags: ["provenance-events"]
 *
 * @param body The replay request.
 * @return model.ProvenanceEventEntity
 */
func (self ProvenanceEvents) SubmitReplay(body model.SubmitReplayRequestEntity) model.ProvenanceEventEntity {
	s := api.Setup()
	res := model.ProvenanceEventEntity{}
	url := "http://localhost:8080/nifi-api/provenance-events/replays"
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
 * Gets a provenance event
 *
 *
 * Tags: ["provenance-events"]
 *
 * @param clusterNodeId The id of the node where this event exists if clustered.
 * @param id The provenance event id.
 * @return model.ProvenanceEventEntity
 */
func (self ProvenanceEvents) GetProvenanceEvent(clusterNodeId string, id string) model.ProvenanceEventEntity {
	s := api.Setup()
	res := model.ProvenanceEventEntity{}
	url := "http://localhost:8080/nifi-api/provenance-events/{id}"
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
 * Gets the input content for a provenance event
 *
 *
 * Tags: ["provenance-events"]
 *
 * @param clusterNodeId The id of the node where the content exists if clustered.
 * @param id The provenance event id.
 * @return
 */
func (self ProvenanceEvents) GetInputContent(clusterNodeId string, id string) {
	s := api.Setup()
	res := struct{}{}
	url := "http://localhost:8080/nifi-api/provenance-events/{id}/content/input"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

}

/**
 * Gets the output content for a provenance event
 *
 *
 * Tags: ["provenance-events"]
 *
 * @param clusterNodeId The id of the node where the content exists if clustered.
 * @param id The provenance event id.
 * @return
 */
func (self ProvenanceEvents) GetOutputContent(clusterNodeId string, id string) {
	s := api.Setup()
	res := struct{}{}
	url := "http://localhost:8080/nifi-api/provenance-events/{id}/content/output"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

}
