package main

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

// STIX Objects:
//   STIX Core Objects:
//     STIX Domain Objects (SDO) -- total 18
//       Attack Pattern, Campaign, Course of Action, Grouping, Identity, Indicator, Infrastructure,
//       Intrusion Set, Location, Malware, Malware Analysis, Note, Observed Data, Opinion,
//       Report, Threat Actor, Tool, Vulnerability,
//     STIX Cyber-observable Object (SCO) -- total 18
//       Artifact, AS, Directory, Domain Name, Email Address, Email Message, File,
//       IPv4 Address, IPv6 Address, MAC Address, Mutex, Network Traffic, Process, Software,
//       URL, User Account, Windows Registry Key, X509 Certificate,
//     STIX Relationship Object (SRO) -- total 2
//       Relationship, Sighting
//   STIX Meta Objects:
//     STIX Language Content Objects
//     STIX Data Markings Objects
//     STIX Bundle
// STIX Bundle:
//   One or more STIX Objects

type StixType string

// SDO
const CampaignType = "campaign"
const CourseOfActionType = "course-of-action"
const GroupingType = "grouping"
const IndicatorType = "indicator"
const InfrastructureType = "infrastructure"
const ThreatActorType = "threat-actor"
const IdentityType = "identity"
const AttackPatternType = "attack-pattern"
const IntrusionSetType = "intrusion-set"
const LocationType = "location"
const MalwareType = "malware"
const MalwareAnalysisType = "malware-analysis"
const NoteType = "note"
const ObservedDataType = "observed-data"
const OpinionType = "opinion"
const ReportType = "report"
const ToolType = "tool"
const VulnerabilityType = "vulnerability"

// SCO
const ArtifactType = "artifact"
const ASType = "as"
const DirectoryType = "directory"
const DomainNameType = "domain-name"
const EmailAddressType = "email-address"
const EmailMessageType = "email-message"
const FileType = "file"
const IPv4AddressType = "ipv4-address"
const IPv6AddressType = "ipv6-address"
const MACTAddressType = "mac-address"
const MutexType = "mutex"
const NetworkTrafficType = "network-traffic"
const ProcessType = "process"
const SoftwareType = "software"
const URLType = "url"
const UserAccountType = "user-account"
const WindowsRegistryKeyType = "windows-registry-key"
const X509CertificateType = "x509-certificate"

// SRO
const RelationshipType = "relationship"
const SightingType = "sighting"

// Meta
const BundleType = "bundle"
const LanguageContentType = "language-content"
const MarkingDefinitionType = "marking-definition"

type STIXObject struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version"`
	ID string `json:"id" binding:"required"`
	CreatedByRef string `json:"created-by-ref"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Revoked bool `json:"revoked"`
	Labels []string `json:"labels"`
	Confidence int `json:"confidence"`
	Lang string `json:"lang"`
	ExternalReferences []ExternalReference `json:"external_references"`
	ObjectMarkingRefs []string `json:"object_marking_refs"`
	GranularMarkings [] GranularMarking `json:"granular-markings"`
	Defanged bool `json:"defanged"`
	Extensions map[string]interface{} `json:"extensions"`
}

type GranularMarking struct {
	Lang string `json:"lang"`
	MarkingRef string `json:"marking-ref, omitempty"`
	Selectors []string `json:"selectors" binding:"required"`
}

type KillChainPhase struct {
	KillChainName string `json:"kill_chain_name"`
	PhaseName     string `json:"phase_name"`
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

