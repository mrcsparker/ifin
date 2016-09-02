package model

type TransactionResultEntity struct {
	FlowFileSent int    `json:"flowFileSent"`
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
