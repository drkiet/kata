package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"
	"github.com/google/uuid"
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

type StixType string
const (
	ThreatActorType = "threat-actor"
	IdentityType = "identity"
	RelationshipType = "relationship"
)


type STIXObject struct {
	Type        string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version"`
	ID          string `json:"id"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
}

func printStixObject(stixObject STIXObject) {
	fmt.Println("\nSTIX Object: ...")
	fmt.Printf("type: %v\n", stixObject.Type)
	fmt.Printf("spec_version: %v\n", stixObject.SpecVersion)
	fmt.Printf("id: %v\n", stixObject.ID)
	fmt.Printf("created: %v\n", stixObject.Created)
	fmt.Printf("modified: %v\n", stixObject.Modified)
}

type ExternalReference struct {
	SourceName  string
	Description string
	Url         url.URL
	Hashes      map[string]interface{}
	ExternalID  string
}

type KillChainPhase struct {
	KillChainName string
	PhaseName     string
}

// New type

type OpenVocab string

type Identifier uuid.UUID

func getUUID() {
    uuidWithHyphen := uuid.New()
    fmt.Println(uuidWithHyphen)
    uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
    fmt.Println(uuid)
}