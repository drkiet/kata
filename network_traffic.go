package kata

import (
	"encoding/json"
	"fmt"
	"time"
)

type NetworkTraffic struct {
	Type string `json:"type" binding:"required"`
	SpecVersion string `json:"spec_version,omitempty"`
	ID string `json:"id" binding:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings [] GranularMarking `json:"granular-markings,omitempty"`
	Defanged bool `json:"defanged,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`

	Start time.Time `json:"start,omitempty"`
	End time.Time `json:"end,omitempty"`
	IsActive bool `json:"is_active,omitempty"`
	SrcRef string `json:"src_ref,omitempty"`
	DstRef string `json:"dst_ref,omitempty"`
	SrcPort int `json:"src_port,omitempty"`
	DstPort int `json:"dst_port,omitempty"`
	Protocols []string `json:"protocols" binding:"required"`
	SrcByteCount int `json:"src_byte_count,omitempty"`
	DstByteCount int `json:"dst_byte_count,omitempty"`
	SrcPackets int `json:"src_packets,omitempty"`
	DstPackets int `json:"dst_packets,omitempty"`
	IpFix map[string]interface{} `json:"ipfix,omitempty"`
	SrcPayloadRef string `json:"src_payload_ref,omitempty"`
	DstPayloadRef string `json:"dst_payload_ref,omitempty"`
	EncapsulatesRefs []string `json:"encapsulates_refs,omitempty"`
	EncapsulatedByRef string `json:"encapsulated_by_ref,omitempty"`
}

func printNetworkTraffic(networkTraffic NetworkTraffic) {
	fmt.Println("Network Traffic:\n", marshalNetworkTraffic(networkTraffic))
}

func marshalNetworkTraffic(networkTraffic NetworkTraffic) (jsonData string){
	data, e := json.MarshalIndent(networkTraffic, "", "  ")
	check(e)
	return string(data)
}

func unmarshalNetworkTraffic(obj json.RawMessage) (networkTraffic NetworkTraffic) {
	json.Unmarshal(obj, &networkTraffic)
	return networkTraffic
}




