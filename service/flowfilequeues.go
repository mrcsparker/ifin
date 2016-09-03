package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type FlowfileQueues struct {
}

/**
 * Creates a request to drop the contents of the queue in this connection.
 *
 *
 * Tags: ["flowfile-queues"]
 *
 * @param id The connection id.
 * @return model.DropRequestEntity
 */
func (self FlowfileQueues) CreateDropRequest(id string) model.DropRequestEntity {
	s := api.Setup()
	res := model.DropRequestEntity{}
	url := "http://localhost:8080/nifi-api/flowfile-queues/{id}/drop-requests"
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
 * Gets the current status of a drop request for the specified connection.
 *
 *
 * Tags: ["flowfile-queues"]
 *
 * @param id The connection id.
 * @param dropRequestId The drop request id.
 * @return model.DropRequestEntity
 */
func (self FlowfileQueues) GetDropRequest(id string, dropRequestId string) model.DropRequestEntity {
	s := api.Setup()
	res := model.DropRequestEntity{}
	url := "http://localhost:8080/nifi-api/flowfile-queues/{id}/drop-requests/{drop-request-id}"
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
 * Cancels and/or removes a request to drop the contents of this connection.
 *
 *
 * Tags: ["flowfile-queues"]
 *
 * @param id The connection id.
 * @param dropRequestId The drop request id.
 * @return model.DropRequestEntity
 */
func (self FlowfileQueues) RemoveDropRequest(id string, dropRequestId string) model.DropRequestEntity {
	s := api.Setup()
	res := model.DropRequestEntity{}
	url := "http://localhost:8080/nifi-api/flowfile-queues/{id}/drop-requests/{drop-request-id}"
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
 * Gets a FlowFile from a Connection.
 *
 *
 * Tags: ["flowfile-queues"]
 *
 * @param id The connection id.
 * @param flowfileUuid The flowfile uuid.
 * @param clusterNodeId The id of the node where the content exists if clustered.
 * @return
 */
func (self FlowfileQueues) GetFlowFile(id string, flowfileUuid string, clusterNodeId string) {
	s := api.Setup()
	res := struct{}{}
	url := "http://localhost:8080/nifi-api/flowfile-queues/{id}/flowfiles/{flowfile-uuid}"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

}

/**
 * Gets the content for a FlowFile in a Connection.
 *
 *
 * Tags: ["flowfile-queues"]
 *
 * @param clientId If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response.
 * @param id The connection id.
 * @param flowfileUuid The flowfile uuid.
 * @param clusterNodeId The id of the node where the content exists if clustered.
 * @return
 */
func (self FlowfileQueues) DownloadFlowFileContent(clientId string, id string, flowfileUuid string, clusterNodeId string) {
	s := api.Setup()
	res := struct{}{}
	url := "http://localhost:8080/nifi-api/flowfile-queues/{id}/flowfiles/{flowfile-uuid}/content"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

}

/**
 * Lists the contents of the queue in this connection.
 *
 *
 * Tags: ["flowfile-queues"]
 *
 * @param id The connection id.
 * @return model.ListingRequestEntity
 */
func (self FlowfileQueues) CreateFlowFileListing(id string) model.ListingRequestEntity {
	s := api.Setup()
	res := model.ListingRequestEntity{}
	url := "http://localhost:8080/nifi-api/flowfile-queues/{id}/listing-requests"
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
 * Gets the current status of a listing request for the specified connection.
 *
 *
 * Tags: ["flowfile-queues"]
 *
 * @param id The connection id.
 * @param listingRequestId The listing request id.
 * @return model.ListingRequestEntity
 */
func (self FlowfileQueues) GetListingRequest(id string, listingRequestId string) model.ListingRequestEntity {
	s := api.Setup()
	res := model.ListingRequestEntity{}
	url := "http://localhost:8080/nifi-api/flowfile-queues/{id}/listing-requests/{listing-request-id}"
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
 * Cancels and/or removes a request to list the contents of this connection.
 *
 *
 * Tags: ["flowfile-queues"]
 *
 * @param id The connection id.
 * @param listingRequestId The listing request id.
 * @return model.DropRequestEntity
 */
func (self FlowfileQueues) DeleteListingRequest(id string, listingRequestId string) model.DropRequestEntity {
	s := api.Setup()
	res := model.DropRequestEntity{}
	url := "http://localhost:8080/nifi-api/flowfile-queues/{id}/listing-requests/{listing-request-id}"
	resp, err := s.Delete(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
