// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/freestix/libstix/stix"
	"github.com/freestix/libstix/ttp"
)

func main() {

	s := stix.New()
	i1 := s.NewIndicator()

	i1.SetTimestampToNow()
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

	var data []byte
	data, _ = json.MarshalIndent(s, "", "    ")

	fmt.Println(string(data))
}
