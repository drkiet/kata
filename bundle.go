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

		// Markings
		case MarkingDefinitionType:
			md := unmarshalMarkingDefinition(obj)
			printMarkingDefinition(md)

		default:
			fmt.Printf("\n** Unknown object %v ***\n", stixObject.Type)
		}
	}
	return bundle
}

