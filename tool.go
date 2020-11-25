package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Tool struct {
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
	GranularMarkings [] GranularMarking `json:"granular_markings,omitempty"`

	Name string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	ToolTypes []OpenVocab `json:"tool_types" binding:"required"`
	Aliases []string `json:"aliases,omitempty"`
	KillChainPhases []KillChainPhase `json:"kill_chain_phases,omitempty"`
	ToolVersion string `json:"tool_version,omitempty"`
}

func unmarshalTool(obj json.RawMessage) (tool Tool) {
	json.Unmarshal(obj, &tool)
	return tool
}

func marshalTool(tool Tool) (jsonData string){
	data, e := json.MarshalIndent(tool, "", "  ")
	check(e)
	return string(data)
}

func printTool(tool Tool) {
	fmt.Println("Tool:\n", marshalTool(tool))
}