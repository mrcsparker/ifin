package model

type SystemDiagnosticsSnapshotDTO struct {
	TotalNonHeap                   string                 `json:"totalNonHeap"`                   // Total size of non heap.
	TotalNonHeapBytes              int64                  `json:"totalNonHeapBytes"`              // Total number of bytes allocated to the JVM not used for heap
	UsedNonHeap                    string                 `json:"usedNonHeap"`                    // Amount of use non heap.
	UsedNonHeapBytes               int64                  `json:"usedNonHeapBytes"`               // Total number of bytes used by the JVM not in the heap space
	FreeNonHeap                    string                 `json:"freeNonHeap"`                    // Amount of free non heap.
	FreeNonHeapBytes               int64                  `json:"freeNonHeapBytes"`               // Total number of free non-heap bytes available to the JVM
	MaxNonHeap                     string                 `json:"maxNonHeap"`                     // Maximum size of non heap.
	MaxNonHeapBytes                int64                  `json:"maxNonHeapBytes"`                // The maximum number of bytes that the JVM can use for non-heap purposes
	NonHeapUtilization             string                 `json:"nonHeapUtilization"`             // Utilization of non heap.
	TotalHeap                      string                 `json:"totalHeap"`                      // Total size of heap.
	TotalHeapBytes                 int64                  `json:"totalHeapBytes"`                 // The total number of bytes that are available for the JVM heap to use
	UsedHeap                       string                 `json:"usedHeap"`                       // Amount of used heap.
	UsedHeapBytes                  int64                  `json:"usedHeapBytes"`                  // The number of bytes of JVM heap that are currently being used
	FreeHeap                       string                 `json:"freeHeap"`                       // Amount of free heap.
	FreeHeapBytes                  int64                  `json:"freeHeapBytes"`                  // The number of bytes that are allocated to the JVM heap but not currently being used
	MaxHeap                        string                 `json:"maxHeap"`                        // Maximum size of heap.
	MaxHeapBytes                   int64                  `json:"maxHeapBytes"`                   // The maximum number of bytes that can be used by the JVM
	HeapUtilization                string                 `json:"heapUtilization"`                // Utilization of heap.
	AvailableProcessors            int32                  `json:"availableProcessors"`            // Number of available processors if supported by the underlying system.
	ProcessorLoadAverage           float64                `json:"processorLoadAverage"`           // The processor load average if supported by the underlying system.
	TotalThreads                   int32                  `json:"totalThreads"`                   // Total number of threads.
	DaemonThreads                  int32                  `json:"daemonThreads"`                  // Number of daemon threads.
	FlowFileRepositoryStorageUsage StorageUsageDTO        `json:"flowFileRepositoryStorageUsage"` // The flowfile repository storage usage.
	ContentRepositoryStorageUsage  []StorageUsageDTO      `json:"contentRepositoryStorageUsage"`  // The content repository storage usage.
	GarbageCollection              []GarbageCollectionDTO `json:"garbageCollection"`              // The garbage collection details.
	StatsLastRefreshed             string                 `json:"statsLastRefreshed"`             // When the diagnostics were generated.
}
