package model

type ProvenanceOptionsDTO struct {
	SearchableFields []ProvenanceSearchableFieldDTO `json:"searchableFields"` // The available searchable field for the NiFi.
}
