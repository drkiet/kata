package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Opinion struct {
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

	Explanation string `json:"explanation,omitempty"`
	Authors []string `json:"authors,omitempty"`
	Opinion string `json:"opinion" binding:"required"`
	ObjectRefs []string `json:"object_refs" binding:"required"`
}

func printOpinion(opinion Opinion) {
	fmt.Println("Opinion:\n", marshalOpinion(opinion))
}

func marshalOpinion(opinion Opinion) (jsonData string){
	data, e := json.MarshalIndent(opinion, "", "  ")
	check(e)
	return string(data)
}

func unmarshalOpinion(obj json.RawMessage) (opinion Opinion) {
	json.Unmarshal(obj, &opinion)
	return opinion
}

