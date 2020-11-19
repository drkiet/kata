package main

import (
	"encoding/json"
	"fmt"
)

var bundle Bundle

type Bundle struct {
	Type string
	ID string
	Objects []json.RawMessage
}

func init() {
	fmt.Println("initializing ...")
}

func unmarshal(data []byte) {
	bundle = Bundle{}
	json.Unmarshal(data, &bundle)
	fmt.Println(bundle.Type)
	fmt.Println(bundle.ID)
	for _, obj := range bundle.Objects {
		var stixObject STIXObject
		json.Unmarshal(obj, &stixObject)
		fmt.Println("\nContent")
		printStixObject(stixObject)
		switch {
		case "threat-actor" == stixObject.Type:
			var ta ThreatActor
			json.Unmarshal(obj, &ta)
			ta.CommonProperties = stixObject
			printThreatActor(ta)
		}
		//if stix.Type == "threat-actor" {
		//	var ta threatActor
		//	json.Unmarshal(obj, &ta)
		//	fmt.Println("threa-actor: ", ta)
		//}
	}
}

