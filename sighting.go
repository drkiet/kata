package kata

import (
	"encoding/json"
	"fmt"
	"time"
)

type Sighting struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version"binding:"required"`
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

	Description string `json:"description,omitempty"`
	FirstSeen time.Time `json:"first_seen,omitempty"`
	LastSeen time.Time `json:"last_seen,omitempty"`
	Count int `json:"count,omitempty"`
	SightingOfRef string `json:"sighting_of_ref" binding:"required"`
	ObservedDataRefs []string `json:"observed_data_refs,omitempty"`
	WhereSightedRef []string `json:"where_sighted_refs,omitempty"`
	Summary bool `json:"summary,omitempty"`
}


func unmarshalSighting(obj json.RawMessage) (sighting Sighting) {
	json.Unmarshal(obj, &sighting)
	return sighting
}

func marshalSighting(sighting Sighting) (jsonData string){
	data, e := json.MarshalIndent(sighting, "", "  ")
	check(e)
	return string(data)
}

func printSighting(sighting Sighting) {
	fmt.Println("Sighting:\n", marshalSighting(sighting))
}
