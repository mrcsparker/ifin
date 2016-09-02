package model

type AccessConfigurationDTO struct {
	SupportsLogin bool `json:"supportsLogin"` // Indicates whether or not this NiFi supports user login. *Read Only*
}
