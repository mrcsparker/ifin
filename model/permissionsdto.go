package model

type PermissionsDTO struct {
	CanRead  bool `json:"canRead"`  // Indicates whether the user can read a given resource.
	CanWrite bool `json:"canWrite"` // Indicates whether the user can write a given resource.
}
