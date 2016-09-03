package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Resources struct {
}

/**
 * Gets the available resources that support access/authorization policies
 *
 *
 * Tags: ["resources"]
 *
 * @return model.ResourcesEntity
 */
func (self Resources) GetResources() model.ResourcesEntity {
	s := api.Setup()
	res := model.ResourcesEntity{}
	url := "http://localhost:8080/nifi-api/resources"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
