package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type ControllerServices struct {
}

/**
 * Gets a controller service
 *
 *
 * Tags: ["controller-services"]
 *
 * @param id The controller service id.
 * @return model.ControllerServiceEntity
 */
func (self ControllerServices) GetControllerService(id string) model.ControllerServiceEntity {
	s := api.Setup()
	res := model.ControllerServiceEntity{}
	url := "http://localhost:8080/nifi-api/controller-services/{id}"
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
 * Updates a controller service
 *
 *
 * Tags: ["controller-services"]
 *
 * @param id The controller service id.
 * @param body The controller service configuration details.
 * @return model.ControllerServiceEntity
 */
func (self ControllerServices) UpdateControllerService(id string, body model.ControllerServiceEntity) model.ControllerServiceEntity {
	s := api.Setup()
	res := model.ControllerServiceEntity{}
	url := "http://localhost:8080/nifi-api/controller-services/{id}"
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
 * Deletes a controller service
 *
 *
 * Tags: ["controller-services"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The controller service id.
 * @return model.ControllerServiceEntity
 */
func (self ControllerServices) RemoveControllerService(version string, clientId string, id string) model.ControllerServiceEntity {
	s := api.Setup()
	res := model.ControllerServiceEntity{}
	url := "http://localhost:8080/nifi-api/controller-services/{id}"
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
 * Gets a controller service property descriptor
 *
 *
 * Tags: ["controller-services"]
 *
 * @param id The controller service id.
 * @param propertyName The property name to return the descriptor for.
 * @return model.PropertyDescriptorEntity
 */
func (self ControllerServices) GetPropertyDescriptor(id string, propertyName string) model.PropertyDescriptorEntity {
	s := api.Setup()
	res := model.PropertyDescriptorEntity{}
	url := "http://localhost:8080/nifi-api/controller-services/{id}/descriptors"
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
 * Gets a controller service
 *
 *
 * Tags: ["controller-services"]
 *
 * @param id The controller service id.
 * @return model.ControllerServiceEntity
 */
func (self ControllerServices) GetControllerServiceReferences(id string) model.ControllerServiceEntity {
	s := api.Setup()
	res := model.ControllerServiceEntity{}
	url := "http://localhost:8080/nifi-api/controller-services/{id}/references"
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
 * Updates a controller services references
 *
 *
 * Tags: ["controller-services"]
 *
 * @param id The controller service id.
 * @param body The controller service request update request.
 * @return model.ControllerServiceReferencingComponentsEntity
 */
func (self ControllerServices) UpdateControllerServiceReferences(id string, body model.UpdateControllerServiceReferenceRequestEntity) model.ControllerServiceReferencingComponentsEntity {
	s := api.Setup()
	res := model.ControllerServiceReferencingComponentsEntity{}
	url := "http://localhost:8080/nifi-api/controller-services/{id}/references"
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
 * Gets the state for a controller service
 *
 *
 * Tags: ["controller-services"]
 *
 * @param id The controller service id.
 * @return model.ComponentStateDTO
 */
func (self ControllerServices) GetState(id string) model.ComponentStateDTO {
	s := api.Setup()
	res := model.ComponentStateDTO{}
	url := "http://localhost:8080/nifi-api/controller-services/{id}/state"
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
 * Clears the state for a controller service
 *
 *
 * Tags: ["controller-services"]
 *
 * @param id The controller service id.
 * @return model.ComponentStateDTO
 */
func (self ControllerServices) ClearState(id string) model.ComponentStateDTO {
	s := api.Setup()
	res := model.ComponentStateDTO{}
	url := "http://localhost:8080/nifi-api/controller-services/{id}/state/clear-requests"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
