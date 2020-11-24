package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type AttackPattern struct {
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
	//Defanged bool `json:"defanged"`
	//Extensions map[string]interface{} `json:"extensions"`

	Name string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	Aliases [] string `json:"aliases,omitempty"`
	KillChainPhases [] KillChainPhase `json:"kill_chain_phases,omitempty"`
}

func unmarshalAttackPattern(obj json.RawMessage) (attackPattern AttackPattern) {
	json.Unmarshal(obj, &attackPattern)
	return attackPattern
}

func marshalAttackPattern(attackPattern AttackPattern) (jsonData string){
	data, e := json.MarshalIndent(attackPattern, "", "  ")
	check(e)
	return string(data)
}

func printAttackPattern(ap AttackPattern) {
	fmt.Println("Attack Pattern:\n", marshalAttackPattern(ap))
}
