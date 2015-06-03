// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

// Version: 0.2

package common

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/jordan2175/freestix/libstix/defs"
)

type MarkingSpecificationType struct {
	Id                  string                 `json:"id,omitempty"`
	IdRef               string                 `json:"idref,omitempty"`
	Version             string                 `json:"version,omitempty"`
	ControlledStructure string                 `json:"controlledStructure,omitempty"`
	MarkingStructure    *MarkingStructureType  `json:"markingStructure,omitempty"`
	InformationSource   *InformationSourceType `json:"informationSource,omitempty"`
}

// ----------------------------------------------------------------------
// Methods MarkingSpecificationType
// ----------------------------------------------------------------------

func (m *MarkingSpecificationType) CreateId() {
	m.Id = defs.COMPANY + ":marking-" + uuid.New()
}

func (m *MarkingSpecificationType) AddIdRef(idref string) {
	m.IdRef = idref
}

// Data_Marking schema version for this content.
func (m *MarkingSpecificationType) AddVersion(ver string) {
	m.Version = ver
}

func (m *MarkingSpecificationType) AddControlledStructure(s string) {
	m.ControlledStructure = s
}

func (m *MarkingSpecificationType) AddMarkingStructure(mark MarkingStructureType) {
	m.MarkingStructure = &mark
}

func (m *MarkingSpecificationType) AddTLPMarking(value string) {
	data := MarkingStructureType{
		MarkingModelName: defs.MARKING_TLP_VOCAB,
		Value:            value,
	}
	m.MarkingStructure = &data
}

func (m *MarkingSpecificationType) AddInformationSource(source InformationSourceType) {
	m.InformationSource = &source
}
