package model

type DimensionsDTO struct {
	Width  float64 `json:"width"`  // The width of the label in pixels when at a 1:1 scale.
	Height float64 `json:"height"` // The height of the label in pixels when at a 1:1 scale.
}
