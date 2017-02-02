// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package ttp

import (
	"github.com/activeshadow/libstix/common"
)

type TTPsType struct {
	TTP        *TTPType               `json:"ttp,omitempty"`
	KillChains []common.KillChainType `json:"kill_chains,omitempty"`
}

type TTPType struct {
	Id               string                     `json:"id,omitempty"`
	IdRef            string                     `json:"idref,omitempty"`
	Timestamp        string                     `json:"timestamp,omitempty"`
	Version          string                     `json:"version,omitempty"`
	Title            string                     `json:"title,omitempty"`
	Description      *common.StructuredTextType `json:"description,omitempty"`
	ShortDescription *common.StructuredTextType `json:"short_description,omitempty"`
}

// ----------------------------------------------------------------------
// Methods TTPsType
// ----------------------------------------------------------------------

func (t *TTPsType) AddKillChains(k common.KillChainType) {
	if t.KillChains == nil {
		a := make([]common.KillChainType, 0)
		t.KillChains = a
	}
	t.KillChains = append(t.KillChains, k)
}
