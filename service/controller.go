package service

import (
	"github.com/mrcsparker/ifin/model"
)

type Payload struct {
	Revision   model.Revision   `json:"revision"`
	Controller model.Controller `json:"controller"`
}

type Controller struct {
}

func (controller *Controller) Get() {

}
