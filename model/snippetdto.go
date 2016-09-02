package model

type SnippetDTO struct {
	Id            string `json:"id"`            // The id of the snippet.
	Uri           string `json:"uri"`           // The URI of the snippet.
	ParentGroupId string `json:"parentGroupId"` // The group id for the components in the snippet.

}
