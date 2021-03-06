package model

type GarbageCollectionDTO struct {
	Name             string `json:"name"`             // The name of the garbage collector.
	CollectionCount  int64  `json:"collectionCount"`  // The number of times garbage collection has run.
	CollectionTime   string `json:"collectionTime"`   // The total amount of time spent garbage collecting.
	CollectionMillis int64  `json:"collectionMillis"` // The total number of milliseconds spent garbage collecting.
}
