package model

type QueueSizeDTO struct {
	ByteCount   int64 `json:"byteCount"`   // The size of objects in a queue.
	ObjectCount int32 `json:"objectCount"` // The count of objects in a queue.
}
