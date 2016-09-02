package model

type PropertyDescriptorDTO struct {
	Name                        string                 `json:"name"`                        // The name for the property.
	DisplayName                 string                 `json:"displayName"`                 // The human readable name for the property.
	Description                 string                 `json:"description"`                 // The descriptoin for the property. Used to relay additional details to a user or provide a mechanism of documenting intent.
	DefaultValue                string                 `json:"defaultValue"`                // The default value for the property.
	AllowableValues             []AllowableValueEntity `json:"allowableValues"`             // Allowable values for the property. If empty then the allowed values are not constrained.
	Required                    bool                   `json:"required"`                    // Whether the property is required.
	Sensitive                   bool                   `json:"sensitive"`                   // Whether the property is sensitive and protected whenever stored or represented.
	Dynamic                     bool                   `json:"dynamic"`                     // Whether the property is dynamic (user-defined).
	SupportsEl                  bool                   `json:"supportsEl"`                  // Whether the property supports expression language.
	IdentifiesControllerService string                 `json:"identifiesControllerService"` // If the property identifies a controller service, this returns the fully qualified type.
}
