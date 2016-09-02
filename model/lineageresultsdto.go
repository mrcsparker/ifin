package model

type LineageResultsDTO struct {
	Errors []string            `json:"errors"` // Any errors that occurred while generating the lineage.
	Nodes  []ProvenanceNodeDTO `json:"nodes"`  // The nodes in the lineage.
	Links  []ProvenanceLinkDTO `json:"links"`  // The links between the nodes in the lineage.
}
