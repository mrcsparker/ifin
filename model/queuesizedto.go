package model

type QueueSizeDTO struct {
	ByteCount   int `json:"byteCount"`   // The size of objects in a queue.
	ObjectCount int `json:"objectCount"` // The count of objects in a queue.
}
