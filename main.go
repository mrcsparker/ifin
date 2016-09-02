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

	var c = controller.ControllerServiceTypes()
	if len(c) > 0 {
		fmt.Println(c[0].Type)
		fmt.Println(c[0].Description)
		fmt.Println(c[0].Tags)
	}

	fmt.Println(controller.Counters())
	fmt.Println(controller.Identity())

	var p = controller.ProcessorTypes()
	if len(p) > 0 {
		fmt.Println(p[0].Type)
		fmt.Println(p[0].Description)
		fmt.Println(p[0].Tags)
	}

	var r = controller.ReportingTaskTypes()
	if len(r) > 0 {
		fmt.Println(r[0].Type)
		fmt.Println(r[0].Description)
		fmt.Println(r[0].Tags)
	}
}
