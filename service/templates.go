package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Templates struct {
}

/**
 * Deletes a template
 *
 *
 * Tags: ["templates"]
 *
 * @param id The template id.
 * @return model.TemplateEntity
 */
func (self Templates) RemoveTemplate(id string) model.TemplateEntity {
	s := api.Setup()
	res := model.TemplateEntity{}
	url := "http://localhost:8080/nifi-api/templates/{id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Exports a template
 *
 *
 * Tags: ["templates"]
 *
 * @param id The template id.
 * @return model.TemplateDTO
 */
func (self Templates) ExportTemplate(id string) model.TemplateDTO {
	s := api.Setup()
	res := model.TemplateDTO{}
	url := "http://localhost:8080/nifi-api/templates/{id}/download"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
