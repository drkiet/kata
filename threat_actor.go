package main

import (
	"fmt"
	"time"
)

type ThreatActor struct {
	CommonProperties STIXObject
	Name string
	Description string
	ThreatActorTypes [] string `json:"threat_actor_types"`
	Aliases [] string
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
	fmt.Printf("\tType: %v\n", ta.CommonProperties.Type)
	fmt.Printf("\tName: %v\n", ta.Name)
	if len(ta.Description) > 0 {
		fmt.Printf("\tDescription: %v\n", ta.Description)
	}

	if ta.ThreatActorTypes != nil {
		fmt.Printf("\tThreatActorTypes: %v\n", ta.ThreatActorTypes)
	}
	if ta.Aliases != nil {
		fmt.Printf("\tAlias: %v\n", ta.Aliases)
	}

	if ta.FirstSeen.Year() != 1 {
		fmt.Printf("\tFirstSeen: %v\n", ta.FirstSeen)
	}
	if ta.LastSeen.Year() != 1 {
		fmt.Printf("\tLastSeen: %v\n", ta.LastSeen)
	}
	if ta.Roles != nil {
		fmt.Printf("\tRoles: %v\n", ta.Roles)
	}
	if ta.Goals != nil {
		fmt.Printf("\tGoals: %v\n", ta.Goals)
	}
	if len(ta.Sophistication) > 0 {
		fmt.Printf("\tSophistication: %v\n", ta.Sophistication)
	}
	if len(ta.ResourceLevel) > 0 {
		fmt.Printf("\tResourceLevel: %v\n", ta.ResourceLevel)
	}
	if len(ta.PrimaryMotivation) > 0 {
		fmt.Printf("\tPrimaryMotivation: %v\n", ta.PrimaryMotivation)
	}
	if len(ta.SecondaryMotivations) > 0 {
		fmt.Printf("\tSecondaryMotivations: %v\n", ta.SecondaryMotivations)
	}
	if len(ta.PersonalMotivations) > 0 {
		fmt.Printf("\tPersonalMotivations: %v\n", ta.PersonalMotivations)
	}
}