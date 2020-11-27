package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Note struct {
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

	Abstract string `json:"abstract,omitempty"'`
	Content string `json:"content" binding:"required"`
	Authors []string `json:"authors,omitempty"`
	ObjectRefs []string `json:"object_refs" binding:"required"`
}

func printNote(note Note) {
	fmt.Println("Note:\n", marshalNote(note))
}

func marshalNote(note Note) (jsonData string){
	data, e := json.MarshalIndent(note, "", "  ")
	check(e)
	return string(data)
}

func unmarshalNote(obj json.RawMessage) (note Note) {
	json.Unmarshal(obj, &note)
	return note
}

