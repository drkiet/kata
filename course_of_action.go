package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"
)

type CourseOfAction struct {
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
	ActionType OpenVocab `json:"action_type,omitempty"`
	OsExecutionEnvs []string `json:"os_execution_envs,omitempty"`
	ActionBin binary.ByteOrder `json:"action_bin,omitempty"`
	ActionReference ExternalReference `json:"action_reference,omitempty"`
}


func unmarshalCourseOfAction(obj json.RawMessage) (courseOfAction CourseOfAction) {
	json.Unmarshal(obj, &courseOfAction)
	return courseOfAction
}

func marshalCourseOfAction(courseOfAction CourseOfAction) (jsonData string){
	data, e := json.MarshalIndent(courseOfAction, "", "  ")
	check(e)
	return string(data)
}

func printCourseOfAction(courseOfAction CourseOfAction) {
	fmt.Println("Course Of Action:\n", marshalCourseOfAction(courseOfAction))
}
