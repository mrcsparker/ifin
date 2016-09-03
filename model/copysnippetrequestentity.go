package model

// attribute: copySnippetRequestEntity
type CopySnippetRequestEntity struct {
	SnippetId string  `json:"snippetId"` // The identifier of the snippet.
	OriginX   float64 `json:"originX"`   // The x coordinate of the origin of the bounding box where the new components will be placed.
	OriginY   float64 `json:"originY"`   // The y coordinate of the origin of the bounding box where the new components will be placed.
}
