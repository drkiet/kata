package main

import (
	"fmt"
	"time"
)

type ThreatActor struct {
	CommonProperties STIXObject
	Name string `json:"name" binding:"required"`
	Description string
	ThreatActorTypes [] string `json:"threat_actor_types"`
	Aliases [] string `json:"aliases"`
	FirstSeen time.Time `json:"first_seen"`
	LastSeen time.Time `json:"last_seen"`
	Roles [] string
	Goals [] string
	Sophistication string
	ResourceLevel string `json:"resource_level"`
	PrimaryMotivation string `json:"primary_motivation"`
	SecondaryMotivations string `json:"secondary_motivations"`
	PersonalMotivations string `json:"personal_motivations"`
}

func printThreatActor(ta ThreatActor) {
	fmt.Printf("\tname: %v\n", ta.Name)
	if len(ta.Description) > 0 {
		fmt.Printf("\tdescription: %v\n", ta.Description)
	}

	if ta.ThreatActorTypes != nil {
		fmt.Printf("\tthreat_actor_types: %v\n", ta.ThreatActorTypes)
	}
	if ta.Aliases != nil {
		fmt.Printf("\talias: %v\n", ta.Aliases)
	}

	if ta.FirstSeen.Year() != 1 {
		fmt.Printf("\tfirst_seen: %v\n", ta.FirstSeen)
	}
	if ta.LastSeen.Year() != 1 {
		fmt.Printf("\tlast_seen: %v\n", ta.LastSeen)
	}
	if ta.Roles != nil {
		fmt.Printf("\troles: %v\n", ta.Roles)
	}
	if ta.Goals != nil {
		fmt.Printf("\tgoals: %v\n", ta.Goals)
	}
	if len(ta.Sophistication) > 0 {
		fmt.Printf("\tsophistication: %v\n", ta.Sophistication)
	}
	if len(ta.ResourceLevel) > 0 {
		fmt.Printf("\tresource_level: %v\n", ta.ResourceLevel)
	}
	if len(ta.PrimaryMotivation) > 0 {
		fmt.Printf("\tprimary_motivation: %v\n", ta.PrimaryMotivation)
	}
	if len(ta.SecondaryMotivations) > 0 {
		fmt.Printf("\tsecondary_motivations: %v\n", ta.SecondaryMotivations)
	}
	if len(ta.PersonalMotivations) > 0 {
		fmt.Printf("\tpersonal_motivations: %v\n", ta.PersonalMotivations)
	}
}