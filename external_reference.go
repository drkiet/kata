package kata

import (
	"encoding/json"
	"fmt"
)

type ExternalReference struct {
	SourceName string `json:"source_name" binding:"required"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	Hashes      map[string]interface{} `json:"hashes,omitempty"`
	ExternalID  string `json:"external_id,omitempty"`
}

func unmarshalExternalReference(obj json.RawMessage) (externalReference ExternalReference) {
	json.Unmarshal(obj, &externalReference)
	return externalReference
}

func marshalExternalReference(externalReference ExternalReference) (jsonData string){
	data, e := json.MarshalIndent(externalReference, "", "  ")
	check(e)
	return string(data)
}

func printExternalReference(er ExternalReference) {
	fmt.Println("External Reference:\n", marshalExternalReference(er))
}