// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"github.com/activeshadow/libstix"
	"github.com/activeshadow/libstix/defs"
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type MarkingSpecificationType struct {
	Id                  string                 `json:"id,omitempty"`
	IdRef               string                 `json:"idref,omitempty"`
	Version             string                 `json:"version,omitempty"`
	ControlledStructure string                 `json:"controlled_structure,omitempty"`
	MarkingStructure    *MarkingStructureType  `json:"marking_structure,omitempty"`
	InformationSource   *InformationSourceType `json:"information_source,omitempty"`
}

// ----------------------------------------------------------------------
// Methods MarkingSpecificationType
// ----------------------------------------------------------------------

func (this *MarkingSpecificationType) CreateId() {
	this.Id = libstix.Company + ":marking-" + uuid.New()
}

func (this *MarkingSpecificationType) AddIdRef(idref string) {
	this.IdRef = idref
}

// Data_Marking schema version for this content.
func (this *MarkingSpecificationType) AddVersion(ver string) {
	this.Version = ver
}

func (this *MarkingSpecificationType) AddControlledStructure(s string) {
	this.ControlledStructure = s
}

func (this *MarkingSpecificationType) AddMarkingStructure(mark MarkingStructureType) {
	this.MarkingStructure = &mark
}

func (this *MarkingSpecificationType) AddTLPMarking(value string) {
	data := MarkingStructureType{
		MarkingModelName: defs.MARKING_TLP_VOCAB,
		Value:            value,
	}
	this.MarkingStructure = &data
}

func (this *MarkingSpecificationType) AddInformationSource(source InformationSourceType) {
	this.InformationSource = &source
}
