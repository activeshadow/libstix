// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package indicator

import (
	"github.com/activeshadow/libstix/common"
	"github.com/activeshadow/libstix/defs"
	"time"
)

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

type SightingsType struct {
	SightingsCount int            `json:"sightings_count,omitempty"`
	Sighting       []SightingType `json:"sighting,omitempty"`
}

type SightingType struct {
	Timestamp          string                        `json:"timestamp,omitempty"`
	TimestampPrecision string                        `json:"timestamp_precision,omitempty"`
	Source             *common.InformationSourceType `json:"source,omitempty"`
	References         []string                      `json:"references,omitempty"`
	Confidence         *common.ConfidenceType        `json:"confidence,omitempty"`
	Description        *common.StructuredTextType    `json:"description,omitempty"`
	RelatedObservables *RelatedObservablesType       `json:"related_observables,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func CreateSighting() SightingType {
	var obj SightingType
	return obj
}

func CreateSightings() SightingsType {
	var obj SightingsType
	return obj
}

// ----------------------------------------------------------------------
// Methods SightingsType
// ----------------------------------------------------------------------

func (this *SightingsType) SetSightingsCount(count int) {
	this.SightingsCount = count
}

func (this *SightingsType) AddSighting(seen SightingType) {
	if this.Sighting == nil {
		a := make([]SightingType, 0)
		this.Sighting = a
	}
	this.Sighting = append(this.Sighting, seen)
}

// ----------------------------------------------------------------------
// Methods SightingType
// ----------------------------------------------------------------------

func (this *SightingType) CreateTimeStamp() {
	this.Timestamp = time.Now().Format(time.RFC3339)
}

func (this *SightingType) AddTimeStamp(t string) {
	// TODO Need to format the string in to ISO 8601 format or check that it is in the right format
	this.Timestamp = t
}

func (this *SightingType) AddTimeStampPrecision(p string) {
	this.TimestampPrecision = p
}

func (this *SightingType) AddSource(source common.InformationSourceType) {
	this.Source = &source
}

func (this *SightingType) AddReference(r string) {
	if this.References == nil {
		a := make([]string, 0)
		this.References = a
	}
	this.References = append(this.References, r)
}

func (this *SightingType) AddConfidence(c common.ConfidenceType) {
	this.Confidence = &c
}

func (this *SightingType) AddDescriptionText(value string) {
	this.AddDescription(defs.RFC6838TEXTPLAIN, value)
}

func (this *SightingType) AddDescription(format, value string) {
	var data common.StructuredTextType
	data.AddFormat(format)
	data.AddValue(value)

	this.Description = &data
}

func (this *SightingType) AddRelatedObservables(r RelatedObservablesType) {
	this.RelatedObservables = &r
}
