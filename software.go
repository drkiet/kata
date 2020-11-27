package main

import (
	"encoding/json"
	"fmt"
)

type Software struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Name string `json:"named" binding:"required"`
	Cpe string `json:"cpe,omitempty"`
	Language []string `json:"language,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
}

func printSoftware(software Software) {
	fmt.Println("Software:\n", marshalSoftware(software))
}

func marshalSoftware(software Software) (jsonData string){
	data, e := json.MarshalIndent(software, "", "  ")
	check(e)
	return string(data)
}

func unmarshalSoftware(obj json.RawMessage) (software Software) {
	json.Unmarshal(obj, &software)
	return software
}
