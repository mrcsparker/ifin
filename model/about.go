package model

type About struct {
	Title   string `json:"title"`   // The title to be used on the page and in the about dialog.
	Version string `json:"version"` // The version of this NiFi.
}
