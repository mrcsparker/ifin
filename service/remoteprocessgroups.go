package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type RemoteProcessGroups struct {
}

/**
 * Gets a remote process group
 *
 *
 * Tags: ["remote-process-groups"]
 *
 * @param id The remote process group id.
 * @return model.RemoteProcessGroupEntity
 */
func (self RemoteProcessGroups) GetRemoteProcessGroup(id string) model.RemoteProcessGroupEntity {
	s := api.Setup()
	res := model.RemoteProcessGroupEntity{}
	url := "http://localhost:8080/nifi-api/remote-process-groups/{id}"
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
 * Updates a remote process group
 *
 *
 * Tags: ["remote-process-groups"]
 *
 * @param id The remote process group id.
 * @param body The remote process group.
 * @return model.RemoteProcessGroupEntity
 */
func (self RemoteProcessGroups) UpdateRemoteProcessGroup(id string, body model.RemoteProcessGroupEntity) model.RemoteProcessGroupEntity {
	s := api.Setup()
	res := model.RemoteProcessGroupEntity{}
	url := "http://localhost:8080/nifi-api/remote-process-groups/{id}"
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
 * Deletes a remote process group
 *
 *
 * Tags: ["remote-process-groups"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The remote process group id.
 * @return model.RemoteProcessGroupEntity
 */
func (self RemoteProcessGroups) RemoveRemoteProcessGroup(version string, clientId string, id string) model.RemoteProcessGroupEntity {
	s := api.Setup()
	res := model.RemoteProcessGroupEntity{}
	url := "http://localhost:8080/nifi-api/remote-process-groups/{id}"
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
 * Updates a remote port
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["remote-process-groups"]
 *
 * @param id The remote process group id.
 * @param portId The remote process group port id.
 * @param body The remote process group port.
 * @return model.RemoteProcessGroupPortEntity
 */
func (self RemoteProcessGroups) UpdateRemoteProcessGroupInputPort(id string, portId string, body model.RemoteProcessGroupPortEntity) model.RemoteProcessGroupPortEntity {
	s := api.Setup()
	res := model.RemoteProcessGroupPortEntity{}
	url := "http://localhost:8080/nifi-api/remote-process-groups/{id}/input-ports/{port-id}"
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
 * Updates a remote port
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["remote-process-groups"]
 *
 * @param id The remote process group id.
 * @param portId The remote process group port id.
 * @param body The remote process group port.
 * @return model.RemoteProcessGroupPortEntity
 */
func (self RemoteProcessGroups) UpdateRemoteProcessGroupOutputPort(id string, portId string, body model.RemoteProcessGroupPortEntity) model.RemoteProcessGroupPortEntity {
	s := api.Setup()
	res := model.RemoteProcessGroupPortEntity{}
	url := "http://localhost:8080/nifi-api/remote-process-groups/{id}/output-ports/{port-id}"
	resp, err := s.Put(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
