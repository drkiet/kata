package main

import (
	"encoding/json"
	"fmt"
)

type Identity struct {
	CommonProperties STIXObject
	Name string `json:"name" binding:"required"`
	Description string
	Roles [] string
	IdentityClass OpenVocab `json:"identity_class"`
	Sectors [] OpenVocab
	ContactInformation string `json:"contact_information"`
}

func printIdentity(id Identity) {
	fmt.Printf("\tname: %v\n", id.Name)
	if len(id.Description) > 0 {
		fmt.Printf("\tdescription: %v\n", id.Description)
	}
	if id.Roles != nil {
		fmt.Printf("\troles: %v\n", id.Roles)
	}
	if len(id.IdentityClass) > 0 {
		fmt.Printf("\tidentity_class: %v\n", id.IdentityClass)
	}
	if id.Sectors != nil {
		fmt.Printf("\tsectors: %v\n", id.Sectors)
	}
	if len(id.ContactInformation) > 0 {
		fmt.Printf("\tcontact_information: %vn", id.ContactInformation)
	}
}

func unmarshalIdentity(obj json.RawMessage) (identity Identity) {
	json.Unmarshal(obj, &identity)
	return identity
}
