package kata

import (
	"encoding/json"
	"fmt"
)

type AutonomousSystem struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Number int `json:"number" binding:"required"`
	Name string `json:"name,omitempty"`
	Rir string `json:"rir,omitempty"`
}

func printAutonomousSystem(autonomousSystem AutonomousSystem) {
	fmt.Println("Autonomous System:\n", marshalAutonomousSystem(autonomousSystem))
}

func marshalAutonomousSystem(autonomousSystem AutonomousSystem) (jsonData string){
	data, e := json.MarshalIndent(autonomousSystem, "", "  ")
	check(e)
	return string(data)
}

func unmarshalAutonomousSystem(obj json.RawMessage) (autonomousSystem AutonomousSystem) {
	json.Unmarshal(obj, &autonomousSystem)
	return autonomousSystem
}
