package model

type AccessConfiguration struct {
	SupportsLogin     bool `json:"supportsLogin"`     // Indicates whether or not this NiFi supports user login.
	SupportsAnonymous bool `json:"supportsAnonymous"` // Indicates whether or not this NiFi supports anonymous.
}
