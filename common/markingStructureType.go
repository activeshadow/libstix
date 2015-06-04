// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.2

package common

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/freestix/libstix/defs"
)

type MarkingStructureType struct {
	MarkingModelName string `json:"markingModelName,omitempty"`
	MarkingModelRef  string `json:"markingModelRef,omitempty"`
	Id               string `json:"id,omitempty"`
	IdRef            string `json:"idref,omitempty"`
	Value            string `json:"value,omitempty"`
}

// ----------------------------------------------------------------------
// Methods MarkingStructureType
// ----------------------------------------------------------------------

func (m *MarkingStructureType) AddModelName(n string) {
	m.MarkingModelName = n
}

func (m *MarkingStructureType) AddModelRef(r string) {
	m.MarkingModelRef = r
}

func (m *MarkingStructureType) CreateId() {
	m.Id = defs.COMPANY + ":marking-" + uuid.New()
}

func (m *MarkingStructureType) AddIdRef(idref string) {
	m.IdRef = idref
}

func (m *MarkingStructureType) AddValue(s string) {
	m.Value = s
}
