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

func marshal(bundle Bundle) (jsonData string){
	data, e := json.MarshalIndent(bundle, "", "  ")
	check(e)
	return string(data)
}

func unmarshal(data []byte) (bundle Bundle) {
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
			ta := unmarshalThreatActor(obj)
			ta.CommonProperties = stixObject
			printThreatActor(ta)
		case IdentityType:
			identity := unmarshalIdentity(obj)
			identity.CommonProperties = stixObject
			printIdentity(identity)
		case RelationshipType:
			relationship := unmarshalRelationship(obj)
			relationship.CommonProperties = stixObject
			printRelationship(relationship)
		}
	}
	return bundle
}

