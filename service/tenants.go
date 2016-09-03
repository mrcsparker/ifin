package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Tenants struct {
}

/**
 * Searches the cluster for a node with the specified address
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @param q Node address to search for.
 * @return model.ClusterSearchResultsEntity
 */
func (self Tenants) SearchCluster(q string) model.ClusterSearchResultsEntity {
	s := api.Setup()
	res := model.ClusterSearchResultsEntity{}
	url := "http://localhost:8080/nifi-api/tenants/search-results"
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
 * Gets all user groups
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @return model.UserGroupsEntity
 */
func (self Tenants) GetUserGroups() model.UserGroupsEntity {
	s := api.Setup()
	res := model.UserGroupsEntity{}
	url := "http://localhost:8080/nifi-api/tenants/user-groups"
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
 * Creates a user group
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @param body The user group configuration details.
 * @return model.UserGroupEntity
 */
func (self Tenants) CreateUserGroup(body model.UserGroupEntity) model.UserGroupEntity {
	s := api.Setup()
	res := model.UserGroupEntity{}
	url := "http://localhost:8080/nifi-api/tenants/user-groups"
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
 * Gets a user group
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @param id The user group id.
 * @return model.UserGroupEntity
 */
func (self Tenants) GetUserGroup(id string) model.UserGroupEntity {
	s := api.Setup()
	res := model.UserGroupEntity{}
	url := "http://localhost:8080/nifi-api/tenants/user-groups/{id}"
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
 * Updates a user group
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @param id The user group id.
 * @param body The user group configuration details.
 * @return model.UserGroupEntity
 */
func (self Tenants) UpdateUserGroup(id string, body model.UserGroupEntity) model.UserGroupEntity {
	s := api.Setup()
	res := model.UserGroupEntity{}
	url := "http://localhost:8080/nifi-api/tenants/user-groups/{id}"
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
 * Deletes a user group
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The user group id.
 * @return model.UserGroupEntity
 */
func (self Tenants) RemoveUserGroup(version string, clientId string, id string) model.UserGroupEntity {
	s := api.Setup()
	res := model.UserGroupEntity{}
	url := "http://localhost:8080/nifi-api/tenants/user-groups/{id}"
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
 * Gets all users
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @return model.UsersEntity
 */
func (self Tenants) GetUsers() model.UsersEntity {
	s := api.Setup()
	res := model.UsersEntity{}
	url := "http://localhost:8080/nifi-api/tenants/users"
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
 * Creates a user
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @param body The user configuration details.
 * @return model.UserEntity
 */
func (self Tenants) CreateUser(body model.UserEntity) model.UserEntity {
	s := api.Setup()
	res := model.UserEntity{}
	url := "http://localhost:8080/nifi-api/tenants/users"
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
 * Gets a user
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @param id The user id.
 * @return model.UserEntity
 */
func (self Tenants) GetUser(id string) model.UserEntity {
	s := api.Setup()
	res := model.UserEntity{}
	url := "http://localhost:8080/nifi-api/tenants/users/{id}"
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
 * Updates a user
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @param id The user id.
 * @param body The user configuration details.
 * @return model.UserEntity
 */
func (self Tenants) UpdateUser(id string, body model.UserEntity) model.UserEntity {
	s := api.Setup()
	res := model.UserEntity{}
	url := "http://localhost:8080/nifi-api/tenants/users/{id}"
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
 * Deletes a user
 *
 * Note: This endpoint is subject to change as NiFi and it's REST API evolve.
 *
 * Tags: ["tenants"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The user id.
 * @return model.UserEntity
 */
func (self Tenants) RemoveUser(version string, clientId string, id string) model.UserEntity {
	s := api.Setup()
	res := model.UserEntity{}
	url := "http://localhost:8080/nifi-api/tenants/users/{id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
