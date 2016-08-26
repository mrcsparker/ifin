package main

import (
	"fmt"
	"github.com/mrcsparker/ifin/service"
)

func welcome() string {
	return "[ifin]"
}

func main() {
	fmt.Println(welcome())

	controller := service.Controller{}
	fmt.Println(controller.Get())
	fmt.Println(controller.About())
	fmt.Println(controller.Authorities())
	fmt.Println(controller.Banners())
	fmt.Println(controller.Config())

	var c = controller.ControllerServiceTypes()[0]
	fmt.Println(c.Type)
	fmt.Println(c.Description)
	fmt.Println(c.Tags)

	fmt.Println(controller.Counters())
	fmt.Println(controller.Identity())

	var p = controller.ProcessorTypes()[0]
	fmt.Println(p.Type)
	fmt.Println(p.Description)
	fmt.Println(p.Tags)

	var r = controller.ReportingTaskTypes()[0]
	fmt.Println(r.Type)
	fmt.Println(r.Description)
	fmt.Println(r.Tags)
}
