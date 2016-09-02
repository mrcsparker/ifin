package model

type AboutDTO struct {
	Title            string `json:"title"`            // The title to be used on the page and in the about dialog.
	Version          string `json:"version"`          // The version of this NiFi.
	Uri              string `json:"uri"`              // The URI for the NiFi.
	ContentViewerUrl string `json:"contentViewerUrl"` // The URL for the content viewer if configured.
	Timezone         string `json:"timezone"`         // The timezone of the NiFi instance.
}
