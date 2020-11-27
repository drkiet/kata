package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Directory struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Path string `json:"path" binding:"required"`
	PathEnc string `json:"path_enc,omitempty"`
	Ctime time.Time `json:"ctime,omitempty"`
	Mtime time.Time `json:"mtime,omitempty"`
	Atime time.Time `json:"atime,omitempty"`
	ContainsRefs []string `json:"contains_refs,omitempty"`
}

func printDirectory(directory Directory) {
	fmt.Println("Directory:\n", marshalDirectory(directory))
}

func marshalDirectory(directory Directory) (jsonData string){
	data, e := json.MarshalIndent(directory, "", "  ")
	check(e)
	return string(data)
}

func unmarshalDirectory(obj json.RawMessage) (directory Directory) {
	json.Unmarshal(obj, &directory)
	return directory
}
