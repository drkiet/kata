package main

import (
	"encoding/json"
	"fmt"
)

type Url struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Name string `json:"named" binding:"required"`
}

func printUrl(url Url) {
	fmt.Println("Url:\n", marshalUrl(url))
}

func marshalUrl(url Url) (jsonData string){
	data, e := json.MarshalIndent(url, "", "  ")
	check(e)
	return string(data)
}

func unmarshalUrl(obj json.RawMessage) (url Url) {
	json.Unmarshal(obj, &url)
	return url
}
