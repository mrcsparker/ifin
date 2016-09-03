package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Labels struct {
}

/**
 * Gets a label
 *
 *
 * Tags: ["labels"]
 *
 * @param id The label id.
 * @return model.LabelEntity
 */
func (self Labels) GetLabel(id string) model.LabelEntity {
	s := api.Setup()
	res := model.LabelEntity{}
	url := "http://localhost:8080/nifi-api/labels/{id}"
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
 * Updates a label
 *
 *
 * Tags: ["labels"]
 *
 * @param id The label id.
 * @param body The label configuration details.
 * @return model.LabelEntity
 */
func (self Labels) UpdateLabel(id string, body model.LabelEntity) model.LabelEntity {
	s := api.Setup()
	res := model.LabelEntity{}
	url := "http://localhost:8080/nifi-api/labels/{id}"
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
 * Deletes a label
 *
 *
 * Tags: ["labels"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The label id.
 * @return model.LabelEntity
 */
func (self Labels) RemoveLabel(version string, clientId string, id string) model.LabelEntity {
	s := api.Setup()
	res := model.LabelEntity{}
	url := "http://localhost:8080/nifi-api/labels/{id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
