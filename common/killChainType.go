// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.1

package common

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/freestix/libstix/defs"
)

type KillChainType struct {
	Id             string               `json:"id,omitempty"`
	Name           string               `json:"name,omitempty"`
	Definer        string               `json:"definer,omitempty"`
	Reference      string               `json:"reference,omitempty"`
	NumberOfPhases int                  `json:"numberOfPhases,omitempty"`
	KillChainPhase []KillChainPhaseType `json:"killChainPhase,omitempty"`
}

type KillChainPhaseType struct {
	Ordinality int    `json:"ordinality,omitempty"`
	Name       string `json:"name,omitempty"`
	PhaseId    string `json:"phaseId,omitempty"`
}

type KillChainPhaseReferenceType struct {
	PhaseId       string `json:"phaseId,omitempty"`
	Name          string `json:"name,omitempty"`
	Ordinality    int    `json:"ordinality,omitempty"`
	KillChainId   string `json:"killChainId,omitempty"`
	KillChainName string `json:"killChainName,omitempty"`
}

// ----------------------------------------------------------------------
// Methods KillChainType
// ----------------------------------------------------------------------

func (k *KillChainType) CreateId() {
	k.Id = defs.COMPANY + ":TTP-" + uuid.New()
}

func (k *KillChainType) AddName(n string) {
	k.Name = n
}

func (k *KillChainType) AddDefiner(d string) {
	k.Definer = d
}

func (k *KillChainType) AddReference(r string) {
	k.Reference = r
}

func (k *KillChainType) AddNumberOfPhases(n int) {
	k.NumberOfPhases = n
}

func (k *KillChainType) AddPhase(n int, name string) {
	if k.KillChainPhase == nil {
		a := make([]KillChainPhaseType, 0)
		k.KillChainPhase = a
	}
	phase := defs.COMPANY + ":TTP-" + uuid.New()
	data := KillChainPhaseType{
		Ordinality: n,
		Name:       name,
		PhaseId:    phase,
	}
	k.KillChainPhase = append(k.KillChainPhase, data)
}
