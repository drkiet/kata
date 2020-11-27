package kata

import (
	"encoding/json"
	"fmt"
	"time"
)

type Grouping struct {
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

	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Context string `json:"context" binding:"required"`
	ObjectRefs []string `json:"object_refs" binding:"required"`
}


func printGrouping(grouping Grouping) {
	fmt.Println("Grouping:\n", marshalGrouping(grouping))
}

func marshalGrouping(grouping Grouping) (jsonData string){
	data, e := json.MarshalIndent(grouping, "", "  ")
	check(e)
	return string(data)
}

func unmarshalGrouping(obj json.RawMessage) (grouping Grouping) {
	json.Unmarshal(obj, &grouping)
	return grouping
}

