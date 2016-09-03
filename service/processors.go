package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Processors struct {
}

/**
 * Gets a processor
 *
 *
 * Tags: ["processors"]
 *
 * @param id The processor id.
 * @return model.ProcessorEntity
 */
func (self Processors) GetProcessor(id string) model.ProcessorEntity {
	s := api.Setup()
	res := model.ProcessorEntity{}
	url := "http://localhost:8080/nifi-api/processors/{id}"
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
 * Updates a processor
 *
 *
 * Tags: ["processors"]
 *
 * @param id The processor id.
 * @param body The processor configuration details.
 * @return model.ProcessorEntity
 */
func (self Processors) UpdateProcessor(id string, body model.ProcessorEntity) model.ProcessorEntity {
	s := api.Setup()
	res := model.ProcessorEntity{}
	url := "http://localhost:8080/nifi-api/processors/{id}"
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
 * Deletes a processor
 *
 *
 * Tags: ["processors"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The processor id.
 * @return model.ProcessorEntity
 */
func (self Processors) DeleteProcessor(version string, clientId string, id string) model.ProcessorEntity {
	s := api.Setup()
	res := model.ProcessorEntity{}
	url := "http://localhost:8080/nifi-api/processors/{id}"
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
 * Gets the descriptor for a processor property
 *
 *
 * Tags: ["processors"]
 *
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The processor id.
 * @param propertyName The property name.
 * @return model.PropertyDescriptorEntity
 */
func (self Processors) GetPropertyDescriptor(clientId string, id string, propertyName string) model.PropertyDescriptorEntity {
	s := api.Setup()
	res := model.PropertyDescriptorEntity{}
	url := "http://localhost:8080/nifi-api/processors/{id}/descriptors"
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
 * Gets the state for a processor
 *
 *
 * Tags: ["processors"]
 *
 * @param id The processor id.
 * @return model.ComponentStateDTO
 */
func (self Processors) GetState(id string) model.ComponentStateDTO {
	s := api.Setup()
	res := model.ComponentStateDTO{}
	url := "http://localhost:8080/nifi-api/processors/{id}/state"
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
 * Clears the state for a processor
 *
 *
 * Tags: ["processors"]
 *
 * @param id The processor id.
 * @return model.ComponentStateDTO
 */
func (self Processors) ClearState(id string) model.ComponentStateDTO {
	s := api.Setup()
	res := model.ComponentStateDTO{}
	url := "http://localhost:8080/nifi-api/processors/{id}/state/clear-requests"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
