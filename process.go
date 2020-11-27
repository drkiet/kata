package kata

import (
	"encoding/json"
	"fmt"
	"time"
)

type Process struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	IsHidden bool `json:"is_hidden,omitempty"`
	CreatedTime time.Time `json:"created_time,omitempty"`
	Cwd string `json:"cwd,omitempty"`
	CommandLine string `json:"command_line,omitempty"`
	EnvironmentVariables map[string]interface{} `json:"environment_variables,omitempty"`
	OpenedConnectionRefs []string `json:"opened_connection_refs,omitempty"`
	CreatorUserRef string `json:"created_user_ref,omitempty"`
	ImageRef string `json:"image_ref,omitempty"`
	ParentRef string `json:"parent_ref,omitempty"`
	ChildRef string `json:"child_ref,omitempty"`
}

func printProcess(process Process) {
	fmt.Println("Process:\n", marshalProcess(process))
}

func marshalProcess(process Process) (jsonData string){
	data, e := json.MarshalIndent(process, "", "  ")
	check(e)
	return string(data)
}

func unmarshalProcess(obj json.RawMessage) (process Process) {
	json.Unmarshal(obj, &process)
	return process
}
