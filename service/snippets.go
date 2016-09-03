package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Snippets struct {
}

/**
 * Creates a snippet
 *
 *
 * Tags: ["snippets"]
 *
 * @param body The snippet configuration details.
 * @return model.SnippetEntity
 */
func (self Snippets) CreateSnippet(body model.SnippetEntity) model.SnippetEntity {
	s := api.Setup()
	res := model.SnippetEntity{}
	url := "http://localhost:8080/nifi-api/snippets"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Move's the components in this Snippet into a new Process Group and drops the snippet
 *
 *
 * Tags: ["snippets"]
 *
 * @param id The snippet id.
 * @param body The snippet configuration details.
 * @return model.SnippetEntity
 */
func (self Snippets) UpdateSnippet(id string, body model.SnippetEntity) model.SnippetEntity {
	s := api.Setup()
	res := model.SnippetEntity{}
	url := "http://localhost:8080/nifi-api/snippets/{id}"
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
 * Deletes the components in a snippet and drops the snippet
 *
 *
 * Tags: ["snippets"]
 *
 * @param id The snippet id.
 * @return model.SnippetEntity
 */
func (self Snippets) DeleteSnippet(id string) model.SnippetEntity {
	s := api.Setup()
	res := model.SnippetEntity{}
	url := "http://localhost:8080/nifi-api/snippets/{id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
