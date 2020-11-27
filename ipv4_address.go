package main

import (
	"encoding/json"
	"fmt"
)

type IpV4Address struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Value string `json:"value" binding:"required"`
	ResolvesToRefs []string `json:"resolves_to_refs,omitempty"` // deprecated
	BelongsToRefs []string `json:"belongs_to_refs,omitempty"` // deprecated
}

func printIpV4Address(ipV4Address IpV4Address) {
	fmt.Println("IpV4 Address:\n", marshalIpV4Address(ipV4Address))
}

func marshalIpV4Address(ipV4Address IpV4Address) (jsonData string){
	data, e := json.MarshalIndent(ipV4Address, "", "  ")
	check(e)
	return string(data)
}

func unmarshalIpV4Address(obj json.RawMessage) (ipV4Address IpV4Address) {
	json.Unmarshal(obj, &ipV4Address)
	return ipV4Address
}


