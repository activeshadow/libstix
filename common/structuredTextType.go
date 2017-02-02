// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package common

import (
	"github.com/activeshadow/libstix"
	"github.com/pborman/uuid"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type StructuredTextType struct {
	IdBaseType
	MarkingIdRefBaseType
	Format string `json:"format,omitempty"`
	Value  string `json:"value,omitempty"`
}

// ----------------------------------------------------------------------
// Methods StructuredTextType
// ----------------------------------------------------------------------

func (this *StructuredTextType) CreateId() {
	this.Id = libstix.Company + ":text-" + uuid.New()
}

func (this *StructuredTextType) AddMarkingIdRef(markingidref string) {
	this.MarkingIdRef = markingidref
}

func (this *StructuredTextType) AddFormat(format string) {
	this.Format = format
}

func (this *StructuredTextType) AddValue(value string) {
	this.Value = value
}
