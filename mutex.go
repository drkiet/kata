package main

import (
	"encoding/json"
	"fmt"
)

type Mutex struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Name string `json:"name" binding:"required"`
}

func printMutex(mutex Mutex) {
	fmt.Println("Mutex:\n", marshalMutex(mutex))
}

func marshalMutex(mutex Mutex) (jsonData string){
	data, e := json.MarshalIndent(mutex, "", "  ")
	check(e)
	return string(data)
}

func unmarshalMutex(obj json.RawMessage) (mutex Mutex) {
	json.Unmarshal(obj, &mutex)
	return mutex
}



