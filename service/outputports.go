package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type OutputPorts struct {
}

/**
 * Gets an output port
 *
 *
 * Tags: ["output-ports"]
 *
 * @param id The output port id.
 * @return model.PortEntity
 */
func (self OutputPorts) GetOutputPort(id string) model.PortEntity {
	s := api.Setup()
	res := model.PortEntity{}
	url := "http://localhost:8080/nifi-api/output-ports/{id}"
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
 * Updates an output port
 *
 *
 * Tags: ["output-ports"]
 *
 * @param id The output port id.
 * @param body The output port configuration details.
 * @return model.PortEntity
 */
func (self OutputPorts) UpdateOutputPort(id string, body model.PortEntity) model.PortEntity {
	s := api.Setup()
	res := model.PortEntity{}
	url := "http://localhost:8080/nifi-api/output-ports/{id}"
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
 * Deletes an output port
 *
 *
 * Tags: ["output-ports"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The output port id.
 * @return model.PortEntity
 */
func (self OutputPorts) RemoveOutputPort(version string, clientId string, id string) model.PortEntity {
	s := api.Setup()
	res := model.PortEntity{}
	url := "http://localhost:8080/nifi-api/output-ports/{id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
