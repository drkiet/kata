package main

import (
	"encoding/json"
	"fmt"
)

type IpV6Address struct {
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

func printIpV6Address(ipV6Address IpV6Address) {
	fmt.Println("IpV6 Address:\n", marshalIpV6Address(ipV6Address))
}

func marshalIpV6Address(ipV6Address IpV6Address) (jsonData string){
	data, e := json.MarshalIndent(ipV6Address, "", "  ")
	check(e)
	return string(data)
}

func unmarshalIpV6Address(obj json.RawMessage) (ipV6Address IpV6Address) {
	json.Unmarshal(obj, &ipV6Address)
	return ipV6Address
}



