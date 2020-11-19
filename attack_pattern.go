package main

// External References
// Kill Chain Phase
type AttackPattern struct {
	Type string
	ExternalReferences [] ExternalReference
	Name string
	Description string
	Aliases [] string
	KillChainPhases [] KillChainPhase
}
