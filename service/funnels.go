package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Funnels struct {
}

/**
 * Gets a funnel
 *
 *
 * Tags: ["funnel"]
 *
 * @param id The funnel id.
 * @return model.FunnelEntity
 */
func (self Funnels) GetFunnel(id string) model.FunnelEntity {
	s := api.Setup()
	res := model.FunnelEntity{}
	url := "http://localhost:8080/nifi-api/funnels/{id}"
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
 * Updates a funnel
 *
 *
 * Tags: ["funnel"]
 *
 * @param id The funnel id.
 * @param body The funnel configuration details.
 * @return model.FunnelEntity
 */
func (self Funnels) UpdateFunnel(id string, body model.FunnelEntity) model.FunnelEntity {
	s := api.Setup()
	res := model.FunnelEntity{}
	url := "http://localhost:8080/nifi-api/funnels/{id}"
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
 * Deletes a funnel
 *
 *
 * Tags: ["funnel"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The funnel id.
 * @return model.FunnelEntity
 */
func (self Funnels) RemoveFunnel(version string, clientId string, id string) model.FunnelEntity {
	s := api.Setup()
	res := model.FunnelEntity{}
	url := "http://localhost:8080/nifi-api/funnels/{id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
