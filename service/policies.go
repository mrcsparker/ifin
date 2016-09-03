package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type Policies struct {
}

/**
 * Creates an access policy
 *
 *
 * Tags: ["policies"]
 *
 * @param body The access policy configuration details.
 * @return model.AccessPolicyEntity
 */
func (self Policies) CreateAccessPolicy(body model.AccessPolicyEntity) model.AccessPolicyEntity {
	s := api.Setup()
	res := model.AccessPolicyEntity{}
	url := "http://localhost:8080/nifi-api/policies"
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
 * Gets an access policy for the specified action and resource
 *
 * Will return the effective policy if no component specific policy exists for the specified action and resource. Must have Read permissions to the policy with the desired action and resource. Permissions for the policy that is returned will be indicated in the response. This means the client could be authorized to get the policy for a given component but the effective policy may be inherited from an ancestor Process Group. If the client does not have permissions to that policy, the response will not include the policy and the permissions in the response will be marked accordingly. If the client does not have permissions to the policy of the desired action and resource a 403 response will be returned.
 *
 * Tags: ["policies"]
 *
 * @param action The request action.
 * @param resource The resource of the policy.
 * @return model.AccessPolicyEntity
 */
func (self Policies) GetAccessPolicyForResource(action string, resource string) model.AccessPolicyEntity {
	s := api.Setup()
	res := model.AccessPolicyEntity{}
	url := "http://localhost:8080/nifi-api/policies/{action}/{resource}"
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
 * Gets an access policy
 *
 *
 * Tags: ["policies"]
 *
 * @param id The access policy id.
 * @return model.AccessPolicyEntity
 */
func (self Policies) GetAccessPolicy(id string) model.AccessPolicyEntity {
	s := api.Setup()
	res := model.AccessPolicyEntity{}
	url := "http://localhost:8080/nifi-api/policies/{id}"
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
 * Updates a access policy
 *
 *
 * Tags: ["policies"]
 *
 * @param id The access policy id.
 * @param body The access policy configuration details.
 * @return model.AccessPolicyEntity
 */
func (self Policies) UpdateAccessPolicy(id string, body model.AccessPolicyEntity) model.AccessPolicyEntity {
	s := api.Setup()
	res := model.AccessPolicyEntity{}
	url := "http://localhost:8080/nifi-api/policies/{id}"
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
 * Deletes an access policy
 *
 *
 * Tags: ["policies"]
 *
 * @param version The revision is used to verify the client is working with the latest version of the flow.
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The access policy id.
 * @return model.AccessPolicyEntity
 */
func (self Policies) RemoveAccessPolicy(version string, clientId string, id string) model.AccessPolicyEntity {
	s := api.Setup()
	res := model.AccessPolicyEntity{}
	url := "http://localhost:8080/nifi-api/policies/{id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
