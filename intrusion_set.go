package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type IntrusionSet struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version", binding:"required"`
	ID string `json:"id" binding:"required"`
	CreatedByRef string `json:"created-by-ref,omitempty"`
	Created time.Time `json:"created" binding="required"`
	Modified time.Time `json:"modified" binding="required"`
	Revoked bool `json:"revoked,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Confidence int `json:"confidence,omitempty"`
	Lang string `json:"lang,omitempty"`
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`

	Name string `json:"name" binding:"required"`
	Description string `json:"description, omitempty"`
	Aliases [] string `json:"aliases,omitempty"`
	FirstSeen time.Time `json:"first_seen,omitempty"`
	Goals []string `json:"goals, omitempty"`
	ResourceLevel OpenVocab `json:"resource_level,omitempty"`
	PrimaryMotivation OpenVocab `json:"primary_motivation,omitempty"`
	SecondaryMotivation OpenVocab `json:"secondary_motivation,omitempty"`
}

func unmarshalIntrusionSet(obj json.RawMessage) (intrusionSet IntrusionSet) {
	json.Unmarshal(obj, &intrusionSet)
	return intrusionSet
}

func marshalIntrusionSet(intrusionSet IntrusionSet) (jsonData string){
	data, e := json.MarshalIndent(intrusionSet, "", "  ")
	check(e)
	return string(data)
}

func printIntrusionSet(intrusionSet IntrusionSet) {
	fmt.Println("Intrusion Set:\n", marshalIntrusionSet(intrusionSet))
}