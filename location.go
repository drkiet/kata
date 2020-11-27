package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Location struct {
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
	Latitude float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	Precision float32 `json:"precision,omitempty"`
	Region OpenVocab `json:"region,omitempty"`
	Country string `json:"country,omitempty"`
	AdministrativeArea string `json:"administrative_area,omitempty"`
	City string `json:"city,omitempty"`
	StreetAddress string `json:"street_address,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

func printLocation(location Location) {
	fmt.Println("Infrastructure:\n", marshalLocation(location))
}

func marshalLocation(location Location) (jsonData string){
	data, e := json.MarshalIndent(location, "", "  ")
	check(e)
	return string(data)
}

func unmarshalLocation(obj json.RawMessage) (location Location) {
	json.Unmarshal(obj, &location)
	return location
}

