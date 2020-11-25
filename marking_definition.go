package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type MarkingDefinition struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version" binding:"required"`
	ID string `json:"id" binding:"required"`
	CreatedByRef string `json:"created-by-ref,omitempty"`
	Created time.Time `json:"created" binding:"required"`
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`

	Name string `json:"name,omitempty"`
	DefinitionType OpenVocab `json:"definition_type" binding:"required"`
	Definition interface{} `json:"definition" binding:"required"`
}
func unmarshalMarkingDefinition(obj json.RawMessage) (markingDefinition MarkingDefinition) {
	json.Unmarshal(obj, &markingDefinition)
	return markingDefinition
}

func marshalMarkingDefinition(markingDefinition MarkingDefinition) (jsonData string){
	data, e := json.MarshalIndent(markingDefinition, "", "  ")
	check(e)
	return string(data)
}

func printMarkingDefinition(markingDefinition MarkingDefinition) {
	fmt.Println("Marking Definition:\n", marshalMarkingDefinition(markingDefinition))
}