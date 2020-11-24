package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Identity struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version", binding:"required"`
	ID string `json:"id" binding:"required"`
	CreatedByRef string `json:"created-by-ref,omitempty"`
	Created time.Time `json:"created" binding="required"`
	Modified time.Time `json:"modified" binding="required"`
	Revoked bool `json:"revoked,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Confidence int `json:"confidence,omitempty"`
	Lang string `json:"lang,omitempty"`
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`

	Name string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	Roles [] string `json:"roles,omitempty"`
	IdentityClass OpenVocab `json:"identity_class" binding:"required"`
	Sectors [] OpenVocab `json:"sectors,omitempty"`
	ContactInformation string `json:"contact_information,omitempty"`
}

func printIdentity(id Identity) {
	fmt.Println("Identity:\n", marshalIdentity(id))
}

func marshalIdentity(identity Identity) (jsonData string){
	data, e := json.MarshalIndent(identity, "", "  ")
	check(e)
	return string(data)
}

func unmarshalIdentity(obj json.RawMessage) (identity Identity) {
	json.Unmarshal(obj, &identity)
	return identity
}
