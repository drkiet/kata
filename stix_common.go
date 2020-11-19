package main

import (
	"fmt"
	"net/url"
	"time"
)

// STIX Objects:
//   STIX Core Objects:
//     STIX Domain Objects (SDO)
//     STIX Cyber-observable Object (SCO)
//     STIX Relationship Object (SRO)
//   STIX Meta Objects:
//     STIX Language Content Objects
//     STIX Marking Definition Objects
// STIX Bundle:
//   One or more STIX Objects

type STIXObject struct {
	Type string
	SpecVersion string
	ID string
	Created time.Time
	Modified time.Time
}

func printStixObject(stixObject STIXObject) {
	fmt.Printf("Type: %v\n", stixObject.Type)
	fmt.Printf("SpecVersion: %v\n", stixObject.SpecVersion)
	fmt.Printf("Id: %v\n", stixObject.ID)
	fmt.Printf("Created: %v\n", stixObject.Created)
	fmt.Printf("Modified: %v\n", stixObject.Modified)
}

type ExternalReference struct {
	SourceName string
	Description string
	Url url.URL
	Hashes map[string]interface{}
	ExternalID string
}

type KillChainPhase struct {
	KillChainName string
	PhaseName string
}

type OpenVocab string
