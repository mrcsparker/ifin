package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Payload struct {
	Revision   model.Revision   `json:"revision"`
	Controller model.Controller `json:"controller"`
	About      model.About      `json:"About"`
}

type Controller struct {
}

// Ugly for now.  Will refactor when I get all the services in

func (controller *Controller) Get() model.Controller {
	s := api.Setup()
	url := "http://localhost:8080/nifi-api/controller"
	res := Payload{}

	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res.Controller

}

func (controller *Controller) About() model.About {
	s := api.Setup()
	url := "http://localhost:8080/nifi-api/controller/about"
	res := Payload{}

	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res.About

}
