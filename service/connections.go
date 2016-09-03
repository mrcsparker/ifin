package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Connections struct {
}

/**
 * Gets a connection
 *
 *
 * Tags: ["connections"]
 *
 * @param id The connection id.
 * @return model.ConnectionEntity
 */
func (self Connections) GetConnection(id string) model.ConnectionEntity {
	s := api.Setup()
	res := model.ConnectionEntity{}
	url := "http://localhost:8080/nifi-api/connections/{id}"
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
 * Updates a connection
 *
 *
 * Tags: ["connections"]
 *
 * @param id The connection id.
 * @param body The connection configuration details.
 * @return model.ConnectionEntity
 */
func (self Connections) UpdateConnection(id string, body model.ConnectionEntity) model.ConnectionEntity {
	s := api.Setup()
	res := model.ConnectionEntity{}
	url := "http://localhost:8080/nifi-api/connections/{id}"
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
 * Deletes a connection
 *
 *
 * Tags: ["connections"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The connection id.
 * @return model.ConnectionEntity
 */
func (self Connections) DeleteConnection(version string, clientId string, id string) model.ConnectionEntity {
	s := api.Setup()
	res := model.ConnectionEntity{}
	url := "http://localhost:8080/nifi-api/connections/{id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
