package main

import (
	"encoding/json"
	"fmt"
)

var bundle Bundle

type Bundle struct {
	Type string `json:"type" binding:"required"`
	ID string `json:"id" binding:"required"`
	Objects []json.RawMessage `json:"objects"`
}

func init() {
	fmt.Println("initializing ...")
}

func printBundle(bundle Bundle) {
	fmt.Printf("\ttype: %v\n", bundle.Type)
	fmt.Printf("\tid: %v\n", bundle.ID)
}

func unmarshal(data []byte) {
	bundle = Bundle{}
	e := json.Unmarshal(data, &bundle)
	fmt.Printf("e: %v\n", e)
	check(e)

	printBundle(bundle)

	for _, obj := range bundle.Objects {
		var stixObject STIXObject
		json.Unmarshal(obj, &stixObject)
		printStixObject(stixObject)
		switch stixObject.Type{
		case ThreatActorType:
			var ta ThreatActor
			json.Unmarshal(obj, &ta)
			ta.CommonProperties = stixObject
			printThreatActor(ta)
		case IdentityType:
			var identity Identity
			json.Unmarshal(obj, &identity)
			identity.CommonProperties = stixObject
			printIdentity(identity)
		case RelationshipType:
			var relationship Relationship
			json.Unmarshal(obj, &relationship)
			relationship.CommonProperties = stixObject
			printRelationship(relationship)
		}

	}
}

