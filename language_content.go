package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type LanguageContent struct {
	SpecVersion string `json:"spec_version" binding:"required"`
	ID string `json:"id" binding:"required"`
	CreatedByRef string `json:"created-by-ref,omitempty"`
	Created time.Time `json:"created" binding:"required"`
	Modified time.Time `json:"modified" binding:"required"`
	Revoked bool `json:"revoked,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Confidence int `json:"confidence,omitempty"`
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`

	ObjectRef string `json:"object-ref" binding:"required"`
	ObjectModified time.Time `json:"object_modified,omitempty"`
	Contents map[string]interface{} `json:"contents" binding:"required"`
}


func unmarshalLanguageContent(obj json.RawMessage) (languageContent LanguageContent) {
	json.Unmarshal(obj, &languageContent)
	return languageContent
}

func marshalLanguageContent(languageContent LanguageContent) (jsonData string){
	data, e := json.MarshalIndent(languageContent, "", "  ")
	check(e)
	return string(data)
}

func printLanguageContent(languageContent LanguageContent) {
	fmt.Println("Marking Definition:\n", marshalLanguageContent(languageContent))
}
