package model

// attribute: transactionResultEntity
type TransactionResultEntity struct {
	FlowFileSent int32  `json:"flowFileSent"`
	ResponseCode int32  `json:"responseCode"`
	Message      string `json:"message"`
}
