package service

import (
	"fmt"
	"github.com/mrcsparker/ifin/api"
	"github.com/mrcsparker/ifin/model"
	"log"
)

type DataTransfer struct {
}

/**
 * Extend transaction TTL
 *
 *
 * Tags: ["data-transfer"]
 *
 * @param portId
 * @param transactionId
 * @return model.TransactionResultEntity
 */
func (self DataTransfer) ExtendInputPortTransactionTTL(portId string, transactionId string) model.TransactionResultEntity {
	s := api.Setup()
	res := model.TransactionResultEntity{}
	url := "http://localhost:8080/nifi-api/data-transfer/input-ports/{portId}/transactions/{transactionId}"
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
 * Commit or cancel the specified transaction
 *
 *
 * Tags: ["data-transfer"]
 *
 * @param responseCode The response code. Available values are BAD_CHECKSUM(19), CONFIRM_TRANSACTION(12) or CANCEL_TRANSACTION(15).
 * @param portId The input port id.
 * @param transactionId The transaction id.
 * @return model.TransactionResultEntity
 */
func (self DataTransfer) CommitInputPortTransaction(responseCode int, portId string, transactionId string) model.TransactionResultEntity {
	s := api.Setup()
	res := model.TransactionResultEntity{}
	url := "http://localhost:8080/nifi-api/data-transfer/input-ports/{portId}/transactions/{transactionId}"
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
 * Transfer flow files to the input port
 *
 *
 * Tags: ["data-transfer"]
 *
 * @param portId The input port id.
 * @param transactionId
 * @return string
 */
func (self DataTransfer) ReceiveFlowFiles(portId string, transactionId string) string {
	s := api.Setup()
	res := ""
	url := "http://localhost:8080/nifi-api/data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files"
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
 * Extend transaction TTL
 *
 *
 * Tags: ["data-transfer"]
 *
 * @param portId
 * @param transactionId
 * @return model.TransactionResultEntity
 */
func (self DataTransfer) ExtendOutputPortTransactionTTL(portId string, transactionId string) model.TransactionResultEntity {
	s := api.Setup()
	res := model.TransactionResultEntity{}
	url := "http://localhost:8080/nifi-api/data-transfer/output-ports/{portId}/transactions/{transactionId}"
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
 * Commit or cancel the specified transaction
 *
 *
 * Tags: ["data-transfer"]
 *
 * @param responseCode The response code. Available values are CONFIRM_TRANSACTION(12) or CANCEL_TRANSACTION(15).
 * @param checksum A checksum calculated at client side using CRC32 to check flow file content integrity. It must match with the value calculated at server side.
 * @param portId The output port id.
 * @param transactionId The transaction id.
 * @return model.TransactionResultEntity
 */
func (self DataTransfer) CommitOutputPortTransaction(responseCode int, checksum string, portId string, transactionId string) model.TransactionResultEntity {
	s := api.Setup()
	res := model.TransactionResultEntity{}
	url := "http://localhost:8080/nifi-api/data-transfer/output-ports/{portId}/transactions/{transactionId}"
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
 * Transfer flow files from the output port
 *
 *
 * Tags: ["data-transfer"]
 *
 * @param portId The output port id.
 * @param transactionId
 * @return
 */
func (self DataTransfer) TransferFlowFiles(portId string, transactionId string) {
	s := api.Setup()
	res := struct{}{}
	url := "http://localhost:8080/nifi-api/data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files"
	resp, err := s.Get(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

}

/**
 * Create a transaction to the specified output port or input port
 *
 *
 * Tags: ["data-transfer"]
 *
 * @param portType The port type.
 * @param portId
 * @return model.TransactionResultEntity
 */
func (self DataTransfer) CreatePortTransaction(portType string, portId string) model.TransactionResultEntity {
	s := api.Setup()
	res := model.TransactionResultEntity{}
	url := "http://localhost:8080/nifi-api/data-transfer/{portType}/{portId}/transactions"
	resp, err := s.Post(url, nil, &res, nil)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status() != 200 {
		fmt.Println(res)
	}

	return res
}
