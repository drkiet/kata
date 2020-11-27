package kata

import (
	"encoding/json"
	"fmt"
)

type EmailAddress struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Value string `json:"value" binding:"required"`
	DisplayName string `json:"display_name,omitempty"`
	BelongsToRef string `json:"belongs_to_ref,omitempty"`
}

func printEmailAddress(emailAddress EmailAddress) {
	fmt.Println("Email Address:\n", marshalEmailAddress(emailAddress))
}

func marshalEmailAddress(emailAddress EmailAddress) (jsonData string){
	data, e := json.MarshalIndent(emailAddress, "", "  ")
	check(e)
	return string(data)
}

func unmarshalEmailAddress(obj json.RawMessage) (emailAddress EmailAddress) {
	json.Unmarshal(obj, &emailAddress)
	return emailAddress
}
