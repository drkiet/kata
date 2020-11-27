package kata

import (
	"encoding/json"
	"fmt"
)

type MacAddress struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Value string `json:"value" binding:"required"`
}

func printMacAddress(macAddress MacAddress) {
	fmt.Println("Mac Address:\n", marshalMacAddress(macAddress))
}

func marshalMacAddress(macAddress MacAddress) (jsonData string){
	data, e := json.MarshalIndent(macAddress, "", "  ")
	check(e)
	return string(data)
}

func unmarshalMacAddress(obj json.RawMessage) (macAddress MacAddress) {
	json.Unmarshal(obj, &macAddress)
	return macAddress
}


