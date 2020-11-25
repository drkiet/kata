package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Indicator struct {
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
	GranularMarkings [] GranularMarking `json:"granular_markings,omitempty"`

	Name string `json:"name, omitempty"`
	Description string `json:"description, omitempty"`
	IndicatorTypes []OpenVocab `json:"indicator_types" binding:"required"`
	Pattern string `json:"pattern" binding:"required"`
	PatternType OpenVocab `json:"pattern_type" binding:"required"`
	PatternVersion string `json:"pattern_version,omitempty"`
	ValidFrom time.Time `json:"valid_from" binding:"required"`
	ValidUntil time.Time `json:"valid_until,omitempty"`
	KillChainPhases []KillChainPhase `json:"kill_chain_phrases,omitempty"`
}

func unmarshalIndicator(obj json.RawMessage) (indicator Indicator) {
	json.Unmarshal(obj, &indicator)
	return indicator
}

func marshalIndicator(indicator Indicator) (jsonData string){
	data, e := json.MarshalIndent(indicator, "", "  ")
	check(e)
	return string(data)
}

func printIndicator(indicator Indicator) {
	fmt.Println("Indicator:\n", marshalIndicator(indicator))
}