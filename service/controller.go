package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Payload struct {
	Revision    model.Revision    `json:"revision"`
	Controller  model.Controller  `json:"controller"`
	About       model.About       `json:"about"`
	Authorities model.Authorities `json:"authorities"`
	Banners     model.Banners     `json:"banners"`
}

type Controller struct {
}

func get(geturl string) Payload {
	s := api.Setup()
	res := Payload{}
	url := "http://localhost:8080/nifi-api/controller" + geturl
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

// Ugly for now.  Will refactor when I get all the services in

// GET controller
func (controller *Controller) Get() model.Controller {
	return get("").Controller
}

// GET controller/about
func (controller *Controller) About() model.About {
	return get("/about").About
}

// POST controller/archive

// GET controller/authorities

func (controller *Controller) Authorities() model.Authorities {
	return get("/authorities").Authorities
}

// GET controller/banners
func (controller *Controller) Banners() model.Banners {
	return get("/banners").Banners
}

// GET controller/bulletin-board

// GET controller/config

// PUT controller/config

// GET controller/controller-service-types

// GET controller/counters

// PUT controller/counters/{id}
