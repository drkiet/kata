package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Report struct {
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
	ReportTypes []OpenVocab `json:"report_types" binding:"required"`
	Published time.Time `json:"published" binding:"required"`
	ObjectRefs []string `json:"object_refs" binding:"required"`

}

func unmarshalReport(obj json.RawMessage) (report Report) {
	json.Unmarshal(obj, &report)
	return report
}

func marshalReport(report Report) (jsonData string){
	data, e := json.MarshalIndent(report, "", "  ")
	check(e)
	return string(data)
}

func printReport(report Report) {
	fmt.Println("Report:\n", marshalReport(report))
}