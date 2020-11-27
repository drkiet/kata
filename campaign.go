package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Campaign struct {
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

	Name string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	Aliases [] string `json:"aliases,omitempty"`
	FirstSeen time.Time `json:"first_seen,omitempty"`
	LastSeen time.Time `json:"last_seen,omitempty"`
	Objective string `json:"objective,omitempty"`
}


func unmarshalCampaign(obj json.RawMessage) (campaign Campaign) {
	json.Unmarshal(obj, &campaign)
	return campaign
}

func marshalCampaign(campaign Campaign) (jsonData string){
	data, e := json.MarshalIndent(campaign, "", "  ")
	check(e)
	return string(data)
}

func printCampaign(campaign Campaign) {
	fmt.Println("Campaign:\n", marshalCampaign(campaign))
}