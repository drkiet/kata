package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type WindowsRegistryKeyValueType struct {
	Name string `json:"name,omitempty"`
	Data string `json:"data,omitempty"`
	DataType string `json:"data_type,omitempty"`
}

type WindowsRegistryKey struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

 	Key string `json:"key,omitempty"`
	Values []WindowsRegistryKeyValueType `json:"values,omitempty"`
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	CreatorUserRef time.Time `json:"creator_user_ref,omitempty"`
	NumberOfSubkeys int `json:"number_of_subkeys,omitempty"`
}

func printWindowsRegistryKey(windowsRegistryKey WindowsRegistryKey) {
	fmt.Println("Windows Registry Key:\n", marshalWindowsRegistryKey(windowsRegistryKey))
}

func marshalWindowsRegistryKey(windowsRegistryKey WindowsRegistryKey) (jsonData string){
	data, e := json.MarshalIndent(windowsRegistryKey, "", "  ")
	check(e)
	return string(data)
}

func unmarshalWindowsRegistryKey(obj json.RawMessage) (windowsRegistryKey WindowsRegistryKey) {
	json.Unmarshal(obj, &windowsRegistryKey)
	return windowsRegistryKey
}
