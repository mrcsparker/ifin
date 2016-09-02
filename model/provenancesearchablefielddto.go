package model

type ProvenanceSearchableFieldDTO struct {
	Id    string `json:"id"`    // The id of the searchable field.
	Field string `json:"field"` // The searchable field.
	Label string `json:"label"` // The label for the searchable field.
	Type  string `json:"type"`  // The type of the searchable field.
}
