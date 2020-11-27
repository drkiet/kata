package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type ThreatActor struct {
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

	Name string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	ThreatActorTypes [] string `json:"threat_actor_types" binding:"required"`
	Aliases [] string `json:"aliases,omitempty"`
	FirstSeen time.Time `json:"first_seen,omitempty"`
	LastSeen time.Time `json:"last_seen,omitempty"`
	Roles [] string `json:"roles,omitempty"`
	Goals [] string `json:"goals,omitempty"`
	Sophistication string `json:"sophistication,omitempty"`
	ResourceLevel string `json:"resource_level,omitempty"`
	PrimaryMotivation string `json:"primary_motivation,omitempty"`
	SecondaryMotivations string `json:"secondary_motivations,omitempty"`
	PersonalMotivations string `json:"personal_motivations,omitempty"`
}

func printThreatActor(ta ThreatActor) {
	fmt.Println("Threat Actor:\n", marshalThreatActor(ta))
}

func marshalThreatActor(threatActor ThreatActor) (jsonData string){
	data, e := json.MarshalIndent(threatActor, "", "  ")
	check(e)
	return string(data)
}

func unmarshalThreatActor(obj json.RawMessage) (ta ThreatActor) {
	json.Unmarshal(obj, &ta)
	return ta
}