// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"github.com/freestix/libstix"
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type KillChainType struct {
	Id             string               `json:"id,omitempty"`
	Name           string               `json:"name,omitempty"`
	Definer        string               `json:"definer,omitempty"`
	Reference      string               `json:"reference,omitempty"`
	NumberOfPhases int                  `json:"number_of_phases,omitempty"`
	KillChainPhase []KillChainPhaseType `json:"kill_chain_phase,omitempty"`
}

type KillChainPhaseType struct {
	Ordinality int    `json:"ordinality,omitempty"`
	Name       string `json:"name,omitempty"`
	PhaseId    string `json:"phase_id,omitempty"`
}

type KillChainPhaseReferenceType struct {
	PhaseId       string `json:"phase_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Ordinality    int    `json:"ordinality,omitempty"`
	KillChainId   string `json:"kill_chain_id,omitempty"`
	KillChainName string `json:"kill_chain_name,omitempty"`
}

// ----------------------------------------------------------------------
// Methods KillChainType
// ----------------------------------------------------------------------

func (this *KillChainType) CreateId() {
	this.Id = libstix.Company + ":TTP-" + uuid.New()
}

func (this *KillChainType) AddName(n string) {
	this.Name = n
}

func (this *KillChainType) AddDefiner(d string) {
	this.Definer = d
}

func (this *KillChainType) AddReference(r string) {
	this.Reference = r
}

func (this *KillChainType) AddNumberOfPhases(n int) {
	this.NumberOfPhases = n
}

func (this *KillChainType) AddPhase(n int, name string) {
	if this.KillChainPhase == nil {
		a := make([]KillChainPhaseType, 0)
		this.KillChainPhase = a
	}
	phase := libstix.Company + ":TTP-" + uuid.New()
	data := KillChainPhaseType{
		Ordinality: n,
		Name:       name,
		PhaseId:    phase,
	}
	this.KillChainPhase = append(this.KillChainPhase, data)
}
