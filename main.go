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
}
