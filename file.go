package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type File struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Hashes map[string]interface{} `json:"hashes,omitempty"`
	Size int `json:"size,omitempty"`
	Name string `json:"name,omitempty"`
	NameEnc string `json:"name_enc,omitempty"`
	MagicNumberHex string `json:"magic_number_hex,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	Ctime time.Time `json:"ctime,omitempty"`
	Mtime time.Time `json:"mtime,omitempty"`
	Atime time.Time `json:"atime,omitempty"`
	ParentDirectoryRef string `json:"parent_directory_ref,omitempty"`
	ContainsRefs []string `json:"contains_refs,omitempty"`
	ContentRef string `json:"content_ref,omitempty"`
}

func printFile(file File) {
	fmt.Println("File:\n", marshalFile(file))
}

func marshalFile(file File) (jsonData string){
	data, e := json.MarshalIndent(file, "", "  ")
	check(e)
	return string(data)
}

func unmarshalFile(obj json.RawMessage) (file File) {
	json.Unmarshal(obj, &file)
	return file
}
