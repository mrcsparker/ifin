package model

type CounterDTO struct {
	Id         string `json:"id"`         // The id of the counter.
	Context    string `json:"context"`    // The context of the counter.
	Name       string `json:"name"`       // The name of the counter.
	ValueCount int    `json:"valueCount"` // The value count.
	Value      string `json:"value"`      // The value of the counter.
}
