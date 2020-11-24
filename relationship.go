package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Relationship struct {
	CommonProperties STIXObject
	RelationshipType string `json:"relationship_type" binding:"required"`
	Description string
	SourceRef string `json:"source_ref" binding:"required"`
	TargetRef string `json:"target_ref" binding:"required"`
	StartTime time.Time `json:"start_time"`
	StopTime time.Time `json:"stop_time"`
}


func printRelationship(rel Relationship) {
	fmt.Printf("\trelationship_type: %v\n", rel.RelationshipType)

	if len(rel.Description) > 0 {
		fmt.Printf("\tdescription: %v\n", rel.Description)
	}
	fmt.Printf("\tsource_ref: %v\n", rel.SourceRef)
	fmt.Printf("\ttarget_ref: %v\n", rel.TargetRef)
	if rel.StartTime.Year() != 1 {
		fmt.Printf("\tstart_time: %v\n", rel.StartTime)
	}
	if rel.StopTime.Year() != 1 {
		fmt.Printf("\tstop_time: %v\n", rel.StopTime)
	}
}

func unmarshalRelationship(obj json.RawMessage) (relationship Relationship) {
	json.Unmarshal(obj, &relationship)
	return relationship
}