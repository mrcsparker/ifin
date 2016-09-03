package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Access struct {
}

/**
 * Gets the status the client's access
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["access"]
 *
 * @return model.AccessStatusEntity
 */
func (self Access) GetAccessStatus() model.AccessStatusEntity {
	s := api.Setup()
	res := model.AccessStatusEntity{}
	url := "http://localhost:8080/nifi-api/access"
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
 * Retrieves the access configuration for this NiFi
 *
 *
 * Tags: ["access"]
 *
 * @return model.AccessConfigurationEntity
 */
func (self Access) GetLoginConfig() model.AccessConfigurationEntity {
	s := api.Setup()
	res := model.AccessConfigurationEntity{}
	url := "http://localhost:8080/nifi-api/access/config"
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
 * Creates a single use access token for downloading FlowFile content.
 *
 * The token returned is a base64 encoded string. It is valid for a single request up to five minutes from being issued. It is used as a query parameter name 'access_token'.
 *
 * Tags: ["access"]
 *
 * @return string
 */
func (self Access) CreateDownloadToken() string {
	s := api.Setup()
	res := ""
	url := "http://localhost:8080/nifi-api/access/download-token"
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
 * Creates a token for accessing the REST API via Kerberos ticket exchange / SPNEGO negotiation
 *
 * The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.
 *
 * Tags: ["access"]
 *
 * @return string
 */
func (self Access) CreateAccessTokenFromTicket() string {
	s := api.Setup()
	res := ""
	url := "http://localhost:8080/nifi-api/access/kerberos"
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
 * Creates a token for accessing the REST API via username/password
 *
 * The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.
 *
 * Tags: ["access"]
 *
 * @param username
 * @param password
 * @return string
 */
func (self Access) CreateAccessToken(username string, password string) string {
	s := api.Setup()
	res := ""
	url := "http://localhost:8080/nifi-api/access/token"
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
 * Creates a single use access token for accessing a NiFi UI extension.
 *
 * The token returned is a base64 encoded string. It is valid for a single request up to five minutes from being issued. It is used as a query parameter name 'access_token'.
 *
 * Tags: ["access"]
 *
 * @return string
 */
func (self Access) CreateUiExtensionToken() string {
	s := api.Setup()
	res := ""
	url := "http://localhost:8080/nifi-api/access/ui-extension-token"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
