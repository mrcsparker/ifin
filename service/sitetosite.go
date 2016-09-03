package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type SiteToSite struct {
}

/**
 * Returns the details about this NiFi necessary to communicate via site to site
 *
 *
 * Tags: ["site-to-site"]
 *
 * @return model.ControllerEntity
 */
func (self SiteToSite) GetSiteToSiteDetails() model.ControllerEntity {
	s := api.Setup()
	res := model.ControllerEntity{}
	url := "http://localhost:8080/nifi-api/site-to-site"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}

/**
 * Returns the available Peers and its status of this NiFi
 *
 *
 * Tags: ["site-to-site"]
 *
 * @return model.PeersEntity
 */
func (self SiteToSite) GetPeers() model.PeersEntity {
	s := api.Setup()
	res := model.PeersEntity{}
	url := "http://localhost:8080/nifi-api/site-to-site/peers"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
