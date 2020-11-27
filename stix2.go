package kata

import (
	"encoding/json"
)


type stixObject struct {
	Type string
	ID string
	SpecVersion string
	Created string
	Modified string
	Name string
	Description string
}

type threatActor struct {
	//StixObject stixObject
	Type string
	ID string
	SpecVersion string
	Created string
	Modified string
	Name string
	Description string
	ThreatActorType []string
	Aliases []string
	Roles []string
	Goals []string
	Sophistication string
	SourceLevel string
	PrimaryMotivation string
}

type stix2 struct {
	Type string
	ID string
	Objects []json.RawMessage
}

//func main() {
//	dat, err := ioutil.ReadFile("./identify_threat_actor.json")
//	check(err)
//	//fmt.Print(string(dat))
//
//	res := stix2{}
//	json.Unmarshal(dat, &res)
//	fmt.Println(res.Type)
//	fmt.Println(res.ID)
//	for _, obj := range res.Objects {
//		var stix stixObject
//		json.Unmarshal(obj, &stix)
//		fmt.Println(stix)
//		if stix.Type == "threat-actor" {
//			var ta threatActor
//			json.Unmarshal(obj, &ta)
//			fmt.Println("threa-actor: ", ta)
//		}
//	}
//}