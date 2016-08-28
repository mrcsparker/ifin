package model

type SystemDiagnostics struct {
	TotalNonHeap                   string              `json:"totalNonHeap"`                   // Total size of non heap.
	UsedNonHeap                    string              `json:"usedNonHeap"`                    // Amount of use non heap.
	FreeNonHeap                    string              `json:"freeNonHeap"`                    // Amount of free non heap.
	MaxNonHeap                     string              `json:"maxNonHeap"`                     // Maximum size of non heap.
	NonHeapUtilization             string              `json:"nonHeapUtilization"`             // Utilization of non heap.
	TotalHeap                      string              `json:"totalHeap"`                      // Total size of heap.
	UsedHeap                       string              `json:"usedHeap"`                       // Amount of used heap.
	FreeHeap                       string              `json:"freeHeap"`                       // Amount of free heap.
	MaxHeap                        string              `json:"maxHeap"`                        // Maximum size of heap.
	HeapUtilization                string              `json:"heapUtilization"`                // Utilization of heap.
	AvailableProcessors            int32               `json:"availableProcessors"`            // Number of available processors if supported by the underlying system.
	ProcessorLoadAverage           float64             `json:"processorLoadAverage"`           // The processor load average if supported by the underlying system.
	TotalThreads                   int32               `json:"totalThreads"`                   // Total number of threads.
	DaemonThreads                  int32               `json:"daemonThreads"`                  // Number of daemon threads.
	FlowFileRepositoryStorageUsage StorageUsage        `json:"flowFileRepositoryStorageUsage"` // The flowfile repository storage usage.
	ContentRepositoryStorageUsage  []StorageUsage      `json:"contentRepositoryStorageUsage"`  // The content repository storage usage.
	GarbageCollection              []GarbageCollection `json:"garbageCollection"`              // The garbage collection details.
	StatsLastRefreshed             string              `json:"statsLastRefreshed"`             // When the diagnostics were generated.
}
