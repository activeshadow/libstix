package main

import (
	"encoding/json"
	"fmt"
	"github.com/jordan2175/freestix/libstix/indicator"
	"github.com/jordan2175/freestix/libstix/stix"
	"github.com/jordan2175/freestix/libstix/ttp"
)

func main() {
	s := stix.Create()
	i1 := indicator.Create()
	i1.CreateId()
	i1.CreateTimeStamp()
	i1.AddTitle("Attack 2015-02")

	// Create TTP and Cyber Kill Chain Definitions
	t := ttp.Create()
	k := stix.CreateKillChain()
	k.CreateId()
	k.AddDefiner("LMCO")
	k.AddName("LM Cyber Kill Chain")
	k.AddNumberOfPhases(7)
	k.AddReference("http://www.lockheedmartin.com/us/what-we-do/information-technology/cyber-security/cyber-kill-chain.html")
	k.AddPhase(1, "Reconnaissance")
	k.AddPhase(2, "Weaponization")
	k.AddPhase(3, "Delivery")
	k.AddPhase(4, "Exploitation")
	k.AddPhase(5, "Installation")
	k.AddPhase(6, "Command and Control")
	k.AddPhase(7, "Actions on Objectives")
	t.AddKillChains(k)

	// Get ID values from kill chain
	chainId := k.Id
	phase3Id := k.KillChainPhase[2].PhaseId

	i1.AddKillChainPhaseAndChain(phase3Id, chainId)

	s.AddTTPs(t)
	s.AddIndicator(i1)

	var data []byte
	data, _ = json.MarshalIndent(s, "", "    ")

	fmt.Println(string(data))
}
