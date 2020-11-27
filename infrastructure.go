package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Infrastructure struct {
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

	Name string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	InfrastructureTypes []OpenVocab `json:"infrastructure_types" binding:"required"`
	Aliases []string `json:"aliases,omitempty"`
	KillChainPhases []KillChainPhase `json:"kill_chain_phases,omitempty"`
	FirstSeen time.Time `json:"first_seen,omitempty"`
	LastSeen time.Time `json:"last_seen,omitempty"`
}


func printInfrastructure(infrastructure Infrastructure) {
	fmt.Println("Infrastructure:\n", marshalInfrastructure(infrastructure))
}

func marshalInfrastructure(infrastructure Infrastructure) (jsonData string){
	data, e := json.MarshalIndent(infrastructure, "", "  ")
	check(e)
	return string(data)
}

func unmarshalInfrastructure(obj json.RawMessage) (infrastructure Infrastructure) {
	json.Unmarshal(obj, &infrastructure)
	return infrastructure
}

