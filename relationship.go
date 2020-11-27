package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Relationship struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version" binding:"required"`
	ID string `json:"id" binding:"required"`
	CreatedByRef string `json:"created-by-ref,omitempty"`
	Created time.Time `json:"created" binding:"required"`
	Modified time.Time `json:"modified" binding:"required"`
	Revoked bool `json:"revoked,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Confidence int `json:"confidence,omitempty"`
	Lang string `json:"lang,omitempty"`
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`

	RelationshipType string `json:"relationship_type" binding:"required"`
	Description string `json:"description,omitempty"`
	SourceRef string `json:"source_ref" binding:"required"`
	TargetRef string `json:"target_ref" binding:"required"`
	StartTime time.Time `json:"start_time,omitempty"`
	StopTime time.Time `json:"stop_time,omitempty"`
}

func unmarshalRelationship(obj json.RawMessage) (relationship Relationship) {
	json.Unmarshal(obj, &relationship)
	return relationship
}

func marshalRelationship(relationship Relationship) (jsonData string){
	data, e := json.MarshalIndent(relationship, "", "  ")
	check(e)
	return string(data)
}

func printRelationship(relationship Relationship) {
	fmt.Println("Relationship:\n", marshalRelationship(relationship))
}