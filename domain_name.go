package main

import (
	"encoding/json"
	"fmt"
)

type DomainName struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Value string `json:"value" binding:"required"`
	ResolvesToRefs []string `json:"resolves_to_refs,omitempty"` // deprecated
}
func printDomainName(domainName DomainName) {
	fmt.Println("Domain Name:\n", marshalDomainName(domainName))
}

func marshalDomainName(domainName DomainName) (jsonData string){
	data, e := json.MarshalIndent(domainName, "", "  ")
	check(e)
	return string(data)
}

func unmarshalDomainName(obj json.RawMessage) (domainName DomainName) {
	json.Unmarshal(obj, &domainName)
	return domainName
}
