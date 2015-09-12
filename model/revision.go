package model

type Revision struct {
	ClientId string `json:"clientId"`
	Version int `json:"version"`
	LastModifier string `json:"lastModifier"`
}
