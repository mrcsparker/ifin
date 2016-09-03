package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Provenance struct {
}

/**
 * Submits a provenance query
 *
 * Provenance queries may be long running so this endpoint submits a request. The response will include the current state of the query. If the request is not completed the URI in the response can be used at a later time to get the updated state of the query. Once the query has completed the provenance request should be deleted by the client who originally submitted it.
 *
 * Tags: ["provenance"]
 *
 * @param body The provenance query details.
 * @return model.ProvenanceEntity
 */
func (self Provenance) SubmitProvenanceRequest(body model.ProvenanceEntity) model.ProvenanceEntity {
	s := api.Setup()
	res := model.ProvenanceEntity{}
	url := "http://localhost:8080/nifi-api/provenance"
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
 * Submits a lineage query
 *
 * Lineage queries may be long running so this endpoint submits a request. The response will include the current state of the query. If the request is not completed the URI in the response can be used at a later time to get the updated state of the query. Once the query has completed the lineage request should be deleted by the client who originally submitted it.
 *
 * Tags: ["provenance"]
 *
 * @param body The lineage query details.
 * @return model.LineageEntity
 */
func (self Provenance) SubmitLineageRequest(body model.LineageEntity) model.LineageEntity {
	s := api.Setup()
	res := model.LineageEntity{}
	url := "http://localhost:8080/nifi-api/provenance/lineage"
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
 * Gets a lineage query
 *
 *
 * Tags: ["provenance"]
 *
 * @param clusterNodeId The id of the node where this query exists if clustered.
 * @param id The id of the lineage query.
 * @return model.LineageEntity
 */
func (self Provenance) GetLineage(clusterNodeId string, id string) model.LineageEntity {
	s := api.Setup()
	res := model.LineageEntity{}
	url := "http://localhost:8080/nifi-api/provenance/lineage/{id}"
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
 * Deletes a lineage query
 *
 *
 * Tags: ["provenance"]
 *
 * @param clusterNodeId The id of the node where this query exists if clustered.
 * @param id The id of the lineage query.
 * @return model.LineageEntity
 */
func (self Provenance) DeleteLineage(clusterNodeId string, id string) model.LineageEntity {
	s := api.Setup()
	res := model.LineageEntity{}
	url := "http://localhost:8080/nifi-api/provenance/lineage/{id}"
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
 * Gets the searchable attributes for provenance events
 *
 *
 * Tags: ["provenance"]
 *
 * @return model.ProvenanceOptionsEntity
 */
func (self Provenance) GetSearchOptions() model.ProvenanceOptionsEntity {
	s := api.Setup()
	res := model.ProvenanceOptionsEntity{}
	url := "http://localhost:8080/nifi-api/provenance/search-options"
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
 * Gets a provenance query
 *
 *
 * Tags: ["provenance"]
 *
 * @param clusterNodeId The id of the node where this query exists if clustered.
 * @param id The id of the provenance query.
 * @return model.ProvenanceEntity
 */
func (self Provenance) GetProvenance(clusterNodeId string, id string) model.ProvenanceEntity {
	s := api.Setup()
	res := model.ProvenanceEntity{}
	url := "http://localhost:8080/nifi-api/provenance/{id}"
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
 * Deletes a provenance query
 *
 *
 * Tags: ["provenance"]
 *
 * @param clusterNodeId The id of the node where this query exists if clustered.
 * @param id The id of the provenance query.
 * @return model.ProvenanceEntity
 */
func (self Provenance) DeleteProvenance(clusterNodeId string, id string) model.ProvenanceEntity {
	s := api.Setup()
	res := model.ProvenanceEntity{}
	url := "http://localhost:8080/nifi-api/provenance/{id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
