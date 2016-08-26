package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Payload struct {
	Revision               model.Revision                `json:"revision"`
	Controller             model.Controller              `json:"controller"`
	About                  model.About                   `json:"about"`
	Authorities            []string                      `json:"authorities"`
	Banners                model.Banners                 `json:"banners"`
	Config                 model.Config                  `json:"config"`
	ControllerServiceTypes []model.ControllerServiceType `json:"controllerServiceTypes"`
	Counters               model.Counters                `json:"counters"`
	Identity               string                        `json:"identity"`
	ProcessorTypes         []model.ProcessorType         `json:"processorTypes"`
	ReportingTaskTypes     []model.ReportingTaskType     `json:"reportingTaskTypes"`
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

func (controller *Controller) Authorities() []string {
	return get("/authorities").Authorities
}

// GET controller/banners
func (controller *Controller) Banners() model.Banners {
	return get("/banners").Banners
}

// GET controller/bulletin-board

// GET controller/config
func (controller *Controller) Config() model.Config {
	return get("/config").Config
}

// PUT controller/config

// GET controller/controller-service-types
func (controller *Controller) ControllerServiceTypes() []model.ControllerServiceType {
	return get("/controller-service-types").ControllerServiceTypes
}

// GET controller/counters
func (controller *Controller) Counters() model.Counters {
	return get("/counters").Counters
}

// PUT controller/counters/{id}

// GET/controller/identity

func (controller *Controller) Identity() string {
	return get("/identity").Identity
}

// GET controller/prioritizers

// GET controller/processor-types
func (controller *Controller) ProcessorTypes() []model.ProcessorType {
	return get("/processor-types").ProcessorTypes
}

// GET controller/reporting-task-types
func (controller *Controller) ReportingTaskTypes() []model.ReportingTaskType {
	return get("/reporting-task-types").ReportingTaskTypes
}

// GET controller/revision

// GET controller/search-results

// POST controller/snippetsGets

// GET controller/snippets/{id}

// PUT controller/snippets/{id}

// DELETE controller/snippets/{id}

// GET controller/status

// POST controller/templates

// GET controller/templates

// GET controller/templates/{id}

// DELETE controller/templates/{id}

// GET system-diagnostics
