package model

type PermissionsDTO struct {
	CanRead  bool `json:"canRead"`  // [Read Only] Indicates whether the user can read a given resource.
	CanWrite bool `json:"canWrite"` // [Read Only] Indicates whether the user can write a given resource.
}
