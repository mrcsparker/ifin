package model

type AccessConfigurationDTO struct {
	SupportsLogin bool `json:"supportsLogin"` // [Read Only] Indicates whether or not this NiFi supports user login.
}
