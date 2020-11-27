package kata

import (
	"encoding/json"
	"fmt"
	"time"
)

type ObservedData struct {
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

	FirstObserved time.Time `json:"first_observed" binding:"required"`
	LastObserved time.Time `json:"last_observed" binding:"required"`
	NumberObserved int `json:"number_observed" binding:"required"`
	Objects map[string]interface{} `json:"objects,depreciated"`
	ObjectRefs []string `json:"object_refs,omitempty"`
}

func unmarshalObservedData(obj json.RawMessage) (observedData ObservedData) {
	json.Unmarshal(obj, &observedData)
	return observedData
}

func marshalObservedData(observedData ObservedData) (jsonData string){
	data, e := json.MarshalIndent(observedData, "", "  ")
	check(e)
	return string(data)
}

func printObservedData(observedData ObservedData) {
	fmt.Println("Observed Data:\n", marshalObservedData(observedData))
}