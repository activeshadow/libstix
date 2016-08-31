// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"github.com/freestix/libstix/defs"
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type MarkingStructureType struct {
	MarkingModelName string `json:"marking_model_name,omitempty"`
	MarkingModelRef  string `json:"marking_model_ref,omitempty"`
	Id               string `json:"id,omitempty"`
	IdRef            string `json:"idref,omitempty"`
	Value            string `json:"value,omitempty"`
}

// ----------------------------------------------------------------------
// Methods MarkingStructureType
// ----------------------------------------------------------------------

func (this *MarkingStructureType) AddModelName(n string) {
	this.MarkingModelName = n
}

func (this *MarkingStructureType) AddModelRef(r string) {
	this.MarkingModelRef = r
}

func (this *MarkingStructureType) CreateId() {
	this.Id = defs.COMPANY + ":marking-" + uuid.New()
}

func (this *MarkingStructureType) AddIdRef(idref string) {
	this.IdRef = idref
}

func (this *MarkingStructureType) AddValue(s string) {
	this.Value = s
}
