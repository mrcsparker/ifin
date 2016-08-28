package model

type Label struct {
	Id            string      `json:"id"`            // The id of the component.
	Uri           string      `json:"uri"`           // The URI for futures requests to the component.
	Position      Position    `json:"position"`      // The position of this component in the UI if applicable.
	ParentGroupId string      `json:"parentGroupId"` // The id of parent process group of this component if applicable.
	Label         string      `json:"label"`         // The text that appears in the label.
	Width         float64     `json:"width"`         // The width of the label in pixels when at a 1:1 scale.
	Height        float64     `json:"height"`        // The height of the label in pixels when at a 1:1 scale.
	Style         interface{} `json:"style"`         // The styles for this label (font-size => 12px, background-color => #eee, etc).
}
