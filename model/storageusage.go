package model

type StorageUsage struct {
	Identifier      string `json:"identifier"`      // The identifier of this storage location. The identifier will correspond to the identifier keyed in the storage configuration.
	FreeSpace       string `json:"freeSpace"`       // Amount of free space.
	TotalSpace      string `json:"totalSpace"`      // Amount of total space.
	UsedSpace       string `json:"usedSpace"`       // Amount of used space.
	FreeSpaceBytes  int64  `json:"freeSpaceBytes"`  // The number of bytes of free space.
	TotalSpaceBytes int64  `json:"totalSpaceBytes"` // The number of bytes of total space.
	UsedSpaceBytes  int64  `json:"usedSpaceBytes"`  // The number of bytes of used space.
	Utilization     string `json:"utilization"`     // Utilization of this storage location.
}
