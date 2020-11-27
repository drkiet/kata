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

func marshalBundle(bundle Bundle) (jsonData string){
	data, e := json.MarshalIndent(bundle, "", "  ")
	check(e)
	return string(data)
}

func unmarshal(data []byte) (bundle Bundle) {
	bundle = Bundle{}
	e := json.Unmarshal(data, &bundle)
	fmt.Printf("e: %v\n", e)
	check(e)
	if bundle.Type != BundleType {
		fmt.Println("Unexpected bundle type ..." + bundle.Type)
	}

	for _, obj := range bundle.Objects {
		var stixObject STIXObject
		json.Unmarshal(obj, &stixObject)
		switch stixObject.Type{
		// SDO
		case ThreatActorType:
			ta := unmarshalThreatActor(obj)
			printThreatActor(ta)
		case IdentityType:
			identity := unmarshalIdentity(obj)
			printIdentity(identity)
		case RelationshipType:
			relationship := unmarshalRelationship(obj)
			printRelationship(relationship)
		case AttackPatternType:
			ap := unmarshalAttackPattern(obj)
			printAttackPattern(ap)
		case CampaignType:
			campaign := unmarshalCampaign(obj)
			printCampaign(campaign)
		case IntrusionSetType:
			is := unmarshalIntrusionSet(obj)
			printIntrusionSet(is)
		case IndicatorType:
			indicator := unmarshalIndicator(obj)
			printIndicator(indicator)
		case MalwareType:
			malware := unmarshalMalware(obj)
			printMalware(malware)
		case SightingType:
			sighting := unmarshalSighting(obj)
			printSighting(sighting)
		case ObservedDataType:
			od := unmarshalObservedData(obj)
			printObservedData(od)
		case ToolType:
			tool := unmarshalTool(obj)
			printTool(tool)
		case ReportType:
			report := unmarshalReport(obj)
			printReport(report)
		case CourseOfActionType:
			coa := unmarshalCourseOfAction(obj)
			printCourseOfAction(coa)
		case VulnerabilityType:
			vul := unmarshalVulnerability(obj)
			printVulnerability(vul)
		case GroupingType:
			grouping := unmarshalGrouping(obj)
			printGrouping(grouping)
		case InfrastructureType:
			infra := unmarshalInfrastructure(obj)
			printInfrastructure(infra)
		case LocationType:
			loc := unmarshalLocation(obj)
			printLocation(loc)
		case MalwareAnalysisType:
			ma := unmarshalMalwareAnalysis(obj)
			printMalwareAnalysis(ma)
		case NoteType:
			note := unmarshalNote(obj)
			printNote(note)
		case OpinionType:
			opinion := unmarshalOpinion(obj)
			printOpinion(opinion)

		// SCO
		case ArtifactType:
			artifact := unmarshalArtifact(obj)
			printArtifact(artifact)
		case AutonomousSystemType:
			as := unmarshalAutonomousSystem(obj)
			printAutonomousSystem(as)
		case DirectoryType:
			dir := unmarshalDirectory(obj)
			printDirectory(dir)
		case DomainNameType:
			domain := unmarshalDomainName(obj)
			printDomainName(domain)
		case EmailAddressType:
			emailAddress := unmarshalEmailAddress(obj)
			printEmailAddress(emailAddress)
		case EmailMessageType:
			emailMessage := unmarshalEmailMessage(obj)
			printEmailMessage(emailMessage)
		case FileType:
			file := unmarshalFile(obj)
			printFile(file)
		case IPv4AddressType:
			ipv4 := unmarshalIpV4Address(obj)
			printIpV4Address(ipv4)
		case IPv6AddressType:
			ipv6 := unmarshalIpV6Address(obj)
			printIpV6Address(ipv6)
		case MACAddressType:
			mac := unmarshalMacAddress(obj)
			printMacAddress(mac)
		case MutexType:
			mutex := unmarshalMutex(obj)
			printMutex(mutex)
		case NetworkTrafficType:
			nt := unmarshalNetworkTraffic(obj)
			printNetworkTraffic(nt)
		case ProcessType:
			proc := unmarshalProcess(obj)
			printProcess(proc)
		case SoftwareType:
			sw := unmarshalSoftware(obj)
			printSoftware(sw)
		case URLType:
			url := unmarshalUrl(obj)
			printUrl(url)
		case UserAccountType:
			ua := unmarshalUserAccount(obj)
			printUserAccount(ua)

		// Markings
		case LanguageContentType:
			lc := unmarshalLanguageContent(obj)
			printLanguageContent(lc)
		case MarkingDefinitionType:
			md := unmarshalMarkingDefinition(obj)
			printMarkingDefinition(md)

		default:
			fmt.Printf("\n** Unknown object %v ***\n", stixObject.Type)
		}
	}
	return bundle
}

